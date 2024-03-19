// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bank

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
	_ = abi.ConvertType
)

// BankMetaData contains all meta data concerning the Bank contract.
var BankMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"EventLog\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AccountBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Balance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"losers\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"anteWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gameFeeWei\",\"type\":\"uint256\"}],\"name\":\"Reconcile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Withdraw\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506040518060400160405280600581526020017f302e312e300000000000000000000000000000000000000000000000000000008152506001908161009591906102eb565b506103bd565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061011c57607f821691505b60208210810361012f5761012e6100d5565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026101977fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8261015a565b6101a1868361015a565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b60006101e86101e36101de846101b9565b6101c3565b6101b9565b9050919050565b6000819050919050565b610202836101cd565b61021661020e826101ef565b848454610167565b825550505050565b600090565b61022b61021e565b6102368184846101f9565b505050565b5b8181101561025a5761024f600082610223565b60018101905061023c565b5050565b601f82111561029f5761027081610135565b6102798461014a565b81016020851015610288578190505b61029c6102948561014a565b83018261023b565b50505b505050565b600082821c905092915050565b60006102c2600019846008026102a4565b1980831691505092915050565b60006102db83836102b1565b9150826002028217905092915050565b6102f48261009b565b67ffffffffffffffff81111561030d5761030c6100a6565b5b6103178254610104565b61032282828561025e565b600060209050601f8311600181146103555760008415610343578287015190505b61034d85826102cf565b8655506103b5565b601f19841661036386610135565b60005b8281101561038b57848901518255600182019150602085019450602081019050610366565b868310156103a857848901516103a4601f8916826102b1565b8355505b6001600288020188555050505b505050505050565b611c3f806103cc6000396000f3fe6080604052600436106100705760003560e01c8063bb62860d1161004e578063bb62860d146100d5578063e63f341f14610100578063ed21248c1461013d578063fa84fd8e1461014757610070565b80630ef678871461007557806357ea89b6146100a0578063b4a99a4e146100aa575b600080fd5b34801561008157600080fd5b5061008a610170565b6040516100979190610f7b565b60405180910390f35b6100a86101b7565b005b3480156100b657600080fd5b506100bf61037a565b6040516100cc9190610fd7565b60405180910390f35b3480156100e157600080fd5b506100ea61039e565b6040516100f79190611082565b60405180910390f35b34801561010c57600080fd5b50610127600480360381019061012291906110da565b61042c565b6040516101349190610f7b565b60405180910390f35b6101456104ce565b005b34801561015357600080fd5b5061016e60048036038101906101699190611198565b6105cd565b005b6000600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905090565b60003390506000600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020540361023e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102359061126c565b60405180910390fd5b6000600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490508173ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f193505050501580156102c8573d6000803e3d6000fd5b506000600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a61033833610bd7565b61034183610d94565b60405160200161035292919061133a565b60405160208183030381529060405260405161036e9190611082565b60405180910390a15050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600180546103ab906113ba565b80601f01602080910402602001604051908101604052809291908181526020018280546103d7906113ba565b80156104245780601f106103f957610100808354040283529160200191610424565b820191906000526020600020905b81548152906001019060200180831161040757829003601f168201915b505050505081565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461048757600080fd5b600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b34600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461051d919061141a565b925050819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a61054e33610bd7565b610596600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610d94565b6040516020016105a792919061149a565b6040516020818303038152906040526040516105c39190611082565b60405180910390a1565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461062557600080fd5b600082905060005b858590508110156108f25783600260008888858181106106505761064f6114eb565b5b905060200201602081019061066591906110da565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541015610859577fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a610736600260008989868181106106e1576106e06114eb565b5b90506020020160208101906106f691906110da565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610d94565b61073f86610d94565b604051602001610750929190611566565b60405160208183030381529060405260405161076c9190611082565b60405180910390a16002600087878481811061078b5761078a6114eb565b5b90506020020160208101906107a091906110da565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054826107e6919061141a565b9150600060026000888885818110610801576108006114eb565b5b905060200201602081019061081691906110da565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506108e5565b8382610865919061141a565b9150836002600088888581811061087f5761087e6114eb565b5b905060200201602081019061089491906110da565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546108dd91906115a8565b925050819055505b808060010191505061062d565b507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a61091d84610d94565b61092684610d94565b61092f84610d94565b6040516020016109419392919061164e565b60405160208183030381529060405260405161095d9190611082565b60405180910390a1600081036109a8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161099f9061172d565b60405180910390fd5b81811015610a8b5780600260008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610a20919061141a565b925050819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a610a5182610d94565b604051602001610a6191906117bf565b604051602081830303815290604052604051610a7d9190611082565b60405180910390a150610bd0565b8181610a9791906115a8565b905080600260008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610ae8919061141a565b9250508190555081600260008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610b5f919061141a565b925050819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a610b9082610d94565b610b9984610d94565b604051602001610baa92919061183c565b604051602081830303815290604052604051610bc69190611082565b60405180910390a1505b5050505050565b60606000602867ffffffffffffffff811115610bf657610bf561188d565b5b6040519080825280601f01601f191660200182016040528015610c285781602001600182028036833780820191505090505b50905060005b6014811015610d8a576000816013610c4691906115a8565b6008610c5291906118bc565b6002610c5e9190611a31565b8573ffffffffffffffffffffffffffffffffffffffff16610c7f9190611aab565b60f81b9050600060108260f81c610c969190611ae9565b60f81b905060008160f81c6010610cad9190611b1a565b8360f81c610cbb9190611b57565b60f81b9050610cc982610f1c565b85856002610cd791906118bc565b81518110610ce857610ce76114eb565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350610d2081610f1c565b856001866002610d3091906118bc565b610d3a919061141a565b81518110610d4b57610d4a6114eb565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053505050508080600101915050610c2e565b5080915050919050565b606060008203610ddb576040518060400160405280600181526020017f30000000000000000000000000000000000000000000000000000000000000008152509050610f17565b600082905060005b60008214610e0d578080610df690611b8c565b915050600a82610e069190611aab565b9150610de3565b60008167ffffffffffffffff811115610e2957610e2861188d565b5b6040519080825280601f01601f191660200182016040528015610e5b5781602001600182028036833780820191505090505b50905060008290505b60008614610f0f57600181610e7991906115a8565b90506000600a8088610e8b9190611aab565b610e9591906118bc565b87610ea091906115a8565b6030610eac9190611bd4565b905060008160f81b905080848481518110610eca57610ec96114eb565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600a88610f069190611aab565b97505050610e64565b819450505050505b919050565b6000600a8260f81c60ff161015610f475760308260f81c610f3d9190611bd4565b60f81b9050610f5d565b60578260f81c610f579190611bd4565b60f81b90505b919050565b6000819050919050565b610f7581610f62565b82525050565b6000602082019050610f906000830184610f6c565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610fc182610f96565b9050919050565b610fd181610fb6565b82525050565b6000602082019050610fec6000830184610fc8565b92915050565b600081519050919050565b600082825260208201905092915050565b60005b8381101561102c578082015181840152602081019050611011565b60008484015250505050565b6000601f19601f8301169050919050565b600061105482610ff2565b61105e8185610ffd565b935061106e81856020860161100e565b61107781611038565b840191505092915050565b6000602082019050818103600083015261109c8184611049565b905092915050565b600080fd5b600080fd5b6110b781610fb6565b81146110c257600080fd5b50565b6000813590506110d4816110ae565b92915050565b6000602082840312156110f0576110ef6110a4565b5b60006110fe848285016110c5565b91505092915050565b600080fd5b600080fd5b600080fd5b60008083601f84011261112c5761112b611107565b5b8235905067ffffffffffffffff8111156111495761114861110c565b5b60208301915083602082028301111561116557611164611111565b5b9250929050565b61117581610f62565b811461118057600080fd5b50565b6000813590506111928161116c565b92915050565b6000806000806000608086880312156111b4576111b36110a4565b5b60006111c2888289016110c5565b955050602086013567ffffffffffffffff8111156111e3576111e26110a9565b5b6111ef88828901611116565b9450945050604061120288828901611183565b925050606061121388828901611183565b9150509295509295909350565b7f6e6f7420656e6f7567682062616c616e63650000000000000000000000000000600082015250565b6000611256601283610ffd565b915061126182611220565b602082019050919050565b6000602082019050818103600083015261128581611249565b9050919050565b7f77697468647261775b0000000000000000000000000000000000000000000000815250565b600081905092915050565b60006112c882610ff2565b6112d281856112b2565b93506112e281856020860161100e565b80840191505092915050565b7f5d20616d6f756e745b0000000000000000000000000000000000000000000000815250565b7f5d00000000000000000000000000000000000000000000000000000000000000815250565b60006113458261128c565b60098201915061135582856112bd565b9150611360826112ee565b60098201915061137082846112bd565b915061137b82611314565b6001820191508190509392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806113d257607f821691505b6020821081036113e5576113e461138b565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061142582610f62565b915061143083610f62565b9250828201905080821115611448576114476113eb565b5b92915050565b7f6465706f7369745b000000000000000000000000000000000000000000000000815250565b7f5d2062616c616e63655b00000000000000000000000000000000000000000000815250565b60006114a58261144e565b6008820191506114b582856112bd565b91506114c082611474565b600a820191506114d082846112bd565b91506114db82611314565b6001820191508190509392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f6163636f756e742062616c616e63652000000000000000000000000000000000815250565b7f206973206c657373207468616e20616e74652000000000000000000000000000815250565b60006115718261151a565b60108201915061158182856112bd565b915061158c82611540565b60138201915061159c82846112bd565b91508190509392505050565b60006115b382610f62565b91506115be83610f62565b92508282039050818111156115d6576115d56113eb565b5b92915050565b7f616e74655b000000000000000000000000000000000000000000000000000000815250565b7f5d2067616d654665655b00000000000000000000000000000000000000000000815250565b7f5d20706f745b0000000000000000000000000000000000000000000000000000815250565b6000611659826115dc565b60058201915061166982866112bd565b915061167482611602565b600a8201915061168482856112bd565b915061168f82611628565b60068201915061169f82846112bd565b91506116aa82611314565b600182019150819050949350505050565b7f6e6f20706f74207761732063726561746564206261736564206f6e206163636f60008201527f756e742062616c616e6365730000000000000000000000000000000000000000602082015250565b6000611717602c83610ffd565b9150611722826116bb565b604082019050919050565b600060208201905081810360008301526117468161170a565b9050919050565b7f706f74206c657373207468616e206665653a2077696e6e65725b305d206f776e60008201527f65725b0000000000000000000000000000000000000000000000000000000000602082015250565b60006117a96023836112b2565b91506117b48261174d565b602382019050919050565b60006117ca8261179c565b91506117d682846112bd565b91506117e182611314565b60018201915081905092915050565b7f77696e6e65725b00000000000000000000000000000000000000000000000000815250565b7f5d206f776e65725b000000000000000000000000000000000000000000000000815250565b6000611847826117f0565b60078201915061185782856112bd565b915061186282611816565b60088201915061187282846112bd565b915061187d82611314565b6001820191508190509392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60006118c782610f62565b91506118d283610f62565b92508282026118e081610f62565b915082820484148315176118f7576118f66113eb565b5b5092915050565b60008160011c9050919050565b6000808291508390505b600185111561195557808604811115611931576119306113eb565b5b60018516156119405780820291505b808102905061194e856118fe565b9450611915565b94509492505050565b60008261196e5760019050611a2a565b8161197c5760009050611a2a565b8160018114611992576002811461199c576119cb565b6001915050611a2a565b60ff8411156119ae576119ad6113eb565b5b8360020a9150848211156119c5576119c46113eb565b5b50611a2a565b5060208310610133831016604e8410600b8410161715611a005782820a9050838111156119fb576119fa6113eb565b5b611a2a565b611a0d848484600161190b565b92509050818404811115611a2457611a236113eb565b5b81810290505b9392505050565b6000611a3c82610f62565b9150611a4783610f62565b9250611a747fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff848461195e565b905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000611ab682610f62565b9150611ac183610f62565b925082611ad157611ad0611a7c565b5b828204905092915050565b600060ff82169050919050565b6000611af482611adc565b9150611aff83611adc565b925082611b0f57611b0e611a7c565b5b828204905092915050565b6000611b2582611adc565b9150611b3083611adc565b9250828202611b3e81611adc565b9150808214611b5057611b4f6113eb565b5b5092915050565b6000611b6282611adc565b9150611b6d83611adc565b9250828203905060ff811115611b8657611b856113eb565b5b92915050565b6000611b9782610f62565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611bc957611bc86113eb565b5b600182019050919050565b6000611bdf82611adc565b9150611bea83611adc565b9250828201905060ff811115611c0357611c026113eb565b5b9291505056fea2646970667358221220884f62408ff065942972f9a3cabc7727a0cfb481092e65cc0b9c01e564b55dd864736f6c63430008190033",
}

// BankABI is the input ABI used to generate the binding from.
// Deprecated: Use BankMetaData.ABI instead.
var BankABI = BankMetaData.ABI

// BankBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BankMetaData.Bin instead.
var BankBin = BankMetaData.Bin

// DeployBank deploys a new Ethereum contract, binding an instance of Bank to it.
func DeployBank(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Bank, error) {
	parsed, err := BankMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BankBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Bank{BankCaller: BankCaller{contract: contract}, BankTransactor: BankTransactor{contract: contract}, BankFilterer: BankFilterer{contract: contract}}, nil
}

// Bank is an auto generated Go binding around an Ethereum contract.
type Bank struct {
	BankCaller     // Read-only binding to the contract
	BankTransactor // Write-only binding to the contract
	BankFilterer   // Log filterer for contract events
}

// BankCaller is an auto generated read-only Go binding around an Ethereum contract.
type BankCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BankTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BankTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BankFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BankFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BankSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BankSession struct {
	Contract     *Bank             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BankCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BankCallerSession struct {
	Contract *BankCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BankTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BankTransactorSession struct {
	Contract     *BankTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BankRaw is an auto generated low-level Go binding around an Ethereum contract.
type BankRaw struct {
	Contract *Bank // Generic contract binding to access the raw methods on
}

// BankCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BankCallerRaw struct {
	Contract *BankCaller // Generic read-only contract binding to access the raw methods on
}

// BankTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BankTransactorRaw struct {
	Contract *BankTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBank creates a new instance of Bank, bound to a specific deployed contract.
func NewBank(address common.Address, backend bind.ContractBackend) (*Bank, error) {
	contract, err := bindBank(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bank{BankCaller: BankCaller{contract: contract}, BankTransactor: BankTransactor{contract: contract}, BankFilterer: BankFilterer{contract: contract}}, nil
}

// NewBankCaller creates a new read-only instance of Bank, bound to a specific deployed contract.
func NewBankCaller(address common.Address, caller bind.ContractCaller) (*BankCaller, error) {
	contract, err := bindBank(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BankCaller{contract: contract}, nil
}

// NewBankTransactor creates a new write-only instance of Bank, bound to a specific deployed contract.
func NewBankTransactor(address common.Address, transactor bind.ContractTransactor) (*BankTransactor, error) {
	contract, err := bindBank(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BankTransactor{contract: contract}, nil
}

// NewBankFilterer creates a new log filterer instance of Bank, bound to a specific deployed contract.
func NewBankFilterer(address common.Address, filterer bind.ContractFilterer) (*BankFilterer, error) {
	contract, err := bindBank(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BankFilterer{contract: contract}, nil
}

// bindBank binds a generic wrapper to an already deployed contract.
func bindBank(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BankMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bank *BankRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bank.Contract.BankCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bank *BankRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bank.Contract.BankTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bank *BankRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bank.Contract.BankTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bank *BankCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bank.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bank *BankTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bank.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bank *BankTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bank.Contract.contract.Transact(opts, method, params...)
}

// AccountBalance is a free data retrieval call binding the contract method 0xe63f341f.
//
// Solidity: function AccountBalance(address account) view returns(uint256)
func (_Bank *BankCaller) AccountBalance(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bank.contract.Call(opts, &out, "AccountBalance", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccountBalance is a free data retrieval call binding the contract method 0xe63f341f.
//
// Solidity: function AccountBalance(address account) view returns(uint256)
func (_Bank *BankSession) AccountBalance(account common.Address) (*big.Int, error) {
	return _Bank.Contract.AccountBalance(&_Bank.CallOpts, account)
}

// AccountBalance is a free data retrieval call binding the contract method 0xe63f341f.
//
// Solidity: function AccountBalance(address account) view returns(uint256)
func (_Bank *BankCallerSession) AccountBalance(account common.Address) (*big.Int, error) {
	return _Bank.Contract.AccountBalance(&_Bank.CallOpts, account)
}

// Balance is a free data retrieval call binding the contract method 0x0ef67887.
//
// Solidity: function Balance() view returns(uint256)
func (_Bank *BankCaller) Balance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bank.contract.Call(opts, &out, "Balance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balance is a free data retrieval call binding the contract method 0x0ef67887.
//
// Solidity: function Balance() view returns(uint256)
func (_Bank *BankSession) Balance() (*big.Int, error) {
	return _Bank.Contract.Balance(&_Bank.CallOpts)
}

// Balance is a free data retrieval call binding the contract method 0x0ef67887.
//
// Solidity: function Balance() view returns(uint256)
func (_Bank *BankCallerSession) Balance() (*big.Int, error) {
	return _Bank.Contract.Balance(&_Bank.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0xb4a99a4e.
//
// Solidity: function Owner() view returns(address)
func (_Bank *BankCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bank.contract.Call(opts, &out, "Owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0xb4a99a4e.
//
// Solidity: function Owner() view returns(address)
func (_Bank *BankSession) Owner() (common.Address, error) {
	return _Bank.Contract.Owner(&_Bank.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0xb4a99a4e.
//
// Solidity: function Owner() view returns(address)
func (_Bank *BankCallerSession) Owner() (common.Address, error) {
	return _Bank.Contract.Owner(&_Bank.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0xbb62860d.
//
// Solidity: function Version() view returns(string)
func (_Bank *BankCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Bank.contract.Call(opts, &out, "Version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0xbb62860d.
//
// Solidity: function Version() view returns(string)
func (_Bank *BankSession) Version() (string, error) {
	return _Bank.Contract.Version(&_Bank.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0xbb62860d.
//
// Solidity: function Version() view returns(string)
func (_Bank *BankCallerSession) Version() (string, error) {
	return _Bank.Contract.Version(&_Bank.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xed21248c.
//
// Solidity: function Deposit() payable returns()
func (_Bank *BankTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bank.contract.Transact(opts, "Deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xed21248c.
//
// Solidity: function Deposit() payable returns()
func (_Bank *BankSession) Deposit() (*types.Transaction, error) {
	return _Bank.Contract.Deposit(&_Bank.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xed21248c.
//
// Solidity: function Deposit() payable returns()
func (_Bank *BankTransactorSession) Deposit() (*types.Transaction, error) {
	return _Bank.Contract.Deposit(&_Bank.TransactOpts)
}

// Reconcile is a paid mutator transaction binding the contract method 0xfa84fd8e.
//
// Solidity: function Reconcile(address winner, address[] losers, uint256 anteWei, uint256 gameFeeWei) returns()
func (_Bank *BankTransactor) Reconcile(opts *bind.TransactOpts, winner common.Address, losers []common.Address, anteWei *big.Int, gameFeeWei *big.Int) (*types.Transaction, error) {
	return _Bank.contract.Transact(opts, "Reconcile", winner, losers, anteWei, gameFeeWei)
}

// Reconcile is a paid mutator transaction binding the contract method 0xfa84fd8e.
//
// Solidity: function Reconcile(address winner, address[] losers, uint256 anteWei, uint256 gameFeeWei) returns()
func (_Bank *BankSession) Reconcile(winner common.Address, losers []common.Address, anteWei *big.Int, gameFeeWei *big.Int) (*types.Transaction, error) {
	return _Bank.Contract.Reconcile(&_Bank.TransactOpts, winner, losers, anteWei, gameFeeWei)
}

// Reconcile is a paid mutator transaction binding the contract method 0xfa84fd8e.
//
// Solidity: function Reconcile(address winner, address[] losers, uint256 anteWei, uint256 gameFeeWei) returns()
func (_Bank *BankTransactorSession) Reconcile(winner common.Address, losers []common.Address, anteWei *big.Int, gameFeeWei *big.Int) (*types.Transaction, error) {
	return _Bank.Contract.Reconcile(&_Bank.TransactOpts, winner, losers, anteWei, gameFeeWei)
}

// Withdraw is a paid mutator transaction binding the contract method 0x57ea89b6.
//
// Solidity: function Withdraw() payable returns()
func (_Bank *BankTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bank.contract.Transact(opts, "Withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x57ea89b6.
//
// Solidity: function Withdraw() payable returns()
func (_Bank *BankSession) Withdraw() (*types.Transaction, error) {
	return _Bank.Contract.Withdraw(&_Bank.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x57ea89b6.
//
// Solidity: function Withdraw() payable returns()
func (_Bank *BankTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Bank.Contract.Withdraw(&_Bank.TransactOpts)
}

// BankEventLogIterator is returned from FilterEventLog and is used to iterate over the raw logs and unpacked data for EventLog events raised by the Bank contract.
type BankEventLogIterator struct {
	Event *BankEventLog // Event containing the contract specifics and raw log

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
func (it *BankEventLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BankEventLog)
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
		it.Event = new(BankEventLog)
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
func (it *BankEventLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BankEventLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BankEventLog represents a EventLog event raised by the Bank contract.
type BankEventLog struct {
	Value string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterEventLog is a free log retrieval operation binding the contract event 0xd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a.
//
// Solidity: event EventLog(string value)
func (_Bank *BankFilterer) FilterEventLog(opts *bind.FilterOpts) (*BankEventLogIterator, error) {

	logs, sub, err := _Bank.contract.FilterLogs(opts, "EventLog")
	if err != nil {
		return nil, err
	}
	return &BankEventLogIterator{contract: _Bank.contract, event: "EventLog", logs: logs, sub: sub}, nil
}

// WatchEventLog is a free log subscription operation binding the contract event 0xd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a.
//
// Solidity: event EventLog(string value)
func (_Bank *BankFilterer) WatchEventLog(opts *bind.WatchOpts, sink chan<- *BankEventLog) (event.Subscription, error) {

	logs, sub, err := _Bank.contract.WatchLogs(opts, "EventLog")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BankEventLog)
				if err := _Bank.contract.UnpackLog(event, "EventLog", log); err != nil {
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
func (_Bank *BankFilterer) ParseEventLog(log types.Log) (*BankEventLog, error) {
	event := new(BankEventLog)
	if err := _Bank.contract.UnpackLog(event, "EventLog", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
