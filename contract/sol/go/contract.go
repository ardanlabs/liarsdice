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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"EventLog\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"}],\"name\":\"EventNewGame\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EventPlaceBet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"}],\"name\":\"GameEnd\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"}],\"name\":\"NewGame\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimum\",\"type\":\"uint256\"}],\"name\":\"PlaceBet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"games\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"created_at\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"finished\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"pot\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"playerbalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506117de806100606000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c8063b4a99a4e1161005b578063b4a99a4e14610113578063e79ba4ee14610131578063ec7ff2cf14610161578063f3fef3a31461017d57610088565b806319bfdc1c1461008d57806347e7ef24146100a95780636763b8a4146100c55780636d20931a146100e1575b600080fd5b6100a760048036038101906100a29190610d0d565b610199565b005b6100c360048036038101906100be9190610d90565b610463565b005b6100df60048036038101906100da9190610dd0565b6104bd565b005b6100fb60048036038101906100f69190610e2c565b6105fd565b60405161010a93929190610e9f565b60405180910390f35b61011b61064a565b6040516101289190610ee5565b60405180910390f35b61014b60048036038101906101469190610f00565b61066e565b6040516101589190610f2d565b60405180910390f35b61017b60048036038101906101769190610e2c565b610686565b005b61019760048036038101906101929190610d90565b610734565b005b6002836040516101a99190610fc2565b908152602001604051809103902060010160009054906101000a900460ff161561022957826040516020016101de9190611025565b6040516020818303038152906040526040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161022091906110a4565b60405180910390fd5b80600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205410156102ab576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102a290611138565b60405180910390fd5b81600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546102fa9190611187565b92505081905550816002846040516103129190610fc2565b9081526020016040518091039020600201600082825461033291906111bb565b925050819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a6103638561078e565b61036c84610951565b8560405160200161037f93929190611283565b60405160208183030381529060405260405161039b91906110a4565b60405180910390a17fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a6103ee6002856040516103d79190610fc2565b908152602001604051809103902060020154610951565b6040516020016103fe9190611307565b60405160208183030381529060405260405161041a91906110a4565b60405180910390a17f29132e0e10ecd40eafe6b5024f603d9798ca0f6ca09efa70276ef86a0011b41c8484846040516104559392919061132d565b60405180910390a150505050565b80600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546104b291906111bb565b925050819055505050565b6002816040516104cd9190610fc2565b908152602001604051809103902060020154600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461052d91906111bb565b9250508190555060016002826040516105469190610fc2565b908152602001604051809103902060010160006101000a81548160ff0219169083151502179055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a816105ba6002846040516105a39190610fc2565b908152602001604051809103902060020154610951565b6105c38561078e565b6040516020016105d5939291906113b7565b6040516020818303038152906040526040516105f191906110a4565b60405180910390a15050565b6002818051602081018201805184825260208301602085012081835280955050505050506000915090508060000154908060010160009054906101000a900460ff16908060020154905083565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60016020528060005260406000206000915090505481565b604051806060016040528042815260200160001515815260200160008152506002826040516106b59190610fc2565b90815260200160405180910390206000820151816000015560208201518160010160006101000a81548160ff021916908315150217905550604082015181600201559050507f74cddd52555f6c3d8aa7d988b2923a84baae675f065350a4924e3ed7407eb8178160405161072991906110a4565b60405180910390a150565b80600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546107839190611187565b925050819055505050565b60606000602867ffffffffffffffff8111156107ad576107ac610bac565b5b6040519080825280601f01601f1916602001820160405280156107df5781602001600182028036833780820191505090505b50905060005b60148110156109475760008160136107fd9190611187565b60086108099190611415565b600261081591906115a2565b8573ffffffffffffffffffffffffffffffffffffffff16610836919061161c565b60f81b9050600060108260f81c61084d919061165a565b60f81b905060008160f81c6010610864919061168b565b8360f81c61087291906116c6565b60f81b905061088082610ad9565b8585600261088e9190611415565b8151811061089f5761089e6116fa565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053506108d781610ad9565b8560018660026108e79190611415565b6108f191906111bb565b81518110610902576109016116fa565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350505050808061093f90611729565b9150506107e5565b5080915050919050565b606060008203610998576040518060400160405280600181526020017f30000000000000000000000000000000000000000000000000000000000000008152509050610ad4565b600082905060005b600082146109ca5780806109b390611729565b915050600a826109c3919061161c565b91506109a0565b60008167ffffffffffffffff8111156109e6576109e5610bac565b5b6040519080825280601f01601f191660200182016040528015610a185781602001600182028036833780820191505090505b50905060008290505b60008614610acc57600181610a369190611187565b90506000600a8088610a48919061161c565b610a529190611415565b87610a5d9190611187565b6030610a699190611771565b905060008160f81b905080848481518110610a8757610a866116fa565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600a88610ac3919061161c565b97505050610a21565b819450505050505b919050565b6000600a8260f81c60ff161015610b045760308260f81c610afa9190611771565b60f81b9050610b1a565b60578260f81c610b149190611771565b60f81b90505b919050565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610b5e82610b33565b9050919050565b610b6e81610b53565b8114610b7957600080fd5b50565b600081359050610b8b81610b65565b92915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610be482610b9b565b810181811067ffffffffffffffff82111715610c0357610c02610bac565b5b80604052505050565b6000610c16610b1f565b9050610c228282610bdb565b919050565b600067ffffffffffffffff821115610c4257610c41610bac565b5b610c4b82610b9b565b9050602081019050919050565b82818337600083830152505050565b6000610c7a610c7584610c27565b610c0c565b905082815260208101848484011115610c9657610c95610b96565b5b610ca1848285610c58565b509392505050565b600082601f830112610cbe57610cbd610b91565b5b8135610cce848260208601610c67565b91505092915050565b6000819050919050565b610cea81610cd7565b8114610cf557600080fd5b50565b600081359050610d0781610ce1565b92915050565b60008060008060808587031215610d2757610d26610b29565b5b6000610d3587828801610b7c565b945050602085013567ffffffffffffffff811115610d5657610d55610b2e565b5b610d6287828801610ca9565b9350506040610d7387828801610cf8565b9250506060610d8487828801610cf8565b91505092959194509250565b60008060408385031215610da757610da6610b29565b5b6000610db585828601610b7c565b9250506020610dc685828601610cf8565b9150509250929050565b60008060408385031215610de757610de6610b29565b5b6000610df585828601610b7c565b925050602083013567ffffffffffffffff811115610e1657610e15610b2e565b5b610e2285828601610ca9565b9150509250929050565b600060208284031215610e4257610e41610b29565b5b600082013567ffffffffffffffff811115610e6057610e5f610b2e565b5b610e6c84828501610ca9565b91505092915050565b610e7e81610cd7565b82525050565b60008115159050919050565b610e9981610e84565b82525050565b6000606082019050610eb46000830186610e75565b610ec16020830185610e90565b610ece6040830184610e75565b949350505050565b610edf81610b53565b82525050565b6000602082019050610efa6000830184610ed6565b92915050565b600060208284031215610f1657610f15610b29565b5b6000610f2484828501610b7c565b91505092915050565b6000602082019050610f426000830184610e75565b92915050565b600081519050919050565b600081905092915050565b60005b83811015610f7c578082015181840152602081019050610f61565b83811115610f8b576000848401525b50505050565b6000610f9c82610f48565b610fa68185610f53565b9350610fb6818560208601610f5e565b80840191505092915050565b6000610fce8284610f91565b915081905092915050565b7f67616d6520000000000000000000000000000000000000000000000000000000815250565b7f206973206e6f7420617661696c61626c6520616e796d6f726500000000000000815250565b600061103082610fd9565b6005820191506110408284610f91565b915061104b82610fff565b60198201915081905092915050565b600082825260208201905092915050565b600061107682610f48565b611080818561105a565b9350611090818560208601610f5e565b61109981610b9b565b840191505092915050565b600060208201905081810360008301526110be818461106b565b905092915050565b7f6e6f7420656e6f7567682062616c616e636520746f20706c616365206120626560008201527f7400000000000000000000000000000000000000000000000000000000000000602082015250565b600061112260218361105a565b915061112d826110c6565b604082019050919050565b6000602082019050818103600083015261115181611115565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061119282610cd7565b915061119d83610cd7565b9250828210156111b0576111af611158565b5b828203905092915050565b60006111c682610cd7565b91506111d183610cd7565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0382111561120657611205611158565b5b828201905092915050565b7f706c617965722000000000000000000000000000000000000000000000000000815250565b7f20706c61636564206120626574206f6620000000000000000000000000000000815250565b7f204c4443206f6e2067616d652000000000000000000000000000000000000000815250565b600061128e82611211565b60078201915061129e8286610f91565b91506112a982611237565b6011820191506112b98285610f91565b91506112c48261125d565b600d820191506112d48284610f91565b9150819050949350505050565b7f63757272656e742067616d6520706f7420000000000000000000000000000000815250565b6000611312826112e1565b6011820191506113228284610f91565b915081905092915050565b60006060820190506113426000830186610ed6565b8181036020830152611354818561106b565b90506113636040830184610e75565b949350505050565b7f206973206f7665722077697468206120706f74206f6620000000000000000000815250565b7f204c44432e205468652077696e6e657220697320000000000000000000000000815250565b60006113c282610fd9565b6005820191506113d28286610f91565b91506113dd8261136b565b6017820191506113ed8285610f91565b91506113f882611391565b6014820191506114088284610f91565b9150819050949350505050565b600061142082610cd7565b915061142b83610cd7565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161561146457611463611158565b5b828202905092915050565b60008160011c9050919050565b6000808291508390505b60018511156114c6578086048111156114a2576114a1611158565b5b60018516156114b15780820291505b80810290506114bf8561146f565b9450611486565b94509492505050565b6000826114df576001905061159b565b816114ed576000905061159b565b8160018114611503576002811461150d5761153c565b600191505061159b565b60ff84111561151f5761151e611158565b5b8360020a91508482111561153657611535611158565b5b5061159b565b5060208310610133831016604e8410600b84101617156115715782820a90508381111561156c5761156b611158565b5b61159b565b61157e848484600161147c565b9250905081840481111561159557611594611158565b5b81810290505b9392505050565b60006115ad82610cd7565b91506115b883610cd7565b92506115e57fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84846114cf565b905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600061162782610cd7565b915061163283610cd7565b925082611642576116416115ed565b5b828204905092915050565b600060ff82169050919050565b60006116658261164d565b91506116708361164d565b9250826116805761167f6115ed565b5b828204905092915050565b60006116968261164d565b91506116a18361164d565b92508160ff04831182151516156116bb576116ba611158565b5b828202905092915050565b60006116d18261164d565b91506116dc8361164d565b9250828210156116ef576116ee611158565b5b828203905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600061173482610cd7565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361176657611765611158565b5b600182019050919050565b600061177c8261164d565b91506117878361164d565b92508260ff0382111561179d5761179c611158565b5b82820190509291505056fea264697066735822122053625d7d3d28e5c4ffc344bf88c9f0cbd584be7964de7e2a56e7138fe3cba31464736f6c634300080f0033",
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

// Games is a free data retrieval call binding the contract method 0x6d20931a.
//
// Solidity: function games(string ) view returns(uint256 created_at, bool finished, uint256 pot)
func (_Contract *ContractCaller) Games(opts *bind.CallOpts, arg0 string) (struct {
	CreatedAt *big.Int
	Finished  bool
	Pot       *big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "games", arg0)

	outstruct := new(struct {
		CreatedAt *big.Int
		Finished  bool
		Pot       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CreatedAt = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Finished = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.Pot = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Games is a free data retrieval call binding the contract method 0x6d20931a.
//
// Solidity: function games(string ) view returns(uint256 created_at, bool finished, uint256 pot)
func (_Contract *ContractSession) Games(arg0 string) (struct {
	CreatedAt *big.Int
	Finished  bool
	Pot       *big.Int
}, error) {
	return _Contract.Contract.Games(&_Contract.CallOpts, arg0)
}

// Games is a free data retrieval call binding the contract method 0x6d20931a.
//
// Solidity: function games(string ) view returns(uint256 created_at, bool finished, uint256 pot)
func (_Contract *ContractCallerSession) Games(arg0 string) (struct {
	CreatedAt *big.Int
	Finished  bool
	Pot       *big.Int
}, error) {
	return _Contract.Contract.Games(&_Contract.CallOpts, arg0)
}

// Playerbalance is a free data retrieval call binding the contract method 0xe79ba4ee.
//
// Solidity: function playerbalance(address ) view returns(uint256)
func (_Contract *ContractCaller) Playerbalance(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "playerbalance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Playerbalance is a free data retrieval call binding the contract method 0xe79ba4ee.
//
// Solidity: function playerbalance(address ) view returns(uint256)
func (_Contract *ContractSession) Playerbalance(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.Playerbalance(&_Contract.CallOpts, arg0)
}

// Playerbalance is a free data retrieval call binding the contract method 0xe79ba4ee.
//
// Solidity: function playerbalance(address ) view returns(uint256)
func (_Contract *ContractCallerSession) Playerbalance(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.Playerbalance(&_Contract.CallOpts, arg0)
}

// GameEnd is a paid mutator transaction binding the contract method 0x6763b8a4.
//
// Solidity: function GameEnd(address player, string uuid) returns()
func (_Contract *ContractTransactor) GameEnd(opts *bind.TransactOpts, player common.Address, uuid string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "GameEnd", player, uuid)
}

// GameEnd is a paid mutator transaction binding the contract method 0x6763b8a4.
//
// Solidity: function GameEnd(address player, string uuid) returns()
func (_Contract *ContractSession) GameEnd(player common.Address, uuid string) (*types.Transaction, error) {
	return _Contract.Contract.GameEnd(&_Contract.TransactOpts, player, uuid)
}

// GameEnd is a paid mutator transaction binding the contract method 0x6763b8a4.
//
// Solidity: function GameEnd(address player, string uuid) returns()
func (_Contract *ContractTransactorSession) GameEnd(player common.Address, uuid string) (*types.Transaction, error) {
	return _Contract.Contract.GameEnd(&_Contract.TransactOpts, player, uuid)
}

// NewGame is a paid mutator transaction binding the contract method 0xec7ff2cf.
//
// Solidity: function NewGame(string uuid) returns()
func (_Contract *ContractTransactor) NewGame(opts *bind.TransactOpts, uuid string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "NewGame", uuid)
}

// NewGame is a paid mutator transaction binding the contract method 0xec7ff2cf.
//
// Solidity: function NewGame(string uuid) returns()
func (_Contract *ContractSession) NewGame(uuid string) (*types.Transaction, error) {
	return _Contract.Contract.NewGame(&_Contract.TransactOpts, uuid)
}

// NewGame is a paid mutator transaction binding the contract method 0xec7ff2cf.
//
// Solidity: function NewGame(string uuid) returns()
func (_Contract *ContractTransactorSession) NewGame(uuid string) (*types.Transaction, error) {
	return _Contract.Contract.NewGame(&_Contract.TransactOpts, uuid)
}

// PlaceBet is a paid mutator transaction binding the contract method 0x19bfdc1c.
//
// Solidity: function PlaceBet(address player, string uuid, uint256 amount, uint256 minimum) returns()
func (_Contract *ContractTransactor) PlaceBet(opts *bind.TransactOpts, player common.Address, uuid string, amount *big.Int, minimum *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "PlaceBet", player, uuid, amount, minimum)
}

// PlaceBet is a paid mutator transaction binding the contract method 0x19bfdc1c.
//
// Solidity: function PlaceBet(address player, string uuid, uint256 amount, uint256 minimum) returns()
func (_Contract *ContractSession) PlaceBet(player common.Address, uuid string, amount *big.Int, minimum *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.PlaceBet(&_Contract.TransactOpts, player, uuid, amount, minimum)
}

// PlaceBet is a paid mutator transaction binding the contract method 0x19bfdc1c.
//
// Solidity: function PlaceBet(address player, string uuid, uint256 amount, uint256 minimum) returns()
func (_Contract *ContractTransactorSession) PlaceBet(player common.Address, uuid string, amount *big.Int, minimum *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.PlaceBet(&_Contract.TransactOpts, player, uuid, amount, minimum)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address player, uint256 amount) returns()
func (_Contract *ContractTransactor) Deposit(opts *bind.TransactOpts, player common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "deposit", player, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address player, uint256 amount) returns()
func (_Contract *ContractSession) Deposit(player common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Deposit(&_Contract.TransactOpts, player, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address player, uint256 amount) returns()
func (_Contract *ContractTransactorSession) Deposit(player common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Deposit(&_Contract.TransactOpts, player, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address player, uint256 amount) returns()
func (_Contract *ContractTransactor) Withdraw(opts *bind.TransactOpts, player common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "withdraw", player, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address player, uint256 amount) returns()
func (_Contract *ContractSession) Withdraw(player common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Withdraw(&_Contract.TransactOpts, player, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address player, uint256 amount) returns()
func (_Contract *ContractTransactorSession) Withdraw(player common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Withdraw(&_Contract.TransactOpts, player, amount)
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

// ContractEventNewGameIterator is returned from FilterEventNewGame and is used to iterate over the raw logs and unpacked data for EventNewGame events raised by the Contract contract.
type ContractEventNewGameIterator struct {
	Event *ContractEventNewGame // Event containing the contract specifics and raw log

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
func (it *ContractEventNewGameIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractEventNewGame)
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
		it.Event = new(ContractEventNewGame)
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
func (it *ContractEventNewGameIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractEventNewGameIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractEventNewGame represents a EventNewGame event raised by the Contract contract.
type ContractEventNewGame struct {
	Uuid string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterEventNewGame is a free log retrieval operation binding the contract event 0x74cddd52555f6c3d8aa7d988b2923a84baae675f065350a4924e3ed7407eb817.
//
// Solidity: event EventNewGame(string uuid)
func (_Contract *ContractFilterer) FilterEventNewGame(opts *bind.FilterOpts) (*ContractEventNewGameIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "EventNewGame")
	if err != nil {
		return nil, err
	}
	return &ContractEventNewGameIterator{contract: _Contract.contract, event: "EventNewGame", logs: logs, sub: sub}, nil
}

// WatchEventNewGame is a free log subscription operation binding the contract event 0x74cddd52555f6c3d8aa7d988b2923a84baae675f065350a4924e3ed7407eb817.
//
// Solidity: event EventNewGame(string uuid)
func (_Contract *ContractFilterer) WatchEventNewGame(opts *bind.WatchOpts, sink chan<- *ContractEventNewGame) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "EventNewGame")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractEventNewGame)
				if err := _Contract.contract.UnpackLog(event, "EventNewGame", log); err != nil {
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

// ParseEventNewGame is a log parse operation binding the contract event 0x74cddd52555f6c3d8aa7d988b2923a84baae675f065350a4924e3ed7407eb817.
//
// Solidity: event EventNewGame(string uuid)
func (_Contract *ContractFilterer) ParseEventNewGame(log types.Log) (*ContractEventNewGame, error) {
	event := new(ContractEventNewGame)
	if err := _Contract.contract.UnpackLog(event, "EventNewGame", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractEventPlaceBetIterator is returned from FilterEventPlaceBet and is used to iterate over the raw logs and unpacked data for EventPlaceBet events raised by the Contract contract.
type ContractEventPlaceBetIterator struct {
	Event *ContractEventPlaceBet // Event containing the contract specifics and raw log

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
func (it *ContractEventPlaceBetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractEventPlaceBet)
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
		it.Event = new(ContractEventPlaceBet)
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
func (it *ContractEventPlaceBetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractEventPlaceBetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractEventPlaceBet represents a EventPlaceBet event raised by the Contract contract.
type ContractEventPlaceBet struct {
	Player common.Address
	Uuid   string
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEventPlaceBet is a free log retrieval operation binding the contract event 0x29132e0e10ecd40eafe6b5024f603d9798ca0f6ca09efa70276ef86a0011b41c.
//
// Solidity: event EventPlaceBet(address player, string uuid, uint256 amount)
func (_Contract *ContractFilterer) FilterEventPlaceBet(opts *bind.FilterOpts) (*ContractEventPlaceBetIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "EventPlaceBet")
	if err != nil {
		return nil, err
	}
	return &ContractEventPlaceBetIterator{contract: _Contract.contract, event: "EventPlaceBet", logs: logs, sub: sub}, nil
}

// WatchEventPlaceBet is a free log subscription operation binding the contract event 0x29132e0e10ecd40eafe6b5024f603d9798ca0f6ca09efa70276ef86a0011b41c.
//
// Solidity: event EventPlaceBet(address player, string uuid, uint256 amount)
func (_Contract *ContractFilterer) WatchEventPlaceBet(opts *bind.WatchOpts, sink chan<- *ContractEventPlaceBet) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "EventPlaceBet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractEventPlaceBet)
				if err := _Contract.contract.UnpackLog(event, "EventPlaceBet", log); err != nil {
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

// ParseEventPlaceBet is a log parse operation binding the contract event 0x29132e0e10ecd40eafe6b5024f603d9798ca0f6ca09efa70276ef86a0011b41c.
//
// Solidity: event EventPlaceBet(address player, string uuid, uint256 amount)
func (_Contract *ContractFilterer) ParseEventPlaceBet(log types.Log) (*ContractEventPlaceBet, error) {
	event := new(ContractEventPlaceBet)
	if err := _Contract.contract.UnpackLog(event, "EventPlaceBet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
