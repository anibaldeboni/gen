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

var printRaw bool

func init() {
	rootCmd.PersistentFlags().BoolVarP(&printRaw, "raw", "r", false, "print the raw code instead of copying to clipboard")
}

var bold = lipgloss.NewStyle().Bold(true).Render

func printValid(code string) {
	if printRaw {
		fmt.Println(true)
		return
	}

	fmt.Println("🟢", bold(code), "is valid")
}

func printInvalid(code string, err error) {
	if printRaw {
		fmt.Println(false)
		return
	}

	fmt.Println("🔴", fmt.Errorf("%s %w", bold(code), err))
}

func sendToClipboard(code string) {
	if printRaw {
		fmt.Println(code)
		return
	}
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
