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

func ccCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cc",
		Short: "Generate a random credit card number",
		Long:  `Generate a random credit card number`,
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			code := ccGenerateFuncs[ccType]()
			cmd.Println(Success(code))
		},
	}

	cmd.Flags().VarP(
		enumflag.New(&ccType, "type", ccTypes, enumflag.EnumCaseInsensitive),
		"type",
		"t",
		"set card type: "+listCCTypes(),
	)
	cmd.Flags().Lookup("type").NoOptDefVal = "visa"

	validator := &cobra.Command{
		Use:   "validate [number]",
		Short: "Validate a credit card number",
		Long:  `Validate a credit card number`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			code := args[0]
			if ccgen.Visa.ValidNumber(code) {
				cmd.Println(Valid(code))
			} else {
				cmd.PrintErrln(Invalid(code, errors.New("invalid credit card number")))
			}
		},
	}

	cmd.AddCommand(validator)

	return cmd
}
