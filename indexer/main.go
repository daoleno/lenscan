package main

import (
	"context"
	"log"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spf13/viper"
)

type Config struct {
	ContractAddress string `mapstructure:"contract_address"`
	RpcURL          string `mapstructure:"rpc_url"`
	PGDSN           string `mapstructure:"pg_dsn"`
	StartBlock      int64  `mapstructure:"start_block"`
	Step            int64  `mapstructure:"step"`
	StepFactor      int64  `mapstructure:"step_factor"`
	MaxRetries      int64  `mapstructure:"max_retries"`
}

func main() {
	conf := loadConfig()

	db := initDB(conf.PGDSN)

	ethclient := initEthClient(conf.RpcURL)

	// fetch raw logs from ethereum
	logsch := make(chan types.Log)
	ctx := context.Background()
	go fetchLogs(ctx, ethclient, db, conf.ContractAddress, conf.StartBlock, conf.Step, conf.StepFactor, conf.MaxRetries, logsch)

	// process logs and save to db
	processLogs(db, logsch)

}

func loadConfig() Config {
	// load config from toml file
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	var conf Config
	err = viper.Unmarshal(&conf)
	if err != nil {
		panic(err)
	}

	return conf
}

func initDB(dsn string) *pgxpool.Pool {
	dbpool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		panic(err)
	}

	return dbpool
}

func initEthClient(rpcURL string) *ethclient.Client {
	ethclient, err := ethclient.Dial(rpcURL)
	if err != nil {
		panic(err)
	}

	return ethclient
}

// / fetchLogs fetches logs from ethereum and sends them to the logsch channel.
// It fetches logs in batches of size `step`, starting from the `startBlock` and moving forward.
// The `contractAddress` parameter specifies the contract address to filter logs.
// The `maxRetries` parameter specifies the maximum number of retries if an error occurs.
// The `concurrency` parameter specifies the maximum number of concurrent requests to make.
func fetchLogs(ctx context.Context, ethclient *ethclient.Client, db *pgxpool.Pool, contractAddress string, startBlock, step, stepFactor, maxRetries int64, logsch chan<- types.Log) {
	if stepFactor <= 0 {
		log.Fatalf("stepFactor must be greater than 0")
		return
	}

	// Fetch last block number from db, if not found, use startBlock
	var lastBlock int64
	err := db.QueryRow(ctx, `SELECT "blockNumber" FROM "LastBlock"`).Scan(&lastBlock)
	if err != nil {
		log.Printf("Error fetching last block number, using start block: %d", startBlock)
	} else {
		startBlock = lastBlock + 1
	}

	for {
		// Fetch logs in batches of `step` blocks
		var wg sync.WaitGroup
		for i := int64(0); i < stepFactor; i++ {
			wg.Add(1)
			go func(i int64) {
				defer wg.Done()
				endBlock := startBlock + step*(i+1) - 1
				if i == stepFactor-1 {
					endBlock = startBlock + step*stepFactor - 1
				}
				startIndex := startBlock + step*i
				logsSlice, err := fetchLogsInRange(ctx, ethclient, contractAddress, startIndex, endBlock, maxRetries)
				if err != nil {
					log.Fatalf("Error fetching logs from block %d to %d: %s", startIndex, endBlock, err)
					return
				}
				log.Printf("Fetched %d logs from block %d to %d", len(logsSlice), startIndex, endBlock)
				for _, l := range logsSlice {
					select {
					case logsch <- l:
					case <-ctx.Done():
						log.Println("Context cancelled. Exiting fetchLogs.")
						return
					}
				}
			}(i)
		}
		wg.Wait()

		// Upsert last block number
		updateLastBlock(ctx, db, startBlock+step*stepFactor-1)

		startBlock += step * stepFactor
	}
}

// fetchLogsInRange fetches logs in the given block range and returns them in a slice.
// The `contractAddress` parameter specifies the contract address to filter logs.
// The `startBlock` and `endBlock` parameters specify the range of blocks to fetch logs from.
// The `maxRetries` parameter specifies the maximum number of retries if an error occurs.
func fetchLogsInRange(ctx context.Context, ethclient *ethclient.Client, contractAddress string, startBlock, endBlock, maxRetries int64) ([]types.Log, error) {
	var logs []types.Log
	var err error

	for i := int64(0); i < maxRetries; i++ {
		logs, err = ethclient.FilterLogs(ctx, ethereum.FilterQuery{
			FromBlock: big.NewInt(startBlock),
			ToBlock:   big.NewInt(endBlock),
			Addresses: []common.Address{common.HexToAddress(contractAddress)},
		})
		if err == nil {
			return logs, nil
		}
		// Log the error and wait for a while before retrying
		log.Printf("Error fetching logs from block %d to %d: %s, retrying in 5 seconds...", startBlock, endBlock, err)
		time.Sleep(time.Second * 5)
	}

	return nil, err
}

func updateLastBlock(ctx context.Context, db *pgxpool.Pool, number int64) {
	// Upsert last block number
	_, err := db.Exec(ctx, `INSERT INTO "LastBlock" ("id", "blockNumber") VALUES (1, $1) ON CONFLICT ("id") DO UPDATE SET "blockNumber" = $1`, number)
	if err != nil {
		log.Fatalf("Error upserting last block number: %s", err)
	}
}

// processLogs parse logs and pass to event channel
// log will be parsed to related event struct
func processLogs(db *pgxpool.Pool, logsch chan types.Log) {
	for l := range logsch {
		p := EventProcessors[l.Topics[0]]
		if p == nil {
			log.Printf("No processor found for event with topic %s at transaction %s", l.Topics[0].String(), l.TxHash.String())
			continue
		}

		err := p.ProcessEvent(db, l)
		if err != nil {
			log.Printf("Error parsing event: %s", err)
			continue
		}
	}
}
