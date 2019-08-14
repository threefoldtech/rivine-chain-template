package main

import (
	"github.com/threefoldtech/rivine/extensions/minting"
	"github.com/threefoldtech/rivine/pkg/client"
	"github.com/threefoldtech/rivine/types"
)

// RegisterStandardTransactions registers the goldchain-specific transactions as required for the standard network.
func RegisterStandardTransactions(cli *client.CommandLineClient) {
	registerTransactions(cli)
}

// RegisterTestnetTransactions registers the goldchain-specific transactions as required for the test network.
func RegisterTestnetTransactions(cli *client.CommandLineClient) {
	registerTransactions(cli)
}

// RegisterDevnetTransactions registers the goldchain-specific transactions as required for the dev network.
func RegisterDevnetTransactions(cli *client.CommandLineClient) {
	registerTransactions(cli)
}

func registerTransactions(cli *client.CommandLineClient) {
	// create minting plugin client...
	mintingCLI := mintingcli.NewPluginConsensusClient(cli)
	// ...and register minting types
	types.RegisterTransactionVersion(gctypes.MinterDefinitionTxVersion, minting.MinterDefinitionTransactionController{
		MintConditionGetter: mintingCLI,
		TransactionVersion:  gctypes.MinterDefinitionTxVersion,
	})
	types.RegisterTransactionVersion(gctypes.CoinCreationTxVersion, minting.CoinCreationTransactionController{
		MintConditionGetter: mintingCLI,
		TransactionVersion:  gctypes.CoinCreationTxVersion,
	})
	types.RegisterTransactionVersion(gctypes.CoinDestructionTxVersion, minting.CoinDestructionTransactionController{
		TransactionVersion: gctypes.CoinDestructionTxVersion,
	})
}
