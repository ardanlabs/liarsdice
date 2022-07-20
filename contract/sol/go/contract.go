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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"EventLog\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"}],\"name\":\"EventNewGame\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EventPlaceAnte\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"}],\"name\":\"GameAnte\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"}],\"name\":\"GameEnd\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"}],\"name\":\"NewGame\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimum\",\"type\":\"uint256\"}],\"name\":\"PlaceAnte\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"games\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"created_at\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"finished\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"pot\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"playerbalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506119ac806100606000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c80636d20931a116100665780636d20931a1461011c578063b4a99a4e1461014e578063e79ba4ee1461016c578063ec7ff2cf1461019c578063f3fef3a3146101b857610093565b8063026853a41461009857806347e7ef24146100b457806367442a3e146100d05780636763b8a414610100575b600080fd5b6100b260048036038101906100ad9190610df4565b6101d4565b005b6100ce60048036038101906100c99190610ec1565b6104a3565b005b6100ea60048036038101906100e59190610f01565b6104fd565b6040516100f79190610f59565b60405180910390f35b61011a60048036038101906101159190610f74565b610602565b005b61013660048036038101906101319190610f01565b610742565b60405161014593929190610feb565b60405180910390f35b61015661078f565b6040516101639190611031565b60405180910390f35b6101866004803603810190610181919061104c565b6107b3565b6040516101939190610f59565b60405180910390f35b6101b660048036038101906101b19190610f01565b6107cb565b005b6101d260048036038101906101cd9190610ec1565b610879565b005b60003390506002846040516101e991906110f3565b908152602001604051809103902060010160009054906101000a900460ff1615610269578360405160200161021e9190611156565b6040516020818303038152906040526040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161026091906111d5565b60405180910390fd5b81600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205410156102eb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102e290611269565b60405180910390fd5b82600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461033a91906112b8565b925050819055508260028560405161035291906110f3565b9081526020016040518091039020600201600082825461037291906112ec565b925050819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a6103a3826108d3565b6103ac85610a96565b866040516020016103bf939291906113b4565b6040516020818303038152906040526040516103db91906111d5565b60405180910390a17fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a61042e60028660405161041791906110f3565b908152602001604051809103902060020154610a96565b60405160200161043e9190611438565b60405160208183030381529060405260405161045a91906111d5565b60405180910390a17feec55407ec32ed157f05d47215b1b78cec719982620d4d02fd3d81085ac72d7d8185856040516104959392919061145e565b60405180910390a150505050565b80600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546104f291906112ec565b925050819055505050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461055857600080fd5b7fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a826105a460028560405161058d91906110f3565b908152602001604051809103902060020154610a96565b6040516020016105b59291906114e8565b6040516020818303038152906040526040516105d191906111d5565b60405180910390a16002826040516105e991906110f3565b9081526020016040518091039020600201549050919050565b60028160405161061291906110f3565b908152602001604051809103902060020154600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461067291906112ec565b92505081905550600160028260405161068b91906110f3565b908152602001604051809103902060010160006101000a81548160ff0219169083151502179055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a816106ff6002846040516106e891906110f3565b908152602001604051809103902060020154610a96565b610708856108d3565b60405160200161071a93929190611585565b60405160208183030381529060405260405161073691906111d5565b60405180910390a15050565b6002818051602081018201805184825260208301602085012081835280955050505050506000915090508060000154908060010160009054906101000a900460ff16908060020154905083565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60016020528060005260406000206000915090505481565b604051806060016040528042815260200160001515815260200160008152506002826040516107fa91906110f3565b90815260200160405180910390206000820151816000015560208201518160010160006101000a81548160ff021916908315150217905550604082015181600201559050507f74cddd52555f6c3d8aa7d988b2923a84baae675f065350a4924e3ed7407eb8178160405161086e91906111d5565b60405180910390a150565b80600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546108c891906112b8565b925050819055505050565b60606000602867ffffffffffffffff8111156108f2576108f1610c93565b5b6040519080825280601f01601f1916602001820160405280156109245781602001600182028036833780820191505090505b50905060005b6014811015610a8c57600081601361094291906112b8565b600861094e91906115e3565b600261095a9190611770565b8573ffffffffffffffffffffffffffffffffffffffff1661097b91906117ea565b60f81b9050600060108260f81c6109929190611828565b60f81b905060008160f81c60106109a99190611859565b8360f81c6109b79190611894565b60f81b90506109c582610c1e565b858560026109d391906115e3565b815181106109e4576109e36118c8565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350610a1c81610c1e565b856001866002610a2c91906115e3565b610a3691906112ec565b81518110610a4757610a466118c8565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053505050508080610a84906118f7565b91505061092a565b5080915050919050565b606060008203610add576040518060400160405280600181526020017f30000000000000000000000000000000000000000000000000000000000000008152509050610c19565b600082905060005b60008214610b0f578080610af8906118f7565b915050600a82610b0891906117ea565b9150610ae5565b60008167ffffffffffffffff811115610b2b57610b2a610c93565b5b6040519080825280601f01601f191660200182016040528015610b5d5781602001600182028036833780820191505090505b50905060008290505b60008614610c1157600181610b7b91906112b8565b90506000600a8088610b8d91906117ea565b610b9791906115e3565b87610ba291906112b8565b6030610bae919061193f565b905060008160f81b905080848481518110610bcc57610bcb6118c8565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600a88610c0891906117ea565b97505050610b66565b819450505050505b919050565b6000600a8260f81c60ff161015610c495760308260f81c610c3f919061193f565b60f81b9050610c5f565b60578260f81c610c59919061193f565b60f81b90505b919050565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610ccb82610c82565b810181811067ffffffffffffffff82111715610cea57610ce9610c93565b5b80604052505050565b6000610cfd610c64565b9050610d098282610cc2565b919050565b600067ffffffffffffffff821115610d2957610d28610c93565b5b610d3282610c82565b9050602081019050919050565b82818337600083830152505050565b6000610d61610d5c84610d0e565b610cf3565b905082815260208101848484011115610d7d57610d7c610c7d565b5b610d88848285610d3f565b509392505050565b600082601f830112610da557610da4610c78565b5b8135610db5848260208601610d4e565b91505092915050565b6000819050919050565b610dd181610dbe565b8114610ddc57600080fd5b50565b600081359050610dee81610dc8565b92915050565b600080600060608486031215610e0d57610e0c610c6e565b5b600084013567ffffffffffffffff811115610e2b57610e2a610c73565b5b610e3786828701610d90565b9350506020610e4886828701610ddf565b9250506040610e5986828701610ddf565b9150509250925092565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610e8e82610e63565b9050919050565b610e9e81610e83565b8114610ea957600080fd5b50565b600081359050610ebb81610e95565b92915050565b60008060408385031215610ed857610ed7610c6e565b5b6000610ee685828601610eac565b9250506020610ef785828601610ddf565b9150509250929050565b600060208284031215610f1757610f16610c6e565b5b600082013567ffffffffffffffff811115610f3557610f34610c73565b5b610f4184828501610d90565b91505092915050565b610f5381610dbe565b82525050565b6000602082019050610f6e6000830184610f4a565b92915050565b60008060408385031215610f8b57610f8a610c6e565b5b6000610f9985828601610eac565b925050602083013567ffffffffffffffff811115610fba57610fb9610c73565b5b610fc685828601610d90565b9150509250929050565b60008115159050919050565b610fe581610fd0565b82525050565b60006060820190506110006000830186610f4a565b61100d6020830185610fdc565b61101a6040830184610f4a565b949350505050565b61102b81610e83565b82525050565b60006020820190506110466000830184611022565b92915050565b60006020828403121561106257611061610c6e565b5b600061107084828501610eac565b91505092915050565b600081519050919050565b600081905092915050565b60005b838110156110ad578082015181840152602081019050611092565b838111156110bc576000848401525b50505050565b60006110cd82611079565b6110d78185611084565b93506110e781856020860161108f565b80840191505092915050565b60006110ff82846110c2565b915081905092915050565b7f67616d6520000000000000000000000000000000000000000000000000000000815250565b7f206973206e6f7420617661696c61626c6520616e796d6f726500000000000000815250565b60006111618261110a565b60058201915061117182846110c2565b915061117c82611130565b60198201915081905092915050565b600082825260208201905092915050565b60006111a782611079565b6111b1818561118b565b93506111c181856020860161108f565b6111ca81610c82565b840191505092915050565b600060208201905081810360008301526111ef818461119c565b905092915050565b7f6e6f7420656e6f7567682062616c616e636520746f20706c616365206120626560008201527f7400000000000000000000000000000000000000000000000000000000000000602082015250565b600061125360218361118b565b915061125e826111f7565b604082019050919050565b6000602082019050818103600083015261128281611246565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006112c382610dbe565b91506112ce83610dbe565b9250828210156112e1576112e0611289565b5b828203905092915050565b60006112f782610dbe565b915061130283610dbe565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0382111561133757611336611289565b5b828201905092915050565b7f706c617965722000000000000000000000000000000000000000000000000000815250565b7f20706c61636564206120626574206f6620000000000000000000000000000000815250565b7f204c4443206f6e2067616d652000000000000000000000000000000000000000815250565b60006113bf82611342565b6007820191506113cf82866110c2565b91506113da82611368565b6011820191506113ea82856110c2565b91506113f58261138e565b600d8201915061140582846110c2565b9150819050949350505050565b7f63757272656e742067616d6520706f7420000000000000000000000000000000815250565b600061144382611412565b60118201915061145382846110c2565b915081905092915050565b60006060820190506114736000830186611022565b8181036020830152611485818561119c565b90506114946040830184610f4a565b949350505050565b7f20686173206120706f74206f6620000000000000000000000000000000000000815250565b7f204c444373000000000000000000000000000000000000000000000000000000815250565b60006114f38261110a565b60058201915061150382856110c2565b915061150e8261149c565b600e8201915061151e82846110c2565b9150611529826114c2565b6005820191508190509392505050565b7f206973206f7665722077697468206120706f74206f6620000000000000000000815250565b7f204c44432e205468652077696e6e657220697320000000000000000000000000815250565b60006115908261110a565b6005820191506115a082866110c2565b91506115ab82611539565b6017820191506115bb82856110c2565b91506115c68261155f565b6014820191506115d682846110c2565b9150819050949350505050565b60006115ee82610dbe565b91506115f983610dbe565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161561163257611631611289565b5b828202905092915050565b60008160011c9050919050565b6000808291508390505b6001851115611694578086048111156116705761166f611289565b5b600185161561167f5780820291505b808102905061168d8561163d565b9450611654565b94509492505050565b6000826116ad5760019050611769565b816116bb5760009050611769565b81600181146116d157600281146116db5761170a565b6001915050611769565b60ff8411156116ed576116ec611289565b5b8360020a91508482111561170457611703611289565b5b50611769565b5060208310610133831016604e8410600b841016171561173f5782820a90508381111561173a57611739611289565b5b611769565b61174c848484600161164a565b9250905081840481111561176357611762611289565b5b81810290505b9392505050565b600061177b82610dbe565b915061178683610dbe565b92506117b37fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff848461169d565b905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b60006117f582610dbe565b915061180083610dbe565b9250826118105761180f6117bb565b5b828204905092915050565b600060ff82169050919050565b60006118338261181b565b915061183e8361181b565b92508261184e5761184d6117bb565b5b828204905092915050565b60006118648261181b565b915061186f8361181b565b92508160ff048311821515161561188957611888611289565b5b828202905092915050565b600061189f8261181b565b91506118aa8361181b565b9250828210156118bd576118bc611289565b5b828203905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600061190282610dbe565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361193457611933611289565b5b600182019050919050565b600061194a8261181b565b91506119558361181b565b92508260ff0382111561196b5761196a611289565b5b82820190509291505056fea2646970667358221220b18b714b55e14dda952d8516926643d0c1c55e65a7b2913be28406f5cdaa4fc964736f6c634300080f0033",
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
