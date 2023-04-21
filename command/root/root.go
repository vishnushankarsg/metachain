package root

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/vishnushankarsg/metachain/command/backup"
	"github.com/vishnushankarsg/metachain/command/bridge"
	"github.com/vishnushankarsg/metachain/command/genesis"
	"github.com/vishnushankarsg/metachain/command/helper"
	"github.com/vishnushankarsg/metachain/command/ibft"
	"github.com/vishnushankarsg/metachain/command/license"
	"github.com/vishnushankarsg/metachain/command/monitor"
	"github.com/vishnushankarsg/metachain/command/peers"
	"github.com/vishnushankarsg/metachain/command/polybft"
	"github.com/vishnushankarsg/metachain/command/polybftmanifest"
	"github.com/vishnushankarsg/metachain/command/polybftsecrets"
	"github.com/vishnushankarsg/metachain/command/regenesis"
	"github.com/vishnushankarsg/metachain/command/rootchain"
	"github.com/vishnushankarsg/metachain/command/secrets"
	"github.com/vishnushankarsg/metachain/command/server"
	"github.com/vishnushankarsg/metachain/command/status"
	"github.com/vishnushankarsg/metachain/command/txpool"
	"github.com/vishnushankarsg/metachain/command/version"
	"github.com/vishnushankarsg/metachain/command/whitelist"
)

type RootCommand struct {
	baseCmd *cobra.Command
}

func NewRootCommand() *RootCommand {
	rootCommand := &RootCommand{
		baseCmd: &cobra.Command{
			Short: "Metachain  is a framework for building Ethereum-compatible Blockchain networks",
		},
	}

	helper.RegisterJSONOutputFlag(rootCommand.baseCmd)

	rootCommand.registerSubCommands()

	return rootCommand
}

func (rc *RootCommand) registerSubCommands() {
	rc.baseCmd.AddCommand(
		version.GetCommand(),
		txpool.GetCommand(),
		status.GetCommand(),
		secrets.GetCommand(),
		peers.GetCommand(),
		rootchain.GetCommand(),
		monitor.GetCommand(),
		ibft.GetCommand(),
		backup.GetCommand(),
		genesis.GetCommand(),
		server.GetCommand(),
		whitelist.GetCommand(),
		license.GetCommand(),
		polybftsecrets.GetCommand(),
		polybft.GetCommand(),
		polybftmanifest.GetCommand(),
		bridge.GetCommand(),
		regenesis.GetCommand(),
	)
}

func (rc *RootCommand) Execute() {
	if err := rc.baseCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)

		os.Exit(1)
	}
}
