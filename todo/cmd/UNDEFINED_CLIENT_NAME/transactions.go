package main

import (
	"github.com/threefoldtech/rivine/pkg/client"
)

// RegisterStandardTransactions registers the goldchain-specific transactions as required for the standard network.
func RegisterStandardTransactions(cli *client.CommandLineClient) {
	// TODO: register custom transaction controllers for CLI if required
}

// RegisterTestnetTransactions registers the goldchain-specific transactions as required for the test network.
func RegisterTestnetTransactions(cli *client.CommandLineClient) {
	// TODO: register custom transaction controllers for CLI if required
}

// RegisterDevnetTransactions registers the goldchain-specific transactions as required for the dev network.
func RegisterDevnetTransactions(cli *client.CommandLineClient) {
	// TODO: register custom transaction controllers for CLI if required
}
