package cmd

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
	"golang.design/x/clipboard"
)

var (
	Version  = "dev"
	printRaw bool
	bold     = lipgloss.NewStyle().Bold(true).Render
)

func Valid(code string, opts ...string) string {
	if printRaw {
		return "true"
	}
	msg := strings.Builder{}
	msg.WriteString("🟢 " + bold(code))
	msg.WriteString(strings.Join(opts, " "))
	msg.WriteString(" is valid")

	return msg.String()
}

func Invalid(code string, err error, opts ...string) string {
	if printRaw {
		return "false"
	}
	msg := strings.Builder{}
	msg.WriteString("🔴 " + bold(code))
	msg.WriteString(strings.Join(opts, " "))
	msg.WriteString(" " + err.Error())

	return msg.String()
}

func Success(code string, opts ...string) string {
	if printRaw {
		return code
	}
	clipboard.Write(clipboard.FmtText, []byte(code))
	msg := strings.Builder{}
	msg.WriteString("🔔 " + bold(code) + " ")
	msg.WriteString(strings.Join(opts, " "))
	msg.WriteString("copied to clipboard")

	return msg.String()
}

func RootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:           "gen",
		Short:         "Gen is a tool to generate random data.",
		Long:          `A fast and simple random data generator for common use cases built in Go.`,
		Example:       "gen name",
		Version:       Version,
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	rootCmd.PersistentFlags().BoolVarP(&printRaw, "raw", "r", false, "print the raw code instead of copying to clipboard")

	rootCmd.AddCommand(ccCmd())
	rootCmd.AddCommand(cpfCmd())
	rootCmd.AddCommand(emailCmd())
	rootCmd.AddCommand(nameCmd())
	rootCmd.AddCommand(uuidCmd())
	rootCmd.AddCommand(cnpjCmd())
	rootCmd.AddCommand(voterRegistrationCmd())
	rootCmd.AddCommand(cnhCmd())
	rootCmd.AddCommand(renavamCmd())

	err := clipboard.Init()
	if err != nil {
		panic(err)
	}

	return rootCmd
}

func isUniformDigit(s string) bool {
	if len(s) == 0 {
		return false
	}
	return strings.Count(s, string(s[0])) == len(s) && s[0] >= '0' && s[0] <= '9'
}
