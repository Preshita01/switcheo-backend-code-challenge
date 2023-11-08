package cli

import (
	"strconv"

	"crude/x/crude/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

func CmdCreateResource() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-resource [r-name] [r-category] [r-colour] [r-size] [r-quantity]",
		Short: "Create a new resource",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argRName := args[0]
			argRCategory := args[1]
			argRColour := args[2]
			argRSize := args[3]
			argRQuantity, err := cast.ToUint64E(args[4])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateResource(clientCtx.GetFromAddress().String(), argRName, argRCategory, argRColour, argRSize, argRQuantity)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateResource() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-resource [id] [r-name] [r-category] [r-colour] [r-size] [r-quantity]",
		Short: "Update a resource",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			argRName := args[1]

			argRCategory := args[2]

			argRColour := args[3]

			argRSize := args[4]

			argRQuantity, err := cast.ToUint64E(args[5])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateResource(clientCtx.GetFromAddress().String(), id, argRName, argRCategory, argRColour, argRSize, argRQuantity)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteResource() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-resource [id]",
		Short: "Delete a resource by id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteResource(clientCtx.GetFromAddress().String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
