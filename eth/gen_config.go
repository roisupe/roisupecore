// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package eth

import (
	"math/big"
	"time"

	"github.com/gochain/gochain/v3/common"
	"github.com/gochain/gochain/v3/common/hexutil"
	"github.com/gochain/gochain/v3/core"
	"github.com/gochain/gochain/v3/eth/downloader"
	"github.com/gochain/gochain/v3/eth/gasprice"
)

var _ = (*configMarshaling)(nil)

// MarshalTOML marshals as TOML.
func (c Config) MarshalTOML() (interface{}, error) {
	type Config struct {
		Genesis                 *core.Genesis `toml:",omitempty"`
		NetworkId               uint64
		SyncMode                downloader.SyncMode
		NoPruning               bool
		LightServ               int  `toml:",omitempty"`
		LightPeers              int  `toml:",omitempty"`
		SkipBcVersionCheck      bool `toml:"-"`
		DatabaseHandles         int  `toml:"-"`
		DatabaseCache           int
		TrieCache               int
		TrieTimeout             time.Duration
		Etherbase               common.Address `toml:",omitempty"`
		MinerExtraData          hexutil.Bytes  `toml:",omitempty"`
		MinerGasPrice           *big.Int
		TxPool                  core.TxPoolConfig
		GPO                     gasprice.Config
		EnablePreimageRecording bool
		DocRoot                 string `toml:"-"`
	}
	var enc Config
	enc.Genesis = c.Genesis
	enc.NetworkId = c.NetworkId
	enc.SyncMode = c.SyncMode
	enc.NoPruning = c.NoPruning
	enc.LightServ = c.LightServ
	enc.LightPeers = c.LightPeers
	enc.SkipBcVersionCheck = c.SkipBcVersionCheck
	enc.DatabaseHandles = c.DatabaseHandles
	enc.DatabaseCache = c.DatabaseCache
	enc.TrieCache = c.TrieCache
	enc.TrieTimeout = c.TrieTimeout
	enc.Etherbase = c.Etherbase
	enc.MinerExtraData = c.MinerExtraData
	enc.MinerGasPrice = c.MinerGasPrice
	enc.TxPool = c.TxPool
	enc.GPO = c.GPO
	enc.EnablePreimageRecording = c.EnablePreimageRecording
	enc.DocRoot = c.DocRoot
	// enc.Archive = c.Archive
	return &enc, nil
}

// UnmarshalTOML unmarshals from TOML.
func (c *Config) UnmarshalTOML(unmarshal func(interface{}) error) error {
	type Config struct {
		Genesis                 *core.Genesis `toml:",omitempty"`
		NetworkId               *uint64
		SyncMode                *downloader.SyncMode
		NoPruning               *bool
		LightServ               *int  `toml:",omitempty"`
		LightPeers              *int  `toml:",omitempty"`
		SkipBcVersionCheck      *bool `toml:"-"`
		DatabaseHandles         *int  `toml:"-"`
		DatabaseCache           *int
		TrieCache               *int
		TrieTimeout             *time.Duration
		Etherbase               *common.Address `toml:",omitempty"`
		MinerExtraData          *hexutil.Bytes  `toml:",omitempty"`
		MinerGasPrice           *big.Int
		TxPool                  *core.TxPoolConfig
		GPO                     *gasprice.Config
		EnablePreimageRecording *bool
		DocRoot                 *string `toml:"-"`
		// Archive                 *archive.Config `toml:",omitempty"`
	}
	var dec Config
	if err := unmarshal(&dec); err != nil {
		return err
	}
	if dec.Genesis != nil {
		c.Genesis = dec.Genesis
	}
	if dec.NetworkId != nil {
		c.NetworkId = *dec.NetworkId
	}
	if dec.SyncMode != nil {
		c.SyncMode = *dec.SyncMode
	}
	if dec.NoPruning != nil {
		c.NoPruning = *dec.NoPruning
	}
	if dec.LightServ != nil {
		c.LightServ = *dec.LightServ
	}
	if dec.LightPeers != nil {
		c.LightPeers = *dec.LightPeers
	}
	if dec.SkipBcVersionCheck != nil {
		c.SkipBcVersionCheck = *dec.SkipBcVersionCheck
	}
	if dec.DatabaseHandles != nil {
		c.DatabaseHandles = *dec.DatabaseHandles
	}
	if dec.DatabaseCache != nil {
		c.DatabaseCache = *dec.DatabaseCache
	}
	if dec.TrieCache != nil {
		c.TrieCache = *dec.TrieCache
	}
	if dec.TrieTimeout != nil {
		c.TrieTimeout = *dec.TrieTimeout
	}
	if dec.Etherbase != nil {
		c.Etherbase = *dec.Etherbase
	}
	if dec.MinerExtraData != nil {
		c.MinerExtraData = *dec.MinerExtraData
	}
	if dec.MinerGasPrice != nil {
		c.MinerGasPrice = dec.MinerGasPrice
	}
	if dec.TxPool != nil {
		c.TxPool = *dec.TxPool
	}
	if dec.GPO != nil {
		c.GPO = *dec.GPO
	}
	if dec.EnablePreimageRecording != nil {
		c.EnablePreimageRecording = *dec.EnablePreimageRecording
	}
	if dec.DocRoot != nil {
		c.DocRoot = *dec.DocRoot
	}
	// if dec.Archive != nil {
	// 	c.Archive = *dec.Archive
	// }
	return nil
}
