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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"EventLog\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"}],\"name\":\"EventNewGame\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EventPlaceAnte\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"}],\"name\":\"GameEnd\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"}],\"name\":\"NewGame\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimum\",\"type\":\"uint256\"}],\"name\":\"PlaceAnte\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"games\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"created_at\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"finished\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"pot\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"playerbalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506117cf806100606000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c8063b4a99a4e1161005b578063b4a99a4e14610113578063e79ba4ee14610131578063ec7ff2cf14610161578063f3fef3a31461017d57610088565b8063026853a41461008d57806347e7ef24146100a95780636763b8a4146100c55780636d20931a146100e1575b600080fd5b6100a760048036038101906100a29190610cb4565b610199565b005b6100c360048036038101906100be9190610d81565b610468565b005b6100df60048036038101906100da9190610dc1565b6104c2565b005b6100fb60048036038101906100f69190610e1d565b610602565b60405161010a93929190610e90565b60405180910390f35b61011b61064f565b6040516101289190610ed6565b60405180910390f35b61014b60048036038101906101469190610ef1565b610673565b6040516101589190610f1e565b60405180910390f35b61017b60048036038101906101769190610e1d565b61068b565b005b61019760048036038101906101929190610d81565b610739565b005b60003390506002846040516101ae9190610fb3565b908152602001604051809103902060010160009054906101000a900460ff161561022e57836040516020016101e39190611016565b6040516020818303038152906040526040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102259190611095565b60405180910390fd5b81600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205410156102b0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102a790611129565b60405180910390fd5b82600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546102ff9190611178565b92505081905550826002856040516103179190610fb3565b9081526020016040518091039020600201600082825461033791906111ac565b925050819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a61036882610793565b61037185610956565b8660405160200161038493929190611274565b6040516020818303038152906040526040516103a09190611095565b60405180910390a17fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a6103f36002866040516103dc9190610fb3565b908152602001604051809103902060020154610956565b60405160200161040391906112f8565b60405160208183030381529060405260405161041f9190611095565b60405180910390a17feec55407ec32ed157f05d47215b1b78cec719982620d4d02fd3d81085ac72d7d81858560405161045a9392919061131e565b60405180910390a150505050565b80600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546104b791906111ac565b925050819055505050565b6002816040516104d29190610fb3565b908152602001604051809103902060020154600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461053291906111ac565b92505081905550600160028260405161054b9190610fb3565b908152602001604051809103902060010160006101000a81548160ff0219169083151502179055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a816105bf6002846040516105a89190610fb3565b908152602001604051809103902060020154610956565b6105c885610793565b6040516020016105da939291906113a8565b6040516020818303038152906040526040516105f69190611095565b60405180910390a15050565b6002818051602081018201805184825260208301602085012081835280955050505050506000915090508060000154908060010160009054906101000a900460ff16908060020154905083565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60016020528060005260406000206000915090505481565b604051806060016040528042815260200160001515815260200160008152506002826040516106ba9190610fb3565b90815260200160405180910390206000820151816000015560208201518160010160006101000a81548160ff021916908315150217905550604082015181600201559050507f74cddd52555f6c3d8aa7d988b2923a84baae675f065350a4924e3ed7407eb8178160405161072e9190611095565b60405180910390a150565b80600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546107889190611178565b925050819055505050565b60606000602867ffffffffffffffff8111156107b2576107b1610b53565b5b6040519080825280601f01601f1916602001820160405280156107e45781602001600182028036833780820191505090505b50905060005b601481101561094c5760008160136108029190611178565b600861080e9190611406565b600261081a9190611593565b8573ffffffffffffffffffffffffffffffffffffffff1661083b919061160d565b60f81b9050600060108260f81c610852919061164b565b60f81b905060008160f81c6010610869919061167c565b8360f81c61087791906116b7565b60f81b905061088582610ade565b858560026108939190611406565b815181106108a4576108a36116eb565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053506108dc81610ade565b8560018660026108ec9190611406565b6108f691906111ac565b81518110610907576109066116eb565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535050505080806109449061171a565b9150506107ea565b5080915050919050565b60606000820361099d576040518060400160405280600181526020017f30000000000000000000000000000000000000000000000000000000000000008152509050610ad9565b600082905060005b600082146109cf5780806109b89061171a565b915050600a826109c8919061160d565b91506109a5565b60008167ffffffffffffffff8111156109eb576109ea610b53565b5b6040519080825280601f01601f191660200182016040528015610a1d5781602001600182028036833780820191505090505b50905060008290505b60008614610ad157600181610a3b9190611178565b90506000600a8088610a4d919061160d565b610a579190611406565b87610a629190611178565b6030610a6e9190611762565b905060008160f81b905080848481518110610a8c57610a8b6116eb565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600a88610ac8919061160d565b97505050610a26565b819450505050505b919050565b6000600a8260f81c60ff161015610b095760308260f81c610aff9190611762565b60f81b9050610b1f565b60578260f81c610b199190611762565b60f81b90505b919050565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610b8b82610b42565b810181811067ffffffffffffffff82111715610baa57610ba9610b53565b5b80604052505050565b6000610bbd610b24565b9050610bc98282610b82565b919050565b600067ffffffffffffffff821115610be957610be8610b53565b5b610bf282610b42565b9050602081019050919050565b82818337600083830152505050565b6000610c21610c1c84610bce565b610bb3565b905082815260208101848484011115610c3d57610c3c610b3d565b5b610c48848285610bff565b509392505050565b600082601f830112610c6557610c64610b38565b5b8135610c75848260208601610c0e565b91505092915050565b6000819050919050565b610c9181610c7e565b8114610c9c57600080fd5b50565b600081359050610cae81610c88565b92915050565b600080600060608486031215610ccd57610ccc610b2e565b5b600084013567ffffffffffffffff811115610ceb57610cea610b33565b5b610cf786828701610c50565b9350506020610d0886828701610c9f565b9250506040610d1986828701610c9f565b9150509250925092565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610d4e82610d23565b9050919050565b610d5e81610d43565b8114610d6957600080fd5b50565b600081359050610d7b81610d55565b92915050565b60008060408385031215610d9857610d97610b2e565b5b6000610da685828601610d6c565b9250506020610db785828601610c9f565b9150509250929050565b60008060408385031215610dd857610dd7610b2e565b5b6000610de685828601610d6c565b925050602083013567ffffffffffffffff811115610e0757610e06610b33565b5b610e1385828601610c50565b9150509250929050565b600060208284031215610e3357610e32610b2e565b5b600082013567ffffffffffffffff811115610e5157610e50610b33565b5b610e5d84828501610c50565b91505092915050565b610e6f81610c7e565b82525050565b60008115159050919050565b610e8a81610e75565b82525050565b6000606082019050610ea56000830186610e66565b610eb26020830185610e81565b610ebf6040830184610e66565b949350505050565b610ed081610d43565b82525050565b6000602082019050610eeb6000830184610ec7565b92915050565b600060208284031215610f0757610f06610b2e565b5b6000610f1584828501610d6c565b91505092915050565b6000602082019050610f336000830184610e66565b92915050565b600081519050919050565b600081905092915050565b60005b83811015610f6d578082015181840152602081019050610f52565b83811115610f7c576000848401525b50505050565b6000610f8d82610f39565b610f978185610f44565b9350610fa7818560208601610f4f565b80840191505092915050565b6000610fbf8284610f82565b915081905092915050565b7f67616d6520000000000000000000000000000000000000000000000000000000815250565b7f206973206e6f7420617661696c61626c6520616e796d6f726500000000000000815250565b600061102182610fca565b6005820191506110318284610f82565b915061103c82610ff0565b60198201915081905092915050565b600082825260208201905092915050565b600061106782610f39565b611071818561104b565b9350611081818560208601610f4f565b61108a81610b42565b840191505092915050565b600060208201905081810360008301526110af818461105c565b905092915050565b7f6e6f7420656e6f7567682062616c616e636520746f20706c616365206120626560008201527f7400000000000000000000000000000000000000000000000000000000000000602082015250565b600061111360218361104b565b915061111e826110b7565b604082019050919050565b6000602082019050818103600083015261114281611106565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061118382610c7e565b915061118e83610c7e565b9250828210156111a1576111a0611149565b5b828203905092915050565b60006111b782610c7e565b91506111c283610c7e565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff038211156111f7576111f6611149565b5b828201905092915050565b7f706c617965722000000000000000000000000000000000000000000000000000815250565b7f20706c61636564206120626574206f6620000000000000000000000000000000815250565b7f204c4443206f6e2067616d652000000000000000000000000000000000000000815250565b600061127f82611202565b60078201915061128f8286610f82565b915061129a82611228565b6011820191506112aa8285610f82565b91506112b58261124e565b600d820191506112c58284610f82565b9150819050949350505050565b7f63757272656e742067616d6520706f7420000000000000000000000000000000815250565b6000611303826112d2565b6011820191506113138284610f82565b915081905092915050565b60006060820190506113336000830186610ec7565b8181036020830152611345818561105c565b90506113546040830184610e66565b949350505050565b7f206973206f7665722077697468206120706f74206f6620000000000000000000815250565b7f204c44432e205468652077696e6e657220697320000000000000000000000000815250565b60006113b382610fca565b6005820191506113c38286610f82565b91506113ce8261135c565b6017820191506113de8285610f82565b91506113e982611382565b6014820191506113f98284610f82565b9150819050949350505050565b600061141182610c7e565b915061141c83610c7e565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161561145557611454611149565b5b828202905092915050565b60008160011c9050919050565b6000808291508390505b60018511156114b75780860481111561149357611492611149565b5b60018516156114a25780820291505b80810290506114b085611460565b9450611477565b94509492505050565b6000826114d0576001905061158c565b816114de576000905061158c565b81600181146114f457600281146114fe5761152d565b600191505061158c565b60ff8411156115105761150f611149565b5b8360020a91508482111561152757611526611149565b5b5061158c565b5060208310610133831016604e8410600b84101617156115625782820a90508381111561155d5761155c611149565b5b61158c565b61156f848484600161146d565b9250905081840481111561158657611585611149565b5b81810290505b9392505050565b600061159e82610c7e565b91506115a983610c7e565b92506115d67fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84846114c0565b905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600061161882610c7e565b915061162383610c7e565b925082611633576116326115de565b5b828204905092915050565b600060ff82169050919050565b60006116568261163e565b91506116618361163e565b925082611671576116706115de565b5b828204905092915050565b60006116878261163e565b91506116928361163e565b92508160ff04831182151516156116ac576116ab611149565b5b828202905092915050565b60006116c28261163e565b91506116cd8361163e565b9250828210156116e0576116df611149565b5b828203905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600061172582610c7e565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361175757611756611149565b5b600182019050919050565b600061176d8261163e565b91506117788361163e565b92508260ff0382111561178e5761178d611149565b5b82820190509291505056fea2646970667358221220de8cb49103c43cbc2588b6009d4c2aff89349a7cdfb218ecb61bde2dce8fc11664736f6c634300080f0033",
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

// PlaceAnte is a paid mutator transaction binding the contract method 0x026853a4.
//
// Solidity: function PlaceAnte(string uuid, uint256 amount, uint256 minimum) returns()
func (_Contract *ContractTransactor) PlaceAnte(opts *bind.TransactOpts, uuid string, amount *big.Int, minimum *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "PlaceAnte", uuid, amount, minimum)
}

// PlaceAnte is a paid mutator transaction binding the contract method 0x026853a4.
//
// Solidity: function PlaceAnte(string uuid, uint256 amount, uint256 minimum) returns()
func (_Contract *ContractSession) PlaceAnte(uuid string, amount *big.Int, minimum *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.PlaceAnte(&_Contract.TransactOpts, uuid, amount, minimum)
}

// PlaceAnte is a paid mutator transaction binding the contract method 0x026853a4.
//
// Solidity: function PlaceAnte(string uuid, uint256 amount, uint256 minimum) returns()
func (_Contract *ContractTransactorSession) PlaceAnte(uuid string, amount *big.Int, minimum *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.PlaceAnte(&_Contract.TransactOpts, uuid, amount, minimum)
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
