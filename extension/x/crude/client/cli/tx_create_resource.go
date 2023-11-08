package cli

import (
	"strconv"

	"crude/x/crude/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCreateResource() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-resource [r-name] [r-category] [r-colour] [r-size] [r-quantity]",
		Short: "Broadcast message create-resource",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argRName := args[0]
			argRCategory := args[1]
			argRColour := args[2]
			argRSize := args[3]
			argRQuantity := args[4]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateResource(
				clientCtx.GetFromAddress().String(),
				argRName,
				argRCategory,
				argRColour,
				argRSize,
				argRQuantity,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
