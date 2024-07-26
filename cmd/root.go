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

func printValid(code string, opts ...string) {
	if printRaw {
		fmt.Println(true)
		return
	}

	msg := append([]string{"🟢", bold(code)}, append(opts, "is valid")...)

	fmt.Println(strings.Join(msg, " "))
}

func printInvalid(code string, err error, opts ...string) {
	if printRaw {
		fmt.Println(false)
		return
	}

	msg := append([]string{"🔴", bold(code)}, append(opts, err.Error())...)
	fmt.Println(strings.Join(msg, " "))
}

func sendToClipboard(code string, opts ...string) {
	if printRaw {
		fmt.Println(code)
		return
	}
	clipboard.Write(clipboard.FmtText, []byte(code))
	msg := append([]string{"🔔", bold(code)}, append(opts, "copied to clipboard")...)
	fmt.Println(strings.Join(msg, " "))
}

func Execute() {
	rootCmd := &cobra.Command{
		Use:           "gen",
		Short:         "Gen is a tool to generate random data.",
		Long:          `A fast and simple random data generator for common use cases built in Go.`,
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
