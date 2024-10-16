// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package DataUploader

import (
	"math/big"
	"strings"

	"github.com/FISCO-BCOS/go-sdk/abi"
	"github.com/FISCO-BCOS/go-sdk/abi/bind"
	"github.com/FISCO-BCOS/go-sdk/core/types"
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
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
)

// DataUploaderABI is the input ABI used to generate the binding from.
const DataUploaderABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"}],\"name\":\"StringStored\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getStringCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"retrieveString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_data\",\"type\":\"string\"}],\"name\":\"storeString\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// DataUploaderBin is the compiled bytecode used for deploying new contracts.
var DataUploaderBin = "0x608060405234801561001057600080fd5b50610724806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80631bb8305d14610046578063508cd1aa146100625780637eddf5b414610080575b600080fd5b610060600480360381019061005b9190610429565b6100b0565b005b61006a61012d565b604051610077919061048b565b60405180910390f35b61009a600480360381019061009591906104d2565b610137565b6040516100a79190610587565b60405180910390f35b600160008154809291906100c3906105d8565b919050555080600080600154815260200190815260200160002090805190602001906100f092919061022c565b506001547e73b9f32bc5b2c870383e7aa1f8fc9bcf80d32f8b95788b9670cca8f05d8dcd826040516101229190610587565b60405180910390a250565b6000600154905090565b606060008211801561014b57506001548211155b61018a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101819061066d565b60405180910390fd5b60008083815260200190815260200160002080546101a7906106bc565b80601f01602080910402602001604051908101604052809291908181526020018280546101d3906106bc565b80156102205780601f106101f557610100808354040283529160200191610220565b820191906000526020600020905b81548152906001019060200180831161020357829003601f168201915b50505050509050919050565b828054610238906106bc565b90600052602060002090601f01602090048101928261025a57600085556102a1565b82601f1061027357805160ff19168380011785556102a1565b828001600101855582156102a1579182015b828111156102a0578251825591602001919060010190610285565b5b5090506102ae91906102b2565b5090565b5b808211156102cb5760008160009055506001016102b3565b5090565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610336826102ed565b810181811067ffffffffffffffff82111715610355576103546102fe565b5b80604052505050565b60006103686102cf565b9050610374828261032d565b919050565b600067ffffffffffffffff821115610394576103936102fe565b5b61039d826102ed565b9050602081019050919050565b82818337600083830152505050565b60006103cc6103c784610379565b61035e565b9050828152602081018484840111156103e8576103e76102e8565b5b6103f38482856103aa565b509392505050565b600082601f8301126104105761040f6102e3565b5b81356104208482602086016103b9565b91505092915050565b60006020828403121561043f5761043e6102d9565b5b600082013567ffffffffffffffff81111561045d5761045c6102de565b5b610469848285016103fb565b91505092915050565b6000819050919050565b61048581610472565b82525050565b60006020820190506104a0600083018461047c565b92915050565b6104af81610472565b81146104ba57600080fd5b50565b6000813590506104cc816104a6565b92915050565b6000602082840312156104e8576104e76102d9565b5b60006104f6848285016104bd565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60005b8381101561053957808201518184015260208101905061051e565b83811115610548576000848401525b50505050565b6000610559826104ff565b610563818561050a565b935061057381856020860161051b565b61057c816102ed565b840191505092915050565b600060208201905081810360008301526105a1818461054e565b905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006105e382610472565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415610616576106156105a9565b5b600182019050919050565b7f4944206f7574206f6620626f756e647300000000000000000000000000000000600082015250565b600061065760108361050a565b915061066282610621565b602082019050919050565b600060208201905081810360008301526106868161064a565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806106d457607f821691505b602082108114156106e8576106e761068d565b5b5091905056fea2646970667358221220651df78514b13dfc8cd1e350cd43c113f18e6115b2eb62b09290d681353eebf764736f6c634300080b0033"

// DeployDataUploader deploys a new contract, binding an instance of DataUploader to it.
func DeployDataUploader(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DataUploader, error) {
	parsed, err := abi.JSON(strings.NewReader(DataUploaderABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DataUploaderBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DataUploader{DataUploaderCaller: DataUploaderCaller{contract: contract}, DataUploaderTransactor: DataUploaderTransactor{contract: contract}, DataUploaderFilterer: DataUploaderFilterer{contract: contract}}, nil
}

func AsyncDeployDataUploader(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(DataUploaderABI))
	if err != nil {
		return nil, err
	}

	tx, err := bind.AsyncDeployContract(auth, handler, parsed, common.FromHex(DataUploaderBin), backend)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// DataUploader is an auto generated Go binding around a Solidity contract.
type DataUploader struct {
	DataUploaderCaller     // Read-only binding to the contract
	DataUploaderTransactor // Write-only binding to the contract
	DataUploaderFilterer   // Log filterer for contract events
}

// DataUploaderCaller is an auto generated read-only Go binding around a Solidity contract.
type DataUploaderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataUploaderTransactor is an auto generated write-only Go binding around a Solidity contract.
type DataUploaderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataUploaderFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type DataUploaderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataUploaderSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type DataUploaderSession struct {
	Contract     *DataUploader     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DataUploaderCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type DataUploaderCallerSession struct {
	Contract *DataUploaderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// DataUploaderTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type DataUploaderTransactorSession struct {
	Contract     *DataUploaderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// DataUploaderRaw is an auto generated low-level Go binding around a Solidity contract.
type DataUploaderRaw struct {
	Contract *DataUploader // Generic contract binding to access the raw methods on
}

// DataUploaderCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type DataUploaderCallerRaw struct {
	Contract *DataUploaderCaller // Generic read-only contract binding to access the raw methods on
}

// DataUploaderTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type DataUploaderTransactorRaw struct {
	Contract *DataUploaderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDataUploader creates a new instance of DataUploader, bound to a specific deployed contract.
func NewDataUploader(address common.Address, backend bind.ContractBackend) (*DataUploader, error) {
	contract, err := bindDataUploader(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DataUploader{DataUploaderCaller: DataUploaderCaller{contract: contract}, DataUploaderTransactor: DataUploaderTransactor{contract: contract}, DataUploaderFilterer: DataUploaderFilterer{contract: contract}}, nil
}

// NewDataUploaderCaller creates a new read-only instance of DataUploader, bound to a specific deployed contract.
func NewDataUploaderCaller(address common.Address, caller bind.ContractCaller) (*DataUploaderCaller, error) {
	contract, err := bindDataUploader(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DataUploaderCaller{contract: contract}, nil
}

// NewDataUploaderTransactor creates a new write-only instance of DataUploader, bound to a specific deployed contract.
func NewDataUploaderTransactor(address common.Address, transactor bind.ContractTransactor) (*DataUploaderTransactor, error) {
	contract, err := bindDataUploader(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DataUploaderTransactor{contract: contract}, nil
}

// NewDataUploaderFilterer creates a new log filterer instance of DataUploader, bound to a specific deployed contract.
func NewDataUploaderFilterer(address common.Address, filterer bind.ContractFilterer) (*DataUploaderFilterer, error) {
	contract, err := bindDataUploader(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DataUploaderFilterer{contract: contract}, nil
}

// bindDataUploader binds a generic wrapper to an already deployed contract.
func bindDataUploader(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DataUploaderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DataUploader *DataUploaderRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DataUploader.Contract.DataUploaderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DataUploader *DataUploaderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _DataUploader.Contract.DataUploaderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DataUploader *DataUploaderRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _DataUploader.Contract.DataUploaderTransactor.contract.TransactWithResult(opts, result, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DataUploader *DataUploaderCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DataUploader.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DataUploader *DataUploaderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _DataUploader.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DataUploader *DataUploaderTransactorRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _DataUploader.Contract.contract.TransactWithResult(opts, result, method, params...)
}

// GetStringCount is a free data retrieval call binding the contract method 0x508cd1aa.
//
// Solidity: function getStringCount() constant returns(uint256)
func (_DataUploader *DataUploaderCaller) GetStringCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DataUploader.contract.Call(opts, out, "getStringCount")
	return *ret0, err
}

// GetStringCount is a free data retrieval call binding the contract method 0x508cd1aa.
//
// Solidity: function getStringCount() constant returns(uint256)
func (_DataUploader *DataUploaderSession) GetStringCount() (*big.Int, error) {
	return _DataUploader.Contract.GetStringCount(&_DataUploader.CallOpts)
}

// GetStringCount is a free data retrieval call binding the contract method 0x508cd1aa.
//
// Solidity: function getStringCount() constant returns(uint256)
func (_DataUploader *DataUploaderCallerSession) GetStringCount() (*big.Int, error) {
	return _DataUploader.Contract.GetStringCount(&_DataUploader.CallOpts)
}

// RetrieveString is a free data retrieval call binding the contract method 0x7eddf5b4.
//
// Solidity: function retrieveString(uint256 _id) constant returns(string)
func (_DataUploader *DataUploaderCaller) RetrieveString(opts *bind.CallOpts, _id *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _DataUploader.contract.Call(opts, out, "retrieveString", _id)
	return *ret0, err
}

// RetrieveString is a free data retrieval call binding the contract method 0x7eddf5b4.
//
// Solidity: function retrieveString(uint256 _id) constant returns(string)
func (_DataUploader *DataUploaderSession) RetrieveString(_id *big.Int) (string, error) {
	return _DataUploader.Contract.RetrieveString(&_DataUploader.CallOpts, _id)
}

// RetrieveString is a free data retrieval call binding the contract method 0x7eddf5b4.
//
// Solidity: function retrieveString(uint256 _id) constant returns(string)
func (_DataUploader *DataUploaderCallerSession) RetrieveString(_id *big.Int) (string, error) {
	return _DataUploader.Contract.RetrieveString(&_DataUploader.CallOpts, _id)
}

// StoreString is a paid mutator transaction binding the contract method 0x1bb8305d.
//
// Solidity: function storeString(string _data) returns()
func (_DataUploader *DataUploaderTransactor) StoreString(opts *bind.TransactOpts, _data string) (*types.Transaction, *types.Receipt, error) {
	var ()
	out := &[]interface{}{}
	transaction, receipt, err := _DataUploader.contract.TransactWithResult(opts, out, "storeString", _data)
	return transaction, receipt, err
}

func (_DataUploader *DataUploaderTransactor) AsyncStoreString(handler func(*types.Receipt, error), opts *bind.TransactOpts, _data string) (*types.Transaction, error) {
	return _DataUploader.contract.AsyncTransact(opts, handler, "storeString", _data)
}

// StoreString is a paid mutator transaction binding the contract method 0x1bb8305d.
//
// Solidity: function storeString(string _data) returns()
func (_DataUploader *DataUploaderSession) StoreString(_data string) (*types.Transaction, *types.Receipt, error) {
	return _DataUploader.Contract.StoreString(&_DataUploader.TransactOpts, _data)
}

func (_DataUploader *DataUploaderSession) AsyncStoreString(handler func(*types.Receipt, error), _data string) (*types.Transaction, error) {
	return _DataUploader.Contract.AsyncStoreString(handler, &_DataUploader.TransactOpts, _data)
}

// StoreString is a paid mutator transaction binding the contract method 0x1bb8305d.
//
// Solidity: function storeString(string _data) returns()
func (_DataUploader *DataUploaderTransactorSession) StoreString(_data string) (*types.Transaction, *types.Receipt, error) {
	return _DataUploader.Contract.StoreString(&_DataUploader.TransactOpts, _data)
}

func (_DataUploader *DataUploaderTransactorSession) AsyncStoreString(handler func(*types.Receipt, error), _data string) (*types.Transaction, error) {
	return _DataUploader.Contract.AsyncStoreString(handler, &_DataUploader.TransactOpts, _data)
}

// DataUploaderStringStored represents a StringStored event raised by the DataUploader contract.
type DataUploaderStringStored struct {
	Id   *big.Int
	Data string
	Raw  types.Log // Blockchain specific contextual infos
}

// WatchStringStored is a free log subscription operation binding the contract event 0x0073b9f32bc5b2c870383e7aa1f8fc9bcf80d32f8b95788b9670cca8f05d8dcd.
//
// Solidity: event StringStored(uint256 indexed id, string data)
func (_DataUploader *DataUploaderFilterer) WatchStringStored(fromBlock *uint64, handler func(int, []types.Log), id *big.Int) (string, error) {
	return _DataUploader.contract.WatchLogs(fromBlock, handler, "StringStored", id)
}

func (_DataUploader *DataUploaderFilterer) WatchAllStringStored(fromBlock *uint64, handler func(int, []types.Log)) (string, error) {
	return _DataUploader.contract.WatchLogs(fromBlock, handler, "StringStored")
}

// ParseStringStored is a log parse operation binding the contract event 0x0073b9f32bc5b2c870383e7aa1f8fc9bcf80d32f8b95788b9670cca8f05d8dcd.
//
// Solidity: event StringStored(uint256 indexed id, string data)
func (_DataUploader *DataUploaderFilterer) ParseStringStored(log types.Log) (*DataUploaderStringStored, error) {
	event := new(DataUploaderStringStored)
	if err := _DataUploader.contract.UnpackLog(event, "StringStored", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchStringStored is a free log subscription operation binding the contract event 0x0073b9f32bc5b2c870383e7aa1f8fc9bcf80d32f8b95788b9670cca8f05d8dcd.
//
// Solidity: event StringStored(uint256 indexed id, string data)
func (_DataUploader *DataUploaderSession) WatchStringStored(fromBlock *uint64, handler func(int, []types.Log), id *big.Int) (string, error) {
	return _DataUploader.Contract.WatchStringStored(fromBlock, handler, id)
}

func (_DataUploader *DataUploaderSession) WatchAllStringStored(fromBlock *uint64, handler func(int, []types.Log)) (string, error) {
	return _DataUploader.Contract.WatchAllStringStored(fromBlock, handler)
}

// ParseStringStored is a log parse operation binding the contract event 0x0073b9f32bc5b2c870383e7aa1f8fc9bcf80d32f8b95788b9670cca8f05d8dcd.
//
// Solidity: event StringStored(uint256 indexed id, string data)
func (_DataUploader *DataUploaderSession) ParseStringStored(log types.Log) (*DataUploaderStringStored, error) {
	return _DataUploader.Contract.ParseStringStored(log)
}
