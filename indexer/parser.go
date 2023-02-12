package main

import (
	"github.com/daoleno/lenscan/indexer/contract"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

var contractFilter, _ = contract.NewContractFilterer(common.Address{}, nil)

// EventParser represents a function that parses a log event
type EventParser func(types.Log) (interface{}, error)

// EventParsers is a map of event signatures to event parsers
var EventParsers = map[common.Hash]EventParser{
	common.HexToHash("0x414cd0b34676984f09a5f76ce9718d4062e50283abe0e7e274a9a5b4e0c99c30"): func(l types.Log) (interface{}, error) {
		return contractFilter.ParseBaseInitialized(l)
	},
	common.HexToHash("0x6cc19a794d6a439023150cd58748eed4353190c0bb060d2e6250e2df4a68b673"): func(l types.Log) (interface{}, error) {
		return contractFilter.ParseCollectModuleWhitelisted(l)
	},
	common.HexToHash("0x0b227b550ffed48af813b32e246f787e99581ee13206ba8f9d90d63615269b3f"): func(l types.Log) (interface{}, error) {
		return contractFilter.ParseCollectNFTDeployed(l)
	},
	common.HexToHash("0x898a2dec95856255977a0fb48cebc30051d50c0d8d33f93dea1e3ddb2e342442"): func(l types.Log) (interface{}, error) {
		return contractFilter.ParseCollectNFTInitialized(l)
	},
	common.HexToHash("0x68edb7ec2c37d21b3b72233960b487f2966f4ac82b7430d39f24d1f8d6f99106"): func(l types.Log) (interface{}, error) {
		return contractFilter.ParseCollectNFTTransferred(l)
	},
	common.HexToHash("0xed39bf0d9afa849610b901c9d7f4d00751ba20de2db023428065bec153833218"): func(l types.Log) (interface{}, error) {
		return contractFilter.ParseCollected(l)
	},
	common.HexToHash("0x7b4d1aa33773161799847429e4fbf29f56dbf1a3fe815f5070231cbfba402c37"): func(l types.Log) (interface{}, error) {
		return contractFilter.ParseCommentCreated(l)
	common.HexToHash("0x0afd7c479e8bc7dcdb856b3cc27d2332dfe1f018fde574ea124919ddcae8a933"): func(l types.Log) (interface{}, error) {
		return contractFilter.ParseDefaultProfileSet(l)
	},
	common.HexToHash("0x22baaec4952f35f59e45bd2ddb287e1ccc6d319375770c09428eb8f8d604e065"): func(l types.Log) (interface{}, error) {
		return contractFilter.ParseDispatcherSet(l)
	},
	common.HexToHash("0x676c0801b0f400762e958ee31cfbb10870e70786f6761f57c8647e766b0db3d9"): func(l types.Log) (interface{}, error) {
		return contractFilter.ParseEmergencyAdminSet(l)
	},
	common.HexToHash("0x4e84a529f4c627b5e787037d117873af1018768804cca3c7f0d47041fe2c89ed"): func(l types.Log) (interface{}, error) {
		return contractFilter.ParseFeeModuleBaseConstructed(l)
	},
	common.HexToHash("0x92d95e400932d129885e627b38b169cbb28443ffaaa282d0fba0cf8797721359"): func(l types.Log) (interface{}, error) {
		return contractFilter.ParseFollowModuleSet(l)
	},
	common.HexToHash("0x52c5b7889df9f12f84ec3da051e854e5876678370d8357959c23ef59dd6486df"): func(l types.Log) (interface{}, error) {
		return contractFilter.ParseFollowModuleWhitelisted(l)
	},
	common.HexToHash("0xd9a6070174f4ccca76ed4896432e9a090b16e07e8fe27f275f50b33500b98e52"): func(l types.Log) (interface{}, error) {
		return contractFilter.ParseFollowNFTDelegatedPowerChanged(l)
	},
	common.HexToHash("0x44403e38baed5e40df7f64ff8708b076c75a0dfda8380e75df5c36f11a476743"): func(l types.Log) (interface{}, error) {
		return contractFilter.ParseFollowNFTDeployed(l)
	},
	common.HexToHash("0xaec15127df11a6b562c87d31bcb8f4cd2f0cf57fb9b663d6334abf41fea94d95"): func(l types.Log) (interface{}, error) {
		return contractFilter.ParseFollowNFTInitialized(l)
	},
	common.HexToHash("0x4996ad2257e7db44908136c43128cc10ca988096f67dc6bb0bcee11d151368fb"): func(l types.Log) (interface{}, error) {
		return contractFilter.ParseFollowNFTTransferred(l)
	},
	common.HexToHash("0xe82886e1af6fcab5caef13815b22f51384e970c367a785f265d13860a7d6966d"): func(l types.Log) (interface{}, error) {
		return contractFilter.ParseFollowNFTURISet(l)
	},
	common.HexToHash("0x40487072dc56f384287d26fbe090f404143c2737d54632177451d1f74bd82c76"): func(l types.Log) (interface{}, error) {
		return contractFilter.ParseFollowed(l)
	},
	common.HexToHash("0xc67fc3972da5d6434ab7b796ba133c240d40ee4e69129963c5aa0f2a6f7c3ad6"): func(l types.Log) (interface{}, error) {
		return contractFilter.ParseFollowsApproved(l)
	},
	common.HexToHash("0x5538c80c8d3bee397d87a7d153f7f085bb12adf2fe25a026c7cc4e83d8c5f1d7"): func(l types.Log) (interface{}, error) {
		return contractFilter.ParseFollowsToggled(l)
	},
	common.HexToHash("0xe552a55455b740845a5c07ed445d1724142fc997b389835495a29b30cddc1ccd"): func(l types.Log) (interface{}, error) {
		return contractFilter.ParseGovernanceSet(l)
	},
	common.HexToHash("0x9ea5dedb85bd9da4e264ee5a39b7ba0982e5d4d035d55edfa98a36b00e770b5a"): func(l types.Log) (interface{}, error) {
		return contractFilter.ParseMirrorCreated(l)
	},
	common.HexToHash("0xf1a1fa6b64aa95186f5a1285e76198d0da80d9c5a88062641d447f1d7c54e56c"): func(l types.Log) (interface{}, error) {
		return contractFilter.ParseModuleBaseConstructed(l)
	},
	common.HexToHash("0x79c3cefc851fd6040f06af202c542818d9fb39bcddcb7a7e3f563b15300a2743"): func(l types.Log) (interface{}, error) {
		return contractFilter.ParseModuleGlobalsCurrencyWhitelisted(l)
	},
	common.HexToHash("0xbf538a2c0db3d440906b8179dd0394f68a65b0b1481da70ffee24e19dccee84c"): func(l types.Log) (interface{}, error) {
		return contractFilter.ParseModuleGlobalsGovernanceSet(l)
	},
	common.HexToHash("0xec936862e6bb897cd711a5f31825057583128c2a482f0a4c9a4e6c3fd7c023f4"): func(l types.Log) (interface{}, error) {
		return contractFilter.ParseModuleGlobalsTreasuryFeeSet(l)
	},
	common.HexToHash("0x3dfc53d6b49bfbc932b215ba515f0d0ab0e17aac17726fba48075f0c16c7ffe3"): func(l types.Log) (interface{}, error) {
		return contractFilter.ParseModuleGlobalsTreasurySet(l)
	},
	common.HexToHash("0xc672c38b4d26c3c978228e99164105280410b144af24dd3ed8e4f9d211d96a50"): func(l types.Log) (interface{}, error) {
		return contractFilter.ParsePostCreated(l)
	},
	common.HexToHash("0x4e14f57cff7910416f2ef43cf05019b5a97a313de71fec9344be11b9b88fed12"): func(l types.Log) (interface{}, error) {
		return contractFilter.ParseProfileCreated(l)
	},
	common.HexToHash("0x8f617843889b94892bd44852d36ca6a7f49ecf4350a01e7b68e22d80f4ed95bc"): func(l types.Log) (interface{}, error) {
		return contractFilter.ParseProfileCreatorWhitelisted(l)
	},
	common.HexToHash("0xd5a5879cad33c830cc1432c1850107029a09c80c60e9bce1ecd08d24880bd46c"): func(l types.Log) (interface{}, error) {
		return contractFilter.ParseProfileImageURISet(l)
	},
	common.HexToHash("0xf901a8b3832914a45999dd4c425fbe42eb4182724d394100401e633d9f0b286a"): func(l types.Log) (interface{}, error) {
		return contractFilter.ParseProfileMetadataSet(l)
	},
	common.HexToHash("0x37872a053ef20cb52defb7c9ec20e1a87cb8dd5098ac9e76a144be263dfef572"): func(l types.Log) (interface{}, error) {
		return contractFilter.ParseReferenceModuleWhitelisted(l)
	},
	common.HexToHash("0xa2f9a1499fc1f9b7796d21fe5761290ccb7e0ef6ccf35fa58b668f304a62a1ca"): func(l types.Log) (interface{}, error) {
		return contractFilter.ParseStateSet(l)
	},
}
