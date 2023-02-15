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
)

// / fetchLogs fetches logs from ethereum and sends them to the logsch channel.
// It fetches logs in batches of size `step`, starting from the `startBlock` and moving forward.
// The `contractAddress` parameter specifies the contract address to filter logs.
// The `maxRetries` parameter specifies the maximum number of retries if an error occurs.
// The `concurrency` parameter specifies the maximum number of concurrent requests to make.
func fetchLogs(ctx context.Context, ethclient *ethclient.Client, db *pgxpool.Pool, contractAddress string, startBlock, step, stepFactor, maxRetries uint64, logsch chan<- types.Log, waitTime time.Duration) {
	if stepFactor <= 0 {
		log.Fatalf("stepFactor must be greater than 0")
		return
	}

	// Fetch last block number from db, if not found, use startBlock
	var lastBlock uint64
	err := db.QueryRow(ctx, `SELECT "blockNumber" FROM "LastBlock"`).Scan(&lastBlock)
	if err != nil {
		log.Printf("Error fetching last block number, using start block: %d", startBlock)
	} else {
		startBlock = lastBlock + 1
	}

	// history fetcher loop
	for {
		// if we reached the latest block, break out of the loop
		latestDBBlock := startBlock + step*stepFactor - 1
		latestChainBlock, err := ethclient.BlockNumber(ctx)
		if err != nil {
			log.Fatalf("Error fetching latest block: %s", err)
			return
		}
		if latestDBBlock >= latestChainBlock {
			log.Printf("Reached latest block %d (latest chain block %d), exiting history fetcher", latestDBBlock, latestChainBlock)
			break
		}

		// Fetch logs in batches of `step` blocks
		var wg sync.WaitGroup
		for i := uint64(0); i < stepFactor; i++ {
			wg.Add(1)
			go func(i uint64) {
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
				log.Printf("History: Fetched %d logs from block %d to %d", len(logsSlice), startIndex, endBlock)
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
		updateLastBlock(ctx, db, latestDBBlock)

		startBlock += step * stepFactor
	}

	// latest fetcher loop
	var latestDBBlock uint64
	if err = db.QueryRow(ctx, `SELECT "blockNumber" FROM "LastBlock"`).Scan(&latestDBBlock); err != nil {
		log.Fatal(err)
	}
	startBlock = latestDBBlock + 1
	for {
		latestChainBlock, err := ethclient.BlockNumber(ctx)
		if err != nil {
			log.Fatalf("Error fetching latest block: %s", err)
			return
		}

		if startBlock < latestChainBlock {
			endBlock := startBlock + step
			if endBlock > latestChainBlock {
				endBlock = latestChainBlock
			}

			logsSlice, err := fetchLogsInRange(ctx, ethclient, contractAddress, startBlock, endBlock, maxRetries)
			if err != nil {
				log.Fatalf("Error fetching logs from block %d to %d: %s", startBlock, endBlock, err)
				return
			}
			log.Printf("Polling: Fetched %d logs from block %d to %d", len(logsSlice), startBlock, endBlock)
			for _, l := range logsSlice {
				select {
				case logsch <- l:
				case <-ctx.Done():
					log.Println("Context cancelled. Exiting fetchLogs.")
					return
				}
			}

			// Upsert last block number
			updateLastBlock(ctx, db, endBlock)

			startBlock = endBlock + 1
		} else {
			log.Printf("Polling latest block %d (latest chain block %d), sleeping for %s", startBlock, latestChainBlock, waitTime)
			time.Sleep(waitTime)
		}
	}
}

// fetchLogsInRange fetches logs in the given block range and returns them in a slice.
// The `contractAddress` parameter specifies the contract address to filter logs.
// The `startBlock` and `endBlock` parameters specify the range of blocks to fetch logs from.
// The `maxRetries` parameter specifies the maximum number of retries if an error occurs.
func fetchLogsInRange(ctx context.Context, ethclient *ethclient.Client, contractAddress string, startBlock, endBlock, maxRetries uint64) ([]types.Log, error) {
	var logs []types.Log
	var err error

	for i := uint64(0); i < maxRetries; i++ {
		logs, err = ethclient.FilterLogs(ctx, ethereum.FilterQuery{
			FromBlock: big.NewInt(int64(startBlock)),
			ToBlock:   big.NewInt(int64(endBlock)),
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

func updateLastBlock(ctx context.Context, db *pgxpool.Pool, number uint64) {
	// Upsert last block number
	_, err := db.Exec(ctx, `INSERT INTO "LastBlock" ("id", "blockNumber") VALUES (1, $1) ON CONFLICT ("id") DO UPDATE SET "blockNumber" = $1`, number)
	if err != nil {
		log.Fatalf("Error upserting last block number: %s", err)
	}
}
