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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"EventLog\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"}],\"name\":\"EventNewGame\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EventPlaceAnte\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"Deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Game\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"created_at\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"finished\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"pot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ante\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"}],\"name\":\"GameEnd\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GamePot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NewGame\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PlaceAnte\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PlayerBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Withdraw\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550611820806100606000396000f3fe6080604052600436106100865760003560e01c80638c354785116100595780638c354785146100ee578063aef99eef14610119578063b4a99a4e14610147578063b676208814610172578063ed21248c1461019b57610086565b8063097e83b51461008b5780633b5bbb24146100a25780633c0fbd85146100b957806357ea89b6146100e4575b600080fd5b34801561009757600080fd5b506100a06101a5565b005b3480156100ae57600080fd5b506100b7610210565b005b3480156100c557600080fd5b506100ce61049e565b6040516100db9190610d36565b60405180910390f35b6100ec610504565b005b3480156100fa57600080fd5b506101036106b7565b6040516101109190610d36565b60405180910390f35b34801561012557600080fd5b5061012e6106fe565b60405161013e9493929190610d6c565b60405180910390f35b34801561015357600080fd5b5061015c610729565b6040516101699190610df2565b60405180910390f35b34801561017e57600080fd5b5061019960048036038101906101949190610e3e565b61074d565b005b6101a361088d565b005b604051806080016040528042815260200160001515815260200160008152602001600581525060016000820151816000015560208201518160010160006101000a81548160ff0219169083151502179055506040820151816002015560608201518160030155905050565b6001800160009054906101000a900460ff1615610262576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161025990610ec8565b60405180910390fd5b6000600160030154036102aa576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102a190610f34565b60405180910390fd5b600160030154600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054101561035f5761030460016003015461098c565b6040516020016103149190611040565b6040516020818303038152906040526040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161035691906110ac565b60405180910390fd5b600160030154600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546103b391906110fd565b92505081905550600160030154600160020160008282546103d49190611131565b925050819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a61040533610b14565b60405160200161041591906111d3565b60405160208183030381529060405260405161043191906110ac565b60405180910390a17fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a61046860016002015461098c565b604051602001610478919061122e565b60405160208183030381529060405260405161049491906110ac565b60405180910390a1565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146104f957600080fd5b600160020154905090565b60003390506000600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541161058b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610582906112a0565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166108fc600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549081150290604051600060405180830381858888f19350505050158015610610573d6000803e3d6000fd5b506000600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a61068033610b14565b60405160200161069091906112e6565b6040516020818303038152906040526040516106ac91906110ac565b60405180910390a150565b6000600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905090565b60018060000154908060010160009054906101000a900460ff16908060020154908060030154905084565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146107a557600080fd5b60018060010160006101000a81548160ff021916908315150217905550600160020154600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546108169190611131565b925050819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a61084c60016002015461098c565b61085583610b14565b604051602001610866929190611358565b60405160208183030381529060405260405161088291906110ac565b60405180910390a150565b34600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546108dc9190611131565b925050819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a61090d33610b14565b610955600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461098c565b6040516020016109669291906113e6565b60405160208183030381529060405260405161098291906110ac565b60405180910390a1565b6060600082036109d3576040518060400160405280600181526020017f30000000000000000000000000000000000000000000000000000000000000008152509050610b0f565b600082905060005b60008214610a055780806109ee90611428565b915050600a826109fe919061149f565b91506109db565b60008167ffffffffffffffff811115610a2157610a206114d0565b5b6040519080825280601f01601f191660200182016040528015610a535781602001600182028036833780820191505090505b50905060008290505b60008614610b0757600181610a7191906110fd565b90506000600a8088610a83919061149f565b610a8d91906114ff565b87610a9891906110fd565b6030610aa49190611566565b905060008160f81b905080848481518110610ac257610ac161159d565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600a88610afe919061149f565b97505050610a5c565b819450505050505b919050565b60606000602867ffffffffffffffff811115610b3357610b326114d0565b5b6040519080825280601f01601f191660200182016040528015610b655781602001600182028036833780820191505090505b50905060005b6014811015610ccd576000816013610b8391906110fd565b6008610b8f91906114ff565b6002610b9b91906116ff565b8573ffffffffffffffffffffffffffffffffffffffff16610bbc919061149f565b60f81b9050600060108260f81c610bd3919061174a565b60f81b905060008160f81c6010610bea919061177b565b8360f81c610bf891906117b6565b60f81b9050610c0682610cd7565b85856002610c1491906114ff565b81518110610c2557610c2461159d565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350610c5d81610cd7565b856001866002610c6d91906114ff565b610c779190611131565b81518110610c8857610c8761159d565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053505050508080610cc590611428565b915050610b6b565b5080915050919050565b6000600a8260f81c60ff161015610d025760308260f81c610cf89190611566565b60f81b9050610d18565b60578260f81c610d129190611566565b60f81b90505b919050565b6000819050919050565b610d3081610d1d565b82525050565b6000602082019050610d4b6000830184610d27565b92915050565b60008115159050919050565b610d6681610d51565b82525050565b6000608082019050610d816000830187610d27565b610d8e6020830186610d5d565b610d9b6040830185610d27565b610da86060830184610d27565b95945050505050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610ddc82610db1565b9050919050565b610dec81610dd1565b82525050565b6000602082019050610e076000830184610de3565b92915050565b600080fd5b610e1b81610dd1565b8114610e2657600080fd5b50565b600081359050610e3881610e12565b92915050565b600060208284031215610e5457610e53610e0d565b5b6000610e6284828501610e29565b91505092915050565b600082825260208201905092915050565b7f67616d65206973206e6f7420617661696c61626c6520616e796d6f7265000000600082015250565b6000610eb2601d83610e6b565b9150610ebd82610e7c565b602082019050919050565b60006020820190508181036000830152610ee181610ea5565b9050919050565b7f67616d65206973206e6f74206372656174656400000000000000000000000000600082015250565b6000610f1e601383610e6b565b9150610f2982610ee8565b602082019050919050565b60006020820190508181036000830152610f4d81610f11565b9050919050565b600081905092915050565b7f6e6f7420656e6f7567682062616c616e636520746f206a6f696e20746865206760008201527f616d652c20697420726571756972657320000000000000000000000000000000602082015250565b6000610fbb603183610f54565b9150610fc682610f5f565b603182019050919050565b600081519050919050565b60005b83811015610ffa578082015181840152602081019050610fdf565b83811115611009576000848401525b50505050565b600061101a82610fd1565b6110248185610f54565b9350611034818560208601610fdc565b80840191505092915050565b600061104b82610fae565b9150611057828461100f565b915081905092915050565b6000601f19601f8301169050919050565b600061107e82610fd1565b6110888185610e6b565b9350611098818560208601610fdc565b6110a181611062565b840191505092915050565b600060208201905081810360008301526110c68184611073565b905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061110882610d1d565b915061111383610d1d565b925082821015611126576111256110ce565b5b828203905092915050565b600061113c82610d1d565b915061114783610d1d565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0382111561117c5761117b6110ce565b5b828201905092915050565b7f706c617965723a20000000000000000000000000000000000000000000000000815250565b7f206a6f696e6564207468652067616d6500000000000000000000000000000000815250565b60006111de82611187565b6008820191506111ee828461100f565b91506111f9826111ad565b60108201915081905092915050565b7f63757272656e742067616d6520706f743a200000000000000000000000000000815250565b600061123982611208565b601282019150611249828461100f565b915081905092915050565b7f6e6f7420656e6f7567682062616c616e63650000000000000000000000000000600082015250565b600061128a601283610e6b565b915061129582611254565b602082019050919050565b600060208201905081810360008301526112b98161127d565b9050919050565b7f77697468647261773a2000000000000000000000000000000000000000000000815250565b60006112f1826112c0565b600a82019150611301828461100f565b915081905092915050565b7f67616d65206973206f7665722077697468206120706f74206f66200000000000815250565b7f204c44432e205468652077696e6e657220697320000000000000000000000000815250565b60006113638261130c565b601b82019150611373828561100f565b915061137e82611332565b60148201915061138e828461100f565b91508190509392505050565b7f6465706f7369743a200000000000000000000000000000000000000000000000815250565b7f202d200000000000000000000000000000000000000000000000000000000000815250565b60006113f18261139a565b600982019150611401828561100f565b915061140c826113c0565b60038201915061141c828461100f565b91508190509392505050565b600061143382610d1d565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611465576114646110ce565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b60006114aa82610d1d565b91506114b583610d1d565b9250826114c5576114c4611470565b5b828204905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600061150a82610d1d565b915061151583610d1d565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161561154e5761154d6110ce565b5b828202905092915050565b600060ff82169050919050565b600061157182611559565b915061157c83611559565b92508260ff03821115611592576115916110ce565b5b828201905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60008160011c9050919050565b6000808291508390505b6001851115611623578086048111156115ff576115fe6110ce565b5b600185161561160e5780820291505b808102905061161c856115cc565b94506115e3565b94509492505050565b60008261163c57600190506116f8565b8161164a57600090506116f8565b8160018114611660576002811461166a57611699565b60019150506116f8565b60ff84111561167c5761167b6110ce565b5b8360020a915084821115611693576116926110ce565b5b506116f8565b5060208310610133831016604e8410600b84101617156116ce5782820a9050838111156116c9576116c86110ce565b5b6116f8565b6116db84848460016115d9565b925090508184048111156116f2576116f16110ce565b5b81810290505b9392505050565b600061170a82610d1d565b915061171583610d1d565b92506117427fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff848461162c565b905092915050565b600061175582611559565b915061176083611559565b9250826117705761176f611470565b5b828204905092915050565b600061178682611559565b915061179183611559565b92508160ff04831182151516156117ab576117aa6110ce565b5b828202905092915050565b60006117c182611559565b91506117cc83611559565b9250828210156117df576117de6110ce565b5b82820390509291505056fea26469706673582212205f6c8babd494f3e7d1e3854c7b88604b02bc2fb6f173fa27697e2bf81e151ca864736f6c634300080f0033",
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

// PlayerBalance is a free data retrieval call binding the contract method 0x8c354785.
//
// Solidity: function PlayerBalance() view returns(uint256)
func (_Contract *ContractCaller) PlayerBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "PlayerBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PlayerBalance is a free data retrieval call binding the contract method 0x8c354785.
//
// Solidity: function PlayerBalance() view returns(uint256)
func (_Contract *ContractSession) PlayerBalance() (*big.Int, error) {
	return _Contract.Contract.PlayerBalance(&_Contract.CallOpts)
}

// PlayerBalance is a free data retrieval call binding the contract method 0x8c354785.
//
// Solidity: function PlayerBalance() view returns(uint256)
func (_Contract *ContractCallerSession) PlayerBalance() (*big.Int, error) {
	return _Contract.Contract.PlayerBalance(&_Contract.CallOpts)
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
