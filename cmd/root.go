package cmd

import (
	"fmt"
	"os"
	"strings"

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

func printValid(code string, opts ...string) {
	if printRaw {
		fmt.Println(true)
		return
	}

	msg := []string{"🟢", bold(code)}
	msg = append(msg, opts...)
	msg = append(msg, "is valid")

	fmt.Println(strings.Join(msg, " "))
}

func printInvalid(code string, err error, opts ...string) {
	if printRaw {
		fmt.Println(false)
		return
	}

	msg := []string{"🔴", bold(code)}
	msg = append(msg, opts...)
	msg = append(msg, err.Error())

	fmt.Println(strings.Join(msg, " "))
}

func sendToClipboard(code string, opts ...string) {
	if printRaw {
		fmt.Println(code)
		return
	}
	clipboard.Write(clipboard.FmtText, []byte(code))
	fmt.Println("🔔", bold(code), strings.Join(opts, " "), "copied to clipboard")
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
