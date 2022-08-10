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
	Bin: "0x608060405234801561001057600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550611a24806100606000396000f3fe60806040526004361061004a5760003560e01c806357ea89b61461004f5780636861542f14610059578063b4a99a4e14610096578063ed21248c146100c1578063fa84fd8e146100cb575b600080fd5b6100576100f4565b005b34801561006557600080fd5b50610080600480360381019061007b9190610e75565b6102a7565b60405161008d9190610ebb565b60405180910390f35b3480156100a257600080fd5b506100ab610349565b6040516100b89190610ee5565b60405180910390f35b6100c961036d565b005b3480156100d757600080fd5b506100f260048036038101906100ed9190610f91565b61046c565b005b60003390506000600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020540361017b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161017290611076565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166108fc600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549081150290604051600060405180830381858888f19350505050158015610200573d6000803e3d6000fd5b506000600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a61027033610a7c565b6040516020016102809190611136565b60405160208183030381529060405260405161029c91906111a6565b60405180910390a150565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461030257600080fd5b600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b34600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546103bc91906111f7565b925050819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a6103ed33610a7c565b610435600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610c3f565b604051602001610446929190611299565b60405160208183030381529060405260405161046291906111a6565b60405180910390a1565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146104c457600080fd5b600080600090505b858590508110156107975783600160008888858181106104ef576104ee6112db565b5b90506020020160208101906105049190610e75565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205410156106f8577fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a6105d5600160008989868181106105805761057f6112db565b5b90506020020160208101906105959190610e75565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610c3f565b6105de86610c3f565b6040516020016105ef929190611356565b60405160208183030381529060405260405161060b91906111a6565b60405180910390a16001600087878481811061062a576106296112db565b5b905060200201602081019061063f9190610e75565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548261068591906111f7565b91506000600160008888858181106106a05761069f6112db565b5b90506020020160208101906106b59190610e75565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550610784565b838261070491906111f7565b9150836001600088888581811061071e5761071d6112db565b5b90506020020160208101906107339190610e75565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461077c9190611398565b925050819055505b808061078f906113cc565b9150506104cc565b507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a6107c284610c3f565b6107cb84610c3f565b6107d484610c3f565b6040516020016107e6939291906114ac565b60405160208183030381529060405260405161080291906111a6565b60405180910390a16000810361084d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108449061158b565b60405180910390fd5b818110156109305780600160008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546108c591906111f7565b925050819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a6108f682610c3f565b60405160200161090691906115d1565b60405160208183030381529060405260405161092291906111a6565b60405180910390a150610a75565b818161093c9190611398565b905080600160008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461098d91906111f7565b9250508190555081600160008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610a0491906111f7565b925050819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a610a3582610c3f565b610a3e84610c3f565b604051602001610a4f929190611652565b604051602081830303815290604052604051610a6b91906111a6565b60405180910390a1505b5050505050565b60606000602867ffffffffffffffff811115610a9b57610a9a6116a3565b5b6040519080825280601f01601f191660200182016040528015610acd5781602001600182028036833780820191505090505b50905060005b6014811015610c35576000816013610aeb9190611398565b6008610af791906116d2565b6002610b03919061185f565b8573ffffffffffffffffffffffffffffffffffffffff16610b2491906118d9565b60f81b9050600060108260f81c610b3b9190611917565b60f81b905060008160f81c6010610b529190611948565b8360f81c610b609190611983565b60f81b9050610b6e82610dc7565b85856002610b7c91906116d2565b81518110610b8d57610b8c6112db565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350610bc581610dc7565b856001866002610bd591906116d2565b610bdf91906111f7565b81518110610bf057610bef6112db565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053505050508080610c2d906113cc565b915050610ad3565b5080915050919050565b606060008203610c86576040518060400160405280600181526020017f30000000000000000000000000000000000000000000000000000000000000008152509050610dc2565b600082905060005b60008214610cb8578080610ca1906113cc565b915050600a82610cb191906118d9565b9150610c8e565b60008167ffffffffffffffff811115610cd457610cd36116a3565b5b6040519080825280601f01601f191660200182016040528015610d065781602001600182028036833780820191505090505b50905060008290505b60008614610dba57600181610d249190611398565b90506000600a8088610d3691906118d9565b610d4091906116d2565b87610d4b9190611398565b6030610d5791906119b7565b905060008160f81b905080848481518110610d7557610d746112db565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600a88610db191906118d9565b97505050610d0f565b819450505050505b919050565b6000600a8260f81c60ff161015610df25760308260f81c610de891906119b7565b60f81b9050610e08565b60578260f81c610e0291906119b7565b60f81b90505b919050565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610e4282610e17565b9050919050565b610e5281610e37565b8114610e5d57600080fd5b50565b600081359050610e6f81610e49565b92915050565b600060208284031215610e8b57610e8a610e0d565b5b6000610e9984828501610e60565b91505092915050565b6000819050919050565b610eb581610ea2565b82525050565b6000602082019050610ed06000830184610eac565b92915050565b610edf81610e37565b82525050565b6000602082019050610efa6000830184610ed6565b92915050565b600080fd5b600080fd5b600080fd5b60008083601f840112610f2557610f24610f00565b5b8235905067ffffffffffffffff811115610f4257610f41610f05565b5b602083019150836020820283011115610f5e57610f5d610f0a565b5b9250929050565b610f6e81610ea2565b8114610f7957600080fd5b50565b600081359050610f8b81610f65565b92915050565b600080600080600060808688031215610fad57610fac610e0d565b5b6000610fbb88828901610e60565b955050602086013567ffffffffffffffff811115610fdc57610fdb610e12565b5b610fe888828901610f0f565b94509450506040610ffb88828901610f7c565b925050606061100c88828901610f7c565b9150509295509295909350565b600082825260208201905092915050565b7f6e6f7420656e6f7567682062616c616e63650000000000000000000000000000600082015250565b6000611060601283611019565b915061106b8261102a565b602082019050919050565b6000602082019050818103600083015261108f81611053565b9050919050565b7f77697468647261773a2000000000000000000000000000000000000000000000815250565b600081519050919050565b600081905092915050565b60005b838110156110f05780820151818401526020810190506110d5565b838111156110ff576000848401525b50505050565b6000611110826110bc565b61111a81856110c7565b935061112a8185602086016110d2565b80840191505092915050565b600061114182611096565b600a820191506111518284611105565b915081905092915050565b6000601f19601f8301169050919050565b6000611178826110bc565b6111828185611019565b93506111928185602086016110d2565b61119b8161115c565b840191505092915050565b600060208201905081810360008301526111c0818461116d565b905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061120282610ea2565b915061120d83610ea2565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115611242576112416111c8565b5b828201905092915050565b7f6465706f7369743a200000000000000000000000000000000000000000000000815250565b7f202d200000000000000000000000000000000000000000000000000000000000815250565b60006112a48261124d565b6009820191506112b48285611105565b91506112bf82611273565b6003820191506112cf8284611105565b91508190509392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f706c617965722062616c616e6365200000000000000000000000000000000000815250565b7f206973206c657373207468616e20616e74652000000000000000000000000000815250565b60006113618261130a565b600f820191506113718285611105565b915061137c82611330565b60138201915061138c8284611105565b91508190509392505050565b60006113a382610ea2565b91506113ae83610ea2565b9250828210156113c1576113c06111c8565b5b828203905092915050565b60006113d782610ea2565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611409576114086111c8565b5b600182019050919050565b7f616e74655b000000000000000000000000000000000000000000000000000000815250565b7f5d2067616d654665655b00000000000000000000000000000000000000000000815250565b7f5d20706f745b0000000000000000000000000000000000000000000000000000815250565b7f5d00000000000000000000000000000000000000000000000000000000000000815250565b60006114b782611414565b6005820191506114c78286611105565b91506114d28261143a565b600a820191506114e28285611105565b91506114ed82611460565b6006820191506114fd8284611105565b915061150882611486565b600182019150819050949350505050565b7f6e6f20706f74207761732063726561746564206261736564206f6e20706c617960008201527f65722062616c616e636573000000000000000000000000000000000000000000602082015250565b6000611575602b83611019565b915061158082611519565b604082019050919050565b600060208201905081810360008301526115a481611568565b9050919050565b7f77696e6e696e67506c617965725b305d206f776e65725b000000000000000000815250565b60006115dc826115ab565b6017820191506115ec8284611105565b91506115f782611486565b60018201915081905092915050565b7f77696e6e696e67506c617965725b000000000000000000000000000000000000815250565b7f5d206f776e65725b000000000000000000000000000000000000000000000000815250565b600061165d82611606565b600e8201915061166d8285611105565b91506116788261162c565b6008820191506116888284611105565b915061169382611486565b6001820191508190509392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60006116dd82610ea2565b91506116e883610ea2565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615611721576117206111c8565b5b828202905092915050565b60008160011c9050919050565b6000808291508390505b60018511156117835780860481111561175f5761175e6111c8565b5b600185161561176e5780820291505b808102905061177c8561172c565b9450611743565b94509492505050565b60008261179c5760019050611858565b816117aa5760009050611858565b81600181146117c057600281146117ca576117f9565b6001915050611858565b60ff8411156117dc576117db6111c8565b5b8360020a9150848211156117f3576117f26111c8565b5b50611858565b5060208310610133831016604e8410600b841016171561182e5782820a905083811115611829576118286111c8565b5b611858565b61183b8484846001611739565b92509050818404811115611852576118516111c8565b5b81810290505b9392505050565b600061186a82610ea2565b915061187583610ea2565b92506118a27fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff848461178c565b905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b60006118e482610ea2565b91506118ef83610ea2565b9250826118ff576118fe6118aa565b5b828204905092915050565b600060ff82169050919050565b60006119228261190a565b915061192d8361190a565b92508261193d5761193c6118aa565b5b828204905092915050565b60006119538261190a565b915061195e8361190a565b92508160ff0483118215151615611978576119776111c8565b5b828202905092915050565b600061198e8261190a565b91506119998361190a565b9250828210156119ac576119ab6111c8565b5b828203905092915050565b60006119c28261190a565b91506119cd8361190a565b92508260ff038211156119e3576119e26111c8565b5b82820190509291505056fea2646970667358221220e465e9a316abcbcecc63c8e5bdb7b0cf376a30261051abb3a0ce55aa9ff4aaa264736f6c634300080f0033",
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