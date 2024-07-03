package cmd

import (
	"github.com/0x6flab/namegenerator"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(nameCmd)
	nameCmd.Flags().BoolP("male", "m", true, "generate a male name")
	nameCmd.Flags().BoolP("female", "f", false, "generate a female name")
	nameCmd.Flags().BoolP("non-binary", "n", false, "generate a non-binary name")
	nameCmd.MarkFlagsMutuallyExclusive("male", "female", "non-binary")
}

var nameCmd = &cobra.Command{
	Use:   "name",
	Short: "Generate a random name",
	Long:  `Generate a random person name`,
	Run: func(cmd *cobra.Command, args []string) {
		name := namegenerator.NewGenerator().
			WithGender(defineGenderFlags(cmd))

		code := name.Generate()
		sendToClipboard(code)
	},
}

func defineGenderFlags(cmd *cobra.Command) namegenerator.Gender {
	switch {
	case cmd.Flags().Lookup("female").Value.String() == "true":
		return namegenerator.Female
	case cmd.Flags().Lookup("non-binary").Value.String() == "true":
		return namegenerator.NonBinary
	default:
		return namegenerator.Male
	}
}
