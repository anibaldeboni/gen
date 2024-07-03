package cmd

import (
	"fmt"
	"regexp"

	"github.com/mvrilo/go-cpf"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(cpfCmd)
	cpfCmd.AddCommand(cpfValidateCmd)
}

var cpfCmd = &cobra.Command{
	Use:   "cpf",
	Short: "Generate a valid CPF",
	Long:  `generate a valid CPF (Brazilian Social Security Number)`,
	Run: func(cmd *cobra.Command, args []string) {
		code := cpf.Generate()
		sendToClipboard(code)
	},
}

var cpfValidateCmd = &cobra.Command{
	Use:   "validate [cpf]",
	Short: "Validate a CPF",
	Long:  `validate a CPF (Brazilian Social Security Number)`,
	Run: func(cmd *cobra.Command, args []string) {
		code := cleanCpfSymbols(args[0])
		if ok, err := cpf.Valid(code); err == nil && ok {
			fmt.Println("🟢", bold(code), "is valid")
		} else {
			fmt.Println("🔴", fmt.Errorf("%s %w", bold(code), err))
		}
	},
}

func cleanCpfSymbols(cpf string) string {
	rg := regexp.MustCompile(`\D`)
	return rg.ReplaceAllString(cpf, "")
}
