package cmd

import (
	"errors"

	bc "github.com/potatowski/brazilcode/v2"
	"github.com/spf13/cobra"
)

func cnpjCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cnpj",
		Short: "Generate a valid CNPJ",
		Long:  `Validate a CNPJ (Cadastro Nacional da Pessoa Jurídica)`,
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			code, err := bc.CNPJ.Generate()
			if err != nil {
				cmd.PrintErrln(Invalid("CNPJ generator", err))
				return
			}
			cmd.Println(Success(code))
		},
	}

	validator := &cobra.Command{
		Use:   "validate [cnpj]",
		Short: "Validate a CNPJ",
		Long:  `Validate a CNPJ (Cadastro Nacional da Pessoa Jurídica)`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			code := args[0]
			if isUniformDigit(code) {
				cmd.PrintErrln(Invalid(code, errors.New("CNPJ is formed by the same digit")))
				return
			}
			if err := bc.CNPJ.IsValid(code); err == nil {
				cmd.Println(Valid(code))
			} else {
				cmd.PrintErrln(Invalid(code, err))
			}
		},
	}

	cmd.AddCommand(validator)

	return cmd
}
