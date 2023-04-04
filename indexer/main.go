package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spf13/viper"
)

var ec *ethclient.Client

var (
	flagConfigFile = flag.String("config", "config.toml", "config file")
)

type Config struct {
	ContractAddress string `mapstructure:"contract_address"`
	RpcURL          string `mapstructure:"rpc_url"`
	PGDSN           string `mapstructure:"pg_dsn"`
	StartBlock      uint64 `mapstructure:"start_block"`
	Step            uint64 `mapstructure:"step"`
	StepFactor      uint64 `mapstructure:"step_factor"`
	MaxRetries      uint64 `mapstructure:"max_retries"`
	PollingInterval uint64 `mapstructure:"polling_interval"`
}

func main() {
	flag.Parse()

	conf := loadConfig(*flagConfigFile)

	db := initDB(conf.PGDSN)

	ethclient := initEthClient(conf.RpcURL)

	// fetch raw logs from ethereum
	logsch := make(chan []types.Log)
	ctx := context.Background()
	go fetchLogs(ctx, ethclient, db, conf.ContractAddress, conf.StartBlock, conf.Step, conf.StepFactor, conf.MaxRetries, logsch, time.Duration(conf.PollingInterval)*time.Second)

	// process logs and save to db
	processLogs(db, logsch)

}

func loadConfig(confPath string) Config {
	// load config from toml file
	viper.SetConfigFile(confPath)

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

	ec = ethclient

	return ethclient
}

// processLogs parse logs and pass to event channel
// log will be parsed to related event struct
func processLogs(db *pgxpool.Pool, logsch chan []types.Log) {
	for logs := range logsch {
		now := time.Now()
		if err := ProcessEvents(db, logs); err != nil {
			log.Printf("Error processing events: %v", err)
		}
		log.Printf("Processed %d events in %s", len(logs), time.Since(now).String())
	}
}
