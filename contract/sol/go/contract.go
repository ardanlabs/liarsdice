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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"EventLog\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"Deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"}],\"name\":\"PlayerBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"winningPlayer\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"losers\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"ante\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gameFee\",\"type\":\"uint256\"}],\"name\":\"Reconcile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Withdraw\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550611564806100606000396000f3fe60806040526004361061004a5760003560e01c806357ea89b61461004f5780636861542f14610059578063b4a99a4e14610096578063ed21248c146100c1578063fa84fd8e146100cb575b600080fd5b6100576100f4565b005b34801561006557600080fd5b50610080600480360381019061007b9190610bd7565b6102a7565b60405161008d9190610c1d565b60405180910390f35b3480156100a257600080fd5b506100ab610349565b6040516100b89190610c47565b60405180910390f35b6100c961036d565b005b3480156100d757600080fd5b506100f260048036038101906100ed9190610de7565b61046c565b005b60003390506000600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020540361017b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161017290610ec7565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166108fc600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549081150290604051600060405180830381858888f19350505050158015610200573d6000803e3d6000fd5b506000600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a610270336107d4565b6040516020016102809190610f87565b60405160208183030381529060405260405161029c9190610fe6565b60405180910390a150565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461030257600080fd5b600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b34600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546103bc9190611037565b925050819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a6103ed336107d4565b610435600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610997565b6040516020016104469291906110d9565b6040516020818303038152906040526040516104629190610fe6565b60405180910390a1565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146104c457600080fd5b600080600090505b84518110156106935783600160008784815181106104ed576104ec61111b565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541015610601576001600086838151811061054d5761054c61111b565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548261059b9190611037565b91506000600160008784815181106105b6576105b561111b565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550610680565b83600160008784815181106106195761061861111b565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461066a919061114a565b92505081905550838261067d9190611037565b91505b808061068b9061117e565b9150506104cc565b5081816106a0919061114a565b905080600160008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546106f19190611037565b9250508190555081600160008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546107689190611037565b925050819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a61079982610997565b6040516020016107a991906111ec565b6040516020818303038152906040526040516107c59190610fe6565b60405180910390a15050505050565b60606000602867ffffffffffffffff8111156107f3576107f2610c78565b5b6040519080825280601f01601f1916602001820160405280156108255781602001600182028036833780820191505090505b50905060005b601481101561098d576000816013610843919061114a565b600861084f9190611212565b600261085b919061139f565b8573ffffffffffffffffffffffffffffffffffffffff1661087c9190611419565b60f81b9050600060108260f81c6108939190611457565b60f81b905060008160f81c60106108aa9190611488565b8360f81c6108b891906114c3565b60f81b90506108c682610b1f565b858560026108d49190611212565b815181106108e5576108e461111b565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535061091d81610b1f565b85600186600261092d9190611212565b6109379190611037565b815181106109485761094761111b565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535050505080806109859061117e565b91505061082b565b5080915050919050565b6060600082036109de576040518060400160405280600181526020017f30000000000000000000000000000000000000000000000000000000000000008152509050610b1a565b600082905060005b60008214610a105780806109f99061117e565b915050600a82610a099190611419565b91506109e6565b60008167ffffffffffffffff811115610a2c57610a2b610c78565b5b6040519080825280601f01601f191660200182016040528015610a5e5781602001600182028036833780820191505090505b50905060008290505b60008614610b1257600181610a7c919061114a565b90506000600a8088610a8e9190611419565b610a989190611212565b87610aa3919061114a565b6030610aaf91906114f7565b905060008160f81b905080848481518110610acd57610acc61111b565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600a88610b099190611419565b97505050610a67565b819450505050505b919050565b6000600a8260f81c60ff161015610b4a5760308260f81c610b4091906114f7565b60f81b9050610b60565b60578260f81c610b5a91906114f7565b60f81b90505b919050565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610ba482610b79565b9050919050565b610bb481610b99565b8114610bbf57600080fd5b50565b600081359050610bd181610bab565b92915050565b600060208284031215610bed57610bec610b6f565b5b6000610bfb84828501610bc2565b91505092915050565b6000819050919050565b610c1781610c04565b82525050565b6000602082019050610c326000830184610c0e565b92915050565b610c4181610b99565b82525050565b6000602082019050610c5c6000830184610c38565b92915050565b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610cb082610c67565b810181811067ffffffffffffffff82111715610ccf57610cce610c78565b5b80604052505050565b6000610ce2610b65565b9050610cee8282610ca7565b919050565b600067ffffffffffffffff821115610d0e57610d0d610c78565b5b602082029050602081019050919050565b600080fd5b6000610d37610d3284610cf3565b610cd8565b90508083825260208201905060208402830185811115610d5a57610d59610d1f565b5b835b81811015610d835780610d6f8882610bc2565b845260208401935050602081019050610d5c565b5050509392505050565b600082601f830112610da257610da1610c62565b5b8135610db2848260208601610d24565b91505092915050565b610dc481610c04565b8114610dcf57600080fd5b50565b600081359050610de181610dbb565b92915050565b60008060008060808587031215610e0157610e00610b6f565b5b6000610e0f87828801610bc2565b945050602085013567ffffffffffffffff811115610e3057610e2f610b74565b5b610e3c87828801610d8d565b9350506040610e4d87828801610dd2565b9250506060610e5e87828801610dd2565b91505092959194509250565b600082825260208201905092915050565b7f6e6f7420656e6f7567682062616c616e63650000000000000000000000000000600082015250565b6000610eb1601283610e6a565b9150610ebc82610e7b565b602082019050919050565b60006020820190508181036000830152610ee081610ea4565b9050919050565b7f77697468647261773a2000000000000000000000000000000000000000000000815250565b600081519050919050565b600081905092915050565b60005b83811015610f41578082015181840152602081019050610f26565b83811115610f50576000848401525b50505050565b6000610f6182610f0d565b610f6b8185610f18565b9350610f7b818560208601610f23565b80840191505092915050565b6000610f9282610ee7565b600a82019150610fa28284610f56565b915081905092915050565b6000610fb882610f0d565b610fc28185610e6a565b9350610fd2818560208601610f23565b610fdb81610c67565b840191505092915050565b600060208201905081810360008301526110008184610fad565b905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061104282610c04565b915061104d83610c04565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0382111561108257611081611008565b5b828201905092915050565b7f6465706f7369743a200000000000000000000000000000000000000000000000815250565b7f202d200000000000000000000000000000000000000000000000000000000000815250565b60006110e48261108d565b6009820191506110f48285610f56565b91506110ff826110b3565b60038201915061110f8284610f56565b91508190509392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600061115582610c04565b915061116083610c04565b92508282101561117357611172611008565b5b828203905092915050565b600061118982610c04565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036111bb576111ba611008565b5b600182019050919050565b7f67616d6520636c6f7365642077697468206120706f74206f6620000000000000815250565b60006111f7826111c6565b601a820191506112078284610f56565b915081905092915050565b600061121d82610c04565b915061122883610c04565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161561126157611260611008565b5b828202905092915050565b60008160011c9050919050565b6000808291508390505b60018511156112c35780860481111561129f5761129e611008565b5b60018516156112ae5780820291505b80810290506112bc8561126c565b9450611283565b94509492505050565b6000826112dc5760019050611398565b816112ea5760009050611398565b8160018114611300576002811461130a57611339565b6001915050611398565b60ff84111561131c5761131b611008565b5b8360020a91508482111561133357611332611008565b5b50611398565b5060208310610133831016604e8410600b841016171561136e5782820a90508381111561136957611368611008565b5b611398565b61137b8484846001611279565b9250905081840481111561139257611391611008565b5b81810290505b9392505050565b60006113aa82610c04565b91506113b583610c04565b92506113e27fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84846112cc565b905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600061142482610c04565b915061142f83610c04565b92508261143f5761143e6113ea565b5b828204905092915050565b600060ff82169050919050565b60006114628261144a565b915061146d8361144a565b92508261147d5761147c6113ea565b5b828204905092915050565b60006114938261144a565b915061149e8361144a565b92508160ff04831182151516156114b8576114b7611008565b5b828202905092915050565b60006114ce8261144a565b91506114d98361144a565b9250828210156114ec576114eb611008565b5b828203905092915050565b60006115028261144a565b915061150d8361144a565b92508260ff0382111561152357611522611008565b5b82820190509291505056fea264697066735822122085f298a28b87fb3b94b4addb747277f810c955ad426b28eb8a47c358272b938064736f6c634300080f0033",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// ContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractMetaData.Bin instead.
var ContractBin = ContractMetaData.Bin

// DeployContract deploys a new Ethereum contract, binding an instance of Contract to it.
func DeployContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Contract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

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

// Owner is a free data retrieval call binding the contract method 0xb4a99a4e.
//
// Solidity: function Owner() view returns(address)
func (_Contract *ContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "Owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0xb4a99a4e.
//
// Solidity: function Owner() view returns(address)
func (_Contract *ContractSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0xb4a99a4e.
//
// Solidity: function Owner() view returns(address)
func (_Contract *ContractCallerSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// PlayerBalance is a free data retrieval call binding the contract method 0x6861542f.
//
// Solidity: function PlayerBalance(address player) view returns(uint256)
func (_Contract *ContractCaller) PlayerBalance(opts *bind.CallOpts, player common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "PlayerBalance", player)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PlayerBalance is a free data retrieval call binding the contract method 0x6861542f.
//
// Solidity: function PlayerBalance(address player) view returns(uint256)
func (_Contract *ContractSession) PlayerBalance(player common.Address) (*big.Int, error) {
	return _Contract.Contract.PlayerBalance(&_Contract.CallOpts, player)
}

// PlayerBalance is a free data retrieval call binding the contract method 0x6861542f.
//
// Solidity: function PlayerBalance(address player) view returns(uint256)
func (_Contract *ContractCallerSession) PlayerBalance(player common.Address) (*big.Int, error) {
	return _Contract.Contract.PlayerBalance(&_Contract.CallOpts, player)
}

// Deposit is a paid mutator transaction binding the contract method 0xed21248c.
//
// Solidity: function Deposit() payable returns()
func (_Contract *ContractTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "Deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xed21248c.
//
// Solidity: function Deposit() payable returns()
func (_Contract *ContractSession) Deposit() (*types.Transaction, error) {
	return _Contract.Contract.Deposit(&_Contract.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xed21248c.
//
// Solidity: function Deposit() payable returns()
func (_Contract *ContractTransactorSession) Deposit() (*types.Transaction, error) {
	return _Contract.Contract.Deposit(&_Contract.TransactOpts)
}

// Reconcile is a paid mutator transaction binding the contract method 0xfa84fd8e.
//
// Solidity: function Reconcile(address winningPlayer, address[] losers, uint256 ante, uint256 gameFee) returns()
func (_Contract *ContractTransactor) Reconcile(opts *bind.TransactOpts, winningPlayer common.Address, losers []common.Address, ante *big.Int, gameFee *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "Reconcile", winningPlayer, losers, ante, gameFee)
}

// Reconcile is a paid mutator transaction binding the contract method 0xfa84fd8e.
//
// Solidity: function Reconcile(address winningPlayer, address[] losers, uint256 ante, uint256 gameFee) returns()
func (_Contract *ContractSession) Reconcile(winningPlayer common.Address, losers []common.Address, ante *big.Int, gameFee *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Reconcile(&_Contract.TransactOpts, winningPlayer, losers, ante, gameFee)
}

// Reconcile is a paid mutator transaction binding the contract method 0xfa84fd8e.
//
// Solidity: function Reconcile(address winningPlayer, address[] losers, uint256 ante, uint256 gameFee) returns()
func (_Contract *ContractTransactorSession) Reconcile(winningPlayer common.Address, losers []common.Address, ante *big.Int, gameFee *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Reconcile(&_Contract.TransactOpts, winningPlayer, losers, ante, gameFee)
}

// Withdraw is a paid mutator transaction binding the contract method 0x57ea89b6.
//
// Solidity: function Withdraw() payable returns()
func (_Contract *ContractTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "Withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x57ea89b6.
//
// Solidity: function Withdraw() payable returns()
func (_Contract *ContractSession) Withdraw() (*types.Transaction, error) {
	return _Contract.Contract.Withdraw(&_Contract.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x57ea89b6.
//
// Solidity: function Withdraw() payable returns()
func (_Contract *ContractTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Contract.Contract.Withdraw(&_Contract.TransactOpts)
}

// ContractEventLogIterator is returned from FilterEventLog and is used to iterate over the raw logs and unpacked data for EventLog events raised by the Contract contract.
type ContractEventLogIterator struct {
	Event *ContractEventLog // Event containing the contract specifics and raw log

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
func (it *ContractEventLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractEventLog)
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
		it.Event = new(ContractEventLog)
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
func (it *ContractEventLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractEventLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractEventLog represents a EventLog event raised by the Contract contract.
type ContractEventLog struct {
	Value string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterEventLog is a free log retrieval operation binding the contract event 0xd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a.
//
// Solidity: event EventLog(string value)
func (_Contract *ContractFilterer) FilterEventLog(opts *bind.FilterOpts) (*ContractEventLogIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "EventLog")
	if err != nil {
		return nil, err
	}
	return &ContractEventLogIterator{contract: _Contract.contract, event: "EventLog", logs: logs, sub: sub}, nil
}

// WatchEventLog is a free log subscription operation binding the contract event 0xd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a.
//
// Solidity: event EventLog(string value)
func (_Contract *ContractFilterer) WatchEventLog(opts *bind.WatchOpts, sink chan<- *ContractEventLog) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "EventLog")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractEventLog)
				if err := _Contract.contract.UnpackLog(event, "EventLog", log); err != nil {
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

// ParseEventLog is a log parse operation binding the contract event 0xd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a.
//
// Solidity: event EventLog(string value)
func (_Contract *ContractFilterer) ParseEventLog(log types.Log) (*ContractEventLog, error) {
	event := new(ContractEventLog)
	if err := _Contract.contract.UnpackLog(event, "EventLog", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
