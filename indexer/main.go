package main

import (
	"context"
	"log"
	"math/big"

	"github.com/daoleno/lenscan/indexer/contract"
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
	StartBlock      int64  `mapstructure:"start_block"`
	PGDSN           string `mapstructure:"pg_dsn"`
	Step            int64  `mapstructure:"step"`
}

func main() {
	conf := loadConfig()

	db := initDB(conf.PGDSN)

	ethclient := initEthClient(conf.RpcURL)

	// fetch raw logs from ethereum
	logsch := make(chan types.Log)
	go fetchLogs(ethclient, conf.ContractAddress, conf.StartBlock, conf.Step, logsch)

	// parse event logs
	eventch := make(chan interface{})
	go parseLogs(logsch, eventch)

	// save to database
	go saveToDB(db, eventch)

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

func initEventFilter(contractAddress string, ethclient *ethclient.Client) *contract.ContractFilterer {
	f, err := contract.NewContractFilterer(common.HexToAddress(contractAddress), ethclient)
	if err != nil {
		panic(err)
	}

	return f
}

// fetchLogs fetch logs from ethereum and pass to logs channel
// step is the number of blocks to fetch at a time
func fetchLogs(ethclient *ethclient.Client, contractAddress string, startBlock, step int64, logsch chan types.Log) {
	for {
		// Fetch logs in batches of `step` blocks
		endBlock := startBlock + step
		logsSlice, err := ethclient.FilterLogs(
			context.Background(),
			ethereum.FilterQuery{
				FromBlock: big.NewInt(startBlock),
				ToBlock:   big.NewInt(endBlock),
				Addresses: []common.Address{common.HexToAddress(contractAddress)},
			})
		if err != nil {
			log.Fatalf("Error fetching logs from block %d to %d: %s", startBlock, endBlock, err)
		}

		// If no logs are found, move on to the next batch
		if len(logsSlice) == 0 {
			startBlock = endBlock + 1
			continue
		}

		for _, log := range logsSlice {
			logsch <- log
		}
		startBlock = endBlock + 1
	}
}

// parseLogs parse logs and pass to event channel
// log will be parsed to related event struct
func parseLogs(logsch chan types.Log, eventch chan interface{}) {
	for l := range logsch {
		p := EventParsers[l.Topics[0]]
		if p == nil {
			log.Printf("No parser found for event with topic %s", l.Topics[0].Hex())
			continue
		}

		event, err := p(l)
		if err != nil {
			log.Printf("Error parsing event: %s", err)
			continue
		}

		eventch <- event
	}
}

// saveToDB save event to database
func saveToDB(db *pgxpool.Pool, eventch chan interface{}) {
	ctx := context.Background()
	for event := range eventch {
		switch e := event.(type) {
		case *contract.ContractBaseInitialized:
			// Save the ContractBaseInitialized event to the database
			_, err := db.Exec(ctx, `
				WITH inserted_event AS (
					INSERT INTO Event (blockNumber, txHash, txIndex, logIndex, removed)
					VALUES ($1, $2, $3, $4, $5)
					RETURNING id
				)
				INSERT INTO BaseInitialized (event_id, name, symbol, timestamp)
				SELECT id, $6, $7, $8
				FROM inserted_event`,
				e.Raw.BlockNumber, e.Raw.TxHash, e.Raw.TxIndex, e.Raw.Index, e.Raw.Removed, e.Name, e.Symbol, e.Timestamp)
			if err != nil {
				log.Println("Error saving BaseInitialized event to database:", err)
			}

		case *contract.ContractCollectModuleWhitelisted:
			// Save the ContractCollectModuleWhitelisted event to the database
			_, err := db.Exec(ctx, `
				WITH inserted_event AS (
					INSERT INTO Event (blockNumber, txHash, txIndex, logIndex, removed)
					VALUES ($1, $2, $3, $4, $5)
					RETURNING id
				)
				INSERT INTO CollectModuleWhitelisted (event_id, collectModule, whitelisted, timestamp)
				SELECT id, $6, $7, $8
				FROM inserted_event`,
				e.Raw.BlockNumber, e.Raw.TxHash, e.Raw.TxIndex, e.Raw.Index, e.Raw.Removed, e.CollectModule, e.Whitelisted, e.Timestamp)
			if err != nil {
				log.Println("Error saving CollectModuleWhitelisted event to database:", err)
			}
		}
	}
}
