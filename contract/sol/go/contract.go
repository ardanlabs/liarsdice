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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"EventLog\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"Deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"winningPlayer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gameFee\",\"type\":\"uint256\"}],\"name\":\"GameEnd\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GamePot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"ante\",\"type\":\"uint256\"}],\"name\":\"PlaceAnte\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"}],\"name\":\"PlayerBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Withdraw\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060405180604001604052806001151581526020016000815250600160008201518160000160006101000a81548160ff02191690831515021790555060208201518160010155905050611803806100a86000396000f3fe6080604052600436106100705760003560e01c80636861542f1161004e5780636861542f146100d3578063a01be7ce14610110578063b4a99a4e14610139578063ed21248c1461016457610070565b80632a235412146100755780633c0fbd851461009e57806357ea89b6146100c9575b600080fd5b34801561008157600080fd5b5061009c60048036038101906100979190610df9565b61016e565b005b3480156100aa57600080fd5b506100b3610368565b6040516100c09190610e48565b60405180910390f35b6100d16103cd565b005b3480156100df57600080fd5b506100fa60048036038101906100f59190610e63565b610580565b6040516101079190610e48565b60405180910390f35b34801561011c57600080fd5b5061013760048036038101906101329190610df9565b610622565b005b34801561014557600080fd5b5061014e6108ac565b60405161015b9190610e9f565b60405180910390f35b61016c6108d0565b005b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146101c657600080fd5b6000600160000160006101000a81548160ff021916908315150217905550806001800160008282546101f89190610ee9565b925050819055506001800154600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546102529190610f1d565b9250508190555080600360008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546102c99190610f1d565b925050819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a6102fe60018001546109cf565b61030784610b57565b604051602001610318929190611039565b60405160208183030381529060405260405161033491906110d6565b60405180910390a16000600160000160006101000a81548160ff021916908315150217905550600060018001819055505050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146103c357600080fd5b6001800154905090565b60003390506000600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205403610454576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161044b90611144565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166108fc600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549081150290604051600060405180830381858888f193505050501580156104d9573d6000803e3d6000fd5b506000600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a61054933610b57565b604051602001610559919061118a565b60405160208183030381529060405260405161057591906110d6565b60405180910390a150565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146105db57600080fd5b600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461067a57600080fd5b600160000160009054906101000a900460ff166106cc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106c3906111fc565b60405180910390fd5b80600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205410156107775761071c816109cf565b60405160200161072c919061128e565b6040516020818303038152906040526040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161076e91906110d6565b60405180910390fd5b80600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546107c69190610ee9565b92505081905550806001800160008282546107e19190610f1d565b925050819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a61081233610b57565b60405160200161082291906112fc565b60405160208183030381529060405260405161083e91906110d6565b60405180910390a17fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a61087460018001546109cf565b6040516020016108849190611357565b6040516020818303038152906040526040516108a091906110d6565b60405180910390a15050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b34600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461091f9190610f1d565b925050819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a61095033610b57565b610998600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546109cf565b6040516020016109a99291906113c9565b6040516020818303038152906040526040516109c591906110d6565b60405180910390a1565b606060008203610a16576040518060400160405280600181526020017f30000000000000000000000000000000000000000000000000000000000000008152509050610b52565b600082905060005b60008214610a48578080610a319061140b565b915050600a82610a419190611482565b9150610a1e565b60008167ffffffffffffffff811115610a6457610a636114b3565b5b6040519080825280601f01601f191660200182016040528015610a965781602001600182028036833780820191505090505b50905060008290505b60008614610b4a57600181610ab49190610ee9565b90506000600a8088610ac69190611482565b610ad091906114e2565b87610adb9190610ee9565b6030610ae79190611549565b905060008160f81b905080848481518110610b0557610b04611580565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600a88610b419190611482565b97505050610a9f565b819450505050505b919050565b60606000602867ffffffffffffffff811115610b7657610b756114b3565b5b6040519080825280601f01601f191660200182016040528015610ba85781602001600182028036833780820191505090505b50905060005b6014811015610d10576000816013610bc69190610ee9565b6008610bd291906114e2565b6002610bde91906116e2565b8573ffffffffffffffffffffffffffffffffffffffff16610bff9190611482565b60f81b9050600060108260f81c610c16919061172d565b60f81b905060008160f81c6010610c2d919061175e565b8360f81c610c3b9190611799565b60f81b9050610c4982610d1a565b85856002610c5791906114e2565b81518110610c6857610c67611580565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350610ca081610d1a565b856001866002610cb091906114e2565b610cba9190610f1d565b81518110610ccb57610cca611580565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053505050508080610d089061140b565b915050610bae565b5080915050919050565b6000600a8260f81c60ff161015610d455760308260f81c610d3b9190611549565b60f81b9050610d5b565b60578260f81c610d559190611549565b60f81b90505b919050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610d9082610d65565b9050919050565b610da081610d85565b8114610dab57600080fd5b50565b600081359050610dbd81610d97565b92915050565b6000819050919050565b610dd681610dc3565b8114610de157600080fd5b50565b600081359050610df381610dcd565b92915050565b60008060408385031215610e1057610e0f610d60565b5b6000610e1e85828601610dae565b9250506020610e2f85828601610de4565b9150509250929050565b610e4281610dc3565b82525050565b6000602082019050610e5d6000830184610e39565b92915050565b600060208284031215610e7957610e78610d60565b5b6000610e8784828501610dae565b91505092915050565b610e9981610d85565b82525050565b6000602082019050610eb46000830184610e90565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610ef482610dc3565b9150610eff83610dc3565b925082821015610f1257610f11610eba565b5b828203905092915050565b6000610f2882610dc3565b9150610f3383610dc3565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115610f6857610f67610eba565b5b828201905092915050565b7f67616d65206973206f7665722077697468206120706f74206f66200000000000815250565b600081519050919050565b600081905092915050565b60005b83811015610fcd578082015181840152602081019050610fb2565b83811115610fdc576000848401525b50505050565b6000610fed82610f99565b610ff78185610fa4565b9350611007818560208601610faf565b80840191505092915050565b7f2e205468652077696e6e65722069732000000000000000000000000000000000815250565b600061104482610f73565b601b820191506110548285610fe2565b915061105f82611013565b60108201915061106f8284610fe2565b91508190509392505050565b600082825260208201905092915050565b6000601f19601f8301169050919050565b60006110a882610f99565b6110b2818561107b565b93506110c2818560208601610faf565b6110cb8161108c565b840191505092915050565b600060208201905081810360008301526110f0818461109d565b905092915050565b7f6e6f7420656e6f7567682062616c616e63650000000000000000000000000000600082015250565b600061112e60128361107b565b9150611139826110f8565b602082019050919050565b6000602082019050818103600083015261115d81611121565b9050919050565b7f77697468647261773a2000000000000000000000000000000000000000000000815250565b600061119582611164565b600a820191506111a58284610fe2565b915081905092915050565b7f67616d65206973206e6f7420617661696c61626c6520616e796d6f7265000000600082015250565b60006111e6601d8361107b565b91506111f1826111b0565b602082019050919050565b60006020820190508181036000830152611215816111d9565b9050919050565b7f6e6f7420656e6f7567682062616c616e636520746f206a6f696e20746865206760008201527f616d652c20697420726571756972657320000000000000000000000000000000602082015250565b6000611278603183610fa4565b91506112838261121c565b603182019050919050565b60006112998261126b565b91506112a58284610fe2565b915081905092915050565b7f706c617965723a20000000000000000000000000000000000000000000000000815250565b7f206a6f696e6564207468652067616d6500000000000000000000000000000000815250565b6000611307826112b0565b6008820191506113178284610fe2565b9150611322826112d6565b60108201915081905092915050565b7f63757272656e742067616d6520706f743a200000000000000000000000000000815250565b600061136282611331565b6012820191506113728284610fe2565b915081905092915050565b7f6465706f7369743a200000000000000000000000000000000000000000000000815250565b7f202d200000000000000000000000000000000000000000000000000000000000815250565b60006113d48261137d565b6009820191506113e48285610fe2565b91506113ef826113a3565b6003820191506113ff8284610fe2565b91508190509392505050565b600061141682610dc3565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361144857611447610eba565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600061148d82610dc3565b915061149883610dc3565b9250826114a8576114a7611453565b5b828204905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60006114ed82610dc3565b91506114f883610dc3565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161561153157611530610eba565b5b828202905092915050565b600060ff82169050919050565b60006115548261153c565b915061155f8361153c565b92508260ff0382111561157557611574610eba565b5b828201905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60008160011c9050919050565b6000808291508390505b6001851115611606578086048111156115e2576115e1610eba565b5b60018516156115f15780820291505b80810290506115ff856115af565b94506115c6565b94509492505050565b60008261161f57600190506116db565b8161162d57600090506116db565b8160018114611643576002811461164d5761167c565b60019150506116db565b60ff84111561165f5761165e610eba565b5b8360020a91508482111561167657611675610eba565b5b506116db565b5060208310610133831016604e8410600b84101617156116b15782820a9050838111156116ac576116ab610eba565b5b6116db565b6116be84848460016115bc565b925090508184048111156116d5576116d4610eba565b5b81810290505b9392505050565b60006116ed82610dc3565b91506116f883610dc3565b92506117257fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff848461160f565b905092915050565b60006117388261153c565b91506117438361153c565b92508261175357611752611453565b5b828204905092915050565b60006117698261153c565b91506117748361153c565b92508160ff048311821515161561178e5761178d610eba565b5b828202905092915050565b60006117a48261153c565b91506117af8361153c565b9250828210156117c2576117c1610eba565b5b82820390509291505056fea2646970667358221220b139478d23b6e2eb6fcde94e05b62b076facf8dcae40be641e009632fb42486864736f6c634300080f0033",
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

// GamePot is a free data retrieval call binding the contract method 0x3c0fbd85.
//
// Solidity: function GamePot() view returns(uint256)
func (_Contract *ContractCaller) GamePot(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "GamePot")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GamePot is a free data retrieval call binding the contract method 0x3c0fbd85.
//
// Solidity: function GamePot() view returns(uint256)
func (_Contract *ContractSession) GamePot() (*big.Int, error) {
	return _Contract.Contract.GamePot(&_Contract.CallOpts)
}

// GamePot is a free data retrieval call binding the contract method 0x3c0fbd85.
//
// Solidity: function GamePot() view returns(uint256)
func (_Contract *ContractCallerSession) GamePot() (*big.Int, error) {
	return _Contract.Contract.GamePot(&_Contract.CallOpts)
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

// GameEnd is a paid mutator transaction binding the contract method 0x2a235412.
//
// Solidity: function GameEnd(address winningPlayer, uint256 gameFee) returns()
func (_Contract *ContractTransactor) GameEnd(opts *bind.TransactOpts, winningPlayer common.Address, gameFee *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "GameEnd", winningPlayer, gameFee)
}

// GameEnd is a paid mutator transaction binding the contract method 0x2a235412.
//
// Solidity: function GameEnd(address winningPlayer, uint256 gameFee) returns()
func (_Contract *ContractSession) GameEnd(winningPlayer common.Address, gameFee *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.GameEnd(&_Contract.TransactOpts, winningPlayer, gameFee)
}

// GameEnd is a paid mutator transaction binding the contract method 0x2a235412.
//
// Solidity: function GameEnd(address winningPlayer, uint256 gameFee) returns()
func (_Contract *ContractTransactorSession) GameEnd(winningPlayer common.Address, gameFee *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.GameEnd(&_Contract.TransactOpts, winningPlayer, gameFee)
}

// PlaceAnte is a paid mutator transaction binding the contract method 0xa01be7ce.
//
// Solidity: function PlaceAnte(address player, uint256 ante) returns()
func (_Contract *ContractTransactor) PlaceAnte(opts *bind.TransactOpts, player common.Address, ante *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "PlaceAnte", player, ante)
}

// PlaceAnte is a paid mutator transaction binding the contract method 0xa01be7ce.
//
// Solidity: function PlaceAnte(address player, uint256 ante) returns()
func (_Contract *ContractSession) PlaceAnte(player common.Address, ante *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.PlaceAnte(&_Contract.TransactOpts, player, ante)
}

// PlaceAnte is a paid mutator transaction binding the contract method 0xa01be7ce.
//
// Solidity: function PlaceAnte(address player, uint256 ante) returns()
func (_Contract *ContractTransactorSession) PlaceAnte(player common.Address, ante *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.PlaceAnte(&_Contract.TransactOpts, player, ante)
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
