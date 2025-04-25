package cmd

import (
	"errors"

	bc "github.com/potatowski/brazilcode/v2"
	"github.com/spf13/cobra"
)

func voterRegistrationCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "voter",
		Short: "Generate a valid Voter Registration Number",
		Long:  `Generate a valid Voter Registration Number`,
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			code, err := bc.VoterRegistration.Generate()
			if err != nil {
				cmd.PrintErrln(Invalid("Voter Registration generator", err))
				return
			}
			cmd.Println(Success(code))
		},
	}

	validator := &cobra.Command{
		Use:   "validate [voter-registration]",
		Short: "Validate a Voter Registration Number",
		Long:  `Validate a Voter Registration Number`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			code := args[0]
			if isUniformDigit(code) {
				cmd.PrintErrln(Invalid(code, errors.New("Voter Registration is formed by the same digit")))
				return
			}
			if err := bc.VoterRegistration.IsValid(code); err == nil {
				cmd.Println(Valid(code))
			} else {
				cmd.PrintErrln(Invalid("Voter Registration validator", err))
			}
		},
	}

	cmd.AddCommand(validator)

	return cmd
}
