// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// OmniPortalInitParams is an auto generated low-level Go binding around an user-defined struct.
type OmniPortalInitParams struct {
	Owner                common.Address
	FeeOracle            common.Address
	OmniChainId          uint64
	OmniCChainId         uint64
	XmsgMaxGasLimit      uint64
	XmsgMinGasLimit      uint64
	XmsgMaxDataSize      uint16
	XreceiptMaxErrorSize uint16
	XsubValsetCutoff     uint8
	CChainXMsgOffset     uint64
	CChainXBlockOffset   uint64
	ValSetId             uint64
	Validators           []XTypesValidator
}

// XTypesBlockHeader is an auto generated low-level Go binding around an user-defined struct.
type XTypesBlockHeader struct {
	SourceChainId     uint64
	ConsensusChainId  uint64
	ConfLevel         uint8
	Offset            uint64
	SourceBlockHeight uint64
	SourceBlockHash   [32]byte
}

// XTypesChain is an auto generated low-level Go binding around an user-defined struct.
type XTypesChain struct {
	ChainId uint64
	Shards  []uint64
}

// XTypesMsg is an auto generated low-level Go binding around an user-defined struct.
type XTypesMsg struct {
	DestChainId uint64
	ShardId     uint64
	Offset      uint64
	Sender      common.Address
	To          common.Address
	Data        []byte
	GasLimit    uint64
}

// XTypesMsgContext is an auto generated low-level Go binding around an user-defined struct.
type XTypesMsgContext struct {
	SourceChainId uint64
	Sender        common.Address
}

// XTypesSigTuple is an auto generated low-level Go binding around an user-defined struct.
type XTypesSigTuple struct {
	ValidatorAddr common.Address
	Signature     []byte
}

// XTypesSubmission is an auto generated low-level Go binding around an user-defined struct.
type XTypesSubmission struct {
	AttestationRoot [32]byte
	ValidatorSetId  uint64
	BlockHeader     XTypesBlockHeader
	Msgs            []XTypesMsg
	Proof           [][32]byte
	ProofFlags      []bool
	Signatures      []XTypesSigTuple
}

// XTypesValidator is an auto generated low-level Go binding around an user-defined struct.
type XTypesValidator struct {
	Addr  common.Address
	Power uint64
}

// OmniPortalMetaData contains all meta data concerning the OmniPortal contract.
var OmniPortalMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ActionXCall\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ActionXSubmit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"KeyPauseAll\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"XSubQuorumDenominator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"XSubQuorumNumerator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addValidatorSet\",\"inputs\":[{\"name\":\"valSetId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"validators\",\"type\":\"tuple[]\",\"internalType\":\"structXTypes.Validator[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"power\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"chainId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"collectFees\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"feeFor\",\"inputs\":[{\"name\":\"destChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"feeOracle\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"inXBlockOffset\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"inXMsgOffset\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"p\",\"type\":\"tuple\",\"internalType\":\"structOmniPortal.InitParams\",\"components\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeOracle\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"omniChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"omniCChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"xmsgMaxGasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"xmsgMinGasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"xmsgMaxDataSize\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"xreceiptMaxErrorSize\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"xsubValsetCutoff\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"cChainXMsgOffset\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"cChainXBlockOffset\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"valSetId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"validators\",\"type\":\"tuple[]\",\"internalType\":\"structXTypes.Validator[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"power\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isPaused\",\"inputs\":[{\"name\":\"actionId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isPaused\",\"inputs\":[{\"name\":\"actionId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"chainId_\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isPaused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedDest\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedShard\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isXCall\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"latestValSetId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"network\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"omniCChainId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"omniChainId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"outXMsgOffset\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseXCall\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseXCallTo\",\"inputs\":[{\"name\":\"chainId_\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseXSubmit\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseXSubmitFrom\",\"inputs\":[{\"name\":\"chainId_\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFeeOracle\",\"inputs\":[{\"name\":\"feeOracle_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setInXBlockOffset\",\"inputs\":[{\"name\":\"sourceChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"shardId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"offset\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setInXMsgOffset\",\"inputs\":[{\"name\":\"sourceChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"shardId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"offset\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setNetwork\",\"inputs\":[{\"name\":\"network_\",\"type\":\"tuple[]\",\"internalType\":\"structXTypes.Chain[]\",\"components\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"shards\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setXMsgMaxDataSize\",\"inputs\":[{\"name\":\"numBytes\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setXMsgMaxGasLimit\",\"inputs\":[{\"name\":\"gasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setXMsgMinGasLimit\",\"inputs\":[{\"name\":\"gasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setXReceiptMaxErrorSize\",\"inputs\":[{\"name\":\"numBytes\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setXSubValsetCutoff\",\"inputs\":[{\"name\":\"xsubValsetCutoff_\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpauseXCall\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpauseXCallTo\",\"inputs\":[{\"name\":\"chainId_\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpauseXSubmit\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpauseXSubmitFrom\",\"inputs\":[{\"name\":\"chainId_\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"valSet\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"valSetTotalPower\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"xcall\",\"inputs\":[{\"name\":\"destChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"conf\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"xmsg\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structXTypes.MsgContext\",\"components\":[{\"name\":\"sourceChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"xmsgMaxDataSize\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"xmsgMaxGasLimit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"xmsgMinGasLimit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"xreceiptMaxErrorSize\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"xsubValsetCutoff\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"xsubmit\",\"inputs\":[{\"name\":\"xsub\",\"type\":\"tuple\",\"internalType\":\"structXTypes.Submission\",\"components\":[{\"name\":\"attestationRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"validatorSetId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"blockHeader\",\"type\":\"tuple\",\"internalType\":\"structXTypes.BlockHeader\",\"components\":[{\"name\":\"sourceChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"consensusChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"confLevel\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"offset\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sourceBlockHeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sourceBlockHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"msgs\",\"type\":\"tuple[]\",\"internalType\":\"structXTypes.Msg[]\",\"components\":[{\"name\":\"destChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"shardId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"offset\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"proof\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"proofFlags\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"},{\"name\":\"signatures\",\"type\":\"tuple[]\",\"internalType\":\"structXTypes.SigTuple[]\",\"components\":[{\"name\":\"validatorAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"FeeOracleSet\",\"inputs\":[{\"name\":\"oracle\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeesCollected\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InXBlockOffsetSet\",\"inputs\":[{\"name\":\"srcChainId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"shardId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"offset\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InXMsgOffsetSet\",\"inputs\":[{\"name\":\"srcChainId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"shardId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"offset\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorSetAdded\",\"inputs\":[{\"name\":\"setId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"XCallPaused\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"XCallToPaused\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"XCallToUnpaused\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"XCallUnpaused\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"XMsg\",\"inputs\":[{\"name\":\"destChainId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"shardId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"offset\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"fees\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"XMsgMaxDataSizeSet\",\"inputs\":[{\"name\":\"size\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"XMsgMaxGasLimitSet\",\"inputs\":[{\"name\":\"gasLimit\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"XMsgMinGasLimitSet\",\"inputs\":[{\"name\":\"gasLimit\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"XReceipt\",\"inputs\":[{\"name\":\"sourceChainId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"shardId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"offset\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"gasUsed\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"relayer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"success\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"error\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"XReceiptMaxErrorSizeSet\",\"inputs\":[{\"name\":\"size\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"XSubValsetCutoffSet\",\"inputs\":[{\"name\":\"cutoff\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"XSubmitFromPaused\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"XSubmitFromUnpaused\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"XSubmitPaused\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"XSubmitUnpaused\",\"inputs\":[],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureS\",\"inputs\":[{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MerkleProofInvalidMultiproof\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]}]",
	Bin: "0x60806040523480156200001157600080fd5b506200001c62000022565b620000d6565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff1615620000735760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b0390811614620000d35780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b614b8b80620000e66000396000f3fe60806040526004361061036b5760003560e01c80638532eb9f116101c6578063b4d5afd1116100f7578063c3d8ad6711610095578063d051c97d1161006f578063d051c97d14610af6578063d533b44514610b37578063f2fde38b14610b57578063f45cc7b814610b7757600080fd5b8063c3d8ad6714610a98578063c4ab80bc14610aad578063cf84c81814610acd57600080fd5b8063bff0e84d116100d1578063bff0e84d14610a25578063c21dda4f14610a45578063c26dfc0514610a58578063c2f9b96814610a7857600080fd5b8063b4d5afd1146109b0578063b521466d146109e5578063bb8590ad14610a0557600080fd5b8063a480ca7911610164578063afe821981161013e578063afe8219814610923578063afe8af9c14610943578063b187bd2614610979578063b2b2f5bd1461098e57600080fd5b8063a480ca79146108b3578063a8a98962146108d3578063aaf1bc97146108f357600080fd5b806397b52062116101a057806397b520621461083c5780639a8a05921461085c578063a10ac97a1461086f578063a32eb7c61461089157600080fd5b80638532eb9f146107b15780638da5cb5b146107d15780638dd9523c1461080e57600080fd5b80633f4ba83a116102a0578063575420501161023e57806374eba9391161021857806374eba9391461074057806378fe53071461076057806383d0cbd9146107875780638456cb591461079c57600080fd5b806357542050146106ca57806366a1eaf31461070b578063715018a61461072b57600080fd5b806349cc3bf61161027a57806349cc3bf614610643578063500b19e71461065d57806354d26bba1461069557806355e2448e146106aa57600080fd5b80633f4ba83a146105cd5780633fd3b15e146105e2578063461ab4881461062357600080fd5b8063241b71bb1161030d57806330632e8b116102e757806330632e8b1461052557806336d219121461054557806336d853f91461056c5780633aa873301461058c57600080fd5b8063241b71bb1461046057806324278bbe146104905780632f32700e146104c057600080fd5b806310a5a7f71161034957806310a5a7f7146103d3578063110ff5f1146103f35780631d3eb6e31461042b57806323dbce501461044b57600080fd5b80630360d20f1461037057806306c3dc5f1461039c578063103ba701146103b1575b600080fd5b34801561037c57600080fd5b50610385600281565b60405160ff90911681526020015b60405180910390f35b3480156103a857600080fd5b50610385600381565b3480156103bd57600080fd5b506103d16103cc366004613d44565b610b9e565b005b3480156103df57600080fd5b506103d16103ee366004613d7f565b610bb2565b3480156103ff57600080fd5b50600154610413906001600160401b031681565b6040516001600160401b039091168152602001610393565b34801561043757600080fd5b506103d1610446366004613d9c565b610c11565b34801561045757600080fd5b506103d1610d2c565b34801561046c57600080fd5b5061048061047b366004613e10565b610d76565b6040519015158152602001610393565b34801561049c57600080fd5b506104806104ab366004613d7f565b60056020526000908152604090205460ff1681565b3480156104cc57600080fd5b50604080518082018252600080825260209182015281518083018352600b546001600160401b0381168083526001600160a01b03600160401b909204821692840192835284519081529151169181019190915201610393565b34801561053157600080fd5b506103d1610540366004613e29565b610d87565b34801561055157600080fd5b5060015461041390600160401b90046001600160401b031681565b34801561057857600080fd5b506103d1610587366004613d7f565b61109e565b34801561059857600080fd5b506104136105a7366004613e64565b60066020908152600092835260408084209091529082529020546001600160401b031681565b3480156105d957600080fd5b506103d16110af565b3480156105ee57600080fd5b506104136105fd366004613e64565b60086020908152600092835260408084209091529082529020546001600160401b031681565b34801561062f57600080fd5b5061048061063e366004613e9d565b6110ea565b34801561064f57600080fd5b506000546103859060ff1681565b34801561066957600080fd5b5060025461067d906001600160a01b031681565b6040516001600160a01b039091168152602001610393565b3480156106a157600080fd5b506103d1611106565b3480156106b657600080fd5b50600b546001600160401b03161515610480565b3480156106d657600080fd5b506104136106e5366004613ed9565b600a6020908152600092835260408084209091529082529020546001600160401b031681565b34801561071757600080fd5b506103d1610726366004613f0e565b611150565b34801561073757600080fd5b506103d16114fe565b34801561074c57600080fd5b5061041361075b366004613e10565b611512565b34801561076c57600080fd5b5060005461041390600160681b90046001600160401b031681565b34801561079357600080fd5b506103d1611541565b3480156107a857600080fd5b506103d161158b565b3480156107bd57600080fd5b506103d16107cc366004613f49565b6115c6565b3480156107dd57600080fd5b507f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031661067d565b34801561081a57600080fd5b5061082e610829366004614018565b6116d9565b604051908152602001610393565b34801561084857600080fd5b506103d161085736600461407f565b61175a565b34801561086857600080fd5b5046610413565b34801561087b57600080fd5b5061082e600080516020614af683398151915281565b34801561089d57600080fd5b5061082e600080516020614b3683398151915281565b3480156108bf57600080fd5b506103d16108ce3660046140ca565b6117d7565b3480156108df57600080fd5b506103d16108ee3660046140ca565b61185f565b3480156108ff57600080fd5b5061048061090e366004613d7f565b60046020526000908152604090205460ff1681565b34801561092f57600080fd5b506103d161093e366004613d7f565b611870565b34801561094f57600080fd5b5061041361095e366004613d7f565b6009602052600090815260409020546001600160401b031681565b34801561098557600080fd5b506104806118ca565b34801561099a57600080fd5b5061082e600080516020614ad683398151915281565b3480156109bc57600080fd5b506000546109d2906301000000900461ffff1681565b60405161ffff9091168152602001610393565b3480156109f157600080fd5b506103d1610a003660046140e5565b611920565b348015610a1157600080fd5b506103d1610a20366004613d7f565b611931565b348015610a3157600080fd5b506103d1610a403660046140e5565b611942565b6103d1610a53366004614109565b611953565b348015610a6457600080fd5b506000546109d290610100900461ffff1681565b348015610a8457600080fd5b506103d1610a93366004613d7f565b611d2d565b348015610aa457600080fd5b506103d1611d8c565b348015610ab957600080fd5b506103d1610ac836600461407f565b611dd6565b348015610ad957600080fd5b50600054610413906501000000000090046001600160401b031681565b348015610b0257600080fd5b50610413610b11366004613e64565b60076020908152600092835260408084209091529082529020546001600160401b031681565b348015610b4357600080fd5b506103d1610b52366004613d7f565b611e4a565b348015610b6357600080fd5b506103d1610b723660046140ca565b611ea4565b348015610b8357600080fd5b5060005461041390600160a81b90046001600160401b031681565b610ba6611edf565b610baf81611f3a565b50565b610bba611edf565b610bda610bd5600080516020614ad683398151915283611fd6565b61201f565b6040516001600160401b038216907fcd7910e1c5569d8433ce4ef8e5d51c1bdc03168f614b576da47dc3d2b51d033a90600090a250565b333014610c5d5760405162461bcd60e51b815260206004820152601560248201527427b6b734a837b93a30b61d1037b7363c9039b2b63360591b60448201526064015b60405180910390fd5b600154600b546001600160401b03908116600160401b9092041614610cbe5760405162461bcd60e51b815260206004820152601760248201527627b6b734a837b93a30b61d1037b7363c9031b1b430b4b760491b6044820152606401610c54565b600b54600160401b90046001600160a01b031615610d1e5760405162461bcd60e51b815260206004820152601e60248201527f4f6d6e69506f7274616c3a206f6e6c792063636861696e2073656e64657200006044820152606401610c54565b610d28828261209a565b5050565b610d34611edf565b610d4b600080516020614b3683398151915261201f565b6040517f3d0f9c56dac46156a2db0aa09ee7804770ad9fc9549d21023164f22d69475ed890600090a1565b6000610d8182612214565b92915050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff1615906001600160401b0316600081158015610dcc5750825b90506000826001600160401b03166001148015610de85750303b155b905081158015610df6575080155b15610e145760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff191660011785558315610e3e57845460ff60401b1916600160401b1785555b610e53610e4e60208801886140ca565b61227b565b610e6b610e6660408801602089016140ca565b61228c565b610e83610e7e60a0880160808901613d7f565b612330565b610e9b610e9660c0880160a08901613d7f565b6123e8565b610eb3610eae60e0880160c089016140e5565b61249c565b610ecc610ec7610100880160e089016140e5565b612540565b610ee6610ee161012088016101008901613d44565b611f3a565b610f0e610efb61018088016101608901613d7f565b610f09610180890189614192565b6125e0565b610f1e6060870160408801613d7f565b6001805467ffffffffffffffff19166001600160401b0392909216919091179055610f4f6080870160608801613d7f565b600180546001600160401b0392909216600160401b026fffffffffffffffff000000000000000019909216919091179055610104610f9561014088016101208901613d7f565b60076000610fa960808b0160608c01613d7f565b6001600160401b0390811682526020808301939093526040918201600090812086831682529093529120805467ffffffffffffffff191692909116919091179055610ffc61016088016101408901613d7f565b6008600061101060808b0160608c01613d7f565b6001600160401b03908116825260208083019390935260409182016000908120958216815294909252909220805467ffffffffffffffff191691909216179055831561109657845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050565b6110a6611edf565b610baf81612330565b6110b7611edf565b6110bf61290f565b6040517fa45f47fdea8a1efdd9029a5691c7f759c32b7c698632b563573e155625d1693390600090a1565b60006110ff836110fa8585611fd6565b612926565b9392505050565b61110e611edf565b611125600080516020614ad68339815191526129ad565b6040517f4c48c7b71557216a3192842746bdfc381f98d7536d9eb1c6764f3b45e679482790600090a1565b600080516020614b3683398151915261116f6060830160408401613d7f565b61117d826110fa8484611fd6565b156111bf5760405162461bcd60e51b815260206004820152601260248201527113db5b9a541bdc9d185b0e881c185d5cd95960721b6044820152606401610c54565b6111c7612a28565b3660006111d86101008601866141db565b90925090506040850160006111f08260208901613d7f565b600154909150600160401b90046001600160401b03166112166040840160208501613d7f565b6001600160401b03161461126c5760405162461bcd60e51b815260206004820152601b60248201527f4f6d6e69506f7274616c3a2077726f6e672063636861696e20494400000000006044820152606401610c54565b826112b05760405162461bcd60e51b81526020600482015260146024820152734f6d6e69506f7274616c3a206e6f20786d73677360601b6044820152606401610c54565b6001600160401b03808216600090815260096020526040902054166113175760405162461bcd60e51b815260206004820152601b60248201527f4f6d6e69506f7274616c3a20756e6b6e6f776e2076616c2073657400000000006044820152606401610c54565b61131f612a72565b6001600160401b0316816001600160401b031610156113805760405162461bcd60e51b815260206004820152601760248201527f4f6d6e69506f7274616c3a206f6c642076616c207365740000000000000000006044820152606401610c54565b6113c487356113936101608a018a6141db565b6001600160401b038086166000908152600a6020908152604080832060099092529091205490911660026003612ac2565b6114085760405162461bcd60e51b81526020600482015260156024820152744f6d6e69506f7274616c3a206e6f2071756f72756d60581b6044820152606401610c54565b611431873583868661141e6101208d018d6141db565b61142c6101408f018f6141db565b612ce4565b61147d5760405162461bcd60e51b815260206004820152601960248201527f4f6d6e69506f7274616c3a20696e76616c69642070726f6f66000000000000006044820152606401610c54565b60005b838110156114cb576114c361149a36859003850185614292565b8686848181106114ac576114ac614333565b90506020028101906114be9190614349565b612d5f565b600101611480565b50505050506114f960017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b505050565b611506611edf565b611510600061324a565b565b6003818154811061152257600080fd5b60009182526020909120600290910201546001600160401b0316905081565b611549611edf565b611560600080516020614ad683398151915261201f565b6040517f5f335a4032d4cfb6aca7835b0c2225f36d4d9eaa4ed43ee59ed537e02dff6b3990600090a1565b611593611edf565b61159b6132bb565b6040517f9e87fac88ff661f02d44f95383c817fece4bce600a3dab7a54406878b965e75290600090a1565b33301461160d5760405162461bcd60e51b815260206004820152601560248201527427b6b734a837b93a30b61d1037b7363c9039b2b63360591b6044820152606401610c54565b600154600b546001600160401b03908116600160401b909204161461166e5760405162461bcd60e51b815260206004820152601760248201527627b6b734a837b93a30b61d1037b7363c9031b1b430b4b760491b6044820152606401610c54565b600b54600160401b90046001600160a01b0316156116ce5760405162461bcd60e51b815260206004820152601e60248201527f4f6d6e69506f7274616c3a206f6e6c792063636861696e2073656e64657200006044820152606401610c54565b6114f98383836125e0565b600254604051632376548f60e21b81526000916001600160a01b031690638dd9523c90611710908890889088908890600401614392565b602060405180830381865afa15801561172d573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061175191906143ca565b95945050505050565b611762611edf565b6001600160401b03838116600081815260086020908152604080832087861680855290835292819020805467ffffffffffffffff191695871695861790555193845290927fe070f08cae8464c91238e8cbea64ccee5e7b48dd79a843f144e3721ee6bdd9b591015b60405180910390a3505050565b6117df611edf565b60405147906001600160a01b0383169082156108fc029083906000818181858888f19350505050158015611817573d6000803e3d6000fd5b50816001600160a01b03167f9dc46f23cfb5ddcad0ae7ea2be38d47fec07bb9382ec7e564efc69e036dd66ce8260405161185391815260200190565b60405180910390a25050565b611867611edf565b610baf8161228c565b611878611edf565b611893610bd5600080516020614b3683398151915283611fd6565b6040516001600160401b038216907fab78810a0515df65f9f10bfbcb92d03d5df71d9fd3b9414e9ad831a5117d6daa90600090a250565b600061191b600080516020614af6833981519152600052600080516020614b168339815191526020527ffae9838a178d7f201aa98e2ce5340158edda60bb1e8f168f46503bf3e99f13be5460ff1690565b905090565b611928611edf565b610baf8161249c565b611939611edf565b610baf816123e8565b61194a611edf565b610baf81612540565b600080516020614ad683398151915286611971826110fa8484611fd6565b156119b35760405162461bcd60e51b815260206004820152601260248201527113db5b9a541bdc9d185b0e881c185d5cd95960721b6044820152606401610c54565b6001600160401b03881660009081526005602052604090205460ff16611a1b5760405162461bcd60e51b815260206004820152601c60248201527f4f6d6e69506f7274616c3a20756e737570706f727465642064657374000000006044820152606401610c54565b6001600160a01b038616611a715760405162461bcd60e51b815260206004820152601b60248201527f4f6d6e69506f7274616c3a206e6f20706f7274616c207863616c6c00000000006044820152606401610c54565b6000546001600160401b036501000000000090910481169084161115611ad95760405162461bcd60e51b815260206004820152601d60248201527f4f6d6e69506f7274616c3a206761734c696d697420746f6f20686967680000006044820152606401610c54565b6000546001600160401b03600160681b90910481169084161015611b3f5760405162461bcd60e51b815260206004820152601c60248201527f4f6d6e69506f7274616c3a206761734c696d697420746f6f206c6f77000000006044820152606401610c54565b6000546301000000900461ffff16841115611b9c5760405162461bcd60e51b815260206004820152601a60248201527f4f6d6e69506f7274616c3a206461746120746f6f206c617267650000000000006044820152606401610c54565b60ff808816600081815260046020526040902054909116611bff5760405162461bcd60e51b815260206004820152601d60248201527f4f6d6e69506f7274616c3a20756e737570706f727465642073686172640000006044820152606401610c54565b6000611c0d8a8888886116d9565b905080341015611c5f5760405162461bcd60e51b815260206004820152601c60248201527f4f6d6e69506f7274616c3a20696e73756666696369656e7420666565000000006044820152606401610c54565b6001600160401b03808b166000908152600660209081526040808320868516845290915281208054600193919291611c99918591166143f9565b82546101009290920a6001600160401b038181021990931691831602179091558b811660008181526006602090815260408083208886168085529252918290205491519190931693507fb7c8eb9d7a7fbcdab809ab7b8a7c41701eb3115e3fe99d30ff490d8552f72bfa90611d199033908e908e908e908e908b90614420565b60405180910390a450505050505050505050565b611d35611edf565b611d55611d50600080516020614b3683398151915283611fd6565b6129ad565b6040516001600160401b038216907fc551305d9bd408be4327b7f8aba28b04ccf6b6c76925392d195ecf9cc764294d90600090a250565b611d94611edf565b611dab600080516020614b368339815191526129ad565b6040517f2cb9d71d4c31860b70e9b707c69aa2f5953e03474f00cfcfff205c4745f8287590600090a1565b611dde611edf565b6001600160401b03838116600081815260076020908152604080832087861680855290835292819020805467ffffffffffffffff191695871695861790555193845290927f8647aae68c8456a1dcbfaf5eaadc94278ae423526d3f09c7b972bff7355d55c791016117ca565b611e52611edf565b611e6d611d50600080516020614ad683398151915283611fd6565b6040516001600160401b038216907f1ed9223556fb0971076c30172f1f00630efd313b6a05290a562aef95928e712590600090a250565b611eac611edf565b6001600160a01b038116611ed657604051631e4fbdf760e01b815260006004820152602401610c54565b610baf8161324a565b33611f117f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b6001600160a01b0316146115105760405163118cdaa760e01b8152336004820152602401610c54565b60008160ff1611611f8d5760405162461bcd60e51b815260206004820152601a60248201527f4f6d6e69506f7274616c3a206e6f207a65726f206375746f66660000000000006044820152606401610c54565b6000805460ff191660ff83169081179091556040519081527f1683dc51426224f6e37a3b41dd5849e2db1bfe22366d1d913fa0ef6f757e828f906020015b60405180910390a150565b6000828260405160200161200192919091825260c01b6001600160c01b031916602082015260280190565b60405160208183030381529060405280519060200120905092915050565b6000818152600080516020614b16833981519152602081905260409091205460ff16156120815760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b6044820152606401610c54565b600091825260205260409020805460ff19166001179055565b6120a26132d2565b3660005b8281101561220e578383828181106120c0576120c0614333565b90506020028101906120d2919061446b565b6003805460018101825560009190915290925082906002027fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b016121168282614507565b505061211f4690565b6001600160401b03166121356020840184613d7f565b6001600160401b031614612183576001600560006121566020860186613d7f565b6001600160401b031681526020810191909152604001600020805460ff1916911515919091179055612206565b60005b61219360208401846141db565b9050811015612204576001600460006121af60208701876141db565b858181106121bf576121bf614333565b90506020020160208101906121d49190613d7f565b6001600160401b031681526020810191909152604001600020805460ff1916911515919091179055600101612186565b505b6001016120a6565b50505050565b600080516020614af68339815191526000908152600080516020614b1683398151915260208190527ffae9838a178d7f201aa98e2ce5340158edda60bb1e8f168f46503bf3e99f13be5460ff16806110ff5750600092835260205250604090205460ff1690565b6122836133d1565b610baf8161341a565b6001600160a01b0381166122e25760405162461bcd60e51b815260206004820152601d60248201527f4f6d6e69506f7274616c3a206e6f207a65726f206665654f7261636c650000006044820152606401610c54565b600280546001600160a01b0319166001600160a01b0383169081179091556040519081527fd97bdb0db82b52a85aa07f8da78033b1d6e159d94f1e3cbd4109d946c3bcfd3290602001611fcb565b6000816001600160401b0316116123895760405162461bcd60e51b815260206004820152601b60248201527f4f6d6e69506f7274616c3a206e6f207a65726f206d61782067617300000000006044820152606401610c54565b600080546cffffffffffffffff00000000001916650100000000006001600160401b038416908102919091179091556040519081527f1153561ac5effc2926ba6c612f86a397c997bc43dfbfc718da08065be0c5fe4d90602001611fcb565b6000816001600160401b0316116124415760405162461bcd60e51b815260206004820152601b60248201527f4f6d6e69506f7274616c3a206e6f207a65726f206d696e2067617300000000006044820152606401610c54565b6000805467ffffffffffffffff60681b1916600160681b6001600160401b038416908102919091179091556040519081527f8c852a6291aa436654b167353bca4a4b0c3d024c7562cb5082e7c869bddabf3e90602001611fcb565b60008161ffff16116124f05760405162461bcd60e51b815260206004820152601c60248201527f4f6d6e69506f7274616c3a206e6f207a65726f206d61782073697a65000000006044820152606401610c54565b6000805464ffff0000001916630100000061ffff8416908102919091179091556040519081527f65923e04419dc810d0ea08a94a7f608d4c4d949818d95c3788f895e575dd206490602001611fcb565b60008161ffff16116125945760405162461bcd60e51b815260206004820152601c60248201527f4f6d6e69506f7274616c3a206e6f207a65726f206d61782073697a65000000006044820152606401610c54565b6000805462ffff00191661010061ffff8416908102919091179091556040519081527f620bbea084306b66a8cc6b5b63830d6b3874f9d2438914e259ffd5065c33f7b090602001611fcb565b808061262e5760405162461bcd60e51b815260206004820152601960248201527f4f6d6e69506f7274616c3a206e6f2076616c696461746f7273000000000000006044820152606401610c54565b6001600160401b0380851660009081526009602052604090205416156126965760405162461bcd60e51b815260206004820152601d60248201527f4f6d6e69506f7274616c3a206475706c69636174652076616c207365740000006044820152606401610c54565b604080518082018252600080825260208083018290526001600160401b0388168252600a9052918220825b8481101561286e578686828181106126db576126db614333565b9050604002018036038101906126f1919061462f565b80519093506001600160a01b031661274b5760405162461bcd60e51b815260206004820152601d60248201527f4f6d6e69506f7274616c3a206e6f207a65726f2076616c696461746f720000006044820152606401610c54565b600083602001516001600160401b0316116127a85760405162461bcd60e51b815260206004820152601960248201527f4f6d6e69506f7274616c3a206e6f207a65726f20706f776572000000000000006044820152606401610c54565b82516001600160a01b03166000908152602083905260409020546001600160401b0316156128185760405162461bcd60e51b815260206004820152601f60248201527f4f6d6e69506f7274616c3a206475706c69636174652076616c696461746f72006044820152606401610c54565b602083015161282790856143f9565b60208481015185516001600160a01b03166000908152918590526040909120805467ffffffffffffffff19166001600160401b0390921691909117905593506001016126c1565b506001600160401b038781166000818152600960205260408120805467ffffffffffffffff191687851617905554600160a81b900490911610156128d2576000805467ffffffffffffffff60a81b1916600160a81b6001600160401b038a16021790555b6040516001600160401b038816907f3a7c2f997a87ba92aedaecd1127f4129cae1283e2809ebf5304d321b943fd10790600090a250505050505050565b611510600080516020614af68339815191526129ad565b600080516020614af68339815191526000908152600080516020614b1683398151915260208190527ffae9838a178d7f201aa98e2ce5340158edda60bb1e8f168f46503bf3e99f13be5460ff168061298c575060008481526020829052604090205460ff165b806129a5575060008381526020829052604090205460ff165b949350505050565b6000818152600080516020614b16833981519152602081905260409091205460ff16612a125760405162461bcd60e51b815260206004820152601460248201527314185d5cd8589b194e881b9bdd081c185d5cd95960621b6044820152606401610c54565b600091825260205260409020805460ff19169055565b7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00805460011901612a6c57604051633ee5aeb560e01b815260040160405180910390fd5b60029055565b6000805460ff8116600160a81b9091046001600160401b031611612a965750600190565b600054612ab79060ff811690600160a81b90046001600160401b031661466e565b61191b9060016143f9565b6000803660005b88811015612cd157898982818110612ae357612ae3614333565b9050602002810190612af5919061446b565b91508015612c175760008a8a612b0c60018561468e565b818110612b1b57612b1b614333565b9050602002810190612b2d919061446b565b612b36906146a1565b80519091506001600160a01b0316612b5160208501856140ca565b6001600160a01b031603612ba75760405162461bcd60e51b815260206004820152601b60248201527f51756f72756d3a206475706c69636174652076616c696461746f7200000000006044820152606401610c54565b80516001600160a01b0316612bbf60208501856140ca565b6001600160a01b031611612c155760405162461bcd60e51b815260206004820152601760248201527f51756f72756d3a2073696773206e6f7420736f727465640000000000000000006044820152606401610c54565b505b612c21828c613422565b612c6d5760405162461bcd60e51b815260206004820152601960248201527f51756f72756d3a20696e76616c6964207369676e6174757265000000000000006044820152606401610c54565b876000612c7d60208501856140ca565b6001600160a01b03168152602081019190915260400160002054612caa906001600160401b0316846143f9565b9250612cb883888888613496565b15612cc95760019350505050612cd9565b600101612ac9565b506000925050505b979650505050505050565b60408051600180825281830190925260009182919060208083019080368337019050509050612d1f86868686612d1a8d8d6134d3565b6135a0565b81600081518110612d3257612d32614333565b602002602001018181525050612d51818b612d4c8c613801565b613819565b9a9950505050505050505050565b81516000612d706020840184613d7f565b90506000612d846040850160208601613d7f565b90506000612d986060860160408701613d7f565b9050466001600160401b0316836001600160401b03161480612dc157506001600160401b038316155b612e0d5760405162461bcd60e51b815260206004820152601c60248201527f4f6d6e69506f7274616c3a2077726f6e67206465737420636861696e000000006044820152606401610c54565b6001600160401b0380851660009081526007602090815260408083208685168452909152902054612e40911660016143f9565b6001600160401b0316816001600160401b031614612ea05760405162461bcd60e51b815260206004820152601860248201527f4f6d6e69506f7274616c3a2077726f6e67206f666673657400000000000000006044820152606401610c54565b856040015160ff16600460ff161480612ec257508160ff16866040015160ff16145b612f0e5760405162461bcd60e51b815260206004820152601c60248201527f4f6d6e69506f7274616c3a2077726f6e6720636f6e66206c6576656c000000006044820152606401610c54565b60608601516001600160401b038581166000908152600860209081526040808320878516845290915290205491811691161015612f855760608601516001600160401b03858116600090815260086020908152604080832087851684529091529020805467ffffffffffffffff1916919092161790555b6001600160401b038085166000908152600760209081526040808320868516845290915281208054600193919291612fbf918591166143f9565b92506101000a8154816001600160401b0302191690836001600160401b03160217905550306001600160a01b031685608001602081019061300091906140ca565b6001600160a01b0316036130da57806001600160401b0316826001600160401b0316856001600160401b03167f8277cab1f0fa69b34674f64a7d43f242b0bacece6f5b7e8652f1e0d88a9b873b6000336000604051602401613093906020808252601e908201527f4f6d6e69506f7274616c3a206e6f207863616c6c20746f20706f7274616c0000604082015260600190565b60408051601f198184030181529181526020820180516001600160e01b031662461bcd60e51b179052516130ca949392919061479c565b60405180910390a4505050505050565b604080518082019091526001600160401b03851681526020810161310460808801606089016140ca565b6001600160a01b039081169091528151600b8054602090940151909216600160401b026001600160e01b03199093166001600160401b0390911617919091179055600080808061315a60a08a0160808b016140ca565b6001600160a01b0316146131ab576131a661317b60a08a0160808b016140ca565b61318b60e08b0160c08c01613d7f565b6001600160401b03166131a160a08c018c6147d8565b61382f565b6131c0565b6131c06131bb60a08a018a6147d8565b6138ef565b600b80546001600160e01b0319169055919450925090506000836131e457826131f5565b604051806020016040528060008152505b9050846001600160401b0316866001600160401b0316896001600160401b03167f8277cab1f0fa69b34674f64a7d43f242b0bacece6f5b7e8652f1e0d88a9b873b85338987604051611d19949392919061479c565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3505050565b611510600080516020614af683398151915261201f565b6000805b6003548110156133c457600381815481106132f3576132f3614333565b9060005260206000209060020201915061330a4690565b82546001600160401b039081169116146133445781546001600160401b03166000908152600560205260409020805460ff191690556133bc565b60005b60018301548110156133ba5760006004600085600101848154811061336e5761336e614333565b6000918252602080832060048304015460039092166008026101000a9091046001600160401b031683528201929092526040019020805460ff1916911515919091179055600101613347565b505b6001016132d6565b50610baf60036000613ca9565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff1661151057604051631afcd79f60e31b815260040160405180910390fd5b611eac6133d1565b600061343160208401846140ca565b6001600160a01b03166134858361344b60208701876147d8565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061398692505050565b6001600160a01b0316149392505050565b60008160ff168360ff16856134ab919061481e565b6134b5919061485f565b6001600160401b0316856001600160401b0316119050949350505050565b60606000826001600160401b038111156134ef576134ef614224565b604051908082528060200260200182016040528015613518578160200160208202803683370190505b50905060005b8381101561359857613573600286868481811061353d5761353d614333565b905060200281019061354f9190614349565b60405160200161355f91906148ca565b6040516020818303038152906040526139b0565b82828151811061358557613585614333565b602090810291909101015260010161351e565b509392505050565b805160009085846135b2816001614999565b6135bc8385614999565b146135da57604051631a8a024960e11b815260040160405180910390fd5b6000816001600160401b038111156135f4576135f4614224565b60405190808252806020026020018201604052801561361d578160200160208202803683370190505b5090506000806000805b8581101561376a576000888510613662578584613643816149ac565b95508151811061365557613655614333565b6020026020010151613688565b8a8561366d816149ac565b96508151811061367f5761367f614333565b60200260200101515b905060008d8d8481811061369e5761369e614333565b90506020020160208101906136b391906149c5565b6136e0578f8f856136c3816149ac565b96508181106136d4576136d4614333565b90506020020135613737565b8986106137115786856136f2816149ac565b96508151811061370457613704614333565b6020026020010151613737565b8b8661371c816149ac565b97508151811061372e5761372e614333565b60200260200101515b905061374382826139e7565b87848151811061375557613755614333565b60209081029190910101525050600101613627565b5084156137bc5785811461379157604051631a8a024960e11b815260040160405180910390fd5b8360018603815181106137a6576137a6614333565b6020026020010151975050505050505050611751565b86156137d557886000815181106137a6576137a6614333565b8c8c60008181106137e8576137e8614333565b9050602002013597505050505050505095945050505050565b6000610d8160018360405160200161355f91906149e7565b6000826138268584613a16565b14949350505050565b600060606000805a90506000806138b28960008060019054906101000a900461ffff168b8b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050508e6001600160a01b0316613a5190949392919063ffffffff16565b9150915060005a90506138c6603f8b614a6c565b81116138ce57fe5b82826138da838761468e565b965096509650505050505b9450945094915050565b600060606000805a9050600080306001600160a01b03168888604051613916929190614a80565b6000604051808303816000865af19150503d8060008114613953576040519150601f19603f3d011682016040523d82523d6000602084013e613958565b606091505b50915091505a613968908461468e565b92508161397757805160208201fd5b909450925090505b9250925092565b6000806000806139968686613adb565b9250925092506139a68282613b25565b5090949350505050565b600082826040516020016139c5929190614a90565b60408051601f1981840301815282825280516020918201209083015201612001565b6000818310613a035760008281526020849052604090206110ff565b60008381526020839052604090206110ff565b600081815b845181101561359857613a4782868381518110613a3a57613a3a614333565b60200260200101516139e7565b9150600101613a1b565b6000606060008060008661ffff166001600160401b03811115613a7657613a76614224565b6040519080825280601f01601f191660200182016040528015613aa0576020820181803683370190505b5090506000808751602089018b8e8ef191503d925086831115613ac1578692505b828152826000602083013e90999098509650505050505050565b60008060008351604103613b155760208401516040850151606086015160001a613b0788828585613bde565b95509550955050505061397f565b505081516000915060029061397f565b6000826003811115613b3957613b39614abf565b03613b42575050565b6001826003811115613b5657613b56614abf565b03613b745760405163f645eedf60e01b815260040160405180910390fd5b6002826003811115613b8857613b88614abf565b03613ba95760405163fce698f760e01b815260048101829052602401610c54565b6003826003811115613bbd57613bbd614abf565b03610d28576040516335e2f38360e21b815260048101829052602401610c54565b600080807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0841115613c1957506000915060039050826138e5565b604080516000808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa158015613c6d573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116613c99575060009250600191508290506138e5565b9760009750879650945050505050565b5080546000825560020290600052602060002090810190610baf91905b80821115613cf557805467ffffffffffffffff191681556000613cec6001830182613cf9565b50600201613cc6565b5090565b508054600082556003016004900490600052602060002090810190610baf91905b80821115613cf55760008155600101613d1a565b803560ff81168114613d3f57600080fd5b919050565b600060208284031215613d5657600080fd5b6110ff82613d2e565b6001600160401b0381168114610baf57600080fd5b8035613d3f81613d5f565b600060208284031215613d9157600080fd5b81356110ff81613d5f565b60008060208385031215613daf57600080fd5b82356001600160401b0380821115613dc657600080fd5b818501915085601f830112613dda57600080fd5b813581811115613de957600080fd5b8660208260051b8501011115613dfe57600080fd5b60209290920196919550909350505050565b600060208284031215613e2257600080fd5b5035919050565b600060208284031215613e3b57600080fd5b81356001600160401b03811115613e5157600080fd5b82016101a081850312156110ff57600080fd5b60008060408385031215613e7757600080fd5b8235613e8281613d5f565b91506020830135613e9281613d5f565b809150509250929050565b60008060408385031215613eb057600080fd5b823591506020830135613e9281613d5f565b80356001600160a01b0381168114613d3f57600080fd5b60008060408385031215613eec57600080fd5b8235613ef781613d5f565b9150613f0560208401613ec2565b90509250929050565b600060208284031215613f2057600080fd5b81356001600160401b03811115613f3657600080fd5b820161018081850312156110ff57600080fd5b600080600060408486031215613f5e57600080fd5b8335613f6981613d5f565b925060208401356001600160401b0380821115613f8557600080fd5b818601915086601f830112613f9957600080fd5b813581811115613fa857600080fd5b8760208260061b8501011115613fbd57600080fd5b6020830194508093505050509250925092565b60008083601f840112613fe257600080fd5b5081356001600160401b03811115613ff957600080fd5b60208301915083602082850101111561401157600080fd5b9250929050565b6000806000806060858703121561402e57600080fd5b843561403981613d5f565b935060208501356001600160401b0381111561405457600080fd5b61406087828801613fd0565b909450925050604085013561407481613d5f565b939692955090935050565b60008060006060848603121561409457600080fd5b833561409f81613d5f565b925060208401356140af81613d5f565b915060408401356140bf81613d5f565b809150509250925092565b6000602082840312156140dc57600080fd5b6110ff82613ec2565b6000602082840312156140f757600080fd5b813561ffff811681146110ff57600080fd5b60008060008060008060a0878903121561412257600080fd5b863561412d81613d5f565b955061413b60208801613d2e565b945061414960408801613ec2565b935060608701356001600160401b0381111561416457600080fd5b61417089828a01613fd0565b909450925050608087013561418481613d5f565b809150509295509295509295565b6000808335601e198436030181126141a957600080fd5b8301803591506001600160401b038211156141c357600080fd5b6020019150600681901b360382131561401157600080fd5b6000808335601e198436030181126141f257600080fd5b8301803591506001600160401b0382111561420c57600080fd5b6020019150600581901b360382131561401157600080fd5b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b038111828210171561425c5761425c614224565b60405290565b604051601f8201601f191681016001600160401b038111828210171561428a5761428a614224565b604052919050565b600060c082840312156142a457600080fd5b60405160c081018181106001600160401b03821117156142c6576142c6614224565b60405282356142d481613d5f565b815260208301356142e481613d5f565b60208201526142f560408401613d2e565b6040820152606083013561430881613d5f565b6060820152608083013561431b81613d5f565b608082015260a0928301359281019290925250919050565b634e487b7160e01b600052603260045260246000fd5b6000823560de1983360301811261435f57600080fd5b9190910192915050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b60006001600160401b038087168352606060208401526143b6606084018688614369565b915080841660408401525095945050505050565b6000602082840312156143dc57600080fd5b5051919050565b634e487b7160e01b600052601160045260246000fd5b6001600160401b03818116838216019080821115614419576144196143e3565b5092915050565b6001600160a01b0387811682528616602082015260a06040820181905260009061444d9083018688614369565b6001600160401b039490941660608301525060800152949350505050565b60008235603e1983360301811261435f57600080fd5b60008135610d8181613d5f565b600160401b8211156144a2576144a2614224565b8054828255808310156114f95760008260005260206000206003850160021c81016003840160021c8201915060188660031b1680156144f2576000198083018054828460200360031b1c16815550505b505b81811015611096578281556001016144f4565b813561451281613d5f565b815467ffffffffffffffff19166001600160401b0391821617825560019081830160208581013536879003601e1901811261454c57600080fd5b860180358481111561455d57600080fd5b6020820194508060051b360385131561457557600080fd5b61457f818561448e565b60009384526020842093600282901c92505b828110156145e8576000805b60048110156145dc576145cf6145b289614481565b6001600160401b03908116600684901b90811b91901b1984161790565b978601979150880161459d565b50858201558601614591565b506003198116808203818314614623576000805b8281101561461d576146106145b28a614481565b98870198915089016145fc565b50868501555b50505050505050505050565b60006040828403121561464157600080fd5b61464961423a565b61465283613ec2565b8152602083013561466281613d5f565b60208201529392505050565b6001600160401b03828116828216039080821115614419576144196143e3565b81810381811115610d8157610d816143e3565b6000604082360312156146b357600080fd5b6146bb61423a565b6146c483613ec2565b81526020808401356001600160401b03808211156146e157600080fd5b9085019036601f8301126146f457600080fd5b81358181111561470657614706614224565b614718601f8201601f19168501614262565b9150808252368482850101111561472e57600080fd5b80848401858401376000908201840152918301919091525092915050565b60005b8381101561476757818101518382015260200161474f565b50506000910152565b6000815180845261478881602086016020860161474c565b601f01601f19169290920160200192915050565b8481526001600160a01b038416602082015282151560408201526080606082018190526000906147ce90830184614770565b9695505050505050565b6000808335601e198436030181126147ef57600080fd5b8301803591506001600160401b0382111561480957600080fd5b60200191503681900382131561401157600080fd5b6001600160401b03818116838216028082169190828114614841576148416143e3565b505092915050565b634e487b7160e01b600052601260045260246000fd5b60006001600160401b038084168061487957614879614849565b92169190910492915050565b6000808335601e1984360301811261489c57600080fd5b83016020810192503590506001600160401b038111156148bb57600080fd5b80360382131561401157600080fd5b60208152600082356148db81613d5f565b6001600160401b038082166020850152602085013591506148fb82613d5f565b80821660408501526040850135915061491382613d5f565b166060838101919091526001600160a01b0390614931908501613ec2565b16608083015261494360808401613ec2565b6001600160a01b03811660a08401525061496060a0840184614885565b60e060c085015261497661010085018284614369565b91505061498560c08501613d74565b6001600160401b03811660e0850152613598565b80820180821115610d8157610d816143e3565b6000600182016149be576149be6143e3565b5060010190565b6000602082840312156149d757600080fd5b813580151581146110ff57600080fd5b60c0810182356149f681613d5f565b6001600160401b039081168352602084013590614a1282613d5f565b808216602085015260ff614a2860408701613d2e565b16604085015260608501359150614a3e82613d5f565b9081166060840152608084013590614a5582613d5f565b16608083015260a092830135929091019190915290565b600082614a7b57614a7b614849565b500490565b8183823760009101908152919050565b60ff60f81b8360f81b16815260008251614ab181600185016020870161474c565b919091016001019392505050565b634e487b7160e01b600052602160045260246000fdfea06a0c1264badca141841b5f52470407dac9adaaa539dd445540986341b73a6876e8952e4b09b8d505aa08998d716721a1dbf0884ac74202e33985da1ed005e9ff37105740f03695c8f3597f3aff2b92fbe1c80abea3c28731ecff2efd693400feccba1cfc4544bf9cd83b76f36ae5c464750b6c43f682e26744ee21ec31fc1ea264697066735822122082007a797748fabc04818a77901716e59f82472d3f9ab0e18d0fe1956de8602a64736f6c63430008180033",
}

// OmniPortalABI is the input ABI used to generate the binding from.
// Deprecated: Use OmniPortalMetaData.ABI instead.
var OmniPortalABI = OmniPortalMetaData.ABI

// OmniPortalBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OmniPortalMetaData.Bin instead.
var OmniPortalBin = OmniPortalMetaData.Bin

// DeployOmniPortal deploys a new Ethereum contract, binding an instance of OmniPortal to it.
func DeployOmniPortal(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OmniPortal, error) {
	parsed, err := OmniPortalMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OmniPortalBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OmniPortal{OmniPortalCaller: OmniPortalCaller{contract: contract}, OmniPortalTransactor: OmniPortalTransactor{contract: contract}, OmniPortalFilterer: OmniPortalFilterer{contract: contract}}, nil
}

// OmniPortal is an auto generated Go binding around an Ethereum contract.
type OmniPortal struct {
	OmniPortalCaller     // Read-only binding to the contract
	OmniPortalTransactor // Write-only binding to the contract
	OmniPortalFilterer   // Log filterer for contract events
}

// OmniPortalCaller is an auto generated read-only Go binding around an Ethereum contract.
type OmniPortalCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OmniPortalTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OmniPortalTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OmniPortalFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OmniPortalFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OmniPortalSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OmniPortalSession struct {
	Contract     *OmniPortal       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OmniPortalCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OmniPortalCallerSession struct {
	Contract *OmniPortalCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// OmniPortalTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OmniPortalTransactorSession struct {
	Contract     *OmniPortalTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// OmniPortalRaw is an auto generated low-level Go binding around an Ethereum contract.
type OmniPortalRaw struct {
	Contract *OmniPortal // Generic contract binding to access the raw methods on
}

// OmniPortalCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OmniPortalCallerRaw struct {
	Contract *OmniPortalCaller // Generic read-only contract binding to access the raw methods on
}

// OmniPortalTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OmniPortalTransactorRaw struct {
	Contract *OmniPortalTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOmniPortal creates a new instance of OmniPortal, bound to a specific deployed contract.
func NewOmniPortal(address common.Address, backend bind.ContractBackend) (*OmniPortal, error) {
	contract, err := bindOmniPortal(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OmniPortal{OmniPortalCaller: OmniPortalCaller{contract: contract}, OmniPortalTransactor: OmniPortalTransactor{contract: contract}, OmniPortalFilterer: OmniPortalFilterer{contract: contract}}, nil
}

// NewOmniPortalCaller creates a new read-only instance of OmniPortal, bound to a specific deployed contract.
func NewOmniPortalCaller(address common.Address, caller bind.ContractCaller) (*OmniPortalCaller, error) {
	contract, err := bindOmniPortal(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OmniPortalCaller{contract: contract}, nil
}

// NewOmniPortalTransactor creates a new write-only instance of OmniPortal, bound to a specific deployed contract.
func NewOmniPortalTransactor(address common.Address, transactor bind.ContractTransactor) (*OmniPortalTransactor, error) {
	contract, err := bindOmniPortal(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OmniPortalTransactor{contract: contract}, nil
}

// NewOmniPortalFilterer creates a new log filterer instance of OmniPortal, bound to a specific deployed contract.
func NewOmniPortalFilterer(address common.Address, filterer bind.ContractFilterer) (*OmniPortalFilterer, error) {
	contract, err := bindOmniPortal(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OmniPortalFilterer{contract: contract}, nil
}

// bindOmniPortal binds a generic wrapper to an already deployed contract.
func bindOmniPortal(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OmniPortalMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OmniPortal *OmniPortalRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OmniPortal.Contract.OmniPortalCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OmniPortal *OmniPortalRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OmniPortal.Contract.OmniPortalTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OmniPortal *OmniPortalRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OmniPortal.Contract.OmniPortalTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OmniPortal *OmniPortalCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OmniPortal.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OmniPortal *OmniPortalTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OmniPortal.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OmniPortal *OmniPortalTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OmniPortal.Contract.contract.Transact(opts, method, params...)
}

// ActionXCall is a free data retrieval call binding the contract method 0xb2b2f5bd.
//
// Solidity: function ActionXCall() view returns(bytes32)
func (_OmniPortal *OmniPortalCaller) ActionXCall(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "ActionXCall")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ActionXCall is a free data retrieval call binding the contract method 0xb2b2f5bd.
//
// Solidity: function ActionXCall() view returns(bytes32)
func (_OmniPortal *OmniPortalSession) ActionXCall() ([32]byte, error) {
	return _OmniPortal.Contract.ActionXCall(&_OmniPortal.CallOpts)
}

// ActionXCall is a free data retrieval call binding the contract method 0xb2b2f5bd.
//
// Solidity: function ActionXCall() view returns(bytes32)
func (_OmniPortal *OmniPortalCallerSession) ActionXCall() ([32]byte, error) {
	return _OmniPortal.Contract.ActionXCall(&_OmniPortal.CallOpts)
}

// ActionXSubmit is a free data retrieval call binding the contract method 0xa32eb7c6.
//
// Solidity: function ActionXSubmit() view returns(bytes32)
func (_OmniPortal *OmniPortalCaller) ActionXSubmit(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "ActionXSubmit")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ActionXSubmit is a free data retrieval call binding the contract method 0xa32eb7c6.
//
// Solidity: function ActionXSubmit() view returns(bytes32)
func (_OmniPortal *OmniPortalSession) ActionXSubmit() ([32]byte, error) {
	return _OmniPortal.Contract.ActionXSubmit(&_OmniPortal.CallOpts)
}

// ActionXSubmit is a free data retrieval call binding the contract method 0xa32eb7c6.
//
// Solidity: function ActionXSubmit() view returns(bytes32)
func (_OmniPortal *OmniPortalCallerSession) ActionXSubmit() ([32]byte, error) {
	return _OmniPortal.Contract.ActionXSubmit(&_OmniPortal.CallOpts)
}

// KeyPauseAll is a free data retrieval call binding the contract method 0xa10ac97a.
//
// Solidity: function KeyPauseAll() view returns(bytes32)
func (_OmniPortal *OmniPortalCaller) KeyPauseAll(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "KeyPauseAll")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// KeyPauseAll is a free data retrieval call binding the contract method 0xa10ac97a.
//
// Solidity: function KeyPauseAll() view returns(bytes32)
func (_OmniPortal *OmniPortalSession) KeyPauseAll() ([32]byte, error) {
	return _OmniPortal.Contract.KeyPauseAll(&_OmniPortal.CallOpts)
}

// KeyPauseAll is a free data retrieval call binding the contract method 0xa10ac97a.
//
// Solidity: function KeyPauseAll() view returns(bytes32)
func (_OmniPortal *OmniPortalCallerSession) KeyPauseAll() ([32]byte, error) {
	return _OmniPortal.Contract.KeyPauseAll(&_OmniPortal.CallOpts)
}

// XSubQuorumDenominator is a free data retrieval call binding the contract method 0x06c3dc5f.
//
// Solidity: function XSubQuorumDenominator() view returns(uint8)
func (_OmniPortal *OmniPortalCaller) XSubQuorumDenominator(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "XSubQuorumDenominator")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// XSubQuorumDenominator is a free data retrieval call binding the contract method 0x06c3dc5f.
//
// Solidity: function XSubQuorumDenominator() view returns(uint8)
func (_OmniPortal *OmniPortalSession) XSubQuorumDenominator() (uint8, error) {
	return _OmniPortal.Contract.XSubQuorumDenominator(&_OmniPortal.CallOpts)
}

// XSubQuorumDenominator is a free data retrieval call binding the contract method 0x06c3dc5f.
//
// Solidity: function XSubQuorumDenominator() view returns(uint8)
func (_OmniPortal *OmniPortalCallerSession) XSubQuorumDenominator() (uint8, error) {
	return _OmniPortal.Contract.XSubQuorumDenominator(&_OmniPortal.CallOpts)
}

// XSubQuorumNumerator is a free data retrieval call binding the contract method 0x0360d20f.
//
// Solidity: function XSubQuorumNumerator() view returns(uint8)
func (_OmniPortal *OmniPortalCaller) XSubQuorumNumerator(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "XSubQuorumNumerator")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// XSubQuorumNumerator is a free data retrieval call binding the contract method 0x0360d20f.
//
// Solidity: function XSubQuorumNumerator() view returns(uint8)
func (_OmniPortal *OmniPortalSession) XSubQuorumNumerator() (uint8, error) {
	return _OmniPortal.Contract.XSubQuorumNumerator(&_OmniPortal.CallOpts)
}

// XSubQuorumNumerator is a free data retrieval call binding the contract method 0x0360d20f.
//
// Solidity: function XSubQuorumNumerator() view returns(uint8)
func (_OmniPortal *OmniPortalCallerSession) XSubQuorumNumerator() (uint8, error) {
	return _OmniPortal.Contract.XSubQuorumNumerator(&_OmniPortal.CallOpts)
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint64)
func (_OmniPortal *OmniPortalCaller) ChainId(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "chainId")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint64)
func (_OmniPortal *OmniPortalSession) ChainId() (uint64, error) {
	return _OmniPortal.Contract.ChainId(&_OmniPortal.CallOpts)
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint64)
func (_OmniPortal *OmniPortalCallerSession) ChainId() (uint64, error) {
	return _OmniPortal.Contract.ChainId(&_OmniPortal.CallOpts)
}

// FeeFor is a free data retrieval call binding the contract method 0x8dd9523c.
//
// Solidity: function feeFor(uint64 destChainId, bytes data, uint64 gasLimit) view returns(uint256)
func (_OmniPortal *OmniPortalCaller) FeeFor(opts *bind.CallOpts, destChainId uint64, data []byte, gasLimit uint64) (*big.Int, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "feeFor", destChainId, data, gasLimit)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeFor is a free data retrieval call binding the contract method 0x8dd9523c.
//
// Solidity: function feeFor(uint64 destChainId, bytes data, uint64 gasLimit) view returns(uint256)
func (_OmniPortal *OmniPortalSession) FeeFor(destChainId uint64, data []byte, gasLimit uint64) (*big.Int, error) {
	return _OmniPortal.Contract.FeeFor(&_OmniPortal.CallOpts, destChainId, data, gasLimit)
}

// FeeFor is a free data retrieval call binding the contract method 0x8dd9523c.
//
// Solidity: function feeFor(uint64 destChainId, bytes data, uint64 gasLimit) view returns(uint256)
func (_OmniPortal *OmniPortalCallerSession) FeeFor(destChainId uint64, data []byte, gasLimit uint64) (*big.Int, error) {
	return _OmniPortal.Contract.FeeFor(&_OmniPortal.CallOpts, destChainId, data, gasLimit)
}

// FeeOracle is a free data retrieval call binding the contract method 0x500b19e7.
//
// Solidity: function feeOracle() view returns(address)
func (_OmniPortal *OmniPortalCaller) FeeOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "feeOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeOracle is a free data retrieval call binding the contract method 0x500b19e7.
//
// Solidity: function feeOracle() view returns(address)
func (_OmniPortal *OmniPortalSession) FeeOracle() (common.Address, error) {
	return _OmniPortal.Contract.FeeOracle(&_OmniPortal.CallOpts)
}

// FeeOracle is a free data retrieval call binding the contract method 0x500b19e7.
//
// Solidity: function feeOracle() view returns(address)
func (_OmniPortal *OmniPortalCallerSession) FeeOracle() (common.Address, error) {
	return _OmniPortal.Contract.FeeOracle(&_OmniPortal.CallOpts)
}

// InXBlockOffset is a free data retrieval call binding the contract method 0x3fd3b15e.
//
// Solidity: function inXBlockOffset(uint64 , uint64 ) view returns(uint64)
func (_OmniPortal *OmniPortalCaller) InXBlockOffset(opts *bind.CallOpts, arg0 uint64, arg1 uint64) (uint64, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "inXBlockOffset", arg0, arg1)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// InXBlockOffset is a free data retrieval call binding the contract method 0x3fd3b15e.
//
// Solidity: function inXBlockOffset(uint64 , uint64 ) view returns(uint64)
func (_OmniPortal *OmniPortalSession) InXBlockOffset(arg0 uint64, arg1 uint64) (uint64, error) {
	return _OmniPortal.Contract.InXBlockOffset(&_OmniPortal.CallOpts, arg0, arg1)
}

// InXBlockOffset is a free data retrieval call binding the contract method 0x3fd3b15e.
//
// Solidity: function inXBlockOffset(uint64 , uint64 ) view returns(uint64)
func (_OmniPortal *OmniPortalCallerSession) InXBlockOffset(arg0 uint64, arg1 uint64) (uint64, error) {
	return _OmniPortal.Contract.InXBlockOffset(&_OmniPortal.CallOpts, arg0, arg1)
}

// InXMsgOffset is a free data retrieval call binding the contract method 0xd051c97d.
//
// Solidity: function inXMsgOffset(uint64 , uint64 ) view returns(uint64)
func (_OmniPortal *OmniPortalCaller) InXMsgOffset(opts *bind.CallOpts, arg0 uint64, arg1 uint64) (uint64, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "inXMsgOffset", arg0, arg1)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// InXMsgOffset is a free data retrieval call binding the contract method 0xd051c97d.
//
// Solidity: function inXMsgOffset(uint64 , uint64 ) view returns(uint64)
func (_OmniPortal *OmniPortalSession) InXMsgOffset(arg0 uint64, arg1 uint64) (uint64, error) {
	return _OmniPortal.Contract.InXMsgOffset(&_OmniPortal.CallOpts, arg0, arg1)
}

// InXMsgOffset is a free data retrieval call binding the contract method 0xd051c97d.
//
// Solidity: function inXMsgOffset(uint64 , uint64 ) view returns(uint64)
func (_OmniPortal *OmniPortalCallerSession) InXMsgOffset(arg0 uint64, arg1 uint64) (uint64, error) {
	return _OmniPortal.Contract.InXMsgOffset(&_OmniPortal.CallOpts, arg0, arg1)
}

// IsPaused is a free data retrieval call binding the contract method 0x241b71bb.
//
// Solidity: function isPaused(bytes32 actionId) view returns(bool)
func (_OmniPortal *OmniPortalCaller) IsPaused(opts *bind.CallOpts, actionId [32]byte) (bool, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "isPaused", actionId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPaused is a free data retrieval call binding the contract method 0x241b71bb.
//
// Solidity: function isPaused(bytes32 actionId) view returns(bool)
func (_OmniPortal *OmniPortalSession) IsPaused(actionId [32]byte) (bool, error) {
	return _OmniPortal.Contract.IsPaused(&_OmniPortal.CallOpts, actionId)
}

// IsPaused is a free data retrieval call binding the contract method 0x241b71bb.
//
// Solidity: function isPaused(bytes32 actionId) view returns(bool)
func (_OmniPortal *OmniPortalCallerSession) IsPaused(actionId [32]byte) (bool, error) {
	return _OmniPortal.Contract.IsPaused(&_OmniPortal.CallOpts, actionId)
}

// IsPaused0 is a free data retrieval call binding the contract method 0x461ab488.
//
// Solidity: function isPaused(bytes32 actionId, uint64 chainId_) view returns(bool)
func (_OmniPortal *OmniPortalCaller) IsPaused0(opts *bind.CallOpts, actionId [32]byte, chainId_ uint64) (bool, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "isPaused0", actionId, chainId_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPaused0 is a free data retrieval call binding the contract method 0x461ab488.
//
// Solidity: function isPaused(bytes32 actionId, uint64 chainId_) view returns(bool)
func (_OmniPortal *OmniPortalSession) IsPaused0(actionId [32]byte, chainId_ uint64) (bool, error) {
	return _OmniPortal.Contract.IsPaused0(&_OmniPortal.CallOpts, actionId, chainId_)
}

// IsPaused0 is a free data retrieval call binding the contract method 0x461ab488.
//
// Solidity: function isPaused(bytes32 actionId, uint64 chainId_) view returns(bool)
func (_OmniPortal *OmniPortalCallerSession) IsPaused0(actionId [32]byte, chainId_ uint64) (bool, error) {
	return _OmniPortal.Contract.IsPaused0(&_OmniPortal.CallOpts, actionId, chainId_)
}

// IsPaused1 is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_OmniPortal *OmniPortalCaller) IsPaused1(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "isPaused1")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPaused1 is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_OmniPortal *OmniPortalSession) IsPaused1() (bool, error) {
	return _OmniPortal.Contract.IsPaused1(&_OmniPortal.CallOpts)
}

// IsPaused1 is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_OmniPortal *OmniPortalCallerSession) IsPaused1() (bool, error) {
	return _OmniPortal.Contract.IsPaused1(&_OmniPortal.CallOpts)
}

// IsSupportedDest is a free data retrieval call binding the contract method 0x24278bbe.
//
// Solidity: function isSupportedDest(uint64 ) view returns(bool)
func (_OmniPortal *OmniPortalCaller) IsSupportedDest(opts *bind.CallOpts, arg0 uint64) (bool, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "isSupportedDest", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSupportedDest is a free data retrieval call binding the contract method 0x24278bbe.
//
// Solidity: function isSupportedDest(uint64 ) view returns(bool)
func (_OmniPortal *OmniPortalSession) IsSupportedDest(arg0 uint64) (bool, error) {
	return _OmniPortal.Contract.IsSupportedDest(&_OmniPortal.CallOpts, arg0)
}

// IsSupportedDest is a free data retrieval call binding the contract method 0x24278bbe.
//
// Solidity: function isSupportedDest(uint64 ) view returns(bool)
func (_OmniPortal *OmniPortalCallerSession) IsSupportedDest(arg0 uint64) (bool, error) {
	return _OmniPortal.Contract.IsSupportedDest(&_OmniPortal.CallOpts, arg0)
}

// IsSupportedShard is a free data retrieval call binding the contract method 0xaaf1bc97.
//
// Solidity: function isSupportedShard(uint64 ) view returns(bool)
func (_OmniPortal *OmniPortalCaller) IsSupportedShard(opts *bind.CallOpts, arg0 uint64) (bool, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "isSupportedShard", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSupportedShard is a free data retrieval call binding the contract method 0xaaf1bc97.
//
// Solidity: function isSupportedShard(uint64 ) view returns(bool)
func (_OmniPortal *OmniPortalSession) IsSupportedShard(arg0 uint64) (bool, error) {
	return _OmniPortal.Contract.IsSupportedShard(&_OmniPortal.CallOpts, arg0)
}

// IsSupportedShard is a free data retrieval call binding the contract method 0xaaf1bc97.
//
// Solidity: function isSupportedShard(uint64 ) view returns(bool)
func (_OmniPortal *OmniPortalCallerSession) IsSupportedShard(arg0 uint64) (bool, error) {
	return _OmniPortal.Contract.IsSupportedShard(&_OmniPortal.CallOpts, arg0)
}

// IsXCall is a free data retrieval call binding the contract method 0x55e2448e.
//
// Solidity: function isXCall() view returns(bool)
func (_OmniPortal *OmniPortalCaller) IsXCall(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "isXCall")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsXCall is a free data retrieval call binding the contract method 0x55e2448e.
//
// Solidity: function isXCall() view returns(bool)
func (_OmniPortal *OmniPortalSession) IsXCall() (bool, error) {
	return _OmniPortal.Contract.IsXCall(&_OmniPortal.CallOpts)
}

// IsXCall is a free data retrieval call binding the contract method 0x55e2448e.
//
// Solidity: function isXCall() view returns(bool)
func (_OmniPortal *OmniPortalCallerSession) IsXCall() (bool, error) {
	return _OmniPortal.Contract.IsXCall(&_OmniPortal.CallOpts)
}

// LatestValSetId is a free data retrieval call binding the contract method 0xf45cc7b8.
//
// Solidity: function latestValSetId() view returns(uint64)
func (_OmniPortal *OmniPortalCaller) LatestValSetId(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "latestValSetId")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LatestValSetId is a free data retrieval call binding the contract method 0xf45cc7b8.
//
// Solidity: function latestValSetId() view returns(uint64)
func (_OmniPortal *OmniPortalSession) LatestValSetId() (uint64, error) {
	return _OmniPortal.Contract.LatestValSetId(&_OmniPortal.CallOpts)
}

// LatestValSetId is a free data retrieval call binding the contract method 0xf45cc7b8.
//
// Solidity: function latestValSetId() view returns(uint64)
func (_OmniPortal *OmniPortalCallerSession) LatestValSetId() (uint64, error) {
	return _OmniPortal.Contract.LatestValSetId(&_OmniPortal.CallOpts)
}

// Network is a free data retrieval call binding the contract method 0x74eba939.
//
// Solidity: function network(uint256 ) view returns(uint64 chainId)
func (_OmniPortal *OmniPortalCaller) Network(opts *bind.CallOpts, arg0 *big.Int) (uint64, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "network", arg0)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// Network is a free data retrieval call binding the contract method 0x74eba939.
//
// Solidity: function network(uint256 ) view returns(uint64 chainId)
func (_OmniPortal *OmniPortalSession) Network(arg0 *big.Int) (uint64, error) {
	return _OmniPortal.Contract.Network(&_OmniPortal.CallOpts, arg0)
}

// Network is a free data retrieval call binding the contract method 0x74eba939.
//
// Solidity: function network(uint256 ) view returns(uint64 chainId)
func (_OmniPortal *OmniPortalCallerSession) Network(arg0 *big.Int) (uint64, error) {
	return _OmniPortal.Contract.Network(&_OmniPortal.CallOpts, arg0)
}

// OmniCChainId is a free data retrieval call binding the contract method 0x36d21912.
//
// Solidity: function omniCChainId() view returns(uint64)
func (_OmniPortal *OmniPortalCaller) OmniCChainId(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "omniCChainId")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// OmniCChainId is a free data retrieval call binding the contract method 0x36d21912.
//
// Solidity: function omniCChainId() view returns(uint64)
func (_OmniPortal *OmniPortalSession) OmniCChainId() (uint64, error) {
	return _OmniPortal.Contract.OmniCChainId(&_OmniPortal.CallOpts)
}

// OmniCChainId is a free data retrieval call binding the contract method 0x36d21912.
//
// Solidity: function omniCChainId() view returns(uint64)
func (_OmniPortal *OmniPortalCallerSession) OmniCChainId() (uint64, error) {
	return _OmniPortal.Contract.OmniCChainId(&_OmniPortal.CallOpts)
}

// OmniChainId is a free data retrieval call binding the contract method 0x110ff5f1.
//
// Solidity: function omniChainId() view returns(uint64)
func (_OmniPortal *OmniPortalCaller) OmniChainId(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "omniChainId")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// OmniChainId is a free data retrieval call binding the contract method 0x110ff5f1.
//
// Solidity: function omniChainId() view returns(uint64)
func (_OmniPortal *OmniPortalSession) OmniChainId() (uint64, error) {
	return _OmniPortal.Contract.OmniChainId(&_OmniPortal.CallOpts)
}

// OmniChainId is a free data retrieval call binding the contract method 0x110ff5f1.
//
// Solidity: function omniChainId() view returns(uint64)
func (_OmniPortal *OmniPortalCallerSession) OmniChainId() (uint64, error) {
	return _OmniPortal.Contract.OmniChainId(&_OmniPortal.CallOpts)
}

// OutXMsgOffset is a free data retrieval call binding the contract method 0x3aa87330.
//
// Solidity: function outXMsgOffset(uint64 , uint64 ) view returns(uint64)
func (_OmniPortal *OmniPortalCaller) OutXMsgOffset(opts *bind.CallOpts, arg0 uint64, arg1 uint64) (uint64, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "outXMsgOffset", arg0, arg1)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// OutXMsgOffset is a free data retrieval call binding the contract method 0x3aa87330.
//
// Solidity: function outXMsgOffset(uint64 , uint64 ) view returns(uint64)
func (_OmniPortal *OmniPortalSession) OutXMsgOffset(arg0 uint64, arg1 uint64) (uint64, error) {
	return _OmniPortal.Contract.OutXMsgOffset(&_OmniPortal.CallOpts, arg0, arg1)
}

// OutXMsgOffset is a free data retrieval call binding the contract method 0x3aa87330.
//
// Solidity: function outXMsgOffset(uint64 , uint64 ) view returns(uint64)
func (_OmniPortal *OmniPortalCallerSession) OutXMsgOffset(arg0 uint64, arg1 uint64) (uint64, error) {
	return _OmniPortal.Contract.OutXMsgOffset(&_OmniPortal.CallOpts, arg0, arg1)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OmniPortal *OmniPortalCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OmniPortal *OmniPortalSession) Owner() (common.Address, error) {
	return _OmniPortal.Contract.Owner(&_OmniPortal.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OmniPortal *OmniPortalCallerSession) Owner() (common.Address, error) {
	return _OmniPortal.Contract.Owner(&_OmniPortal.CallOpts)
}

// ValSet is a free data retrieval call binding the contract method 0x57542050.
//
// Solidity: function valSet(uint64 , address ) view returns(uint64)
func (_OmniPortal *OmniPortalCaller) ValSet(opts *bind.CallOpts, arg0 uint64, arg1 common.Address) (uint64, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "valSet", arg0, arg1)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ValSet is a free data retrieval call binding the contract method 0x57542050.
//
// Solidity: function valSet(uint64 , address ) view returns(uint64)
func (_OmniPortal *OmniPortalSession) ValSet(arg0 uint64, arg1 common.Address) (uint64, error) {
	return _OmniPortal.Contract.ValSet(&_OmniPortal.CallOpts, arg0, arg1)
}

// ValSet is a free data retrieval call binding the contract method 0x57542050.
//
// Solidity: function valSet(uint64 , address ) view returns(uint64)
func (_OmniPortal *OmniPortalCallerSession) ValSet(arg0 uint64, arg1 common.Address) (uint64, error) {
	return _OmniPortal.Contract.ValSet(&_OmniPortal.CallOpts, arg0, arg1)
}

// ValSetTotalPower is a free data retrieval call binding the contract method 0xafe8af9c.
//
// Solidity: function valSetTotalPower(uint64 ) view returns(uint64)
func (_OmniPortal *OmniPortalCaller) ValSetTotalPower(opts *bind.CallOpts, arg0 uint64) (uint64, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "valSetTotalPower", arg0)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ValSetTotalPower is a free data retrieval call binding the contract method 0xafe8af9c.
//
// Solidity: function valSetTotalPower(uint64 ) view returns(uint64)
func (_OmniPortal *OmniPortalSession) ValSetTotalPower(arg0 uint64) (uint64, error) {
	return _OmniPortal.Contract.ValSetTotalPower(&_OmniPortal.CallOpts, arg0)
}

// ValSetTotalPower is a free data retrieval call binding the contract method 0xafe8af9c.
//
// Solidity: function valSetTotalPower(uint64 ) view returns(uint64)
func (_OmniPortal *OmniPortalCallerSession) ValSetTotalPower(arg0 uint64) (uint64, error) {
	return _OmniPortal.Contract.ValSetTotalPower(&_OmniPortal.CallOpts, arg0)
}

// Xmsg is a free data retrieval call binding the contract method 0x2f32700e.
//
// Solidity: function xmsg() view returns((uint64,address))
func (_OmniPortal *OmniPortalCaller) Xmsg(opts *bind.CallOpts) (XTypesMsgContext, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "xmsg")

	if err != nil {
		return *new(XTypesMsgContext), err
	}

	out0 := *abi.ConvertType(out[0], new(XTypesMsgContext)).(*XTypesMsgContext)

	return out0, err

}

// Xmsg is a free data retrieval call binding the contract method 0x2f32700e.
//
// Solidity: function xmsg() view returns((uint64,address))
func (_OmniPortal *OmniPortalSession) Xmsg() (XTypesMsgContext, error) {
	return _OmniPortal.Contract.Xmsg(&_OmniPortal.CallOpts)
}

// Xmsg is a free data retrieval call binding the contract method 0x2f32700e.
//
// Solidity: function xmsg() view returns((uint64,address))
func (_OmniPortal *OmniPortalCallerSession) Xmsg() (XTypesMsgContext, error) {
	return _OmniPortal.Contract.Xmsg(&_OmniPortal.CallOpts)
}

// XmsgMaxDataSize is a free data retrieval call binding the contract method 0xb4d5afd1.
//
// Solidity: function xmsgMaxDataSize() view returns(uint16)
func (_OmniPortal *OmniPortalCaller) XmsgMaxDataSize(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "xmsgMaxDataSize")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// XmsgMaxDataSize is a free data retrieval call binding the contract method 0xb4d5afd1.
//
// Solidity: function xmsgMaxDataSize() view returns(uint16)
func (_OmniPortal *OmniPortalSession) XmsgMaxDataSize() (uint16, error) {
	return _OmniPortal.Contract.XmsgMaxDataSize(&_OmniPortal.CallOpts)
}

// XmsgMaxDataSize is a free data retrieval call binding the contract method 0xb4d5afd1.
//
// Solidity: function xmsgMaxDataSize() view returns(uint16)
func (_OmniPortal *OmniPortalCallerSession) XmsgMaxDataSize() (uint16, error) {
	return _OmniPortal.Contract.XmsgMaxDataSize(&_OmniPortal.CallOpts)
}

// XmsgMaxGasLimit is a free data retrieval call binding the contract method 0xcf84c818.
//
// Solidity: function xmsgMaxGasLimit() view returns(uint64)
func (_OmniPortal *OmniPortalCaller) XmsgMaxGasLimit(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "xmsgMaxGasLimit")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// XmsgMaxGasLimit is a free data retrieval call binding the contract method 0xcf84c818.
//
// Solidity: function xmsgMaxGasLimit() view returns(uint64)
func (_OmniPortal *OmniPortalSession) XmsgMaxGasLimit() (uint64, error) {
	return _OmniPortal.Contract.XmsgMaxGasLimit(&_OmniPortal.CallOpts)
}

// XmsgMaxGasLimit is a free data retrieval call binding the contract method 0xcf84c818.
//
// Solidity: function xmsgMaxGasLimit() view returns(uint64)
func (_OmniPortal *OmniPortalCallerSession) XmsgMaxGasLimit() (uint64, error) {
	return _OmniPortal.Contract.XmsgMaxGasLimit(&_OmniPortal.CallOpts)
}

// XmsgMinGasLimit is a free data retrieval call binding the contract method 0x78fe5307.
//
// Solidity: function xmsgMinGasLimit() view returns(uint64)
func (_OmniPortal *OmniPortalCaller) XmsgMinGasLimit(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "xmsgMinGasLimit")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// XmsgMinGasLimit is a free data retrieval call binding the contract method 0x78fe5307.
//
// Solidity: function xmsgMinGasLimit() view returns(uint64)
func (_OmniPortal *OmniPortalSession) XmsgMinGasLimit() (uint64, error) {
	return _OmniPortal.Contract.XmsgMinGasLimit(&_OmniPortal.CallOpts)
}

// XmsgMinGasLimit is a free data retrieval call binding the contract method 0x78fe5307.
//
// Solidity: function xmsgMinGasLimit() view returns(uint64)
func (_OmniPortal *OmniPortalCallerSession) XmsgMinGasLimit() (uint64, error) {
	return _OmniPortal.Contract.XmsgMinGasLimit(&_OmniPortal.CallOpts)
}

// XreceiptMaxErrorSize is a free data retrieval call binding the contract method 0xc26dfc05.
//
// Solidity: function xreceiptMaxErrorSize() view returns(uint16)
func (_OmniPortal *OmniPortalCaller) XreceiptMaxErrorSize(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "xreceiptMaxErrorSize")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// XreceiptMaxErrorSize is a free data retrieval call binding the contract method 0xc26dfc05.
//
// Solidity: function xreceiptMaxErrorSize() view returns(uint16)
func (_OmniPortal *OmniPortalSession) XreceiptMaxErrorSize() (uint16, error) {
	return _OmniPortal.Contract.XreceiptMaxErrorSize(&_OmniPortal.CallOpts)
}

// XreceiptMaxErrorSize is a free data retrieval call binding the contract method 0xc26dfc05.
//
// Solidity: function xreceiptMaxErrorSize() view returns(uint16)
func (_OmniPortal *OmniPortalCallerSession) XreceiptMaxErrorSize() (uint16, error) {
	return _OmniPortal.Contract.XreceiptMaxErrorSize(&_OmniPortal.CallOpts)
}

// XsubValsetCutoff is a free data retrieval call binding the contract method 0x49cc3bf6.
//
// Solidity: function xsubValsetCutoff() view returns(uint8)
func (_OmniPortal *OmniPortalCaller) XsubValsetCutoff(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "xsubValsetCutoff")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// XsubValsetCutoff is a free data retrieval call binding the contract method 0x49cc3bf6.
//
// Solidity: function xsubValsetCutoff() view returns(uint8)
func (_OmniPortal *OmniPortalSession) XsubValsetCutoff() (uint8, error) {
	return _OmniPortal.Contract.XsubValsetCutoff(&_OmniPortal.CallOpts)
}

// XsubValsetCutoff is a free data retrieval call binding the contract method 0x49cc3bf6.
//
// Solidity: function xsubValsetCutoff() view returns(uint8)
func (_OmniPortal *OmniPortalCallerSession) XsubValsetCutoff() (uint8, error) {
	return _OmniPortal.Contract.XsubValsetCutoff(&_OmniPortal.CallOpts)
}

// AddValidatorSet is a paid mutator transaction binding the contract method 0x8532eb9f.
//
// Solidity: function addValidatorSet(uint64 valSetId, (address,uint64)[] validators) returns()
func (_OmniPortal *OmniPortalTransactor) AddValidatorSet(opts *bind.TransactOpts, valSetId uint64, validators []XTypesValidator) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "addValidatorSet", valSetId, validators)
}

// AddValidatorSet is a paid mutator transaction binding the contract method 0x8532eb9f.
//
// Solidity: function addValidatorSet(uint64 valSetId, (address,uint64)[] validators) returns()
func (_OmniPortal *OmniPortalSession) AddValidatorSet(valSetId uint64, validators []XTypesValidator) (*types.Transaction, error) {
	return _OmniPortal.Contract.AddValidatorSet(&_OmniPortal.TransactOpts, valSetId, validators)
}

// AddValidatorSet is a paid mutator transaction binding the contract method 0x8532eb9f.
//
// Solidity: function addValidatorSet(uint64 valSetId, (address,uint64)[] validators) returns()
func (_OmniPortal *OmniPortalTransactorSession) AddValidatorSet(valSetId uint64, validators []XTypesValidator) (*types.Transaction, error) {
	return _OmniPortal.Contract.AddValidatorSet(&_OmniPortal.TransactOpts, valSetId, validators)
}

// CollectFees is a paid mutator transaction binding the contract method 0xa480ca79.
//
// Solidity: function collectFees(address to) returns()
func (_OmniPortal *OmniPortalTransactor) CollectFees(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "collectFees", to)
}

// CollectFees is a paid mutator transaction binding the contract method 0xa480ca79.
//
// Solidity: function collectFees(address to) returns()
func (_OmniPortal *OmniPortalSession) CollectFees(to common.Address) (*types.Transaction, error) {
	return _OmniPortal.Contract.CollectFees(&_OmniPortal.TransactOpts, to)
}

// CollectFees is a paid mutator transaction binding the contract method 0xa480ca79.
//
// Solidity: function collectFees(address to) returns()
func (_OmniPortal *OmniPortalTransactorSession) CollectFees(to common.Address) (*types.Transaction, error) {
	return _OmniPortal.Contract.CollectFees(&_OmniPortal.TransactOpts, to)
}

// Initialize is a paid mutator transaction binding the contract method 0x30632e8b.
//
// Solidity: function initialize((address,address,uint64,uint64,uint64,uint64,uint16,uint16,uint8,uint64,uint64,uint64,(address,uint64)[]) p) returns()
func (_OmniPortal *OmniPortalTransactor) Initialize(opts *bind.TransactOpts, p OmniPortalInitParams) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "initialize", p)
}

// Initialize is a paid mutator transaction binding the contract method 0x30632e8b.
//
// Solidity: function initialize((address,address,uint64,uint64,uint64,uint64,uint16,uint16,uint8,uint64,uint64,uint64,(address,uint64)[]) p) returns()
func (_OmniPortal *OmniPortalSession) Initialize(p OmniPortalInitParams) (*types.Transaction, error) {
	return _OmniPortal.Contract.Initialize(&_OmniPortal.TransactOpts, p)
}

// Initialize is a paid mutator transaction binding the contract method 0x30632e8b.
//
// Solidity: function initialize((address,address,uint64,uint64,uint64,uint64,uint16,uint16,uint8,uint64,uint64,uint64,(address,uint64)[]) p) returns()
func (_OmniPortal *OmniPortalTransactorSession) Initialize(p OmniPortalInitParams) (*types.Transaction, error) {
	return _OmniPortal.Contract.Initialize(&_OmniPortal.TransactOpts, p)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_OmniPortal *OmniPortalTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_OmniPortal *OmniPortalSession) Pause() (*types.Transaction, error) {
	return _OmniPortal.Contract.Pause(&_OmniPortal.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_OmniPortal *OmniPortalTransactorSession) Pause() (*types.Transaction, error) {
	return _OmniPortal.Contract.Pause(&_OmniPortal.TransactOpts)
}

// PauseXCall is a paid mutator transaction binding the contract method 0x83d0cbd9.
//
// Solidity: function pauseXCall() returns()
func (_OmniPortal *OmniPortalTransactor) PauseXCall(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "pauseXCall")
}

// PauseXCall is a paid mutator transaction binding the contract method 0x83d0cbd9.
//
// Solidity: function pauseXCall() returns()
func (_OmniPortal *OmniPortalSession) PauseXCall() (*types.Transaction, error) {
	return _OmniPortal.Contract.PauseXCall(&_OmniPortal.TransactOpts)
}

// PauseXCall is a paid mutator transaction binding the contract method 0x83d0cbd9.
//
// Solidity: function pauseXCall() returns()
func (_OmniPortal *OmniPortalTransactorSession) PauseXCall() (*types.Transaction, error) {
	return _OmniPortal.Contract.PauseXCall(&_OmniPortal.TransactOpts)
}

// PauseXCallTo is a paid mutator transaction binding the contract method 0x10a5a7f7.
//
// Solidity: function pauseXCallTo(uint64 chainId_) returns()
func (_OmniPortal *OmniPortalTransactor) PauseXCallTo(opts *bind.TransactOpts, chainId_ uint64) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "pauseXCallTo", chainId_)
}

// PauseXCallTo is a paid mutator transaction binding the contract method 0x10a5a7f7.
//
// Solidity: function pauseXCallTo(uint64 chainId_) returns()
func (_OmniPortal *OmniPortalSession) PauseXCallTo(chainId_ uint64) (*types.Transaction, error) {
	return _OmniPortal.Contract.PauseXCallTo(&_OmniPortal.TransactOpts, chainId_)
}

// PauseXCallTo is a paid mutator transaction binding the contract method 0x10a5a7f7.
//
// Solidity: function pauseXCallTo(uint64 chainId_) returns()
func (_OmniPortal *OmniPortalTransactorSession) PauseXCallTo(chainId_ uint64) (*types.Transaction, error) {
	return _OmniPortal.Contract.PauseXCallTo(&_OmniPortal.TransactOpts, chainId_)
}

// PauseXSubmit is a paid mutator transaction binding the contract method 0x23dbce50.
//
// Solidity: function pauseXSubmit() returns()
func (_OmniPortal *OmniPortalTransactor) PauseXSubmit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "pauseXSubmit")
}

// PauseXSubmit is a paid mutator transaction binding the contract method 0x23dbce50.
//
// Solidity: function pauseXSubmit() returns()
func (_OmniPortal *OmniPortalSession) PauseXSubmit() (*types.Transaction, error) {
	return _OmniPortal.Contract.PauseXSubmit(&_OmniPortal.TransactOpts)
}

// PauseXSubmit is a paid mutator transaction binding the contract method 0x23dbce50.
//
// Solidity: function pauseXSubmit() returns()
func (_OmniPortal *OmniPortalTransactorSession) PauseXSubmit() (*types.Transaction, error) {
	return _OmniPortal.Contract.PauseXSubmit(&_OmniPortal.TransactOpts)
}

// PauseXSubmitFrom is a paid mutator transaction binding the contract method 0xafe82198.
//
// Solidity: function pauseXSubmitFrom(uint64 chainId_) returns()
func (_OmniPortal *OmniPortalTransactor) PauseXSubmitFrom(opts *bind.TransactOpts, chainId_ uint64) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "pauseXSubmitFrom", chainId_)
}

// PauseXSubmitFrom is a paid mutator transaction binding the contract method 0xafe82198.
//
// Solidity: function pauseXSubmitFrom(uint64 chainId_) returns()
func (_OmniPortal *OmniPortalSession) PauseXSubmitFrom(chainId_ uint64) (*types.Transaction, error) {
	return _OmniPortal.Contract.PauseXSubmitFrom(&_OmniPortal.TransactOpts, chainId_)
}

// PauseXSubmitFrom is a paid mutator transaction binding the contract method 0xafe82198.
//
// Solidity: function pauseXSubmitFrom(uint64 chainId_) returns()
func (_OmniPortal *OmniPortalTransactorSession) PauseXSubmitFrom(chainId_ uint64) (*types.Transaction, error) {
	return _OmniPortal.Contract.PauseXSubmitFrom(&_OmniPortal.TransactOpts, chainId_)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OmniPortal *OmniPortalTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OmniPortal *OmniPortalSession) RenounceOwnership() (*types.Transaction, error) {
	return _OmniPortal.Contract.RenounceOwnership(&_OmniPortal.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OmniPortal *OmniPortalTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _OmniPortal.Contract.RenounceOwnership(&_OmniPortal.TransactOpts)
}

// SetFeeOracle is a paid mutator transaction binding the contract method 0xa8a98962.
//
// Solidity: function setFeeOracle(address feeOracle_) returns()
func (_OmniPortal *OmniPortalTransactor) SetFeeOracle(opts *bind.TransactOpts, feeOracle_ common.Address) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "setFeeOracle", feeOracle_)
}

// SetFeeOracle is a paid mutator transaction binding the contract method 0xa8a98962.
//
// Solidity: function setFeeOracle(address feeOracle_) returns()
func (_OmniPortal *OmniPortalSession) SetFeeOracle(feeOracle_ common.Address) (*types.Transaction, error) {
	return _OmniPortal.Contract.SetFeeOracle(&_OmniPortal.TransactOpts, feeOracle_)
}

// SetFeeOracle is a paid mutator transaction binding the contract method 0xa8a98962.
//
// Solidity: function setFeeOracle(address feeOracle_) returns()
func (_OmniPortal *OmniPortalTransactorSession) SetFeeOracle(feeOracle_ common.Address) (*types.Transaction, error) {
	return _OmniPortal.Contract.SetFeeOracle(&_OmniPortal.TransactOpts, feeOracle_)
}

// SetInXBlockOffset is a paid mutator transaction binding the contract method 0x97b52062.
//
// Solidity: function setInXBlockOffset(uint64 sourceChainId, uint64 shardId, uint64 offset) returns()
func (_OmniPortal *OmniPortalTransactor) SetInXBlockOffset(opts *bind.TransactOpts, sourceChainId uint64, shardId uint64, offset uint64) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "setInXBlockOffset", sourceChainId, shardId, offset)
}

// SetInXBlockOffset is a paid mutator transaction binding the contract method 0x97b52062.
//
// Solidity: function setInXBlockOffset(uint64 sourceChainId, uint64 shardId, uint64 offset) returns()
func (_OmniPortal *OmniPortalSession) SetInXBlockOffset(sourceChainId uint64, shardId uint64, offset uint64) (*types.Transaction, error) {
	return _OmniPortal.Contract.SetInXBlockOffset(&_OmniPortal.TransactOpts, sourceChainId, shardId, offset)
}

// SetInXBlockOffset is a paid mutator transaction binding the contract method 0x97b52062.
//
// Solidity: function setInXBlockOffset(uint64 sourceChainId, uint64 shardId, uint64 offset) returns()
func (_OmniPortal *OmniPortalTransactorSession) SetInXBlockOffset(sourceChainId uint64, shardId uint64, offset uint64) (*types.Transaction, error) {
	return _OmniPortal.Contract.SetInXBlockOffset(&_OmniPortal.TransactOpts, sourceChainId, shardId, offset)
}

// SetInXMsgOffset is a paid mutator transaction binding the contract method 0xc4ab80bc.
//
// Solidity: function setInXMsgOffset(uint64 sourceChainId, uint64 shardId, uint64 offset) returns()
func (_OmniPortal *OmniPortalTransactor) SetInXMsgOffset(opts *bind.TransactOpts, sourceChainId uint64, shardId uint64, offset uint64) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "setInXMsgOffset", sourceChainId, shardId, offset)
}

// SetInXMsgOffset is a paid mutator transaction binding the contract method 0xc4ab80bc.
//
// Solidity: function setInXMsgOffset(uint64 sourceChainId, uint64 shardId, uint64 offset) returns()
func (_OmniPortal *OmniPortalSession) SetInXMsgOffset(sourceChainId uint64, shardId uint64, offset uint64) (*types.Transaction, error) {
	return _OmniPortal.Contract.SetInXMsgOffset(&_OmniPortal.TransactOpts, sourceChainId, shardId, offset)
}

// SetInXMsgOffset is a paid mutator transaction binding the contract method 0xc4ab80bc.
//
// Solidity: function setInXMsgOffset(uint64 sourceChainId, uint64 shardId, uint64 offset) returns()
func (_OmniPortal *OmniPortalTransactorSession) SetInXMsgOffset(sourceChainId uint64, shardId uint64, offset uint64) (*types.Transaction, error) {
	return _OmniPortal.Contract.SetInXMsgOffset(&_OmniPortal.TransactOpts, sourceChainId, shardId, offset)
}

// SetNetwork is a paid mutator transaction binding the contract method 0x1d3eb6e3.
//
// Solidity: function setNetwork((uint64,uint64[])[] network_) returns()
func (_OmniPortal *OmniPortalTransactor) SetNetwork(opts *bind.TransactOpts, network_ []XTypesChain) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "setNetwork", network_)
}

// SetNetwork is a paid mutator transaction binding the contract method 0x1d3eb6e3.
//
// Solidity: function setNetwork((uint64,uint64[])[] network_) returns()
func (_OmniPortal *OmniPortalSession) SetNetwork(network_ []XTypesChain) (*types.Transaction, error) {
	return _OmniPortal.Contract.SetNetwork(&_OmniPortal.TransactOpts, network_)
}

// SetNetwork is a paid mutator transaction binding the contract method 0x1d3eb6e3.
//
// Solidity: function setNetwork((uint64,uint64[])[] network_) returns()
func (_OmniPortal *OmniPortalTransactorSession) SetNetwork(network_ []XTypesChain) (*types.Transaction, error) {
	return _OmniPortal.Contract.SetNetwork(&_OmniPortal.TransactOpts, network_)
}

// SetXMsgMaxDataSize is a paid mutator transaction binding the contract method 0xb521466d.
//
// Solidity: function setXMsgMaxDataSize(uint16 numBytes) returns()
func (_OmniPortal *OmniPortalTransactor) SetXMsgMaxDataSize(opts *bind.TransactOpts, numBytes uint16) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "setXMsgMaxDataSize", numBytes)
}

// SetXMsgMaxDataSize is a paid mutator transaction binding the contract method 0xb521466d.
//
// Solidity: function setXMsgMaxDataSize(uint16 numBytes) returns()
func (_OmniPortal *OmniPortalSession) SetXMsgMaxDataSize(numBytes uint16) (*types.Transaction, error) {
	return _OmniPortal.Contract.SetXMsgMaxDataSize(&_OmniPortal.TransactOpts, numBytes)
}

// SetXMsgMaxDataSize is a paid mutator transaction binding the contract method 0xb521466d.
//
// Solidity: function setXMsgMaxDataSize(uint16 numBytes) returns()
func (_OmniPortal *OmniPortalTransactorSession) SetXMsgMaxDataSize(numBytes uint16) (*types.Transaction, error) {
	return _OmniPortal.Contract.SetXMsgMaxDataSize(&_OmniPortal.TransactOpts, numBytes)
}

// SetXMsgMaxGasLimit is a paid mutator transaction binding the contract method 0x36d853f9.
//
// Solidity: function setXMsgMaxGasLimit(uint64 gasLimit) returns()
func (_OmniPortal *OmniPortalTransactor) SetXMsgMaxGasLimit(opts *bind.TransactOpts, gasLimit uint64) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "setXMsgMaxGasLimit", gasLimit)
}

// SetXMsgMaxGasLimit is a paid mutator transaction binding the contract method 0x36d853f9.
//
// Solidity: function setXMsgMaxGasLimit(uint64 gasLimit) returns()
func (_OmniPortal *OmniPortalSession) SetXMsgMaxGasLimit(gasLimit uint64) (*types.Transaction, error) {
	return _OmniPortal.Contract.SetXMsgMaxGasLimit(&_OmniPortal.TransactOpts, gasLimit)
}

// SetXMsgMaxGasLimit is a paid mutator transaction binding the contract method 0x36d853f9.
//
// Solidity: function setXMsgMaxGasLimit(uint64 gasLimit) returns()
func (_OmniPortal *OmniPortalTransactorSession) SetXMsgMaxGasLimit(gasLimit uint64) (*types.Transaction, error) {
	return _OmniPortal.Contract.SetXMsgMaxGasLimit(&_OmniPortal.TransactOpts, gasLimit)
}

// SetXMsgMinGasLimit is a paid mutator transaction binding the contract method 0xbb8590ad.
//
// Solidity: function setXMsgMinGasLimit(uint64 gasLimit) returns()
func (_OmniPortal *OmniPortalTransactor) SetXMsgMinGasLimit(opts *bind.TransactOpts, gasLimit uint64) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "setXMsgMinGasLimit", gasLimit)
}

// SetXMsgMinGasLimit is a paid mutator transaction binding the contract method 0xbb8590ad.
//
// Solidity: function setXMsgMinGasLimit(uint64 gasLimit) returns()
func (_OmniPortal *OmniPortalSession) SetXMsgMinGasLimit(gasLimit uint64) (*types.Transaction, error) {
	return _OmniPortal.Contract.SetXMsgMinGasLimit(&_OmniPortal.TransactOpts, gasLimit)
}

// SetXMsgMinGasLimit is a paid mutator transaction binding the contract method 0xbb8590ad.
//
// Solidity: function setXMsgMinGasLimit(uint64 gasLimit) returns()
func (_OmniPortal *OmniPortalTransactorSession) SetXMsgMinGasLimit(gasLimit uint64) (*types.Transaction, error) {
	return _OmniPortal.Contract.SetXMsgMinGasLimit(&_OmniPortal.TransactOpts, gasLimit)
}

// SetXReceiptMaxErrorSize is a paid mutator transaction binding the contract method 0xbff0e84d.
//
// Solidity: function setXReceiptMaxErrorSize(uint16 numBytes) returns()
func (_OmniPortal *OmniPortalTransactor) SetXReceiptMaxErrorSize(opts *bind.TransactOpts, numBytes uint16) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "setXReceiptMaxErrorSize", numBytes)
}

// SetXReceiptMaxErrorSize is a paid mutator transaction binding the contract method 0xbff0e84d.
//
// Solidity: function setXReceiptMaxErrorSize(uint16 numBytes) returns()
func (_OmniPortal *OmniPortalSession) SetXReceiptMaxErrorSize(numBytes uint16) (*types.Transaction, error) {
	return _OmniPortal.Contract.SetXReceiptMaxErrorSize(&_OmniPortal.TransactOpts, numBytes)
}

// SetXReceiptMaxErrorSize is a paid mutator transaction binding the contract method 0xbff0e84d.
//
// Solidity: function setXReceiptMaxErrorSize(uint16 numBytes) returns()
func (_OmniPortal *OmniPortalTransactorSession) SetXReceiptMaxErrorSize(numBytes uint16) (*types.Transaction, error) {
	return _OmniPortal.Contract.SetXReceiptMaxErrorSize(&_OmniPortal.TransactOpts, numBytes)
}

// SetXSubValsetCutoff is a paid mutator transaction binding the contract method 0x103ba701.
//
// Solidity: function setXSubValsetCutoff(uint8 xsubValsetCutoff_) returns()
func (_OmniPortal *OmniPortalTransactor) SetXSubValsetCutoff(opts *bind.TransactOpts, xsubValsetCutoff_ uint8) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "setXSubValsetCutoff", xsubValsetCutoff_)
}

// SetXSubValsetCutoff is a paid mutator transaction binding the contract method 0x103ba701.
//
// Solidity: function setXSubValsetCutoff(uint8 xsubValsetCutoff_) returns()
func (_OmniPortal *OmniPortalSession) SetXSubValsetCutoff(xsubValsetCutoff_ uint8) (*types.Transaction, error) {
	return _OmniPortal.Contract.SetXSubValsetCutoff(&_OmniPortal.TransactOpts, xsubValsetCutoff_)
}

// SetXSubValsetCutoff is a paid mutator transaction binding the contract method 0x103ba701.
//
// Solidity: function setXSubValsetCutoff(uint8 xsubValsetCutoff_) returns()
func (_OmniPortal *OmniPortalTransactorSession) SetXSubValsetCutoff(xsubValsetCutoff_ uint8) (*types.Transaction, error) {
	return _OmniPortal.Contract.SetXSubValsetCutoff(&_OmniPortal.TransactOpts, xsubValsetCutoff_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OmniPortal *OmniPortalTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OmniPortal *OmniPortalSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OmniPortal.Contract.TransferOwnership(&_OmniPortal.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OmniPortal *OmniPortalTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OmniPortal.Contract.TransferOwnership(&_OmniPortal.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_OmniPortal *OmniPortalTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_OmniPortal *OmniPortalSession) Unpause() (*types.Transaction, error) {
	return _OmniPortal.Contract.Unpause(&_OmniPortal.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_OmniPortal *OmniPortalTransactorSession) Unpause() (*types.Transaction, error) {
	return _OmniPortal.Contract.Unpause(&_OmniPortal.TransactOpts)
}

// UnpauseXCall is a paid mutator transaction binding the contract method 0x54d26bba.
//
// Solidity: function unpauseXCall() returns()
func (_OmniPortal *OmniPortalTransactor) UnpauseXCall(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "unpauseXCall")
}

// UnpauseXCall is a paid mutator transaction binding the contract method 0x54d26bba.
//
// Solidity: function unpauseXCall() returns()
func (_OmniPortal *OmniPortalSession) UnpauseXCall() (*types.Transaction, error) {
	return _OmniPortal.Contract.UnpauseXCall(&_OmniPortal.TransactOpts)
}

// UnpauseXCall is a paid mutator transaction binding the contract method 0x54d26bba.
//
// Solidity: function unpauseXCall() returns()
func (_OmniPortal *OmniPortalTransactorSession) UnpauseXCall() (*types.Transaction, error) {
	return _OmniPortal.Contract.UnpauseXCall(&_OmniPortal.TransactOpts)
}

// UnpauseXCallTo is a paid mutator transaction binding the contract method 0xd533b445.
//
// Solidity: function unpauseXCallTo(uint64 chainId_) returns()
func (_OmniPortal *OmniPortalTransactor) UnpauseXCallTo(opts *bind.TransactOpts, chainId_ uint64) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "unpauseXCallTo", chainId_)
}

// UnpauseXCallTo is a paid mutator transaction binding the contract method 0xd533b445.
//
// Solidity: function unpauseXCallTo(uint64 chainId_) returns()
func (_OmniPortal *OmniPortalSession) UnpauseXCallTo(chainId_ uint64) (*types.Transaction, error) {
	return _OmniPortal.Contract.UnpauseXCallTo(&_OmniPortal.TransactOpts, chainId_)
}

// UnpauseXCallTo is a paid mutator transaction binding the contract method 0xd533b445.
//
// Solidity: function unpauseXCallTo(uint64 chainId_) returns()
func (_OmniPortal *OmniPortalTransactorSession) UnpauseXCallTo(chainId_ uint64) (*types.Transaction, error) {
	return _OmniPortal.Contract.UnpauseXCallTo(&_OmniPortal.TransactOpts, chainId_)
}

// UnpauseXSubmit is a paid mutator transaction binding the contract method 0xc3d8ad67.
//
// Solidity: function unpauseXSubmit() returns()
func (_OmniPortal *OmniPortalTransactor) UnpauseXSubmit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "unpauseXSubmit")
}

// UnpauseXSubmit is a paid mutator transaction binding the contract method 0xc3d8ad67.
//
// Solidity: function unpauseXSubmit() returns()
func (_OmniPortal *OmniPortalSession) UnpauseXSubmit() (*types.Transaction, error) {
	return _OmniPortal.Contract.UnpauseXSubmit(&_OmniPortal.TransactOpts)
}

// UnpauseXSubmit is a paid mutator transaction binding the contract method 0xc3d8ad67.
//
// Solidity: function unpauseXSubmit() returns()
func (_OmniPortal *OmniPortalTransactorSession) UnpauseXSubmit() (*types.Transaction, error) {
	return _OmniPortal.Contract.UnpauseXSubmit(&_OmniPortal.TransactOpts)
}

// UnpauseXSubmitFrom is a paid mutator transaction binding the contract method 0xc2f9b968.
//
// Solidity: function unpauseXSubmitFrom(uint64 chainId_) returns()
func (_OmniPortal *OmniPortalTransactor) UnpauseXSubmitFrom(opts *bind.TransactOpts, chainId_ uint64) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "unpauseXSubmitFrom", chainId_)
}

// UnpauseXSubmitFrom is a paid mutator transaction binding the contract method 0xc2f9b968.
//
// Solidity: function unpauseXSubmitFrom(uint64 chainId_) returns()
func (_OmniPortal *OmniPortalSession) UnpauseXSubmitFrom(chainId_ uint64) (*types.Transaction, error) {
	return _OmniPortal.Contract.UnpauseXSubmitFrom(&_OmniPortal.TransactOpts, chainId_)
}

// UnpauseXSubmitFrom is a paid mutator transaction binding the contract method 0xc2f9b968.
//
// Solidity: function unpauseXSubmitFrom(uint64 chainId_) returns()
func (_OmniPortal *OmniPortalTransactorSession) UnpauseXSubmitFrom(chainId_ uint64) (*types.Transaction, error) {
	return _OmniPortal.Contract.UnpauseXSubmitFrom(&_OmniPortal.TransactOpts, chainId_)
}

// Xcall is a paid mutator transaction binding the contract method 0xc21dda4f.
//
// Solidity: function xcall(uint64 destChainId, uint8 conf, address to, bytes data, uint64 gasLimit) payable returns()
func (_OmniPortal *OmniPortalTransactor) Xcall(opts *bind.TransactOpts, destChainId uint64, conf uint8, to common.Address, data []byte, gasLimit uint64) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "xcall", destChainId, conf, to, data, gasLimit)
}

// Xcall is a paid mutator transaction binding the contract method 0xc21dda4f.
//
// Solidity: function xcall(uint64 destChainId, uint8 conf, address to, bytes data, uint64 gasLimit) payable returns()
func (_OmniPortal *OmniPortalSession) Xcall(destChainId uint64, conf uint8, to common.Address, data []byte, gasLimit uint64) (*types.Transaction, error) {
	return _OmniPortal.Contract.Xcall(&_OmniPortal.TransactOpts, destChainId, conf, to, data, gasLimit)
}

// Xcall is a paid mutator transaction binding the contract method 0xc21dda4f.
//
// Solidity: function xcall(uint64 destChainId, uint8 conf, address to, bytes data, uint64 gasLimit) payable returns()
func (_OmniPortal *OmniPortalTransactorSession) Xcall(destChainId uint64, conf uint8, to common.Address, data []byte, gasLimit uint64) (*types.Transaction, error) {
	return _OmniPortal.Contract.Xcall(&_OmniPortal.TransactOpts, destChainId, conf, to, data, gasLimit)
}

// Xsubmit is a paid mutator transaction binding the contract method 0x66a1eaf3.
//
// Solidity: function xsubmit((bytes32,uint64,(uint64,uint64,uint8,uint64,uint64,bytes32),(uint64,uint64,uint64,address,address,bytes,uint64)[],bytes32[],bool[],(address,bytes)[]) xsub) returns()
func (_OmniPortal *OmniPortalTransactor) Xsubmit(opts *bind.TransactOpts, xsub XTypesSubmission) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "xsubmit", xsub)
}

// Xsubmit is a paid mutator transaction binding the contract method 0x66a1eaf3.
//
// Solidity: function xsubmit((bytes32,uint64,(uint64,uint64,uint8,uint64,uint64,bytes32),(uint64,uint64,uint64,address,address,bytes,uint64)[],bytes32[],bool[],(address,bytes)[]) xsub) returns()
func (_OmniPortal *OmniPortalSession) Xsubmit(xsub XTypesSubmission) (*types.Transaction, error) {
	return _OmniPortal.Contract.Xsubmit(&_OmniPortal.TransactOpts, xsub)
}

// Xsubmit is a paid mutator transaction binding the contract method 0x66a1eaf3.
//
// Solidity: function xsubmit((bytes32,uint64,(uint64,uint64,uint8,uint64,uint64,bytes32),(uint64,uint64,uint64,address,address,bytes,uint64)[],bytes32[],bool[],(address,bytes)[]) xsub) returns()
func (_OmniPortal *OmniPortalTransactorSession) Xsubmit(xsub XTypesSubmission) (*types.Transaction, error) {
	return _OmniPortal.Contract.Xsubmit(&_OmniPortal.TransactOpts, xsub)
}

// OmniPortalFeeOracleSetIterator is returned from FilterFeeOracleSet and is used to iterate over the raw logs and unpacked data for FeeOracleSet events raised by the OmniPortal contract.
type OmniPortalFeeOracleSetIterator struct {
	Event *OmniPortalFeeOracleSet // Event containing the contract specifics and raw log

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
func (it *OmniPortalFeeOracleSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniPortalFeeOracleSet)
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
		it.Event = new(OmniPortalFeeOracleSet)
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
func (it *OmniPortalFeeOracleSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniPortalFeeOracleSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniPortalFeeOracleSet represents a FeeOracleSet event raised by the OmniPortal contract.
type OmniPortalFeeOracleSet struct {
	Oracle common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFeeOracleSet is a free log retrieval operation binding the contract event 0xd97bdb0db82b52a85aa07f8da78033b1d6e159d94f1e3cbd4109d946c3bcfd32.
//
// Solidity: event FeeOracleSet(address oracle)
func (_OmniPortal *OmniPortalFilterer) FilterFeeOracleSet(opts *bind.FilterOpts) (*OmniPortalFeeOracleSetIterator, error) {

	logs, sub, err := _OmniPortal.contract.FilterLogs(opts, "FeeOracleSet")
	if err != nil {
		return nil, err
	}
	return &OmniPortalFeeOracleSetIterator{contract: _OmniPortal.contract, event: "FeeOracleSet", logs: logs, sub: sub}, nil
}

// WatchFeeOracleSet is a free log subscription operation binding the contract event 0xd97bdb0db82b52a85aa07f8da78033b1d6e159d94f1e3cbd4109d946c3bcfd32.
//
// Solidity: event FeeOracleSet(address oracle)
func (_OmniPortal *OmniPortalFilterer) WatchFeeOracleSet(opts *bind.WatchOpts, sink chan<- *OmniPortalFeeOracleSet) (event.Subscription, error) {

	logs, sub, err := _OmniPortal.contract.WatchLogs(opts, "FeeOracleSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniPortalFeeOracleSet)
				if err := _OmniPortal.contract.UnpackLog(event, "FeeOracleSet", log); err != nil {
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

// ParseFeeOracleSet is a log parse operation binding the contract event 0xd97bdb0db82b52a85aa07f8da78033b1d6e159d94f1e3cbd4109d946c3bcfd32.
//
// Solidity: event FeeOracleSet(address oracle)
func (_OmniPortal *OmniPortalFilterer) ParseFeeOracleSet(log types.Log) (*OmniPortalFeeOracleSet, error) {
	event := new(OmniPortalFeeOracleSet)
	if err := _OmniPortal.contract.UnpackLog(event, "FeeOracleSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniPortalFeesCollectedIterator is returned from FilterFeesCollected and is used to iterate over the raw logs and unpacked data for FeesCollected events raised by the OmniPortal contract.
type OmniPortalFeesCollectedIterator struct {
	Event *OmniPortalFeesCollected // Event containing the contract specifics and raw log

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
func (it *OmniPortalFeesCollectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniPortalFeesCollected)
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
		it.Event = new(OmniPortalFeesCollected)
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
func (it *OmniPortalFeesCollectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniPortalFeesCollectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniPortalFeesCollected represents a FeesCollected event raised by the OmniPortal contract.
type OmniPortalFeesCollected struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFeesCollected is a free log retrieval operation binding the contract event 0x9dc46f23cfb5ddcad0ae7ea2be38d47fec07bb9382ec7e564efc69e036dd66ce.
//
// Solidity: event FeesCollected(address indexed to, uint256 amount)
func (_OmniPortal *OmniPortalFilterer) FilterFeesCollected(opts *bind.FilterOpts, to []common.Address) (*OmniPortalFeesCollectedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OmniPortal.contract.FilterLogs(opts, "FeesCollected", toRule)
	if err != nil {
		return nil, err
	}
	return &OmniPortalFeesCollectedIterator{contract: _OmniPortal.contract, event: "FeesCollected", logs: logs, sub: sub}, nil
}

// WatchFeesCollected is a free log subscription operation binding the contract event 0x9dc46f23cfb5ddcad0ae7ea2be38d47fec07bb9382ec7e564efc69e036dd66ce.
//
// Solidity: event FeesCollected(address indexed to, uint256 amount)
func (_OmniPortal *OmniPortalFilterer) WatchFeesCollected(opts *bind.WatchOpts, sink chan<- *OmniPortalFeesCollected, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OmniPortal.contract.WatchLogs(opts, "FeesCollected", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniPortalFeesCollected)
				if err := _OmniPortal.contract.UnpackLog(event, "FeesCollected", log); err != nil {
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

// ParseFeesCollected is a log parse operation binding the contract event 0x9dc46f23cfb5ddcad0ae7ea2be38d47fec07bb9382ec7e564efc69e036dd66ce.
//
// Solidity: event FeesCollected(address indexed to, uint256 amount)
func (_OmniPortal *OmniPortalFilterer) ParseFeesCollected(log types.Log) (*OmniPortalFeesCollected, error) {
	event := new(OmniPortalFeesCollected)
	if err := _OmniPortal.contract.UnpackLog(event, "FeesCollected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniPortalInXBlockOffsetSetIterator is returned from FilterInXBlockOffsetSet and is used to iterate over the raw logs and unpacked data for InXBlockOffsetSet events raised by the OmniPortal contract.
type OmniPortalInXBlockOffsetSetIterator struct {
	Event *OmniPortalInXBlockOffsetSet // Event containing the contract specifics and raw log

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
func (it *OmniPortalInXBlockOffsetSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniPortalInXBlockOffsetSet)
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
		it.Event = new(OmniPortalInXBlockOffsetSet)
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
func (it *OmniPortalInXBlockOffsetSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniPortalInXBlockOffsetSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniPortalInXBlockOffsetSet represents a InXBlockOffsetSet event raised by the OmniPortal contract.
type OmniPortalInXBlockOffsetSet struct {
	SrcChainId uint64
	ShardId    uint64
	Offset     uint64
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInXBlockOffsetSet is a free log retrieval operation binding the contract event 0xe070f08cae8464c91238e8cbea64ccee5e7b48dd79a843f144e3721ee6bdd9b5.
//
// Solidity: event InXBlockOffsetSet(uint64 indexed srcChainId, uint64 indexed shardId, uint64 offset)
func (_OmniPortal *OmniPortalFilterer) FilterInXBlockOffsetSet(opts *bind.FilterOpts, srcChainId []uint64, shardId []uint64) (*OmniPortalInXBlockOffsetSetIterator, error) {

	var srcChainIdRule []interface{}
	for _, srcChainIdItem := range srcChainId {
		srcChainIdRule = append(srcChainIdRule, srcChainIdItem)
	}
	var shardIdRule []interface{}
	for _, shardIdItem := range shardId {
		shardIdRule = append(shardIdRule, shardIdItem)
	}

	logs, sub, err := _OmniPortal.contract.FilterLogs(opts, "InXBlockOffsetSet", srcChainIdRule, shardIdRule)
	if err != nil {
		return nil, err
	}
	return &OmniPortalInXBlockOffsetSetIterator{contract: _OmniPortal.contract, event: "InXBlockOffsetSet", logs: logs, sub: sub}, nil
}

// WatchInXBlockOffsetSet is a free log subscription operation binding the contract event 0xe070f08cae8464c91238e8cbea64ccee5e7b48dd79a843f144e3721ee6bdd9b5.
//
// Solidity: event InXBlockOffsetSet(uint64 indexed srcChainId, uint64 indexed shardId, uint64 offset)
func (_OmniPortal *OmniPortalFilterer) WatchInXBlockOffsetSet(opts *bind.WatchOpts, sink chan<- *OmniPortalInXBlockOffsetSet, srcChainId []uint64, shardId []uint64) (event.Subscription, error) {

	var srcChainIdRule []interface{}
	for _, srcChainIdItem := range srcChainId {
		srcChainIdRule = append(srcChainIdRule, srcChainIdItem)
	}
	var shardIdRule []interface{}
	for _, shardIdItem := range shardId {
		shardIdRule = append(shardIdRule, shardIdItem)
	}

	logs, sub, err := _OmniPortal.contract.WatchLogs(opts, "InXBlockOffsetSet", srcChainIdRule, shardIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniPortalInXBlockOffsetSet)
				if err := _OmniPortal.contract.UnpackLog(event, "InXBlockOffsetSet", log); err != nil {
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

// ParseInXBlockOffsetSet is a log parse operation binding the contract event 0xe070f08cae8464c91238e8cbea64ccee5e7b48dd79a843f144e3721ee6bdd9b5.
//
// Solidity: event InXBlockOffsetSet(uint64 indexed srcChainId, uint64 indexed shardId, uint64 offset)
func (_OmniPortal *OmniPortalFilterer) ParseInXBlockOffsetSet(log types.Log) (*OmniPortalInXBlockOffsetSet, error) {
	event := new(OmniPortalInXBlockOffsetSet)
	if err := _OmniPortal.contract.UnpackLog(event, "InXBlockOffsetSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniPortalInXMsgOffsetSetIterator is returned from FilterInXMsgOffsetSet and is used to iterate over the raw logs and unpacked data for InXMsgOffsetSet events raised by the OmniPortal contract.
type OmniPortalInXMsgOffsetSetIterator struct {
	Event *OmniPortalInXMsgOffsetSet // Event containing the contract specifics and raw log

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
func (it *OmniPortalInXMsgOffsetSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniPortalInXMsgOffsetSet)
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
		it.Event = new(OmniPortalInXMsgOffsetSet)
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
func (it *OmniPortalInXMsgOffsetSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniPortalInXMsgOffsetSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniPortalInXMsgOffsetSet represents a InXMsgOffsetSet event raised by the OmniPortal contract.
type OmniPortalInXMsgOffsetSet struct {
	SrcChainId uint64
	ShardId    uint64
	Offset     uint64
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInXMsgOffsetSet is a free log retrieval operation binding the contract event 0x8647aae68c8456a1dcbfaf5eaadc94278ae423526d3f09c7b972bff7355d55c7.
//
// Solidity: event InXMsgOffsetSet(uint64 indexed srcChainId, uint64 indexed shardId, uint64 offset)
func (_OmniPortal *OmniPortalFilterer) FilterInXMsgOffsetSet(opts *bind.FilterOpts, srcChainId []uint64, shardId []uint64) (*OmniPortalInXMsgOffsetSetIterator, error) {

	var srcChainIdRule []interface{}
	for _, srcChainIdItem := range srcChainId {
		srcChainIdRule = append(srcChainIdRule, srcChainIdItem)
	}
	var shardIdRule []interface{}
	for _, shardIdItem := range shardId {
		shardIdRule = append(shardIdRule, shardIdItem)
	}

	logs, sub, err := _OmniPortal.contract.FilterLogs(opts, "InXMsgOffsetSet", srcChainIdRule, shardIdRule)
	if err != nil {
		return nil, err
	}
	return &OmniPortalInXMsgOffsetSetIterator{contract: _OmniPortal.contract, event: "InXMsgOffsetSet", logs: logs, sub: sub}, nil
}

// WatchInXMsgOffsetSet is a free log subscription operation binding the contract event 0x8647aae68c8456a1dcbfaf5eaadc94278ae423526d3f09c7b972bff7355d55c7.
//
// Solidity: event InXMsgOffsetSet(uint64 indexed srcChainId, uint64 indexed shardId, uint64 offset)
func (_OmniPortal *OmniPortalFilterer) WatchInXMsgOffsetSet(opts *bind.WatchOpts, sink chan<- *OmniPortalInXMsgOffsetSet, srcChainId []uint64, shardId []uint64) (event.Subscription, error) {

	var srcChainIdRule []interface{}
	for _, srcChainIdItem := range srcChainId {
		srcChainIdRule = append(srcChainIdRule, srcChainIdItem)
	}
	var shardIdRule []interface{}
	for _, shardIdItem := range shardId {
		shardIdRule = append(shardIdRule, shardIdItem)
	}

	logs, sub, err := _OmniPortal.contract.WatchLogs(opts, "InXMsgOffsetSet", srcChainIdRule, shardIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniPortalInXMsgOffsetSet)
				if err := _OmniPortal.contract.UnpackLog(event, "InXMsgOffsetSet", log); err != nil {
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

// ParseInXMsgOffsetSet is a log parse operation binding the contract event 0x8647aae68c8456a1dcbfaf5eaadc94278ae423526d3f09c7b972bff7355d55c7.
//
// Solidity: event InXMsgOffsetSet(uint64 indexed srcChainId, uint64 indexed shardId, uint64 offset)
func (_OmniPortal *OmniPortalFilterer) ParseInXMsgOffsetSet(log types.Log) (*OmniPortalInXMsgOffsetSet, error) {
	event := new(OmniPortalInXMsgOffsetSet)
	if err := _OmniPortal.contract.UnpackLog(event, "InXMsgOffsetSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniPortalInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the OmniPortal contract.
type OmniPortalInitializedIterator struct {
	Event *OmniPortalInitialized // Event containing the contract specifics and raw log

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
func (it *OmniPortalInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniPortalInitialized)
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
		it.Event = new(OmniPortalInitialized)
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
func (it *OmniPortalInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniPortalInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniPortalInitialized represents a Initialized event raised by the OmniPortal contract.
type OmniPortalInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_OmniPortal *OmniPortalFilterer) FilterInitialized(opts *bind.FilterOpts) (*OmniPortalInitializedIterator, error) {

	logs, sub, err := _OmniPortal.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &OmniPortalInitializedIterator{contract: _OmniPortal.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_OmniPortal *OmniPortalFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *OmniPortalInitialized) (event.Subscription, error) {

	logs, sub, err := _OmniPortal.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniPortalInitialized)
				if err := _OmniPortal.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_OmniPortal *OmniPortalFilterer) ParseInitialized(log types.Log) (*OmniPortalInitialized, error) {
	event := new(OmniPortalInitialized)
	if err := _OmniPortal.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniPortalOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the OmniPortal contract.
type OmniPortalOwnershipTransferredIterator struct {
	Event *OmniPortalOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OmniPortalOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniPortalOwnershipTransferred)
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
		it.Event = new(OmniPortalOwnershipTransferred)
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
func (it *OmniPortalOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniPortalOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniPortalOwnershipTransferred represents a OwnershipTransferred event raised by the OmniPortal contract.
type OmniPortalOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OmniPortal *OmniPortalFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OmniPortalOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OmniPortal.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OmniPortalOwnershipTransferredIterator{contract: _OmniPortal.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OmniPortal *OmniPortalFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OmniPortalOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OmniPortal.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniPortalOwnershipTransferred)
				if err := _OmniPortal.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_OmniPortal *OmniPortalFilterer) ParseOwnershipTransferred(log types.Log) (*OmniPortalOwnershipTransferred, error) {
	event := new(OmniPortalOwnershipTransferred)
	if err := _OmniPortal.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniPortalPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the OmniPortal contract.
type OmniPortalPausedIterator struct {
	Event *OmniPortalPaused // Event containing the contract specifics and raw log

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
func (it *OmniPortalPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniPortalPaused)
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
		it.Event = new(OmniPortalPaused)
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
func (it *OmniPortalPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniPortalPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniPortalPaused represents a Paused event raised by the OmniPortal contract.
type OmniPortalPaused struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x9e87fac88ff661f02d44f95383c817fece4bce600a3dab7a54406878b965e752.
//
// Solidity: event Paused()
func (_OmniPortal *OmniPortalFilterer) FilterPaused(opts *bind.FilterOpts) (*OmniPortalPausedIterator, error) {

	logs, sub, err := _OmniPortal.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &OmniPortalPausedIterator{contract: _OmniPortal.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x9e87fac88ff661f02d44f95383c817fece4bce600a3dab7a54406878b965e752.
//
// Solidity: event Paused()
func (_OmniPortal *OmniPortalFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *OmniPortalPaused) (event.Subscription, error) {

	logs, sub, err := _OmniPortal.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniPortalPaused)
				if err := _OmniPortal.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x9e87fac88ff661f02d44f95383c817fece4bce600a3dab7a54406878b965e752.
//
// Solidity: event Paused()
func (_OmniPortal *OmniPortalFilterer) ParsePaused(log types.Log) (*OmniPortalPaused, error) {
	event := new(OmniPortalPaused)
	if err := _OmniPortal.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniPortalUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the OmniPortal contract.
type OmniPortalUnpausedIterator struct {
	Event *OmniPortalUnpaused // Event containing the contract specifics and raw log

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
func (it *OmniPortalUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniPortalUnpaused)
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
		it.Event = new(OmniPortalUnpaused)
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
func (it *OmniPortalUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniPortalUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniPortalUnpaused represents a Unpaused event raised by the OmniPortal contract.
type OmniPortalUnpaused struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0xa45f47fdea8a1efdd9029a5691c7f759c32b7c698632b563573e155625d16933.
//
// Solidity: event Unpaused()
func (_OmniPortal *OmniPortalFilterer) FilterUnpaused(opts *bind.FilterOpts) (*OmniPortalUnpausedIterator, error) {

	logs, sub, err := _OmniPortal.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &OmniPortalUnpausedIterator{contract: _OmniPortal.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0xa45f47fdea8a1efdd9029a5691c7f759c32b7c698632b563573e155625d16933.
//
// Solidity: event Unpaused()
func (_OmniPortal *OmniPortalFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *OmniPortalUnpaused) (event.Subscription, error) {

	logs, sub, err := _OmniPortal.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniPortalUnpaused)
				if err := _OmniPortal.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0xa45f47fdea8a1efdd9029a5691c7f759c32b7c698632b563573e155625d16933.
//
// Solidity: event Unpaused()
func (_OmniPortal *OmniPortalFilterer) ParseUnpaused(log types.Log) (*OmniPortalUnpaused, error) {
	event := new(OmniPortalUnpaused)
	if err := _OmniPortal.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniPortalValidatorSetAddedIterator is returned from FilterValidatorSetAdded and is used to iterate over the raw logs and unpacked data for ValidatorSetAdded events raised by the OmniPortal contract.
type OmniPortalValidatorSetAddedIterator struct {
	Event *OmniPortalValidatorSetAdded // Event containing the contract specifics and raw log

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
func (it *OmniPortalValidatorSetAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniPortalValidatorSetAdded)
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
		it.Event = new(OmniPortalValidatorSetAdded)
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
func (it *OmniPortalValidatorSetAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniPortalValidatorSetAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniPortalValidatorSetAdded represents a ValidatorSetAdded event raised by the OmniPortal contract.
type OmniPortalValidatorSetAdded struct {
	SetId uint64
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterValidatorSetAdded is a free log retrieval operation binding the contract event 0x3a7c2f997a87ba92aedaecd1127f4129cae1283e2809ebf5304d321b943fd107.
//
// Solidity: event ValidatorSetAdded(uint64 indexed setId)
func (_OmniPortal *OmniPortalFilterer) FilterValidatorSetAdded(opts *bind.FilterOpts, setId []uint64) (*OmniPortalValidatorSetAddedIterator, error) {

	var setIdRule []interface{}
	for _, setIdItem := range setId {
		setIdRule = append(setIdRule, setIdItem)
	}

	logs, sub, err := _OmniPortal.contract.FilterLogs(opts, "ValidatorSetAdded", setIdRule)
	if err != nil {
		return nil, err
	}
	return &OmniPortalValidatorSetAddedIterator{contract: _OmniPortal.contract, event: "ValidatorSetAdded", logs: logs, sub: sub}, nil
}

// WatchValidatorSetAdded is a free log subscription operation binding the contract event 0x3a7c2f997a87ba92aedaecd1127f4129cae1283e2809ebf5304d321b943fd107.
//
// Solidity: event ValidatorSetAdded(uint64 indexed setId)
func (_OmniPortal *OmniPortalFilterer) WatchValidatorSetAdded(opts *bind.WatchOpts, sink chan<- *OmniPortalValidatorSetAdded, setId []uint64) (event.Subscription, error) {

	var setIdRule []interface{}
	for _, setIdItem := range setId {
		setIdRule = append(setIdRule, setIdItem)
	}

	logs, sub, err := _OmniPortal.contract.WatchLogs(opts, "ValidatorSetAdded", setIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniPortalValidatorSetAdded)
				if err := _OmniPortal.contract.UnpackLog(event, "ValidatorSetAdded", log); err != nil {
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

// ParseValidatorSetAdded is a log parse operation binding the contract event 0x3a7c2f997a87ba92aedaecd1127f4129cae1283e2809ebf5304d321b943fd107.
//
// Solidity: event ValidatorSetAdded(uint64 indexed setId)
func (_OmniPortal *OmniPortalFilterer) ParseValidatorSetAdded(log types.Log) (*OmniPortalValidatorSetAdded, error) {
	event := new(OmniPortalValidatorSetAdded)
	if err := _OmniPortal.contract.UnpackLog(event, "ValidatorSetAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniPortalXCallPausedIterator is returned from FilterXCallPaused and is used to iterate over the raw logs and unpacked data for XCallPaused events raised by the OmniPortal contract.
type OmniPortalXCallPausedIterator struct {
	Event *OmniPortalXCallPaused // Event containing the contract specifics and raw log

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
func (it *OmniPortalXCallPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniPortalXCallPaused)
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
		it.Event = new(OmniPortalXCallPaused)
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
func (it *OmniPortalXCallPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniPortalXCallPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniPortalXCallPaused represents a XCallPaused event raised by the OmniPortal contract.
type OmniPortalXCallPaused struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterXCallPaused is a free log retrieval operation binding the contract event 0x5f335a4032d4cfb6aca7835b0c2225f36d4d9eaa4ed43ee59ed537e02dff6b39.
//
// Solidity: event XCallPaused()
func (_OmniPortal *OmniPortalFilterer) FilterXCallPaused(opts *bind.FilterOpts) (*OmniPortalXCallPausedIterator, error) {

	logs, sub, err := _OmniPortal.contract.FilterLogs(opts, "XCallPaused")
	if err != nil {
		return nil, err
	}
	return &OmniPortalXCallPausedIterator{contract: _OmniPortal.contract, event: "XCallPaused", logs: logs, sub: sub}, nil
}

// WatchXCallPaused is a free log subscription operation binding the contract event 0x5f335a4032d4cfb6aca7835b0c2225f36d4d9eaa4ed43ee59ed537e02dff6b39.
//
// Solidity: event XCallPaused()
func (_OmniPortal *OmniPortalFilterer) WatchXCallPaused(opts *bind.WatchOpts, sink chan<- *OmniPortalXCallPaused) (event.Subscription, error) {

	logs, sub, err := _OmniPortal.contract.WatchLogs(opts, "XCallPaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniPortalXCallPaused)
				if err := _OmniPortal.contract.UnpackLog(event, "XCallPaused", log); err != nil {
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

// ParseXCallPaused is a log parse operation binding the contract event 0x5f335a4032d4cfb6aca7835b0c2225f36d4d9eaa4ed43ee59ed537e02dff6b39.
//
// Solidity: event XCallPaused()
func (_OmniPortal *OmniPortalFilterer) ParseXCallPaused(log types.Log) (*OmniPortalXCallPaused, error) {
	event := new(OmniPortalXCallPaused)
	if err := _OmniPortal.contract.UnpackLog(event, "XCallPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniPortalXCallToPausedIterator is returned from FilterXCallToPaused and is used to iterate over the raw logs and unpacked data for XCallToPaused events raised by the OmniPortal contract.
type OmniPortalXCallToPausedIterator struct {
	Event *OmniPortalXCallToPaused // Event containing the contract specifics and raw log

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
func (it *OmniPortalXCallToPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniPortalXCallToPaused)
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
		it.Event = new(OmniPortalXCallToPaused)
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
func (it *OmniPortalXCallToPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniPortalXCallToPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniPortalXCallToPaused represents a XCallToPaused event raised by the OmniPortal contract.
type OmniPortalXCallToPaused struct {
	ChainId uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterXCallToPaused is a free log retrieval operation binding the contract event 0xcd7910e1c5569d8433ce4ef8e5d51c1bdc03168f614b576da47dc3d2b51d033a.
//
// Solidity: event XCallToPaused(uint64 indexed chainId)
func (_OmniPortal *OmniPortalFilterer) FilterXCallToPaused(opts *bind.FilterOpts, chainId []uint64) (*OmniPortalXCallToPausedIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _OmniPortal.contract.FilterLogs(opts, "XCallToPaused", chainIdRule)
	if err != nil {
		return nil, err
	}
	return &OmniPortalXCallToPausedIterator{contract: _OmniPortal.contract, event: "XCallToPaused", logs: logs, sub: sub}, nil
}

// WatchXCallToPaused is a free log subscription operation binding the contract event 0xcd7910e1c5569d8433ce4ef8e5d51c1bdc03168f614b576da47dc3d2b51d033a.
//
// Solidity: event XCallToPaused(uint64 indexed chainId)
func (_OmniPortal *OmniPortalFilterer) WatchXCallToPaused(opts *bind.WatchOpts, sink chan<- *OmniPortalXCallToPaused, chainId []uint64) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _OmniPortal.contract.WatchLogs(opts, "XCallToPaused", chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniPortalXCallToPaused)
				if err := _OmniPortal.contract.UnpackLog(event, "XCallToPaused", log); err != nil {
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

// ParseXCallToPaused is a log parse operation binding the contract event 0xcd7910e1c5569d8433ce4ef8e5d51c1bdc03168f614b576da47dc3d2b51d033a.
//
// Solidity: event XCallToPaused(uint64 indexed chainId)
func (_OmniPortal *OmniPortalFilterer) ParseXCallToPaused(log types.Log) (*OmniPortalXCallToPaused, error) {
	event := new(OmniPortalXCallToPaused)
	if err := _OmniPortal.contract.UnpackLog(event, "XCallToPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniPortalXCallToUnpausedIterator is returned from FilterXCallToUnpaused and is used to iterate over the raw logs and unpacked data for XCallToUnpaused events raised by the OmniPortal contract.
type OmniPortalXCallToUnpausedIterator struct {
	Event *OmniPortalXCallToUnpaused // Event containing the contract specifics and raw log

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
func (it *OmniPortalXCallToUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniPortalXCallToUnpaused)
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
		it.Event = new(OmniPortalXCallToUnpaused)
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
func (it *OmniPortalXCallToUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniPortalXCallToUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniPortalXCallToUnpaused represents a XCallToUnpaused event raised by the OmniPortal contract.
type OmniPortalXCallToUnpaused struct {
	ChainId uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterXCallToUnpaused is a free log retrieval operation binding the contract event 0x1ed9223556fb0971076c30172f1f00630efd313b6a05290a562aef95928e7125.
//
// Solidity: event XCallToUnpaused(uint64 indexed chainId)
func (_OmniPortal *OmniPortalFilterer) FilterXCallToUnpaused(opts *bind.FilterOpts, chainId []uint64) (*OmniPortalXCallToUnpausedIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _OmniPortal.contract.FilterLogs(opts, "XCallToUnpaused", chainIdRule)
	if err != nil {
		return nil, err
	}
	return &OmniPortalXCallToUnpausedIterator{contract: _OmniPortal.contract, event: "XCallToUnpaused", logs: logs, sub: sub}, nil
}

// WatchXCallToUnpaused is a free log subscription operation binding the contract event 0x1ed9223556fb0971076c30172f1f00630efd313b6a05290a562aef95928e7125.
//
// Solidity: event XCallToUnpaused(uint64 indexed chainId)
func (_OmniPortal *OmniPortalFilterer) WatchXCallToUnpaused(opts *bind.WatchOpts, sink chan<- *OmniPortalXCallToUnpaused, chainId []uint64) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _OmniPortal.contract.WatchLogs(opts, "XCallToUnpaused", chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniPortalXCallToUnpaused)
				if err := _OmniPortal.contract.UnpackLog(event, "XCallToUnpaused", log); err != nil {
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

// ParseXCallToUnpaused is a log parse operation binding the contract event 0x1ed9223556fb0971076c30172f1f00630efd313b6a05290a562aef95928e7125.
//
// Solidity: event XCallToUnpaused(uint64 indexed chainId)
func (_OmniPortal *OmniPortalFilterer) ParseXCallToUnpaused(log types.Log) (*OmniPortalXCallToUnpaused, error) {
	event := new(OmniPortalXCallToUnpaused)
	if err := _OmniPortal.contract.UnpackLog(event, "XCallToUnpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniPortalXCallUnpausedIterator is returned from FilterXCallUnpaused and is used to iterate over the raw logs and unpacked data for XCallUnpaused events raised by the OmniPortal contract.
type OmniPortalXCallUnpausedIterator struct {
	Event *OmniPortalXCallUnpaused // Event containing the contract specifics and raw log

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
func (it *OmniPortalXCallUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniPortalXCallUnpaused)
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
		it.Event = new(OmniPortalXCallUnpaused)
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
func (it *OmniPortalXCallUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniPortalXCallUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniPortalXCallUnpaused represents a XCallUnpaused event raised by the OmniPortal contract.
type OmniPortalXCallUnpaused struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterXCallUnpaused is a free log retrieval operation binding the contract event 0x4c48c7b71557216a3192842746bdfc381f98d7536d9eb1c6764f3b45e6794827.
//
// Solidity: event XCallUnpaused()
func (_OmniPortal *OmniPortalFilterer) FilterXCallUnpaused(opts *bind.FilterOpts) (*OmniPortalXCallUnpausedIterator, error) {

	logs, sub, err := _OmniPortal.contract.FilterLogs(opts, "XCallUnpaused")
	if err != nil {
		return nil, err
	}
	return &OmniPortalXCallUnpausedIterator{contract: _OmniPortal.contract, event: "XCallUnpaused", logs: logs, sub: sub}, nil
}

// WatchXCallUnpaused is a free log subscription operation binding the contract event 0x4c48c7b71557216a3192842746bdfc381f98d7536d9eb1c6764f3b45e6794827.
//
// Solidity: event XCallUnpaused()
func (_OmniPortal *OmniPortalFilterer) WatchXCallUnpaused(opts *bind.WatchOpts, sink chan<- *OmniPortalXCallUnpaused) (event.Subscription, error) {

	logs, sub, err := _OmniPortal.contract.WatchLogs(opts, "XCallUnpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniPortalXCallUnpaused)
				if err := _OmniPortal.contract.UnpackLog(event, "XCallUnpaused", log); err != nil {
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

// ParseXCallUnpaused is a log parse operation binding the contract event 0x4c48c7b71557216a3192842746bdfc381f98d7536d9eb1c6764f3b45e6794827.
//
// Solidity: event XCallUnpaused()
func (_OmniPortal *OmniPortalFilterer) ParseXCallUnpaused(log types.Log) (*OmniPortalXCallUnpaused, error) {
	event := new(OmniPortalXCallUnpaused)
	if err := _OmniPortal.contract.UnpackLog(event, "XCallUnpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniPortalXMsgIterator is returned from FilterXMsg and is used to iterate over the raw logs and unpacked data for XMsg events raised by the OmniPortal contract.
type OmniPortalXMsgIterator struct {
	Event *OmniPortalXMsg // Event containing the contract specifics and raw log

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
func (it *OmniPortalXMsgIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniPortalXMsg)
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
		it.Event = new(OmniPortalXMsg)
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
func (it *OmniPortalXMsgIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniPortalXMsgIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniPortalXMsg represents a XMsg event raised by the OmniPortal contract.
type OmniPortalXMsg struct {
	DestChainId uint64
	ShardId     uint64
	Offset      uint64
	Sender      common.Address
	To          common.Address
	Data        []byte
	GasLimit    uint64
	Fees        *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterXMsg is a free log retrieval operation binding the contract event 0xb7c8eb9d7a7fbcdab809ab7b8a7c41701eb3115e3fe99d30ff490d8552f72bfa.
//
// Solidity: event XMsg(uint64 indexed destChainId, uint64 indexed shardId, uint64 indexed offset, address sender, address to, bytes data, uint64 gasLimit, uint256 fees)
func (_OmniPortal *OmniPortalFilterer) FilterXMsg(opts *bind.FilterOpts, destChainId []uint64, shardId []uint64, offset []uint64) (*OmniPortalXMsgIterator, error) {

	var destChainIdRule []interface{}
	for _, destChainIdItem := range destChainId {
		destChainIdRule = append(destChainIdRule, destChainIdItem)
	}
	var shardIdRule []interface{}
	for _, shardIdItem := range shardId {
		shardIdRule = append(shardIdRule, shardIdItem)
	}
	var offsetRule []interface{}
	for _, offsetItem := range offset {
		offsetRule = append(offsetRule, offsetItem)
	}

	logs, sub, err := _OmniPortal.contract.FilterLogs(opts, "XMsg", destChainIdRule, shardIdRule, offsetRule)
	if err != nil {
		return nil, err
	}
	return &OmniPortalXMsgIterator{contract: _OmniPortal.contract, event: "XMsg", logs: logs, sub: sub}, nil
}

// WatchXMsg is a free log subscription operation binding the contract event 0xb7c8eb9d7a7fbcdab809ab7b8a7c41701eb3115e3fe99d30ff490d8552f72bfa.
//
// Solidity: event XMsg(uint64 indexed destChainId, uint64 indexed shardId, uint64 indexed offset, address sender, address to, bytes data, uint64 gasLimit, uint256 fees)
func (_OmniPortal *OmniPortalFilterer) WatchXMsg(opts *bind.WatchOpts, sink chan<- *OmniPortalXMsg, destChainId []uint64, shardId []uint64, offset []uint64) (event.Subscription, error) {

	var destChainIdRule []interface{}
	for _, destChainIdItem := range destChainId {
		destChainIdRule = append(destChainIdRule, destChainIdItem)
	}
	var shardIdRule []interface{}
	for _, shardIdItem := range shardId {
		shardIdRule = append(shardIdRule, shardIdItem)
	}
	var offsetRule []interface{}
	for _, offsetItem := range offset {
		offsetRule = append(offsetRule, offsetItem)
	}

	logs, sub, err := _OmniPortal.contract.WatchLogs(opts, "XMsg", destChainIdRule, shardIdRule, offsetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniPortalXMsg)
				if err := _OmniPortal.contract.UnpackLog(event, "XMsg", log); err != nil {
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

// ParseXMsg is a log parse operation binding the contract event 0xb7c8eb9d7a7fbcdab809ab7b8a7c41701eb3115e3fe99d30ff490d8552f72bfa.
//
// Solidity: event XMsg(uint64 indexed destChainId, uint64 indexed shardId, uint64 indexed offset, address sender, address to, bytes data, uint64 gasLimit, uint256 fees)
func (_OmniPortal *OmniPortalFilterer) ParseXMsg(log types.Log) (*OmniPortalXMsg, error) {
	event := new(OmniPortalXMsg)
	if err := _OmniPortal.contract.UnpackLog(event, "XMsg", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniPortalXMsgMaxDataSizeSetIterator is returned from FilterXMsgMaxDataSizeSet and is used to iterate over the raw logs and unpacked data for XMsgMaxDataSizeSet events raised by the OmniPortal contract.
type OmniPortalXMsgMaxDataSizeSetIterator struct {
	Event *OmniPortalXMsgMaxDataSizeSet // Event containing the contract specifics and raw log

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
func (it *OmniPortalXMsgMaxDataSizeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniPortalXMsgMaxDataSizeSet)
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
		it.Event = new(OmniPortalXMsgMaxDataSizeSet)
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
func (it *OmniPortalXMsgMaxDataSizeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniPortalXMsgMaxDataSizeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniPortalXMsgMaxDataSizeSet represents a XMsgMaxDataSizeSet event raised by the OmniPortal contract.
type OmniPortalXMsgMaxDataSizeSet struct {
	Size uint16
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterXMsgMaxDataSizeSet is a free log retrieval operation binding the contract event 0x65923e04419dc810d0ea08a94a7f608d4c4d949818d95c3788f895e575dd2064.
//
// Solidity: event XMsgMaxDataSizeSet(uint16 size)
func (_OmniPortal *OmniPortalFilterer) FilterXMsgMaxDataSizeSet(opts *bind.FilterOpts) (*OmniPortalXMsgMaxDataSizeSetIterator, error) {

	logs, sub, err := _OmniPortal.contract.FilterLogs(opts, "XMsgMaxDataSizeSet")
	if err != nil {
		return nil, err
	}
	return &OmniPortalXMsgMaxDataSizeSetIterator{contract: _OmniPortal.contract, event: "XMsgMaxDataSizeSet", logs: logs, sub: sub}, nil
}

// WatchXMsgMaxDataSizeSet is a free log subscription operation binding the contract event 0x65923e04419dc810d0ea08a94a7f608d4c4d949818d95c3788f895e575dd2064.
//
// Solidity: event XMsgMaxDataSizeSet(uint16 size)
func (_OmniPortal *OmniPortalFilterer) WatchXMsgMaxDataSizeSet(opts *bind.WatchOpts, sink chan<- *OmniPortalXMsgMaxDataSizeSet) (event.Subscription, error) {

	logs, sub, err := _OmniPortal.contract.WatchLogs(opts, "XMsgMaxDataSizeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniPortalXMsgMaxDataSizeSet)
				if err := _OmniPortal.contract.UnpackLog(event, "XMsgMaxDataSizeSet", log); err != nil {
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

// ParseXMsgMaxDataSizeSet is a log parse operation binding the contract event 0x65923e04419dc810d0ea08a94a7f608d4c4d949818d95c3788f895e575dd2064.
//
// Solidity: event XMsgMaxDataSizeSet(uint16 size)
func (_OmniPortal *OmniPortalFilterer) ParseXMsgMaxDataSizeSet(log types.Log) (*OmniPortalXMsgMaxDataSizeSet, error) {
	event := new(OmniPortalXMsgMaxDataSizeSet)
	if err := _OmniPortal.contract.UnpackLog(event, "XMsgMaxDataSizeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniPortalXMsgMaxGasLimitSetIterator is returned from FilterXMsgMaxGasLimitSet and is used to iterate over the raw logs and unpacked data for XMsgMaxGasLimitSet events raised by the OmniPortal contract.
type OmniPortalXMsgMaxGasLimitSetIterator struct {
	Event *OmniPortalXMsgMaxGasLimitSet // Event containing the contract specifics and raw log

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
func (it *OmniPortalXMsgMaxGasLimitSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniPortalXMsgMaxGasLimitSet)
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
		it.Event = new(OmniPortalXMsgMaxGasLimitSet)
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
func (it *OmniPortalXMsgMaxGasLimitSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniPortalXMsgMaxGasLimitSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniPortalXMsgMaxGasLimitSet represents a XMsgMaxGasLimitSet event raised by the OmniPortal contract.
type OmniPortalXMsgMaxGasLimitSet struct {
	GasLimit uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterXMsgMaxGasLimitSet is a free log retrieval operation binding the contract event 0x1153561ac5effc2926ba6c612f86a397c997bc43dfbfc718da08065be0c5fe4d.
//
// Solidity: event XMsgMaxGasLimitSet(uint64 gasLimit)
func (_OmniPortal *OmniPortalFilterer) FilterXMsgMaxGasLimitSet(opts *bind.FilterOpts) (*OmniPortalXMsgMaxGasLimitSetIterator, error) {

	logs, sub, err := _OmniPortal.contract.FilterLogs(opts, "XMsgMaxGasLimitSet")
	if err != nil {
		return nil, err
	}
	return &OmniPortalXMsgMaxGasLimitSetIterator{contract: _OmniPortal.contract, event: "XMsgMaxGasLimitSet", logs: logs, sub: sub}, nil
}

// WatchXMsgMaxGasLimitSet is a free log subscription operation binding the contract event 0x1153561ac5effc2926ba6c612f86a397c997bc43dfbfc718da08065be0c5fe4d.
//
// Solidity: event XMsgMaxGasLimitSet(uint64 gasLimit)
func (_OmniPortal *OmniPortalFilterer) WatchXMsgMaxGasLimitSet(opts *bind.WatchOpts, sink chan<- *OmniPortalXMsgMaxGasLimitSet) (event.Subscription, error) {

	logs, sub, err := _OmniPortal.contract.WatchLogs(opts, "XMsgMaxGasLimitSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniPortalXMsgMaxGasLimitSet)
				if err := _OmniPortal.contract.UnpackLog(event, "XMsgMaxGasLimitSet", log); err != nil {
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

// ParseXMsgMaxGasLimitSet is a log parse operation binding the contract event 0x1153561ac5effc2926ba6c612f86a397c997bc43dfbfc718da08065be0c5fe4d.
//
// Solidity: event XMsgMaxGasLimitSet(uint64 gasLimit)
func (_OmniPortal *OmniPortalFilterer) ParseXMsgMaxGasLimitSet(log types.Log) (*OmniPortalXMsgMaxGasLimitSet, error) {
	event := new(OmniPortalXMsgMaxGasLimitSet)
	if err := _OmniPortal.contract.UnpackLog(event, "XMsgMaxGasLimitSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniPortalXMsgMinGasLimitSetIterator is returned from FilterXMsgMinGasLimitSet and is used to iterate over the raw logs and unpacked data for XMsgMinGasLimitSet events raised by the OmniPortal contract.
type OmniPortalXMsgMinGasLimitSetIterator struct {
	Event *OmniPortalXMsgMinGasLimitSet // Event containing the contract specifics and raw log

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
func (it *OmniPortalXMsgMinGasLimitSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniPortalXMsgMinGasLimitSet)
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
		it.Event = new(OmniPortalXMsgMinGasLimitSet)
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
func (it *OmniPortalXMsgMinGasLimitSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniPortalXMsgMinGasLimitSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniPortalXMsgMinGasLimitSet represents a XMsgMinGasLimitSet event raised by the OmniPortal contract.
type OmniPortalXMsgMinGasLimitSet struct {
	GasLimit uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterXMsgMinGasLimitSet is a free log retrieval operation binding the contract event 0x8c852a6291aa436654b167353bca4a4b0c3d024c7562cb5082e7c869bddabf3e.
//
// Solidity: event XMsgMinGasLimitSet(uint64 gasLimit)
func (_OmniPortal *OmniPortalFilterer) FilterXMsgMinGasLimitSet(opts *bind.FilterOpts) (*OmniPortalXMsgMinGasLimitSetIterator, error) {

	logs, sub, err := _OmniPortal.contract.FilterLogs(opts, "XMsgMinGasLimitSet")
	if err != nil {
		return nil, err
	}
	return &OmniPortalXMsgMinGasLimitSetIterator{contract: _OmniPortal.contract, event: "XMsgMinGasLimitSet", logs: logs, sub: sub}, nil
}

// WatchXMsgMinGasLimitSet is a free log subscription operation binding the contract event 0x8c852a6291aa436654b167353bca4a4b0c3d024c7562cb5082e7c869bddabf3e.
//
// Solidity: event XMsgMinGasLimitSet(uint64 gasLimit)
func (_OmniPortal *OmniPortalFilterer) WatchXMsgMinGasLimitSet(opts *bind.WatchOpts, sink chan<- *OmniPortalXMsgMinGasLimitSet) (event.Subscription, error) {

	logs, sub, err := _OmniPortal.contract.WatchLogs(opts, "XMsgMinGasLimitSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniPortalXMsgMinGasLimitSet)
				if err := _OmniPortal.contract.UnpackLog(event, "XMsgMinGasLimitSet", log); err != nil {
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

// ParseXMsgMinGasLimitSet is a log parse operation binding the contract event 0x8c852a6291aa436654b167353bca4a4b0c3d024c7562cb5082e7c869bddabf3e.
//
// Solidity: event XMsgMinGasLimitSet(uint64 gasLimit)
func (_OmniPortal *OmniPortalFilterer) ParseXMsgMinGasLimitSet(log types.Log) (*OmniPortalXMsgMinGasLimitSet, error) {
	event := new(OmniPortalXMsgMinGasLimitSet)
	if err := _OmniPortal.contract.UnpackLog(event, "XMsgMinGasLimitSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniPortalXReceiptIterator is returned from FilterXReceipt and is used to iterate over the raw logs and unpacked data for XReceipt events raised by the OmniPortal contract.
type OmniPortalXReceiptIterator struct {
	Event *OmniPortalXReceipt // Event containing the contract specifics and raw log

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
func (it *OmniPortalXReceiptIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniPortalXReceipt)
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
		it.Event = new(OmniPortalXReceipt)
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
func (it *OmniPortalXReceiptIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniPortalXReceiptIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniPortalXReceipt represents a XReceipt event raised by the OmniPortal contract.
type OmniPortalXReceipt struct {
	SourceChainId uint64
	ShardId       uint64
	Offset        uint64
	GasUsed       *big.Int
	Relayer       common.Address
	Success       bool
	Error         []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterXReceipt is a free log retrieval operation binding the contract event 0x8277cab1f0fa69b34674f64a7d43f242b0bacece6f5b7e8652f1e0d88a9b873b.
//
// Solidity: event XReceipt(uint64 indexed sourceChainId, uint64 indexed shardId, uint64 indexed offset, uint256 gasUsed, address relayer, bool success, bytes error)
func (_OmniPortal *OmniPortalFilterer) FilterXReceipt(opts *bind.FilterOpts, sourceChainId []uint64, shardId []uint64, offset []uint64) (*OmniPortalXReceiptIterator, error) {

	var sourceChainIdRule []interface{}
	for _, sourceChainIdItem := range sourceChainId {
		sourceChainIdRule = append(sourceChainIdRule, sourceChainIdItem)
	}
	var shardIdRule []interface{}
	for _, shardIdItem := range shardId {
		shardIdRule = append(shardIdRule, shardIdItem)
	}
	var offsetRule []interface{}
	for _, offsetItem := range offset {
		offsetRule = append(offsetRule, offsetItem)
	}

	logs, sub, err := _OmniPortal.contract.FilterLogs(opts, "XReceipt", sourceChainIdRule, shardIdRule, offsetRule)
	if err != nil {
		return nil, err
	}
	return &OmniPortalXReceiptIterator{contract: _OmniPortal.contract, event: "XReceipt", logs: logs, sub: sub}, nil
}

// WatchXReceipt is a free log subscription operation binding the contract event 0x8277cab1f0fa69b34674f64a7d43f242b0bacece6f5b7e8652f1e0d88a9b873b.
//
// Solidity: event XReceipt(uint64 indexed sourceChainId, uint64 indexed shardId, uint64 indexed offset, uint256 gasUsed, address relayer, bool success, bytes error)
func (_OmniPortal *OmniPortalFilterer) WatchXReceipt(opts *bind.WatchOpts, sink chan<- *OmniPortalXReceipt, sourceChainId []uint64, shardId []uint64, offset []uint64) (event.Subscription, error) {

	var sourceChainIdRule []interface{}
	for _, sourceChainIdItem := range sourceChainId {
		sourceChainIdRule = append(sourceChainIdRule, sourceChainIdItem)
	}
	var shardIdRule []interface{}
	for _, shardIdItem := range shardId {
		shardIdRule = append(shardIdRule, shardIdItem)
	}
	var offsetRule []interface{}
	for _, offsetItem := range offset {
		offsetRule = append(offsetRule, offsetItem)
	}

	logs, sub, err := _OmniPortal.contract.WatchLogs(opts, "XReceipt", sourceChainIdRule, shardIdRule, offsetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniPortalXReceipt)
				if err := _OmniPortal.contract.UnpackLog(event, "XReceipt", log); err != nil {
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

// ParseXReceipt is a log parse operation binding the contract event 0x8277cab1f0fa69b34674f64a7d43f242b0bacece6f5b7e8652f1e0d88a9b873b.
//
// Solidity: event XReceipt(uint64 indexed sourceChainId, uint64 indexed shardId, uint64 indexed offset, uint256 gasUsed, address relayer, bool success, bytes error)
func (_OmniPortal *OmniPortalFilterer) ParseXReceipt(log types.Log) (*OmniPortalXReceipt, error) {
	event := new(OmniPortalXReceipt)
	if err := _OmniPortal.contract.UnpackLog(event, "XReceipt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniPortalXReceiptMaxErrorSizeSetIterator is returned from FilterXReceiptMaxErrorSizeSet and is used to iterate over the raw logs and unpacked data for XReceiptMaxErrorSizeSet events raised by the OmniPortal contract.
type OmniPortalXReceiptMaxErrorSizeSetIterator struct {
	Event *OmniPortalXReceiptMaxErrorSizeSet // Event containing the contract specifics and raw log

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
func (it *OmniPortalXReceiptMaxErrorSizeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniPortalXReceiptMaxErrorSizeSet)
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
		it.Event = new(OmniPortalXReceiptMaxErrorSizeSet)
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
func (it *OmniPortalXReceiptMaxErrorSizeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniPortalXReceiptMaxErrorSizeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniPortalXReceiptMaxErrorSizeSet represents a XReceiptMaxErrorSizeSet event raised by the OmniPortal contract.
type OmniPortalXReceiptMaxErrorSizeSet struct {
	Size uint16
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterXReceiptMaxErrorSizeSet is a free log retrieval operation binding the contract event 0x620bbea084306b66a8cc6b5b63830d6b3874f9d2438914e259ffd5065c33f7b0.
//
// Solidity: event XReceiptMaxErrorSizeSet(uint16 size)
func (_OmniPortal *OmniPortalFilterer) FilterXReceiptMaxErrorSizeSet(opts *bind.FilterOpts) (*OmniPortalXReceiptMaxErrorSizeSetIterator, error) {

	logs, sub, err := _OmniPortal.contract.FilterLogs(opts, "XReceiptMaxErrorSizeSet")
	if err != nil {
		return nil, err
	}
	return &OmniPortalXReceiptMaxErrorSizeSetIterator{contract: _OmniPortal.contract, event: "XReceiptMaxErrorSizeSet", logs: logs, sub: sub}, nil
}

// WatchXReceiptMaxErrorSizeSet is a free log subscription operation binding the contract event 0x620bbea084306b66a8cc6b5b63830d6b3874f9d2438914e259ffd5065c33f7b0.
//
// Solidity: event XReceiptMaxErrorSizeSet(uint16 size)
func (_OmniPortal *OmniPortalFilterer) WatchXReceiptMaxErrorSizeSet(opts *bind.WatchOpts, sink chan<- *OmniPortalXReceiptMaxErrorSizeSet) (event.Subscription, error) {

	logs, sub, err := _OmniPortal.contract.WatchLogs(opts, "XReceiptMaxErrorSizeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniPortalXReceiptMaxErrorSizeSet)
				if err := _OmniPortal.contract.UnpackLog(event, "XReceiptMaxErrorSizeSet", log); err != nil {
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

// ParseXReceiptMaxErrorSizeSet is a log parse operation binding the contract event 0x620bbea084306b66a8cc6b5b63830d6b3874f9d2438914e259ffd5065c33f7b0.
//
// Solidity: event XReceiptMaxErrorSizeSet(uint16 size)
func (_OmniPortal *OmniPortalFilterer) ParseXReceiptMaxErrorSizeSet(log types.Log) (*OmniPortalXReceiptMaxErrorSizeSet, error) {
	event := new(OmniPortalXReceiptMaxErrorSizeSet)
	if err := _OmniPortal.contract.UnpackLog(event, "XReceiptMaxErrorSizeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniPortalXSubValsetCutoffSetIterator is returned from FilterXSubValsetCutoffSet and is used to iterate over the raw logs and unpacked data for XSubValsetCutoffSet events raised by the OmniPortal contract.
type OmniPortalXSubValsetCutoffSetIterator struct {
	Event *OmniPortalXSubValsetCutoffSet // Event containing the contract specifics and raw log

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
func (it *OmniPortalXSubValsetCutoffSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniPortalXSubValsetCutoffSet)
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
		it.Event = new(OmniPortalXSubValsetCutoffSet)
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
func (it *OmniPortalXSubValsetCutoffSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniPortalXSubValsetCutoffSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniPortalXSubValsetCutoffSet represents a XSubValsetCutoffSet event raised by the OmniPortal contract.
type OmniPortalXSubValsetCutoffSet struct {
	Cutoff uint8
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterXSubValsetCutoffSet is a free log retrieval operation binding the contract event 0x1683dc51426224f6e37a3b41dd5849e2db1bfe22366d1d913fa0ef6f757e828f.
//
// Solidity: event XSubValsetCutoffSet(uint8 cutoff)
func (_OmniPortal *OmniPortalFilterer) FilterXSubValsetCutoffSet(opts *bind.FilterOpts) (*OmniPortalXSubValsetCutoffSetIterator, error) {

	logs, sub, err := _OmniPortal.contract.FilterLogs(opts, "XSubValsetCutoffSet")
	if err != nil {
		return nil, err
	}
	return &OmniPortalXSubValsetCutoffSetIterator{contract: _OmniPortal.contract, event: "XSubValsetCutoffSet", logs: logs, sub: sub}, nil
}

// WatchXSubValsetCutoffSet is a free log subscription operation binding the contract event 0x1683dc51426224f6e37a3b41dd5849e2db1bfe22366d1d913fa0ef6f757e828f.
//
// Solidity: event XSubValsetCutoffSet(uint8 cutoff)
func (_OmniPortal *OmniPortalFilterer) WatchXSubValsetCutoffSet(opts *bind.WatchOpts, sink chan<- *OmniPortalXSubValsetCutoffSet) (event.Subscription, error) {

	logs, sub, err := _OmniPortal.contract.WatchLogs(opts, "XSubValsetCutoffSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniPortalXSubValsetCutoffSet)
				if err := _OmniPortal.contract.UnpackLog(event, "XSubValsetCutoffSet", log); err != nil {
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

// ParseXSubValsetCutoffSet is a log parse operation binding the contract event 0x1683dc51426224f6e37a3b41dd5849e2db1bfe22366d1d913fa0ef6f757e828f.
//
// Solidity: event XSubValsetCutoffSet(uint8 cutoff)
func (_OmniPortal *OmniPortalFilterer) ParseXSubValsetCutoffSet(log types.Log) (*OmniPortalXSubValsetCutoffSet, error) {
	event := new(OmniPortalXSubValsetCutoffSet)
	if err := _OmniPortal.contract.UnpackLog(event, "XSubValsetCutoffSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniPortalXSubmitFromPausedIterator is returned from FilterXSubmitFromPaused and is used to iterate over the raw logs and unpacked data for XSubmitFromPaused events raised by the OmniPortal contract.
type OmniPortalXSubmitFromPausedIterator struct {
	Event *OmniPortalXSubmitFromPaused // Event containing the contract specifics and raw log

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
func (it *OmniPortalXSubmitFromPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniPortalXSubmitFromPaused)
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
		it.Event = new(OmniPortalXSubmitFromPaused)
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
func (it *OmniPortalXSubmitFromPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniPortalXSubmitFromPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniPortalXSubmitFromPaused represents a XSubmitFromPaused event raised by the OmniPortal contract.
type OmniPortalXSubmitFromPaused struct {
	ChainId uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterXSubmitFromPaused is a free log retrieval operation binding the contract event 0xab78810a0515df65f9f10bfbcb92d03d5df71d9fd3b9414e9ad831a5117d6daa.
//
// Solidity: event XSubmitFromPaused(uint64 indexed chainId)
func (_OmniPortal *OmniPortalFilterer) FilterXSubmitFromPaused(opts *bind.FilterOpts, chainId []uint64) (*OmniPortalXSubmitFromPausedIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _OmniPortal.contract.FilterLogs(opts, "XSubmitFromPaused", chainIdRule)
	if err != nil {
		return nil, err
	}
	return &OmniPortalXSubmitFromPausedIterator{contract: _OmniPortal.contract, event: "XSubmitFromPaused", logs: logs, sub: sub}, nil
}

// WatchXSubmitFromPaused is a free log subscription operation binding the contract event 0xab78810a0515df65f9f10bfbcb92d03d5df71d9fd3b9414e9ad831a5117d6daa.
//
// Solidity: event XSubmitFromPaused(uint64 indexed chainId)
func (_OmniPortal *OmniPortalFilterer) WatchXSubmitFromPaused(opts *bind.WatchOpts, sink chan<- *OmniPortalXSubmitFromPaused, chainId []uint64) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _OmniPortal.contract.WatchLogs(opts, "XSubmitFromPaused", chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniPortalXSubmitFromPaused)
				if err := _OmniPortal.contract.UnpackLog(event, "XSubmitFromPaused", log); err != nil {
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

// ParseXSubmitFromPaused is a log parse operation binding the contract event 0xab78810a0515df65f9f10bfbcb92d03d5df71d9fd3b9414e9ad831a5117d6daa.
//
// Solidity: event XSubmitFromPaused(uint64 indexed chainId)
func (_OmniPortal *OmniPortalFilterer) ParseXSubmitFromPaused(log types.Log) (*OmniPortalXSubmitFromPaused, error) {
	event := new(OmniPortalXSubmitFromPaused)
	if err := _OmniPortal.contract.UnpackLog(event, "XSubmitFromPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniPortalXSubmitFromUnpausedIterator is returned from FilterXSubmitFromUnpaused and is used to iterate over the raw logs and unpacked data for XSubmitFromUnpaused events raised by the OmniPortal contract.
type OmniPortalXSubmitFromUnpausedIterator struct {
	Event *OmniPortalXSubmitFromUnpaused // Event containing the contract specifics and raw log

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
func (it *OmniPortalXSubmitFromUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniPortalXSubmitFromUnpaused)
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
		it.Event = new(OmniPortalXSubmitFromUnpaused)
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
func (it *OmniPortalXSubmitFromUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniPortalXSubmitFromUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniPortalXSubmitFromUnpaused represents a XSubmitFromUnpaused event raised by the OmniPortal contract.
type OmniPortalXSubmitFromUnpaused struct {
	ChainId uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterXSubmitFromUnpaused is a free log retrieval operation binding the contract event 0xc551305d9bd408be4327b7f8aba28b04ccf6b6c76925392d195ecf9cc764294d.
//
// Solidity: event XSubmitFromUnpaused(uint64 indexed chainId)
func (_OmniPortal *OmniPortalFilterer) FilterXSubmitFromUnpaused(opts *bind.FilterOpts, chainId []uint64) (*OmniPortalXSubmitFromUnpausedIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _OmniPortal.contract.FilterLogs(opts, "XSubmitFromUnpaused", chainIdRule)
	if err != nil {
		return nil, err
	}
	return &OmniPortalXSubmitFromUnpausedIterator{contract: _OmniPortal.contract, event: "XSubmitFromUnpaused", logs: logs, sub: sub}, nil
}

// WatchXSubmitFromUnpaused is a free log subscription operation binding the contract event 0xc551305d9bd408be4327b7f8aba28b04ccf6b6c76925392d195ecf9cc764294d.
//
// Solidity: event XSubmitFromUnpaused(uint64 indexed chainId)
func (_OmniPortal *OmniPortalFilterer) WatchXSubmitFromUnpaused(opts *bind.WatchOpts, sink chan<- *OmniPortalXSubmitFromUnpaused, chainId []uint64) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _OmniPortal.contract.WatchLogs(opts, "XSubmitFromUnpaused", chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniPortalXSubmitFromUnpaused)
				if err := _OmniPortal.contract.UnpackLog(event, "XSubmitFromUnpaused", log); err != nil {
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

// ParseXSubmitFromUnpaused is a log parse operation binding the contract event 0xc551305d9bd408be4327b7f8aba28b04ccf6b6c76925392d195ecf9cc764294d.
//
// Solidity: event XSubmitFromUnpaused(uint64 indexed chainId)
func (_OmniPortal *OmniPortalFilterer) ParseXSubmitFromUnpaused(log types.Log) (*OmniPortalXSubmitFromUnpaused, error) {
	event := new(OmniPortalXSubmitFromUnpaused)
	if err := _OmniPortal.contract.UnpackLog(event, "XSubmitFromUnpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniPortalXSubmitPausedIterator is returned from FilterXSubmitPaused and is used to iterate over the raw logs and unpacked data for XSubmitPaused events raised by the OmniPortal contract.
type OmniPortalXSubmitPausedIterator struct {
	Event *OmniPortalXSubmitPaused // Event containing the contract specifics and raw log

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
func (it *OmniPortalXSubmitPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniPortalXSubmitPaused)
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
		it.Event = new(OmniPortalXSubmitPaused)
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
func (it *OmniPortalXSubmitPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniPortalXSubmitPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniPortalXSubmitPaused represents a XSubmitPaused event raised by the OmniPortal contract.
type OmniPortalXSubmitPaused struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterXSubmitPaused is a free log retrieval operation binding the contract event 0x3d0f9c56dac46156a2db0aa09ee7804770ad9fc9549d21023164f22d69475ed8.
//
// Solidity: event XSubmitPaused()
func (_OmniPortal *OmniPortalFilterer) FilterXSubmitPaused(opts *bind.FilterOpts) (*OmniPortalXSubmitPausedIterator, error) {

	logs, sub, err := _OmniPortal.contract.FilterLogs(opts, "XSubmitPaused")
	if err != nil {
		return nil, err
	}
	return &OmniPortalXSubmitPausedIterator{contract: _OmniPortal.contract, event: "XSubmitPaused", logs: logs, sub: sub}, nil
}

// WatchXSubmitPaused is a free log subscription operation binding the contract event 0x3d0f9c56dac46156a2db0aa09ee7804770ad9fc9549d21023164f22d69475ed8.
//
// Solidity: event XSubmitPaused()
func (_OmniPortal *OmniPortalFilterer) WatchXSubmitPaused(opts *bind.WatchOpts, sink chan<- *OmniPortalXSubmitPaused) (event.Subscription, error) {

	logs, sub, err := _OmniPortal.contract.WatchLogs(opts, "XSubmitPaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniPortalXSubmitPaused)
				if err := _OmniPortal.contract.UnpackLog(event, "XSubmitPaused", log); err != nil {
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

// ParseXSubmitPaused is a log parse operation binding the contract event 0x3d0f9c56dac46156a2db0aa09ee7804770ad9fc9549d21023164f22d69475ed8.
//
// Solidity: event XSubmitPaused()
func (_OmniPortal *OmniPortalFilterer) ParseXSubmitPaused(log types.Log) (*OmniPortalXSubmitPaused, error) {
	event := new(OmniPortalXSubmitPaused)
	if err := _OmniPortal.contract.UnpackLog(event, "XSubmitPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniPortalXSubmitUnpausedIterator is returned from FilterXSubmitUnpaused and is used to iterate over the raw logs and unpacked data for XSubmitUnpaused events raised by the OmniPortal contract.
type OmniPortalXSubmitUnpausedIterator struct {
	Event *OmniPortalXSubmitUnpaused // Event containing the contract specifics and raw log

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
func (it *OmniPortalXSubmitUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniPortalXSubmitUnpaused)
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
		it.Event = new(OmniPortalXSubmitUnpaused)
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
func (it *OmniPortalXSubmitUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniPortalXSubmitUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniPortalXSubmitUnpaused represents a XSubmitUnpaused event raised by the OmniPortal contract.
type OmniPortalXSubmitUnpaused struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterXSubmitUnpaused is a free log retrieval operation binding the contract event 0x2cb9d71d4c31860b70e9b707c69aa2f5953e03474f00cfcfff205c4745f82875.
//
// Solidity: event XSubmitUnpaused()
func (_OmniPortal *OmniPortalFilterer) FilterXSubmitUnpaused(opts *bind.FilterOpts) (*OmniPortalXSubmitUnpausedIterator, error) {

	logs, sub, err := _OmniPortal.contract.FilterLogs(opts, "XSubmitUnpaused")
	if err != nil {
		return nil, err
	}
	return &OmniPortalXSubmitUnpausedIterator{contract: _OmniPortal.contract, event: "XSubmitUnpaused", logs: logs, sub: sub}, nil
}

// WatchXSubmitUnpaused is a free log subscription operation binding the contract event 0x2cb9d71d4c31860b70e9b707c69aa2f5953e03474f00cfcfff205c4745f82875.
//
// Solidity: event XSubmitUnpaused()
func (_OmniPortal *OmniPortalFilterer) WatchXSubmitUnpaused(opts *bind.WatchOpts, sink chan<- *OmniPortalXSubmitUnpaused) (event.Subscription, error) {

	logs, sub, err := _OmniPortal.contract.WatchLogs(opts, "XSubmitUnpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniPortalXSubmitUnpaused)
				if err := _OmniPortal.contract.UnpackLog(event, "XSubmitUnpaused", log); err != nil {
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

// ParseXSubmitUnpaused is a log parse operation binding the contract event 0x2cb9d71d4c31860b70e9b707c69aa2f5953e03474f00cfcfff205c4745f82875.
//
// Solidity: event XSubmitUnpaused()
func (_OmniPortal *OmniPortalFilterer) ParseXSubmitUnpaused(log types.Log) (*OmniPortalXSubmitUnpaused, error) {
	event := new(OmniPortalXSubmitUnpaused)
	if err := _OmniPortal.contract.UnpackLog(event, "XSubmitUnpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
