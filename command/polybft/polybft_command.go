package polybft

import (
	"github.com/vishnushankarsg/metachain/command/sidechain/registration"
	"github.com/vishnushankarsg/metachain/command/sidechain/staking"
	"github.com/vishnushankarsg/metachain/command/sidechain/unstaking"
	"github.com/vishnushankarsg/metachain/command/sidechain/validators"

	"github.com/vishnushankarsg/metachain/command/sidechain/whitelist"
	"github.com/vishnushankarsg/metachain/command/sidechain/withdraw"
	"github.com/spf13/cobra"
)

func GetCommand() *cobra.Command {
	polybftCmd := &cobra.Command{
		Use:   "validator",
		Short: "Validator command",
	}

	polybftCmd.AddCommand(
		staking.GetCommand(),
		unstaking.GetCommand(),
		withdraw.GetCommand(),
		validators.GetCommand(),
		whitelist.GetCommand(),
		registration.GetCommand(),
	)

	return polybftCmd
}
