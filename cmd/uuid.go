package cmd

import (
	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(uuidCmd)
}

var uuidCmd = &cobra.Command{
	Use:   "uuid",
	Short: "Generate a valid UUID",
	Long:  `generate a valid UUID (Universally Unique Identifier)`,
	Run: func(cmd *cobra.Command, args []string) {
		code := uuid.New().String()
		sendToClipboard(code)
	},
}
