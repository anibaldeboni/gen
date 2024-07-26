package cmd

import (
	"github.com/0x6flab/namegenerator"
	"github.com/spf13/cobra"
)

func nameCmd() *cobra.Command {
	cmd := &cobra.Command{
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

	cmd.Flags().BoolP("male", "m", true, "generate a male name")
	cmd.Flags().BoolP("female", "f", false, "generate a female name")
	cmd.Flags().BoolP("non-binary", "n", false, "generate a non-binary name")
	cmd.MarkFlagsMutuallyExclusive("male", "female", "non-binary")

	return cmd
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
