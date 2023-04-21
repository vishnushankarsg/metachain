package snapshot

import (
	"math"

	"github.com/vishnushankarsg/metachain/command"
	"github.com/vishnushankarsg/metachain/command/helper"
	"github.com/spf13/cobra"
)

func GetCommand() *cobra.Command {
	ibftSnapshotCmd := &cobra.Command{
		Use:   "snapshot",
		Short: "Returns the IBFT snapshot at the latest block number, unless a block number is specified",
		Run:   runCommand,
	}

	setFlags(ibftSnapshotCmd)

	return ibftSnapshotCmd
}

func setFlags(cmd *cobra.Command) {
	cmd.Flags().Uint64Var(
		&params.blockNumber,
		numberFlag,
		math.MaxUint64,
		"the block height (number) for the snapshot",
	)
}

func runCommand(cmd *cobra.Command, _ []string) {
	outputter := command.InitializeOutputter(cmd)
	defer outputter.WriteOutput()

	if err := params.initSnapshot(helper.GetGRPCAddress(cmd)); err != nil {
		outputter.SetError(err)

		return
	}

	result, err := newIBFTSnapshotResult(params.snapshot)
	if err != nil {
		outputter.SetError(err)

		return
	}

	outputter.SetCommandResult(result)
}
