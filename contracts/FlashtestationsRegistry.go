// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// IFlashtestationRegistryRegisteredTEE is an auto generated low-level Go binding around an user-defined struct.
type IFlashtestationRegistryRegisteredTEE struct {
	IsValid                  bool
	RawQuote                 []byte
	ParsedReportBody         TD10ReportBody
	ExtendedRegistrationData []byte
	QuoteHash                [32]byte
}

// FlashtestationsRegistryMetaData contains all meta data concerning the FlashtestationsRegistry contract.
var FlashtestationsRegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"MAX_BYTES_SIZE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"REGISTER_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TD_REPORTDATA_LENGTH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"attestationContract\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAttestation\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"computeStructHash\",\"inputs\":[{\"name\":\"rawQuote\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extendedRegistrationData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"domainSeparator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"eip712Domain\",\"inputs\":[],\"outputs\":[{\"name\":\"fields\",\"type\":\"bytes1\",\"internalType\":\"bytes1\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"verifyingContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"extensions\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRegistration\",\"inputs\":[{\"name\":\"teeAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIFlashtestationRegistry.RegisteredTEE\",\"components\":[{\"name\":\"isValid\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"rawQuote\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"parsedReportBody\",\"type\":\"tuple\",\"internalType\":\"structTD10ReportBody\",\"components\":[{\"name\":\"teeTcbSvn\",\"type\":\"bytes16\",\"internalType\":\"bytes16\"},{\"name\":\"mrSeam\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"mrsignerSeam\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"seamAttributes\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"},{\"name\":\"tdAttributes\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"},{\"name\":\"xFAM\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"},{\"name\":\"mrTd\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"mrConfigId\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"mrOwner\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"mrOwnerConfig\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"rtMr0\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"rtMr1\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"rtMr2\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"rtMr3\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"reportData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"extendedRegistrationData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quoteHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRegistrationStatus\",\"inputs\":[{\"name\":\"teeAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"isValid\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"quoteHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hashTypedDataV4\",\"inputs\":[{\"name\":\"structHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_attestationContract\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"invalidateAttestation\",\"inputs\":[{\"name\":\"teeAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"invalidatePreviousSignature\",\"inputs\":[{\"name\":\"_nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"nonces\",\"inputs\":[{\"name\":\"teeAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"permitNonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"permitRegisterTEEService\",\"inputs\":[{\"name\":\"rawQuote\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extendedRegistrationData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerTEEService\",\"inputs\":[{\"name\":\"rawQuote\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extendedRegistrationData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"registeredTEEs\",\"inputs\":[{\"name\":\"teeAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"isValid\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"rawQuote\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"parsedReportBody\",\"type\":\"tuple\",\"internalType\":\"structTD10ReportBody\",\"components\":[{\"name\":\"teeTcbSvn\",\"type\":\"bytes16\",\"internalType\":\"bytes16\"},{\"name\":\"mrSeam\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"mrsignerSeam\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"seamAttributes\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"},{\"name\":\"tdAttributes\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"},{\"name\":\"xFAM\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"},{\"name\":\"mrTd\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"mrConfigId\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"mrOwner\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"mrOwnerConfig\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"rtMr0\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"rtMr1\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"rtMr2\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"rtMr3\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"reportData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"extendedRegistrationData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quoteHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"EIP712DomainChanged\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PreviousSignatureInvalidated\",\"inputs\":[{\"name\":\"teeAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"invalidatedNonce\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TEEServiceInvalidated\",\"inputs\":[{\"name\":\"teeAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TEEServiceRegistered\",\"inputs\":[{\"name\":\"teeAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"rawQuote\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"alreadyExists\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ByteSizeExceeded\",\"inputs\":[{\"name\":\"size\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureS\",\"inputs\":[{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpiredSignature\",\"inputs\":[{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"FailedCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAttestationContract\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidNonce\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidQuote\",\"inputs\":[{\"name\":\"output\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidQuoteLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidRegistrationDataHash\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"received\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidReportDataLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidTEEType\",\"inputs\":[{\"name\":\"teeType\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"InvalidTEEVersion\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SignerMustMatchTEEAddress\",\"inputs\":[{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"teeAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TEEIsStillValid\",\"inputs\":[{\"name\":\"teeAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TEEServiceAlreadyInvalid\",\"inputs\":[{\"name\":\"teeAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TEEServiceAlreadyRegistered\",\"inputs\":[{\"name\":\"teeAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TEEServiceNotRegistered\",\"inputs\":[{\"name\":\"teeAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
}

// FlashtestationsRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use FlashtestationsRegistryMetaData.ABI instead.
var FlashtestationsRegistryABI = FlashtestationsRegistryMetaData.ABI

// FlashtestationsRegistry is an auto generated Go binding around an Ethereum contract.
type FlashtestationsRegistry struct {
	FlashtestationsRegistryCaller     // Read-only binding to the contract
	FlashtestationsRegistryTransactor // Write-only binding to the contract
	FlashtestationsRegistryFilterer   // Log filterer for contract events
}

// FlashtestationsRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type FlashtestationsRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FlashtestationsRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FlashtestationsRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FlashtestationsRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FlashtestationsRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FlashtestationsRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FlashtestationsRegistrySession struct {
	Contract     *FlashtestationsRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// FlashtestationsRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FlashtestationsRegistryCallerSession struct {
	Contract *FlashtestationsRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// FlashtestationsRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FlashtestationsRegistryTransactorSession struct {
	Contract     *FlashtestationsRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// FlashtestationsRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type FlashtestationsRegistryRaw struct {
	Contract *FlashtestationsRegistry // Generic contract binding to access the raw methods on
}

// FlashtestationsRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FlashtestationsRegistryCallerRaw struct {
	Contract *FlashtestationsRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// FlashtestationsRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FlashtestationsRegistryTransactorRaw struct {
	Contract *FlashtestationsRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFlashtestationsRegistry creates a new instance of FlashtestationsRegistry, bound to a specific deployed contract.
func NewFlashtestationsRegistry(address common.Address, backend bind.ContractBackend) (*FlashtestationsRegistry, error) {
	contract, err := bindFlashtestationsRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FlashtestationsRegistry{FlashtestationsRegistryCaller: FlashtestationsRegistryCaller{contract: contract}, FlashtestationsRegistryTransactor: FlashtestationsRegistryTransactor{contract: contract}, FlashtestationsRegistryFilterer: FlashtestationsRegistryFilterer{contract: contract}}, nil
}

// NewFlashtestationsRegistryCaller creates a new read-only instance of FlashtestationsRegistry, bound to a specific deployed contract.
func NewFlashtestationsRegistryCaller(address common.Address, caller bind.ContractCaller) (*FlashtestationsRegistryCaller, error) {
	contract, err := bindFlashtestationsRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FlashtestationsRegistryCaller{contract: contract}, nil
}

// NewFlashtestationsRegistryTransactor creates a new write-only instance of FlashtestationsRegistry, bound to a specific deployed contract.
func NewFlashtestationsRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*FlashtestationsRegistryTransactor, error) {
	contract, err := bindFlashtestationsRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FlashtestationsRegistryTransactor{contract: contract}, nil
}

// NewFlashtestationsRegistryFilterer creates a new log filterer instance of FlashtestationsRegistry, bound to a specific deployed contract.
func NewFlashtestationsRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*FlashtestationsRegistryFilterer, error) {
	contract, err := bindFlashtestationsRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FlashtestationsRegistryFilterer{contract: contract}, nil
}

// bindFlashtestationsRegistry binds a generic wrapper to an already deployed contract.
func bindFlashtestationsRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FlashtestationsRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FlashtestationsRegistry *FlashtestationsRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FlashtestationsRegistry.Contract.FlashtestationsRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FlashtestationsRegistry *FlashtestationsRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FlashtestationsRegistry.Contract.FlashtestationsRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FlashtestationsRegistry *FlashtestationsRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FlashtestationsRegistry.Contract.FlashtestationsRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FlashtestationsRegistry *FlashtestationsRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FlashtestationsRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FlashtestationsRegistry *FlashtestationsRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FlashtestationsRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FlashtestationsRegistry *FlashtestationsRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FlashtestationsRegistry.Contract.contract.Transact(opts, method, params...)
}

// MAXBYTESSIZE is a free data retrieval call binding the contract method 0xaaae748e.
//
// Solidity: function MAX_BYTES_SIZE() view returns(uint256)
func (_FlashtestationsRegistry *FlashtestationsRegistryCaller) MAXBYTESSIZE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FlashtestationsRegistry.contract.Call(opts, &out, "MAX_BYTES_SIZE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXBYTESSIZE is a free data retrieval call binding the contract method 0xaaae748e.
//
// Solidity: function MAX_BYTES_SIZE() view returns(uint256)
func (_FlashtestationsRegistry *FlashtestationsRegistrySession) MAXBYTESSIZE() (*big.Int, error) {
	return _FlashtestationsRegistry.Contract.MAXBYTESSIZE(&_FlashtestationsRegistry.CallOpts)
}

// MAXBYTESSIZE is a free data retrieval call binding the contract method 0xaaae748e.
//
// Solidity: function MAX_BYTES_SIZE() view returns(uint256)
func (_FlashtestationsRegistry *FlashtestationsRegistryCallerSession) MAXBYTESSIZE() (*big.Int, error) {
	return _FlashtestationsRegistry.Contract.MAXBYTESSIZE(&_FlashtestationsRegistry.CallOpts)
}

// REGISTERTYPEHASH is a free data retrieval call binding the contract method 0x6a5306a3.
//
// Solidity: function REGISTER_TYPEHASH() view returns(bytes32)
func (_FlashtestationsRegistry *FlashtestationsRegistryCaller) REGISTERTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _FlashtestationsRegistry.contract.Call(opts, &out, "REGISTER_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// REGISTERTYPEHASH is a free data retrieval call binding the contract method 0x6a5306a3.
//
// Solidity: function REGISTER_TYPEHASH() view returns(bytes32)
func (_FlashtestationsRegistry *FlashtestationsRegistrySession) REGISTERTYPEHASH() ([32]byte, error) {
	return _FlashtestationsRegistry.Contract.REGISTERTYPEHASH(&_FlashtestationsRegistry.CallOpts)
}

// REGISTERTYPEHASH is a free data retrieval call binding the contract method 0x6a5306a3.
//
// Solidity: function REGISTER_TYPEHASH() view returns(bytes32)
func (_FlashtestationsRegistry *FlashtestationsRegistryCallerSession) REGISTERTYPEHASH() ([32]byte, error) {
	return _FlashtestationsRegistry.Contract.REGISTERTYPEHASH(&_FlashtestationsRegistry.CallOpts)
}

// TDREPORTDATALENGTH is a free data retrieval call binding the contract method 0xe4168952.
//
// Solidity: function TD_REPORTDATA_LENGTH() view returns(uint256)
func (_FlashtestationsRegistry *FlashtestationsRegistryCaller) TDREPORTDATALENGTH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FlashtestationsRegistry.contract.Call(opts, &out, "TD_REPORTDATA_LENGTH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TDREPORTDATALENGTH is a free data retrieval call binding the contract method 0xe4168952.
//
// Solidity: function TD_REPORTDATA_LENGTH() view returns(uint256)
func (_FlashtestationsRegistry *FlashtestationsRegistrySession) TDREPORTDATALENGTH() (*big.Int, error) {
	return _FlashtestationsRegistry.Contract.TDREPORTDATALENGTH(&_FlashtestationsRegistry.CallOpts)
}

// TDREPORTDATALENGTH is a free data retrieval call binding the contract method 0xe4168952.
//
// Solidity: function TD_REPORTDATA_LENGTH() view returns(uint256)
func (_FlashtestationsRegistry *FlashtestationsRegistryCallerSession) TDREPORTDATALENGTH() (*big.Int, error) {
	return _FlashtestationsRegistry.Contract.TDREPORTDATALENGTH(&_FlashtestationsRegistry.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_FlashtestationsRegistry *FlashtestationsRegistryCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FlashtestationsRegistry.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_FlashtestationsRegistry *FlashtestationsRegistrySession) UPGRADEINTERFACEVERSION() (string, error) {
	return _FlashtestationsRegistry.Contract.UPGRADEINTERFACEVERSION(&_FlashtestationsRegistry.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_FlashtestationsRegistry *FlashtestationsRegistryCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _FlashtestationsRegistry.Contract.UPGRADEINTERFACEVERSION(&_FlashtestationsRegistry.CallOpts)
}

// AttestationContract is a free data retrieval call binding the contract method 0x87be6d4e.
//
// Solidity: function attestationContract() view returns(address)
func (_FlashtestationsRegistry *FlashtestationsRegistryCaller) AttestationContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FlashtestationsRegistry.contract.Call(opts, &out, "attestationContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AttestationContract is a free data retrieval call binding the contract method 0x87be6d4e.
//
// Solidity: function attestationContract() view returns(address)
func (_FlashtestationsRegistry *FlashtestationsRegistrySession) AttestationContract() (common.Address, error) {
	return _FlashtestationsRegistry.Contract.AttestationContract(&_FlashtestationsRegistry.CallOpts)
}

// AttestationContract is a free data retrieval call binding the contract method 0x87be6d4e.
//
// Solidity: function attestationContract() view returns(address)
func (_FlashtestationsRegistry *FlashtestationsRegistryCallerSession) AttestationContract() (common.Address, error) {
	return _FlashtestationsRegistry.Contract.AttestationContract(&_FlashtestationsRegistry.CallOpts)
}

// ComputeStructHash is a free data retrieval call binding the contract method 0x0634434a.
//
// Solidity: function computeStructHash(bytes rawQuote, bytes extendedRegistrationData, uint256 nonce, uint256 deadline) pure returns(bytes32)
func (_FlashtestationsRegistry *FlashtestationsRegistryCaller) ComputeStructHash(opts *bind.CallOpts, rawQuote []byte, extendedRegistrationData []byte, nonce *big.Int, deadline *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _FlashtestationsRegistry.contract.Call(opts, &out, "computeStructHash", rawQuote, extendedRegistrationData, nonce, deadline)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ComputeStructHash is a free data retrieval call binding the contract method 0x0634434a.
//
// Solidity: function computeStructHash(bytes rawQuote, bytes extendedRegistrationData, uint256 nonce, uint256 deadline) pure returns(bytes32)
func (_FlashtestationsRegistry *FlashtestationsRegistrySession) ComputeStructHash(rawQuote []byte, extendedRegistrationData []byte, nonce *big.Int, deadline *big.Int) ([32]byte, error) {
	return _FlashtestationsRegistry.Contract.ComputeStructHash(&_FlashtestationsRegistry.CallOpts, rawQuote, extendedRegistrationData, nonce, deadline)
}

// ComputeStructHash is a free data retrieval call binding the contract method 0x0634434a.
//
// Solidity: function computeStructHash(bytes rawQuote, bytes extendedRegistrationData, uint256 nonce, uint256 deadline) pure returns(bytes32)
func (_FlashtestationsRegistry *FlashtestationsRegistryCallerSession) ComputeStructHash(rawQuote []byte, extendedRegistrationData []byte, nonce *big.Int, deadline *big.Int) ([32]byte, error) {
	return _FlashtestationsRegistry.Contract.ComputeStructHash(&_FlashtestationsRegistry.CallOpts, rawQuote, extendedRegistrationData, nonce, deadline)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_FlashtestationsRegistry *FlashtestationsRegistryCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _FlashtestationsRegistry.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_FlashtestationsRegistry *FlashtestationsRegistrySession) DomainSeparator() ([32]byte, error) {
	return _FlashtestationsRegistry.Contract.DomainSeparator(&_FlashtestationsRegistry.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_FlashtestationsRegistry *FlashtestationsRegistryCallerSession) DomainSeparator() ([32]byte, error) {
	return _FlashtestationsRegistry.Contract.DomainSeparator(&_FlashtestationsRegistry.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_FlashtestationsRegistry *FlashtestationsRegistryCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _FlashtestationsRegistry.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_FlashtestationsRegistry *FlashtestationsRegistrySession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _FlashtestationsRegistry.Contract.Eip712Domain(&_FlashtestationsRegistry.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_FlashtestationsRegistry *FlashtestationsRegistryCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _FlashtestationsRegistry.Contract.Eip712Domain(&_FlashtestationsRegistry.CallOpts)
}

// GetRegistration is a free data retrieval call binding the contract method 0x72731062.
//
// Solidity: function getRegistration(address teeAddress) view returns(bool, (bool,bytes,(bytes16,bytes,bytes,bytes8,bytes8,bytes8,bytes,bytes,bytes,bytes,bytes,bytes,bytes,bytes,bytes),bytes,bytes32))
func (_FlashtestationsRegistry *FlashtestationsRegistryCaller) GetRegistration(opts *bind.CallOpts, teeAddress common.Address) (bool, IFlashtestationRegistryRegisteredTEE, error) {
	var out []interface{}
	err := _FlashtestationsRegistry.contract.Call(opts, &out, "getRegistration", teeAddress)

	if err != nil {
		return *new(bool), *new(IFlashtestationRegistryRegisteredTEE), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(IFlashtestationRegistryRegisteredTEE)).(*IFlashtestationRegistryRegisteredTEE)

	return out0, out1, err

}

// GetRegistration is a free data retrieval call binding the contract method 0x72731062.
//
// Solidity: function getRegistration(address teeAddress) view returns(bool, (bool,bytes,(bytes16,bytes,bytes,bytes8,bytes8,bytes8,bytes,bytes,bytes,bytes,bytes,bytes,bytes,bytes,bytes),bytes,bytes32))
func (_FlashtestationsRegistry *FlashtestationsRegistrySession) GetRegistration(teeAddress common.Address) (bool, IFlashtestationRegistryRegisteredTEE, error) {
	return _FlashtestationsRegistry.Contract.GetRegistration(&_FlashtestationsRegistry.CallOpts, teeAddress)
}

// GetRegistration is a free data retrieval call binding the contract method 0x72731062.
//
// Solidity: function getRegistration(address teeAddress) view returns(bool, (bool,bytes,(bytes16,bytes,bytes,bytes8,bytes8,bytes8,bytes,bytes,bytes,bytes,bytes,bytes,bytes,bytes,bytes),bytes,bytes32))
func (_FlashtestationsRegistry *FlashtestationsRegistryCallerSession) GetRegistration(teeAddress common.Address) (bool, IFlashtestationRegistryRegisteredTEE, error) {
	return _FlashtestationsRegistry.Contract.GetRegistration(&_FlashtestationsRegistry.CallOpts, teeAddress)
}

// GetRegistrationStatus is a free data retrieval call binding the contract method 0xa8af4ff5.
//
// Solidity: function getRegistrationStatus(address teeAddress) view returns(bool isValid, bytes32 quoteHash)
func (_FlashtestationsRegistry *FlashtestationsRegistryCaller) GetRegistrationStatus(opts *bind.CallOpts, teeAddress common.Address) (struct {
	IsValid   bool
	QuoteHash [32]byte
}, error) {
	var out []interface{}
	err := _FlashtestationsRegistry.contract.Call(opts, &out, "getRegistrationStatus", teeAddress)

	outstruct := new(struct {
		IsValid   bool
		QuoteHash [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsValid = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.QuoteHash = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// GetRegistrationStatus is a free data retrieval call binding the contract method 0xa8af4ff5.
//
// Solidity: function getRegistrationStatus(address teeAddress) view returns(bool isValid, bytes32 quoteHash)
func (_FlashtestationsRegistry *FlashtestationsRegistrySession) GetRegistrationStatus(teeAddress common.Address) (struct {
	IsValid   bool
	QuoteHash [32]byte
}, error) {
	return _FlashtestationsRegistry.Contract.GetRegistrationStatus(&_FlashtestationsRegistry.CallOpts, teeAddress)
}

// GetRegistrationStatus is a free data retrieval call binding the contract method 0xa8af4ff5.
//
// Solidity: function getRegistrationStatus(address teeAddress) view returns(bool isValid, bytes32 quoteHash)
func (_FlashtestationsRegistry *FlashtestationsRegistryCallerSession) GetRegistrationStatus(teeAddress common.Address) (struct {
	IsValid   bool
	QuoteHash [32]byte
}, error) {
	return _FlashtestationsRegistry.Contract.GetRegistrationStatus(&_FlashtestationsRegistry.CallOpts, teeAddress)
}

// HashTypedDataV4 is a free data retrieval call binding the contract method 0x4980f288.
//
// Solidity: function hashTypedDataV4(bytes32 structHash) view returns(bytes32)
func (_FlashtestationsRegistry *FlashtestationsRegistryCaller) HashTypedDataV4(opts *bind.CallOpts, structHash [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _FlashtestationsRegistry.contract.Call(opts, &out, "hashTypedDataV4", structHash)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashTypedDataV4 is a free data retrieval call binding the contract method 0x4980f288.
//
// Solidity: function hashTypedDataV4(bytes32 structHash) view returns(bytes32)
func (_FlashtestationsRegistry *FlashtestationsRegistrySession) HashTypedDataV4(structHash [32]byte) ([32]byte, error) {
	return _FlashtestationsRegistry.Contract.HashTypedDataV4(&_FlashtestationsRegistry.CallOpts, structHash)
}

// HashTypedDataV4 is a free data retrieval call binding the contract method 0x4980f288.
//
// Solidity: function hashTypedDataV4(bytes32 structHash) view returns(bytes32)
func (_FlashtestationsRegistry *FlashtestationsRegistryCallerSession) HashTypedDataV4(structHash [32]byte) ([32]byte, error) {
	return _FlashtestationsRegistry.Contract.HashTypedDataV4(&_FlashtestationsRegistry.CallOpts, structHash)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address teeAddress) view returns(uint256 permitNonce)
func (_FlashtestationsRegistry *FlashtestationsRegistryCaller) Nonces(opts *bind.CallOpts, teeAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FlashtestationsRegistry.contract.Call(opts, &out, "nonces", teeAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address teeAddress) view returns(uint256 permitNonce)
func (_FlashtestationsRegistry *FlashtestationsRegistrySession) Nonces(teeAddress common.Address) (*big.Int, error) {
	return _FlashtestationsRegistry.Contract.Nonces(&_FlashtestationsRegistry.CallOpts, teeAddress)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address teeAddress) view returns(uint256 permitNonce)
func (_FlashtestationsRegistry *FlashtestationsRegistryCallerSession) Nonces(teeAddress common.Address) (*big.Int, error) {
	return _FlashtestationsRegistry.Contract.Nonces(&_FlashtestationsRegistry.CallOpts, teeAddress)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FlashtestationsRegistry *FlashtestationsRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FlashtestationsRegistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FlashtestationsRegistry *FlashtestationsRegistrySession) Owner() (common.Address, error) {
	return _FlashtestationsRegistry.Contract.Owner(&_FlashtestationsRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FlashtestationsRegistry *FlashtestationsRegistryCallerSession) Owner() (common.Address, error) {
	return _FlashtestationsRegistry.Contract.Owner(&_FlashtestationsRegistry.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_FlashtestationsRegistry *FlashtestationsRegistryCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _FlashtestationsRegistry.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_FlashtestationsRegistry *FlashtestationsRegistrySession) ProxiableUUID() ([32]byte, error) {
	return _FlashtestationsRegistry.Contract.ProxiableUUID(&_FlashtestationsRegistry.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_FlashtestationsRegistry *FlashtestationsRegistryCallerSession) ProxiableUUID() ([32]byte, error) {
	return _FlashtestationsRegistry.Contract.ProxiableUUID(&_FlashtestationsRegistry.CallOpts)
}

// RegisteredTEEs is a free data retrieval call binding the contract method 0xf745cb30.
//
// Solidity: function registeredTEEs(address teeAddress) view returns(bool isValid, bytes rawQuote, (bytes16,bytes,bytes,bytes8,bytes8,bytes8,bytes,bytes,bytes,bytes,bytes,bytes,bytes,bytes,bytes) parsedReportBody, bytes extendedRegistrationData, bytes32 quoteHash)
func (_FlashtestationsRegistry *FlashtestationsRegistryCaller) RegisteredTEEs(opts *bind.CallOpts, teeAddress common.Address) (struct {
	IsValid                  bool
	RawQuote                 []byte
	ParsedReportBody         TD10ReportBody
	ExtendedRegistrationData []byte
	QuoteHash                [32]byte
}, error) {
	var out []interface{}
	err := _FlashtestationsRegistry.contract.Call(opts, &out, "registeredTEEs", teeAddress)

	outstruct := new(struct {
		IsValid                  bool
		RawQuote                 []byte
		ParsedReportBody         TD10ReportBody
		ExtendedRegistrationData []byte
		QuoteHash                [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsValid = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.RawQuote = *abi.ConvertType(out[1], new([]byte)).(*[]byte)
	outstruct.ParsedReportBody = *abi.ConvertType(out[2], new(TD10ReportBody)).(*TD10ReportBody)
	outstruct.ExtendedRegistrationData = *abi.ConvertType(out[3], new([]byte)).(*[]byte)
	outstruct.QuoteHash = *abi.ConvertType(out[4], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// RegisteredTEEs is a free data retrieval call binding the contract method 0xf745cb30.
//
// Solidity: function registeredTEEs(address teeAddress) view returns(bool isValid, bytes rawQuote, (bytes16,bytes,bytes,bytes8,bytes8,bytes8,bytes,bytes,bytes,bytes,bytes,bytes,bytes,bytes,bytes) parsedReportBody, bytes extendedRegistrationData, bytes32 quoteHash)
func (_FlashtestationsRegistry *FlashtestationsRegistrySession) RegisteredTEEs(teeAddress common.Address) (struct {
	IsValid                  bool
	RawQuote                 []byte
	ParsedReportBody         TD10ReportBody
	ExtendedRegistrationData []byte
	QuoteHash                [32]byte
}, error) {
	return _FlashtestationsRegistry.Contract.RegisteredTEEs(&_FlashtestationsRegistry.CallOpts, teeAddress)
}

// RegisteredTEEs is a free data retrieval call binding the contract method 0xf745cb30.
//
// Solidity: function registeredTEEs(address teeAddress) view returns(bool isValid, bytes rawQuote, (bytes16,bytes,bytes,bytes8,bytes8,bytes8,bytes,bytes,bytes,bytes,bytes,bytes,bytes,bytes,bytes) parsedReportBody, bytes extendedRegistrationData, bytes32 quoteHash)
func (_FlashtestationsRegistry *FlashtestationsRegistryCallerSession) RegisteredTEEs(teeAddress common.Address) (struct {
	IsValid                  bool
	RawQuote                 []byte
	ParsedReportBody         TD10ReportBody
	ExtendedRegistrationData []byte
	QuoteHash                [32]byte
}, error) {
	return _FlashtestationsRegistry.Contract.RegisteredTEEs(&_FlashtestationsRegistry.CallOpts, teeAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address owner, address _attestationContract) returns()
func (_FlashtestationsRegistry *FlashtestationsRegistryTransactor) Initialize(opts *bind.TransactOpts, owner common.Address, _attestationContract common.Address) (*types.Transaction, error) {
	return _FlashtestationsRegistry.contract.Transact(opts, "initialize", owner, _attestationContract)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address owner, address _attestationContract) returns()
func (_FlashtestationsRegistry *FlashtestationsRegistrySession) Initialize(owner common.Address, _attestationContract common.Address) (*types.Transaction, error) {
	return _FlashtestationsRegistry.Contract.Initialize(&_FlashtestationsRegistry.TransactOpts, owner, _attestationContract)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address owner, address _attestationContract) returns()
func (_FlashtestationsRegistry *FlashtestationsRegistryTransactorSession) Initialize(owner common.Address, _attestationContract common.Address) (*types.Transaction, error) {
	return _FlashtestationsRegistry.Contract.Initialize(&_FlashtestationsRegistry.TransactOpts, owner, _attestationContract)
}

// InvalidateAttestation is a paid mutator transaction binding the contract method 0xf9b68b31.
//
// Solidity: function invalidateAttestation(address teeAddress) payable returns()
func (_FlashtestationsRegistry *FlashtestationsRegistryTransactor) InvalidateAttestation(opts *bind.TransactOpts, teeAddress common.Address) (*types.Transaction, error) {
	return _FlashtestationsRegistry.contract.Transact(opts, "invalidateAttestation", teeAddress)
}

// InvalidateAttestation is a paid mutator transaction binding the contract method 0xf9b68b31.
//
// Solidity: function invalidateAttestation(address teeAddress) payable returns()
func (_FlashtestationsRegistry *FlashtestationsRegistrySession) InvalidateAttestation(teeAddress common.Address) (*types.Transaction, error) {
	return _FlashtestationsRegistry.Contract.InvalidateAttestation(&_FlashtestationsRegistry.TransactOpts, teeAddress)
}

// InvalidateAttestation is a paid mutator transaction binding the contract method 0xf9b68b31.
//
// Solidity: function invalidateAttestation(address teeAddress) payable returns()
func (_FlashtestationsRegistry *FlashtestationsRegistryTransactorSession) InvalidateAttestation(teeAddress common.Address) (*types.Transaction, error) {
	return _FlashtestationsRegistry.Contract.InvalidateAttestation(&_FlashtestationsRegistry.TransactOpts, teeAddress)
}

// InvalidatePreviousSignature is a paid mutator transaction binding the contract method 0x87811112.
//
// Solidity: function invalidatePreviousSignature(uint256 _nonce) returns()
func (_FlashtestationsRegistry *FlashtestationsRegistryTransactor) InvalidatePreviousSignature(opts *bind.TransactOpts, _nonce *big.Int) (*types.Transaction, error) {
	return _FlashtestationsRegistry.contract.Transact(opts, "invalidatePreviousSignature", _nonce)
}

// InvalidatePreviousSignature is a paid mutator transaction binding the contract method 0x87811112.
//
// Solidity: function invalidatePreviousSignature(uint256 _nonce) returns()
func (_FlashtestationsRegistry *FlashtestationsRegistrySession) InvalidatePreviousSignature(_nonce *big.Int) (*types.Transaction, error) {
	return _FlashtestationsRegistry.Contract.InvalidatePreviousSignature(&_FlashtestationsRegistry.TransactOpts, _nonce)
}

// InvalidatePreviousSignature is a paid mutator transaction binding the contract method 0x87811112.
//
// Solidity: function invalidatePreviousSignature(uint256 _nonce) returns()
func (_FlashtestationsRegistry *FlashtestationsRegistryTransactorSession) InvalidatePreviousSignature(_nonce *big.Int) (*types.Transaction, error) {
	return _FlashtestationsRegistry.Contract.InvalidatePreviousSignature(&_FlashtestationsRegistry.TransactOpts, _nonce)
}

// PermitRegisterTEEService is a paid mutator transaction binding the contract method 0x0ac3302b.
//
// Solidity: function permitRegisterTEEService(bytes rawQuote, bytes extendedRegistrationData, uint256 nonce, uint256 deadline, bytes signature) payable returns()
func (_FlashtestationsRegistry *FlashtestationsRegistryTransactor) PermitRegisterTEEService(opts *bind.TransactOpts, rawQuote []byte, extendedRegistrationData []byte, nonce *big.Int, deadline *big.Int, signature []byte) (*types.Transaction, error) {
	return _FlashtestationsRegistry.contract.Transact(opts, "permitRegisterTEEService", rawQuote, extendedRegistrationData, nonce, deadline, signature)
}

// PermitRegisterTEEService is a paid mutator transaction binding the contract method 0x0ac3302b.
//
// Solidity: function permitRegisterTEEService(bytes rawQuote, bytes extendedRegistrationData, uint256 nonce, uint256 deadline, bytes signature) payable returns()
func (_FlashtestationsRegistry *FlashtestationsRegistrySession) PermitRegisterTEEService(rawQuote []byte, extendedRegistrationData []byte, nonce *big.Int, deadline *big.Int, signature []byte) (*types.Transaction, error) {
	return _FlashtestationsRegistry.Contract.PermitRegisterTEEService(&_FlashtestationsRegistry.TransactOpts, rawQuote, extendedRegistrationData, nonce, deadline, signature)
}

// PermitRegisterTEEService is a paid mutator transaction binding the contract method 0x0ac3302b.
//
// Solidity: function permitRegisterTEEService(bytes rawQuote, bytes extendedRegistrationData, uint256 nonce, uint256 deadline, bytes signature) payable returns()
func (_FlashtestationsRegistry *FlashtestationsRegistryTransactorSession) PermitRegisterTEEService(rawQuote []byte, extendedRegistrationData []byte, nonce *big.Int, deadline *big.Int, signature []byte) (*types.Transaction, error) {
	return _FlashtestationsRegistry.Contract.PermitRegisterTEEService(&_FlashtestationsRegistry.TransactOpts, rawQuote, extendedRegistrationData, nonce, deadline, signature)
}

// RegisterTEEService is a paid mutator transaction binding the contract method 0x22ba2bbf.
//
// Solidity: function registerTEEService(bytes rawQuote, bytes extendedRegistrationData) payable returns()
func (_FlashtestationsRegistry *FlashtestationsRegistryTransactor) RegisterTEEService(opts *bind.TransactOpts, rawQuote []byte, extendedRegistrationData []byte) (*types.Transaction, error) {
	return _FlashtestationsRegistry.contract.Transact(opts, "registerTEEService", rawQuote, extendedRegistrationData)
}

// RegisterTEEService is a paid mutator transaction binding the contract method 0x22ba2bbf.
//
// Solidity: function registerTEEService(bytes rawQuote, bytes extendedRegistrationData) payable returns()
func (_FlashtestationsRegistry *FlashtestationsRegistrySession) RegisterTEEService(rawQuote []byte, extendedRegistrationData []byte) (*types.Transaction, error) {
	return _FlashtestationsRegistry.Contract.RegisterTEEService(&_FlashtestationsRegistry.TransactOpts, rawQuote, extendedRegistrationData)
}

// RegisterTEEService is a paid mutator transaction binding the contract method 0x22ba2bbf.
//
// Solidity: function registerTEEService(bytes rawQuote, bytes extendedRegistrationData) payable returns()
func (_FlashtestationsRegistry *FlashtestationsRegistryTransactorSession) RegisterTEEService(rawQuote []byte, extendedRegistrationData []byte) (*types.Transaction, error) {
	return _FlashtestationsRegistry.Contract.RegisterTEEService(&_FlashtestationsRegistry.TransactOpts, rawQuote, extendedRegistrationData)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FlashtestationsRegistry *FlashtestationsRegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FlashtestationsRegistry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FlashtestationsRegistry *FlashtestationsRegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _FlashtestationsRegistry.Contract.RenounceOwnership(&_FlashtestationsRegistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FlashtestationsRegistry *FlashtestationsRegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _FlashtestationsRegistry.Contract.RenounceOwnership(&_FlashtestationsRegistry.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FlashtestationsRegistry *FlashtestationsRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _FlashtestationsRegistry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FlashtestationsRegistry *FlashtestationsRegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FlashtestationsRegistry.Contract.TransferOwnership(&_FlashtestationsRegistry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FlashtestationsRegistry *FlashtestationsRegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FlashtestationsRegistry.Contract.TransferOwnership(&_FlashtestationsRegistry.TransactOpts, newOwner)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_FlashtestationsRegistry *FlashtestationsRegistryTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _FlashtestationsRegistry.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_FlashtestationsRegistry *FlashtestationsRegistrySession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _FlashtestationsRegistry.Contract.UpgradeToAndCall(&_FlashtestationsRegistry.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_FlashtestationsRegistry *FlashtestationsRegistryTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _FlashtestationsRegistry.Contract.UpgradeToAndCall(&_FlashtestationsRegistry.TransactOpts, newImplementation, data)
}

// FlashtestationsRegistryEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the FlashtestationsRegistry contract.
type FlashtestationsRegistryEIP712DomainChangedIterator struct {
	Event *FlashtestationsRegistryEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *FlashtestationsRegistryEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlashtestationsRegistryEIP712DomainChanged)
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
		it.Event = new(FlashtestationsRegistryEIP712DomainChanged)
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
func (it *FlashtestationsRegistryEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlashtestationsRegistryEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlashtestationsRegistryEIP712DomainChanged represents a EIP712DomainChanged event raised by the FlashtestationsRegistry contract.
type FlashtestationsRegistryEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_FlashtestationsRegistry *FlashtestationsRegistryFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*FlashtestationsRegistryEIP712DomainChangedIterator, error) {

	logs, sub, err := _FlashtestationsRegistry.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &FlashtestationsRegistryEIP712DomainChangedIterator{contract: _FlashtestationsRegistry.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_FlashtestationsRegistry *FlashtestationsRegistryFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *FlashtestationsRegistryEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _FlashtestationsRegistry.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlashtestationsRegistryEIP712DomainChanged)
				if err := _FlashtestationsRegistry.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_FlashtestationsRegistry *FlashtestationsRegistryFilterer) ParseEIP712DomainChanged(log types.Log) (*FlashtestationsRegistryEIP712DomainChanged, error) {
	event := new(FlashtestationsRegistryEIP712DomainChanged)
	if err := _FlashtestationsRegistry.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FlashtestationsRegistryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the FlashtestationsRegistry contract.
type FlashtestationsRegistryInitializedIterator struct {
	Event *FlashtestationsRegistryInitialized // Event containing the contract specifics and raw log

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
func (it *FlashtestationsRegistryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlashtestationsRegistryInitialized)
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
		it.Event = new(FlashtestationsRegistryInitialized)
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
func (it *FlashtestationsRegistryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlashtestationsRegistryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlashtestationsRegistryInitialized represents a Initialized event raised by the FlashtestationsRegistry contract.
type FlashtestationsRegistryInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_FlashtestationsRegistry *FlashtestationsRegistryFilterer) FilterInitialized(opts *bind.FilterOpts) (*FlashtestationsRegistryInitializedIterator, error) {

	logs, sub, err := _FlashtestationsRegistry.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &FlashtestationsRegistryInitializedIterator{contract: _FlashtestationsRegistry.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_FlashtestationsRegistry *FlashtestationsRegistryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *FlashtestationsRegistryInitialized) (event.Subscription, error) {

	logs, sub, err := _FlashtestationsRegistry.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlashtestationsRegistryInitialized)
				if err := _FlashtestationsRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_FlashtestationsRegistry *FlashtestationsRegistryFilterer) ParseInitialized(log types.Log) (*FlashtestationsRegistryInitialized, error) {
	event := new(FlashtestationsRegistryInitialized)
	if err := _FlashtestationsRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FlashtestationsRegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the FlashtestationsRegistry contract.
type FlashtestationsRegistryOwnershipTransferredIterator struct {
	Event *FlashtestationsRegistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FlashtestationsRegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlashtestationsRegistryOwnershipTransferred)
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
		it.Event = new(FlashtestationsRegistryOwnershipTransferred)
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
func (it *FlashtestationsRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlashtestationsRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlashtestationsRegistryOwnershipTransferred represents a OwnershipTransferred event raised by the FlashtestationsRegistry contract.
type FlashtestationsRegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FlashtestationsRegistry *FlashtestationsRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FlashtestationsRegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FlashtestationsRegistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FlashtestationsRegistryOwnershipTransferredIterator{contract: _FlashtestationsRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FlashtestationsRegistry *FlashtestationsRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FlashtestationsRegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FlashtestationsRegistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlashtestationsRegistryOwnershipTransferred)
				if err := _FlashtestationsRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FlashtestationsRegistry *FlashtestationsRegistryFilterer) ParseOwnershipTransferred(log types.Log) (*FlashtestationsRegistryOwnershipTransferred, error) {
	event := new(FlashtestationsRegistryOwnershipTransferred)
	if err := _FlashtestationsRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FlashtestationsRegistryPreviousSignatureInvalidatedIterator is returned from FilterPreviousSignatureInvalidated and is used to iterate over the raw logs and unpacked data for PreviousSignatureInvalidated events raised by the FlashtestationsRegistry contract.
type FlashtestationsRegistryPreviousSignatureInvalidatedIterator struct {
	Event *FlashtestationsRegistryPreviousSignatureInvalidated // Event containing the contract specifics and raw log

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
func (it *FlashtestationsRegistryPreviousSignatureInvalidatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlashtestationsRegistryPreviousSignatureInvalidated)
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
		it.Event = new(FlashtestationsRegistryPreviousSignatureInvalidated)
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
func (it *FlashtestationsRegistryPreviousSignatureInvalidatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlashtestationsRegistryPreviousSignatureInvalidatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlashtestationsRegistryPreviousSignatureInvalidated represents a PreviousSignatureInvalidated event raised by the FlashtestationsRegistry contract.
type FlashtestationsRegistryPreviousSignatureInvalidated struct {
	TeeAddress       common.Address
	InvalidatedNonce *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterPreviousSignatureInvalidated is a free log retrieval operation binding the contract event 0xaba960b001cf41ae7d1278e08bf0afa5081bfad043326cfe1e1d5ee266c9ac52.
//
// Solidity: event PreviousSignatureInvalidated(address indexed teeAddress, uint256 invalidatedNonce)
func (_FlashtestationsRegistry *FlashtestationsRegistryFilterer) FilterPreviousSignatureInvalidated(opts *bind.FilterOpts, teeAddress []common.Address) (*FlashtestationsRegistryPreviousSignatureInvalidatedIterator, error) {

	var teeAddressRule []interface{}
	for _, teeAddressItem := range teeAddress {
		teeAddressRule = append(teeAddressRule, teeAddressItem)
	}

	logs, sub, err := _FlashtestationsRegistry.contract.FilterLogs(opts, "PreviousSignatureInvalidated", teeAddressRule)
	if err != nil {
		return nil, err
	}
	return &FlashtestationsRegistryPreviousSignatureInvalidatedIterator{contract: _FlashtestationsRegistry.contract, event: "PreviousSignatureInvalidated", logs: logs, sub: sub}, nil
}

// WatchPreviousSignatureInvalidated is a free log subscription operation binding the contract event 0xaba960b001cf41ae7d1278e08bf0afa5081bfad043326cfe1e1d5ee266c9ac52.
//
// Solidity: event PreviousSignatureInvalidated(address indexed teeAddress, uint256 invalidatedNonce)
func (_FlashtestationsRegistry *FlashtestationsRegistryFilterer) WatchPreviousSignatureInvalidated(opts *bind.WatchOpts, sink chan<- *FlashtestationsRegistryPreviousSignatureInvalidated, teeAddress []common.Address) (event.Subscription, error) {

	var teeAddressRule []interface{}
	for _, teeAddressItem := range teeAddress {
		teeAddressRule = append(teeAddressRule, teeAddressItem)
	}

	logs, sub, err := _FlashtestationsRegistry.contract.WatchLogs(opts, "PreviousSignatureInvalidated", teeAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlashtestationsRegistryPreviousSignatureInvalidated)
				if err := _FlashtestationsRegistry.contract.UnpackLog(event, "PreviousSignatureInvalidated", log); err != nil {
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

// ParsePreviousSignatureInvalidated is a log parse operation binding the contract event 0xaba960b001cf41ae7d1278e08bf0afa5081bfad043326cfe1e1d5ee266c9ac52.
//
// Solidity: event PreviousSignatureInvalidated(address indexed teeAddress, uint256 invalidatedNonce)
func (_FlashtestationsRegistry *FlashtestationsRegistryFilterer) ParsePreviousSignatureInvalidated(log types.Log) (*FlashtestationsRegistryPreviousSignatureInvalidated, error) {
	event := new(FlashtestationsRegistryPreviousSignatureInvalidated)
	if err := _FlashtestationsRegistry.contract.UnpackLog(event, "PreviousSignatureInvalidated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FlashtestationsRegistryTEEServiceInvalidatedIterator is returned from FilterTEEServiceInvalidated and is used to iterate over the raw logs and unpacked data for TEEServiceInvalidated events raised by the FlashtestationsRegistry contract.
type FlashtestationsRegistryTEEServiceInvalidatedIterator struct {
	Event *FlashtestationsRegistryTEEServiceInvalidated // Event containing the contract specifics and raw log

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
func (it *FlashtestationsRegistryTEEServiceInvalidatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlashtestationsRegistryTEEServiceInvalidated)
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
		it.Event = new(FlashtestationsRegistryTEEServiceInvalidated)
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
func (it *FlashtestationsRegistryTEEServiceInvalidatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlashtestationsRegistryTEEServiceInvalidatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlashtestationsRegistryTEEServiceInvalidated represents a TEEServiceInvalidated event raised by the FlashtestationsRegistry contract.
type FlashtestationsRegistryTEEServiceInvalidated struct {
	TeeAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTEEServiceInvalidated is a free log retrieval operation binding the contract event 0x5bb0bbb0993a623e10dd3579bf5b9403deba943e0bfe950b740d60209c9135ef.
//
// Solidity: event TEEServiceInvalidated(address indexed teeAddress)
func (_FlashtestationsRegistry *FlashtestationsRegistryFilterer) FilterTEEServiceInvalidated(opts *bind.FilterOpts, teeAddress []common.Address) (*FlashtestationsRegistryTEEServiceInvalidatedIterator, error) {

	var teeAddressRule []interface{}
	for _, teeAddressItem := range teeAddress {
		teeAddressRule = append(teeAddressRule, teeAddressItem)
	}

	logs, sub, err := _FlashtestationsRegistry.contract.FilterLogs(opts, "TEEServiceInvalidated", teeAddressRule)
	if err != nil {
		return nil, err
	}
	return &FlashtestationsRegistryTEEServiceInvalidatedIterator{contract: _FlashtestationsRegistry.contract, event: "TEEServiceInvalidated", logs: logs, sub: sub}, nil
}

// WatchTEEServiceInvalidated is a free log subscription operation binding the contract event 0x5bb0bbb0993a623e10dd3579bf5b9403deba943e0bfe950b740d60209c9135ef.
//
// Solidity: event TEEServiceInvalidated(address indexed teeAddress)
func (_FlashtestationsRegistry *FlashtestationsRegistryFilterer) WatchTEEServiceInvalidated(opts *bind.WatchOpts, sink chan<- *FlashtestationsRegistryTEEServiceInvalidated, teeAddress []common.Address) (event.Subscription, error) {

	var teeAddressRule []interface{}
	for _, teeAddressItem := range teeAddress {
		teeAddressRule = append(teeAddressRule, teeAddressItem)
	}

	logs, sub, err := _FlashtestationsRegistry.contract.WatchLogs(opts, "TEEServiceInvalidated", teeAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlashtestationsRegistryTEEServiceInvalidated)
				if err := _FlashtestationsRegistry.contract.UnpackLog(event, "TEEServiceInvalidated", log); err != nil {
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

// ParseTEEServiceInvalidated is a log parse operation binding the contract event 0x5bb0bbb0993a623e10dd3579bf5b9403deba943e0bfe950b740d60209c9135ef.
//
// Solidity: event TEEServiceInvalidated(address indexed teeAddress)
func (_FlashtestationsRegistry *FlashtestationsRegistryFilterer) ParseTEEServiceInvalidated(log types.Log) (*FlashtestationsRegistryTEEServiceInvalidated, error) {
	event := new(FlashtestationsRegistryTEEServiceInvalidated)
	if err := _FlashtestationsRegistry.contract.UnpackLog(event, "TEEServiceInvalidated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FlashtestationsRegistryTEEServiceRegisteredIterator is returned from FilterTEEServiceRegistered and is used to iterate over the raw logs and unpacked data for TEEServiceRegistered events raised by the FlashtestationsRegistry contract.
type FlashtestationsRegistryTEEServiceRegisteredIterator struct {
	Event *FlashtestationsRegistryTEEServiceRegistered // Event containing the contract specifics and raw log

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
func (it *FlashtestationsRegistryTEEServiceRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlashtestationsRegistryTEEServiceRegistered)
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
		it.Event = new(FlashtestationsRegistryTEEServiceRegistered)
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
func (it *FlashtestationsRegistryTEEServiceRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlashtestationsRegistryTEEServiceRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlashtestationsRegistryTEEServiceRegistered represents a TEEServiceRegistered event raised by the FlashtestationsRegistry contract.
type FlashtestationsRegistryTEEServiceRegistered struct {
	TeeAddress    common.Address
	RawQuote      []byte
	AlreadyExists bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTEEServiceRegistered is a free log retrieval operation binding the contract event 0x206fdb1a74851a8542447b8b6704db24a36b906a7297cc23c2b984dc357b9978.
//
// Solidity: event TEEServiceRegistered(address indexed teeAddress, bytes rawQuote, bool alreadyExists)
func (_FlashtestationsRegistry *FlashtestationsRegistryFilterer) FilterTEEServiceRegistered(opts *bind.FilterOpts, teeAddress []common.Address) (*FlashtestationsRegistryTEEServiceRegisteredIterator, error) {

	var teeAddressRule []interface{}
	for _, teeAddressItem := range teeAddress {
		teeAddressRule = append(teeAddressRule, teeAddressItem)
	}

	logs, sub, err := _FlashtestationsRegistry.contract.FilterLogs(opts, "TEEServiceRegistered", teeAddressRule)
	if err != nil {
		return nil, err
	}
	return &FlashtestationsRegistryTEEServiceRegisteredIterator{contract: _FlashtestationsRegistry.contract, event: "TEEServiceRegistered", logs: logs, sub: sub}, nil
}

// WatchTEEServiceRegistered is a free log subscription operation binding the contract event 0x206fdb1a74851a8542447b8b6704db24a36b906a7297cc23c2b984dc357b9978.
//
// Solidity: event TEEServiceRegistered(address indexed teeAddress, bytes rawQuote, bool alreadyExists)
func (_FlashtestationsRegistry *FlashtestationsRegistryFilterer) WatchTEEServiceRegistered(opts *bind.WatchOpts, sink chan<- *FlashtestationsRegistryTEEServiceRegistered, teeAddress []common.Address) (event.Subscription, error) {

	var teeAddressRule []interface{}
	for _, teeAddressItem := range teeAddress {
		teeAddressRule = append(teeAddressRule, teeAddressItem)
	}

	logs, sub, err := _FlashtestationsRegistry.contract.WatchLogs(opts, "TEEServiceRegistered", teeAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlashtestationsRegistryTEEServiceRegistered)
				if err := _FlashtestationsRegistry.contract.UnpackLog(event, "TEEServiceRegistered", log); err != nil {
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

// ParseTEEServiceRegistered is a log parse operation binding the contract event 0x206fdb1a74851a8542447b8b6704db24a36b906a7297cc23c2b984dc357b9978.
//
// Solidity: event TEEServiceRegistered(address indexed teeAddress, bytes rawQuote, bool alreadyExists)
func (_FlashtestationsRegistry *FlashtestationsRegistryFilterer) ParseTEEServiceRegistered(log types.Log) (*FlashtestationsRegistryTEEServiceRegistered, error) {
	event := new(FlashtestationsRegistryTEEServiceRegistered)
	if err := _FlashtestationsRegistry.contract.UnpackLog(event, "TEEServiceRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FlashtestationsRegistryUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the FlashtestationsRegistry contract.
type FlashtestationsRegistryUpgradedIterator struct {
	Event *FlashtestationsRegistryUpgraded // Event containing the contract specifics and raw log

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
func (it *FlashtestationsRegistryUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlashtestationsRegistryUpgraded)
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
		it.Event = new(FlashtestationsRegistryUpgraded)
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
func (it *FlashtestationsRegistryUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlashtestationsRegistryUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlashtestationsRegistryUpgraded represents a Upgraded event raised by the FlashtestationsRegistry contract.
type FlashtestationsRegistryUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_FlashtestationsRegistry *FlashtestationsRegistryFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*FlashtestationsRegistryUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _FlashtestationsRegistry.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &FlashtestationsRegistryUpgradedIterator{contract: _FlashtestationsRegistry.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_FlashtestationsRegistry *FlashtestationsRegistryFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *FlashtestationsRegistryUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _FlashtestationsRegistry.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlashtestationsRegistryUpgraded)
				if err := _FlashtestationsRegistry.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_FlashtestationsRegistry *FlashtestationsRegistryFilterer) ParseUpgraded(log types.Log) (*FlashtestationsRegistryUpgraded, error) {
	event := new(FlashtestationsRegistryUpgraded)
	if err := _FlashtestationsRegistry.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
