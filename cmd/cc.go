package cmd

import (
	"errors"
	"strings"

	"github.com/nsuprun/ccgen"
	"github.com/spf13/cobra"
	"github.com/thediveo/enumflag/v2"
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
	ccgen.AmericanExpress: {"amex", "AmericanExpress"},
	ccgen.DinersClub:      {"diners", "DinersClub"},
	ccgen.DinersClubUS:    {"diners-us", "DinersClubUS"},
	ccgen.Discover:        {"discover", "Discover"},
	ccgen.JCB:             {"jcb", "JCB"},
	ccgen.Maestro:         {"maestro", "Maestro"},
	ccgen.Mastercard:      {"mastercard", "Mastercard"},
	ccgen.Solo:            {"solo", "Solo"},
	ccgen.Unionpay:        {"unionpay", "UnionPay"},
	ccgen.Visa:            {"visa", "Visa"},
	ccgen.Mir:             {"mir", "Mir"},
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
	for _, v := range ccTypes {
		list = append(list, v[0])
	}
	return strings.Join(list, ", ")
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
	Run: func(cmd *cobra.Command, args []string) {
		code := args[0]
		if ccgen.Visa.ValidNumber(code) {
			printValid(code)
		} else {
			printInvalid(code, errors.New("invalid number"))
		}
	},
}
