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
		Run: func(cmd *cobra.Command, args []string) {
			code := uuid.New().String()
			sendToClipboard(code)
		},
	}
	cmd.AddCommand(validateUUIDCmd)

	return cmd
}

var validateUUIDCmd = &cobra.Command{
	Use:   "validate [uuid]",
	Short: "Validate a UUID",
	Long:  `Validate a UUID (Universally Unique Identifier)`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		code := args[0]
		if _, err := uuid.Parse(code); err == nil {
			printValid(code)
		} else {
			printInvalid(code, err)
		}
	},
}
