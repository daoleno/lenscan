package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"reflect"

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
	common.HexToHash("0x414cd0b34676984f09a5f76ce9718d4062e50283abe0e7e274a9a5b4e0c99c30"): NewEventProcessor(
		func(l types.Log) (interface{}, error) {
			return contractFilter.ParseBaseInitialized(l)
		},
		func(db *pgxpool.Pool, e interface{}) error {
			baseInitialized := e.(*contract.ContractBaseInitialized)
			return saveEventToDB(db, "BaseInitialized", e, baseInitialized.Raw)
		},
	),
	common.HexToHash("0x6cc19a794d6a439023150cd58748eed4353190c0bb060d2e6250e2df4a68b673"): NewEventProcessor(
		func(l types.Log) (interface{}, error) {
			return contractFilter.ParseCollectModuleWhitelisted(l)
		},
		func(db *pgxpool.Pool, e interface{}) error {
			collectModuleWhitelisted := e.(*contract.ContractCollectModuleWhitelisted)
			return saveEventToDB(db, "CollectModuleWhitelisted", e, collectModuleWhitelisted.Raw)
		},
	),
	common.HexToHash("0x0b227b550ffed48af813b32e246f787e99581ee13206ba8f9d90d63615269b3f"): NewEventProcessor(
		func(l types.Log) (interface{}, error) {
			return contractFilter.ParseCollectNFTDeployed(l)
		},
		func(db *pgxpool.Pool, e interface{}) error {
			collectNFTDeployed := e.(*contract.ContractCollectNFTDeployed)
			return saveEventToDB(db, "CollectNFTDeployed", e, collectNFTDeployed.Raw)
		},
	),
	common.HexToHash("0x898a2dec95856255977a0fb48cebc30051d50c0d8d33f93dea1e3ddb2e342442"): NewEventProcessor(
		func(l types.Log) (interface{}, error) {
			return contractFilter.ParseCollectNFTInitialized(l)
		},
		func(db *pgxpool.Pool, e interface{}) error {
			collectNFTInitialized := e.(*contract.ContractCollectNFTInitialized)
			return saveEventToDB(db, "CollectNFTInitialized", e, collectNFTInitialized.Raw)
		},
	),
	common.HexToHash("0x68edb7ec2c37d21b3b72233960b487f2966f4ac82b7430d39f24d1f8d6f99106"): NewEventProcessor(
		func(l types.Log) (interface{}, error) {
			return contractFilter.ParseCollectNFTTransferred(l)
		},
		func(db *pgxpool.Pool, e interface{}) error {
			collectNFTTransferred := e.(*contract.ContractCollectNFTTransferred)
			return saveEventToDB(db, "CollectNFTTransferred", e, collectNFTTransferred.Raw)
		},
	),

	common.HexToHash("0xed39bf0d9afa849610b901c9d7f4d00751ba20de2db023428065bec153833218"): NewEventProcessor(
		func(l types.Log) (interface{}, error) {
			return contractFilter.ParseCollected(l)
		},
		func(db *pgxpool.Pool, e interface{}) error {
			collected := e.(*contract.ContractCollected)
			return saveEventToDB(db, "Collected", e, collected.Raw)
		},
	),
	common.HexToHash("0x7b4d1aa33773161799847429e4fbf29f56dbf1a3fe815f5070231cbfba402c37"): NewEventProcessor(
		func(l types.Log) (interface{}, error) {
			return contractFilter.ParseCommentCreated(l)
		},
		func(db *pgxpool.Pool, e interface{}) error {
			commentCreated := e.(*contract.ContractCommentCreated)
			return saveEventToDB(db, "CommentCreated", e, commentCreated.Raw)
		},
	),
	common.HexToHash("0x0afd7c479e8bc7dcdb856b3cc27d2332dfe1f018fde574ea124919ddcae8a933"): NewEventProcessor(
		func(l types.Log) (interface{}, error) {
			return contractFilter.ParseDefaultProfileSet(l)
		},
		func(db *pgxpool.Pool, e interface{}) error {
			defaultProfileSet := e.(*contract.ContractDefaultProfileSet)
			return saveEventToDB(db, "DefaultProfileSet", e, defaultProfileSet.Raw)
		},
	),

	common.HexToHash("0x22baaec4952f35f59e45bd2ddb287e1ccc6d319375770c09428eb8f8d604e065"): NewEventProcessor(
		func(l types.Log) (interface{}, error) {
			return contractFilter.ParseDispatcherSet(l)
		},
		func(db *pgxpool.Pool, e interface{}) error {
			dispatcherSet := e.(*contract.ContractDispatcherSet)
			return saveEventToDB(db, "DispatcherSet", e, dispatcherSet.Raw)
		},
	),

	common.HexToHash("0x676c0801b0f400762e958ee31cfbb10870e70786f6761f57c8647e766b0db3d9"): NewEventProcessor(
		func(l types.Log) (interface{}, error) { return contractFilter.ParseEmergencyAdminSet(l) },
		func(db *pgxpool.Pool, e interface{}) error {
			emergencyAdminSet := e.(*contract.ContractEmergencyAdminSet)
			return saveEventToDB(db, "EmergencyAdminSet", e, emergencyAdminSet.Raw)
		},
	),

	common.HexToHash("0x4e84a529f4c627b5e787037d117873af1018768804cca3c7f0d47041fe2c89ed"): NewEventProcessor(
		func(l types.Log) (interface{}, error) {
			return contractFilter.ParseFeeModuleBaseConstructed(l)
		},
		func(db *pgxpool.Pool, e interface{}) error {
			feeModuleBaseConstructed := e.(*contract.ContractFeeModuleBaseConstructed)
			return saveEventToDB(db, "FeeModuleBaseConstructed", e, feeModuleBaseConstructed.Raw)
		},
	),

	common.HexToHash("0x92d95e400932d129885e627b38b169cbb28443ffaaa282d0fba0cf8797721359"): NewEventProcessor(
		func(l types.Log) (interface{}, error) {
			return contractFilter.ParseFollowModuleSet(l)
		},
		func(db *pgxpool.Pool, e interface{}) error {
			followModuleSet := e.(*contract.ContractFollowModuleSet)
			return saveEventToDB(db, "FollowModuleSet", e, followModuleSet.Raw)
		},
	),

	common.HexToHash("0x52c5b7889df9f12f84ec3da051e854e5876678370d8357959c23ef59dd6486df"): NewEventProcessor(
		func(l types.Log) (interface{}, error) {
			return contractFilter.ParseFollowModuleWhitelisted(l)
		},
		func(db *pgxpool.Pool, e interface{}) error {
			followModuleWhitelisted := e.(*contract.ContractFollowModuleWhitelisted)
			return saveEventToDB(db, "FollowModuleWhitelisted", e, followModuleWhitelisted.Raw)
		},
	),

	common.HexToHash("0xd9a6070174f4ccca76ed4896432e9a090b16e07e8fe27f275f50b33500b98e52"): NewEventProcessor(
		func(l types.Log) (interface{}, error) {
			return contractFilter.ParseFollowNFTDelegatedPowerChanged(l)
		},
		func(db *pgxpool.Pool, e interface{}) error {
			followNFTDelegatedPowerChanged := e.(*contract.ContractFollowNFTDelegatedPowerChanged)
			return saveEventToDB(db, "FollowNFTDelegatedPowerChanged", e, followNFTDelegatedPowerChanged.Raw)
		},
	),

	common.HexToHash("0x44403e38baed5e40df7f64ff8708b076c75a0dfda8380e75df5c36f11a476743"): NewEventProcessor(
		func(l types.Log) (interface{}, error) {
			return contractFilter.ParseFollowNFTDeployed(l)
		},
		func(db *pgxpool.Pool, e interface{}) error {
			followNFTDeployed := e.(*contract.ContractFollowNFTDeployed)
			return saveEventToDB(db, "FollowNFTDeployed", e, followNFTDeployed.Raw)
		},
	),

	common.HexToHash("0xaec15127df11a6b562c87d31bcb8f4cd2f0cf57fb9b663d6334abf41fea94d95"): NewEventProcessor(
		func(l types.Log) (interface{}, error) {
			return contractFilter.ParseFollowNFTInitialized(l)
		},
		func(db *pgxpool.Pool, e interface{}) error {
			followNFTInitialized := e.(*contract.ContractFollowNFTInitialized)
			return saveEventToDB(db, "FollowNFTInitialized", e, followNFTInitialized.Raw)
		},
	),

	common.HexToHash("0x4996ad2257e7db44908136c43128cc10ca988096f67dc6bb0bcee11d151368fb"): NewEventProcessor(
		func(l types.Log) (interface{}, error) {
			return contractFilter.ParseFollowNFTTransferred(l)
		},
		func(db *pgxpool.Pool, e interface{}) error {
			followNFTTransferred := e.(*contract.ContractFollowNFTTransferred)
			return saveEventToDB(db, "FollowNFTTransferred", e, followNFTTransferred.Raw)
		},
	),

	common.HexToHash("0xe82886e1af6fcab5caef13815b22f51384e970c367a785f265d13860a7d6966d"): NewEventProcessor(
		func(l types.Log) (interface{}, error) {
			return contractFilter.ParseFollowNFTURISet(l)
		},
		func(db *pgxpool.Pool, e interface{}) error {
			followNFTURISet := e.(*contract.ContractFollowNFTURISet)
			return saveEventToDB(db, "FollowNFTURISet", e, followNFTURISet.Raw)
		}),

	common.HexToHash("0x40487072dc56f384287d26fbe090f404143c2737d54632177451d1f74bd82c76"): NewEventProcessor(
		func(l types.Log) (interface{}, error) {
			return contractFilter.ParseFollowed(l)
		},
		func(db *pgxpool.Pool, e interface{}) error {
			followed := e.(*contract.ContractFollowed)
			return saveEventToDB(db, "Followed", e, followed.Raw)
		}),

	common.HexToHash("0xc67fc3972da5d6434ab7b796ba133c240d40ee4e69129963c5aa0f2a6f7c3ad6"): NewEventProcessor(
		func(l types.Log) (interface{}, error) {
			return contractFilter.ParseFollowsApproved(l)
		},
		func(db *pgxpool.Pool, e interface{}) error {
			followsApproved := e.(*contract.ContractFollowsApproved)
			return saveEventToDB(db, "FollowsApproved", e, followsApproved.Raw)
		}),

	common.HexToHash("0x5538c80c8d3bee397d87a7d153f7f085bb12adf2fe25a026c7cc4e83d8c5f1d7"): NewEventProcessor(
		func(l types.Log) (interface{}, error) {
			return contractFilter.ParseFollowsToggled(l)
		},
		func(db *pgxpool.Pool, e interface{}) error {
			followsToggled := e.(*contract.ContractFollowsToggled)
			return saveEventToDB(db, "FollowsToggled", e, followsToggled.Raw)
		}),

	common.HexToHash("0xe552a55455b740845a5c07ed445d1724142fc997b389835495a29b30cddc1ccd"): NewEventProcessor(
		func(l types.Log) (interface{}, error) {
			return contractFilter.ParseGovernanceSet(l)
		},
		func(db *pgxpool.Pool, e interface{}) error {
			governanceSet := e.(*contract.ContractGovernanceSet)
			return saveEventToDB(db, "GovernanceSet", e, governanceSet.Raw)
		}),

	common.HexToHash("0x9ea5dedb85bd9da4e264ee5a39b7ba0982e5d4d035d55edfa98a36b00e770b5a"): NewEventProcessor(
		func(l types.Log) (interface{}, error) {
			return contractFilter.ParseMirrorCreated(l)
		},
		func(db *pgxpool.Pool, e interface{}) error {
			mirrorCreated := e.(*contract.ContractMirrorCreated)
			return saveEventToDB(db, "MirrorCreated", e, mirrorCreated.Raw)
		}),

	common.HexToHash("0xf1a1fa6b64aa95186f5a1285e76198d0da80d9c5a88062641d447f1d7c54e56c"): NewEventProcessor(
		func(l types.Log) (interface{}, error) {
			return contractFilter.ParseModuleBaseConstructed(l)
		},
		func(db *pgxpool.Pool, e interface{}) error {
			moduleBaseConstructed := e.(*contract.ContractModuleBaseConstructed)
			return saveEventToDB(db, "ModuleBaseConstructed", e, moduleBaseConstructed.Raw)
		}),

	common.HexToHash("0x79c3cefc851fd6040f06af202c542818d9fb39bcddcb7a7e3f563b15300a2743"): NewEventProcessor(
		func(l types.Log) (interface{}, error) {
			return contractFilter.ParseModuleGlobalsCurrencyWhitelisted(l)
		},
		func(db *pgxpool.Pool, e interface{}) error {
			moduleGlobalsCurrencyWhitelisted := e.(*contract.ContractModuleGlobalsCurrencyWhitelisted)
			return saveEventToDB(db, "ModuleGlobalsCurrencyWhitelisted", e, moduleGlobalsCurrencyWhitelisted.Raw)
		}),

	common.HexToHash("0xbf538a2c0db3d440906b8179dd0394f68a65b0b1481da70ffee24e19dccee84c"): NewEventProcessor(
		func(l types.Log) (interface{}, error) {
			return contractFilter.ParseModuleGlobalsGovernanceSet(l)
		},
		func(db *pgxpool.Pool, e interface{}) error {
			governanceSet := e.(*contract.ContractModuleGlobalsGovernanceSet)
			return saveEventToDB(db, "ModuleGlobalsGovernanceSet", e, governanceSet.Raw)
		}),

	common.HexToHash("0xec936862e6bb897cd711a5f31825057583128c2a482f0a4c9a4e6c3fd7c023f4"): NewEventProcessor(
		func(l types.Log) (interface{}, error) {
			return contractFilter.ParseModuleGlobalsTreasuryFeeSet(l)
		},
		func(db *pgxpool.Pool, e interface{}) error {
			moduleGlobalsTreasuryFeeSet := e.(*contract.ContractModuleGlobalsTreasuryFeeSet)
			return saveEventToDB(db, "ModuleGlobalsTreasuryFeeSet", e, moduleGlobalsTreasuryFeeSet.Raw)
		}),

	common.HexToHash("0x3dfc53d6b49bfbc932b215ba515f0d0ab0e17aac17726fba48075f0c16c7ffe3"): NewEventProcessor(
		func(l types.Log) (interface{}, error) {
			return contractFilter.ParseModuleGlobalsTreasurySet(l)
		},
		func(db *pgxpool.Pool, e interface{}) error {
			moduleGlobalsTreasurySet := e.(*contract.ContractModuleGlobalsTreasurySet)
			return saveEventToDB(db, "ModuleGlobalsTreasurySet", e, moduleGlobalsTreasurySet.Raw)
		}),

	common.HexToHash("0xc672c38b4d26c3c978228e99164105280410b144af24dd3ed8e4f9d211d96a50"): NewEventProcessor(
		func(l types.Log) (interface{}, error) {
			return contractFilter.ParsePostCreated(l)
		},
		func(db *pgxpool.Pool, e interface{}) error {
			postCreated := e.(*contract.ContractPostCreated)
			return saveEventToDB(db, "PostCreated", e, postCreated.Raw)
		}),

	common.HexToHash("0x4e14f57cff7910416f2ef43cf05019b5a97a313de71fec9344be11b9b88fed12"): NewEventProcessor(
		func(l types.Log) (interface{}, error) {
			return contractFilter.ParseProfileCreated(l)
		},
		func(db *pgxpool.Pool, e interface{}) error {
			profileCreated := e.(*contract.ContractProfileCreated)
			return saveEventToDB(db, "ProfileCreated", e, profileCreated.Raw)
		}),

	common.HexToHash("0x8f617843889b94892bd44852d36ca6a7f49ecf4350a01e7b68e22d80f4ed95bc"): NewEventProcessor(
		func(l types.Log) (interface{}, error) {
			return contractFilter.ParseProfileCreatorWhitelisted(l)
		},
		func(db *pgxpool.Pool, e interface{}) error {
			profileCreatorWhitelisted := e.(*contract.ContractProfileCreatorWhitelisted)
			return saveEventToDB(db, "ProfileCreatorWhitelisted", e, profileCreatorWhitelisted.Raw)
		}),

	common.HexToHash("0xd5a5879cad33c830cc1432c1850107029a09c80c60e9bce1ecd08d24880bd46c"): NewEventProcessor(
		func(l types.Log) (interface{}, error) {
			return contractFilter.ParseProfileImageURISet(l)
		},
		func(db *pgxpool.Pool, e interface{}) error {
			profileImageURISet := e.(*contract.ContractProfileImageURISet)
			return saveEventToDB(db, "ProfileImageURISet", e, profileImageURISet.Raw)
		}),

	common.HexToHash("0xf901a8b3832914a45999dd4c425fbe42eb4182724d394100401e633d9f0b286a"): NewEventProcessor(
		func(l types.Log) (interface{}, error) {
			return contractFilter.ParseProfileMetadataSet(l)
		},
		func(db *pgxpool.Pool, e interface{}) error {
			profileMetadataSet := e.(*contract.ContractProfileMetadataSet)
			return saveEventToDB(db, "ProfileMetadataSet", e, profileMetadataSet.Raw)
		}),

	common.HexToHash("0x37872a053ef20cb52defb7c9ec20e1a87cb8dd5098ac9e76a144be263dfef572"): NewEventProcessor(
		func(l types.Log) (interface{}, error) {
			return contractFilter.ParseReferenceModuleWhitelisted(l)
		},
		func(db *pgxpool.Pool, e interface{}) error {
			referenceModuleWhitelisted := e.(*contract.ContractReferenceModuleWhitelisted)
			return saveEventToDB(db, "ReferenceModuleWhitelisted", e, referenceModuleWhitelisted.Raw)
		}),

	common.HexToHash("0xa2f9a1499fc1f9b7796d21fe5761290ccb7e0ef6ccf35fa58b668f304a62a1ca"): NewEventProcessor(
		func(l types.Log) (interface{}, error) {
			return contractFilter.ParseStateSet(l)
		},
		func(db *pgxpool.Pool, e interface{}) error {
			stateSet := e.(*contract.ContractStateSet)
			return saveEventToDB(db, "StateSet", e, stateSet.Raw)
		}),
	common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"): NewEventProcessor(
		func(l types.Log) (interface{}, error) {
			return contractFilter.ParseTransfer(l)
		},
		func(db *pgxpool.Pool, e interface{}) error {
			transfer := e.(*contract.ContractTransfer)
			return saveEventToDB(db, "Transfer", e, transfer.Raw)
		}),

	common.HexToHash("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925"): NewEventProcessor(
		func(l types.Log) (interface{}, error) {
			return contractFilter.ParseApproval(l)
		},
		func(db *pgxpool.Pool, e interface{}) error {
			approval := e.(*contract.ContractApproval)
			return saveEventToDB(db, "Approval", e, approval.Raw)
		}),

	common.HexToHash("0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31"): NewEventProcessor(
		func(l types.Log) (interface{}, error) {
			return contractFilter.ParseApprovalForAll(l)
		},
		func(db *pgxpool.Pool, e interface{}) error {
			approvalForAll := e.(*contract.ContractApprovalForAll)
			return saveEventToDB(db, "ApprovalForAll", e, approvalForAll.Raw)
		}),
}

func saveEventToDB(db *pgxpool.Pool, eventType string, eventData interface{}, rawEvent types.Log) error {
	dataBytes, err := MarshalEvent(eventData)
	if err != nil {
		return fmt.Errorf("error marshalling data for %s event: %w", eventType, err)
	}

	ts := getTimestamp(rawEvent.BlockNumber, eventData)

	_, err = db.Exec(context.Background(), `
		INSERT INTO "Event_New" ("blockNumber", "txHash", "txIndex", "logIndex", "removed", "type", "data", "timestamp")
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		ON CONFLICT ("blockNumber", "logIndex") DO UPDATE
		SET "txHash" = excluded."txHash",
			"txIndex" = excluded."txIndex",
			"removed" = excluded."removed",
			"data" = excluded."data"
		`,
		rawEvent.BlockNumber, rawEvent.TxHash.Hex(), rawEvent.TxIndex, rawEvent.Index, rawEvent.Removed, eventType, dataBytes, ts)
	if err != nil {
		return fmt.Errorf("error saving %s event to database: %w", eventType, err)
	}
	return nil
}

func getTimestamp(blockNumber uint64, eventData interface{}) *big.Int {
	// get timestamp from eventData, most events have a Timestamp(*big.Int) field
	ts, err1 := getTimestampFromEvent(eventData)
	if ts != nil {
		return ts
	}

	// if timestamp is nil, get timestamp from block header
	tss, err2 := getTimestampFromBlock(blockNumber)
	if tss != 0 {
		return big.NewInt(int64(tss))
	}

	if err1 != nil || err2 != nil {
		log.Printf("error getting timestamp for block %d: %v %v", blockNumber, err1, err2)
	}

	return nil
}

func getTimestampFromEvent(v interface{}) (*big.Int, error) {
	// Get the type of the struct
	typ := reflect.TypeOf(v).Elem()

	// Iterate over the struct fields
	for i := 0; i < typ.NumField(); i++ {
		// Get the field type and value
		field := typ.Field(i)
		fieldValue := reflect.ValueOf(v).Elem().FieldByName(field.Name)

		// Check if the field is a Timestamp
		if field.Type.String() == "*big.Int" && field.Name == "Timestamp" {
			// Convert the field value to an interface{}
			return fieldValue.Interface().(*big.Int), nil
		}
	}

	return nil, fmt.Errorf("no Timestamp field found in event [%s]", typ.Name())
}

func MarshalEvent(v interface{}) ([]byte, error) {
	return MarshalIgnoreField(v, "Raw")
}

func MarshalIgnoreField(v interface{}, fieldName string) ([]byte, error) {
	// Get the type of the struct
	typ := reflect.TypeOf(v).Elem()

	// Create a map to hold the non-ignored fields
	objMap := make(map[string]interface{})

	// Iterate over the struct fields
	for i := 0; i < typ.NumField(); i++ {
		// Get the field type and value
		field := typ.Field(i)
		fieldValue := reflect.ValueOf(v).Elem().FieldByName(field.Name)

		// Ignore the specified field
		if field.Name == fieldName {
			continue
		}

		// Convert the field value to an interface{}
		objMap[field.Name] = fieldValue.Interface()
	}

	// Create a buffer to hold the JSON-encoded data
	buf := new(bytes.Buffer)

	// Create a JSON encoder that writes to the buffer
	enc := json.NewEncoder(buf)

	// Encode the object map as JSON and write it to the buffer
	err := enc.Encode(objMap)
	if err != nil {
		return nil, err
	}

	// Unmarshal the buffer data into a map
	var m map[string]interface{}
	err = json.Unmarshal(buf.Bytes(), &m)
	if err != nil {
		return nil, err
	}

	// Delete the specified field from the map
	delete(m, fieldName)

	// Marshal the remaining fields as JSON
	return json.Marshal(m)
}
