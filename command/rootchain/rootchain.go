package rootchain

import (
	"github.com/spf13/cobra"

	"github.com/vishnushankarsg/metachain/command/rootchain/initcontracts"
)

// GetCommand creates "rootchain" helper command
func GetCommand() *cobra.Command {
	rootchainCmd := &cobra.Command{
		Use:   "rootchain",
		Short: "Top level rootchain helper command.",
	}

	registerSubcommands(rootchainCmd)

	return rootchainCmd
}

func registerSubcommands(baseCmd *cobra.Command) {
	baseCmd.AddCommand(
		// init-contracts
		initcontracts.GetCommand(),
	)
}
