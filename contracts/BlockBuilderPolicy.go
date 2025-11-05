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

// IBlockBuilderPolicyWorkloadMetadata is an auto generated low-level Go binding around an user-defined struct.
type IBlockBuilderPolicyWorkloadMetadata struct {
	CommitHash     string
	SourceLocators []string
}

// TD10ReportBody is an auto generated low-level Go binding around an user-defined struct.
type TD10ReportBody struct {
	TeeTcbSvn      [16]byte
	MrSeam         []byte
	MrsignerSeam   []byte
	SeamAttributes [8]byte
	TdAttributes   [8]byte
	XFAM           [8]byte
	MrTd           []byte
	MrConfigId     []byte
	MrOwner        []byte
	MrOwnerConfig  []byte
	RtMr0          []byte
	RtMr1          []byte
	RtMr2          []byte
	RtMr3          []byte
	ReportData     []byte
}

// BlockBuilderPolicyMetaData contains all meta data concerning the BlockBuilderPolicy contract.
var BlockBuilderPolicyMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"VERIFY_BLOCK_BUILDER_PROOF_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addWorkloadToPolicy\",\"inputs\":[{\"name\":\"workloadId\",\"type\":\"bytes32\",\"internalType\":\"WorkloadId\"},{\"name\":\"commitHash\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"sourceLocators\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"computeStructHash\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"blockContentHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"domainSeparator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"eip712Domain\",\"inputs\":[],\"outputs\":[{\"name\":\"fields\",\"type\":\"bytes1\",\"internalType\":\"bytes1\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"verifyingContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"extensions\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getHashedTypeDataV4\",\"inputs\":[{\"name\":\"structHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWorkloadMetadata\",\"inputs\":[{\"name\":\"workloadId\",\"type\":\"bytes32\",\"internalType\":\"WorkloadId\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIBlockBuilderPolicy.WorkloadMetadata\",\"components\":[{\"name\":\"commitHash\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"sourceLocators\",\"type\":\"string[]\",\"internalType\":\"string[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_registry\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isAllowedPolicy\",\"inputs\":[{\"name\":\"teeAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"allowed\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"WorkloadId\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nonces\",\"inputs\":[{\"name\":\"teeAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"permitNonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"permitVerifyBlockBuilderProof\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"blockContentHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"eip712Sig\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeWorkloadFromPolicy\",\"inputs\":[{\"name\":\"workloadId\",\"type\":\"bytes32\",\"internalType\":\"WorkloadId\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"verifyBlockBuilderProof\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"blockContentHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"workloadIdForTDRegistration\",\"inputs\":[{\"name\":\"registration\",\"type\":\"tuple\",\"internalType\":\"structIFlashtestationRegistry.RegisteredTEE\",\"components\":[{\"name\":\"isValid\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"rawQuote\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"parsedReportBody\",\"type\":\"tuple\",\"internalType\":\"structTD10ReportBody\",\"components\":[{\"name\":\"teeTcbSvn\",\"type\":\"bytes16\",\"internalType\":\"bytes16\"},{\"name\":\"mrSeam\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"mrsignerSeam\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"seamAttributes\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"},{\"name\":\"tdAttributes\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"},{\"name\":\"xFAM\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"},{\"name\":\"mrTd\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"mrConfigId\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"mrOwner\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"mrOwnerConfig\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"rtMr0\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"rtMr1\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"rtMr2\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"rtMr3\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"reportData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"extendedRegistrationData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quoteHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"WorkloadId\"}],\"stateMutability\":\"pure\"},{\"type\":\"event\",\"name\":\"BlockBuilderProofVerified\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"workloadId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"blockContentHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"commitHash\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EIP712DomainChanged\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RegistrySet\",\"inputs\":[{\"name\":\"registry\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WorkloadAddedToPolicy\",\"inputs\":[{\"name\":\"workloadId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WorkloadRemovedFromPolicy\",\"inputs\":[{\"name\":\"workloadId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureS\",\"inputs\":[{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EmptyCommitHash\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EmptySourceLocators\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidNonce\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidRegistry\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"UnauthorizedBlockBuilder\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"WorkloadAlreadyInPolicy\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WorkloadNotInPolicy\",\"inputs\":[]}]",
}

// BlockBuilderPolicyABI is the input ABI used to generate the binding from.
// Deprecated: Use BlockBuilderPolicyMetaData.ABI instead.
var BlockBuilderPolicyABI = BlockBuilderPolicyMetaData.ABI

// BlockBuilderPolicy is an auto generated Go binding around an Ethereum contract.
type BlockBuilderPolicy struct {
	BlockBuilderPolicyCaller     // Read-only binding to the contract
	BlockBuilderPolicyTransactor // Write-only binding to the contract
	BlockBuilderPolicyFilterer   // Log filterer for contract events
}

// BlockBuilderPolicyCaller is an auto generated read-only Go binding around an Ethereum contract.
type BlockBuilderPolicyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockBuilderPolicyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BlockBuilderPolicyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockBuilderPolicyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BlockBuilderPolicyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockBuilderPolicySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BlockBuilderPolicySession struct {
	Contract     *BlockBuilderPolicy // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BlockBuilderPolicyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BlockBuilderPolicyCallerSession struct {
	Contract *BlockBuilderPolicyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// BlockBuilderPolicyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BlockBuilderPolicyTransactorSession struct {
	Contract     *BlockBuilderPolicyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// BlockBuilderPolicyRaw is an auto generated low-level Go binding around an Ethereum contract.
type BlockBuilderPolicyRaw struct {
	Contract *BlockBuilderPolicy // Generic contract binding to access the raw methods on
}

// BlockBuilderPolicyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BlockBuilderPolicyCallerRaw struct {
	Contract *BlockBuilderPolicyCaller // Generic read-only contract binding to access the raw methods on
}

// BlockBuilderPolicyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BlockBuilderPolicyTransactorRaw struct {
	Contract *BlockBuilderPolicyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBlockBuilderPolicy creates a new instance of BlockBuilderPolicy, bound to a specific deployed contract.
func NewBlockBuilderPolicy(address common.Address, backend bind.ContractBackend) (*BlockBuilderPolicy, error) {
	contract, err := bindBlockBuilderPolicy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BlockBuilderPolicy{BlockBuilderPolicyCaller: BlockBuilderPolicyCaller{contract: contract}, BlockBuilderPolicyTransactor: BlockBuilderPolicyTransactor{contract: contract}, BlockBuilderPolicyFilterer: BlockBuilderPolicyFilterer{contract: contract}}, nil
}

// NewBlockBuilderPolicyCaller creates a new read-only instance of BlockBuilderPolicy, bound to a specific deployed contract.
func NewBlockBuilderPolicyCaller(address common.Address, caller bind.ContractCaller) (*BlockBuilderPolicyCaller, error) {
	contract, err := bindBlockBuilderPolicy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BlockBuilderPolicyCaller{contract: contract}, nil
}

// NewBlockBuilderPolicyTransactor creates a new write-only instance of BlockBuilderPolicy, bound to a specific deployed contract.
func NewBlockBuilderPolicyTransactor(address common.Address, transactor bind.ContractTransactor) (*BlockBuilderPolicyTransactor, error) {
	contract, err := bindBlockBuilderPolicy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BlockBuilderPolicyTransactor{contract: contract}, nil
}

// NewBlockBuilderPolicyFilterer creates a new log filterer instance of BlockBuilderPolicy, bound to a specific deployed contract.
func NewBlockBuilderPolicyFilterer(address common.Address, filterer bind.ContractFilterer) (*BlockBuilderPolicyFilterer, error) {
	contract, err := bindBlockBuilderPolicy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BlockBuilderPolicyFilterer{contract: contract}, nil
}

// bindBlockBuilderPolicy binds a generic wrapper to an already deployed contract.
func bindBlockBuilderPolicy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BlockBuilderPolicyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BlockBuilderPolicy *BlockBuilderPolicyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BlockBuilderPolicy.Contract.BlockBuilderPolicyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BlockBuilderPolicy *BlockBuilderPolicyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlockBuilderPolicy.Contract.BlockBuilderPolicyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BlockBuilderPolicy *BlockBuilderPolicyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BlockBuilderPolicy.Contract.BlockBuilderPolicyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BlockBuilderPolicy *BlockBuilderPolicyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BlockBuilderPolicy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BlockBuilderPolicy *BlockBuilderPolicyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlockBuilderPolicy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BlockBuilderPolicy *BlockBuilderPolicyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BlockBuilderPolicy.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_BlockBuilderPolicy *BlockBuilderPolicyCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BlockBuilderPolicy.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_BlockBuilderPolicy *BlockBuilderPolicySession) UPGRADEINTERFACEVERSION() (string, error) {
	return _BlockBuilderPolicy.Contract.UPGRADEINTERFACEVERSION(&_BlockBuilderPolicy.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_BlockBuilderPolicy *BlockBuilderPolicyCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _BlockBuilderPolicy.Contract.UPGRADEINTERFACEVERSION(&_BlockBuilderPolicy.CallOpts)
}

// VERIFYBLOCKBUILDERPROOFTYPEHASH is a free data retrieval call binding the contract method 0x73016923.
//
// Solidity: function VERIFY_BLOCK_BUILDER_PROOF_TYPEHASH() view returns(bytes32)
func (_BlockBuilderPolicy *BlockBuilderPolicyCaller) VERIFYBLOCKBUILDERPROOFTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BlockBuilderPolicy.contract.Call(opts, &out, "VERIFY_BLOCK_BUILDER_PROOF_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VERIFYBLOCKBUILDERPROOFTYPEHASH is a free data retrieval call binding the contract method 0x73016923.
//
// Solidity: function VERIFY_BLOCK_BUILDER_PROOF_TYPEHASH() view returns(bytes32)
func (_BlockBuilderPolicy *BlockBuilderPolicySession) VERIFYBLOCKBUILDERPROOFTYPEHASH() ([32]byte, error) {
	return _BlockBuilderPolicy.Contract.VERIFYBLOCKBUILDERPROOFTYPEHASH(&_BlockBuilderPolicy.CallOpts)
}

// VERIFYBLOCKBUILDERPROOFTYPEHASH is a free data retrieval call binding the contract method 0x73016923.
//
// Solidity: function VERIFY_BLOCK_BUILDER_PROOF_TYPEHASH() view returns(bytes32)
func (_BlockBuilderPolicy *BlockBuilderPolicyCallerSession) VERIFYBLOCKBUILDERPROOFTYPEHASH() ([32]byte, error) {
	return _BlockBuilderPolicy.Contract.VERIFYBLOCKBUILDERPROOFTYPEHASH(&_BlockBuilderPolicy.CallOpts)
}

// ComputeStructHash is a free data retrieval call binding the contract method 0x7dec71a9.
//
// Solidity: function computeStructHash(uint8 version, bytes32 blockContentHash, uint256 nonce) pure returns(bytes32)
func (_BlockBuilderPolicy *BlockBuilderPolicyCaller) ComputeStructHash(opts *bind.CallOpts, version uint8, blockContentHash [32]byte, nonce *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _BlockBuilderPolicy.contract.Call(opts, &out, "computeStructHash", version, blockContentHash, nonce)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ComputeStructHash is a free data retrieval call binding the contract method 0x7dec71a9.
//
// Solidity: function computeStructHash(uint8 version, bytes32 blockContentHash, uint256 nonce) pure returns(bytes32)
func (_BlockBuilderPolicy *BlockBuilderPolicySession) ComputeStructHash(version uint8, blockContentHash [32]byte, nonce *big.Int) ([32]byte, error) {
	return _BlockBuilderPolicy.Contract.ComputeStructHash(&_BlockBuilderPolicy.CallOpts, version, blockContentHash, nonce)
}

// ComputeStructHash is a free data retrieval call binding the contract method 0x7dec71a9.
//
// Solidity: function computeStructHash(uint8 version, bytes32 blockContentHash, uint256 nonce) pure returns(bytes32)
func (_BlockBuilderPolicy *BlockBuilderPolicyCallerSession) ComputeStructHash(version uint8, blockContentHash [32]byte, nonce *big.Int) ([32]byte, error) {
	return _BlockBuilderPolicy.Contract.ComputeStructHash(&_BlockBuilderPolicy.CallOpts, version, blockContentHash, nonce)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_BlockBuilderPolicy *BlockBuilderPolicyCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BlockBuilderPolicy.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_BlockBuilderPolicy *BlockBuilderPolicySession) DomainSeparator() ([32]byte, error) {
	return _BlockBuilderPolicy.Contract.DomainSeparator(&_BlockBuilderPolicy.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_BlockBuilderPolicy *BlockBuilderPolicyCallerSession) DomainSeparator() ([32]byte, error) {
	return _BlockBuilderPolicy.Contract.DomainSeparator(&_BlockBuilderPolicy.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_BlockBuilderPolicy *BlockBuilderPolicyCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _BlockBuilderPolicy.contract.Call(opts, &out, "eip712Domain")

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
func (_BlockBuilderPolicy *BlockBuilderPolicySession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _BlockBuilderPolicy.Contract.Eip712Domain(&_BlockBuilderPolicy.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_BlockBuilderPolicy *BlockBuilderPolicyCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _BlockBuilderPolicy.Contract.Eip712Domain(&_BlockBuilderPolicy.CallOpts)
}

// GetHashedTypeDataV4 is a free data retrieval call binding the contract method 0x6931164e.
//
// Solidity: function getHashedTypeDataV4(bytes32 structHash) view returns(bytes32)
func (_BlockBuilderPolicy *BlockBuilderPolicyCaller) GetHashedTypeDataV4(opts *bind.CallOpts, structHash [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _BlockBuilderPolicy.contract.Call(opts, &out, "getHashedTypeDataV4", structHash)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetHashedTypeDataV4 is a free data retrieval call binding the contract method 0x6931164e.
//
// Solidity: function getHashedTypeDataV4(bytes32 structHash) view returns(bytes32)
func (_BlockBuilderPolicy *BlockBuilderPolicySession) GetHashedTypeDataV4(structHash [32]byte) ([32]byte, error) {
	return _BlockBuilderPolicy.Contract.GetHashedTypeDataV4(&_BlockBuilderPolicy.CallOpts, structHash)
}

// GetHashedTypeDataV4 is a free data retrieval call binding the contract method 0x6931164e.
//
// Solidity: function getHashedTypeDataV4(bytes32 structHash) view returns(bytes32)
func (_BlockBuilderPolicy *BlockBuilderPolicyCallerSession) GetHashedTypeDataV4(structHash [32]byte) ([32]byte, error) {
	return _BlockBuilderPolicy.Contract.GetHashedTypeDataV4(&_BlockBuilderPolicy.CallOpts, structHash)
}

// GetWorkloadMetadata is a free data retrieval call binding the contract method 0xabd45d21.
//
// Solidity: function getWorkloadMetadata(bytes32 workloadId) view returns((string,string[]))
func (_BlockBuilderPolicy *BlockBuilderPolicyCaller) GetWorkloadMetadata(opts *bind.CallOpts, workloadId [32]byte) (IBlockBuilderPolicyWorkloadMetadata, error) {
	var out []interface{}
	err := _BlockBuilderPolicy.contract.Call(opts, &out, "getWorkloadMetadata", workloadId)

	if err != nil {
		return *new(IBlockBuilderPolicyWorkloadMetadata), err
	}

	out0 := *abi.ConvertType(out[0], new(IBlockBuilderPolicyWorkloadMetadata)).(*IBlockBuilderPolicyWorkloadMetadata)

	return out0, err

}

// GetWorkloadMetadata is a free data retrieval call binding the contract method 0xabd45d21.
//
// Solidity: function getWorkloadMetadata(bytes32 workloadId) view returns((string,string[]))
func (_BlockBuilderPolicy *BlockBuilderPolicySession) GetWorkloadMetadata(workloadId [32]byte) (IBlockBuilderPolicyWorkloadMetadata, error) {
	return _BlockBuilderPolicy.Contract.GetWorkloadMetadata(&_BlockBuilderPolicy.CallOpts, workloadId)
}

// GetWorkloadMetadata is a free data retrieval call binding the contract method 0xabd45d21.
//
// Solidity: function getWorkloadMetadata(bytes32 workloadId) view returns((string,string[]))
func (_BlockBuilderPolicy *BlockBuilderPolicyCallerSession) GetWorkloadMetadata(workloadId [32]byte) (IBlockBuilderPolicyWorkloadMetadata, error) {
	return _BlockBuilderPolicy.Contract.GetWorkloadMetadata(&_BlockBuilderPolicy.CallOpts, workloadId)
}

// IsAllowedPolicy is a free data retrieval call binding the contract method 0xd2753561.
//
// Solidity: function isAllowedPolicy(address teeAddress) view returns(bool allowed, bytes32)
func (_BlockBuilderPolicy *BlockBuilderPolicyCaller) IsAllowedPolicy(opts *bind.CallOpts, teeAddress common.Address) (bool, [32]byte, error) {
	var out []interface{}
	err := _BlockBuilderPolicy.contract.Call(opts, &out, "isAllowedPolicy", teeAddress)

	if err != nil {
		return *new(bool), *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return out0, out1, err

}

// IsAllowedPolicy is a free data retrieval call binding the contract method 0xd2753561.
//
// Solidity: function isAllowedPolicy(address teeAddress) view returns(bool allowed, bytes32)
func (_BlockBuilderPolicy *BlockBuilderPolicySession) IsAllowedPolicy(teeAddress common.Address) (bool, [32]byte, error) {
	return _BlockBuilderPolicy.Contract.IsAllowedPolicy(&_BlockBuilderPolicy.CallOpts, teeAddress)
}

// IsAllowedPolicy is a free data retrieval call binding the contract method 0xd2753561.
//
// Solidity: function isAllowedPolicy(address teeAddress) view returns(bool allowed, bytes32)
func (_BlockBuilderPolicy *BlockBuilderPolicyCallerSession) IsAllowedPolicy(teeAddress common.Address) (bool, [32]byte, error) {
	return _BlockBuilderPolicy.Contract.IsAllowedPolicy(&_BlockBuilderPolicy.CallOpts, teeAddress)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address teeAddress) view returns(uint256 permitNonce)
func (_BlockBuilderPolicy *BlockBuilderPolicyCaller) Nonces(opts *bind.CallOpts, teeAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BlockBuilderPolicy.contract.Call(opts, &out, "nonces", teeAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address teeAddress) view returns(uint256 permitNonce)
func (_BlockBuilderPolicy *BlockBuilderPolicySession) Nonces(teeAddress common.Address) (*big.Int, error) {
	return _BlockBuilderPolicy.Contract.Nonces(&_BlockBuilderPolicy.CallOpts, teeAddress)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address teeAddress) view returns(uint256 permitNonce)
func (_BlockBuilderPolicy *BlockBuilderPolicyCallerSession) Nonces(teeAddress common.Address) (*big.Int, error) {
	return _BlockBuilderPolicy.Contract.Nonces(&_BlockBuilderPolicy.CallOpts, teeAddress)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BlockBuilderPolicy *BlockBuilderPolicyCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BlockBuilderPolicy.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BlockBuilderPolicy *BlockBuilderPolicySession) Owner() (common.Address, error) {
	return _BlockBuilderPolicy.Contract.Owner(&_BlockBuilderPolicy.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BlockBuilderPolicy *BlockBuilderPolicyCallerSession) Owner() (common.Address, error) {
	return _BlockBuilderPolicy.Contract.Owner(&_BlockBuilderPolicy.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_BlockBuilderPolicy *BlockBuilderPolicyCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BlockBuilderPolicy.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_BlockBuilderPolicy *BlockBuilderPolicySession) ProxiableUUID() ([32]byte, error) {
	return _BlockBuilderPolicy.Contract.ProxiableUUID(&_BlockBuilderPolicy.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_BlockBuilderPolicy *BlockBuilderPolicyCallerSession) ProxiableUUID() ([32]byte, error) {
	return _BlockBuilderPolicy.Contract.ProxiableUUID(&_BlockBuilderPolicy.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_BlockBuilderPolicy *BlockBuilderPolicyCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BlockBuilderPolicy.contract.Call(opts, &out, "registry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_BlockBuilderPolicy *BlockBuilderPolicySession) Registry() (common.Address, error) {
	return _BlockBuilderPolicy.Contract.Registry(&_BlockBuilderPolicy.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_BlockBuilderPolicy *BlockBuilderPolicyCallerSession) Registry() (common.Address, error) {
	return _BlockBuilderPolicy.Contract.Registry(&_BlockBuilderPolicy.CallOpts)
}

// WorkloadIdForTDRegistration is a free data retrieval call binding the contract method 0x4d37fc7a.
//
// Solidity: function workloadIdForTDRegistration((bool,bytes,(bytes16,bytes,bytes,bytes8,bytes8,bytes8,bytes,bytes,bytes,bytes,bytes,bytes,bytes,bytes,bytes),bytes,bytes32) registration) pure returns(bytes32)
func (_BlockBuilderPolicy *BlockBuilderPolicyCaller) WorkloadIdForTDRegistration(opts *bind.CallOpts, registration IFlashtestationRegistryRegisteredTEE) ([32]byte, error) {
	var out []interface{}
	err := _BlockBuilderPolicy.contract.Call(opts, &out, "workloadIdForTDRegistration", registration)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WorkloadIdForTDRegistration is a free data retrieval call binding the contract method 0x4d37fc7a.
//
// Solidity: function workloadIdForTDRegistration((bool,bytes,(bytes16,bytes,bytes,bytes8,bytes8,bytes8,bytes,bytes,bytes,bytes,bytes,bytes,bytes,bytes,bytes),bytes,bytes32) registration) pure returns(bytes32)
func (_BlockBuilderPolicy *BlockBuilderPolicySession) WorkloadIdForTDRegistration(registration IFlashtestationRegistryRegisteredTEE) ([32]byte, error) {
	return _BlockBuilderPolicy.Contract.WorkloadIdForTDRegistration(&_BlockBuilderPolicy.CallOpts, registration)
}

// WorkloadIdForTDRegistration is a free data retrieval call binding the contract method 0x4d37fc7a.
//
// Solidity: function workloadIdForTDRegistration((bool,bytes,(bytes16,bytes,bytes,bytes8,bytes8,bytes8,bytes,bytes,bytes,bytes,bytes,bytes,bytes,bytes,bytes),bytes,bytes32) registration) pure returns(bytes32)
func (_BlockBuilderPolicy *BlockBuilderPolicyCallerSession) WorkloadIdForTDRegistration(registration IFlashtestationRegistryRegisteredTEE) ([32]byte, error) {
	return _BlockBuilderPolicy.Contract.WorkloadIdForTDRegistration(&_BlockBuilderPolicy.CallOpts, registration)
}

// AddWorkloadToPolicy is a paid mutator transaction binding the contract method 0x4f3a415a.
//
// Solidity: function addWorkloadToPolicy(bytes32 workloadId, string commitHash, string[] sourceLocators) returns()
func (_BlockBuilderPolicy *BlockBuilderPolicyTransactor) AddWorkloadToPolicy(opts *bind.TransactOpts, workloadId [32]byte, commitHash string, sourceLocators []string) (*types.Transaction, error) {
	return _BlockBuilderPolicy.contract.Transact(opts, "addWorkloadToPolicy", workloadId, commitHash, sourceLocators)
}

// AddWorkloadToPolicy is a paid mutator transaction binding the contract method 0x4f3a415a.
//
// Solidity: function addWorkloadToPolicy(bytes32 workloadId, string commitHash, string[] sourceLocators) returns()
func (_BlockBuilderPolicy *BlockBuilderPolicySession) AddWorkloadToPolicy(workloadId [32]byte, commitHash string, sourceLocators []string) (*types.Transaction, error) {
	return _BlockBuilderPolicy.Contract.AddWorkloadToPolicy(&_BlockBuilderPolicy.TransactOpts, workloadId, commitHash, sourceLocators)
}

// AddWorkloadToPolicy is a paid mutator transaction binding the contract method 0x4f3a415a.
//
// Solidity: function addWorkloadToPolicy(bytes32 workloadId, string commitHash, string[] sourceLocators) returns()
func (_BlockBuilderPolicy *BlockBuilderPolicyTransactorSession) AddWorkloadToPolicy(workloadId [32]byte, commitHash string, sourceLocators []string) (*types.Transaction, error) {
	return _BlockBuilderPolicy.Contract.AddWorkloadToPolicy(&_BlockBuilderPolicy.TransactOpts, workloadId, commitHash, sourceLocators)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _initialOwner, address _registry) returns()
func (_BlockBuilderPolicy *BlockBuilderPolicyTransactor) Initialize(opts *bind.TransactOpts, _initialOwner common.Address, _registry common.Address) (*types.Transaction, error) {
	return _BlockBuilderPolicy.contract.Transact(opts, "initialize", _initialOwner, _registry)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _initialOwner, address _registry) returns()
func (_BlockBuilderPolicy *BlockBuilderPolicySession) Initialize(_initialOwner common.Address, _registry common.Address) (*types.Transaction, error) {
	return _BlockBuilderPolicy.Contract.Initialize(&_BlockBuilderPolicy.TransactOpts, _initialOwner, _registry)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _initialOwner, address _registry) returns()
func (_BlockBuilderPolicy *BlockBuilderPolicyTransactorSession) Initialize(_initialOwner common.Address, _registry common.Address) (*types.Transaction, error) {
	return _BlockBuilderPolicy.Contract.Initialize(&_BlockBuilderPolicy.TransactOpts, _initialOwner, _registry)
}

// PermitVerifyBlockBuilderProof is a paid mutator transaction binding the contract method 0x2dd8abfe.
//
// Solidity: function permitVerifyBlockBuilderProof(uint8 version, bytes32 blockContentHash, uint256 nonce, bytes eip712Sig) returns()
func (_BlockBuilderPolicy *BlockBuilderPolicyTransactor) PermitVerifyBlockBuilderProof(opts *bind.TransactOpts, version uint8, blockContentHash [32]byte, nonce *big.Int, eip712Sig []byte) (*types.Transaction, error) {
	return _BlockBuilderPolicy.contract.Transact(opts, "permitVerifyBlockBuilderProof", version, blockContentHash, nonce, eip712Sig)
}

// PermitVerifyBlockBuilderProof is a paid mutator transaction binding the contract method 0x2dd8abfe.
//
// Solidity: function permitVerifyBlockBuilderProof(uint8 version, bytes32 blockContentHash, uint256 nonce, bytes eip712Sig) returns()
func (_BlockBuilderPolicy *BlockBuilderPolicySession) PermitVerifyBlockBuilderProof(version uint8, blockContentHash [32]byte, nonce *big.Int, eip712Sig []byte) (*types.Transaction, error) {
	return _BlockBuilderPolicy.Contract.PermitVerifyBlockBuilderProof(&_BlockBuilderPolicy.TransactOpts, version, blockContentHash, nonce, eip712Sig)
}

// PermitVerifyBlockBuilderProof is a paid mutator transaction binding the contract method 0x2dd8abfe.
//
// Solidity: function permitVerifyBlockBuilderProof(uint8 version, bytes32 blockContentHash, uint256 nonce, bytes eip712Sig) returns()
func (_BlockBuilderPolicy *BlockBuilderPolicyTransactorSession) PermitVerifyBlockBuilderProof(version uint8, blockContentHash [32]byte, nonce *big.Int, eip712Sig []byte) (*types.Transaction, error) {
	return _BlockBuilderPolicy.Contract.PermitVerifyBlockBuilderProof(&_BlockBuilderPolicy.TransactOpts, version, blockContentHash, nonce, eip712Sig)
}

// RemoveWorkloadFromPolicy is a paid mutator transaction binding the contract method 0x5c40e542.
//
// Solidity: function removeWorkloadFromPolicy(bytes32 workloadId) returns()
func (_BlockBuilderPolicy *BlockBuilderPolicyTransactor) RemoveWorkloadFromPolicy(opts *bind.TransactOpts, workloadId [32]byte) (*types.Transaction, error) {
	return _BlockBuilderPolicy.contract.Transact(opts, "removeWorkloadFromPolicy", workloadId)
}

// RemoveWorkloadFromPolicy is a paid mutator transaction binding the contract method 0x5c40e542.
//
// Solidity: function removeWorkloadFromPolicy(bytes32 workloadId) returns()
func (_BlockBuilderPolicy *BlockBuilderPolicySession) RemoveWorkloadFromPolicy(workloadId [32]byte) (*types.Transaction, error) {
	return _BlockBuilderPolicy.Contract.RemoveWorkloadFromPolicy(&_BlockBuilderPolicy.TransactOpts, workloadId)
}

// RemoveWorkloadFromPolicy is a paid mutator transaction binding the contract method 0x5c40e542.
//
// Solidity: function removeWorkloadFromPolicy(bytes32 workloadId) returns()
func (_BlockBuilderPolicy *BlockBuilderPolicyTransactorSession) RemoveWorkloadFromPolicy(workloadId [32]byte) (*types.Transaction, error) {
	return _BlockBuilderPolicy.Contract.RemoveWorkloadFromPolicy(&_BlockBuilderPolicy.TransactOpts, workloadId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BlockBuilderPolicy *BlockBuilderPolicyTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlockBuilderPolicy.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BlockBuilderPolicy *BlockBuilderPolicySession) RenounceOwnership() (*types.Transaction, error) {
	return _BlockBuilderPolicy.Contract.RenounceOwnership(&_BlockBuilderPolicy.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BlockBuilderPolicy *BlockBuilderPolicyTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BlockBuilderPolicy.Contract.RenounceOwnership(&_BlockBuilderPolicy.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BlockBuilderPolicy *BlockBuilderPolicyTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BlockBuilderPolicy.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BlockBuilderPolicy *BlockBuilderPolicySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BlockBuilderPolicy.Contract.TransferOwnership(&_BlockBuilderPolicy.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BlockBuilderPolicy *BlockBuilderPolicyTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BlockBuilderPolicy.Contract.TransferOwnership(&_BlockBuilderPolicy.TransactOpts, newOwner)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_BlockBuilderPolicy *BlockBuilderPolicyTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BlockBuilderPolicy.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_BlockBuilderPolicy *BlockBuilderPolicySession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BlockBuilderPolicy.Contract.UpgradeToAndCall(&_BlockBuilderPolicy.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_BlockBuilderPolicy *BlockBuilderPolicyTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BlockBuilderPolicy.Contract.UpgradeToAndCall(&_BlockBuilderPolicy.TransactOpts, newImplementation, data)
}

// VerifyBlockBuilderProof is a paid mutator transaction binding the contract method 0xb33d59da.
//
// Solidity: function verifyBlockBuilderProof(uint8 version, bytes32 blockContentHash) returns()
func (_BlockBuilderPolicy *BlockBuilderPolicyTransactor) VerifyBlockBuilderProof(opts *bind.TransactOpts, version uint8, blockContentHash [32]byte) (*types.Transaction, error) {
	return _BlockBuilderPolicy.contract.Transact(opts, "verifyBlockBuilderProof", version, blockContentHash)
}

// VerifyBlockBuilderProof is a paid mutator transaction binding the contract method 0xb33d59da.
//
// Solidity: function verifyBlockBuilderProof(uint8 version, bytes32 blockContentHash) returns()
func (_BlockBuilderPolicy *BlockBuilderPolicySession) VerifyBlockBuilderProof(version uint8, blockContentHash [32]byte) (*types.Transaction, error) {
	return _BlockBuilderPolicy.Contract.VerifyBlockBuilderProof(&_BlockBuilderPolicy.TransactOpts, version, blockContentHash)
}

// VerifyBlockBuilderProof is a paid mutator transaction binding the contract method 0xb33d59da.
//
// Solidity: function verifyBlockBuilderProof(uint8 version, bytes32 blockContentHash) returns()
func (_BlockBuilderPolicy *BlockBuilderPolicyTransactorSession) VerifyBlockBuilderProof(version uint8, blockContentHash [32]byte) (*types.Transaction, error) {
	return _BlockBuilderPolicy.Contract.VerifyBlockBuilderProof(&_BlockBuilderPolicy.TransactOpts, version, blockContentHash)
}

// BlockBuilderPolicyBlockBuilderProofVerifiedIterator is returned from FilterBlockBuilderProofVerified and is used to iterate over the raw logs and unpacked data for BlockBuilderProofVerified events raised by the BlockBuilderPolicy contract.
type BlockBuilderPolicyBlockBuilderProofVerifiedIterator struct {
	Event *BlockBuilderPolicyBlockBuilderProofVerified // Event containing the contract specifics and raw log

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
func (it *BlockBuilderPolicyBlockBuilderProofVerifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockBuilderPolicyBlockBuilderProofVerified)
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
		it.Event = new(BlockBuilderPolicyBlockBuilderProofVerified)
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
func (it *BlockBuilderPolicyBlockBuilderProofVerifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockBuilderPolicyBlockBuilderProofVerifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockBuilderPolicyBlockBuilderProofVerified represents a BlockBuilderProofVerified event raised by the BlockBuilderPolicy contract.
type BlockBuilderPolicyBlockBuilderProofVerified struct {
	Caller           common.Address
	WorkloadId       [32]byte
	Version          uint8
	BlockContentHash [32]byte
	CommitHash       string
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterBlockBuilderProofVerified is a free log retrieval operation binding the contract event 0x3fa039a23466a52e08acb25376ac7d81de184fa6549ffffb2fc920c47cb623ed.
//
// Solidity: event BlockBuilderProofVerified(address caller, bytes32 workloadId, uint8 version, bytes32 blockContentHash, string commitHash)
func (_BlockBuilderPolicy *BlockBuilderPolicyFilterer) FilterBlockBuilderProofVerified(opts *bind.FilterOpts) (*BlockBuilderPolicyBlockBuilderProofVerifiedIterator, error) {

	logs, sub, err := _BlockBuilderPolicy.contract.FilterLogs(opts, "BlockBuilderProofVerified")
	if err != nil {
		return nil, err
	}
	return &BlockBuilderPolicyBlockBuilderProofVerifiedIterator{contract: _BlockBuilderPolicy.contract, event: "BlockBuilderProofVerified", logs: logs, sub: sub}, nil
}

// WatchBlockBuilderProofVerified is a free log subscription operation binding the contract event 0x3fa039a23466a52e08acb25376ac7d81de184fa6549ffffb2fc920c47cb623ed.
//
// Solidity: event BlockBuilderProofVerified(address caller, bytes32 workloadId, uint8 version, bytes32 blockContentHash, string commitHash)
func (_BlockBuilderPolicy *BlockBuilderPolicyFilterer) WatchBlockBuilderProofVerified(opts *bind.WatchOpts, sink chan<- *BlockBuilderPolicyBlockBuilderProofVerified) (event.Subscription, error) {

	logs, sub, err := _BlockBuilderPolicy.contract.WatchLogs(opts, "BlockBuilderProofVerified")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockBuilderPolicyBlockBuilderProofVerified)
				if err := _BlockBuilderPolicy.contract.UnpackLog(event, "BlockBuilderProofVerified", log); err != nil {
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

// ParseBlockBuilderProofVerified is a log parse operation binding the contract event 0x3fa039a23466a52e08acb25376ac7d81de184fa6549ffffb2fc920c47cb623ed.
//
// Solidity: event BlockBuilderProofVerified(address caller, bytes32 workloadId, uint8 version, bytes32 blockContentHash, string commitHash)
func (_BlockBuilderPolicy *BlockBuilderPolicyFilterer) ParseBlockBuilderProofVerified(log types.Log) (*BlockBuilderPolicyBlockBuilderProofVerified, error) {
	event := new(BlockBuilderPolicyBlockBuilderProofVerified)
	if err := _BlockBuilderPolicy.contract.UnpackLog(event, "BlockBuilderProofVerified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlockBuilderPolicyEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the BlockBuilderPolicy contract.
type BlockBuilderPolicyEIP712DomainChangedIterator struct {
	Event *BlockBuilderPolicyEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *BlockBuilderPolicyEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockBuilderPolicyEIP712DomainChanged)
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
		it.Event = new(BlockBuilderPolicyEIP712DomainChanged)
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
func (it *BlockBuilderPolicyEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockBuilderPolicyEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockBuilderPolicyEIP712DomainChanged represents a EIP712DomainChanged event raised by the BlockBuilderPolicy contract.
type BlockBuilderPolicyEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_BlockBuilderPolicy *BlockBuilderPolicyFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*BlockBuilderPolicyEIP712DomainChangedIterator, error) {

	logs, sub, err := _BlockBuilderPolicy.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &BlockBuilderPolicyEIP712DomainChangedIterator{contract: _BlockBuilderPolicy.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_BlockBuilderPolicy *BlockBuilderPolicyFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *BlockBuilderPolicyEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _BlockBuilderPolicy.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockBuilderPolicyEIP712DomainChanged)
				if err := _BlockBuilderPolicy.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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
func (_BlockBuilderPolicy *BlockBuilderPolicyFilterer) ParseEIP712DomainChanged(log types.Log) (*BlockBuilderPolicyEIP712DomainChanged, error) {
	event := new(BlockBuilderPolicyEIP712DomainChanged)
	if err := _BlockBuilderPolicy.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlockBuilderPolicyInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the BlockBuilderPolicy contract.
type BlockBuilderPolicyInitializedIterator struct {
	Event *BlockBuilderPolicyInitialized // Event containing the contract specifics and raw log

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
func (it *BlockBuilderPolicyInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockBuilderPolicyInitialized)
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
		it.Event = new(BlockBuilderPolicyInitialized)
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
func (it *BlockBuilderPolicyInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockBuilderPolicyInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockBuilderPolicyInitialized represents a Initialized event raised by the BlockBuilderPolicy contract.
type BlockBuilderPolicyInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_BlockBuilderPolicy *BlockBuilderPolicyFilterer) FilterInitialized(opts *bind.FilterOpts) (*BlockBuilderPolicyInitializedIterator, error) {

	logs, sub, err := _BlockBuilderPolicy.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BlockBuilderPolicyInitializedIterator{contract: _BlockBuilderPolicy.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_BlockBuilderPolicy *BlockBuilderPolicyFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BlockBuilderPolicyInitialized) (event.Subscription, error) {

	logs, sub, err := _BlockBuilderPolicy.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockBuilderPolicyInitialized)
				if err := _BlockBuilderPolicy.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_BlockBuilderPolicy *BlockBuilderPolicyFilterer) ParseInitialized(log types.Log) (*BlockBuilderPolicyInitialized, error) {
	event := new(BlockBuilderPolicyInitialized)
	if err := _BlockBuilderPolicy.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlockBuilderPolicyOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BlockBuilderPolicy contract.
type BlockBuilderPolicyOwnershipTransferredIterator struct {
	Event *BlockBuilderPolicyOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BlockBuilderPolicyOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockBuilderPolicyOwnershipTransferred)
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
		it.Event = new(BlockBuilderPolicyOwnershipTransferred)
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
func (it *BlockBuilderPolicyOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockBuilderPolicyOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockBuilderPolicyOwnershipTransferred represents a OwnershipTransferred event raised by the BlockBuilderPolicy contract.
type BlockBuilderPolicyOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BlockBuilderPolicy *BlockBuilderPolicyFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BlockBuilderPolicyOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BlockBuilderPolicy.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BlockBuilderPolicyOwnershipTransferredIterator{contract: _BlockBuilderPolicy.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BlockBuilderPolicy *BlockBuilderPolicyFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BlockBuilderPolicyOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BlockBuilderPolicy.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockBuilderPolicyOwnershipTransferred)
				if err := _BlockBuilderPolicy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_BlockBuilderPolicy *BlockBuilderPolicyFilterer) ParseOwnershipTransferred(log types.Log) (*BlockBuilderPolicyOwnershipTransferred, error) {
	event := new(BlockBuilderPolicyOwnershipTransferred)
	if err := _BlockBuilderPolicy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlockBuilderPolicyRegistrySetIterator is returned from FilterRegistrySet and is used to iterate over the raw logs and unpacked data for RegistrySet events raised by the BlockBuilderPolicy contract.
type BlockBuilderPolicyRegistrySetIterator struct {
	Event *BlockBuilderPolicyRegistrySet // Event containing the contract specifics and raw log

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
func (it *BlockBuilderPolicyRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockBuilderPolicyRegistrySet)
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
		it.Event = new(BlockBuilderPolicyRegistrySet)
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
func (it *BlockBuilderPolicyRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockBuilderPolicyRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockBuilderPolicyRegistrySet represents a RegistrySet event raised by the BlockBuilderPolicy contract.
type BlockBuilderPolicyRegistrySet struct {
	Registry common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRegistrySet is a free log retrieval operation binding the contract event 0x27fe5f0c1c3b1ed427cc63d0f05759ffdecf9aec9e18d31ef366fc8a6cb5dc3b.
//
// Solidity: event RegistrySet(address indexed registry)
func (_BlockBuilderPolicy *BlockBuilderPolicyFilterer) FilterRegistrySet(opts *bind.FilterOpts, registry []common.Address) (*BlockBuilderPolicyRegistrySetIterator, error) {

	var registryRule []interface{}
	for _, registryItem := range registry {
		registryRule = append(registryRule, registryItem)
	}

	logs, sub, err := _BlockBuilderPolicy.contract.FilterLogs(opts, "RegistrySet", registryRule)
	if err != nil {
		return nil, err
	}
	return &BlockBuilderPolicyRegistrySetIterator{contract: _BlockBuilderPolicy.contract, event: "RegistrySet", logs: logs, sub: sub}, nil
}

// WatchRegistrySet is a free log subscription operation binding the contract event 0x27fe5f0c1c3b1ed427cc63d0f05759ffdecf9aec9e18d31ef366fc8a6cb5dc3b.
//
// Solidity: event RegistrySet(address indexed registry)
func (_BlockBuilderPolicy *BlockBuilderPolicyFilterer) WatchRegistrySet(opts *bind.WatchOpts, sink chan<- *BlockBuilderPolicyRegistrySet, registry []common.Address) (event.Subscription, error) {

	var registryRule []interface{}
	for _, registryItem := range registry {
		registryRule = append(registryRule, registryItem)
	}

	logs, sub, err := _BlockBuilderPolicy.contract.WatchLogs(opts, "RegistrySet", registryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockBuilderPolicyRegistrySet)
				if err := _BlockBuilderPolicy.contract.UnpackLog(event, "RegistrySet", log); err != nil {
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

// ParseRegistrySet is a log parse operation binding the contract event 0x27fe5f0c1c3b1ed427cc63d0f05759ffdecf9aec9e18d31ef366fc8a6cb5dc3b.
//
// Solidity: event RegistrySet(address indexed registry)
func (_BlockBuilderPolicy *BlockBuilderPolicyFilterer) ParseRegistrySet(log types.Log) (*BlockBuilderPolicyRegistrySet, error) {
	event := new(BlockBuilderPolicyRegistrySet)
	if err := _BlockBuilderPolicy.contract.UnpackLog(event, "RegistrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlockBuilderPolicyUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the BlockBuilderPolicy contract.
type BlockBuilderPolicyUpgradedIterator struct {
	Event *BlockBuilderPolicyUpgraded // Event containing the contract specifics and raw log

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
func (it *BlockBuilderPolicyUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockBuilderPolicyUpgraded)
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
		it.Event = new(BlockBuilderPolicyUpgraded)
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
func (it *BlockBuilderPolicyUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockBuilderPolicyUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockBuilderPolicyUpgraded represents a Upgraded event raised by the BlockBuilderPolicy contract.
type BlockBuilderPolicyUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_BlockBuilderPolicy *BlockBuilderPolicyFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*BlockBuilderPolicyUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _BlockBuilderPolicy.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &BlockBuilderPolicyUpgradedIterator{contract: _BlockBuilderPolicy.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_BlockBuilderPolicy *BlockBuilderPolicyFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *BlockBuilderPolicyUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _BlockBuilderPolicy.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockBuilderPolicyUpgraded)
				if err := _BlockBuilderPolicy.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_BlockBuilderPolicy *BlockBuilderPolicyFilterer) ParseUpgraded(log types.Log) (*BlockBuilderPolicyUpgraded, error) {
	event := new(BlockBuilderPolicyUpgraded)
	if err := _BlockBuilderPolicy.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlockBuilderPolicyWorkloadAddedToPolicyIterator is returned from FilterWorkloadAddedToPolicy and is used to iterate over the raw logs and unpacked data for WorkloadAddedToPolicy events raised by the BlockBuilderPolicy contract.
type BlockBuilderPolicyWorkloadAddedToPolicyIterator struct {
	Event *BlockBuilderPolicyWorkloadAddedToPolicy // Event containing the contract specifics and raw log

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
func (it *BlockBuilderPolicyWorkloadAddedToPolicyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockBuilderPolicyWorkloadAddedToPolicy)
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
		it.Event = new(BlockBuilderPolicyWorkloadAddedToPolicy)
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
func (it *BlockBuilderPolicyWorkloadAddedToPolicyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockBuilderPolicyWorkloadAddedToPolicyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockBuilderPolicyWorkloadAddedToPolicy represents a WorkloadAddedToPolicy event raised by the BlockBuilderPolicy contract.
type BlockBuilderPolicyWorkloadAddedToPolicy struct {
	WorkloadId [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterWorkloadAddedToPolicy is a free log retrieval operation binding the contract event 0xcbb92e241e191fed6d0b0da0a918c7dcf595e77d868e2e3bf9e6b0b91589c7ad.
//
// Solidity: event WorkloadAddedToPolicy(bytes32 indexed workloadId)
func (_BlockBuilderPolicy *BlockBuilderPolicyFilterer) FilterWorkloadAddedToPolicy(opts *bind.FilterOpts, workloadId [][32]byte) (*BlockBuilderPolicyWorkloadAddedToPolicyIterator, error) {

	var workloadIdRule []interface{}
	for _, workloadIdItem := range workloadId {
		workloadIdRule = append(workloadIdRule, workloadIdItem)
	}

	logs, sub, err := _BlockBuilderPolicy.contract.FilterLogs(opts, "WorkloadAddedToPolicy", workloadIdRule)
	if err != nil {
		return nil, err
	}
	return &BlockBuilderPolicyWorkloadAddedToPolicyIterator{contract: _BlockBuilderPolicy.contract, event: "WorkloadAddedToPolicy", logs: logs, sub: sub}, nil
}

// WatchWorkloadAddedToPolicy is a free log subscription operation binding the contract event 0xcbb92e241e191fed6d0b0da0a918c7dcf595e77d868e2e3bf9e6b0b91589c7ad.
//
// Solidity: event WorkloadAddedToPolicy(bytes32 indexed workloadId)
func (_BlockBuilderPolicy *BlockBuilderPolicyFilterer) WatchWorkloadAddedToPolicy(opts *bind.WatchOpts, sink chan<- *BlockBuilderPolicyWorkloadAddedToPolicy, workloadId [][32]byte) (event.Subscription, error) {

	var workloadIdRule []interface{}
	for _, workloadIdItem := range workloadId {
		workloadIdRule = append(workloadIdRule, workloadIdItem)
	}

	logs, sub, err := _BlockBuilderPolicy.contract.WatchLogs(opts, "WorkloadAddedToPolicy", workloadIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockBuilderPolicyWorkloadAddedToPolicy)
				if err := _BlockBuilderPolicy.contract.UnpackLog(event, "WorkloadAddedToPolicy", log); err != nil {
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

// ParseWorkloadAddedToPolicy is a log parse operation binding the contract event 0xcbb92e241e191fed6d0b0da0a918c7dcf595e77d868e2e3bf9e6b0b91589c7ad.
//
// Solidity: event WorkloadAddedToPolicy(bytes32 indexed workloadId)
func (_BlockBuilderPolicy *BlockBuilderPolicyFilterer) ParseWorkloadAddedToPolicy(log types.Log) (*BlockBuilderPolicyWorkloadAddedToPolicy, error) {
	event := new(BlockBuilderPolicyWorkloadAddedToPolicy)
	if err := _BlockBuilderPolicy.contract.UnpackLog(event, "WorkloadAddedToPolicy", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlockBuilderPolicyWorkloadRemovedFromPolicyIterator is returned from FilterWorkloadRemovedFromPolicy and is used to iterate over the raw logs and unpacked data for WorkloadRemovedFromPolicy events raised by the BlockBuilderPolicy contract.
type BlockBuilderPolicyWorkloadRemovedFromPolicyIterator struct {
	Event *BlockBuilderPolicyWorkloadRemovedFromPolicy // Event containing the contract specifics and raw log

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
func (it *BlockBuilderPolicyWorkloadRemovedFromPolicyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockBuilderPolicyWorkloadRemovedFromPolicy)
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
		it.Event = new(BlockBuilderPolicyWorkloadRemovedFromPolicy)
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
func (it *BlockBuilderPolicyWorkloadRemovedFromPolicyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockBuilderPolicyWorkloadRemovedFromPolicyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockBuilderPolicyWorkloadRemovedFromPolicy represents a WorkloadRemovedFromPolicy event raised by the BlockBuilderPolicy contract.
type BlockBuilderPolicyWorkloadRemovedFromPolicy struct {
	WorkloadId [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterWorkloadRemovedFromPolicy is a free log retrieval operation binding the contract event 0x56c387a9be1bf0e0e4f852c577a225db98e8253ad401d1b4ea73926f27d6af09.
//
// Solidity: event WorkloadRemovedFromPolicy(bytes32 indexed workloadId)
func (_BlockBuilderPolicy *BlockBuilderPolicyFilterer) FilterWorkloadRemovedFromPolicy(opts *bind.FilterOpts, workloadId [][32]byte) (*BlockBuilderPolicyWorkloadRemovedFromPolicyIterator, error) {

	var workloadIdRule []interface{}
	for _, workloadIdItem := range workloadId {
		workloadIdRule = append(workloadIdRule, workloadIdItem)
	}

	logs, sub, err := _BlockBuilderPolicy.contract.FilterLogs(opts, "WorkloadRemovedFromPolicy", workloadIdRule)
	if err != nil {
		return nil, err
	}
	return &BlockBuilderPolicyWorkloadRemovedFromPolicyIterator{contract: _BlockBuilderPolicy.contract, event: "WorkloadRemovedFromPolicy", logs: logs, sub: sub}, nil
}

// WatchWorkloadRemovedFromPolicy is a free log subscription operation binding the contract event 0x56c387a9be1bf0e0e4f852c577a225db98e8253ad401d1b4ea73926f27d6af09.
//
// Solidity: event WorkloadRemovedFromPolicy(bytes32 indexed workloadId)
func (_BlockBuilderPolicy *BlockBuilderPolicyFilterer) WatchWorkloadRemovedFromPolicy(opts *bind.WatchOpts, sink chan<- *BlockBuilderPolicyWorkloadRemovedFromPolicy, workloadId [][32]byte) (event.Subscription, error) {

	var workloadIdRule []interface{}
	for _, workloadIdItem := range workloadId {
		workloadIdRule = append(workloadIdRule, workloadIdItem)
	}

	logs, sub, err := _BlockBuilderPolicy.contract.WatchLogs(opts, "WorkloadRemovedFromPolicy", workloadIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockBuilderPolicyWorkloadRemovedFromPolicy)
				if err := _BlockBuilderPolicy.contract.UnpackLog(event, "WorkloadRemovedFromPolicy", log); err != nil {
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

// ParseWorkloadRemovedFromPolicy is a log parse operation binding the contract event 0x56c387a9be1bf0e0e4f852c577a225db98e8253ad401d1b4ea73926f27d6af09.
//
// Solidity: event WorkloadRemovedFromPolicy(bytes32 indexed workloadId)
func (_BlockBuilderPolicy *BlockBuilderPolicyFilterer) ParseWorkloadRemovedFromPolicy(log types.Log) (*BlockBuilderPolicyWorkloadRemovedFromPolicy, error) {
	event := new(BlockBuilderPolicyWorkloadRemovedFromPolicy)
	if err := _BlockBuilderPolicy.contract.UnpackLog(event, "WorkloadRemovedFromPolicy", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
