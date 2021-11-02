// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// NFTMarketV2MarketItem is an auto generated low-level Go binding around an user-defined struct.
type NFTMarketV2MarketItem struct {
	ItemId      *big.Int
	NftContract common.Address
	TokenId     *big.Int
	Seller      common.Address
	Owner       common.Address
	Price       *big.Int
	Sold        bool
}

// NFTMarketMetaData contains all meta data concerning the NFTMarket contract.
var NFTMarketMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"itemId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"sold\",\"type\":\"bool\"}],\"name\":\"MarketItemCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"itemId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"MarketItemSold\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"createMarketItem\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"itemId\",\"type\":\"uint256\"}],\"name\":\"createMarketSale\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fetchItemsCreated\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"itemId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"sold\",\"type\":\"bool\"}],\"internalType\":\"structNFTMarketV2.MarketItem[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fetchMarketItems\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"itemId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"sold\",\"type\":\"bool\"}],\"internalType\":\"structNFTMarketV2.MarketItem[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fetchMyNFTs\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"itemId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"sold\",\"type\":\"bool\"}],\"internalType\":\"structNFTMarketV2.MarketItem[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getListingPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// NFTMarketABI is the input ABI used to generate the binding from.
// Deprecated: Use NFTMarketMetaData.ABI instead.
var NFTMarketABI = NFTMarketMetaData.ABI

// NFTMarket is an auto generated Go binding around an Ethereum contract.
type NFTMarket struct {
	NFTMarketCaller     // Read-only binding to the contract
	NFTMarketTransactor // Write-only binding to the contract
	NFTMarketFilterer   // Log filterer for contract events
}

// NFTMarketCaller is an auto generated read-only Go binding around an Ethereum contract.
type NFTMarketCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NFTMarketTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NFTMarketTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NFTMarketFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NFTMarketFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NFTMarketSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NFTMarketSession struct {
	Contract     *NFTMarket        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NFTMarketCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NFTMarketCallerSession struct {
	Contract *NFTMarketCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// NFTMarketTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NFTMarketTransactorSession struct {
	Contract     *NFTMarketTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// NFTMarketRaw is an auto generated low-level Go binding around an Ethereum contract.
type NFTMarketRaw struct {
	Contract *NFTMarket // Generic contract binding to access the raw methods on
}

// NFTMarketCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NFTMarketCallerRaw struct {
	Contract *NFTMarketCaller // Generic read-only contract binding to access the raw methods on
}

// NFTMarketTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NFTMarketTransactorRaw struct {
	Contract *NFTMarketTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNFTMarket creates a new instance of NFTMarket, bound to a specific deployed contract.
func NewNFTMarket(address common.Address, backend bind.ContractBackend) (*NFTMarket, error) {
	contract, err := bindNFTMarket(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NFTMarket{NFTMarketCaller: NFTMarketCaller{contract: contract}, NFTMarketTransactor: NFTMarketTransactor{contract: contract}, NFTMarketFilterer: NFTMarketFilterer{contract: contract}}, nil
}

// NewNFTMarketCaller creates a new read-only instance of NFTMarket, bound to a specific deployed contract.
func NewNFTMarketCaller(address common.Address, caller bind.ContractCaller) (*NFTMarketCaller, error) {
	contract, err := bindNFTMarket(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NFTMarketCaller{contract: contract}, nil
}

// NewNFTMarketTransactor creates a new write-only instance of NFTMarket, bound to a specific deployed contract.
func NewNFTMarketTransactor(address common.Address, transactor bind.ContractTransactor) (*NFTMarketTransactor, error) {
	contract, err := bindNFTMarket(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NFTMarketTransactor{contract: contract}, nil
}

// NewNFTMarketFilterer creates a new log filterer instance of NFTMarket, bound to a specific deployed contract.
func NewNFTMarketFilterer(address common.Address, filterer bind.ContractFilterer) (*NFTMarketFilterer, error) {
	contract, err := bindNFTMarket(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NFTMarketFilterer{contract: contract}, nil
}

// bindNFTMarket binds a generic wrapper to an already deployed contract.
func bindNFTMarket(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NFTMarketABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NFTMarket *NFTMarketRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NFTMarket.Contract.NFTMarketCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NFTMarket *NFTMarketRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NFTMarket.Contract.NFTMarketTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NFTMarket *NFTMarketRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NFTMarket.Contract.NFTMarketTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NFTMarket *NFTMarketCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NFTMarket.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NFTMarket *NFTMarketTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NFTMarket.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NFTMarket *NFTMarketTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NFTMarket.Contract.contract.Transact(opts, method, params...)
}

// FetchItemsCreated is a free data retrieval call binding the contract method 0xf064c32e.
//
// Solidity: function fetchItemsCreated() view returns((uint256,address,uint256,address,address,uint256,bool)[])
func (_NFTMarket *NFTMarketCaller) FetchItemsCreated(opts *bind.CallOpts) ([]NFTMarketV2MarketItem, error) {
	var out []interface{}
	err := _NFTMarket.contract.Call(opts, &out, "fetchItemsCreated")

	if err != nil {
		return *new([]NFTMarketV2MarketItem), err
	}

	out0 := *abi.ConvertType(out[0], new([]NFTMarketV2MarketItem)).(*[]NFTMarketV2MarketItem)

	return out0, err

}

// FetchItemsCreated is a free data retrieval call binding the contract method 0xf064c32e.
//
// Solidity: function fetchItemsCreated() view returns((uint256,address,uint256,address,address,uint256,bool)[])
func (_NFTMarket *NFTMarketSession) FetchItemsCreated() ([]NFTMarketV2MarketItem, error) {
	return _NFTMarket.Contract.FetchItemsCreated(&_NFTMarket.CallOpts)
}

// FetchItemsCreated is a free data retrieval call binding the contract method 0xf064c32e.
//
// Solidity: function fetchItemsCreated() view returns((uint256,address,uint256,address,address,uint256,bool)[])
func (_NFTMarket *NFTMarketCallerSession) FetchItemsCreated() ([]NFTMarketV2MarketItem, error) {
	return _NFTMarket.Contract.FetchItemsCreated(&_NFTMarket.CallOpts)
}

// FetchMarketItems is a free data retrieval call binding the contract method 0x0f08efe0.
//
// Solidity: function fetchMarketItems() view returns((uint256,address,uint256,address,address,uint256,bool)[])
func (_NFTMarket *NFTMarketCaller) FetchMarketItems(opts *bind.CallOpts) ([]NFTMarketV2MarketItem, error) {
	var out []interface{}
	err := _NFTMarket.contract.Call(opts, &out, "fetchMarketItems")

	if err != nil {
		return *new([]NFTMarketV2MarketItem), err
	}

	out0 := *abi.ConvertType(out[0], new([]NFTMarketV2MarketItem)).(*[]NFTMarketV2MarketItem)

	return out0, err

}

// FetchMarketItems is a free data retrieval call binding the contract method 0x0f08efe0.
//
// Solidity: function fetchMarketItems() view returns((uint256,address,uint256,address,address,uint256,bool)[])
func (_NFTMarket *NFTMarketSession) FetchMarketItems() ([]NFTMarketV2MarketItem, error) {
	return _NFTMarket.Contract.FetchMarketItems(&_NFTMarket.CallOpts)
}

// FetchMarketItems is a free data retrieval call binding the contract method 0x0f08efe0.
//
// Solidity: function fetchMarketItems() view returns((uint256,address,uint256,address,address,uint256,bool)[])
func (_NFTMarket *NFTMarketCallerSession) FetchMarketItems() ([]NFTMarketV2MarketItem, error) {
	return _NFTMarket.Contract.FetchMarketItems(&_NFTMarket.CallOpts)
}

// FetchMyNFTs is a free data retrieval call binding the contract method 0x202e3740.
//
// Solidity: function fetchMyNFTs() view returns((uint256,address,uint256,address,address,uint256,bool)[])
func (_NFTMarket *NFTMarketCaller) FetchMyNFTs(opts *bind.CallOpts) ([]NFTMarketV2MarketItem, error) {
	var out []interface{}
	err := _NFTMarket.contract.Call(opts, &out, "fetchMyNFTs")

	if err != nil {
		return *new([]NFTMarketV2MarketItem), err
	}

	out0 := *abi.ConvertType(out[0], new([]NFTMarketV2MarketItem)).(*[]NFTMarketV2MarketItem)

	return out0, err

}

// FetchMyNFTs is a free data retrieval call binding the contract method 0x202e3740.
//
// Solidity: function fetchMyNFTs() view returns((uint256,address,uint256,address,address,uint256,bool)[])
func (_NFTMarket *NFTMarketSession) FetchMyNFTs() ([]NFTMarketV2MarketItem, error) {
	return _NFTMarket.Contract.FetchMyNFTs(&_NFTMarket.CallOpts)
}

// FetchMyNFTs is a free data retrieval call binding the contract method 0x202e3740.
//
// Solidity: function fetchMyNFTs() view returns((uint256,address,uint256,address,address,uint256,bool)[])
func (_NFTMarket *NFTMarketCallerSession) FetchMyNFTs() ([]NFTMarketV2MarketItem, error) {
	return _NFTMarket.Contract.FetchMyNFTs(&_NFTMarket.CallOpts)
}

// GetListingPrice is a free data retrieval call binding the contract method 0x12e85585.
//
// Solidity: function getListingPrice() view returns(uint256)
func (_NFTMarket *NFTMarketCaller) GetListingPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NFTMarket.contract.Call(opts, &out, "getListingPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetListingPrice is a free data retrieval call binding the contract method 0x12e85585.
//
// Solidity: function getListingPrice() view returns(uint256)
func (_NFTMarket *NFTMarketSession) GetListingPrice() (*big.Int, error) {
	return _NFTMarket.Contract.GetListingPrice(&_NFTMarket.CallOpts)
}

// GetListingPrice is a free data retrieval call binding the contract method 0x12e85585.
//
// Solidity: function getListingPrice() view returns(uint256)
func (_NFTMarket *NFTMarketCallerSession) GetListingPrice() (*big.Int, error) {
	return _NFTMarket.Contract.GetListingPrice(&_NFTMarket.CallOpts)
}

// CreateMarketItem is a paid mutator transaction binding the contract method 0x58eb2df5.
//
// Solidity: function createMarketItem(address nftContract, uint256 tokenId, uint256 price) payable returns()
func (_NFTMarket *NFTMarketTransactor) CreateMarketItem(opts *bind.TransactOpts, nftContract common.Address, tokenId *big.Int, price *big.Int) (*types.Transaction, error) {
	return _NFTMarket.contract.Transact(opts, "createMarketItem", nftContract, tokenId, price)
}

// CreateMarketItem is a paid mutator transaction binding the contract method 0x58eb2df5.
//
// Solidity: function createMarketItem(address nftContract, uint256 tokenId, uint256 price) payable returns()
func (_NFTMarket *NFTMarketSession) CreateMarketItem(nftContract common.Address, tokenId *big.Int, price *big.Int) (*types.Transaction, error) {
	return _NFTMarket.Contract.CreateMarketItem(&_NFTMarket.TransactOpts, nftContract, tokenId, price)
}

// CreateMarketItem is a paid mutator transaction binding the contract method 0x58eb2df5.
//
// Solidity: function createMarketItem(address nftContract, uint256 tokenId, uint256 price) payable returns()
func (_NFTMarket *NFTMarketTransactorSession) CreateMarketItem(nftContract common.Address, tokenId *big.Int, price *big.Int) (*types.Transaction, error) {
	return _NFTMarket.Contract.CreateMarketItem(&_NFTMarket.TransactOpts, nftContract, tokenId, price)
}

// CreateMarketSale is a paid mutator transaction binding the contract method 0xc23b139e.
//
// Solidity: function createMarketSale(address nftContract, uint256 itemId) payable returns()
func (_NFTMarket *NFTMarketTransactor) CreateMarketSale(opts *bind.TransactOpts, nftContract common.Address, itemId *big.Int) (*types.Transaction, error) {
	return _NFTMarket.contract.Transact(opts, "createMarketSale", nftContract, itemId)
}

// CreateMarketSale is a paid mutator transaction binding the contract method 0xc23b139e.
//
// Solidity: function createMarketSale(address nftContract, uint256 itemId) payable returns()
func (_NFTMarket *NFTMarketSession) CreateMarketSale(nftContract common.Address, itemId *big.Int) (*types.Transaction, error) {
	return _NFTMarket.Contract.CreateMarketSale(&_NFTMarket.TransactOpts, nftContract, itemId)
}

// CreateMarketSale is a paid mutator transaction binding the contract method 0xc23b139e.
//
// Solidity: function createMarketSale(address nftContract, uint256 itemId) payable returns()
func (_NFTMarket *NFTMarketTransactorSession) CreateMarketSale(nftContract common.Address, itemId *big.Int) (*types.Transaction, error) {
	return _NFTMarket.Contract.CreateMarketSale(&_NFTMarket.TransactOpts, nftContract, itemId)
}

// NFTMarketMarketItemCreatedIterator is returned from FilterMarketItemCreated and is used to iterate over the raw logs and unpacked data for MarketItemCreated events raised by the NFTMarket contract.
type NFTMarketMarketItemCreatedIterator struct {
	Event *NFTMarketMarketItemCreated // Event containing the contract specifics and raw log

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
func (it *NFTMarketMarketItemCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTMarketMarketItemCreated)
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
		it.Event = new(NFTMarketMarketItemCreated)
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
func (it *NFTMarketMarketItemCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTMarketMarketItemCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTMarketMarketItemCreated represents a MarketItemCreated event raised by the NFTMarket contract.
type NFTMarketMarketItemCreated struct {
	ItemId      *big.Int
	NftContract common.Address
	TokenId     *big.Int
	Seller      common.Address
	Owner       common.Address
	Price       *big.Int
	Sold        bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMarketItemCreated is a free log retrieval operation binding the contract event 0x045dfa01dcba2b36aba1d3dc4a874f4b0c5d2fbeb8d2c4b34a7d88c8d8f929d1.
//
// Solidity: event MarketItemCreated(uint256 indexed itemId, address indexed nftContract, uint256 indexed tokenId, address seller, address owner, uint256 price, bool sold)
func (_NFTMarket *NFTMarketFilterer) FilterMarketItemCreated(opts *bind.FilterOpts, itemId []*big.Int, nftContract []common.Address, tokenId []*big.Int) (*NFTMarketMarketItemCreatedIterator, error) {

	var itemIdRule []interface{}
	for _, itemIdItem := range itemId {
		itemIdRule = append(itemIdRule, itemIdItem)
	}
	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _NFTMarket.contract.FilterLogs(opts, "MarketItemCreated", itemIdRule, nftContractRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &NFTMarketMarketItemCreatedIterator{contract: _NFTMarket.contract, event: "MarketItemCreated", logs: logs, sub: sub}, nil
}

// WatchMarketItemCreated is a free log subscription operation binding the contract event 0x045dfa01dcba2b36aba1d3dc4a874f4b0c5d2fbeb8d2c4b34a7d88c8d8f929d1.
//
// Solidity: event MarketItemCreated(uint256 indexed itemId, address indexed nftContract, uint256 indexed tokenId, address seller, address owner, uint256 price, bool sold)
func (_NFTMarket *NFTMarketFilterer) WatchMarketItemCreated(opts *bind.WatchOpts, sink chan<- *NFTMarketMarketItemCreated, itemId []*big.Int, nftContract []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var itemIdRule []interface{}
	for _, itemIdItem := range itemId {
		itemIdRule = append(itemIdRule, itemIdItem)
	}
	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _NFTMarket.contract.WatchLogs(opts, "MarketItemCreated", itemIdRule, nftContractRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTMarketMarketItemCreated)
				if err := _NFTMarket.contract.UnpackLog(event, "MarketItemCreated", log); err != nil {
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

// ParseMarketItemCreated is a log parse operation binding the contract event 0x045dfa01dcba2b36aba1d3dc4a874f4b0c5d2fbeb8d2c4b34a7d88c8d8f929d1.
//
// Solidity: event MarketItemCreated(uint256 indexed itemId, address indexed nftContract, uint256 indexed tokenId, address seller, address owner, uint256 price, bool sold)
func (_NFTMarket *NFTMarketFilterer) ParseMarketItemCreated(log types.Log) (*NFTMarketMarketItemCreated, error) {
	event := new(NFTMarketMarketItemCreated)
	if err := _NFTMarket.contract.UnpackLog(event, "MarketItemCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTMarketMarketItemSoldIterator is returned from FilterMarketItemSold and is used to iterate over the raw logs and unpacked data for MarketItemSold events raised by the NFTMarket contract.
type NFTMarketMarketItemSoldIterator struct {
	Event *NFTMarketMarketItemSold // Event containing the contract specifics and raw log

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
func (it *NFTMarketMarketItemSoldIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTMarketMarketItemSold)
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
		it.Event = new(NFTMarketMarketItemSold)
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
func (it *NFTMarketMarketItemSoldIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTMarketMarketItemSoldIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTMarketMarketItemSold represents a MarketItemSold event raised by the NFTMarket contract.
type NFTMarketMarketItemSold struct {
	ItemId      *big.Int
	NftContract common.Address
	TokenId     *big.Int
	Seller      common.Address
	Owner       common.Address
	Price       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMarketItemSold is a free log retrieval operation binding the contract event 0x6a5f9d2b6ff1f375ebe095556dc91e68fce80a6187d7527cce81c3ab02726c9c.
//
// Solidity: event MarketItemSold(uint256 indexed itemId, address indexed nftContract, uint256 indexed tokenId, address seller, address owner, uint256 price)
func (_NFTMarket *NFTMarketFilterer) FilterMarketItemSold(opts *bind.FilterOpts, itemId []*big.Int, nftContract []common.Address, tokenId []*big.Int) (*NFTMarketMarketItemSoldIterator, error) {

	var itemIdRule []interface{}
	for _, itemIdItem := range itemId {
		itemIdRule = append(itemIdRule, itemIdItem)
	}
	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _NFTMarket.contract.FilterLogs(opts, "MarketItemSold", itemIdRule, nftContractRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &NFTMarketMarketItemSoldIterator{contract: _NFTMarket.contract, event: "MarketItemSold", logs: logs, sub: sub}, nil
}

// WatchMarketItemSold is a free log subscription operation binding the contract event 0x6a5f9d2b6ff1f375ebe095556dc91e68fce80a6187d7527cce81c3ab02726c9c.
//
// Solidity: event MarketItemSold(uint256 indexed itemId, address indexed nftContract, uint256 indexed tokenId, address seller, address owner, uint256 price)
func (_NFTMarket *NFTMarketFilterer) WatchMarketItemSold(opts *bind.WatchOpts, sink chan<- *NFTMarketMarketItemSold, itemId []*big.Int, nftContract []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var itemIdRule []interface{}
	for _, itemIdItem := range itemId {
		itemIdRule = append(itemIdRule, itemIdItem)
	}
	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _NFTMarket.contract.WatchLogs(opts, "MarketItemSold", itemIdRule, nftContractRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTMarketMarketItemSold)
				if err := _NFTMarket.contract.UnpackLog(event, "MarketItemSold", log); err != nil {
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

// ParseMarketItemSold is a log parse operation binding the contract event 0x6a5f9d2b6ff1f375ebe095556dc91e68fce80a6187d7527cce81c3ab02726c9c.
//
// Solidity: event MarketItemSold(uint256 indexed itemId, address indexed nftContract, uint256 indexed tokenId, address seller, address owner, uint256 price)
func (_NFTMarket *NFTMarketFilterer) ParseMarketItemSold(log types.Log) (*NFTMarketMarketItemSold, error) {
	event := new(NFTMarketMarketItemSold)
	if err := _NFTMarket.contract.UnpackLog(event, "MarketItemSold", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
