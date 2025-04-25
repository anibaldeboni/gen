package cmd

import (
	"errors"

	bc "github.com/potatowski/brazilcode/v2"
	"github.com/spf13/cobra"
)

func cnhCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cnh",
		Short: "Generate a valid CNH (National Driver's License)",
		Long:  `Generate a valid CNH (National Driver's License)`,
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			code, err := bc.CNH.Generate()
			if err != nil {
				cmd.PrintErrln(Invalid("CNH generator", err))
				return
			}
			cmd.Println(Success(code))
		},
	}

	validator := &cobra.Command{
		Use:   "validate [cnh]",
		Short: "Validate a CNH (National Driver's License)",
		Long:  `Validate a CNH (National Driver's License)`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			code := args[0]
			if isUniformDigit(code) {
				cmd.PrintErrln(Invalid(code, errors.New("CNH is formed by the same digit")))
				return
			}
			if err := bc.CNH.IsValid(code); err == nil {
				cmd.Println(Valid(code))
			} else {
				cmd.PrintErrln(Invalid("CNH validator", err))
			}
		},
	}

	cmd.AddCommand(validator)

	return cmd
}
