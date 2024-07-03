package cmd

import (
	"regexp"
	"strings"

	"github.com/0x6flab/namegenerator"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(emailCmd)
	emailCmd.Flags().BoolP("male", "m", true, "generate a male name")
	emailCmd.Flags().BoolP("female", "f", false, "generate a female name")
	emailCmd.Flags().BoolP("non-binary", "n", false, "generate a non-binary name")
	emailCmd.MarkFlagsMutuallyExclusive("male", "female", "non-binary")
	emailCmd.Flags().StringP("domain", "d", "example.com", "define the domain of the email")
}

var emailCmd = &cobra.Command{
	Use:   "email",
	Short: "Generate a random email",
	Long:  `Generate a random email address`,
	Run: func(cmd *cobra.Command, args []string) {
		name := namegenerator.NewGenerator().
			WithGender(defineGenderFlags(cmd)).
			WithSuffix("@" + cmd.Flag("domain").Value.String())

		code := replaceDashWithDot(strings.ToLower(name.Generate()))
		sendToClipboard(code)
	},
}

func replaceDashWithDot(input string) string {
	re := regexp.MustCompile(`-`)
	return re.ReplaceAllString(input, ".")
}
