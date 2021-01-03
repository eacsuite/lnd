package lnd

import (
	"github.com/eacsuite/eacd/chaincfg"
	bitcoinCfg "github.com/eacsuite/eacd/chaincfg"
	"github.com/eacsuite/eacd/chaincfg/chainhash"
	bitcoinWire "github.com/eacsuite/eacd/wire"
	"github.com/eacsuite/lnd/keychain"
	earthcoinCfg "github.com/eacsuite/eacd/chaincfg"
	earthcoinWire "github.com/eacsuite/eacd/wire"
)

// activeNetParams is a pointer to the parameters specific to the currently
// active bitcoin network.
var activeNetParams = earthcoinTestNetParams

// bitcoinNetParams couples the p2p parameters of a network with the
// corresponding RPC port of a daemon running on the particular network.
type bitcoinNetParams struct {
	*bitcoinCfg.Params
	rpcPort  string
	CoinType uint32
}

// earthcoinNetParams couples the p2p parameters of a network with the
// corresponding RPC port of a daemon running on the particular network.
type earthcoinNetParams struct {
	*earthcoinCfg.Params
	rpcPort  string
	CoinType uint32
}

// bitcoinTestNetParams contains parameters specific to the 3rd version of the
// test network.
var bitcoinTestNetParams = bitcoinNetParams{
	Params:   &bitcoinCfg.TestNet4Params,
	rpcPort:  "19334",
	CoinType: keychain.CoinTypeTestnet,
}

// bitcoinMainNetParams contains parameters specific to the current Bitcoin
// mainnet.
var bitcoinMainNetParams = bitcoinNetParams{
	Params:   &bitcoinCfg.MainNetParams,
	rpcPort:  "8334",
	CoinType: keychain.CoinTypeBitcoin,
}

// bitcoinSimNetParams contains parameters specific to the simulation test
// network.
var bitcoinSimNetParams = bitcoinNetParams{
	Params:   &bitcoinCfg.SimNetParams,
	rpcPort:  "18556",
	CoinType: keychain.CoinTypeTestnet,
}

// earthcoinSimNetParams contains parameters specific to the simulation test
// network.
var earthcoinSimNetParams = earthcoinNetParams{
	Params:   &earthcoinCfg.SimNetParams,
	rpcPort:  "18556",
	CoinType: keychain.CoinTypeTestnet,
}

// earthcoinTestNetParams contains parameters specific to the 4th version of the
// test network.
var earthcoinTestNetParams = earthcoinNetParams{
	Params:   &earthcoinCfg.TestNet4Params,
	rpcPort:  "19334",
	CoinType: keychain.CoinTypeTestnet,
}

// earthcoinMainNetParams contains the parameters specific to the current
// Earthcoin mainnet.
var earthcoinMainNetParams = earthcoinNetParams{
	Params:   &earthcoinCfg.MainNetParams,
	rpcPort:  "9334",
	CoinType: keychain.CoinTypeEarthcoin,
}

// earthcoinRegTestNetParams contains parameters specific to a local earthcoin
// regtest network.
var earthcoinRegTestNetParams = earthcoinNetParams{
	Params:   &earthcoinCfg.RegressionNetParams,
	rpcPort:  "18334",
	CoinType: keychain.CoinTypeTestnet,
}

// bitcoinRegTestNetParams contains parameters specific to a local bitcoin
// regtest network.
var bitcoinRegTestNetParams = bitcoinNetParams{
	Params:   &bitcoinCfg.RegressionNetParams,
	rpcPort:  "18334",
	CoinType: keychain.CoinTypeTestnet,
}

// applyEarthcoinParams applies the relevant chain configuration parameters that
// differ for earthcoin to the chain parameters typed for eacsuite derivation.
// This function is used in place of using something like interface{} to
// abstract over _which_ chain (or fork) the parameters are for.
func applyEarthcoinParams(params *bitcoinNetParams, earthcoinParams *earthcoinNetParams) {
	params.Name = earthcoinParams.Name
	params.Net = bitcoinWire.BitcoinNet(earthcoinParams.Net)
	params.DefaultPort = earthcoinParams.DefaultPort
	params.CoinbaseMaturity = earthcoinParams.CoinbaseMaturity

	copy(params.GenesisHash[:], earthcoinParams.GenesisHash[:])
	copy(params.GenesisBlock.Header.MerkleRoot[:],
		earthcoinParams.GenesisBlock.Header.MerkleRoot[:])
	params.GenesisBlock.Header.Version =
		earthcoinParams.GenesisBlock.Header.Version
	params.GenesisBlock.Header.Timestamp =
		earthcoinParams.GenesisBlock.Header.Timestamp
	params.GenesisBlock.Header.Bits =
		earthcoinParams.GenesisBlock.Header.Bits
	params.GenesisBlock.Header.Nonce =
		earthcoinParams.GenesisBlock.Header.Nonce
	params.GenesisBlock.Transactions[0].Version =
		earthcoinParams.GenesisBlock.Transactions[0].Version
	params.GenesisBlock.Transactions[0].LockTime =
		earthcoinParams.GenesisBlock.Transactions[0].LockTime
	params.GenesisBlock.Transactions[0].TxIn[0].Sequence =
		earthcoinParams.GenesisBlock.Transactions[0].TxIn[0].Sequence
	params.GenesisBlock.Transactions[0].TxIn[0].PreviousOutPoint.Index =
		earthcoinParams.GenesisBlock.Transactions[0].TxIn[0].PreviousOutPoint.Index
	copy(params.GenesisBlock.Transactions[0].TxIn[0].SignatureScript[:],
		earthcoinParams.GenesisBlock.Transactions[0].TxIn[0].SignatureScript[:])
	copy(params.GenesisBlock.Transactions[0].TxOut[0].PkScript[:],
		earthcoinParams.GenesisBlock.Transactions[0].TxOut[0].PkScript[:])
	params.GenesisBlock.Transactions[0].TxOut[0].Value =
		earthcoinParams.GenesisBlock.Transactions[0].TxOut[0].Value
	params.GenesisBlock.Transactions[0].TxIn[0].PreviousOutPoint.Hash =
		chainhash.Hash{}
	params.PowLimitBits = earthcoinParams.PowLimitBits
	params.PowLimit = earthcoinParams.PowLimit

	dnsSeeds := make([]chaincfg.DNSSeed, len(earthcoinParams.DNSSeeds))
	for i := 0; i < len(earthcoinParams.DNSSeeds); i++ {
		dnsSeeds[i] = chaincfg.DNSSeed{
			Host:         earthcoinParams.DNSSeeds[i].Host,
			HasFiltering: earthcoinParams.DNSSeeds[i].HasFiltering,
		}
	}
	params.DNSSeeds = dnsSeeds

	// Address encoding magics
	params.PubKeyHashAddrID = earthcoinParams.PubKeyHashAddrID
	params.ScriptHashAddrID = earthcoinParams.ScriptHashAddrID
	params.PrivateKeyID = earthcoinParams.PrivateKeyID
	params.WitnessPubKeyHashAddrID = earthcoinParams.WitnessPubKeyHashAddrID
	params.WitnessScriptHashAddrID = earthcoinParams.WitnessScriptHashAddrID
	params.Bech32HRPSegwit = earthcoinParams.Bech32HRPSegwit

	copy(params.HDPrivateKeyID[:], earthcoinParams.HDPrivateKeyID[:])
	copy(params.HDPublicKeyID[:], earthcoinParams.HDPublicKeyID[:])

	params.HDCoinType = earthcoinParams.HDCoinType

	checkPoints := make([]chaincfg.Checkpoint, len(earthcoinParams.Checkpoints))
	for i := 0; i < len(earthcoinParams.Checkpoints); i++ {
		var chainHash chainhash.Hash
		copy(chainHash[:], earthcoinParams.Checkpoints[i].Hash[:])

		checkPoints[i] = chaincfg.Checkpoint{
			Height: earthcoinParams.Checkpoints[i].Height,
			Hash:   &chainHash,
		}
	}
	params.Checkpoints = checkPoints

	params.rpcPort = earthcoinParams.rpcPort
	params.CoinType = earthcoinParams.CoinType
	chaincfg.Register(params.Params); // JMC
}

// isTestnet tests if the given params correspond to a testnet
// parameter configuration.
func isTestnet(params *earthcoinNetParams) bool {
	switch params.Params.Net {
	case earthcoinWire.TestNet4:
		return true
	default:
		return false
	}
}
