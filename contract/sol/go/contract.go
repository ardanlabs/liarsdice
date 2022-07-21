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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"EventLog\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"}],\"name\":\"EventNewGame\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EventPlaceAnte\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"}],\"name\":\"GameAnte\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"}],\"name\":\"GameEnd\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"}],\"name\":\"NewGame\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"PlaceAnte\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"games\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"created_at\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"finished\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"pot\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"playerbalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506118a4806100606000396000f3fe6080604052600436106100865760003560e01c8063b4a99a4e11610059578063b4a99a4e14610159578063bcb77bd614610184578063e79ba4ee146101a0578063ec7ff2cf146101dd578063f3fef3a31461020657610086565b806347e7ef241461008b57806367442a3e146100b45780636763b8a4146100f15780636d20931a1461011a575b600080fd5b34801561009757600080fd5b506100b260048036038101906100ad9190610c85565b61022f565b005b3480156100c057600080fd5b506100db60048036038101906100d69190610e0b565b610289565b6040516100e89190610e63565b60405180910390f35b3480156100fd57600080fd5b5061011860048036038101906101139190610e7e565b61038e565b005b34801561012657600080fd5b50610141600480360381019061013c9190610e0b565b610457565b60405161015093929190610ef5565b60405180910390f35b34801561016557600080fd5b5061016e6104a4565b60405161017b9190610f3b565b60405180910390f35b61019e60048036038101906101999190610f56565b6104c8565b005b3480156101ac57600080fd5b506101c760048036038101906101c29190610fb2565b61072c565b6040516101d49190610e63565b60405180910390f35b3480156101e957600080fd5b5061020460048036038101906101ff9190610e0b565b610744565b005b34801561021257600080fd5b5061022d60048036038101906102289190610c85565b6107f2565b005b80600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461027e919061100e565b925050819055505050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146102e457600080fd5b7fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a8261033060028560405161031991906110de565b90815260200160405180910390206002015461084c565b604051602001610341929190611167565b60405160208183030381529060405260405161035d9190611202565b60405180910390a160028260405161037591906110de565b9081526020016040518091039020600201549050919050565b60016002826040516103a091906110de565b908152602001604051809103902060010160006101000a81548160ff0219169083151502179055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a816104146002846040516103fd91906110de565b90815260200160405180910390206002015461084c565b61041d856109d4565b60405160200161042f93929190611270565b60405160208183030381529060405260405161044b9190611202565b60405180910390a15050565b6002818051602081018201805184825260208301602085012081835280955050505050506000915090508060000154908060010160009054906101000a900460ff16908060020154905083565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600033905060005a90507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a6104fc3461084c565b6105058361084c565b6040516020016105169291906112ce565b6040516020818303038152906040526040516105329190611202565b60405180910390a160028460405161054a91906110de565b908152602001604051809103902060010160009054906101000a900460ff16156105ca578360405160200161057f9190611318565b6040516020818303038152906040526040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105c19190611202565b60405180910390fd5b826002856040516105db91906110de565b908152602001604051809103902060020160008282546105fb919061100e565b925050819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a61062c836109d4565b6106358561084c565b86604051602001610648939291906113bf565b6040516020818303038152906040526040516106649190611202565b60405180910390a17fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a6106b76002866040516106a091906110de565b90815260200160405180910390206002015461084c565b6040516020016106c79190611443565b6040516020818303038152906040526040516106e39190611202565b60405180910390a17feec55407ec32ed157f05d47215b1b78cec719982620d4d02fd3d81085ac72d7d82858560405161071e93929190611469565b60405180910390a150505050565b60016020528060005260406000206000915090505481565b6040518060600160405280428152602001600015158152602001600081525060028260405161077391906110de565b90815260200160405180910390206000820151816000015560208201518160010160006101000a81548160ff021916908315150217905550604082015181600201559050507f74cddd52555f6c3d8aa7d988b2923a84baae675f065350a4924e3ed7407eb817816040516107e79190611202565b60405180910390a150565b80600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461084191906114a7565b925050819055505050565b606060008203610893576040518060400160405280600181526020017f300000000000000000000000000000000000000000000000000000000000000081525090506109cf565b600082905060005b600082146108c55780806108ae906114db565b915050600a826108be9190611552565b915061089b565b60008167ffffffffffffffff8111156108e1576108e0610ce0565b5b6040519080825280601f01601f1916602001820160405280156109135781602001600182028036833780820191505090505b50905060008290505b600086146109c75760018161093191906114a7565b90506000600a80886109439190611552565b61094d9190611583565b8761095891906114a7565b603061096491906115ea565b905060008160f81b90508084848151811061098257610981611621565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600a886109be9190611552565b9750505061091c565b819450505050505b919050565b60606000602867ffffffffffffffff8111156109f3576109f2610ce0565b5b6040519080825280601f01601f191660200182016040528015610a255781602001600182028036833780820191505090505b50905060005b6014811015610b8d576000816013610a4391906114a7565b6008610a4f9190611583565b6002610a5b9190611783565b8573ffffffffffffffffffffffffffffffffffffffff16610a7c9190611552565b60f81b9050600060108260f81c610a9391906117ce565b60f81b905060008160f81c6010610aaa91906117ff565b8360f81c610ab8919061183a565b60f81b9050610ac682610b97565b85856002610ad49190611583565b81518110610ae557610ae4611621565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350610b1d81610b97565b856001866002610b2d9190611583565b610b37919061100e565b81518110610b4857610b47611621565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053505050508080610b85906114db565b915050610a2b565b5080915050919050565b6000600a8260f81c60ff161015610bc25760308260f81c610bb891906115ea565b60f81b9050610bd8565b60578260f81c610bd291906115ea565b60f81b90505b919050565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610c1c82610bf1565b9050919050565b610c2c81610c11565b8114610c3757600080fd5b50565b600081359050610c4981610c23565b92915050565b6000819050919050565b610c6281610c4f565b8114610c6d57600080fd5b50565b600081359050610c7f81610c59565b92915050565b60008060408385031215610c9c57610c9b610be7565b5b6000610caa85828601610c3a565b9250506020610cbb85828601610c70565b9150509250929050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610d1882610ccf565b810181811067ffffffffffffffff82111715610d3757610d36610ce0565b5b80604052505050565b6000610d4a610bdd565b9050610d568282610d0f565b919050565b600067ffffffffffffffff821115610d7657610d75610ce0565b5b610d7f82610ccf565b9050602081019050919050565b82818337600083830152505050565b6000610dae610da984610d5b565b610d40565b905082815260208101848484011115610dca57610dc9610cca565b5b610dd5848285610d8c565b509392505050565b600082601f830112610df257610df1610cc5565b5b8135610e02848260208601610d9b565b91505092915050565b600060208284031215610e2157610e20610be7565b5b600082013567ffffffffffffffff811115610e3f57610e3e610bec565b5b610e4b84828501610ddd565b91505092915050565b610e5d81610c4f565b82525050565b6000602082019050610e786000830184610e54565b92915050565b60008060408385031215610e9557610e94610be7565b5b6000610ea385828601610c3a565b925050602083013567ffffffffffffffff811115610ec457610ec3610bec565b5b610ed085828601610ddd565b9150509250929050565b60008115159050919050565b610eef81610eda565b82525050565b6000606082019050610f0a6000830186610e54565b610f176020830185610ee6565b610f246040830184610e54565b949350505050565b610f3581610c11565b82525050565b6000602082019050610f506000830184610f2c565b92915050565b60008060408385031215610f6d57610f6c610be7565b5b600083013567ffffffffffffffff811115610f8b57610f8a610bec565b5b610f9785828601610ddd565b9250506020610fa885828601610c70565b9150509250929050565b600060208284031215610fc857610fc7610be7565b5b6000610fd684828501610c3a565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061101982610c4f565b915061102483610c4f565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0382111561105957611058610fdf565b5b828201905092915050565b600081519050919050565b600081905092915050565b60005b8381101561109857808201518184015260208101905061107d565b838111156110a7576000848401525b50505050565b60006110b882611064565b6110c2818561106f565b93506110d281856020860161107a565b80840191505092915050565b60006110ea82846110ad565b915081905092915050565b7f67616d6520000000000000000000000000000000000000000000000000000000815250565b7f20686173206120706f74206f6620000000000000000000000000000000000000815250565b7f204c444373000000000000000000000000000000000000000000000000000000815250565b6000611172826110f5565b60058201915061118282856110ad565b915061118d8261111b565b600e8201915061119d82846110ad565b91506111a882611141565b6005820191508190509392505050565b600082825260208201905092915050565b60006111d482611064565b6111de81856111b8565b93506111ee81856020860161107a565b6111f781610ccf565b840191505092915050565b6000602082019050818103600083015261121c81846111c9565b905092915050565b7f206973206f7665722077697468206120706f74206f6620000000000000000000815250565b7f204c44432e205468652077696e6e657220697320000000000000000000000000815250565b600061127b826110f5565b60058201915061128b82866110ad565b915061129682611224565b6017820191506112a682856110ad565b91506112b18261124a565b6014820191506112c182846110ad565b9150819050949350505050565b60006112da82856110ad565b91506112e682846110ad565b91508190509392505050565b7f206973206e6f7420617661696c61626c6520616e796d6f726500000000000000815250565b6000611323826110f5565b60058201915061133382846110ad565b915061133e826112f2565b60198201915081905092915050565b7f706c617965722000000000000000000000000000000000000000000000000000815250565b7f20706c61636564206120626574206f6620000000000000000000000000000000815250565b7f204c4443206f6e2067616d652000000000000000000000000000000000000000815250565b60006113ca8261134d565b6007820191506113da82866110ad565b91506113e582611373565b6011820191506113f582856110ad565b915061140082611399565b600d8201915061141082846110ad565b9150819050949350505050565b7f63757272656e742067616d6520706f7420000000000000000000000000000000815250565b600061144e8261141d565b60118201915061145e82846110ad565b915081905092915050565b600060608201905061147e6000830186610f2c565b818103602083015261149081856111c9565b905061149f6040830184610e54565b949350505050565b60006114b282610c4f565b91506114bd83610c4f565b9250828210156114d0576114cf610fdf565b5b828203905092915050565b60006114e682610c4f565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361151857611517610fdf565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600061155d82610c4f565b915061156883610c4f565b92508261157857611577611523565b5b828204905092915050565b600061158e82610c4f565b915061159983610c4f565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04831182151516156115d2576115d1610fdf565b5b828202905092915050565b600060ff82169050919050565b60006115f5826115dd565b9150611600836115dd565b92508260ff0382111561161657611615610fdf565b5b828201905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60008160011c9050919050565b6000808291508390505b60018511156116a75780860481111561168357611682610fdf565b5b60018516156116925780820291505b80810290506116a085611650565b9450611667565b94509492505050565b6000826116c0576001905061177c565b816116ce576000905061177c565b81600181146116e457600281146116ee5761171d565b600191505061177c565b60ff841115611700576116ff610fdf565b5b8360020a91508482111561171757611716610fdf565b5b5061177c565b5060208310610133831016604e8410600b84101617156117525782820a90508381111561174d5761174c610fdf565b5b61177c565b61175f848484600161165d565b9250905081840481111561177657611775610fdf565b5b81810290505b9392505050565b600061178e82610c4f565b915061179983610c4f565b92506117c67fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84846116b0565b905092915050565b60006117d9826115dd565b91506117e4836115dd565b9250826117f4576117f3611523565b5b828204905092915050565b600061180a826115dd565b9150611815836115dd565b92508160ff048311821515161561182f5761182e610fdf565b5b828202905092915050565b6000611845826115dd565b9150611850836115dd565b92508282101561186357611862610fdf565b5b82820390509291505056fea2646970667358221220d8542d3be29c93b01ee7c90a83b8300946c9dd2a77528d836fef4ef5ad2831b664736f6c634300080f0033",
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

// GameAnte is a paid mutator transaction binding the contract method 0x67442a3e.
//
// Solidity: function GameAnte(string uuid) returns(uint256)
func (_Contract *ContractTransactor) GameAnte(opts *bind.TransactOpts, uuid string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "GameAnte", uuid)
}

// GameAnte is a paid mutator transaction binding the contract method 0x67442a3e.
//
// Solidity: function GameAnte(string uuid) returns(uint256)
func (_Contract *ContractSession) GameAnte(uuid string) (*types.Transaction, error) {
	return _Contract.Contract.GameAnte(&_Contract.TransactOpts, uuid)
}

// GameAnte is a paid mutator transaction binding the contract method 0x67442a3e.
//
// Solidity: function GameAnte(string uuid) returns(uint256)
func (_Contract *ContractTransactorSession) GameAnte(uuid string) (*types.Transaction, error) {
	return _Contract.Contract.GameAnte(&_Contract.TransactOpts, uuid)
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

// PlaceAnte is a paid mutator transaction binding the contract method 0xbcb77bd6.
//
// Solidity: function PlaceAnte(string uuid, uint256 amount) payable returns()
func (_Contract *ContractTransactor) PlaceAnte(opts *bind.TransactOpts, uuid string, amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "PlaceAnte", uuid, amount)
}

// PlaceAnte is a paid mutator transaction binding the contract method 0xbcb77bd6.
//
// Solidity: function PlaceAnte(string uuid, uint256 amount) payable returns()
func (_Contract *ContractSession) PlaceAnte(uuid string, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.PlaceAnte(&_Contract.TransactOpts, uuid, amount)
}

// PlaceAnte is a paid mutator transaction binding the contract method 0xbcb77bd6.
//
// Solidity: function PlaceAnte(string uuid, uint256 amount) payable returns()
func (_Contract *ContractTransactorSession) PlaceAnte(uuid string, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.PlaceAnte(&_Contract.TransactOpts, uuid, amount)
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
