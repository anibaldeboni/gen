package cmd

import (
	"errors"

	bc "github.com/potatowski/brazilcode/v2"
	"github.com/spf13/cobra"
)

func renavamCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "renavam",
		Short: "Generate a valid RENAVAM number",
		Long:  `Generate a valid RENAVAM number`,
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			code, err := bc.RENAVAM.Generate()
			if err != nil {
				cmd.PrintErrln(Invalid("RENAVAM generator", err))
				return
			}
			cmd.Println(Success(code))
		},
	}

	validator := &cobra.Command{
		Use:   "validate [renavam]",
		Short: "Validate a RENAVAM number",
		Long:  `Validate a RENAVAM number`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			code := args[0]

			if isUniformDigit(code) {
				cmd.PrintErrln(Invalid(code, errors.New("RENAVAM is formed by the same digit")))
				return
			}

			if err := bc.RENAVAM.IsValid(code); err == nil {
				cmd.Println(Valid(code))
			} else {
				cmd.PrintErrln(Invalid("RENAVAM validator", err))
			}
		},
	}

	cmd.AddCommand(validator)

	return cmd
}
