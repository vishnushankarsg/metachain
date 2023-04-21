package ibft

import (
	"github.com/vishnushankarsg/metachain/command/helper"
	"github.com/vishnushankarsg/metachain/command/ibft/candidates"
	"github.com/vishnushankarsg/metachain/command/ibft/propose"
	"github.com/vishnushankarsg/metachain/command/ibft/quorum"
	"github.com/vishnushankarsg/metachain/command/ibft/snapshot"
	"github.com/vishnushankarsg/metachain/command/ibft/status"
	_switch "github.com/vishnushankarsg/metachain/command/ibft/switch"
	"github.com/spf13/cobra"
)

func GetCommand() *cobra.Command {
	ibftCmd := &cobra.Command{
		Use:   "ibft",
		Short: "Top level IBFT command for interacting with the IBFT consensus. Only accepts subcommands.",
	}

	helper.RegisterGRPCAddressFlag(ibftCmd)

	registerSubcommands(ibftCmd)

	return ibftCmd
}

func registerSubcommands(baseCmd *cobra.Command) {
	baseCmd.AddCommand(
		// ibft status
		status.GetCommand(),
		// ibft snapshot
		snapshot.GetCommand(),
		// ibft propose
		propose.GetCommand(),
		// ibft candidates
		candidates.GetCommand(),
		// ibft switch
		_switch.GetCommand(),
		// ibft quorum
		quorum.GetCommand(),
	)
}
