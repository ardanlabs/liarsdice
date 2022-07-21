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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"EventLog\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"}],\"name\":\"EventNewGame\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EventPlaceAnte\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"Deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Game\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"created_at\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"finished\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"pot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ante\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GameAnte\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"}],\"name\":\"GameEnd\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NewGame\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PlaceAnte\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550611775806100606000396000f3fe60806040526004361061007b5760003560e01c8063b4a99a4e1161004e578063b4a99a4e14610105578063b676208814610130578063ed21248c14610159578063fcaa632e146101635761007b565b8063097e83b5146100805780633b5bbb24146100975780635b6b431d146100ae578063aef99eef146100d7575b600080fd5b34801561008c57600080fd5b5061009561018e565b005b3480156100a357600080fd5b506100ac6101f9565b005b3480156100ba57600080fd5b506100d560048036038101906100d09190610c85565b61045e565b005b3480156100e357600080fd5b506100ec6105a1565b6040516100fc9493929190610cdc565b60405180910390f35b34801561011157600080fd5b5061011a6105cc565b6040516101279190610d62565b60405180910390f35b34801561013c57600080fd5b5061015760048036038101906101529190610da9565b6105f0565b005b610161610730565b005b34801561016f57600080fd5b506101786107f0565b6040516101859190610dd6565b60405180910390f35b604051806080016040528042815260200160001515815260200160008152602001600581525060016000820151816000015560208201518160010160006101000a81548160ff0219169083151502179055506040820151816002015560608201518160030155905050565b6001800160009054906101000a900460ff161561026a5760405160200161021f90610e17565b6040516020818303038152906040526040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102619190610ec9565b60405180910390fd5b600160030154600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054101561031f576102c46001600301546108b9565b6040516020016102d49190610f99565b6040516020818303038152906040526040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103169190610ec9565b60405180910390fd5b600160030154600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546103739190610fea565b9250508190555060016003015460016002016000828254610394919061101e565b925050819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a6103c533610a41565b6040516020016103d591906110c0565b6040516020818303038152906040526040516103f19190610ec9565b60405180910390a17fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a6104286001600201546108b9565b604051602001610438919061111b565b6040516020818303038152906040526040516104549190610ec9565b60405180910390a1565b80600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205410156104e0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104d79061118d565b60405180910390fd5b80600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461052f9190610fea565b925050819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a61056033610a41565b610569836108b9565b60405160200161057a9291906111f9565b6040516020818303038152906040526040516105969190610ec9565b60405180910390a150565b60018060000154908060010160009054906101000a900460ff16908060020154908060030154905084565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461064857600080fd5b60018060010160006101000a81548160ff021916908315150217905550600160020154600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546106b9919061101e565b925050819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a6106ef6001600201546108b9565b6106f883610a41565b604051602001610709929190611287565b6040516020818303038152906040526040516107259190610ec9565b60405180910390a150565b34600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461077f919061101e565b925050819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a6107b033610a41565b6107b9346108b9565b6040516020016107ca9291906112ef565b6040516020818303038152906040526040516107e69190610ec9565b60405180910390a1565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461084b57600080fd5b7fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a61087a6001600201546108b9565b60405160200161088a9190611357565b6040516020818303038152906040526040516108a69190610ec9565b60405180910390a1600160020154905090565b606060008203610900576040518060400160405280600181526020017f30000000000000000000000000000000000000000000000000000000000000008152509050610a3c565b600082905060005b6000821461093257808061091b9061137d565b915050600a8261092b91906113f4565b9150610908565b60008167ffffffffffffffff81111561094e5761094d611425565b5b6040519080825280601f01601f1916602001820160405280156109805781602001600182028036833780820191505090505b50905060008290505b60008614610a345760018161099e9190610fea565b90506000600a80886109b091906113f4565b6109ba9190611454565b876109c59190610fea565b60306109d191906114bb565b905060008160f81b9050808484815181106109ef576109ee6114f2565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600a88610a2b91906113f4565b97505050610989565b819450505050505b919050565b60606000602867ffffffffffffffff811115610a6057610a5f611425565b5b6040519080825280601f01601f191660200182016040528015610a925781602001600182028036833780820191505090505b50905060005b6014811015610bfa576000816013610ab09190610fea565b6008610abc9190611454565b6002610ac89190611654565b8573ffffffffffffffffffffffffffffffffffffffff16610ae991906113f4565b60f81b9050600060108260f81c610b00919061169f565b60f81b905060008160f81c6010610b1791906116d0565b8360f81c610b25919061170b565b60f81b9050610b3382610c04565b85856002610b419190611454565b81518110610b5257610b516114f2565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350610b8a81610c04565b856001866002610b9a9190611454565b610ba4919061101e565b81518110610bb557610bb46114f2565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053505050508080610bf29061137d565b915050610a98565b5080915050919050565b6000600a8260f81c60ff161015610c2f5760308260f81c610c2591906114bb565b60f81b9050610c45565b60578260f81c610c3f91906114bb565b60f81b90505b919050565b600080fd5b6000819050919050565b610c6281610c4f565b8114610c6d57600080fd5b50565b600081359050610c7f81610c59565b92915050565b600060208284031215610c9b57610c9a610c4a565b5b6000610ca984828501610c70565b91505092915050565b610cbb81610c4f565b82525050565b60008115159050919050565b610cd681610cc1565b82525050565b6000608082019050610cf16000830187610cb2565b610cfe6020830186610ccd565b610d0b6040830185610cb2565b610d186060830184610cb2565b95945050505050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610d4c82610d21565b9050919050565b610d5c81610d41565b82525050565b6000602082019050610d776000830184610d53565b92915050565b610d8681610d41565b8114610d9157600080fd5b50565b600081359050610da381610d7d565b92915050565b600060208284031215610dbf57610dbe610c4a565b5b6000610dcd84828501610d94565b91505092915050565b6000602082019050610deb6000830184610cb2565b92915050565b7f67616d65206973206e6f7420617661696c61626c6520616e796d6f7265000000815250565b6000610e2282610df1565b601d82019150819050919050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610e6a578082015181840152602081019050610e4f565b83811115610e79576000848401525b50505050565b6000601f19601f8301169050919050565b6000610e9b82610e30565b610ea58185610e3b565b9350610eb5818560208601610e4c565b610ebe81610e7f565b840191505092915050565b60006020820190508181036000830152610ee38184610e90565b905092915050565b600081905092915050565b7f6e6f7420656e6f7567682062616c616e636520746f206a6f696e20746865206760008201527f616d652c20697420726571756972657320000000000000000000000000000000602082015250565b6000610f52603183610eeb565b9150610f5d82610ef6565b603182019050919050565b6000610f7382610e30565b610f7d8185610eeb565b9350610f8d818560208601610e4c565b80840191505092915050565b6000610fa482610f45565b9150610fb08284610f68565b915081905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610ff582610c4f565b915061100083610c4f565b92508282101561101357611012610fbb565b5b828203905092915050565b600061102982610c4f565b915061103483610c4f565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0382111561106957611068610fbb565b5b828201905092915050565b7f706c617965722000000000000000000000000000000000000000000000000000815250565b7f206a6f696e6564207468652067616d6500000000000000000000000000000000815250565b60006110cb82611074565b6007820191506110db8284610f68565b91506110e68261109a565b60108201915081905092915050565b7f63757272656e742067616d6520706f7420000000000000000000000000000000815250565b6000611126826110f5565b6011820191506111368284610f68565b915081905092915050565b7f6e6f7420656e6f7567682062616c616e63650000000000000000000000000000600082015250565b6000611177601283610e3b565b915061118282611141565b602082019050919050565b600060208201905081810360008301526111a68161116a565b9050919050565b7f77697468647261773a2000000000000000000000000000000000000000000000815250565b7f202d200000000000000000000000000000000000000000000000000000000000815250565b6000611204826111ad565b600a820191506112148285610f68565b915061121f826111d3565b60038201915061122f8284610f68565b91508190509392505050565b7f67616d65206973206f7665722077697468206120706f74206f66200000000000815250565b7f204c44432e205468652077696e6e657220697320000000000000000000000000815250565b60006112928261123b565b601b820191506112a28285610f68565b91506112ad82611261565b6014820191506112bd8284610f68565b91508190509392505050565b7f6465706f7369743a200000000000000000000000000000000000000000000000815250565b60006112fa826112c9565b60098201915061130a8285610f68565b9150611315826111d3565b6003820191506113258284610f68565b91508190509392505050565b7f67616d652063757272656e7420706f743a200000000000000000000000000000815250565b600061136282611331565b6012820191506113728284610f68565b915081905092915050565b600061138882610c4f565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036113ba576113b9610fbb565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b60006113ff82610c4f565b915061140a83610c4f565b92508261141a576114196113c5565b5b828204905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600061145f82610c4f565b915061146a83610c4f565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04831182151516156114a3576114a2610fbb565b5b828202905092915050565b600060ff82169050919050565b60006114c6826114ae565b91506114d1836114ae565b92508260ff038211156114e7576114e6610fbb565b5b828201905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60008160011c9050919050565b6000808291508390505b60018511156115785780860481111561155457611553610fbb565b5b60018516156115635780820291505b808102905061157185611521565b9450611538565b94509492505050565b600082611591576001905061164d565b8161159f576000905061164d565b81600181146115b557600281146115bf576115ee565b600191505061164d565b60ff8411156115d1576115d0610fbb565b5b8360020a9150848211156115e8576115e7610fbb565b5b5061164d565b5060208310610133831016604e8410600b84101617156116235782820a90508381111561161e5761161d610fbb565b5b61164d565b611630848484600161152e565b9250905081840481111561164757611646610fbb565b5b81810290505b9392505050565b600061165f82610c4f565b915061166a83610c4f565b92506116977fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8484611581565b905092915050565b60006116aa826114ae565b91506116b5836114ae565b9250826116c5576116c46113c5565b5b828204905092915050565b60006116db826114ae565b91506116e6836114ae565b92508160ff0483118215151615611700576116ff610fbb565b5b828202905092915050565b6000611716826114ae565b9150611721836114ae565b92508282101561173457611733610fbb565b5b82820390509291505056fea26469706673582212207eec1658ebc54a6c25cd233285b76ad755f864cc377e2e2bb71d911aa8de1f6564736f6c634300080f0033",
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

// Game is a free data retrieval call binding the contract method 0xaef99eef.
//
// Solidity: function Game() view returns(uint256 created_at, bool finished, uint256 pot, uint256 ante)
func (_Contract *ContractCaller) Game(opts *bind.CallOpts) (struct {
	CreatedAt *big.Int
	Finished  bool
	Pot       *big.Int
	Ante      *big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "Game")

	outstruct := new(struct {
		CreatedAt *big.Int
		Finished  bool
		Pot       *big.Int
		Ante      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CreatedAt = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Finished = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.Pot = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Ante = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Game is a free data retrieval call binding the contract method 0xaef99eef.
//
// Solidity: function Game() view returns(uint256 created_at, bool finished, uint256 pot, uint256 ante)
func (_Contract *ContractSession) Game() (struct {
	CreatedAt *big.Int
	Finished  bool
	Pot       *big.Int
	Ante      *big.Int
}, error) {
	return _Contract.Contract.Game(&_Contract.CallOpts)
}

// Game is a free data retrieval call binding the contract method 0xaef99eef.
//
// Solidity: function Game() view returns(uint256 created_at, bool finished, uint256 pot, uint256 ante)
func (_Contract *ContractCallerSession) Game() (struct {
	CreatedAt *big.Int
	Finished  bool
	Pot       *big.Int
	Ante      *big.Int
}, error) {
	return _Contract.Contract.Game(&_Contract.CallOpts)
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

// GameAnte is a paid mutator transaction binding the contract method 0xfcaa632e.
//
// Solidity: function GameAnte() returns(uint256)
func (_Contract *ContractTransactor) GameAnte(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "GameAnte")
}

// GameAnte is a paid mutator transaction binding the contract method 0xfcaa632e.
//
// Solidity: function GameAnte() returns(uint256)
func (_Contract *ContractSession) GameAnte() (*types.Transaction, error) {
	return _Contract.Contract.GameAnte(&_Contract.TransactOpts)
}

// GameAnte is a paid mutator transaction binding the contract method 0xfcaa632e.
//
// Solidity: function GameAnte() returns(uint256)
func (_Contract *ContractTransactorSession) GameAnte() (*types.Transaction, error) {
	return _Contract.Contract.GameAnte(&_Contract.TransactOpts)
}

// GameEnd is a paid mutator transaction binding the contract method 0xb6762088.
//
// Solidity: function GameEnd(address player) returns()
func (_Contract *ContractTransactor) GameEnd(opts *bind.TransactOpts, player common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "GameEnd", player)
}

// GameEnd is a paid mutator transaction binding the contract method 0xb6762088.
//
// Solidity: function GameEnd(address player) returns()
func (_Contract *ContractSession) GameEnd(player common.Address) (*types.Transaction, error) {
	return _Contract.Contract.GameEnd(&_Contract.TransactOpts, player)
}

// GameEnd is a paid mutator transaction binding the contract method 0xb6762088.
//
// Solidity: function GameEnd(address player) returns()
func (_Contract *ContractTransactorSession) GameEnd(player common.Address) (*types.Transaction, error) {
	return _Contract.Contract.GameEnd(&_Contract.TransactOpts, player)
}

// NewGame is a paid mutator transaction binding the contract method 0x097e83b5.
//
// Solidity: function NewGame() returns()
func (_Contract *ContractTransactor) NewGame(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "NewGame")
}

// NewGame is a paid mutator transaction binding the contract method 0x097e83b5.
//
// Solidity: function NewGame() returns()
func (_Contract *ContractSession) NewGame() (*types.Transaction, error) {
	return _Contract.Contract.NewGame(&_Contract.TransactOpts)
}

// NewGame is a paid mutator transaction binding the contract method 0x097e83b5.
//
// Solidity: function NewGame() returns()
func (_Contract *ContractTransactorSession) NewGame() (*types.Transaction, error) {
	return _Contract.Contract.NewGame(&_Contract.TransactOpts)
}

// PlaceAnte is a paid mutator transaction binding the contract method 0x3b5bbb24.
//
// Solidity: function PlaceAnte() returns()
func (_Contract *ContractTransactor) PlaceAnte(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "PlaceAnte")
}

// PlaceAnte is a paid mutator transaction binding the contract method 0x3b5bbb24.
//
// Solidity: function PlaceAnte() returns()
func (_Contract *ContractSession) PlaceAnte() (*types.Transaction, error) {
	return _Contract.Contract.PlaceAnte(&_Contract.TransactOpts)
}

// PlaceAnte is a paid mutator transaction binding the contract method 0x3b5bbb24.
//
// Solidity: function PlaceAnte() returns()
func (_Contract *ContractTransactorSession) PlaceAnte() (*types.Transaction, error) {
	return _Contract.Contract.PlaceAnte(&_Contract.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x5b6b431d.
//
// Solidity: function Withdraw(uint256 amount) returns()
func (_Contract *ContractTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "Withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x5b6b431d.
//
// Solidity: function Withdraw(uint256 amount) returns()
func (_Contract *ContractSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Withdraw(&_Contract.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x5b6b431d.
//
// Solidity: function Withdraw(uint256 amount) returns()
func (_Contract *ContractTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Withdraw(&_Contract.TransactOpts, amount)
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

// ContractEventPlaceAnteIterator is returned from FilterEventPlaceAnte and is used to iterate over the raw logs and unpacked data for EventPlaceAnte events raised by the Contract contract.
type ContractEventPlaceAnteIterator struct {
	Event *ContractEventPlaceAnte // Event containing the contract specifics and raw log

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
func (it *ContractEventPlaceAnteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractEventPlaceAnte)
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
		it.Event = new(ContractEventPlaceAnte)
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
func (it *ContractEventPlaceAnteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractEventPlaceAnteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractEventPlaceAnte represents a EventPlaceAnte event raised by the Contract contract.
type ContractEventPlaceAnte struct {
	Player common.Address
	Uuid   string
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEventPlaceAnte is a free log retrieval operation binding the contract event 0xeec55407ec32ed157f05d47215b1b78cec719982620d4d02fd3d81085ac72d7d.
//
// Solidity: event EventPlaceAnte(address player, string uuid, uint256 amount)
func (_Contract *ContractFilterer) FilterEventPlaceAnte(opts *bind.FilterOpts) (*ContractEventPlaceAnteIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "EventPlaceAnte")
	if err != nil {
		return nil, err
	}
	return &ContractEventPlaceAnteIterator{contract: _Contract.contract, event: "EventPlaceAnte", logs: logs, sub: sub}, nil
}

// WatchEventPlaceAnte is a free log subscription operation binding the contract event 0xeec55407ec32ed157f05d47215b1b78cec719982620d4d02fd3d81085ac72d7d.
//
// Solidity: event EventPlaceAnte(address player, string uuid, uint256 amount)
func (_Contract *ContractFilterer) WatchEventPlaceAnte(opts *bind.WatchOpts, sink chan<- *ContractEventPlaceAnte) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "EventPlaceAnte")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractEventPlaceAnte)
				if err := _Contract.contract.UnpackLog(event, "EventPlaceAnte", log); err != nil {
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

// ParseEventPlaceAnte is a log parse operation binding the contract event 0xeec55407ec32ed157f05d47215b1b78cec719982620d4d02fd3d81085ac72d7d.
//
// Solidity: event EventPlaceAnte(address player, string uuid, uint256 amount)
func (_Contract *ContractFilterer) ParseEventPlaceAnte(log types.Log) (*ContractEventPlaceAnte, error) {
	event := new(ContractEventPlaceAnte)
	if err := _Contract.contract.UnpackLog(event, "EventPlaceAnte", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
