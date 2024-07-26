package cmd

import (
	"regexp"

	"github.com/mvrilo/go-cpf"
	"github.com/spf13/cobra"
)

func cpfCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cpf",
		Short: "Generate a valid CPF",
		Long:  `Generate a valid CPF (Brazilian Social Security Number)`,
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			code := cpf.Generate()
			cmd.Println(Success(code))
		},
	}

	validator := &cobra.Command{
		Use:   "validate [cpf]",
		Short: "Validate a CPF",
		Long:  `validate a CPF (Brazilian Social Security Number)`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			code := cleanCpfSymbols(args[0])
			if ok, err := cpf.Valid(code); err == nil && ok {
				cmd.Println(Valid(code))
			} else {
				cmd.PrintErrln(Invalid(code, err))
			}
		},
	}

	cmd.AddCommand(validator)

	return cmd
}

func cleanCpfSymbols(cpf string) string {
	rg := regexp.MustCompile(`\D`)
	return rg.ReplaceAllString(cpf, "")
}
