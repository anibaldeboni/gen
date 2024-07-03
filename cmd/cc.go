package cmd

import (
	"errors"
	"slices"
	"strings"

	"github.com/nsuprun/ccgen"
	"github.com/spf13/cobra"
	"github.com/thediveo/enumflag/v2"
	"golang.org/x/exp/maps"
)

var (
	ccType ccgen.CardType = ccgen.Visa
)

func init() {
	rootCmd.AddCommand(ccCmd)
	ccCmd.Flags().VarP(
		enumflag.New(&ccType, "type", ccTypes, enumflag.EnumCaseInsensitive),
		"type",
		"t",
		"set card type: "+listCCTypes(),
	)
	ccCmd.Flags().Lookup("type").NoOptDefVal = "visa"

	ccCmd.AddCommand(validateCCCmd)
}

var ccTypes = map[ccgen.CardType][]string{
	ccgen.AmericanExpress: {"amex"},
	ccgen.DinersClub:      {"diners"},
	ccgen.DinersClubUS:    {"diners-us"},
	ccgen.Discover:        {"discover"},
	ccgen.JCB:             {"jcb"},
	ccgen.Maestro:         {"maestro"},
	ccgen.Mastercard:      {"mastercard"},
	ccgen.Solo:            {"solo"},
	ccgen.Unionpay:        {"unionpay"},
	ccgen.Visa:            {"visa"},
	ccgen.Mir:             {"mir"},
}

var ccGenerateFuncs = map[ccgen.CardType]func() string{
	ccgen.AmericanExpress: ccgen.AmericanExpress.Generate,
	ccgen.DinersClub:      ccgen.DinersClub.Generate,
	ccgen.DinersClubUS:    ccgen.DinersClubUS.Generate,
	ccgen.Discover:        ccgen.Discover.Generate,
	ccgen.JCB:             ccgen.JCB.Generate,
	ccgen.Maestro:         ccgen.Maestro.Generate,
	ccgen.Mastercard:      ccgen.Mastercard.Generate,
	ccgen.Solo:            ccgen.Solo.Generate,
	ccgen.Unionpay:        ccgen.Unionpay.Generate,
	ccgen.Visa:            ccgen.Visa.Generate,
	ccgen.Mir:             ccgen.Mir.Generate,
}

func listCCTypes() string {
	var list []string
	for _, v := range maps.Values(ccTypes) {
		list = append(list, v...)
	}
	slices.Sort(list)
	return strings.Join(list, "\n")
}

var ccCmd = &cobra.Command{
	Use:   "cc",
	Short: "Generate a random credit card number",
	Long:  `Generate a random credit card number`,
	Run: func(cmd *cobra.Command, args []string) {
		code := ccGenerateFuncs[ccType]()
		sendToClipboard(code, ccTypes[ccType][0])
	},
}

var validateCCCmd = &cobra.Command{
	Use:   "validate [number]",
	Short: "Validate a credit card number",
	Long:  `Validate a credit card number`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		code := args[0]
		if ccgen.Visa.ValidNumber(code) {
			printValid(code)
		} else {
			printInvalid(code, errors.New("invalid number"))
		}
	},
}
