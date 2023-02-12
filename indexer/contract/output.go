// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"BaseInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collectModule\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"whitelisted\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"CollectModuleWhitelisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pubId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collectNFT\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"CollectNFTDeployed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pubId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"CollectNFTInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pubId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"collectNFTId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"CollectNFTTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collector\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pubId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rootProfileId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rootPubId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"collectModuleData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"Collected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pubId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"contentURI\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"profileIdPointed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pubIdPointed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"referenceModuleData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collectModule\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"collectModuleReturnData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"referenceModule\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"referenceModuleReturnData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"CommentCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"DefaultProfileSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dispatcher\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"DispatcherSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldEmergencyAdmin\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newEmergencyAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"EmergencyAdminSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"moduleGlobals\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"FeeModuleBaseConstructed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"followModule\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"followModuleReturnData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"FollowModuleSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"followModule\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"whitelisted\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"FollowModuleWhitelisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newPower\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"FollowNFTDelegatedPowerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"followNFT\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"FollowNFTDeployed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"FollowNFTInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"followNFTId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"FollowNFTTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"followNFTURI\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"FollowNFTURISet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"follower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"profileIds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"followModuleDatas\",\"type\":\"bytes[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"Followed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"bool[]\",\"name\":\"approved\",\"type\":\"bool[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"FollowsApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"profileIds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"bool[]\",\"name\":\"enabled\",\"type\":\"bool[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"FollowsToggled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"prevGovernance\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newGovernance\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"GovernanceSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pubId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"profileIdPointed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pubIdPointed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"referenceModuleData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"referenceModule\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"referenceModuleReturnData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"MirrorCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"hub\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"ModuleBaseConstructed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"prevWhitelisted\",\"type\":\"bool\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"whitelisted\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"ModuleGlobalsCurrencyWhitelisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"prevGovernance\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newGovernance\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"ModuleGlobalsGovernanceSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"prevTreasuryFee\",\"type\":\"uint16\"},{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"newTreasuryFee\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"ModuleGlobalsTreasuryFeeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"prevTreasury\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newTreasury\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"ModuleGlobalsTreasurySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pubId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"contentURI\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collectModule\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"collectModuleReturnData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"referenceModule\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"referenceModuleReturnData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"PostCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"handle\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"imageURI\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"followModule\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"followModuleReturnData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"followNFTURI\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"ProfileCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"profileCreator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"whitelisted\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"ProfileCreatorWhitelisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"imageURI\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"ProfileImageURISet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"ProfileMetadataSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"referenceModule\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"whitelisted\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"ReferenceModuleWhitelisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"enumDataTypes.ProtocolState\",\"name\":\"prevState\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"enumDataTypes.ProtocolState\",\"name\":\"newState\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"StateSet\",\"type\":\"event\"}]",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// ContractBaseInitializedIterator is returned from FilterBaseInitialized and is used to iterate over the raw logs and unpacked data for BaseInitialized events raised by the Contract contract.
type ContractBaseInitializedIterator struct {
	Event *ContractBaseInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractBaseInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractBaseInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractBaseInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractBaseInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractBaseInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractBaseInitialized represents a BaseInitialized event raised by the Contract contract.
type ContractBaseInitialized struct {
	Name      string
	Symbol    string
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBaseInitialized is a free log retrieval operation binding the contract event 0x414cd0b34676984f09a5f76ce9718d4062e50283abe0e7e274a9a5b4e0c99c30.
//
// Solidity: event BaseInitialized(string name, string symbol, uint256 timestamp)
func (_Contract *ContractFilterer) FilterBaseInitialized(opts *bind.FilterOpts) (*ContractBaseInitializedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "BaseInitialized")
	if err != nil {
		return nil, err
	}
	return &ContractBaseInitializedIterator{contract: _Contract.contract, event: "BaseInitialized", logs: logs, sub: sub}, nil
}

// WatchBaseInitialized is a free log subscription operation binding the contract event 0x414cd0b34676984f09a5f76ce9718d4062e50283abe0e7e274a9a5b4e0c99c30.
//
// Solidity: event BaseInitialized(string name, string symbol, uint256 timestamp)
func (_Contract *ContractFilterer) WatchBaseInitialized(opts *bind.WatchOpts, sink chan<- *ContractBaseInitialized) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "BaseInitialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractBaseInitialized)
				if err := _Contract.contract.UnpackLog(event, "BaseInitialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBaseInitialized is a log parse operation binding the contract event 0x414cd0b34676984f09a5f76ce9718d4062e50283abe0e7e274a9a5b4e0c99c30.
//
// Solidity: event BaseInitialized(string name, string symbol, uint256 timestamp)
func (_Contract *ContractFilterer) ParseBaseInitialized(log types.Log) (*ContractBaseInitialized, error) {
	event := new(ContractBaseInitialized)
	if err := _Contract.contract.UnpackLog(event, "BaseInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractCollectModuleWhitelistedIterator is returned from FilterCollectModuleWhitelisted and is used to iterate over the raw logs and unpacked data for CollectModuleWhitelisted events raised by the Contract contract.
type ContractCollectModuleWhitelistedIterator struct {
	Event *ContractCollectModuleWhitelisted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractCollectModuleWhitelistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractCollectModuleWhitelisted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractCollectModuleWhitelisted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractCollectModuleWhitelistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractCollectModuleWhitelistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractCollectModuleWhitelisted represents a CollectModuleWhitelisted event raised by the Contract contract.
type ContractCollectModuleWhitelisted struct {
	CollectModule common.Address
	Whitelisted   bool
	Timestamp     *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCollectModuleWhitelisted is a free log retrieval operation binding the contract event 0x6cc19a794d6a439023150cd58748eed4353190c0bb060d2e6250e2df4a68b673.
//
// Solidity: event CollectModuleWhitelisted(address indexed collectModule, bool indexed whitelisted, uint256 timestamp)
func (_Contract *ContractFilterer) FilterCollectModuleWhitelisted(opts *bind.FilterOpts, collectModule []common.Address, whitelisted []bool) (*ContractCollectModuleWhitelistedIterator, error) {

	var collectModuleRule []interface{}
	for _, collectModuleItem := range collectModule {
		collectModuleRule = append(collectModuleRule, collectModuleItem)
	}
	var whitelistedRule []interface{}
	for _, whitelistedItem := range whitelisted {
		whitelistedRule = append(whitelistedRule, whitelistedItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "CollectModuleWhitelisted", collectModuleRule, whitelistedRule)
	if err != nil {
		return nil, err
	}
	return &ContractCollectModuleWhitelistedIterator{contract: _Contract.contract, event: "CollectModuleWhitelisted", logs: logs, sub: sub}, nil
}

// WatchCollectModuleWhitelisted is a free log subscription operation binding the contract event 0x6cc19a794d6a439023150cd58748eed4353190c0bb060d2e6250e2df4a68b673.
//
// Solidity: event CollectModuleWhitelisted(address indexed collectModule, bool indexed whitelisted, uint256 timestamp)
func (_Contract *ContractFilterer) WatchCollectModuleWhitelisted(opts *bind.WatchOpts, sink chan<- *ContractCollectModuleWhitelisted, collectModule []common.Address, whitelisted []bool) (event.Subscription, error) {

	var collectModuleRule []interface{}
	for _, collectModuleItem := range collectModule {
		collectModuleRule = append(collectModuleRule, collectModuleItem)
	}
	var whitelistedRule []interface{}
	for _, whitelistedItem := range whitelisted {
		whitelistedRule = append(whitelistedRule, whitelistedItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "CollectModuleWhitelisted", collectModuleRule, whitelistedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractCollectModuleWhitelisted)
				if err := _Contract.contract.UnpackLog(event, "CollectModuleWhitelisted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCollectModuleWhitelisted is a log parse operation binding the contract event 0x6cc19a794d6a439023150cd58748eed4353190c0bb060d2e6250e2df4a68b673.
//
// Solidity: event CollectModuleWhitelisted(address indexed collectModule, bool indexed whitelisted, uint256 timestamp)
func (_Contract *ContractFilterer) ParseCollectModuleWhitelisted(log types.Log) (*ContractCollectModuleWhitelisted, error) {
	event := new(ContractCollectModuleWhitelisted)
	if err := _Contract.contract.UnpackLog(event, "CollectModuleWhitelisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractCollectNFTDeployedIterator is returned from FilterCollectNFTDeployed and is used to iterate over the raw logs and unpacked data for CollectNFTDeployed events raised by the Contract contract.
type ContractCollectNFTDeployedIterator struct {
	Event *ContractCollectNFTDeployed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractCollectNFTDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractCollectNFTDeployed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractCollectNFTDeployed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractCollectNFTDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractCollectNFTDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractCollectNFTDeployed represents a CollectNFTDeployed event raised by the Contract contract.
type ContractCollectNFTDeployed struct {
	ProfileId  *big.Int
	PubId      *big.Int
	CollectNFT common.Address
	Timestamp  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCollectNFTDeployed is a free log retrieval operation binding the contract event 0x0b227b550ffed48af813b32e246f787e99581ee13206ba8f9d90d63615269b3f.
//
// Solidity: event CollectNFTDeployed(uint256 indexed profileId, uint256 indexed pubId, address indexed collectNFT, uint256 timestamp)
func (_Contract *ContractFilterer) FilterCollectNFTDeployed(opts *bind.FilterOpts, profileId []*big.Int, pubId []*big.Int, collectNFT []common.Address) (*ContractCollectNFTDeployedIterator, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var pubIdRule []interface{}
	for _, pubIdItem := range pubId {
		pubIdRule = append(pubIdRule, pubIdItem)
	}
	var collectNFTRule []interface{}
	for _, collectNFTItem := range collectNFT {
		collectNFTRule = append(collectNFTRule, collectNFTItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "CollectNFTDeployed", profileIdRule, pubIdRule, collectNFTRule)
	if err != nil {
		return nil, err
	}
	return &ContractCollectNFTDeployedIterator{contract: _Contract.contract, event: "CollectNFTDeployed", logs: logs, sub: sub}, nil
}

// WatchCollectNFTDeployed is a free log subscription operation binding the contract event 0x0b227b550ffed48af813b32e246f787e99581ee13206ba8f9d90d63615269b3f.
//
// Solidity: event CollectNFTDeployed(uint256 indexed profileId, uint256 indexed pubId, address indexed collectNFT, uint256 timestamp)
func (_Contract *ContractFilterer) WatchCollectNFTDeployed(opts *bind.WatchOpts, sink chan<- *ContractCollectNFTDeployed, profileId []*big.Int, pubId []*big.Int, collectNFT []common.Address) (event.Subscription, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var pubIdRule []interface{}
	for _, pubIdItem := range pubId {
		pubIdRule = append(pubIdRule, pubIdItem)
	}
	var collectNFTRule []interface{}
	for _, collectNFTItem := range collectNFT {
		collectNFTRule = append(collectNFTRule, collectNFTItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "CollectNFTDeployed", profileIdRule, pubIdRule, collectNFTRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractCollectNFTDeployed)
				if err := _Contract.contract.UnpackLog(event, "CollectNFTDeployed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCollectNFTDeployed is a log parse operation binding the contract event 0x0b227b550ffed48af813b32e246f787e99581ee13206ba8f9d90d63615269b3f.
//
// Solidity: event CollectNFTDeployed(uint256 indexed profileId, uint256 indexed pubId, address indexed collectNFT, uint256 timestamp)
func (_Contract *ContractFilterer) ParseCollectNFTDeployed(log types.Log) (*ContractCollectNFTDeployed, error) {
	event := new(ContractCollectNFTDeployed)
	if err := _Contract.contract.UnpackLog(event, "CollectNFTDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractCollectNFTInitializedIterator is returned from FilterCollectNFTInitialized and is used to iterate over the raw logs and unpacked data for CollectNFTInitialized events raised by the Contract contract.
type ContractCollectNFTInitializedIterator struct {
	Event *ContractCollectNFTInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractCollectNFTInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractCollectNFTInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractCollectNFTInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractCollectNFTInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractCollectNFTInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractCollectNFTInitialized represents a CollectNFTInitialized event raised by the Contract contract.
type ContractCollectNFTInitialized struct {
	ProfileId *big.Int
	PubId     *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCollectNFTInitialized is a free log retrieval operation binding the contract event 0x898a2dec95856255977a0fb48cebc30051d50c0d8d33f93dea1e3ddb2e342442.
//
// Solidity: event CollectNFTInitialized(uint256 indexed profileId, uint256 indexed pubId, uint256 timestamp)
func (_Contract *ContractFilterer) FilterCollectNFTInitialized(opts *bind.FilterOpts, profileId []*big.Int, pubId []*big.Int) (*ContractCollectNFTInitializedIterator, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var pubIdRule []interface{}
	for _, pubIdItem := range pubId {
		pubIdRule = append(pubIdRule, pubIdItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "CollectNFTInitialized", profileIdRule, pubIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractCollectNFTInitializedIterator{contract: _Contract.contract, event: "CollectNFTInitialized", logs: logs, sub: sub}, nil
}

// WatchCollectNFTInitialized is a free log subscription operation binding the contract event 0x898a2dec95856255977a0fb48cebc30051d50c0d8d33f93dea1e3ddb2e342442.
//
// Solidity: event CollectNFTInitialized(uint256 indexed profileId, uint256 indexed pubId, uint256 timestamp)
func (_Contract *ContractFilterer) WatchCollectNFTInitialized(opts *bind.WatchOpts, sink chan<- *ContractCollectNFTInitialized, profileId []*big.Int, pubId []*big.Int) (event.Subscription, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var pubIdRule []interface{}
	for _, pubIdItem := range pubId {
		pubIdRule = append(pubIdRule, pubIdItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "CollectNFTInitialized", profileIdRule, pubIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractCollectNFTInitialized)
				if err := _Contract.contract.UnpackLog(event, "CollectNFTInitialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCollectNFTInitialized is a log parse operation binding the contract event 0x898a2dec95856255977a0fb48cebc30051d50c0d8d33f93dea1e3ddb2e342442.
//
// Solidity: event CollectNFTInitialized(uint256 indexed profileId, uint256 indexed pubId, uint256 timestamp)
func (_Contract *ContractFilterer) ParseCollectNFTInitialized(log types.Log) (*ContractCollectNFTInitialized, error) {
	event := new(ContractCollectNFTInitialized)
	if err := _Contract.contract.UnpackLog(event, "CollectNFTInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractCollectNFTTransferredIterator is returned from FilterCollectNFTTransferred and is used to iterate over the raw logs and unpacked data for CollectNFTTransferred events raised by the Contract contract.
type ContractCollectNFTTransferredIterator struct {
	Event *ContractCollectNFTTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractCollectNFTTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractCollectNFTTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractCollectNFTTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractCollectNFTTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractCollectNFTTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractCollectNFTTransferred represents a CollectNFTTransferred event raised by the Contract contract.
type ContractCollectNFTTransferred struct {
	ProfileId    *big.Int
	PubId        *big.Int
	CollectNFTId *big.Int
	From         common.Address
	To           common.Address
	Timestamp    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCollectNFTTransferred is a free log retrieval operation binding the contract event 0x68edb7ec2c37d21b3b72233960b487f2966f4ac82b7430d39f24d1f8d6f99106.
//
// Solidity: event CollectNFTTransferred(uint256 indexed profileId, uint256 indexed pubId, uint256 indexed collectNFTId, address from, address to, uint256 timestamp)
func (_Contract *ContractFilterer) FilterCollectNFTTransferred(opts *bind.FilterOpts, profileId []*big.Int, pubId []*big.Int, collectNFTId []*big.Int) (*ContractCollectNFTTransferredIterator, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var pubIdRule []interface{}
	for _, pubIdItem := range pubId {
		pubIdRule = append(pubIdRule, pubIdItem)
	}
	var collectNFTIdRule []interface{}
	for _, collectNFTIdItem := range collectNFTId {
		collectNFTIdRule = append(collectNFTIdRule, collectNFTIdItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "CollectNFTTransferred", profileIdRule, pubIdRule, collectNFTIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractCollectNFTTransferredIterator{contract: _Contract.contract, event: "CollectNFTTransferred", logs: logs, sub: sub}, nil
}

// WatchCollectNFTTransferred is a free log subscription operation binding the contract event 0x68edb7ec2c37d21b3b72233960b487f2966f4ac82b7430d39f24d1f8d6f99106.
//
// Solidity: event CollectNFTTransferred(uint256 indexed profileId, uint256 indexed pubId, uint256 indexed collectNFTId, address from, address to, uint256 timestamp)
func (_Contract *ContractFilterer) WatchCollectNFTTransferred(opts *bind.WatchOpts, sink chan<- *ContractCollectNFTTransferred, profileId []*big.Int, pubId []*big.Int, collectNFTId []*big.Int) (event.Subscription, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var pubIdRule []interface{}
	for _, pubIdItem := range pubId {
		pubIdRule = append(pubIdRule, pubIdItem)
	}
	var collectNFTIdRule []interface{}
	for _, collectNFTIdItem := range collectNFTId {
		collectNFTIdRule = append(collectNFTIdRule, collectNFTIdItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "CollectNFTTransferred", profileIdRule, pubIdRule, collectNFTIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractCollectNFTTransferred)
				if err := _Contract.contract.UnpackLog(event, "CollectNFTTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCollectNFTTransferred is a log parse operation binding the contract event 0x68edb7ec2c37d21b3b72233960b487f2966f4ac82b7430d39f24d1f8d6f99106.
//
// Solidity: event CollectNFTTransferred(uint256 indexed profileId, uint256 indexed pubId, uint256 indexed collectNFTId, address from, address to, uint256 timestamp)
func (_Contract *ContractFilterer) ParseCollectNFTTransferred(log types.Log) (*ContractCollectNFTTransferred, error) {
	event := new(ContractCollectNFTTransferred)
	if err := _Contract.contract.UnpackLog(event, "CollectNFTTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractCollectedIterator is returned from FilterCollected and is used to iterate over the raw logs and unpacked data for Collected events raised by the Contract contract.
type ContractCollectedIterator struct {
	Event *ContractCollected // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractCollectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractCollected)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractCollected)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractCollectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractCollectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractCollected represents a Collected event raised by the Contract contract.
type ContractCollected struct {
	Collector         common.Address
	ProfileId         *big.Int
	PubId             *big.Int
	RootProfileId     *big.Int
	RootPubId         *big.Int
	CollectModuleData []byte
	Timestamp         *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterCollected is a free log retrieval operation binding the contract event 0xed39bf0d9afa849610b901c9d7f4d00751ba20de2db023428065bec153833218.
//
// Solidity: event Collected(address indexed collector, uint256 indexed profileId, uint256 indexed pubId, uint256 rootProfileId, uint256 rootPubId, bytes collectModuleData, uint256 timestamp)
func (_Contract *ContractFilterer) FilterCollected(opts *bind.FilterOpts, collector []common.Address, profileId []*big.Int, pubId []*big.Int) (*ContractCollectedIterator, error) {

	var collectorRule []interface{}
	for _, collectorItem := range collector {
		collectorRule = append(collectorRule, collectorItem)
	}
	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var pubIdRule []interface{}
	for _, pubIdItem := range pubId {
		pubIdRule = append(pubIdRule, pubIdItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Collected", collectorRule, profileIdRule, pubIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractCollectedIterator{contract: _Contract.contract, event: "Collected", logs: logs, sub: sub}, nil
}

// WatchCollected is a free log subscription operation binding the contract event 0xed39bf0d9afa849610b901c9d7f4d00751ba20de2db023428065bec153833218.
//
// Solidity: event Collected(address indexed collector, uint256 indexed profileId, uint256 indexed pubId, uint256 rootProfileId, uint256 rootPubId, bytes collectModuleData, uint256 timestamp)
func (_Contract *ContractFilterer) WatchCollected(opts *bind.WatchOpts, sink chan<- *ContractCollected, collector []common.Address, profileId []*big.Int, pubId []*big.Int) (event.Subscription, error) {

	var collectorRule []interface{}
	for _, collectorItem := range collector {
		collectorRule = append(collectorRule, collectorItem)
	}
	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var pubIdRule []interface{}
	for _, pubIdItem := range pubId {
		pubIdRule = append(pubIdRule, pubIdItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Collected", collectorRule, profileIdRule, pubIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractCollected)
				if err := _Contract.contract.UnpackLog(event, "Collected", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCollected is a log parse operation binding the contract event 0xed39bf0d9afa849610b901c9d7f4d00751ba20de2db023428065bec153833218.
//
// Solidity: event Collected(address indexed collector, uint256 indexed profileId, uint256 indexed pubId, uint256 rootProfileId, uint256 rootPubId, bytes collectModuleData, uint256 timestamp)
func (_Contract *ContractFilterer) ParseCollected(log types.Log) (*ContractCollected, error) {
	event := new(ContractCollected)
	if err := _Contract.contract.UnpackLog(event, "Collected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractCommentCreatedIterator is returned from FilterCommentCreated and is used to iterate over the raw logs and unpacked data for CommentCreated events raised by the Contract contract.
type ContractCommentCreatedIterator struct {
	Event *ContractCommentCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractCommentCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractCommentCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractCommentCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractCommentCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractCommentCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractCommentCreated represents a CommentCreated event raised by the Contract contract.
type ContractCommentCreated struct {
	ProfileId                 *big.Int
	PubId                     *big.Int
	ContentURI                string
	ProfileIdPointed          *big.Int
	PubIdPointed              *big.Int
	ReferenceModuleData       []byte
	CollectModule             common.Address
	CollectModuleReturnData   []byte
	ReferenceModule           common.Address
	ReferenceModuleReturnData []byte
	Timestamp                 *big.Int
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterCommentCreated is a free log retrieval operation binding the contract event 0x7b4d1aa33773161799847429e4fbf29f56dbf1a3fe815f5070231cbfba402c37.
//
// Solidity: event CommentCreated(uint256 indexed profileId, uint256 indexed pubId, string contentURI, uint256 profileIdPointed, uint256 pubIdPointed, bytes referenceModuleData, address collectModule, bytes collectModuleReturnData, address referenceModule, bytes referenceModuleReturnData, uint256 timestamp)
func (_Contract *ContractFilterer) FilterCommentCreated(opts *bind.FilterOpts, profileId []*big.Int, pubId []*big.Int) (*ContractCommentCreatedIterator, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var pubIdRule []interface{}
	for _, pubIdItem := range pubId {
		pubIdRule = append(pubIdRule, pubIdItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "CommentCreated", profileIdRule, pubIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractCommentCreatedIterator{contract: _Contract.contract, event: "CommentCreated", logs: logs, sub: sub}, nil
}

// WatchCommentCreated is a free log subscription operation binding the contract event 0x7b4d1aa33773161799847429e4fbf29f56dbf1a3fe815f5070231cbfba402c37.
//
// Solidity: event CommentCreated(uint256 indexed profileId, uint256 indexed pubId, string contentURI, uint256 profileIdPointed, uint256 pubIdPointed, bytes referenceModuleData, address collectModule, bytes collectModuleReturnData, address referenceModule, bytes referenceModuleReturnData, uint256 timestamp)
func (_Contract *ContractFilterer) WatchCommentCreated(opts *bind.WatchOpts, sink chan<- *ContractCommentCreated, profileId []*big.Int, pubId []*big.Int) (event.Subscription, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var pubIdRule []interface{}
	for _, pubIdItem := range pubId {
		pubIdRule = append(pubIdRule, pubIdItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "CommentCreated", profileIdRule, pubIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractCommentCreated)
				if err := _Contract.contract.UnpackLog(event, "CommentCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCommentCreated is a log parse operation binding the contract event 0x7b4d1aa33773161799847429e4fbf29f56dbf1a3fe815f5070231cbfba402c37.
//
// Solidity: event CommentCreated(uint256 indexed profileId, uint256 indexed pubId, string contentURI, uint256 profileIdPointed, uint256 pubIdPointed, bytes referenceModuleData, address collectModule, bytes collectModuleReturnData, address referenceModule, bytes referenceModuleReturnData, uint256 timestamp)
func (_Contract *ContractFilterer) ParseCommentCreated(log types.Log) (*ContractCommentCreated, error) {
	event := new(ContractCommentCreated)
	if err := _Contract.contract.UnpackLog(event, "CommentCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDefaultProfileSetIterator is returned from FilterDefaultProfileSet and is used to iterate over the raw logs and unpacked data for DefaultProfileSet events raised by the Contract contract.
type ContractDefaultProfileSetIterator struct {
	Event *ContractDefaultProfileSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractDefaultProfileSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDefaultProfileSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractDefaultProfileSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractDefaultProfileSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDefaultProfileSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDefaultProfileSet represents a DefaultProfileSet event raised by the Contract contract.
type ContractDefaultProfileSet struct {
	Wallet    common.Address
	ProfileId *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDefaultProfileSet is a free log retrieval operation binding the contract event 0x0afd7c479e8bc7dcdb856b3cc27d2332dfe1f018fde574ea124919ddcae8a933.
//
// Solidity: event DefaultProfileSet(address indexed wallet, uint256 indexed profileId, uint256 timestamp)
func (_Contract *ContractFilterer) FilterDefaultProfileSet(opts *bind.FilterOpts, wallet []common.Address, profileId []*big.Int) (*ContractDefaultProfileSetIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "DefaultProfileSet", walletRule, profileIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractDefaultProfileSetIterator{contract: _Contract.contract, event: "DefaultProfileSet", logs: logs, sub: sub}, nil
}

// WatchDefaultProfileSet is a free log subscription operation binding the contract event 0x0afd7c479e8bc7dcdb856b3cc27d2332dfe1f018fde574ea124919ddcae8a933.
//
// Solidity: event DefaultProfileSet(address indexed wallet, uint256 indexed profileId, uint256 timestamp)
func (_Contract *ContractFilterer) WatchDefaultProfileSet(opts *bind.WatchOpts, sink chan<- *ContractDefaultProfileSet, wallet []common.Address, profileId []*big.Int) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "DefaultProfileSet", walletRule, profileIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDefaultProfileSet)
				if err := _Contract.contract.UnpackLog(event, "DefaultProfileSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDefaultProfileSet is a log parse operation binding the contract event 0x0afd7c479e8bc7dcdb856b3cc27d2332dfe1f018fde574ea124919ddcae8a933.
//
// Solidity: event DefaultProfileSet(address indexed wallet, uint256 indexed profileId, uint256 timestamp)
func (_Contract *ContractFilterer) ParseDefaultProfileSet(log types.Log) (*ContractDefaultProfileSet, error) {
	event := new(ContractDefaultProfileSet)
	if err := _Contract.contract.UnpackLog(event, "DefaultProfileSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDispatcherSetIterator is returned from FilterDispatcherSet and is used to iterate over the raw logs and unpacked data for DispatcherSet events raised by the Contract contract.
type ContractDispatcherSetIterator struct {
	Event *ContractDispatcherSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractDispatcherSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDispatcherSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractDispatcherSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractDispatcherSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDispatcherSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDispatcherSet represents a DispatcherSet event raised by the Contract contract.
type ContractDispatcherSet struct {
	ProfileId  *big.Int
	Dispatcher common.Address
	Timestamp  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDispatcherSet is a free log retrieval operation binding the contract event 0x22baaec4952f35f59e45bd2ddb287e1ccc6d319375770c09428eb8f8d604e065.
//
// Solidity: event DispatcherSet(uint256 indexed profileId, address indexed dispatcher, uint256 timestamp)
func (_Contract *ContractFilterer) FilterDispatcherSet(opts *bind.FilterOpts, profileId []*big.Int, dispatcher []common.Address) (*ContractDispatcherSetIterator, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var dispatcherRule []interface{}
	for _, dispatcherItem := range dispatcher {
		dispatcherRule = append(dispatcherRule, dispatcherItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "DispatcherSet", profileIdRule, dispatcherRule)
	if err != nil {
		return nil, err
	}
	return &ContractDispatcherSetIterator{contract: _Contract.contract, event: "DispatcherSet", logs: logs, sub: sub}, nil
}

// WatchDispatcherSet is a free log subscription operation binding the contract event 0x22baaec4952f35f59e45bd2ddb287e1ccc6d319375770c09428eb8f8d604e065.
//
// Solidity: event DispatcherSet(uint256 indexed profileId, address indexed dispatcher, uint256 timestamp)
func (_Contract *ContractFilterer) WatchDispatcherSet(opts *bind.WatchOpts, sink chan<- *ContractDispatcherSet, profileId []*big.Int, dispatcher []common.Address) (event.Subscription, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var dispatcherRule []interface{}
	for _, dispatcherItem := range dispatcher {
		dispatcherRule = append(dispatcherRule, dispatcherItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "DispatcherSet", profileIdRule, dispatcherRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDispatcherSet)
				if err := _Contract.contract.UnpackLog(event, "DispatcherSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDispatcherSet is a log parse operation binding the contract event 0x22baaec4952f35f59e45bd2ddb287e1ccc6d319375770c09428eb8f8d604e065.
//
// Solidity: event DispatcherSet(uint256 indexed profileId, address indexed dispatcher, uint256 timestamp)
func (_Contract *ContractFilterer) ParseDispatcherSet(log types.Log) (*ContractDispatcherSet, error) {
	event := new(ContractDispatcherSet)
	if err := _Contract.contract.UnpackLog(event, "DispatcherSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractEmergencyAdminSetIterator is returned from FilterEmergencyAdminSet and is used to iterate over the raw logs and unpacked data for EmergencyAdminSet events raised by the Contract contract.
type ContractEmergencyAdminSetIterator struct {
	Event *ContractEmergencyAdminSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractEmergencyAdminSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractEmergencyAdminSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractEmergencyAdminSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractEmergencyAdminSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractEmergencyAdminSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractEmergencyAdminSet represents a EmergencyAdminSet event raised by the Contract contract.
type ContractEmergencyAdminSet struct {
	Caller            common.Address
	OldEmergencyAdmin common.Address
	NewEmergencyAdmin common.Address
	Timestamp         *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterEmergencyAdminSet is a free log retrieval operation binding the contract event 0x676c0801b0f400762e958ee31cfbb10870e70786f6761f57c8647e766b0db3d9.
//
// Solidity: event EmergencyAdminSet(address indexed caller, address indexed oldEmergencyAdmin, address indexed newEmergencyAdmin, uint256 timestamp)
func (_Contract *ContractFilterer) FilterEmergencyAdminSet(opts *bind.FilterOpts, caller []common.Address, oldEmergencyAdmin []common.Address, newEmergencyAdmin []common.Address) (*ContractEmergencyAdminSetIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var oldEmergencyAdminRule []interface{}
	for _, oldEmergencyAdminItem := range oldEmergencyAdmin {
		oldEmergencyAdminRule = append(oldEmergencyAdminRule, oldEmergencyAdminItem)
	}
	var newEmergencyAdminRule []interface{}
	for _, newEmergencyAdminItem := range newEmergencyAdmin {
		newEmergencyAdminRule = append(newEmergencyAdminRule, newEmergencyAdminItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "EmergencyAdminSet", callerRule, oldEmergencyAdminRule, newEmergencyAdminRule)
	if err != nil {
		return nil, err
	}
	return &ContractEmergencyAdminSetIterator{contract: _Contract.contract, event: "EmergencyAdminSet", logs: logs, sub: sub}, nil
}

// WatchEmergencyAdminSet is a free log subscription operation binding the contract event 0x676c0801b0f400762e958ee31cfbb10870e70786f6761f57c8647e766b0db3d9.
//
// Solidity: event EmergencyAdminSet(address indexed caller, address indexed oldEmergencyAdmin, address indexed newEmergencyAdmin, uint256 timestamp)
func (_Contract *ContractFilterer) WatchEmergencyAdminSet(opts *bind.WatchOpts, sink chan<- *ContractEmergencyAdminSet, caller []common.Address, oldEmergencyAdmin []common.Address, newEmergencyAdmin []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var oldEmergencyAdminRule []interface{}
	for _, oldEmergencyAdminItem := range oldEmergencyAdmin {
		oldEmergencyAdminRule = append(oldEmergencyAdminRule, oldEmergencyAdminItem)
	}
	var newEmergencyAdminRule []interface{}
	for _, newEmergencyAdminItem := range newEmergencyAdmin {
		newEmergencyAdminRule = append(newEmergencyAdminRule, newEmergencyAdminItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "EmergencyAdminSet", callerRule, oldEmergencyAdminRule, newEmergencyAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractEmergencyAdminSet)
				if err := _Contract.contract.UnpackLog(event, "EmergencyAdminSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEmergencyAdminSet is a log parse operation binding the contract event 0x676c0801b0f400762e958ee31cfbb10870e70786f6761f57c8647e766b0db3d9.
//
// Solidity: event EmergencyAdminSet(address indexed caller, address indexed oldEmergencyAdmin, address indexed newEmergencyAdmin, uint256 timestamp)
func (_Contract *ContractFilterer) ParseEmergencyAdminSet(log types.Log) (*ContractEmergencyAdminSet, error) {
	event := new(ContractEmergencyAdminSet)
	if err := _Contract.contract.UnpackLog(event, "EmergencyAdminSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractFeeModuleBaseConstructedIterator is returned from FilterFeeModuleBaseConstructed and is used to iterate over the raw logs and unpacked data for FeeModuleBaseConstructed events raised by the Contract contract.
type ContractFeeModuleBaseConstructedIterator struct {
	Event *ContractFeeModuleBaseConstructed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractFeeModuleBaseConstructedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractFeeModuleBaseConstructed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractFeeModuleBaseConstructed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractFeeModuleBaseConstructedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractFeeModuleBaseConstructedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractFeeModuleBaseConstructed represents a FeeModuleBaseConstructed event raised by the Contract contract.
type ContractFeeModuleBaseConstructed struct {
	ModuleGlobals common.Address
	Timestamp     *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterFeeModuleBaseConstructed is a free log retrieval operation binding the contract event 0x4e84a529f4c627b5e787037d117873af1018768804cca3c7f0d47041fe2c89ed.
//
// Solidity: event FeeModuleBaseConstructed(address indexed moduleGlobals, uint256 timestamp)
func (_Contract *ContractFilterer) FilterFeeModuleBaseConstructed(opts *bind.FilterOpts, moduleGlobals []common.Address) (*ContractFeeModuleBaseConstructedIterator, error) {

	var moduleGlobalsRule []interface{}
	for _, moduleGlobalsItem := range moduleGlobals {
		moduleGlobalsRule = append(moduleGlobalsRule, moduleGlobalsItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "FeeModuleBaseConstructed", moduleGlobalsRule)
	if err != nil {
		return nil, err
	}
	return &ContractFeeModuleBaseConstructedIterator{contract: _Contract.contract, event: "FeeModuleBaseConstructed", logs: logs, sub: sub}, nil
}

// WatchFeeModuleBaseConstructed is a free log subscription operation binding the contract event 0x4e84a529f4c627b5e787037d117873af1018768804cca3c7f0d47041fe2c89ed.
//
// Solidity: event FeeModuleBaseConstructed(address indexed moduleGlobals, uint256 timestamp)
func (_Contract *ContractFilterer) WatchFeeModuleBaseConstructed(opts *bind.WatchOpts, sink chan<- *ContractFeeModuleBaseConstructed, moduleGlobals []common.Address) (event.Subscription, error) {

	var moduleGlobalsRule []interface{}
	for _, moduleGlobalsItem := range moduleGlobals {
		moduleGlobalsRule = append(moduleGlobalsRule, moduleGlobalsItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "FeeModuleBaseConstructed", moduleGlobalsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractFeeModuleBaseConstructed)
				if err := _Contract.contract.UnpackLog(event, "FeeModuleBaseConstructed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFeeModuleBaseConstructed is a log parse operation binding the contract event 0x4e84a529f4c627b5e787037d117873af1018768804cca3c7f0d47041fe2c89ed.
//
// Solidity: event FeeModuleBaseConstructed(address indexed moduleGlobals, uint256 timestamp)
func (_Contract *ContractFilterer) ParseFeeModuleBaseConstructed(log types.Log) (*ContractFeeModuleBaseConstructed, error) {
	event := new(ContractFeeModuleBaseConstructed)
	if err := _Contract.contract.UnpackLog(event, "FeeModuleBaseConstructed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractFollowModuleSetIterator is returned from FilterFollowModuleSet and is used to iterate over the raw logs and unpacked data for FollowModuleSet events raised by the Contract contract.
type ContractFollowModuleSetIterator struct {
	Event *ContractFollowModuleSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractFollowModuleSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractFollowModuleSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractFollowModuleSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractFollowModuleSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractFollowModuleSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractFollowModuleSet represents a FollowModuleSet event raised by the Contract contract.
type ContractFollowModuleSet struct {
	ProfileId              *big.Int
	FollowModule           common.Address
	FollowModuleReturnData []byte
	Timestamp              *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterFollowModuleSet is a free log retrieval operation binding the contract event 0x92d95e400932d129885e627b38b169cbb28443ffaaa282d0fba0cf8797721359.
//
// Solidity: event FollowModuleSet(uint256 indexed profileId, address followModule, bytes followModuleReturnData, uint256 timestamp)
func (_Contract *ContractFilterer) FilterFollowModuleSet(opts *bind.FilterOpts, profileId []*big.Int) (*ContractFollowModuleSetIterator, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "FollowModuleSet", profileIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractFollowModuleSetIterator{contract: _Contract.contract, event: "FollowModuleSet", logs: logs, sub: sub}, nil
}

// WatchFollowModuleSet is a free log subscription operation binding the contract event 0x92d95e400932d129885e627b38b169cbb28443ffaaa282d0fba0cf8797721359.
//
// Solidity: event FollowModuleSet(uint256 indexed profileId, address followModule, bytes followModuleReturnData, uint256 timestamp)
func (_Contract *ContractFilterer) WatchFollowModuleSet(opts *bind.WatchOpts, sink chan<- *ContractFollowModuleSet, profileId []*big.Int) (event.Subscription, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "FollowModuleSet", profileIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractFollowModuleSet)
				if err := _Contract.contract.UnpackLog(event, "FollowModuleSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFollowModuleSet is a log parse operation binding the contract event 0x92d95e400932d129885e627b38b169cbb28443ffaaa282d0fba0cf8797721359.
//
// Solidity: event FollowModuleSet(uint256 indexed profileId, address followModule, bytes followModuleReturnData, uint256 timestamp)
func (_Contract *ContractFilterer) ParseFollowModuleSet(log types.Log) (*ContractFollowModuleSet, error) {
	event := new(ContractFollowModuleSet)
	if err := _Contract.contract.UnpackLog(event, "FollowModuleSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractFollowModuleWhitelistedIterator is returned from FilterFollowModuleWhitelisted and is used to iterate over the raw logs and unpacked data for FollowModuleWhitelisted events raised by the Contract contract.
type ContractFollowModuleWhitelistedIterator struct {
	Event *ContractFollowModuleWhitelisted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractFollowModuleWhitelistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractFollowModuleWhitelisted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractFollowModuleWhitelisted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractFollowModuleWhitelistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractFollowModuleWhitelistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractFollowModuleWhitelisted represents a FollowModuleWhitelisted event raised by the Contract contract.
type ContractFollowModuleWhitelisted struct {
	FollowModule common.Address
	Whitelisted  bool
	Timestamp    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterFollowModuleWhitelisted is a free log retrieval operation binding the contract event 0x52c5b7889df9f12f84ec3da051e854e5876678370d8357959c23ef59dd6486df.
//
// Solidity: event FollowModuleWhitelisted(address indexed followModule, bool indexed whitelisted, uint256 timestamp)
func (_Contract *ContractFilterer) FilterFollowModuleWhitelisted(opts *bind.FilterOpts, followModule []common.Address, whitelisted []bool) (*ContractFollowModuleWhitelistedIterator, error) {

	var followModuleRule []interface{}
	for _, followModuleItem := range followModule {
		followModuleRule = append(followModuleRule, followModuleItem)
	}
	var whitelistedRule []interface{}
	for _, whitelistedItem := range whitelisted {
		whitelistedRule = append(whitelistedRule, whitelistedItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "FollowModuleWhitelisted", followModuleRule, whitelistedRule)
	if err != nil {
		return nil, err
	}
	return &ContractFollowModuleWhitelistedIterator{contract: _Contract.contract, event: "FollowModuleWhitelisted", logs: logs, sub: sub}, nil
}

// WatchFollowModuleWhitelisted is a free log subscription operation binding the contract event 0x52c5b7889df9f12f84ec3da051e854e5876678370d8357959c23ef59dd6486df.
//
// Solidity: event FollowModuleWhitelisted(address indexed followModule, bool indexed whitelisted, uint256 timestamp)
func (_Contract *ContractFilterer) WatchFollowModuleWhitelisted(opts *bind.WatchOpts, sink chan<- *ContractFollowModuleWhitelisted, followModule []common.Address, whitelisted []bool) (event.Subscription, error) {

	var followModuleRule []interface{}
	for _, followModuleItem := range followModule {
		followModuleRule = append(followModuleRule, followModuleItem)
	}
	var whitelistedRule []interface{}
	for _, whitelistedItem := range whitelisted {
		whitelistedRule = append(whitelistedRule, whitelistedItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "FollowModuleWhitelisted", followModuleRule, whitelistedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractFollowModuleWhitelisted)
				if err := _Contract.contract.UnpackLog(event, "FollowModuleWhitelisted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFollowModuleWhitelisted is a log parse operation binding the contract event 0x52c5b7889df9f12f84ec3da051e854e5876678370d8357959c23ef59dd6486df.
//
// Solidity: event FollowModuleWhitelisted(address indexed followModule, bool indexed whitelisted, uint256 timestamp)
func (_Contract *ContractFilterer) ParseFollowModuleWhitelisted(log types.Log) (*ContractFollowModuleWhitelisted, error) {
	event := new(ContractFollowModuleWhitelisted)
	if err := _Contract.contract.UnpackLog(event, "FollowModuleWhitelisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractFollowNFTDelegatedPowerChangedIterator is returned from FilterFollowNFTDelegatedPowerChanged and is used to iterate over the raw logs and unpacked data for FollowNFTDelegatedPowerChanged events raised by the Contract contract.
type ContractFollowNFTDelegatedPowerChangedIterator struct {
	Event *ContractFollowNFTDelegatedPowerChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractFollowNFTDelegatedPowerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractFollowNFTDelegatedPowerChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractFollowNFTDelegatedPowerChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractFollowNFTDelegatedPowerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractFollowNFTDelegatedPowerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractFollowNFTDelegatedPowerChanged represents a FollowNFTDelegatedPowerChanged event raised by the Contract contract.
type ContractFollowNFTDelegatedPowerChanged struct {
	Delegate  common.Address
	NewPower  *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFollowNFTDelegatedPowerChanged is a free log retrieval operation binding the contract event 0xd9a6070174f4ccca76ed4896432e9a090b16e07e8fe27f275f50b33500b98e52.
//
// Solidity: event FollowNFTDelegatedPowerChanged(address indexed delegate, uint256 indexed newPower, uint256 timestamp)
func (_Contract *ContractFilterer) FilterFollowNFTDelegatedPowerChanged(opts *bind.FilterOpts, delegate []common.Address, newPower []*big.Int) (*ContractFollowNFTDelegatedPowerChangedIterator, error) {

	var delegateRule []interface{}
	for _, delegateItem := range delegate {
		delegateRule = append(delegateRule, delegateItem)
	}
	var newPowerRule []interface{}
	for _, newPowerItem := range newPower {
		newPowerRule = append(newPowerRule, newPowerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "FollowNFTDelegatedPowerChanged", delegateRule, newPowerRule)
	if err != nil {
		return nil, err
	}
	return &ContractFollowNFTDelegatedPowerChangedIterator{contract: _Contract.contract, event: "FollowNFTDelegatedPowerChanged", logs: logs, sub: sub}, nil
}

// WatchFollowNFTDelegatedPowerChanged is a free log subscription operation binding the contract event 0xd9a6070174f4ccca76ed4896432e9a090b16e07e8fe27f275f50b33500b98e52.
//
// Solidity: event FollowNFTDelegatedPowerChanged(address indexed delegate, uint256 indexed newPower, uint256 timestamp)
func (_Contract *ContractFilterer) WatchFollowNFTDelegatedPowerChanged(opts *bind.WatchOpts, sink chan<- *ContractFollowNFTDelegatedPowerChanged, delegate []common.Address, newPower []*big.Int) (event.Subscription, error) {

	var delegateRule []interface{}
	for _, delegateItem := range delegate {
		delegateRule = append(delegateRule, delegateItem)
	}
	var newPowerRule []interface{}
	for _, newPowerItem := range newPower {
		newPowerRule = append(newPowerRule, newPowerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "FollowNFTDelegatedPowerChanged", delegateRule, newPowerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractFollowNFTDelegatedPowerChanged)
				if err := _Contract.contract.UnpackLog(event, "FollowNFTDelegatedPowerChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFollowNFTDelegatedPowerChanged is a log parse operation binding the contract event 0xd9a6070174f4ccca76ed4896432e9a090b16e07e8fe27f275f50b33500b98e52.
//
// Solidity: event FollowNFTDelegatedPowerChanged(address indexed delegate, uint256 indexed newPower, uint256 timestamp)
func (_Contract *ContractFilterer) ParseFollowNFTDelegatedPowerChanged(log types.Log) (*ContractFollowNFTDelegatedPowerChanged, error) {
	event := new(ContractFollowNFTDelegatedPowerChanged)
	if err := _Contract.contract.UnpackLog(event, "FollowNFTDelegatedPowerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractFollowNFTDeployedIterator is returned from FilterFollowNFTDeployed and is used to iterate over the raw logs and unpacked data for FollowNFTDeployed events raised by the Contract contract.
type ContractFollowNFTDeployedIterator struct {
	Event *ContractFollowNFTDeployed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractFollowNFTDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractFollowNFTDeployed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractFollowNFTDeployed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractFollowNFTDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractFollowNFTDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractFollowNFTDeployed represents a FollowNFTDeployed event raised by the Contract contract.
type ContractFollowNFTDeployed struct {
	ProfileId *big.Int
	FollowNFT common.Address
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFollowNFTDeployed is a free log retrieval operation binding the contract event 0x44403e38baed5e40df7f64ff8708b076c75a0dfda8380e75df5c36f11a476743.
//
// Solidity: event FollowNFTDeployed(uint256 indexed profileId, address indexed followNFT, uint256 timestamp)
func (_Contract *ContractFilterer) FilterFollowNFTDeployed(opts *bind.FilterOpts, profileId []*big.Int, followNFT []common.Address) (*ContractFollowNFTDeployedIterator, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var followNFTRule []interface{}
	for _, followNFTItem := range followNFT {
		followNFTRule = append(followNFTRule, followNFTItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "FollowNFTDeployed", profileIdRule, followNFTRule)
	if err != nil {
		return nil, err
	}
	return &ContractFollowNFTDeployedIterator{contract: _Contract.contract, event: "FollowNFTDeployed", logs: logs, sub: sub}, nil
}

// WatchFollowNFTDeployed is a free log subscription operation binding the contract event 0x44403e38baed5e40df7f64ff8708b076c75a0dfda8380e75df5c36f11a476743.
//
// Solidity: event FollowNFTDeployed(uint256 indexed profileId, address indexed followNFT, uint256 timestamp)
func (_Contract *ContractFilterer) WatchFollowNFTDeployed(opts *bind.WatchOpts, sink chan<- *ContractFollowNFTDeployed, profileId []*big.Int, followNFT []common.Address) (event.Subscription, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var followNFTRule []interface{}
	for _, followNFTItem := range followNFT {
		followNFTRule = append(followNFTRule, followNFTItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "FollowNFTDeployed", profileIdRule, followNFTRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractFollowNFTDeployed)
				if err := _Contract.contract.UnpackLog(event, "FollowNFTDeployed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFollowNFTDeployed is a log parse operation binding the contract event 0x44403e38baed5e40df7f64ff8708b076c75a0dfda8380e75df5c36f11a476743.
//
// Solidity: event FollowNFTDeployed(uint256 indexed profileId, address indexed followNFT, uint256 timestamp)
func (_Contract *ContractFilterer) ParseFollowNFTDeployed(log types.Log) (*ContractFollowNFTDeployed, error) {
	event := new(ContractFollowNFTDeployed)
	if err := _Contract.contract.UnpackLog(event, "FollowNFTDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractFollowNFTInitializedIterator is returned from FilterFollowNFTInitialized and is used to iterate over the raw logs and unpacked data for FollowNFTInitialized events raised by the Contract contract.
type ContractFollowNFTInitializedIterator struct {
	Event *ContractFollowNFTInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractFollowNFTInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractFollowNFTInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractFollowNFTInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractFollowNFTInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractFollowNFTInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractFollowNFTInitialized represents a FollowNFTInitialized event raised by the Contract contract.
type ContractFollowNFTInitialized struct {
	ProfileId *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFollowNFTInitialized is a free log retrieval operation binding the contract event 0xaec15127df11a6b562c87d31bcb8f4cd2f0cf57fb9b663d6334abf41fea94d95.
//
// Solidity: event FollowNFTInitialized(uint256 indexed profileId, uint256 timestamp)
func (_Contract *ContractFilterer) FilterFollowNFTInitialized(opts *bind.FilterOpts, profileId []*big.Int) (*ContractFollowNFTInitializedIterator, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "FollowNFTInitialized", profileIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractFollowNFTInitializedIterator{contract: _Contract.contract, event: "FollowNFTInitialized", logs: logs, sub: sub}, nil
}

// WatchFollowNFTInitialized is a free log subscription operation binding the contract event 0xaec15127df11a6b562c87d31bcb8f4cd2f0cf57fb9b663d6334abf41fea94d95.
//
// Solidity: event FollowNFTInitialized(uint256 indexed profileId, uint256 timestamp)
func (_Contract *ContractFilterer) WatchFollowNFTInitialized(opts *bind.WatchOpts, sink chan<- *ContractFollowNFTInitialized, profileId []*big.Int) (event.Subscription, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "FollowNFTInitialized", profileIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractFollowNFTInitialized)
				if err := _Contract.contract.UnpackLog(event, "FollowNFTInitialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFollowNFTInitialized is a log parse operation binding the contract event 0xaec15127df11a6b562c87d31bcb8f4cd2f0cf57fb9b663d6334abf41fea94d95.
//
// Solidity: event FollowNFTInitialized(uint256 indexed profileId, uint256 timestamp)
func (_Contract *ContractFilterer) ParseFollowNFTInitialized(log types.Log) (*ContractFollowNFTInitialized, error) {
	event := new(ContractFollowNFTInitialized)
	if err := _Contract.contract.UnpackLog(event, "FollowNFTInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractFollowNFTTransferredIterator is returned from FilterFollowNFTTransferred and is used to iterate over the raw logs and unpacked data for FollowNFTTransferred events raised by the Contract contract.
type ContractFollowNFTTransferredIterator struct {
	Event *ContractFollowNFTTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractFollowNFTTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractFollowNFTTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractFollowNFTTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractFollowNFTTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractFollowNFTTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractFollowNFTTransferred represents a FollowNFTTransferred event raised by the Contract contract.
type ContractFollowNFTTransferred struct {
	ProfileId   *big.Int
	FollowNFTId *big.Int
	From        common.Address
	To          common.Address
	Timestamp   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFollowNFTTransferred is a free log retrieval operation binding the contract event 0x4996ad2257e7db44908136c43128cc10ca988096f67dc6bb0bcee11d151368fb.
//
// Solidity: event FollowNFTTransferred(uint256 indexed profileId, uint256 indexed followNFTId, address from, address to, uint256 timestamp)
func (_Contract *ContractFilterer) FilterFollowNFTTransferred(opts *bind.FilterOpts, profileId []*big.Int, followNFTId []*big.Int) (*ContractFollowNFTTransferredIterator, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var followNFTIdRule []interface{}
	for _, followNFTIdItem := range followNFTId {
		followNFTIdRule = append(followNFTIdRule, followNFTIdItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "FollowNFTTransferred", profileIdRule, followNFTIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractFollowNFTTransferredIterator{contract: _Contract.contract, event: "FollowNFTTransferred", logs: logs, sub: sub}, nil
}

// WatchFollowNFTTransferred is a free log subscription operation binding the contract event 0x4996ad2257e7db44908136c43128cc10ca988096f67dc6bb0bcee11d151368fb.
//
// Solidity: event FollowNFTTransferred(uint256 indexed profileId, uint256 indexed followNFTId, address from, address to, uint256 timestamp)
func (_Contract *ContractFilterer) WatchFollowNFTTransferred(opts *bind.WatchOpts, sink chan<- *ContractFollowNFTTransferred, profileId []*big.Int, followNFTId []*big.Int) (event.Subscription, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var followNFTIdRule []interface{}
	for _, followNFTIdItem := range followNFTId {
		followNFTIdRule = append(followNFTIdRule, followNFTIdItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "FollowNFTTransferred", profileIdRule, followNFTIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractFollowNFTTransferred)
				if err := _Contract.contract.UnpackLog(event, "FollowNFTTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFollowNFTTransferred is a log parse operation binding the contract event 0x4996ad2257e7db44908136c43128cc10ca988096f67dc6bb0bcee11d151368fb.
//
// Solidity: event FollowNFTTransferred(uint256 indexed profileId, uint256 indexed followNFTId, address from, address to, uint256 timestamp)
func (_Contract *ContractFilterer) ParseFollowNFTTransferred(log types.Log) (*ContractFollowNFTTransferred, error) {
	event := new(ContractFollowNFTTransferred)
	if err := _Contract.contract.UnpackLog(event, "FollowNFTTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractFollowNFTURISetIterator is returned from FilterFollowNFTURISet and is used to iterate over the raw logs and unpacked data for FollowNFTURISet events raised by the Contract contract.
type ContractFollowNFTURISetIterator struct {
	Event *ContractFollowNFTURISet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractFollowNFTURISetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractFollowNFTURISet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractFollowNFTURISet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractFollowNFTURISetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractFollowNFTURISetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractFollowNFTURISet represents a FollowNFTURISet event raised by the Contract contract.
type ContractFollowNFTURISet struct {
	ProfileId    *big.Int
	FollowNFTURI string
	Timestamp    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterFollowNFTURISet is a free log retrieval operation binding the contract event 0xe82886e1af6fcab5caef13815b22f51384e970c367a785f265d13860a7d6966d.
//
// Solidity: event FollowNFTURISet(uint256 indexed profileId, string followNFTURI, uint256 timestamp)
func (_Contract *ContractFilterer) FilterFollowNFTURISet(opts *bind.FilterOpts, profileId []*big.Int) (*ContractFollowNFTURISetIterator, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "FollowNFTURISet", profileIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractFollowNFTURISetIterator{contract: _Contract.contract, event: "FollowNFTURISet", logs: logs, sub: sub}, nil
}

// WatchFollowNFTURISet is a free log subscription operation binding the contract event 0xe82886e1af6fcab5caef13815b22f51384e970c367a785f265d13860a7d6966d.
//
// Solidity: event FollowNFTURISet(uint256 indexed profileId, string followNFTURI, uint256 timestamp)
func (_Contract *ContractFilterer) WatchFollowNFTURISet(opts *bind.WatchOpts, sink chan<- *ContractFollowNFTURISet, profileId []*big.Int) (event.Subscription, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "FollowNFTURISet", profileIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractFollowNFTURISet)
				if err := _Contract.contract.UnpackLog(event, "FollowNFTURISet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFollowNFTURISet is a log parse operation binding the contract event 0xe82886e1af6fcab5caef13815b22f51384e970c367a785f265d13860a7d6966d.
//
// Solidity: event FollowNFTURISet(uint256 indexed profileId, string followNFTURI, uint256 timestamp)
func (_Contract *ContractFilterer) ParseFollowNFTURISet(log types.Log) (*ContractFollowNFTURISet, error) {
	event := new(ContractFollowNFTURISet)
	if err := _Contract.contract.UnpackLog(event, "FollowNFTURISet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractFollowedIterator is returned from FilterFollowed and is used to iterate over the raw logs and unpacked data for Followed events raised by the Contract contract.
type ContractFollowedIterator struct {
	Event *ContractFollowed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractFollowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractFollowed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractFollowed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractFollowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractFollowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractFollowed represents a Followed event raised by the Contract contract.
type ContractFollowed struct {
	Follower          common.Address
	ProfileIds        []*big.Int
	FollowModuleDatas [][]byte
	Timestamp         *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterFollowed is a free log retrieval operation binding the contract event 0x40487072dc56f384287d26fbe090f404143c2737d54632177451d1f74bd82c76.
//
// Solidity: event Followed(address indexed follower, uint256[] profileIds, bytes[] followModuleDatas, uint256 timestamp)
func (_Contract *ContractFilterer) FilterFollowed(opts *bind.FilterOpts, follower []common.Address) (*ContractFollowedIterator, error) {

	var followerRule []interface{}
	for _, followerItem := range follower {
		followerRule = append(followerRule, followerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Followed", followerRule)
	if err != nil {
		return nil, err
	}
	return &ContractFollowedIterator{contract: _Contract.contract, event: "Followed", logs: logs, sub: sub}, nil
}

// WatchFollowed is a free log subscription operation binding the contract event 0x40487072dc56f384287d26fbe090f404143c2737d54632177451d1f74bd82c76.
//
// Solidity: event Followed(address indexed follower, uint256[] profileIds, bytes[] followModuleDatas, uint256 timestamp)
func (_Contract *ContractFilterer) WatchFollowed(opts *bind.WatchOpts, sink chan<- *ContractFollowed, follower []common.Address) (event.Subscription, error) {

	var followerRule []interface{}
	for _, followerItem := range follower {
		followerRule = append(followerRule, followerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Followed", followerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractFollowed)
				if err := _Contract.contract.UnpackLog(event, "Followed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFollowed is a log parse operation binding the contract event 0x40487072dc56f384287d26fbe090f404143c2737d54632177451d1f74bd82c76.
//
// Solidity: event Followed(address indexed follower, uint256[] profileIds, bytes[] followModuleDatas, uint256 timestamp)
func (_Contract *ContractFilterer) ParseFollowed(log types.Log) (*ContractFollowed, error) {
	event := new(ContractFollowed)
	if err := _Contract.contract.UnpackLog(event, "Followed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractFollowsApprovedIterator is returned from FilterFollowsApproved and is used to iterate over the raw logs and unpacked data for FollowsApproved events raised by the Contract contract.
type ContractFollowsApprovedIterator struct {
	Event *ContractFollowsApproved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractFollowsApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractFollowsApproved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractFollowsApproved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractFollowsApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractFollowsApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractFollowsApproved represents a FollowsApproved event raised by the Contract contract.
type ContractFollowsApproved struct {
	Owner     common.Address
	ProfileId *big.Int
	Addresses []common.Address
	Approved  []bool
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFollowsApproved is a free log retrieval operation binding the contract event 0xc67fc3972da5d6434ab7b796ba133c240d40ee4e69129963c5aa0f2a6f7c3ad6.
//
// Solidity: event FollowsApproved(address indexed owner, uint256 indexed profileId, address[] addresses, bool[] approved, uint256 timestamp)
func (_Contract *ContractFilterer) FilterFollowsApproved(opts *bind.FilterOpts, owner []common.Address, profileId []*big.Int) (*ContractFollowsApprovedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "FollowsApproved", ownerRule, profileIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractFollowsApprovedIterator{contract: _Contract.contract, event: "FollowsApproved", logs: logs, sub: sub}, nil
}

// WatchFollowsApproved is a free log subscription operation binding the contract event 0xc67fc3972da5d6434ab7b796ba133c240d40ee4e69129963c5aa0f2a6f7c3ad6.
//
// Solidity: event FollowsApproved(address indexed owner, uint256 indexed profileId, address[] addresses, bool[] approved, uint256 timestamp)
func (_Contract *ContractFilterer) WatchFollowsApproved(opts *bind.WatchOpts, sink chan<- *ContractFollowsApproved, owner []common.Address, profileId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "FollowsApproved", ownerRule, profileIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractFollowsApproved)
				if err := _Contract.contract.UnpackLog(event, "FollowsApproved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFollowsApproved is a log parse operation binding the contract event 0xc67fc3972da5d6434ab7b796ba133c240d40ee4e69129963c5aa0f2a6f7c3ad6.
//
// Solidity: event FollowsApproved(address indexed owner, uint256 indexed profileId, address[] addresses, bool[] approved, uint256 timestamp)
func (_Contract *ContractFilterer) ParseFollowsApproved(log types.Log) (*ContractFollowsApproved, error) {
	event := new(ContractFollowsApproved)
	if err := _Contract.contract.UnpackLog(event, "FollowsApproved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractFollowsToggledIterator is returned from FilterFollowsToggled and is used to iterate over the raw logs and unpacked data for FollowsToggled events raised by the Contract contract.
type ContractFollowsToggledIterator struct {
	Event *ContractFollowsToggled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractFollowsToggledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractFollowsToggled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractFollowsToggled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractFollowsToggledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractFollowsToggledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractFollowsToggled represents a FollowsToggled event raised by the Contract contract.
type ContractFollowsToggled struct {
	Owner      common.Address
	ProfileIds []*big.Int
	Enabled    []bool
	Timestamp  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFollowsToggled is a free log retrieval operation binding the contract event 0x5538c80c8d3bee397d87a7d153f7f085bb12adf2fe25a026c7cc4e83d8c5f1d7.
//
// Solidity: event FollowsToggled(address indexed owner, uint256[] profileIds, bool[] enabled, uint256 timestamp)
func (_Contract *ContractFilterer) FilterFollowsToggled(opts *bind.FilterOpts, owner []common.Address) (*ContractFollowsToggledIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "FollowsToggled", ownerRule)
	if err != nil {
		return nil, err
	}
	return &ContractFollowsToggledIterator{contract: _Contract.contract, event: "FollowsToggled", logs: logs, sub: sub}, nil
}

// WatchFollowsToggled is a free log subscription operation binding the contract event 0x5538c80c8d3bee397d87a7d153f7f085bb12adf2fe25a026c7cc4e83d8c5f1d7.
//
// Solidity: event FollowsToggled(address indexed owner, uint256[] profileIds, bool[] enabled, uint256 timestamp)
func (_Contract *ContractFilterer) WatchFollowsToggled(opts *bind.WatchOpts, sink chan<- *ContractFollowsToggled, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "FollowsToggled", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractFollowsToggled)
				if err := _Contract.contract.UnpackLog(event, "FollowsToggled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFollowsToggled is a log parse operation binding the contract event 0x5538c80c8d3bee397d87a7d153f7f085bb12adf2fe25a026c7cc4e83d8c5f1d7.
//
// Solidity: event FollowsToggled(address indexed owner, uint256[] profileIds, bool[] enabled, uint256 timestamp)
func (_Contract *ContractFilterer) ParseFollowsToggled(log types.Log) (*ContractFollowsToggled, error) {
	event := new(ContractFollowsToggled)
	if err := _Contract.contract.UnpackLog(event, "FollowsToggled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractGovernanceSetIterator is returned from FilterGovernanceSet and is used to iterate over the raw logs and unpacked data for GovernanceSet events raised by the Contract contract.
type ContractGovernanceSetIterator struct {
	Event *ContractGovernanceSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractGovernanceSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractGovernanceSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractGovernanceSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractGovernanceSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractGovernanceSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractGovernanceSet represents a GovernanceSet event raised by the Contract contract.
type ContractGovernanceSet struct {
	Caller         common.Address
	PrevGovernance common.Address
	NewGovernance  common.Address
	Timestamp      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterGovernanceSet is a free log retrieval operation binding the contract event 0xe552a55455b740845a5c07ed445d1724142fc997b389835495a29b30cddc1ccd.
//
// Solidity: event GovernanceSet(address indexed caller, address indexed prevGovernance, address indexed newGovernance, uint256 timestamp)
func (_Contract *ContractFilterer) FilterGovernanceSet(opts *bind.FilterOpts, caller []common.Address, prevGovernance []common.Address, newGovernance []common.Address) (*ContractGovernanceSetIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var prevGovernanceRule []interface{}
	for _, prevGovernanceItem := range prevGovernance {
		prevGovernanceRule = append(prevGovernanceRule, prevGovernanceItem)
	}
	var newGovernanceRule []interface{}
	for _, newGovernanceItem := range newGovernance {
		newGovernanceRule = append(newGovernanceRule, newGovernanceItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "GovernanceSet", callerRule, prevGovernanceRule, newGovernanceRule)
	if err != nil {
		return nil, err
	}
	return &ContractGovernanceSetIterator{contract: _Contract.contract, event: "GovernanceSet", logs: logs, sub: sub}, nil
}

// WatchGovernanceSet is a free log subscription operation binding the contract event 0xe552a55455b740845a5c07ed445d1724142fc997b389835495a29b30cddc1ccd.
//
// Solidity: event GovernanceSet(address indexed caller, address indexed prevGovernance, address indexed newGovernance, uint256 timestamp)
func (_Contract *ContractFilterer) WatchGovernanceSet(opts *bind.WatchOpts, sink chan<- *ContractGovernanceSet, caller []common.Address, prevGovernance []common.Address, newGovernance []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var prevGovernanceRule []interface{}
	for _, prevGovernanceItem := range prevGovernance {
		prevGovernanceRule = append(prevGovernanceRule, prevGovernanceItem)
	}
	var newGovernanceRule []interface{}
	for _, newGovernanceItem := range newGovernance {
		newGovernanceRule = append(newGovernanceRule, newGovernanceItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "GovernanceSet", callerRule, prevGovernanceRule, newGovernanceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractGovernanceSet)
				if err := _Contract.contract.UnpackLog(event, "GovernanceSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseGovernanceSet is a log parse operation binding the contract event 0xe552a55455b740845a5c07ed445d1724142fc997b389835495a29b30cddc1ccd.
//
// Solidity: event GovernanceSet(address indexed caller, address indexed prevGovernance, address indexed newGovernance, uint256 timestamp)
func (_Contract *ContractFilterer) ParseGovernanceSet(log types.Log) (*ContractGovernanceSet, error) {
	event := new(ContractGovernanceSet)
	if err := _Contract.contract.UnpackLog(event, "GovernanceSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractMirrorCreatedIterator is returned from FilterMirrorCreated and is used to iterate over the raw logs and unpacked data for MirrorCreated events raised by the Contract contract.
type ContractMirrorCreatedIterator struct {
	Event *ContractMirrorCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractMirrorCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractMirrorCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractMirrorCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractMirrorCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractMirrorCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractMirrorCreated represents a MirrorCreated event raised by the Contract contract.
type ContractMirrorCreated struct {
	ProfileId                 *big.Int
	PubId                     *big.Int
	ProfileIdPointed          *big.Int
	PubIdPointed              *big.Int
	ReferenceModuleData       []byte
	ReferenceModule           common.Address
	ReferenceModuleReturnData []byte
	Timestamp                 *big.Int
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterMirrorCreated is a free log retrieval operation binding the contract event 0x9ea5dedb85bd9da4e264ee5a39b7ba0982e5d4d035d55edfa98a36b00e770b5a.
//
// Solidity: event MirrorCreated(uint256 indexed profileId, uint256 indexed pubId, uint256 profileIdPointed, uint256 pubIdPointed, bytes referenceModuleData, address referenceModule, bytes referenceModuleReturnData, uint256 timestamp)
func (_Contract *ContractFilterer) FilterMirrorCreated(opts *bind.FilterOpts, profileId []*big.Int, pubId []*big.Int) (*ContractMirrorCreatedIterator, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var pubIdRule []interface{}
	for _, pubIdItem := range pubId {
		pubIdRule = append(pubIdRule, pubIdItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "MirrorCreated", profileIdRule, pubIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractMirrorCreatedIterator{contract: _Contract.contract, event: "MirrorCreated", logs: logs, sub: sub}, nil
}

// WatchMirrorCreated is a free log subscription operation binding the contract event 0x9ea5dedb85bd9da4e264ee5a39b7ba0982e5d4d035d55edfa98a36b00e770b5a.
//
// Solidity: event MirrorCreated(uint256 indexed profileId, uint256 indexed pubId, uint256 profileIdPointed, uint256 pubIdPointed, bytes referenceModuleData, address referenceModule, bytes referenceModuleReturnData, uint256 timestamp)
func (_Contract *ContractFilterer) WatchMirrorCreated(opts *bind.WatchOpts, sink chan<- *ContractMirrorCreated, profileId []*big.Int, pubId []*big.Int) (event.Subscription, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var pubIdRule []interface{}
	for _, pubIdItem := range pubId {
		pubIdRule = append(pubIdRule, pubIdItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "MirrorCreated", profileIdRule, pubIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractMirrorCreated)
				if err := _Contract.contract.UnpackLog(event, "MirrorCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMirrorCreated is a log parse operation binding the contract event 0x9ea5dedb85bd9da4e264ee5a39b7ba0982e5d4d035d55edfa98a36b00e770b5a.
//
// Solidity: event MirrorCreated(uint256 indexed profileId, uint256 indexed pubId, uint256 profileIdPointed, uint256 pubIdPointed, bytes referenceModuleData, address referenceModule, bytes referenceModuleReturnData, uint256 timestamp)
func (_Contract *ContractFilterer) ParseMirrorCreated(log types.Log) (*ContractMirrorCreated, error) {
	event := new(ContractMirrorCreated)
	if err := _Contract.contract.UnpackLog(event, "MirrorCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractModuleBaseConstructedIterator is returned from FilterModuleBaseConstructed and is used to iterate over the raw logs and unpacked data for ModuleBaseConstructed events raised by the Contract contract.
type ContractModuleBaseConstructedIterator struct {
	Event *ContractModuleBaseConstructed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractModuleBaseConstructedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractModuleBaseConstructed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractModuleBaseConstructed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractModuleBaseConstructedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractModuleBaseConstructedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractModuleBaseConstructed represents a ModuleBaseConstructed event raised by the Contract contract.
type ContractModuleBaseConstructed struct {
	Hub       common.Address
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterModuleBaseConstructed is a free log retrieval operation binding the contract event 0xf1a1fa6b64aa95186f5a1285e76198d0da80d9c5a88062641d447f1d7c54e56c.
//
// Solidity: event ModuleBaseConstructed(address indexed hub, uint256 timestamp)
func (_Contract *ContractFilterer) FilterModuleBaseConstructed(opts *bind.FilterOpts, hub []common.Address) (*ContractModuleBaseConstructedIterator, error) {

	var hubRule []interface{}
	for _, hubItem := range hub {
		hubRule = append(hubRule, hubItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ModuleBaseConstructed", hubRule)
	if err != nil {
		return nil, err
	}
	return &ContractModuleBaseConstructedIterator{contract: _Contract.contract, event: "ModuleBaseConstructed", logs: logs, sub: sub}, nil
}

// WatchModuleBaseConstructed is a free log subscription operation binding the contract event 0xf1a1fa6b64aa95186f5a1285e76198d0da80d9c5a88062641d447f1d7c54e56c.
//
// Solidity: event ModuleBaseConstructed(address indexed hub, uint256 timestamp)
func (_Contract *ContractFilterer) WatchModuleBaseConstructed(opts *bind.WatchOpts, sink chan<- *ContractModuleBaseConstructed, hub []common.Address) (event.Subscription, error) {

	var hubRule []interface{}
	for _, hubItem := range hub {
		hubRule = append(hubRule, hubItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ModuleBaseConstructed", hubRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractModuleBaseConstructed)
				if err := _Contract.contract.UnpackLog(event, "ModuleBaseConstructed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseModuleBaseConstructed is a log parse operation binding the contract event 0xf1a1fa6b64aa95186f5a1285e76198d0da80d9c5a88062641d447f1d7c54e56c.
//
// Solidity: event ModuleBaseConstructed(address indexed hub, uint256 timestamp)
func (_Contract *ContractFilterer) ParseModuleBaseConstructed(log types.Log) (*ContractModuleBaseConstructed, error) {
	event := new(ContractModuleBaseConstructed)
	if err := _Contract.contract.UnpackLog(event, "ModuleBaseConstructed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractModuleGlobalsCurrencyWhitelistedIterator is returned from FilterModuleGlobalsCurrencyWhitelisted and is used to iterate over the raw logs and unpacked data for ModuleGlobalsCurrencyWhitelisted events raised by the Contract contract.
type ContractModuleGlobalsCurrencyWhitelistedIterator struct {
	Event *ContractModuleGlobalsCurrencyWhitelisted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractModuleGlobalsCurrencyWhitelistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractModuleGlobalsCurrencyWhitelisted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractModuleGlobalsCurrencyWhitelisted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractModuleGlobalsCurrencyWhitelistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractModuleGlobalsCurrencyWhitelistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractModuleGlobalsCurrencyWhitelisted represents a ModuleGlobalsCurrencyWhitelisted event raised by the Contract contract.
type ContractModuleGlobalsCurrencyWhitelisted struct {
	Currency        common.Address
	PrevWhitelisted bool
	Whitelisted     bool
	Timestamp       *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterModuleGlobalsCurrencyWhitelisted is a free log retrieval operation binding the contract event 0x79c3cefc851fd6040f06af202c542818d9fb39bcddcb7a7e3f563b15300a2743.
//
// Solidity: event ModuleGlobalsCurrencyWhitelisted(address indexed currency, bool indexed prevWhitelisted, bool indexed whitelisted, uint256 timestamp)
func (_Contract *ContractFilterer) FilterModuleGlobalsCurrencyWhitelisted(opts *bind.FilterOpts, currency []common.Address, prevWhitelisted []bool, whitelisted []bool) (*ContractModuleGlobalsCurrencyWhitelistedIterator, error) {

	var currencyRule []interface{}
	for _, currencyItem := range currency {
		currencyRule = append(currencyRule, currencyItem)
	}
	var prevWhitelistedRule []interface{}
	for _, prevWhitelistedItem := range prevWhitelisted {
		prevWhitelistedRule = append(prevWhitelistedRule, prevWhitelistedItem)
	}
	var whitelistedRule []interface{}
	for _, whitelistedItem := range whitelisted {
		whitelistedRule = append(whitelistedRule, whitelistedItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ModuleGlobalsCurrencyWhitelisted", currencyRule, prevWhitelistedRule, whitelistedRule)
	if err != nil {
		return nil, err
	}
	return &ContractModuleGlobalsCurrencyWhitelistedIterator{contract: _Contract.contract, event: "ModuleGlobalsCurrencyWhitelisted", logs: logs, sub: sub}, nil
}

// WatchModuleGlobalsCurrencyWhitelisted is a free log subscription operation binding the contract event 0x79c3cefc851fd6040f06af202c542818d9fb39bcddcb7a7e3f563b15300a2743.
//
// Solidity: event ModuleGlobalsCurrencyWhitelisted(address indexed currency, bool indexed prevWhitelisted, bool indexed whitelisted, uint256 timestamp)
func (_Contract *ContractFilterer) WatchModuleGlobalsCurrencyWhitelisted(opts *bind.WatchOpts, sink chan<- *ContractModuleGlobalsCurrencyWhitelisted, currency []common.Address, prevWhitelisted []bool, whitelisted []bool) (event.Subscription, error) {

	var currencyRule []interface{}
	for _, currencyItem := range currency {
		currencyRule = append(currencyRule, currencyItem)
	}
	var prevWhitelistedRule []interface{}
	for _, prevWhitelistedItem := range prevWhitelisted {
		prevWhitelistedRule = append(prevWhitelistedRule, prevWhitelistedItem)
	}
	var whitelistedRule []interface{}
	for _, whitelistedItem := range whitelisted {
		whitelistedRule = append(whitelistedRule, whitelistedItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ModuleGlobalsCurrencyWhitelisted", currencyRule, prevWhitelistedRule, whitelistedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractModuleGlobalsCurrencyWhitelisted)
				if err := _Contract.contract.UnpackLog(event, "ModuleGlobalsCurrencyWhitelisted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseModuleGlobalsCurrencyWhitelisted is a log parse operation binding the contract event 0x79c3cefc851fd6040f06af202c542818d9fb39bcddcb7a7e3f563b15300a2743.
//
// Solidity: event ModuleGlobalsCurrencyWhitelisted(address indexed currency, bool indexed prevWhitelisted, bool indexed whitelisted, uint256 timestamp)
func (_Contract *ContractFilterer) ParseModuleGlobalsCurrencyWhitelisted(log types.Log) (*ContractModuleGlobalsCurrencyWhitelisted, error) {
	event := new(ContractModuleGlobalsCurrencyWhitelisted)
	if err := _Contract.contract.UnpackLog(event, "ModuleGlobalsCurrencyWhitelisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractModuleGlobalsGovernanceSetIterator is returned from FilterModuleGlobalsGovernanceSet and is used to iterate over the raw logs and unpacked data for ModuleGlobalsGovernanceSet events raised by the Contract contract.
type ContractModuleGlobalsGovernanceSetIterator struct {
	Event *ContractModuleGlobalsGovernanceSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractModuleGlobalsGovernanceSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractModuleGlobalsGovernanceSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractModuleGlobalsGovernanceSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractModuleGlobalsGovernanceSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractModuleGlobalsGovernanceSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractModuleGlobalsGovernanceSet represents a ModuleGlobalsGovernanceSet event raised by the Contract contract.
type ContractModuleGlobalsGovernanceSet struct {
	PrevGovernance common.Address
	NewGovernance  common.Address
	Timestamp      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterModuleGlobalsGovernanceSet is a free log retrieval operation binding the contract event 0xbf538a2c0db3d440906b8179dd0394f68a65b0b1481da70ffee24e19dccee84c.
//
// Solidity: event ModuleGlobalsGovernanceSet(address indexed prevGovernance, address indexed newGovernance, uint256 timestamp)
func (_Contract *ContractFilterer) FilterModuleGlobalsGovernanceSet(opts *bind.FilterOpts, prevGovernance []common.Address, newGovernance []common.Address) (*ContractModuleGlobalsGovernanceSetIterator, error) {

	var prevGovernanceRule []interface{}
	for _, prevGovernanceItem := range prevGovernance {
		prevGovernanceRule = append(prevGovernanceRule, prevGovernanceItem)
	}
	var newGovernanceRule []interface{}
	for _, newGovernanceItem := range newGovernance {
		newGovernanceRule = append(newGovernanceRule, newGovernanceItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ModuleGlobalsGovernanceSet", prevGovernanceRule, newGovernanceRule)
	if err != nil {
		return nil, err
	}
	return &ContractModuleGlobalsGovernanceSetIterator{contract: _Contract.contract, event: "ModuleGlobalsGovernanceSet", logs: logs, sub: sub}, nil
}

// WatchModuleGlobalsGovernanceSet is a free log subscription operation binding the contract event 0xbf538a2c0db3d440906b8179dd0394f68a65b0b1481da70ffee24e19dccee84c.
//
// Solidity: event ModuleGlobalsGovernanceSet(address indexed prevGovernance, address indexed newGovernance, uint256 timestamp)
func (_Contract *ContractFilterer) WatchModuleGlobalsGovernanceSet(opts *bind.WatchOpts, sink chan<- *ContractModuleGlobalsGovernanceSet, prevGovernance []common.Address, newGovernance []common.Address) (event.Subscription, error) {

	var prevGovernanceRule []interface{}
	for _, prevGovernanceItem := range prevGovernance {
		prevGovernanceRule = append(prevGovernanceRule, prevGovernanceItem)
	}
	var newGovernanceRule []interface{}
	for _, newGovernanceItem := range newGovernance {
		newGovernanceRule = append(newGovernanceRule, newGovernanceItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ModuleGlobalsGovernanceSet", prevGovernanceRule, newGovernanceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractModuleGlobalsGovernanceSet)
				if err := _Contract.contract.UnpackLog(event, "ModuleGlobalsGovernanceSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseModuleGlobalsGovernanceSet is a log parse operation binding the contract event 0xbf538a2c0db3d440906b8179dd0394f68a65b0b1481da70ffee24e19dccee84c.
//
// Solidity: event ModuleGlobalsGovernanceSet(address indexed prevGovernance, address indexed newGovernance, uint256 timestamp)
func (_Contract *ContractFilterer) ParseModuleGlobalsGovernanceSet(log types.Log) (*ContractModuleGlobalsGovernanceSet, error) {
	event := new(ContractModuleGlobalsGovernanceSet)
	if err := _Contract.contract.UnpackLog(event, "ModuleGlobalsGovernanceSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractModuleGlobalsTreasuryFeeSetIterator is returned from FilterModuleGlobalsTreasuryFeeSet and is used to iterate over the raw logs and unpacked data for ModuleGlobalsTreasuryFeeSet events raised by the Contract contract.
type ContractModuleGlobalsTreasuryFeeSetIterator struct {
	Event *ContractModuleGlobalsTreasuryFeeSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractModuleGlobalsTreasuryFeeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractModuleGlobalsTreasuryFeeSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractModuleGlobalsTreasuryFeeSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractModuleGlobalsTreasuryFeeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractModuleGlobalsTreasuryFeeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractModuleGlobalsTreasuryFeeSet represents a ModuleGlobalsTreasuryFeeSet event raised by the Contract contract.
type ContractModuleGlobalsTreasuryFeeSet struct {
	PrevTreasuryFee uint16
	NewTreasuryFee  uint16
	Timestamp       *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterModuleGlobalsTreasuryFeeSet is a free log retrieval operation binding the contract event 0xec936862e6bb897cd711a5f31825057583128c2a482f0a4c9a4e6c3fd7c023f4.
//
// Solidity: event ModuleGlobalsTreasuryFeeSet(uint16 indexed prevTreasuryFee, uint16 indexed newTreasuryFee, uint256 timestamp)
func (_Contract *ContractFilterer) FilterModuleGlobalsTreasuryFeeSet(opts *bind.FilterOpts, prevTreasuryFee []uint16, newTreasuryFee []uint16) (*ContractModuleGlobalsTreasuryFeeSetIterator, error) {

	var prevTreasuryFeeRule []interface{}
	for _, prevTreasuryFeeItem := range prevTreasuryFee {
		prevTreasuryFeeRule = append(prevTreasuryFeeRule, prevTreasuryFeeItem)
	}
	var newTreasuryFeeRule []interface{}
	for _, newTreasuryFeeItem := range newTreasuryFee {
		newTreasuryFeeRule = append(newTreasuryFeeRule, newTreasuryFeeItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ModuleGlobalsTreasuryFeeSet", prevTreasuryFeeRule, newTreasuryFeeRule)
	if err != nil {
		return nil, err
	}
	return &ContractModuleGlobalsTreasuryFeeSetIterator{contract: _Contract.contract, event: "ModuleGlobalsTreasuryFeeSet", logs: logs, sub: sub}, nil
}

// WatchModuleGlobalsTreasuryFeeSet is a free log subscription operation binding the contract event 0xec936862e6bb897cd711a5f31825057583128c2a482f0a4c9a4e6c3fd7c023f4.
//
// Solidity: event ModuleGlobalsTreasuryFeeSet(uint16 indexed prevTreasuryFee, uint16 indexed newTreasuryFee, uint256 timestamp)
func (_Contract *ContractFilterer) WatchModuleGlobalsTreasuryFeeSet(opts *bind.WatchOpts, sink chan<- *ContractModuleGlobalsTreasuryFeeSet, prevTreasuryFee []uint16, newTreasuryFee []uint16) (event.Subscription, error) {

	var prevTreasuryFeeRule []interface{}
	for _, prevTreasuryFeeItem := range prevTreasuryFee {
		prevTreasuryFeeRule = append(prevTreasuryFeeRule, prevTreasuryFeeItem)
	}
	var newTreasuryFeeRule []interface{}
	for _, newTreasuryFeeItem := range newTreasuryFee {
		newTreasuryFeeRule = append(newTreasuryFeeRule, newTreasuryFeeItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ModuleGlobalsTreasuryFeeSet", prevTreasuryFeeRule, newTreasuryFeeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractModuleGlobalsTreasuryFeeSet)
				if err := _Contract.contract.UnpackLog(event, "ModuleGlobalsTreasuryFeeSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseModuleGlobalsTreasuryFeeSet is a log parse operation binding the contract event 0xec936862e6bb897cd711a5f31825057583128c2a482f0a4c9a4e6c3fd7c023f4.
//
// Solidity: event ModuleGlobalsTreasuryFeeSet(uint16 indexed prevTreasuryFee, uint16 indexed newTreasuryFee, uint256 timestamp)
func (_Contract *ContractFilterer) ParseModuleGlobalsTreasuryFeeSet(log types.Log) (*ContractModuleGlobalsTreasuryFeeSet, error) {
	event := new(ContractModuleGlobalsTreasuryFeeSet)
	if err := _Contract.contract.UnpackLog(event, "ModuleGlobalsTreasuryFeeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractModuleGlobalsTreasurySetIterator is returned from FilterModuleGlobalsTreasurySet and is used to iterate over the raw logs and unpacked data for ModuleGlobalsTreasurySet events raised by the Contract contract.
type ContractModuleGlobalsTreasurySetIterator struct {
	Event *ContractModuleGlobalsTreasurySet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractModuleGlobalsTreasurySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractModuleGlobalsTreasurySet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractModuleGlobalsTreasurySet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractModuleGlobalsTreasurySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractModuleGlobalsTreasurySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractModuleGlobalsTreasurySet represents a ModuleGlobalsTreasurySet event raised by the Contract contract.
type ContractModuleGlobalsTreasurySet struct {
	PrevTreasury common.Address
	NewTreasury  common.Address
	Timestamp    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterModuleGlobalsTreasurySet is a free log retrieval operation binding the contract event 0x3dfc53d6b49bfbc932b215ba515f0d0ab0e17aac17726fba48075f0c16c7ffe3.
//
// Solidity: event ModuleGlobalsTreasurySet(address indexed prevTreasury, address indexed newTreasury, uint256 timestamp)
func (_Contract *ContractFilterer) FilterModuleGlobalsTreasurySet(opts *bind.FilterOpts, prevTreasury []common.Address, newTreasury []common.Address) (*ContractModuleGlobalsTreasurySetIterator, error) {

	var prevTreasuryRule []interface{}
	for _, prevTreasuryItem := range prevTreasury {
		prevTreasuryRule = append(prevTreasuryRule, prevTreasuryItem)
	}
	var newTreasuryRule []interface{}
	for _, newTreasuryItem := range newTreasury {
		newTreasuryRule = append(newTreasuryRule, newTreasuryItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ModuleGlobalsTreasurySet", prevTreasuryRule, newTreasuryRule)
	if err != nil {
		return nil, err
	}
	return &ContractModuleGlobalsTreasurySetIterator{contract: _Contract.contract, event: "ModuleGlobalsTreasurySet", logs: logs, sub: sub}, nil
}

// WatchModuleGlobalsTreasurySet is a free log subscription operation binding the contract event 0x3dfc53d6b49bfbc932b215ba515f0d0ab0e17aac17726fba48075f0c16c7ffe3.
//
// Solidity: event ModuleGlobalsTreasurySet(address indexed prevTreasury, address indexed newTreasury, uint256 timestamp)
func (_Contract *ContractFilterer) WatchModuleGlobalsTreasurySet(opts *bind.WatchOpts, sink chan<- *ContractModuleGlobalsTreasurySet, prevTreasury []common.Address, newTreasury []common.Address) (event.Subscription, error) {

	var prevTreasuryRule []interface{}
	for _, prevTreasuryItem := range prevTreasury {
		prevTreasuryRule = append(prevTreasuryRule, prevTreasuryItem)
	}
	var newTreasuryRule []interface{}
	for _, newTreasuryItem := range newTreasury {
		newTreasuryRule = append(newTreasuryRule, newTreasuryItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ModuleGlobalsTreasurySet", prevTreasuryRule, newTreasuryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractModuleGlobalsTreasurySet)
				if err := _Contract.contract.UnpackLog(event, "ModuleGlobalsTreasurySet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseModuleGlobalsTreasurySet is a log parse operation binding the contract event 0x3dfc53d6b49bfbc932b215ba515f0d0ab0e17aac17726fba48075f0c16c7ffe3.
//
// Solidity: event ModuleGlobalsTreasurySet(address indexed prevTreasury, address indexed newTreasury, uint256 timestamp)
func (_Contract *ContractFilterer) ParseModuleGlobalsTreasurySet(log types.Log) (*ContractModuleGlobalsTreasurySet, error) {
	event := new(ContractModuleGlobalsTreasurySet)
	if err := _Contract.contract.UnpackLog(event, "ModuleGlobalsTreasurySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractPostCreatedIterator is returned from FilterPostCreated and is used to iterate over the raw logs and unpacked data for PostCreated events raised by the Contract contract.
type ContractPostCreatedIterator struct {
	Event *ContractPostCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractPostCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractPostCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractPostCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractPostCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractPostCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractPostCreated represents a PostCreated event raised by the Contract contract.
type ContractPostCreated struct {
	ProfileId                 *big.Int
	PubId                     *big.Int
	ContentURI                string
	CollectModule             common.Address
	CollectModuleReturnData   []byte
	ReferenceModule           common.Address
	ReferenceModuleReturnData []byte
	Timestamp                 *big.Int
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterPostCreated is a free log retrieval operation binding the contract event 0xc672c38b4d26c3c978228e99164105280410b144af24dd3ed8e4f9d211d96a50.
//
// Solidity: event PostCreated(uint256 indexed profileId, uint256 indexed pubId, string contentURI, address collectModule, bytes collectModuleReturnData, address referenceModule, bytes referenceModuleReturnData, uint256 timestamp)
func (_Contract *ContractFilterer) FilterPostCreated(opts *bind.FilterOpts, profileId []*big.Int, pubId []*big.Int) (*ContractPostCreatedIterator, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var pubIdRule []interface{}
	for _, pubIdItem := range pubId {
		pubIdRule = append(pubIdRule, pubIdItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "PostCreated", profileIdRule, pubIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractPostCreatedIterator{contract: _Contract.contract, event: "PostCreated", logs: logs, sub: sub}, nil
}

// WatchPostCreated is a free log subscription operation binding the contract event 0xc672c38b4d26c3c978228e99164105280410b144af24dd3ed8e4f9d211d96a50.
//
// Solidity: event PostCreated(uint256 indexed profileId, uint256 indexed pubId, string contentURI, address collectModule, bytes collectModuleReturnData, address referenceModule, bytes referenceModuleReturnData, uint256 timestamp)
func (_Contract *ContractFilterer) WatchPostCreated(opts *bind.WatchOpts, sink chan<- *ContractPostCreated, profileId []*big.Int, pubId []*big.Int) (event.Subscription, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var pubIdRule []interface{}
	for _, pubIdItem := range pubId {
		pubIdRule = append(pubIdRule, pubIdItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "PostCreated", profileIdRule, pubIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractPostCreated)
				if err := _Contract.contract.UnpackLog(event, "PostCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePostCreated is a log parse operation binding the contract event 0xc672c38b4d26c3c978228e99164105280410b144af24dd3ed8e4f9d211d96a50.
//
// Solidity: event PostCreated(uint256 indexed profileId, uint256 indexed pubId, string contentURI, address collectModule, bytes collectModuleReturnData, address referenceModule, bytes referenceModuleReturnData, uint256 timestamp)
func (_Contract *ContractFilterer) ParsePostCreated(log types.Log) (*ContractPostCreated, error) {
	event := new(ContractPostCreated)
	if err := _Contract.contract.UnpackLog(event, "PostCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractProfileCreatedIterator is returned from FilterProfileCreated and is used to iterate over the raw logs and unpacked data for ProfileCreated events raised by the Contract contract.
type ContractProfileCreatedIterator struct {
	Event *ContractProfileCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractProfileCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractProfileCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractProfileCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractProfileCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractProfileCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractProfileCreated represents a ProfileCreated event raised by the Contract contract.
type ContractProfileCreated struct {
	ProfileId              *big.Int
	Creator                common.Address
	To                     common.Address
	Handle                 string
	ImageURI               string
	FollowModule           common.Address
	FollowModuleReturnData []byte
	FollowNFTURI           string
	Timestamp              *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterProfileCreated is a free log retrieval operation binding the contract event 0x4e14f57cff7910416f2ef43cf05019b5a97a313de71fec9344be11b9b88fed12.
//
// Solidity: event ProfileCreated(uint256 indexed profileId, address indexed creator, address indexed to, string handle, string imageURI, address followModule, bytes followModuleReturnData, string followNFTURI, uint256 timestamp)
func (_Contract *ContractFilterer) FilterProfileCreated(opts *bind.FilterOpts, profileId []*big.Int, creator []common.Address, to []common.Address) (*ContractProfileCreatedIterator, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ProfileCreated", profileIdRule, creatorRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ContractProfileCreatedIterator{contract: _Contract.contract, event: "ProfileCreated", logs: logs, sub: sub}, nil
}

// WatchProfileCreated is a free log subscription operation binding the contract event 0x4e14f57cff7910416f2ef43cf05019b5a97a313de71fec9344be11b9b88fed12.
//
// Solidity: event ProfileCreated(uint256 indexed profileId, address indexed creator, address indexed to, string handle, string imageURI, address followModule, bytes followModuleReturnData, string followNFTURI, uint256 timestamp)
func (_Contract *ContractFilterer) WatchProfileCreated(opts *bind.WatchOpts, sink chan<- *ContractProfileCreated, profileId []*big.Int, creator []common.Address, to []common.Address) (event.Subscription, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ProfileCreated", profileIdRule, creatorRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractProfileCreated)
				if err := _Contract.contract.UnpackLog(event, "ProfileCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProfileCreated is a log parse operation binding the contract event 0x4e14f57cff7910416f2ef43cf05019b5a97a313de71fec9344be11b9b88fed12.
//
// Solidity: event ProfileCreated(uint256 indexed profileId, address indexed creator, address indexed to, string handle, string imageURI, address followModule, bytes followModuleReturnData, string followNFTURI, uint256 timestamp)
func (_Contract *ContractFilterer) ParseProfileCreated(log types.Log) (*ContractProfileCreated, error) {
	event := new(ContractProfileCreated)
	if err := _Contract.contract.UnpackLog(event, "ProfileCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractProfileCreatorWhitelistedIterator is returned from FilterProfileCreatorWhitelisted and is used to iterate over the raw logs and unpacked data for ProfileCreatorWhitelisted events raised by the Contract contract.
type ContractProfileCreatorWhitelistedIterator struct {
	Event *ContractProfileCreatorWhitelisted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractProfileCreatorWhitelistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractProfileCreatorWhitelisted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractProfileCreatorWhitelisted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractProfileCreatorWhitelistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractProfileCreatorWhitelistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractProfileCreatorWhitelisted represents a ProfileCreatorWhitelisted event raised by the Contract contract.
type ContractProfileCreatorWhitelisted struct {
	ProfileCreator common.Address
	Whitelisted    bool
	Timestamp      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterProfileCreatorWhitelisted is a free log retrieval operation binding the contract event 0x8f617843889b94892bd44852d36ca6a7f49ecf4350a01e7b68e22d80f4ed95bc.
//
// Solidity: event ProfileCreatorWhitelisted(address indexed profileCreator, bool indexed whitelisted, uint256 timestamp)
func (_Contract *ContractFilterer) FilterProfileCreatorWhitelisted(opts *bind.FilterOpts, profileCreator []common.Address, whitelisted []bool) (*ContractProfileCreatorWhitelistedIterator, error) {

	var profileCreatorRule []interface{}
	for _, profileCreatorItem := range profileCreator {
		profileCreatorRule = append(profileCreatorRule, profileCreatorItem)
	}
	var whitelistedRule []interface{}
	for _, whitelistedItem := range whitelisted {
		whitelistedRule = append(whitelistedRule, whitelistedItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ProfileCreatorWhitelisted", profileCreatorRule, whitelistedRule)
	if err != nil {
		return nil, err
	}
	return &ContractProfileCreatorWhitelistedIterator{contract: _Contract.contract, event: "ProfileCreatorWhitelisted", logs: logs, sub: sub}, nil
}

// WatchProfileCreatorWhitelisted is a free log subscription operation binding the contract event 0x8f617843889b94892bd44852d36ca6a7f49ecf4350a01e7b68e22d80f4ed95bc.
//
// Solidity: event ProfileCreatorWhitelisted(address indexed profileCreator, bool indexed whitelisted, uint256 timestamp)
func (_Contract *ContractFilterer) WatchProfileCreatorWhitelisted(opts *bind.WatchOpts, sink chan<- *ContractProfileCreatorWhitelisted, profileCreator []common.Address, whitelisted []bool) (event.Subscription, error) {

	var profileCreatorRule []interface{}
	for _, profileCreatorItem := range profileCreator {
		profileCreatorRule = append(profileCreatorRule, profileCreatorItem)
	}
	var whitelistedRule []interface{}
	for _, whitelistedItem := range whitelisted {
		whitelistedRule = append(whitelistedRule, whitelistedItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ProfileCreatorWhitelisted", profileCreatorRule, whitelistedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractProfileCreatorWhitelisted)
				if err := _Contract.contract.UnpackLog(event, "ProfileCreatorWhitelisted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProfileCreatorWhitelisted is a log parse operation binding the contract event 0x8f617843889b94892bd44852d36ca6a7f49ecf4350a01e7b68e22d80f4ed95bc.
//
// Solidity: event ProfileCreatorWhitelisted(address indexed profileCreator, bool indexed whitelisted, uint256 timestamp)
func (_Contract *ContractFilterer) ParseProfileCreatorWhitelisted(log types.Log) (*ContractProfileCreatorWhitelisted, error) {
	event := new(ContractProfileCreatorWhitelisted)
	if err := _Contract.contract.UnpackLog(event, "ProfileCreatorWhitelisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractProfileImageURISetIterator is returned from FilterProfileImageURISet and is used to iterate over the raw logs and unpacked data for ProfileImageURISet events raised by the Contract contract.
type ContractProfileImageURISetIterator struct {
	Event *ContractProfileImageURISet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractProfileImageURISetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractProfileImageURISet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractProfileImageURISet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractProfileImageURISetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractProfileImageURISetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractProfileImageURISet represents a ProfileImageURISet event raised by the Contract contract.
type ContractProfileImageURISet struct {
	ProfileId *big.Int
	ImageURI  string
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterProfileImageURISet is a free log retrieval operation binding the contract event 0xd5a5879cad33c830cc1432c1850107029a09c80c60e9bce1ecd08d24880bd46c.
//
// Solidity: event ProfileImageURISet(uint256 indexed profileId, string imageURI, uint256 timestamp)
func (_Contract *ContractFilterer) FilterProfileImageURISet(opts *bind.FilterOpts, profileId []*big.Int) (*ContractProfileImageURISetIterator, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ProfileImageURISet", profileIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractProfileImageURISetIterator{contract: _Contract.contract, event: "ProfileImageURISet", logs: logs, sub: sub}, nil
}

// WatchProfileImageURISet is a free log subscription operation binding the contract event 0xd5a5879cad33c830cc1432c1850107029a09c80c60e9bce1ecd08d24880bd46c.
//
// Solidity: event ProfileImageURISet(uint256 indexed profileId, string imageURI, uint256 timestamp)
func (_Contract *ContractFilterer) WatchProfileImageURISet(opts *bind.WatchOpts, sink chan<- *ContractProfileImageURISet, profileId []*big.Int) (event.Subscription, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ProfileImageURISet", profileIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractProfileImageURISet)
				if err := _Contract.contract.UnpackLog(event, "ProfileImageURISet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProfileImageURISet is a log parse operation binding the contract event 0xd5a5879cad33c830cc1432c1850107029a09c80c60e9bce1ecd08d24880bd46c.
//
// Solidity: event ProfileImageURISet(uint256 indexed profileId, string imageURI, uint256 timestamp)
func (_Contract *ContractFilterer) ParseProfileImageURISet(log types.Log) (*ContractProfileImageURISet, error) {
	event := new(ContractProfileImageURISet)
	if err := _Contract.contract.UnpackLog(event, "ProfileImageURISet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractProfileMetadataSetIterator is returned from FilterProfileMetadataSet and is used to iterate over the raw logs and unpacked data for ProfileMetadataSet events raised by the Contract contract.
type ContractProfileMetadataSetIterator struct {
	Event *ContractProfileMetadataSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractProfileMetadataSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractProfileMetadataSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractProfileMetadataSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractProfileMetadataSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractProfileMetadataSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractProfileMetadataSet represents a ProfileMetadataSet event raised by the Contract contract.
type ContractProfileMetadataSet struct {
	ProfileId *big.Int
	Metadata  string
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterProfileMetadataSet is a free log retrieval operation binding the contract event 0xf901a8b3832914a45999dd4c425fbe42eb4182724d394100401e633d9f0b286a.
//
// Solidity: event ProfileMetadataSet(uint256 indexed profileId, string metadata, uint256 timestamp)
func (_Contract *ContractFilterer) FilterProfileMetadataSet(opts *bind.FilterOpts, profileId []*big.Int) (*ContractProfileMetadataSetIterator, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ProfileMetadataSet", profileIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractProfileMetadataSetIterator{contract: _Contract.contract, event: "ProfileMetadataSet", logs: logs, sub: sub}, nil
}

// WatchProfileMetadataSet is a free log subscription operation binding the contract event 0xf901a8b3832914a45999dd4c425fbe42eb4182724d394100401e633d9f0b286a.
//
// Solidity: event ProfileMetadataSet(uint256 indexed profileId, string metadata, uint256 timestamp)
func (_Contract *ContractFilterer) WatchProfileMetadataSet(opts *bind.WatchOpts, sink chan<- *ContractProfileMetadataSet, profileId []*big.Int) (event.Subscription, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ProfileMetadataSet", profileIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractProfileMetadataSet)
				if err := _Contract.contract.UnpackLog(event, "ProfileMetadataSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProfileMetadataSet is a log parse operation binding the contract event 0xf901a8b3832914a45999dd4c425fbe42eb4182724d394100401e633d9f0b286a.
//
// Solidity: event ProfileMetadataSet(uint256 indexed profileId, string metadata, uint256 timestamp)
func (_Contract *ContractFilterer) ParseProfileMetadataSet(log types.Log) (*ContractProfileMetadataSet, error) {
	event := new(ContractProfileMetadataSet)
	if err := _Contract.contract.UnpackLog(event, "ProfileMetadataSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractReferenceModuleWhitelistedIterator is returned from FilterReferenceModuleWhitelisted and is used to iterate over the raw logs and unpacked data for ReferenceModuleWhitelisted events raised by the Contract contract.
type ContractReferenceModuleWhitelistedIterator struct {
	Event *ContractReferenceModuleWhitelisted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractReferenceModuleWhitelistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractReferenceModuleWhitelisted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractReferenceModuleWhitelisted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractReferenceModuleWhitelistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractReferenceModuleWhitelistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractReferenceModuleWhitelisted represents a ReferenceModuleWhitelisted event raised by the Contract contract.
type ContractReferenceModuleWhitelisted struct {
	ReferenceModule common.Address
	Whitelisted     bool
	Timestamp       *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterReferenceModuleWhitelisted is a free log retrieval operation binding the contract event 0x37872a053ef20cb52defb7c9ec20e1a87cb8dd5098ac9e76a144be263dfef572.
//
// Solidity: event ReferenceModuleWhitelisted(address indexed referenceModule, bool indexed whitelisted, uint256 timestamp)
func (_Contract *ContractFilterer) FilterReferenceModuleWhitelisted(opts *bind.FilterOpts, referenceModule []common.Address, whitelisted []bool) (*ContractReferenceModuleWhitelistedIterator, error) {

	var referenceModuleRule []interface{}
	for _, referenceModuleItem := range referenceModule {
		referenceModuleRule = append(referenceModuleRule, referenceModuleItem)
	}
	var whitelistedRule []interface{}
	for _, whitelistedItem := range whitelisted {
		whitelistedRule = append(whitelistedRule, whitelistedItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ReferenceModuleWhitelisted", referenceModuleRule, whitelistedRule)
	if err != nil {
		return nil, err
	}
	return &ContractReferenceModuleWhitelistedIterator{contract: _Contract.contract, event: "ReferenceModuleWhitelisted", logs: logs, sub: sub}, nil
}

// WatchReferenceModuleWhitelisted is a free log subscription operation binding the contract event 0x37872a053ef20cb52defb7c9ec20e1a87cb8dd5098ac9e76a144be263dfef572.
//
// Solidity: event ReferenceModuleWhitelisted(address indexed referenceModule, bool indexed whitelisted, uint256 timestamp)
func (_Contract *ContractFilterer) WatchReferenceModuleWhitelisted(opts *bind.WatchOpts, sink chan<- *ContractReferenceModuleWhitelisted, referenceModule []common.Address, whitelisted []bool) (event.Subscription, error) {

	var referenceModuleRule []interface{}
	for _, referenceModuleItem := range referenceModule {
		referenceModuleRule = append(referenceModuleRule, referenceModuleItem)
	}
	var whitelistedRule []interface{}
	for _, whitelistedItem := range whitelisted {
		whitelistedRule = append(whitelistedRule, whitelistedItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ReferenceModuleWhitelisted", referenceModuleRule, whitelistedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractReferenceModuleWhitelisted)
				if err := _Contract.contract.UnpackLog(event, "ReferenceModuleWhitelisted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseReferenceModuleWhitelisted is a log parse operation binding the contract event 0x37872a053ef20cb52defb7c9ec20e1a87cb8dd5098ac9e76a144be263dfef572.
//
// Solidity: event ReferenceModuleWhitelisted(address indexed referenceModule, bool indexed whitelisted, uint256 timestamp)
func (_Contract *ContractFilterer) ParseReferenceModuleWhitelisted(log types.Log) (*ContractReferenceModuleWhitelisted, error) {
	event := new(ContractReferenceModuleWhitelisted)
	if err := _Contract.contract.UnpackLog(event, "ReferenceModuleWhitelisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractStateSetIterator is returned from FilterStateSet and is used to iterate over the raw logs and unpacked data for StateSet events raised by the Contract contract.
type ContractStateSetIterator struct {
	Event *ContractStateSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractStateSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractStateSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractStateSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractStateSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractStateSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractStateSet represents a StateSet event raised by the Contract contract.
type ContractStateSet struct {
	Caller    common.Address
	PrevState uint8
	NewState  uint8
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStateSet is a free log retrieval operation binding the contract event 0xa2f9a1499fc1f9b7796d21fe5761290ccb7e0ef6ccf35fa58b668f304a62a1ca.
//
// Solidity: event StateSet(address indexed caller, uint8 indexed prevState, uint8 indexed newState, uint256 timestamp)
func (_Contract *ContractFilterer) FilterStateSet(opts *bind.FilterOpts, caller []common.Address, prevState []uint8, newState []uint8) (*ContractStateSetIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var prevStateRule []interface{}
	for _, prevStateItem := range prevState {
		prevStateRule = append(prevStateRule, prevStateItem)
	}
	var newStateRule []interface{}
	for _, newStateItem := range newState {
		newStateRule = append(newStateRule, newStateItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "StateSet", callerRule, prevStateRule, newStateRule)
	if err != nil {
		return nil, err
	}
	return &ContractStateSetIterator{contract: _Contract.contract, event: "StateSet", logs: logs, sub: sub}, nil
}

// WatchStateSet is a free log subscription operation binding the contract event 0xa2f9a1499fc1f9b7796d21fe5761290ccb7e0ef6ccf35fa58b668f304a62a1ca.
//
// Solidity: event StateSet(address indexed caller, uint8 indexed prevState, uint8 indexed newState, uint256 timestamp)
func (_Contract *ContractFilterer) WatchStateSet(opts *bind.WatchOpts, sink chan<- *ContractStateSet, caller []common.Address, prevState []uint8, newState []uint8) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var prevStateRule []interface{}
	for _, prevStateItem := range prevState {
		prevStateRule = append(prevStateRule, prevStateItem)
	}
	var newStateRule []interface{}
	for _, newStateItem := range newState {
		newStateRule = append(newStateRule, newStateItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "StateSet", callerRule, prevStateRule, newStateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractStateSet)
				if err := _Contract.contract.UnpackLog(event, "StateSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStateSet is a log parse operation binding the contract event 0xa2f9a1499fc1f9b7796d21fe5761290ccb7e0ef6ccf35fa58b668f304a62a1ca.
//
// Solidity: event StateSet(address indexed caller, uint8 indexed prevState, uint8 indexed newState, uint256 timestamp)
func (_Contract *ContractFilterer) ParseStateSet(log types.Log) (*ContractStateSet, error) {
	event := new(ContractStateSet)
	if err := _Contract.contract.UnpackLog(event, "StateSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
