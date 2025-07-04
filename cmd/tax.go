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
	Use:   "tax",
	Short: "Tells you how much income tax you owe based on your earnings. Will tell you if you are breaching the higher bracket or how close you are to it",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		earnings, _ := strconv.ParseFloat(args[0], 64)
		taxOwed := internal.ReturnOwedTax(earnings)
		fmt.Printf("Your owed tax for the year is %v", taxOwed)
	},
}

func init() {
	rootCmd.AddCommand(taxCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// taxCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// taxCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
