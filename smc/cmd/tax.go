/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
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
		fmt.Printf("Your owed tax for the year is %v \n", taxOwed)
	},
}

func init() {
	rootCmd.AddCommand(taxCmd)
}
