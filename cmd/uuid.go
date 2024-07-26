package cmd

import (
	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

func uuidCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "uuid",
		Short: "Generate a valid UUID",
		Long:  `Generate a valid UUID (Universally Unique Identifier)`,
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			code := uuid.New().String()
			cmd.Println(Success(code))
		},
	}

	validator := &cobra.Command{
		Use:   "validate [uuid]",
		Short: "Validate a UUID",
		Long:  `Validate a UUID (Universally Unique Identifier)`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			code := args[0]
			if _, err := uuid.Parse(code); err == nil {

				cmd.Println(Valid(code))
			} else {
				cmd.PrintErrln(Invalid(code, err))
			}
		},
	}

	cmd.AddCommand(validator)

	return cmd
}
