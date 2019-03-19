// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package store

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// StoreABI is the input ABI used to generate the binding from.
const StoreABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"nearest_pole_id\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"latitude\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_nearest_pole_id\",\"type\":\"string\"},{\"name\":\"_latitude\",\"type\":\"string\"},{\"name\":\"_longitude\",\"type\":\"string\"},{\"name\":\"_timestamp\",\"type\":\"string\"}],\"name\":\"setData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"longitude\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"timestamp\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_nearest_pole_id\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_latitude\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_longitude\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_timestamp\",\"type\":\"string\"}],\"name\":\"dataSetter\",\"type\":\"event\"}]"

// StoreBin is the compiled bytecode used for deploying new contracts.
const StoreBin = `608060405234801561001057600080fd5b5061077b806100206000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c8063397955e31461005c5780634fd7d76a146100df57806353104d4414610162578063589af69c146102da578063b80777ea1461035d575b600080fd5b6100646103e0565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156100a4578082015181840152602081019050610089565b50505050905090810190601f1680156100d15780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6100e761047e565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561012757808201518184015260208101905061010c565b50505050905090810190601f1680156101545780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6102d86004803603608081101561017857600080fd5b810190808035906020019064010000000081111561019557600080fd5b8201836020820111156101a757600080fd5b803590602001918460018302840111640100000000831117156101c957600080fd5b9091929391929390803590602001906401000000008111156101ea57600080fd5b8201836020820111156101fc57600080fd5b8035906020019184600183028401116401000000008311171561021e57600080fd5b90919293919293908035906020019064010000000081111561023f57600080fd5b82018360208201111561025157600080fd5b8035906020019184600183028401116401000000008311171561027357600080fd5b90919293919293908035906020019064010000000081111561029457600080fd5b8201836020820111156102a657600080fd5b803590602001918460018302840111640100000000831117156102c857600080fd5b909192939192939050505061051c565b005b6102e261056e565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610322578082015181840152602081019050610307565b50505050905090810190601f16801561034f5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b61036561060c565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156103a557808201518184015260208101905061038a565b50505050905090810190601f1680156103d25780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b60038054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156104765780601f1061044b57610100808354040283529160200191610476565b820191906000526020600020905b81548152906001019060200180831161045957829003601f168201915b505050505081565b60008054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156105145780601f106104e957610100808354040283529160200191610514565b820191906000526020600020905b8154815290600101906020018083116104f757829003601f168201915b505050505081565b85856000919061052d9291906106aa565b5083836001919061053f9291906106aa565b508787600391906105519291906106aa565b508181600291906105639291906106aa565b505050505050505050565b60018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156106045780601f106105d957610100808354040283529160200191610604565b820191906000526020600020905b8154815290600101906020018083116105e757829003601f168201915b505050505081565b60028054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156106a25780601f10610677576101008083540402835291602001916106a2565b820191906000526020600020905b81548152906001019060200180831161068557829003601f168201915b505050505081565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106106eb57803560ff1916838001178555610719565b82800160010185558215610719579182015b828111156107185782358255916020019190600101906106fd565b5b509050610726919061072a565b5090565b61074c91905b80821115610748576000816000905550600101610730565b5090565b9056fea165627a7a723058207e8f0bb116be727fae1cf36f2813a1f65eb95213ee5bfc6c0d17a5f70bdf50310029`

// DeployStore deploys a new Ethereum contract, binding an instance of Store to it.
func DeployStore(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Store, error) {
	parsed, err := abi.JSON(strings.NewReader(StoreABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StoreBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Store{StoreCaller: StoreCaller{contract: contract}, StoreTransactor: StoreTransactor{contract: contract}, StoreFilterer: StoreFilterer{contract: contract}}, nil
}

// Store is an auto generated Go binding around an Ethereum contract.
type Store struct {
	StoreCaller     // Read-only binding to the contract
	StoreTransactor // Write-only binding to the contract
	StoreFilterer   // Log filterer for contract events
}

// StoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type StoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StoreSession struct {
	Contract     *Store            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StoreCallerSession struct {
	Contract *StoreCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// StoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StoreTransactorSession struct {
	Contract     *StoreTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type StoreRaw struct {
	Contract *Store // Generic contract binding to access the raw methods on
}

// StoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StoreCallerRaw struct {
	Contract *StoreCaller // Generic read-only contract binding to access the raw methods on
}

// StoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StoreTransactorRaw struct {
	Contract *StoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStore creates a new instance of Store, bound to a specific deployed contract.
func NewStore(address common.Address, backend bind.ContractBackend) (*Store, error) {
	contract, err := bindStore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Store{StoreCaller: StoreCaller{contract: contract}, StoreTransactor: StoreTransactor{contract: contract}, StoreFilterer: StoreFilterer{contract: contract}}, nil
}

// NewStoreCaller creates a new read-only instance of Store, bound to a specific deployed contract.
func NewStoreCaller(address common.Address, caller bind.ContractCaller) (*StoreCaller, error) {
	contract, err := bindStore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StoreCaller{contract: contract}, nil
}

// NewStoreTransactor creates a new write-only instance of Store, bound to a specific deployed contract.
func NewStoreTransactor(address common.Address, transactor bind.ContractTransactor) (*StoreTransactor, error) {
	contract, err := bindStore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StoreTransactor{contract: contract}, nil
}

// NewStoreFilterer creates a new log filterer instance of Store, bound to a specific deployed contract.
func NewStoreFilterer(address common.Address, filterer bind.ContractFilterer) (*StoreFilterer, error) {
	contract, err := bindStore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StoreFilterer{contract: contract}, nil
}

// bindStore binds a generic wrapper to an already deployed contract.
func bindStore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StoreABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Store.Contract.StoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Store.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.contract.Transact(opts, method, params...)
}

// Latitude is a free data retrieval call binding the contract method 0x4fd7d76a.
//
// Solidity: function latitude() constant returns(string)
func (_Store *StoreCaller) Latitude(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Store.contract.Call(opts, out, "latitude")
	return *ret0, err
}

// Latitude is a free data retrieval call binding the contract method 0x4fd7d76a.
//
// Solidity: function latitude() constant returns(string)
func (_Store *StoreSession) Latitude() (string, error) {
	return _Store.Contract.Latitude(&_Store.CallOpts)
}

// Latitude is a free data retrieval call binding the contract method 0x4fd7d76a.
//
// Solidity: function latitude() constant returns(string)
func (_Store *StoreCallerSession) Latitude() (string, error) {
	return _Store.Contract.Latitude(&_Store.CallOpts)
}

// Longitude is a free data retrieval call binding the contract method 0x589af69c.
//
// Solidity: function longitude() constant returns(string)
func (_Store *StoreCaller) Longitude(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Store.contract.Call(opts, out, "longitude")
	return *ret0, err
}

// Longitude is a free data retrieval call binding the contract method 0x589af69c.
//
// Solidity: function longitude() constant returns(string)
func (_Store *StoreSession) Longitude() (string, error) {
	return _Store.Contract.Longitude(&_Store.CallOpts)
}

// Longitude is a free data retrieval call binding the contract method 0x589af69c.
//
// Solidity: function longitude() constant returns(string)
func (_Store *StoreCallerSession) Longitude() (string, error) {
	return _Store.Contract.Longitude(&_Store.CallOpts)
}

// NearestPoleId is a free data retrieval call binding the contract method 0x397955e3.
//
// Solidity: function nearest_pole_id() constant returns(string)
func (_Store *StoreCaller) NearestPoleId(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Store.contract.Call(opts, out, "nearest_pole_id")
	return *ret0, err
}

// NearestPoleId is a free data retrieval call binding the contract method 0x397955e3.
//
// Solidity: function nearest_pole_id() constant returns(string)
func (_Store *StoreSession) NearestPoleId() (string, error) {
	return _Store.Contract.NearestPoleId(&_Store.CallOpts)
}

// NearestPoleId is a free data retrieval call binding the contract method 0x397955e3.
//
// Solidity: function nearest_pole_id() constant returns(string)
func (_Store *StoreCallerSession) NearestPoleId() (string, error) {
	return _Store.Contract.NearestPoleId(&_Store.CallOpts)
}

// Timestamp is a free data retrieval call binding the contract method 0xb80777ea.
//
// Solidity: function timestamp() constant returns(string)
func (_Store *StoreCaller) Timestamp(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Store.contract.Call(opts, out, "timestamp")
	return *ret0, err
}

// Timestamp is a free data retrieval call binding the contract method 0xb80777ea.
//
// Solidity: function timestamp() constant returns(string)
func (_Store *StoreSession) Timestamp() (string, error) {
	return _Store.Contract.Timestamp(&_Store.CallOpts)
}

// Timestamp is a free data retrieval call binding the contract method 0xb80777ea.
//
// Solidity: function timestamp() constant returns(string)
func (_Store *StoreCallerSession) Timestamp() (string, error) {
	return _Store.Contract.Timestamp(&_Store.CallOpts)
}

// SetData is a paid mutator transaction binding the contract method 0x53104d44.
//
// Solidity: function setData(string _nearest_pole_id, string _latitude, string _longitude, string _timestamp) returns()
func (_Store *StoreTransactor) SetData(opts *bind.TransactOpts, _nearest_pole_id string, _latitude string, _longitude string, _timestamp string) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "setData", _nearest_pole_id, _latitude, _longitude, _timestamp)
}

// SetData is a paid mutator transaction binding the contract method 0x53104d44.
//
// Solidity: function setData(string _nearest_pole_id, string _latitude, string _longitude, string _timestamp) returns()
func (_Store *StoreSession) SetData(_nearest_pole_id string, _latitude string, _longitude string, _timestamp string) (*types.Transaction, error) {
	return _Store.Contract.SetData(&_Store.TransactOpts, _nearest_pole_id, _latitude, _longitude, _timestamp)
}

// SetData is a paid mutator transaction binding the contract method 0x53104d44.
//
// Solidity: function setData(string _nearest_pole_id, string _latitude, string _longitude, string _timestamp) returns()
func (_Store *StoreTransactorSession) SetData(_nearest_pole_id string, _latitude string, _longitude string, _timestamp string) (*types.Transaction, error) {
	return _Store.Contract.SetData(&_Store.TransactOpts, _nearest_pole_id, _latitude, _longitude, _timestamp)
}

// StoreDataSetterIterator is returned from FilterDataSetter and is used to iterate over the raw logs and unpacked data for DataSetter events raised by the Store contract.
type StoreDataSetterIterator struct {
	Event *StoreDataSetter // Event containing the contract specifics and raw log

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
func (it *StoreDataSetterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreDataSetter)
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
		it.Event = new(StoreDataSetter)
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
func (it *StoreDataSetterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreDataSetterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreDataSetter represents a DataSetter event raised by the Store contract.
type StoreDataSetter struct {
	NearestPoleId string
	Latitude      string
	Longitude     string
	Timestamp     string
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDataSetter is a free log retrieval operation binding the contract event 0xddadf5f419c9519542ffd9195795f0f36b8c024dc4c74a9a08f3acf95c78d153.
//
// Solidity: event dataSetter(string _nearest_pole_id, string _latitude, string _longitude, string _timestamp)
func (_Store *StoreFilterer) FilterDataSetter(opts *bind.FilterOpts) (*StoreDataSetterIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "dataSetter")
	if err != nil {
		return nil, err
	}
	return &StoreDataSetterIterator{contract: _Store.contract, event: "dataSetter", logs: logs, sub: sub}, nil
}

// WatchDataSetter is a free log subscription operation binding the contract event 0xddadf5f419c9519542ffd9195795f0f36b8c024dc4c74a9a08f3acf95c78d153.
//
// Solidity: event dataSetter(string _nearest_pole_id, string _latitude, string _longitude, string _timestamp)
func (_Store *StoreFilterer) WatchDataSetter(opts *bind.WatchOpts, sink chan<- *StoreDataSetter) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "dataSetter")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreDataSetter)
				if err := _Store.contract.UnpackLog(event, "dataSetter", log); err != nil {
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
