package cmd

import (
	"fmt"
	"os"
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
	msg.WriteString("is valid")

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

func Execute() {
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

	err := clipboard.Init()
	if err != nil {
		panic(err)
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
