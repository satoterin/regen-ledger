package cli

import (
	"encoding/hex"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtxb "github.com/cosmos/cosmos-sdk/x/auth/client/txbuilder"
	"github.com/spf13/cobra"
	"gitlab.com/regen-network/regen-ledger/x/agent"
	"github.com/satori/go.uuid"
)

func addrsFromHexArray(arr []string) []sdk.AccAddress {
	n := len(arr)
	res := make([]sdk.AccAddress, n)
	for i := 0; i < n; i++ {
		str := arr[i]
		bz, err := hex.DecodeString(str)
		if err != nil {
			panic(err)
		}
		res[i] = bz
	}
	return res
}

func agentsFromHexArray(arr []string) []agent.AgentId {
	n := len(arr)
	res := make([]agent.AgentId, n)
	for i := 0; i < n; i++ {
		str := arr[i]
		bz, err := hex.DecodeString(str)
		if err != nil {
			panic(err)
		}
		res[i] = bz
	}
	return res
}

func GetCmdCreateAgent(cdc *codec.Codec) *cobra.Command {
	var threshold int
	var addrs []string
	var agents []string

	cmd := &cobra.Command{
		Use:   "create-agent [OPTIONS]",
		Short: "create an agent",
		//Args:  cobra.MinimumNArgs(1),
		PreRun: func(cmd *cobra.Command, args []string) {

		},
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)

			txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			if err := cliCtx.EnsureAccountExists(); err != nil {
				return err
			}

			account, err := cliCtx.GetFromAddress()

			if err != nil {
				return err
			}

			info := agent.AgentInfo{
				AuthPolicy:        agent.MultiSig,
				MultisigThreshold: threshold,
				Addresses:         addrsFromHexArray(addrs),
				Agents:            agentsFromHexArray(agents),
			}


			msg := agent.NewMsgCreateAgent(uuid.NewV4().Bytes(), info, account)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			cliCtx.PrintResponse = true

			return utils.CompleteAndBroadcastTxCli(txBldr, cliCtx, []sdk.Msg{msg})
		},
	}

	cmd.Flags().IntVar(&threshold, "threshold", 0, "Multisig threshold")
	cmd.Flags().StringArrayVar(&addrs, "addrs", []string{}, "Address")
	cmd.Flags().StringArrayVar(&agents, "agents", []string{}, "Agents")

	return cmd
}