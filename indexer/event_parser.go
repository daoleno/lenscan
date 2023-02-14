package main

import (
	"context"
	"fmt"

	"github.com/daoleno/lenscan/indexer/contract"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/jackc/pgx/v5/pgxpool"
)

var contractFilter, _ = contract.NewContractFilterer(common.Address{}, nil)

// EventProcessor includes the logic to parse and write an event to the database
type EventProcessor struct {
	parser EventParser
	writer EventWriter
}

// EventParser represents a function that parses an event
type EventParser func(types.Log) (interface{}, error)
type EventWriter func(*pgxpool.Pool, interface{}) error

// NewEventProcessor creates a new EventProcessor
func NewEventProcessor(parser EventParser, writer EventWriter) *EventProcessor {
	return &EventProcessor{
		parser: parser,
		writer: writer,
	}
}

// ProcessEvent parses and writes an event to the database
func (p *EventProcessor) ProcessEvent(db *pgxpool.Pool, l types.Log) error {
	// Parse the event
	e, err := p.parser(l)
	if err != nil {
		return err
	}

	// Write the event to the database
	return p.writer(db, e)
}

func (p *EventProcessor) Parse(l types.Log) (interface{}, error) {
	return p.parser(l)
}

func (p *EventProcessor) Write(db *pgxpool.Pool, e interface{}) error {
	return p.writer(db, e)
}

// EventProcessors is a map of event signatures to event parsers
var EventProcessors = map[common.Hash]*EventProcessor{
	// first example
	common.HexToHash("0x414cd0b34676984f09a5f76ce9718d4062e50283abe0e7e274a9a5b4e0c99c30"): NewEventProcessor(
		func(l types.Log) (interface{}, error) {
			return contractFilter.ParseBaseInitialized(l)
		},
		func(db *pgxpool.Pool, e interface{}) error {
			baseInitialized := e.(*contract.ContractBaseInitialized)
			_, err := db.Exec(context.Background(), `
			WITH inserted_event AS (
				INSERT INTO "Event" ("blockNumber", "txHash", "txIndex", "logIndex", "removed", "event")
				VALUES ($1, $2, $3, $4, $5, 'BaseInitialized')
				ON CONFLICT ("blockNumber", "logIndex") DO UPDATE
				SET "txHash" = excluded."txHash",
					"txIndex" = excluded."txIndex",
					"removed" = excluded."removed",
					"event" = excluded."event"
				RETURNING id
			)
			INSERT INTO "BaseInitialized" ("event_id", "name", "symbol", "timestamp")
			SELECT id, $6, $7, $8
			FROM inserted_event
			ON CONFLICT ("event_id") DO UPDATE
			SET "name" = excluded."name",
				"symbol" = excluded."symbol",
				"timestamp" = excluded."timestamp"
			`,
				baseInitialized.Raw.BlockNumber, baseInitialized.Raw.TxHash.Hex(), baseInitialized.Raw.TxIndex, baseInitialized.Raw.Index, baseInitialized.Raw.Removed, baseInitialized.Name, baseInitialized.Symbol, baseInitialized.Timestamp)
			if err != nil {
				return fmt.Errorf("error saving BaseInitialized event to database: %w", err)
			}
			return nil
		},
	),
	common.HexToHash("0x6cc19a794d6a439023150cd58748eed4353190c0bb060d2e6250e2df4a68b673"): NewEventProcessor(
		func(l types.Log) (interface{}, error) {
			return contractFilter.ParseCollectModuleWhitelisted(l)
		},
		func(db *pgxpool.Pool, e interface{}) error {
			collectModuleWhitelisted := e.(*contract.ContractCollectModuleWhitelisted)
			_, err := db.Exec(context.Background(), `
			WITH inserted_event AS (
				INSERT INTO "Event" ("blockNumber", "txHash", "txIndex", "logIndex", "removed", "event")
				VALUES ($1, $2, $3, $4, $5, 'CollectModuleWhitelisted')
				ON CONFLICT ("blockNumber", "logIndex") DO UPDATE
				SET "txHash" = excluded."txHash",
					"txIndex" = excluded."txIndex",
					"removed" = excluded."removed",
					"event" = excluded."event"
				RETURNING id
			)
			INSERT INTO "CollectModuleWhitelisted" ("event_id", "collectModule", "whitelisted", "timestamp")
			SELECT id, $6, $7, $8
			FROM inserted_event
			ON CONFLICT ("event_id") DO UPDATE
			SET "collectModule" = excluded."collectModule",
				"whitelisted" = excluded."whitelisted",
				"timestamp" = excluded."timestamp"
			`,
				collectModuleWhitelisted.Raw.BlockNumber, collectModuleWhitelisted.Raw.TxHash.Hex(), collectModuleWhitelisted.Raw.TxIndex, collectModuleWhitelisted.Raw.Index, collectModuleWhitelisted.Raw.Removed, collectModuleWhitelisted.CollectModule.Hex(), collectModuleWhitelisted.Whitelisted, collectModuleWhitelisted.Timestamp)
			if err != nil {
				return fmt.Errorf("error saving CollectModuleWhitelisted event to database: %w", err)
			}
			return nil
		},
	),
	common.HexToHash("0x0b227b550ffed48af813b32e246f787e99581ee13206ba8f9d90d63615269b3f"): NewEventProcessor(
		func(l types.Log) (interface{}, error) {
			return contractFilter.ParseCollectNFTDeployed(l)
		},
		func(db *pgxpool.Pool, e interface{}) error {
			collectNFTDeployed := e.(*contract.ContractCollectNFTDeployed)
			_, err := db.Exec(context.Background(), `
				WITH inserted_event AS (
					INSERT INTO "Event" ("blockNumber", "txHash", "txIndex", "logIndex", "removed", "event")
					VALUES ($1, $2, $3, $4, $5, 'CollectNFTDeployed')
					ON CONFLICT ("blockNumber", "logIndex") DO UPDATE
				SET "txHash" = excluded."txHash",
					"txIndex" = excluded."txIndex",
					"removed" = excluded."removed",
					"event" = excluded."event"
				RETURNING id
				)
				INSERT INTO "CollectNFTDeployed" ("event_id", "profileId", "pubId", "collectNFT", "timestamp")
				SELECT id, $6, $7, $8, $9
				FROM inserted_event
				ON CONFLICT ("event_id") DO UPDATE
				SET "profileId" = excluded."profileId",
					"pubId" = excluded."pubId",
					"collectNFT" = excluded."collectNFT",
					"timestamp" = excluded."timestamp"
				`,
				collectNFTDeployed.Raw.BlockNumber, collectNFTDeployed.Raw.TxHash.Hex(), collectNFTDeployed.Raw.TxIndex, collectNFTDeployed.Raw.Index, collectNFTDeployed.Raw.Removed, collectNFTDeployed.ProfileId, collectNFTDeployed.PubId, collectNFTDeployed.CollectNFT.Hex(), collectNFTDeployed.Timestamp)
			if err != nil {
				return fmt.Errorf("error saving CollectNFTDeployed event to database: %w", err)
			}
			return nil
		},
	),
	common.HexToHash("0x898a2dec95856255977a0fb48cebc30051d50c0d8d33f93dea1e3ddb2e342442"): NewEventProcessor(
		func(l types.Log) (interface{}, error) {
			return contractFilter.ParseCollectNFTInitialized(l)
		},
		func(db *pgxpool.Pool, e interface{}) error {
			collectNFTInitialized := e.(*contract.ContractCollectNFTInitialized)
			_, err := db.Exec(context.Background(), `
				WITH inserted_event AS (
					INSERT INTO "Event" ("blockNumber", "txHash", "txIndex", "logIndex", "removed", "event")
					VALUES ($1, $2, $3, $4, $5, 'CollectNFTInitialized')
					ON CONFLICT ("blockNumber", "logIndex") DO UPDATE
				SET "txHash" = excluded."txHash",
					"txIndex" = excluded."txIndex",
					"removed" = excluded."removed",
					"event" = excluded."event"
				RETURNING id
				)
				INSERT INTO "CollectNFTInitialized" ("event_id", "profileId", "pubId", "timestamp")
				SELECT id, $6, $7, $8
				FROM inserted_event
				ON CONFLICT ("event_id") DO UPDATE
				SET "profileId" = excluded."profileId",
					"pubId" = excluded."pubId",
					"timestamp" = excluded."timestamp"
				`,
				collectNFTInitialized.Raw.BlockNumber, collectNFTInitialized.Raw.TxHash.Hex(), collectNFTInitialized.Raw.TxIndex, collectNFTInitialized.Raw.Index, collectNFTInitialized.Raw.Removed, collectNFTInitialized.ProfileId, collectNFTInitialized.PubId, collectNFTInitialized.Timestamp)
			if err != nil {
				return fmt.Errorf("error saving CollectNFTInitialized event to database: %w", err)
			}
			return nil
		},
	),
	common.HexToHash("0x68edb7ec2c37d21b3b72233960b487f2966f4ac82b7430d39f24d1f8d6f99106"): NewEventProcessor(
		func(l types.Log) (interface{}, error) {
			return contractFilter.ParseCollectNFTTransferred(l)
		},
		func(db *pgxpool.Pool, e interface{}) error {
			collectNFTTransferred := e.(*contract.ContractCollectNFTTransferred)
			_, err := db.Exec(context.Background(), `
				WITH inserted_event AS (
					INSERT INTO "Event" ("blockNumber", "txHash", "txIndex", "logIndex", "removed", "event")
					VALUES ($1, $2, $3, $4, $5, 'CollectNFTTransferred')
					ON CONFLICT ("blockNumber", "logIndex") DO UPDATE
				SET "txHash" = excluded."txHash",
					"txIndex" = excluded."txIndex",
					"removed" = excluded."removed",
					"event" = excluded."event"
				RETURNING id
				)
				INSERT INTO "CollectNFTTransferred" ("event_id", "profileId", "pubId", "collectNFTId", "from", "to", "timestamp")
				SELECT id, $6, $7, $8, $9, $10, $11
				FROM inserted_event
				ON CONFLICT ("event_id") DO UPDATE
				SET "profileId" = excluded."profileId",
					"pubId" = excluded."pubId",
					"collectNFTId" = excluded."collectNFTId",
					"from" = excluded."from",
					"to" = excluded."to",
					"timestamp" = excluded."timestamp"
				`,
				collectNFTTransferred.Raw.BlockNumber, collectNFTTransferred.Raw.TxHash.Hex(), collectNFTTransferred.Raw.TxIndex, collectNFTTransferred.Raw.Index, collectNFTTransferred.Raw.Removed, collectNFTTransferred.ProfileId, collectNFTTransferred.PubId, collectNFTTransferred.CollectNFTId, collectNFTTransferred.From.Hex(), collectNFTTransferred.To.Hex(), collectNFTTransferred.Timestamp)
			if err != nil {
				return fmt.Errorf("error saving CollectNFTTransferred event to database: %w", err)
			}
			return nil
		},
	),

	common.HexToHash("0xed39bf0d9afa849610b901c9d7f4d00751ba20de2db023428065bec153833218"): NewEventProcessor(
		func(l types.Log) (interface{}, error) {
			return contractFilter.ParseCollected(l)
		},
		func(db *pgxpool.Pool, e interface{}) error {
			collected := e.(*contract.ContractCollected)
			_, err := db.Exec(context.Background(), `
				WITH inserted_event AS (
					INSERT INTO "Event" ("blockNumber", "txHash", "txIndex", "logIndex", "removed", "event")
					VALUES ($1, $2, $3, $4, $5, 'Collected')
					ON CONFLICT ("blockNumber", "logIndex") DO UPDATE
				SET "txHash" = excluded."txHash",
					"txIndex" = excluded."txIndex",
					"removed" = excluded."removed",
					"event" = excluded."event"
				RETURNING id
				)
				INSERT INTO "Collected" ("event_id", "collector", "profileId", "pubId", "rootProfileId", "rootPubId", "collectModuleData", "timestamp")
				SELECT id, $6, $7, $8, $9, $10, $11, $12
				FROM inserted_event
				ON CONFLICT ("event_id") DO UPDATE
				SET "collector" = excluded."collector",
					"profileId" = excluded."profileId",
					"pubId" = excluded."pubId",
					"rootProfileId" = excluded."rootProfileId",
					"rootPubId" = excluded."rootPubId",
					"collectModuleData" = excluded."collectModuleData",
					"timestamp" = excluded."timestamp"
				`,
				collected.Raw.BlockNumber, collected.Raw.TxHash.Hex(), collected.Raw.TxIndex, collected.Raw.Index, collected.Raw.Removed, collected.Collector.Hex(), collected.ProfileId, collected.PubId, collected.RootProfileId, collected.RootPubId, collected.CollectModuleData, collected.Timestamp)
			if err != nil {
				return fmt.Errorf("error saving Collected event to database: %w", err)
			}
			return nil
		},
	),
	common.HexToHash("0x7b4d1aa33773161799847429e4fbf29f56dbf1a3fe815f5070231cbfba402c37"): NewEventProcessor(
		func(l types.Log) (interface{}, error) {
			return contractFilter.ParseCommentCreated(l)
		},
		func(db *pgxpool.Pool, e interface{}) error {
			commentCreated := e.(*contract.ContractCommentCreated)
			_, err := db.Exec(context.Background(), `
				WITH inserted_event AS (
					INSERT INTO "Event" ("blockNumber", "txHash", "txIndex", "logIndex", "removed", "event")
					VALUES ($1, $2, $3, $4, $5, 'CommentCreated')
					ON CONFLICT ("blockNumber", "logIndex") DO UPDATE
				SET "txHash" = excluded."txHash",
					"txIndex" = excluded."txIndex",
					"removed" = excluded."removed",
					"event" = excluded."event"
				RETURNING id
				)
				INSERT INTO "CommentCreated" ("event_id", "profileId", "pubId", "contentURI", "profileIdPointed", "pubIdPointed","referenceModuleData", "collectModule", "collectModuleReturnData", "referenceModule", "referenceModuleReturnData", "timestamp")
				SELECT id, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16
				FROM inserted_event
				ON CONFLICT ("event_id") DO UPDATE
				SET "profileId" = excluded."profileId",
					"pubId" = excluded."pubId",
					"contentURI" = excluded."contentURI",
					"profileIdPointed" = excluded."profileIdPointed",
					"pubIdPointed" = excluded."pubIdPointed",
					"referenceModuleData" = excluded."referenceModuleData",
					"collectModule" = excluded."collectModule",
					"collectModuleReturnData" = excluded."collectModuleReturnData",
					"referenceModule" = excluded."referenceModule",
					"referenceModuleReturnData" = excluded."referenceModuleReturnData",
					"timestamp" = excluded."timestamp"
				`,
				commentCreated.Raw.BlockNumber, commentCreated.Raw.TxHash.Hex(), commentCreated.Raw.TxIndex, commentCreated.Raw.Index, commentCreated.Raw.Removed, commentCreated.ProfileId, commentCreated.PubId, commentCreated.ContentURI, commentCreated.ProfileIdPointed, commentCreated.PubIdPointed, commentCreated.ReferenceModuleData, commentCreated.CollectModule.Hex(), commentCreated.CollectModuleReturnData, commentCreated.ReferenceModule.Hex(), commentCreated.ReferenceModuleReturnData, commentCreated.Timestamp)
			if err != nil {
				return fmt.Errorf("error saving CommentCreated event to database: %w", err)
			}
			return nil
		},
	),
	common.HexToHash("0x0afd7c479e8bc7dcdb856b3cc27d2332dfe1f018fde574ea124919ddcae8a933"): NewEventProcessor(
		func(l types.Log) (interface{}, error) {
			return contractFilter.ParseDefaultProfileSet(l)
		},
		func(db *pgxpool.Pool, e interface{}) error {
			defaultProfileSet := e.(*contract.ContractDefaultProfileSet)
			_, err := db.Exec(context.Background(), `
				WITH inserted_event AS (
					INSERT INTO "Event" ("blockNumber", "txHash", "txIndex", "logIndex", "removed", "event")
					VALUES ($1, $2, $3, $4, $5, 'DefaultProfileSet')
					ON CONFLICT ("blockNumber", "logIndex") DO UPDATE
				SET "txHash" = excluded."txHash",
					"txIndex" = excluded."txIndex",
					"removed" = excluded."removed",
					"event" = excluded."event"
				RETURNING id
				)
				INSERT INTO "DefaultProfileSet" ("event_id", "wallet", "profileId", "timestamp")
				SELECT id, $6, $7, $8
				FROM inserted_event
				ON CONFLICT ("event_id") DO UPDATE
				SET "wallet" = excluded."wallet",
					"profileId" = excluded."profileId",
					"timestamp" = excluded."timestamp"
				`,
				defaultProfileSet.Raw.BlockNumber, defaultProfileSet.Raw.TxHash.Hex(), defaultProfileSet.Raw.TxIndex, defaultProfileSet.Raw.Index, defaultProfileSet.Raw.Removed, defaultProfileSet.Wallet.Hex(), defaultProfileSet.ProfileId, defaultProfileSet.Timestamp)
			if err != nil {
				return fmt.Errorf("error saving DefaultProfileSet event to database: %w", err)
			}
			return nil
		},
	),

	common.HexToHash("0x22baaec4952f35f59e45bd2ddb287e1ccc6d319375770c09428eb8f8d604e065"): NewEventProcessor(
		func(l types.Log) (interface{}, error) {
			return contractFilter.ParseDispatcherSet(l)
		},
		func(db *pgxpool.Pool, e interface{}) error {
			dispatcherSet := e.(*contract.ContractDispatcherSet)
			_, err := db.Exec(context.Background(), `
				WITH inserted_event AS ( 
					INSERT INTO "Event" ("blockNumber", "txHash", "txIndex", "logIndex", "removed", "event") 
					VALUES ($1, $2, $3, $4, $5, 'DispatcherSet') ON CONFLICT ("blockNumber", "logIndex") DO UPDATE
				SET "txHash" = excluded."txHash",
					"txIndex" = excluded."txIndex",
					"removed" = excluded."removed",
					"event" = excluded."event"
				RETURNING id ) 
					INSERT INTO "DispatcherSet" ("event_id", "profileId", "dispatcher", "timestamp") 
					SELECT id, $6, $7, $8 FROM inserted_event
				ON CONFLICT ("event_id") DO UPDATE
				SET "profileId" = excluded."profileId",
					"dispatcher" = excluded."dispatcher",
					"timestamp" = excluded."timestamp"
				`,
				dispatcherSet.Raw.BlockNumber, dispatcherSet.Raw.TxHash.Hex(), dispatcherSet.Raw.TxIndex, dispatcherSet.Raw.Index, dispatcherSet.Raw.Removed, dispatcherSet.ProfileId, dispatcherSet.Dispatcher.Hex(), dispatcherSet.Timestamp)
			if err != nil {
				return fmt.Errorf("error saving DispatcherSet event to database: %w", err)
			}
			return nil
		},
	),

	common.HexToHash("0x676c0801b0f400762e958ee31cfbb10870e70786f6761f57c8647e766b0db3d9"): NewEventProcessor(
		func(l types.Log) (interface{}, error) { return contractFilter.ParseEmergencyAdminSet(l) },
		func(db *pgxpool.Pool, e interface{}) error {
			emergencyAdminSet := e.(*contract.ContractEmergencyAdminSet)
			_, err := db.Exec(context.Background(), ` 
			WITH inserted_event AS ( INSERT INTO "Event" ("blockNumber", "txHash", "txIndex", "logIndex", "removed", "event") 
			VALUES ($1, $2, $3, $4, $5, 'EmergencyAdminSet') ON CONFLICT ("blockNumber", "logIndex") DO UPDATE
				SET "txHash" = excluded."txHash",
					"txIndex" = excluded."txIndex",
					"removed" = excluded."removed",
					"event" = excluded."event"
				RETURNING id ) 
			INSERT INTO "EmergencyAdminSet" ("event_id", "caller", "oldEmergencyAdmin", "newEmergencyAdmin", "timestamp") 
			SELECT id, $6, $7, $8, $9 FROM inserted_event
			ON CONFLICT ("event_id") DO UPDATE
				SET "caller" = excluded."caller",
					"oldEmergencyAdmin" = excluded."oldEmergencyAdmin",
					"newEmergencyAdmin" = excluded."newEmergencyAdmin",
					"timestamp" = excluded."timestamp"
			`,
				emergencyAdminSet.Raw.BlockNumber, emergencyAdminSet.Raw.TxHash.Hex(), emergencyAdminSet.Raw.TxIndex, emergencyAdminSet.Raw.Index, emergencyAdminSet.Raw.Removed, emergencyAdminSet.Caller.Hex(), emergencyAdminSet.OldEmergencyAdmin.Hex(), emergencyAdminSet.NewEmergencyAdmin.Hex(), emergencyAdminSet.Timestamp)
			if err != nil {
				return fmt.Errorf("error saving EmergencyAdminSet event to database: %w", err)
			}
			return nil
		},
	),

	common.HexToHash("0x4e84a529f4c627b5e787037d117873af1018768804cca3c7f0d47041fe2c89ed"): NewEventProcessor(
		func(l types.Log) (interface{}, error) {
			return contractFilter.ParseFeeModuleBaseConstructed(l)
		},
		func(db *pgxpool.Pool, e interface{}) error {
			feeModuleBaseConstructed := e.(*contract.ContractFeeModuleBaseConstructed)
			_, err := db.Exec(context.Background(), `
					WITH inserted_event AS (
						INSERT INTO "Event" ("blockNumber", "txHash", "txIndex", "logIndex", "removed", "event")
						VALUES ($1, $2, $3, $4, $5, 'FeeModuleBaseConstructed')
						ON CONFLICT ("blockNumber", "logIndex") DO UPDATE
				SET "txHash" = excluded."txHash",
					"txIndex" = excluded."txIndex",
					"removed" = excluded."removed",
					"event" = excluded."event"
				RETURNING id
					)
					INSERT INTO "FeeModuleBaseConstructed" ("event_id", "moduleGlobals", "timestamp")
					SELECT id, $6, $7
					FROM inserted_event
				ON CONFLICT ("event_id") DO UPDATE
				SET "moduleGlobals" = excluded."moduleGlobals",
					"timestamp" = excluded."timestamp"
					`,
				feeModuleBaseConstructed.Raw.BlockNumber, feeModuleBaseConstructed.Raw.TxHash.Hex(), feeModuleBaseConstructed.Raw.TxIndex, feeModuleBaseConstructed.Raw.Index, feeModuleBaseConstructed.Raw.Removed, feeModuleBaseConstructed.ModuleGlobals.Hex(), feeModuleBaseConstructed.Timestamp)
			if err != nil {
				return fmt.Errorf("error saving FeeModuleBaseConstructed event to database: %w", err)
			}
			return nil
		},
	),

	common.HexToHash("0x92d95e400932d129885e627b38b169cbb28443ffaaa282d0fba0cf8797721359"): NewEventProcessor(
		func(l types.Log) (interface{}, error) {
			return contractFilter.ParseFollowModuleSet(l)
		},
		func(db *pgxpool.Pool, e interface{}) error {
			followModuleSet := e.(*contract.ContractFollowModuleSet)
			_, err := db.Exec(context.Background(), `
					WITH inserted_event AS (
						INSERT INTO "Event" ("blockNumber", "txHash", "txIndex", "logIndex", "removed", "event")
						VALUES ($1, $2, $3, $4, $5, 'FollowModuleSet')
						ON CONFLICT ("blockNumber", "logIndex") DO UPDATE
				SET "txHash" = excluded."txHash",
					"txIndex" = excluded."txIndex",
					"removed" = excluded."removed",
					"event" = excluded."event"
				RETURNING id
					)
					INSERT INTO "FollowModuleSet" ("event_id", "profileId", "followModule", "followModuleReturnData", "timestamp")
					SELECT id, $6, $7, $8, $9
					FROM inserted_event
				ON CONFLICT ("event_id") DO UPDATE
				SET "profileId" = excluded."profileId",
					"followModule" = excluded."followModule",
					"followModuleReturnData" = excluded."followModuleReturnData",
					"timestamp" = excluded."timestamp"
					`,
				followModuleSet.Raw.BlockNumber, followModuleSet.Raw.TxHash.Hex(), followModuleSet.Raw.TxIndex, followModuleSet.Raw.Index, followModuleSet.Raw.Removed, followModuleSet.ProfileId, followModuleSet.FollowModule.Hex(), followModuleSet.FollowModuleReturnData, followModuleSet.Timestamp)
			if err != nil {
				return fmt.Errorf("error saving FollowModuleSet event to database: %w", err)
			}
			return nil
		},
	),

	common.HexToHash("0x52c5b7889df9f12f84ec3da051e854e5876678370d8357959c23ef59dd6486df"): NewEventProcessor(
		func(l types.Log) (interface{}, error) {
			return contractFilter.ParseFollowModuleWhitelisted(l)
		},
		func(db *pgxpool.Pool, e interface{}) error {
			followModuleWhitelisted := e.(*contract.ContractFollowModuleWhitelisted)
			_, err := db.Exec(context.Background(), `
				WITH inserted_event AS (
					INSERT INTO "Event" ("blockNumber", "txHash", "txIndex", "logIndex", "removed", "event")
					VALUES ($1, $2, $3, $4, $5, 'FollowModuleWhitelisted')
					ON CONFLICT ("blockNumber", "logIndex") DO UPDATE
				SET "txHash" = excluded."txHash",
					"txIndex" = excluded."txIndex",
					"removed" = excluded."removed",
					"event" = excluded."event"
				RETURNING id
				)
				INSERT INTO "FollowModuleWhitelisted" ("event_id", "followModule", "whitelisted", "timestamp")
				SELECT id, $6, $7, $8
				FROM inserted_event
				ON CONFLICT ("event_id") DO UPDATE
				SET "followModule" = excluded."followModule",
					"whitelisted" = excluded."whitelisted",
					"timestamp" = excluded."timestamp"
				`,
				followModuleWhitelisted.Raw.BlockNumber, followModuleWhitelisted.Raw.TxHash.Hex(), followModuleWhitelisted.Raw.TxIndex, followModuleWhitelisted.Raw.Index, followModuleWhitelisted.Raw.Removed, followModuleWhitelisted.FollowModule.Hex(), followModuleWhitelisted.Whitelisted, followModuleWhitelisted.Timestamp)
			if err != nil {
				return fmt.Errorf("error saving FollowModuleWhitelisted event to database: %w", err)
			}
			return nil
		},
	),

	common.HexToHash("0xd9a6070174f4ccca76ed4896432e9a090b16e07e8fe27f275f50b33500b98e52"): NewEventProcessor(
		func(l types.Log) (interface{}, error) {
			return contractFilter.ParseFollowNFTDelegatedPowerChanged(l)
		},
		func(db *pgxpool.Pool, e interface{}) error {
			followNFTDelegatedPowerChanged := e.(*contract.ContractFollowNFTDelegatedPowerChanged)
			_, err := db.Exec(context.Background(), `
					WITH inserted_event AS (
						INSERT INTO "Event" ("blockNumber", "txHash", "txIndex", "logIndex", "removed", "event")
						VALUES ($1, $2, $3, $4, $5, 'FollowNFTDelegatedPowerChanged')
						ON CONFLICT ("blockNumber", "logIndex") DO UPDATE
				SET "txHash" = excluded."txHash",
					"txIndex" = excluded."txIndex",
					"removed" = excluded."removed",
					"event" = excluded."event"
				RETURNING id
					)
					INSERT INTO "FollowNFTDelegatedPowerChanged" ("event_id", "delegate", "newPower", "timestamp")
					SELECT id, $6, $7, $8
					FROM inserted_event
				ON CONFLICT ("event_id") DO UPDATE
				SET "delegate" = excluded."delegate",
					"newPower" = excluded."newPower",
					"timestamp" = excluded."timestamp"
					`,
				followNFTDelegatedPowerChanged.Raw.BlockNumber, followNFTDelegatedPowerChanged.Raw.TxHash.Hex(), followNFTDelegatedPowerChanged.Raw.TxIndex, followNFTDelegatedPowerChanged.Raw.Index, followNFTDelegatedPowerChanged.Raw.Removed, followNFTDelegatedPowerChanged.Delegate.Hex(), followNFTDelegatedPowerChanged.NewPower, followNFTDelegatedPowerChanged.Timestamp)
			if err != nil {
				return fmt.Errorf("error saving FollowNFTDelegatedPowerChanged event to database: %w", err)
			}
			return nil
		},
	),

	common.HexToHash("0x44403e38baed5e40df7f64ff8708b076c75a0dfda8380e75df5c36f11a476743"): NewEventProcessor(
		func(l types.Log) (interface{}, error) {
			return contractFilter.ParseFollowNFTDeployed(l)
		},
		func(db *pgxpool.Pool, e interface{}) error {
			followNFTDeployed := e.(*contract.ContractFollowNFTDeployed)
			_, err := db.Exec(context.Background(), `
					WITH inserted_event AS (
						INSERT INTO "Event" ("blockNumber", "txHash", "txIndex", "logIndex", "removed", "event")
						VALUES ($1, $2, $3, $4, $5, 'FollowNFTDeployed')
						ON CONFLICT ("blockNumber", "logIndex") DO UPDATE
				SET "txHash" = excluded."txHash",
					"txIndex" = excluded."txIndex",
					"removed" = excluded."removed",
					"event" = excluded."event"
				RETURNING id
					)
					INSERT INTO "FollowNFTDeployed" ("event_id", "profileId", "followNFT", "timestamp")
					SELECT id, $6, $7, $8
					FROM inserted_event
				ON CONFLICT ("event_id") DO UPDATE
				SET "profileId" = excluded."profileId",
					"followNFT" = excluded."followNFT",
					"timestamp" = excluded."timestamp"
					`,
				followNFTDeployed.Raw.BlockNumber, followNFTDeployed.Raw.TxHash.Hex(), followNFTDeployed.Raw.TxIndex, followNFTDeployed.Raw.Index, followNFTDeployed.Raw.Removed, followNFTDeployed.ProfileId, followNFTDeployed.FollowNFT.Hex(), followNFTDeployed.Timestamp)
			if err != nil {
				return fmt.Errorf("error saving FollowNFTDeployed event to database: %w", err)
			}
			return nil
		},
	),

	common.HexToHash("0xaec15127df11a6b562c87d31bcb8f4cd2f0cf57fb9b663d6334abf41fea94d95"): NewEventProcessor(
		func(l types.Log) (interface{}, error) {
			return contractFilter.ParseFollowNFTInitialized(l)
		},
		func(db *pgxpool.Pool, e interface{}) error {
			followNFTInitialized := e.(*contract.ContractFollowNFTInitialized)
			_, err := db.Exec(context.Background(), `
					WITH inserted_event AS (
						INSERT INTO "Event" ("blockNumber", "txHash", "txIndex", "logIndex", "removed", "event")
						VALUES ($1, $2, $3, $4, $5, 'FollowNFTInitialized')
						ON CONFLICT ("blockNumber", "logIndex") DO UPDATE
				SET "txHash" = excluded."txHash",
					"txIndex" = excluded."txIndex",
					"removed" = excluded."removed",
					"event" = excluded."event"
				RETURNING id
					)
					INSERT INTO "FollowNFTInitialized" ("event_id", "profileId", "timestamp")
					SELECT id, $6, $7
					FROM inserted_event
				ON CONFLICT ("event_id") DO UPDATE
				SET "profileId" = excluded."profileId",
					"timestamp" = excluded."timestamp"
					`,
				followNFTInitialized.Raw.BlockNumber, followNFTInitialized.Raw.TxHash.Hex(), followNFTInitialized.Raw.TxIndex, followNFTInitialized.Raw.Index, followNFTInitialized.Raw.Removed, followNFTInitialized.ProfileId, followNFTInitialized.Timestamp)
			if err != nil {
				return fmt.Errorf("error saving FollowNFTInitialized event to database: %w", err)
			}
			return nil
		},
	),

	common.HexToHash("0x4996ad2257e7db44908136c43128cc10ca988096f67dc6bb0bcee11d151368fb"): NewEventProcessor(
		func(l types.Log) (interface{}, error) {
			return contractFilter.ParseFollowNFTTransferred(l)
		},
		func(db *pgxpool.Pool, e interface{}) error {
			followNFTTransferred := e.(*contract.ContractFollowNFTTransferred)
			_, err := db.Exec(context.Background(), `
			WITH inserted_event AS ( 
				INSERT INTO "Event" ("blockNumber", "txHash", "txIndex", "logIndex", "removed", "event"
			) 
			VALUES ($1, $2, $3, $4, $5, 'FollowNFTTransferred') ON CONFLICT ("blockNumber", "logIndex") DO UPDATE
				SET "txHash" = excluded."txHash",
					"txIndex" = excluded."txIndex",
					"removed" = excluded."removed",
					"event" = excluded."event"
				RETURNING id ) 
			INSERT INTO "FollowNFTTransferred" ("event_id", "profileId", "followNFTId", "from", "to", "timestamp") 
			SELECT id, $6, $7, $8, $9, $10 FROM inserted_event
			ON CONFLICT ("event_id") DO UPDATE
				SET "profileId" = excluded."profileId",
					"followNFTId" = excluded."followNFTId",
					"from" = excluded."from",
					"to" = excluded."to",
					"timestamp" = excluded."timestamp"
			`,
				followNFTTransferred.Raw.BlockNumber, followNFTTransferred.Raw.TxHash.Hex(), followNFTTransferred.Raw.TxIndex, followNFTTransferred.Raw.Index, followNFTTransferred.Raw.Removed, followNFTTransferred.ProfileId, followNFTTransferred.FollowNFTId, followNFTTransferred.From.Hex(), followNFTTransferred.To.Hex(), followNFTTransferred.Timestamp)
			if err != nil {
				return fmt.Errorf("error saving FollowNFTTransferred event to database: %w", err)
			}
			return nil
		},
	),

	common.HexToHash("0xe82886e1af6fcab5caef13815b22f51384e970c367a785f265d13860a7d6966d"): NewEventProcessor(
		func(l types.Log) (interface{}, error) {
			return contractFilter.ParseFollowNFTURISet(l)
		},
		func(db *pgxpool.Pool, e interface{}) error {
			followNFTURISet := e.(*contract.ContractFollowNFTURISet)
			_, err := db.Exec(context.Background(), `
				WITH inserted_event AS ( 
					INSERT INTO "Event" ("blockNumber", "txHash", "txIndex", "logIndex", "removed", "event" ) 
					VALUES ($1, $2, $3, $4, $5, 'FollowNFTURISet') ON CONFLICT ("blockNumber", "logIndex") DO UPDATE
				SET "txHash" = excluded."txHash",
					"txIndex" = excluded."txIndex",
					"removed" = excluded."removed",
					"event" = excluded."event"
				RETURNING id ) 
				INSERT INTO "FollowNFTURISet" ("event_id", "profileId", "followNFTURI", "timestamp") 
				SELECT id, $6, $7, $8 FROM inserted_event
				ON CONFLICT ("event_id") DO UPDATE
				SET "profileId" = excluded."profileId",
					"followNFTURI" = excluded."followNFTURI",
					"timestamp" = excluded."timestamp"
				`,
				followNFTURISet.Raw.BlockNumber, followNFTURISet.Raw.TxHash.Hex(), followNFTURISet.Raw.TxIndex, followNFTURISet.Raw.Index, followNFTURISet.Raw.Removed, followNFTURISet.ProfileId, followNFTURISet.FollowNFTURI, followNFTURISet.Timestamp)
			if err != nil {
				return fmt.Errorf("error saving FollowNFTURISet event to database: %w", err)
			}
			return nil
		}),

	common.HexToHash("0x40487072dc56f384287d26fbe090f404143c2737d54632177451d1f74bd82c76"): NewEventProcessor(
		func(l types.Log) (interface{}, error) {
			return contractFilter.ParseFollowed(l)
		},
		func(db *pgxpool.Pool, e interface{}) error {
			followed := e.(*contract.ContractFollowed)
			profileIds := make([]uint64, len(followed.ProfileIds))
			for i, profileId := range followed.ProfileIds {
				profileIds[i] = profileId.Uint64()
			}
			_, err := db.Exec(context.Background(), `
					WITH inserted_event AS ( 
						INSERT INTO "Event" ("blockNumber", "txHash", "txIndex", "logIndex", "removed", "event" ) 
						VALUES ($1, $2, $3, $4, $5, 'Followed') ON CONFLICT ("blockNumber", "logIndex") DO UPDATE
				SET "txHash" = excluded."txHash",
					"txIndex" = excluded."txIndex",
					"removed" = excluded."removed",
					"event" = excluded."event"
				RETURNING id ) 
					INSERT INTO "Followed" ("event_id", "follower", "profileIds", "followModuleDatas", "timestamp") 
					SELECT id, $6, $7, $8, $9 FROM inserted_event
					ON CONFLICT ("event_id") DO UPDATE
				SET "follower" = excluded."follower",
					"profileIds" = excluded."profileIds",
					"followModuleDatas" = excluded."followModuleDatas",
					"timestamp" = excluded."timestamp"
					`,
				followed.Raw.BlockNumber, followed.Raw.TxHash.Hex(), followed.Raw.TxIndex, followed.Raw.Index, followed.Raw.Removed, followed.Follower.Hex(), profileIds, followed.FollowModuleDatas, followed.Timestamp)
			if err != nil {
				return fmt.Errorf("error saving Followed event to database: %w", err)
			}
			return nil
		}),

	common.HexToHash("0xc67fc3972da5d6434ab7b796ba133c240d40ee4e69129963c5aa0f2a6f7c3ad6"): NewEventProcessor(
		func(l types.Log) (interface{}, error) {
			return contractFilter.ParseFollowsApproved(l)
		},
		func(db *pgxpool.Pool, e interface{}) error {
			followsApproved := e.(*contract.ContractFollowsApproved)
			addresses := make([]string, len(followsApproved.Addresses))
			for i, address := range followsApproved.Addresses {
				addresses[i] = address.Hex()
			}
			_, err := db.Exec(context.Background(), `
					WITH inserted_event AS ( 
						INSERT INTO "Event" ("blockNumber", "txHash", "txIndex", "logIndex", "removed", "event" ) 
						VALUES ($1, $2, $3, $4, $5, 'FollowsApproved') ON CONFLICT ("blockNumber", "logIndex") DO UPDATE
				SET "txHash" = excluded."txHash",
					"txIndex" = excluded."txIndex",
					"removed" = excluded."removed",
					"event" = excluded."event"
				RETURNING id ) 
					INSERT INTO "FollowsApproved" ("event_id", "owner", "profileId", "addresses", "approved", "timestamp") 
					SELECT id, $6, $7, $8, $9, $10 FROM inserted_event
					ON CONFLICT ("event_id") DO UPDATE
				SET "owner" = excluded."owner",
					"profileId" = excluded."profileId",
					"addresses" = excluded."addresses",
					"approved" = excluded."approved",
					"timestamp" = excluded."timestamp"
					`,
				followsApproved.Raw.BlockNumber, followsApproved.Raw.TxHash.Hex(), followsApproved.Raw.TxIndex, followsApproved.Raw.Index, followsApproved.Raw.Removed, followsApproved.Owner.Hex(), followsApproved.ProfileId, addresses, followsApproved.Approved, followsApproved.Timestamp)
			if err != nil {
				return fmt.Errorf("error saving FollowsApproved event to database: %w", err)
			}
			return nil
		}),

	common.HexToHash("0x5538c80c8d3bee397d87a7d153f7f085bb12adf2fe25a026c7cc4e83d8c5f1d7"): NewEventProcessor(
		func(l types.Log) (interface{}, error) {
			return contractFilter.ParseFollowsToggled(l)
		},
		func(db *pgxpool.Pool, e interface{}) error {
			followsToggled := e.(*contract.ContractFollowsToggled)
			profileIds := make([]uint64, len(followsToggled.ProfileIds))
			for i, profileId := range followsToggled.ProfileIds {
				profileIds[i] = profileId.Uint64()
			}
			_, err := db.Exec(context.Background(), `
					WITH inserted_event AS (
						INSERT INTO "Event" ("blockNumber", "txHash", "txIndex", "logIndex", "removed", "event" )
						VALUES ($1, $2, $3, $4, $5, 'FollowsToggled') ON CONFLICT ("blockNumber", "logIndex") DO UPDATE
				SET "txHash" = excluded."txHash",
					"txIndex" = excluded."txIndex",
					"removed" = excluded."removed",
					"event" = excluded."event"
				RETURNING id )
					INSERT INTO "FollowsToggled" ("event_id", "owner", "profileIds", "enabled", "timestamp")
					SELECT id, $6, $7, $8, $9 FROM inserted_event
					ON CONFLICT ("event_id") DO UPDATE
				SET "owner" = excluded."owner",
					"profileIds" = excluded."profileIds",
					"enabled" = excluded."enabled",
					"timestamp" = excluded."timestamp"
					`,
				followsToggled.Raw.BlockNumber, followsToggled.Raw.TxHash.Hex(), followsToggled.Raw.TxIndex, followsToggled.Raw.Index, followsToggled.Raw.Removed, followsToggled.Owner.Hex(), profileIds, followsToggled.Enabled, followsToggled.Timestamp)
			if err != nil {
				return fmt.Errorf("error saving FollowsToggled event to database: %w", err)
			}
			return nil
		}),

	common.HexToHash("0xe552a55455b740845a5c07ed445d1724142fc997b389835495a29b30cddc1ccd"): NewEventProcessor(
		func(l types.Log) (interface{}, error) {
			return contractFilter.ParseGovernanceSet(l)
		},
		func(db *pgxpool.Pool, e interface{}) error {
			governanceSet := e.(*contract.ContractGovernanceSet)
			_, err := db.Exec(context.Background(), `
					WITH inserted_event AS (
						INSERT INTO "Event" ("blockNumber", "txHash", "txIndex", "logIndex", "removed", "event" )
						VALUES ($1, $2, $3, $4, $5, 'GovernanceSet') ON CONFLICT ("blockNumber", "logIndex") DO UPDATE
				SET "txHash" = excluded."txHash",
					"txIndex" = excluded."txIndex",
					"removed" = excluded."removed",
					"event" = excluded."event"
				RETURNING id )
					INSERT INTO "GovernanceSet" ("event_id", "caller", "prevGovernance", "newGovernance", "timestamp")
					SELECT id, $6, $7, $8, $9 FROM inserted_event
					ON CONFLICT ("event_id") DO UPDATE
				SET "caller" = excluded."caller",
					"prevGovernance" = excluded."prevGovernance",
					"newGovernance" = excluded."newGovernance",
					"timestamp" = excluded."timestamp"
					`,
				governanceSet.Raw.BlockNumber, governanceSet.Raw.TxHash.Hex(), governanceSet.Raw.TxIndex, governanceSet.Raw.Index, governanceSet.Raw.Removed, governanceSet.Caller.Hex(), governanceSet.PrevGovernance.Hex(), governanceSet.NewGovernance.Hex(), governanceSet.Timestamp)
			if err != nil {
				return fmt.Errorf("error saving GovernanceSet event to database: %w", err)
			}
			return nil
		}),

	common.HexToHash("0x9ea5dedb85bd9da4e264ee5a39b7ba0982e5d4d035d55edfa98a36b00e770b5a"): NewEventProcessor(
		func(l types.Log) (interface{}, error) {
			return contractFilter.ParseMirrorCreated(l)
		},
		func(db *pgxpool.Pool, e interface{}) error {
			mirrorCreated := e.(*contract.ContractMirrorCreated)
			_, err := db.Exec(context.Background(), `
					WITH inserted_event AS (
						INSERT INTO "Event" ("blockNumber", "txHash", "txIndex", "logIndex", "removed", "event" )
						VALUES ($1, $2, $3, $4, $5, 'MirrorCreated') ON CONFLICT ("blockNumber", "logIndex") DO UPDATE
				SET "txHash" = excluded."txHash",
					"txIndex" = excluded."txIndex",
					"removed" = excluded."removed",
					"event" = excluded."event"
				RETURNING id )
					INSERT INTO "MirrorCreated" ("event_id", "profileId", "pubId", "profileIdPointed", "pubIdPointed", "referenceModuleData", "referenceModule", "referenceModuleReturnData", "timestamp")
					SELECT id, $6, $7, $8, $9, $10, $11, $12, $13 FROM inserted_event
					ON CONFLICT ("event_id") DO UPDATE
				SET "profileId" = excluded."profileId",
					"pubId" = excluded."pubId",
					"profileIdPointed" = excluded."profileIdPointed",
					"pubIdPointed" = excluded."pubIdPointed",
					"referenceModuleData" = excluded."referenceModuleData",
					"referenceModule" = excluded."referenceModule",
					"referenceModuleReturnData" = excluded."referenceModuleReturnData",
					"timestamp" = excluded."timestamp"
					`,
				mirrorCreated.Raw.BlockNumber, mirrorCreated.Raw.TxHash.Hex(), mirrorCreated.Raw.TxIndex, mirrorCreated.Raw.Index, mirrorCreated.Raw.Removed, mirrorCreated.ProfileId, mirrorCreated.PubId, mirrorCreated.ProfileIdPointed, mirrorCreated.PubIdPointed, mirrorCreated.ReferenceModuleData, mirrorCreated.ReferenceModule.Hex(), mirrorCreated.ReferenceModuleReturnData, mirrorCreated.Timestamp)
			if err != nil {
				return fmt.Errorf("error saving MirrorCreated event to database: %w", err)
			}
			return nil
		}),

	common.HexToHash("0xf1a1fa6b64aa95186f5a1285e76198d0da80d9c5a88062641d447f1d7c54e56c"): NewEventProcessor(
		func(l types.Log) (interface{}, error) {
			return contractFilter.ParseModuleBaseConstructed(l)
		},
		func(db *pgxpool.Pool, e interface{}) error {
			moduleBaseConstructed := e.(*contract.ContractModuleBaseConstructed)
			_, err := db.Exec(context.Background(), `
				WITH inserted_event AS ( 
					INSERT INTO "Event" ("blockNumber", "txHash", "txIndex", "logIndex", "removed", "event" 
				) 
				VALUES ($1, $2, $3, $4, $5, 'ModuleBaseConstructed') ON CONFLICT ("blockNumber", "logIndex") DO UPDATE
				SET "txHash" = excluded."txHash",
					"txIndex" = excluded."txIndex",
					"removed" = excluded."removed",
					"event" = excluded."event"
				RETURNING id )
				INSERT INTO "ModuleBaseConstructed" ("event_id", "hub", "timestamp")
				SELECT id, $6, $7 FROM inserted_event
				ON CONFLICT ("event_id") DO UPDATE
				SET "hub" = excluded."hub",
					"timestamp" = excluded."timestamp"
				`,
				moduleBaseConstructed.Raw.BlockNumber, moduleBaseConstructed.Raw.TxHash.Hex(), moduleBaseConstructed.Raw.TxIndex, moduleBaseConstructed.Raw.Index, moduleBaseConstructed.Raw.Removed, moduleBaseConstructed.Hub.Hex(), moduleBaseConstructed.Timestamp)
			if err != nil {
				return fmt.Errorf("error saving ModuleBaseConstructed event to database: %w", err)
			}
			return nil
		}),

	common.HexToHash("0x79c3cefc851fd6040f06af202c542818d9fb39bcddcb7a7e3f563b15300a2743"): NewEventProcessor(
		func(l types.Log) (interface{}, error) {
			return contractFilter.ParseModuleGlobalsCurrencyWhitelisted(l)
		},
		func(db *pgxpool.Pool, e interface{}) error {
			moduleGlobalsCurrencyWhitelisted := e.(*contract.ContractModuleGlobalsCurrencyWhitelisted)
			_, err := db.Exec(context.Background(), `WITH inserted_event AS ( INSERT INTO "Event" ("blockNumber", "txHash", "txIndex", "logIndex", "removed", "event" ) VALUES ($1, $2, $3, $4, $5, 'ModuleGlobalsCurrencyWhitelisted') ON CONFLICT ("blockNumber", "logIndex") DO UPDATE
				SET "txHash" = excluded."txHash",
					"txIndex" = excluded."txIndex",
					"removed" = excluded."removed",
					"event" = excluded."event"
				RETURNING id ) INSERT INTO "ModuleGlobalsCurrencyWhitelisted" ("event_id", "currency", "prevWhitelisted", "whitelisted", "timestamp") SELECT id, $6, $7, $8, $9 FROM inserted_event
				ON CONFLICT ("event_id") DO UPDATE
				SET "currency" = excluded."currency",
					"prevWhitelisted" = excluded."prevWhitelisted",
					"whitelisted" = excluded."whitelisted",
					"timestamp" = excluded."timestamp"
				`,
				moduleGlobalsCurrencyWhitelisted.Raw.BlockNumber, moduleGlobalsCurrencyWhitelisted.Raw.TxHash.Hex(), moduleGlobalsCurrencyWhitelisted.Raw.TxIndex, moduleGlobalsCurrencyWhitelisted.Raw.Index, moduleGlobalsCurrencyWhitelisted.Raw.Removed, moduleGlobalsCurrencyWhitelisted.Currency.Hex(), moduleGlobalsCurrencyWhitelisted.PrevWhitelisted, moduleGlobalsCurrencyWhitelisted.Whitelisted, moduleGlobalsCurrencyWhitelisted.Timestamp)
			if err != nil {
				return fmt.Errorf("error saving ModuleGlobalsCurrencyWhitelisted event to database: %w", err)
			}
			return nil
		}),

	common.HexToHash("0xbf538a2c0db3d440906b8179dd0394f68a65b0b1481da70ffee24e19dccee84c"): NewEventProcessor(
		func(l types.Log) (interface{}, error) {
			return contractFilter.ParseModuleGlobalsGovernanceSet(l)
		},
		func(db *pgxpool.Pool, e interface{}) error {
			governanceSet := e.(*contract.ContractModuleGlobalsGovernanceSet)
			_, err := db.Exec(context.Background(), `WITH inserted_event AS ( INSERT INTO "Event" ("blockNumber", "txHash", "txIndex", "logIndex", "removed", "event") VALUES ($1, $2, $3, $4, $5, 'ModuleGlobalsGovernanceSet') ON CONFLICT ("blockNumber", "logIndex") DO UPDATE
				SET "txHash" = excluded."txHash",
					"txIndex" = excluded."txIndex",
					"removed" = excluded."removed",
					"event" = excluded."event"
				RETURNING id) INSERT INTO "ModuleGlobalsGovernanceSet" ("event_id", "prevGovernance", "newGovernance", "timestamp") SELECT id, $6, $7, $8 FROM inserted_event
				ON CONFLICT ("event_id") DO UPDATE
				SET "prevGovernance" = excluded."prevGovernance",
					"newGovernance" = excluded."newGovernance",
					"timestamp" = excluded."timestamp"
				`,
				governanceSet.Raw.BlockNumber, governanceSet.Raw.TxHash.Hex(), governanceSet.Raw.TxIndex, governanceSet.Raw.Index, governanceSet.Raw.Removed, governanceSet.PrevGovernance.Hex(), governanceSet.NewGovernance.Hex(), governanceSet.Timestamp)
			if err != nil {
				return fmt.Errorf("error saving ModuleGlobalsGovernanceSet event to database: %w", err)
			}
			return nil
		}),

	common.HexToHash("0xec936862e6bb897cd711a5f31825057583128c2a482f0a4c9a4e6c3fd7c023f4"): NewEventProcessor(
		func(l types.Log) (interface{}, error) {
			return contractFilter.ParseModuleGlobalsTreasuryFeeSet(l)
		},
		func(db *pgxpool.Pool, e interface{}) error {
			moduleGlobalsTreasuryFeeSet := e.(*contract.ContractModuleGlobalsTreasuryFeeSet)
			_, err := db.Exec(context.Background(), `WITH inserted_event AS ( INSERT INTO "Event" ("blockNumber", "txHash", "txIndex", "logIndex", "removed", "event" ) VALUES ($1, $2, $3, $4, $5, 'ModuleGlobalsTreasuryFeeSet') ON CONFLICT ("blockNumber", "logIndex") DO UPDATE
				SET "txHash" = excluded."txHash",
					"txIndex" = excluded."txIndex",
					"removed" = excluded."removed",
					"event" = excluded."event"
				RETURNING id ) INSERT INTO "ModuleGlobalsTreasuryFeeSet" ("event_id", "prevTreasuryFee", "newTreasuryFee", "timestamp") SELECT id, $6, $7, $8 FROM inserted_event
				ON CONFLICT ("event_id") DO UPDATE
				SET "prevTreasuryFee" = excluded."prevTreasuryFee",
					"newTreasuryFee" = excluded."newTreasuryFee",
					"timestamp" = excluded."timestamp"
				`,
				moduleGlobalsTreasuryFeeSet.Raw.BlockNumber, moduleGlobalsTreasuryFeeSet.Raw.TxHash.Hex(), moduleGlobalsTreasuryFeeSet.Raw.TxIndex, moduleGlobalsTreasuryFeeSet.Raw.Index, moduleGlobalsTreasuryFeeSet.Raw.Removed, moduleGlobalsTreasuryFeeSet.PrevTreasuryFee, moduleGlobalsTreasuryFeeSet.NewTreasuryFee, moduleGlobalsTreasuryFeeSet.Timestamp)
			if err != nil {
				return fmt.Errorf("error saving ModuleGlobalsTreasuryFeeSet event to database: %w", err)
			}
			return nil
		}),

	common.HexToHash("0x3dfc53d6b49bfbc932b215ba515f0d0ab0e17aac17726fba48075f0c16c7ffe3"): NewEventProcessor(
		func(l types.Log) (interface{}, error) {
			return contractFilter.ParseModuleGlobalsTreasurySet(l)
		},
		func(db *pgxpool.Pool, e interface{}) error {
			moduleGlobalsTreasurySet := e.(*contract.ContractModuleGlobalsTreasurySet)
			_, err := db.Exec(context.Background(), `WITH inserted_event AS ( INSERT INTO "Event" ("blockNumber", "txHash", "txIndex", "logIndex", "removed", "event" ) VALUES ($1, $2, $3, $4, $5, 'ModuleGlobalsTreasurySet') ON CONFLICT ("blockNumber", "logIndex") DO UPDATE
				SET "txHash" = excluded."txHash",
					"txIndex" = excluded."txIndex",
					"removed" = excluded."removed",
					"event" = excluded."event"
				RETURNING id ) INSERT INTO "ModuleGlobalsTreasurySet" ("event_id", "prevTreasury", "newTreasury", "timestamp") SELECT id, $6, $7, $8 FROM inserted_event
				ON CONFLICT ("event_id") DO UPDATE
				SET "prevTreasury" = excluded."prevTreasury",
					"newTreasury" = excluded."newTreasury",
					"timestamp" = excluded."timestamp"
				`,
				moduleGlobalsTreasurySet.Raw.BlockNumber, moduleGlobalsTreasurySet.Raw.TxHash.Hex(), moduleGlobalsTreasurySet.Raw.TxIndex, moduleGlobalsTreasurySet.Raw.Index, moduleGlobalsTreasurySet.Raw.Removed, moduleGlobalsTreasurySet.PrevTreasury.Hex(), moduleGlobalsTreasurySet.NewTreasury.Hex(), moduleGlobalsTreasurySet.Timestamp)
			if err != nil {
				return fmt.Errorf("error saving ModuleGlobalsTreasurySet event to database: %w", err)
			}
			return nil
		}),

	common.HexToHash("0xc672c38b4d26c3c978228e99164105280410b144af24dd3ed8e4f9d211d96a50"): NewEventProcessor(
		func(l types.Log) (interface{}, error) {
			return contractFilter.ParsePostCreated(l)
		},
		func(db *pgxpool.Pool, e interface{}) error {
			postCreated := e.(*contract.ContractPostCreated)
			_, err := db.Exec(context.Background(), `WITH inserted_event AS ( INSERT INTO "Event" ("blockNumber", "txHash", "txIndex", "logIndex", "removed", "event") VALUES ($1, $2, $3, $4, $5, 'PostCreated') ON CONFLICT ("blockNumber", "logIndex") DO UPDATE
				SET "txHash" = excluded."txHash",
					"txIndex" = excluded."txIndex",
					"removed" = excluded."removed",
					"event" = excluded."event"
				RETURNING id ) INSERT INTO "PostCreated" ("event_id", "profileId", "pubId", "contentURI", "collectModule", "collectModuleReturnData", "referenceModule", "referenceModuleReturnData", "timestamp") SELECT id, $6, $7, $8, $9, $10, $11, $12, $13 FROM inserted_event
				ON CONFLICT ("event_id") DO UPDATE
				SET "profileId" = excluded."profileId",
					"pubId" = excluded."pubId",
					"contentURI" = excluded."contentURI",
					"collectModule" = excluded."collectModule",
					"collectModuleReturnData" = excluded."collectModuleReturnData",
					"referenceModule" = excluded."referenceModule",
					"referenceModuleReturnData" = excluded."referenceModuleReturnData",
					"timestamp" = excluded."timestamp"
				`,
				postCreated.Raw.BlockNumber, postCreated.Raw.TxHash.Hex(), postCreated.Raw.TxIndex, postCreated.Raw.Index, postCreated.Raw.Removed, postCreated.ProfileId, postCreated.PubId, postCreated.ContentURI, postCreated.CollectModule.Hex(), postCreated.CollectModuleReturnData, postCreated.ReferenceModule.Hex(), postCreated.ReferenceModuleReturnData, postCreated.Timestamp)
			if err != nil {
				return fmt.Errorf("error saving PostCreated event to database: %w", err)
			}
			return nil
		}),

	common.HexToHash("0x4e14f57cff7910416f2ef43cf05019b5a97a313de71fec9344be11b9b88fed12"): NewEventProcessor(
		func(l types.Log) (interface{}, error) {
			return contractFilter.ParseProfileCreated(l)
		},
		func(db *pgxpool.Pool, e interface{}) error {
			profileCreated := e.(*contract.ContractProfileCreated)
			_, err := db.Exec(context.Background(), `WITH inserted_event AS ( INSERT INTO "Event" ("blockNumber", "txHash", "txIndex", "logIndex", "removed", "event" ) VALUES ($1, $2, $3, $4, $5, 'ProfileCreated') ON CONFLICT ("blockNumber", "logIndex") DO UPDATE
				SET "txHash" = excluded."txHash",
					"txIndex" = excluded."txIndex",
					"removed" = excluded."removed",
					"event" = excluded."event"
				RETURNING id ) INSERT INTO "ProfileCreated" ("event_id", "profileId", "creator", "to", "handle", "imageURI", "followModule", "followModuleReturnData", "followNFTURI", "timestamp") SELECT id, $6, $7, $8, $9, $10, $11, $12, $13, $14 FROM inserted_event
				ON CONFLICT ("event_id") DO UPDATE
				SET "profileId" = excluded."profileId",
					"creator" = excluded."creator",
					"to" = excluded."to",
					"handle" = excluded."handle",
					"imageURI" = excluded."imageURI",
					"followModule" = excluded."followModule",
					"followModuleReturnData" = excluded."followModuleReturnData",
					"followNFTURI" = excluded."followNFTURI",
					"timestamp" = excluded."timestamp"
				`,
				profileCreated.Raw.BlockNumber, profileCreated.Raw.TxHash.Hex(), profileCreated.Raw.TxIndex, profileCreated.Raw.Index, profileCreated.Raw.Removed, profileCreated.ProfileId, profileCreated.Creator.Hex(), profileCreated.To.Hex(), profileCreated.Handle, profileCreated.ImageURI, profileCreated.FollowModule.Hex(), profileCreated.FollowModuleReturnData, profileCreated.FollowNFTURI, profileCreated.Timestamp)
			if err != nil {
				return fmt.Errorf("error saving ProfileCreated event to database: %w", err)
			}
			return nil
		}),

	common.HexToHash("0x8f617843889b94892bd44852d36ca6a7f49ecf4350a01e7b68e22d80f4ed95bc"): NewEventProcessor(
		func(l types.Log) (interface{}, error) {
			return contractFilter.ParseProfileCreatorWhitelisted(l)
		},
		func(db *pgxpool.Pool, e interface{}) error {
			profileCreatorWhitelisted := e.(*contract.ContractProfileCreatorWhitelisted)
			_, err := db.Exec(context.Background(), `WITH inserted_event AS ( INSERT INTO "Event" ("blockNumber", "txHash", "txIndex", "logIndex", "removed", "event") VALUES ($1, $2, $3, $4, $5, 'ProfileCreatorWhitelisted') ON CONFLICT ("blockNumber", "logIndex") DO UPDATE
				SET "txHash" = excluded."txHash",
					"txIndex" = excluded."txIndex",
					"removed" = excluded."removed",
					"event" = excluded."event"
				RETURNING id ) INSERT INTO "ProfileCreatorWhitelisted" ("event_id", "profileCreator", "whitelisted", "timestamp") SELECT id, $6, $7, $8 FROM inserted_event
				ON CONFLICT ("event_id") DO UPDATE
				SET "profileCreator" = excluded."profileCreator",
					"whitelisted" = excluded."whitelisted",
					"timestamp" = excluded."timestamp"
				`,
				profileCreatorWhitelisted.Raw.BlockNumber, profileCreatorWhitelisted.Raw.TxHash.Hex(), profileCreatorWhitelisted.Raw.TxIndex, profileCreatorWhitelisted.Raw.Index, profileCreatorWhitelisted.Raw.Removed, profileCreatorWhitelisted.ProfileCreator.Hex(), profileCreatorWhitelisted.Whitelisted, profileCreatorWhitelisted.Timestamp)
			if err != nil {
				return fmt.Errorf("error saving ProfileCreatorWhitelisted event to database: %w", err)
			}
			return nil
		}),

	common.HexToHash("0xd5a5879cad33c830cc1432c1850107029a09c80c60e9bce1ecd08d24880bd46c"): NewEventProcessor(
		func(l types.Log) (interface{}, error) {
			return contractFilter.ParseProfileImageURISet(l)
		},
		func(db *pgxpool.Pool, e interface{}) error {
			profileImageURISet := e.(*contract.ContractProfileImageURISet)
			_, err := db.Exec(context.Background(), `WITH inserted_event AS ( INSERT INTO "Event" ("blockNumber", "txHash", "txIndex", "logIndex", "removed", "event") VALUES ($1, $2, $3, $4, $5, 'ProfileImageURISet') ON CONFLICT ("blockNumber", "logIndex") DO UPDATE
				SET "txHash" = excluded."txHash",
					"txIndex" = excluded."txIndex",
					"removed" = excluded."removed",
					"event" = excluded."event"
				RETURNING id ) INSERT INTO "ProfileImageURISet" ("event_id", "profileId", "imageURI", "timestamp") SELECT id, $6, $7, $8 FROM inserted_event
				ON CONFLICT ("event_id") DO UPDATE
				SET "profileId" = excluded."profileId",
					"imageURI" = excluded."imageURI",
					"timestamp" = excluded."timestamp"
				`,
				profileImageURISet.Raw.BlockNumber, profileImageURISet.Raw.TxHash.Hex(), profileImageURISet.Raw.TxIndex, profileImageURISet.Raw.Index, profileImageURISet.Raw.Removed, profileImageURISet.ProfileId, profileImageURISet.ImageURI, profileImageURISet.Timestamp)
			if err != nil {
				return fmt.Errorf("error saving ProfileImageURISet event to database: %w", err)
			}
			return nil
		}),

	common.HexToHash("0xf901a8b3832914a45999dd4c425fbe42eb4182724d394100401e633d9f0b286a"): NewEventProcessor(
		func(l types.Log) (interface{}, error) {
			return contractFilter.ParseProfileMetadataSet(l)
		},
		func(db *pgxpool.Pool, e interface{}) error {
			profileMetadataSet := e.(*contract.ContractProfileMetadataSet)
			_, err := db.Exec(context.Background(), `WITH inserted_event AS ( INSERT INTO "Event" ("blockNumber", "txHash", "txIndex", "logIndex", "removed", "event") VALUES ($1, $2, $3, $4, $5, 'ProfileMetadataSet') ON CONFLICT ("blockNumber", "logIndex") DO UPDATE
				SET "txHash" = excluded."txHash",
					"txIndex" = excluded."txIndex",
					"removed" = excluded."removed",
					"event" = excluded."event"
				RETURNING id ) INSERT INTO "ProfileMetadataSet" ("event_id", "profileId", "metadata", "timestamp") SELECT id, $6, $7, $8 FROM inserted_event
				ON CONFLICT ("event_id") DO UPDATE
				SET "profileId" = excluded."profileId",
					"metadata" = excluded."metadata",
					"timestamp" = excluded."timestamp"
				`,
				profileMetadataSet.Raw.BlockNumber, profileMetadataSet.Raw.TxHash.Hex(), profileMetadataSet.Raw.TxIndex, profileMetadataSet.Raw.Index, profileMetadataSet.Raw.Removed, profileMetadataSet.ProfileId, profileMetadataSet.Metadata, profileMetadataSet.Timestamp)
			if err != nil {
				return fmt.Errorf("error saving ProfileMetadataSet event to database: %w", err)
			}
			return nil
		}),

	common.HexToHash("0x37872a053ef20cb52defb7c9ec20e1a87cb8dd5098ac9e76a144be263dfef572"): NewEventProcessor(
		func(l types.Log) (interface{}, error) {
			return contractFilter.ParseReferenceModuleWhitelisted(l)
		},
		func(db *pgxpool.Pool, e interface{}) error {
			referenceModuleWhitelisted := e.(*contract.ContractReferenceModuleWhitelisted)
			_, err := db.Exec(context.Background(), `WITH inserted_event AS ( INSERT INTO "Event" ("blockNumber", "txHash", "txIndex", "logIndex", "removed", "event") VALUES ($1, $2, $3, $4, $5, 'ReferenceModuleWhitelisted') ON CONFLICT ("blockNumber", "logIndex") DO UPDATE
				SET "txHash" = excluded."txHash",
					"txIndex" = excluded."txIndex",
					"removed" = excluded."removed",
					"event" = excluded."event"
				RETURNING id ) INSERT INTO "ReferenceModuleWhitelisted" ("event_id", "referenceModule", "whitelisted", "timestamp") SELECT id, $6, $7, $8 FROM inserted_event
				ON CONFLICT ("event_id") DO UPDATE
				SET "referenceModule" = excluded."referenceModule",
					"whitelisted" = excluded."whitelisted",
					"timestamp" = excluded."timestamp"
				`,
				referenceModuleWhitelisted.Raw.BlockNumber, referenceModuleWhitelisted.Raw.TxHash.Hex(), referenceModuleWhitelisted.Raw.TxIndex, referenceModuleWhitelisted.Raw.Index, referenceModuleWhitelisted.Raw.Removed, referenceModuleWhitelisted.ReferenceModule.Hex(), referenceModuleWhitelisted.Whitelisted, referenceModuleWhitelisted.Timestamp)
			if err != nil {
				return fmt.Errorf("error saving ReferenceModuleWhitelisted event to database: %w", err)
			}
			return nil
		}),

	common.HexToHash("0xa2f9a1499fc1f9b7796d21fe5761290ccb7e0ef6ccf35fa58b668f304a62a1ca"): NewEventProcessor(
		func(l types.Log) (interface{}, error) {
			return contractFilter.ParseStateSet(l)
		},
		func(db *pgxpool.Pool, e interface{}) error {
			stateSet := e.(*contract.ContractStateSet)
			_, err := db.Exec(context.Background(), `WITH inserted_event AS ( INSERT INTO "Event" ("blockNumber", "txHash", "txIndex", "logIndex", "removed", "event") VALUES ($1, $2, $3, $4, $5, 'StateSet') ON CONFLICT ("blockNumber", "logIndex") DO UPDATE
				SET "txHash" = excluded."txHash",
					"txIndex" = excluded."txIndex",
					"removed" = excluded."removed",
					"event" = excluded."event"
				RETURNING id ) INSERT INTO "StateSet" ("event_id", "caller", "prevState", "newState", "timestamp") SELECT id, $6, $7, $8, $9 FROM inserted_event
				ON CONFLICT ("event_id") DO UPDATE
				SET "caller" = excluded."caller",
					"prevState" = excluded."prevState",
					"newState" = excluded."newState",
					"timestamp" = excluded."timestamp"
				`,
				stateSet.Raw.BlockNumber, stateSet.Raw.TxHash.Hex(), stateSet.Raw.TxIndex, stateSet.Raw.Index, stateSet.Raw.Removed, stateSet.Caller.Hex(), stateSet.PrevState, stateSet.NewState, stateSet.Timestamp)
			if err != nil {
				return fmt.Errorf("error saving StateSet event to database: %w", err)
			}
			return nil
		}),
	common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"): NewEventProcessor(
		func(l types.Log) (interface{}, error) {
			return contractFilter.ParseTransfer(l)
		},
		func(db *pgxpool.Pool, e interface{}) error {
			transfer := e.(*contract.ContractTransfer)
			_, err := db.Exec(context.Background(), `WITH inserted_event AS ( INSERT INTO "Event" ("blockNumber", "txHash", "txIndex", "logIndex", "removed", "event") VALUES ($1, $2, $3, $4, $5, 'Transfer') ON CONFLICT ("blockNumber", "logIndex") DO UPDATE
				SET "txHash" = excluded."txHash",
					"txIndex" = excluded."txIndex",
					"removed" = excluded."removed",
					"event" = excluded."event"
				RETURNING id ) INSERT INTO "Transfer" ("event_id", "from", "to", "tokenId") SELECT id, $6, $7, $8 FROM inserted_event
				ON CONFLICT ("event_id") DO UPDATE
				SET "from" = excluded."from",
					"to" = excluded."to",
					"tokenId" = excluded."tokenId"
				`,
				transfer.Raw.BlockNumber, transfer.Raw.TxHash.Hex(), transfer.Raw.TxIndex, transfer.Raw.Index, transfer.Raw.Removed, transfer.From.Hex(), transfer.To.Hex(), transfer.TokenId)
			if err != nil {
				return fmt.Errorf("error saving Transfer event to database: %w", err)
			}
			return nil
		}),

	common.HexToHash("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925"): NewEventProcessor(
		func(l types.Log) (interface{}, error) {
			return contractFilter.ParseApproval(l)
		},
		func(db *pgxpool.Pool, e interface{}) error {
			approval := e.(*contract.ContractApproval)
			_, err := db.Exec(context.Background(), `WITH inserted_event AS ( INSERT INTO "Event" ("blockNumber", "txHash", "txIndex", "logIndex", "removed", "event") VALUES ($1, $2, $3, $4, $5, 'Approval') ON CONFLICT ("blockNumber", "logIndex") DO UPDATE
				SET "txHash" = excluded."txHash",
					"txIndex" = excluded."txIndex",
					"removed" = excluded."removed",
					"event" = excluded."event"
				RETURNING id ) INSERT INTO "Approval" ("event_id", "owner", "approved", "tokenId") SELECT id, $6, $7, $8 FROM inserted_event
				ON CONFLICT ("event_id") DO UPDATE
				SET "owner" = excluded."owner",
					"approved" = excluded."approved",
					"tokenId" = excluded."tokenId"
				`,
				approval.Raw.BlockNumber, approval.Raw.TxHash.Hex(), approval.Raw.TxIndex, approval.Raw.Index, approval.Raw.Removed, approval.Owner.Hex(), approval.Approved.Hex(), approval.TokenId)
			if err != nil {
				return fmt.Errorf("error saving Approval event to database: %w", err)
			}
			return nil
		}),

	common.HexToHash("0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31"): NewEventProcessor(
		func(l types.Log) (interface{}, error) {
			return contractFilter.ParseApprovalForAll(l)
		},
		func(db *pgxpool.Pool, e interface{}) error {
			approvalForAll := e.(*contract.ContractApprovalForAll)
			_, err := db.Exec(context.Background(), `WITH inserted_event AS ( INSERT INTO "Event" ("blockNumber", "txHash", "txIndex", "logIndex", "removed", "event") VALUES ($1, $2, $3, $4, $5, 'ApprovalForAll') ON CONFLICT ("blockNumber", "logIndex") DO UPDATE
				SET "txHash" = excluded."txHash",
					"txIndex" = excluded."txIndex",
					"removed" = excluded."removed",
					"event" = excluded."event"
				RETURNING id ) INSERT INTO "ApprovalForAll" ("event_id", "owner", "operator", "approved") SELECT id, $6, $7, $8 FROM inserted_event
				ON CONFLICT ("event_id") DO UPDATE
				SET "owner" = excluded."owner",
					"operator" = excluded."operator",
					"approved" = excluded."approved"
				`,
				approvalForAll.Raw.BlockNumber, approvalForAll.Raw.TxHash.Hex(), approvalForAll.Raw.TxIndex, approvalForAll.Raw.Index, approvalForAll.Raw.Removed, approvalForAll.Owner.Hex(), approvalForAll.Operator.Hex(), approvalForAll.Approved)
			if err != nil {
				return fmt.Errorf("error saving ApprovalForAll event to database: %w", err)
			}
			return nil
		}),
}
