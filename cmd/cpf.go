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
		Run: func(cmd *cobra.Command, args []string) {
			code := cpf.Generate()
			sendToClipboard(code)
		},
	}
	cmd.AddCommand(cpfValidateCmd)

	return cmd
}

var cpfValidateCmd = &cobra.Command{
	Use:   "validate [cpf]",
	Short: "Validate a CPF",
	Long:  `validate a CPF (Brazilian Social Security Number)`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		code := cleanCpfSymbols(args[0])
		if ok, err := cpf.Valid(code); err == nil && ok {
			printValid(code)
		} else {
			printInvalid(code, err)
		}
	},
}

func cleanCpfSymbols(cpf string) string {
	rg := regexp.MustCompile(`\D`)
	return rg.ReplaceAllString(cpf, "")
}
