package cmd

import (
	"fmt"
	"github.com/LiamH97/SMC-CLI-Tool/internal"
	"github.com/spf13/cobra"
	"strconv"
)

// taxCmd represents the tax command
var taxCmd = &cobra.Command{
	Use:   "tax [earnings for the year]",
	Short: "Tells you how much income tax you owe based on your earnings. Will tell you if you are breaching the higher bracket or how close you are to it",
	Run: func(cmd *cobra.Command, args []string) {
		earnings, _ := strconv.ParseFloat(args[0], 64)
		taxOwed := internal.ReturnOwedTax(earnings)
		diffFromHigherBracket, _ := internal.ReturnAmountFromHigherBracket(earnings)
		amountOverBracket := earnings - 44000
		fmt.Printf("Your owed tax for the year is %v \n", taxOwed)
		if diffFromHigherBracket != 0 {
			fmt.Printf("You are %v away from the higher bracket of tax \n", diffFromHigherBracket)
		}
		fmt.Printf("You are already paying %v of your income at the higher rate \n", amountOverBracket)
	},
}

func init() {
	rootCmd.AddCommand(taxCmd)
}
