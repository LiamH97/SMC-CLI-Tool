package cmd

import (
	"strconv"

	"github.com/LiamH97/SMC-CLI-Tool/internal"
	"github.com/spf13/cobra"
)

var vatCmd = &cobra.Command{
	Use:   "vat [earnings]",
	Short: "Based on your earnings, details whether or not you must pay VAT and the amount owed.",
	Run: func(cmd *cobra.Command, args []string) {
		earnings, _ := strconv.ParseFloat(args[0], 64)
		internal.ReturnVatInformation(earnings)
	},
}

func init() {
	rootCmd.AddCommand(vatCmd)
}
