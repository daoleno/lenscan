package main

import (
	"context"
	"database/sql"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	_ "github.com/lib/pq"
)

const (
	profileCreatedEventTopic       = "0xf642d82f9bf073e3403d88853e8ee1a91d4fff05e11bcdf593f09ce442c6b247"
	forcedStartEnv                 = "FORCED_START_BLOCK"
	stepSize                       = 1000
	permissionlessCreatorInitBlock = 53476993
)

func main() {
	pgDsn := os.Getenv("DATABASE_URL")
	rpcUrl := os.Getenv("RPC_URL")
	forcedStartBlock := os.Getenv(forcedStartEnv)

	log.Println("Starting the event listener...")

	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		log.Fatalf("Failed to connect to the client: %v", err)
	}
	log.Println("Connected to RPC node.")

	db, err := sql.Open("postgres", pgDsn)
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL database: %v", err)
	}
	defer db.Close()
	log.Println("Connected to PostgreSQL database.")

	var startBlockNumber *big.Int
	if forcedStartBlock != "" {
		startBlockNumber, _ = new(big.Int).SetString(forcedStartBlock, 10)
		log.Printf("Forced start block specified: %s", startBlockNumber.String())
	} else {
		startBlockNumber = recoverLastIndexedHeight(db)
		log.Printf("Recovered start block: %s", startBlockNumber.String())
	}

	for {
		latestBlock, err := client.BlockByNumber(context.Background(), nil)
		if err != nil {
			log.Fatalf("Failed to fetch the latest block: %v", err)
		}

		endBlockNumber := new(big.Int).Add(startBlockNumber, big.NewInt(stepSize-1))
		if endBlockNumber.Cmp(latestBlock.Number()) > 0 {
			endBlockNumber.Set(latestBlock.Number())
		}

		log.Printf("Processing blocks from %s to %s", startBlockNumber.String(), endBlockNumber.String())

		query := ethereum.FilterQuery{
			FromBlock: startBlockNumber,
			ToBlock:   endBlockNumber,
			Topics:    [][]common.Hash{{common.HexToHash(profileCreatedEventTopic)}},
		}

		logs, err := client.FilterLogs(context.Background(), query)
		if err != nil {
			log.Fatalf("Failed to filter logs: %v", err)
		}

		log.Printf("Found %d logs in the block range.", len(logs))

		processLogs(logs, client, db)

		if endBlockNumber.Cmp(latestBlock.Number()) >= 0 {
			log.Println("Reached the latest block. Waiting for new blocks...")
			time.Sleep(10 * time.Second)
		} else {
			startBlockNumber.Add(endBlockNumber, big.NewInt(1))
		}
	}
}

func processLogs(logs []types.Log, client *ethclient.Client, db *sql.DB) {
	for _, vLog := range logs {
		log.Printf("Processing log: BlockNumber=%d, TxHash=%s", vLog.BlockNumber, vLog.TxHash.Hex())

		// Get the transaction by its hash
		tx, isPending, err := client.TransactionByHash(context.Background(), vLog.TxHash)
		if err != nil {
			log.Fatalf("Failed to fetch transaction: %v", err)
		}
		if isPending {
			log.Printf("Transaction %s is still pending, skipping", vLog.TxHash.Hex())
			continue
		}

		// https://github.com/ethereum/go-ethereum/issues/22918#issuecomment-1378157861
		from, err := types.Sender(types.LatestSignerForChainID(tx.ChainId()), tx)
		if err != nil {
			log.Fatalf("Failed to get sender address: %v", err)
		}

		// Get the recipient address of the transaction (if it's a contract creation transaction, the recipient address may be nil)
		var toAddress common.Address
		if tx.To() != nil {
			toAddress = *tx.To()
		}

		// Get the timestamp of the block where the transaction is included
		block, err := client.BlockByNumber(context.Background(), big.NewInt(int64(vLog.BlockNumber)))
		if err != nil {
			log.Fatalf("Failed to fetch block: %v", err)
		}
		blockTimestamp := time.Unix(int64(block.Time()), 0)

		// Insert the transaction information into the database
		_, err = db.Exec("INSERT INTO lens_profile_created_transaction (block_number, block_timestamp, transaction_hash, from_address, to_address, value) VALUES ($1, $2, $3, $4, $5, $6) ON CONFLICT (transaction_hash) DO NOTHING",
			vLog.BlockNumber, blockTimestamp, vLog.TxHash.Hex(), from.Hex(), toAddress.Hex(), tx.Value().String())
		if err != nil {
			log.Fatalf("Failed to insert data into database: %v", err)
		}

		log.Printf("Successfully processed and stored transaction: %s", vLog.TxHash.Hex())
	}
}

func recoverLastIndexedHeight(db *sql.DB) *big.Int {
	var lastBlock sql.NullInt64
	err := db.QueryRow("SELECT MAX(block_number) FROM lens_profile_created_transaction").Scan(&lastBlock)
	if err != nil {
		log.Fatalf("Failed to query the latest block number from database: %v", err)
	}
	if lastBlock.Valid {
		return big.NewInt(lastBlock.Int64 + 1)
	}
	return big.NewInt(permissionlessCreatorInitBlock)
}
