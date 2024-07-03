package cmd

import (
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
	"golang.design/x/clipboard"
)

var rootCmd = &cobra.Command{
	Use:   "gen",
	Short: "Gen is a tool to generate random data.",
	Long:  `A fast and simple random data generator for common use cases built in Go.`,
}

var bold = lipgloss.NewStyle().Bold(true).Render

func printValid(code string) {
	fmt.Println("🟢", bold(code), "is valid")
}

func printInvalid(code string, err error) {
	fmt.Println("🔴", fmt.Errorf("%s %w", bold(code), err))
}

func sendToClipboard(code string) {
	clipboard.Write(clipboard.FmtText, []byte(code))
	fmt.Println("🔔", bold(code), "copied to clipboard")
}

func Execute() {
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
