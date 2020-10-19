// nolint
// autogenerated code using github.com/rigelrozanski/multitool
// aliases generated for the following subdirectories:
// ALIASGEN: github.com/hbtc-chain/bhchain/x/auth/types
package custodianunit

import (
	"github.com/hbtc-chain/bhchain/x/custodianunit/exported"
	"github.com/hbtc-chain/bhchain/x/custodianunit/types"
)

const (
	ModuleName                    = types.ModuleName
	StoreKey                      = types.StoreKey
	FeeCollectorName              = types.FeeCollectorName
	QuerierRoute                  = types.QuerierRoute
	DefaultParamspace             = types.DefaultParamspace
	DefaultMaxMemoCharacters      = types.DefaultMaxMemoCharacters
	DefaultTxSigLimit             = types.DefaultTxSigLimit
	DefaultTxSizeCostPerByte      = types.DefaultTxSizeCostPerByte
	DefaultSigVerifyCostED25519   = types.DefaultSigVerifyCostED25519
	DefaultSigVerifyCostSecp256k1 = types.DefaultSigVerifyCostSecp256k1
	QueryCU                       = types.QueryCU
	QueryOpCU                     = types.QueryOpCU
)

var (
	// functions aliases
	NewBaseCU             = types.NewBaseCU
	ProtoBaseCU           = types.ProtoBaseCU
	NewBaseCUWithAddress  = types.NewBaseCUWithAddress
	RegisterCodec         = types.RegisterCodec
	NewGenesisState       = types.NewGenesisState
	DefaultGenesisState   = types.DefaultGenesisState
	ValidateGenesis       = types.ValidateGenesis
	AddressStoreKey       = types.AddressStoreKey
	NewParams             = types.NewParams
	ParamKeyTable         = types.ParamKeyTable
	DefaultParams         = types.DefaultParams
	NewQueryAccountParams = types.NewQueryCUParams
	NewStdTx              = types.NewStdTx
	CountSubKeys          = types.CountSubKeys
	NewStdFee             = types.NewStdFee
	StdSignBytes          = types.StdSignBytes
	DefaultTxDecoder      = types.DefaultTxDecoder
	DefaultTxEncoder      = types.DefaultTxEncoder
	NewTxBuilder          = types.NewTxBuilder
	NewTxBuilderFromCLI   = types.NewTxBuilderFromCLI
	MakeSignature         = types.MakeSignature
	NewAccountRetriever   = types.NewCURetriever

	// variable aliases
	ModuleCdc                 = types.ModuleCdc
	AddressStoreKeyPrefix     = types.AddressStoreKeyPrefix
	GlobalCUNumberKey         = types.GlobalCUNumberKey
	KeyMaxMemoCharacters      = types.KeyMaxMemoCharacters
	KeyTxSigLimit             = types.KeyTxSigLimit
	KeyTxSizeCostPerByte      = types.KeyTxSizeCostPerByte
	KeySigVerifyCostED25519   = types.KeySigVerifyCostED25519
	KeySigVerifyCostSecp256k1 = types.KeySigVerifyCostSecp256k1

	QueryOpCUForTest = queryOpCUInfo
	QueryCUForTest   = queryCU
)

type (
	CU                 = exported.CustodianUnit
	VestingAccount     = exported.VestingCU
	BaseCU             = types.BaseCU
	GenesisState       = types.GenesisState
	Params             = types.Params
	QueryAccountParams = types.QueryCUParams
	QueryOpCUParams    = types.QueryOpCUParams
	StdSignMsg         = types.StdSignMsg
	StdTx              = types.StdTx
	StdFee             = types.StdFee
	StdSignDoc         = types.StdSignDoc
	StdSignature       = types.StdSignature
	TxBuilder          = types.TxBuilder
)