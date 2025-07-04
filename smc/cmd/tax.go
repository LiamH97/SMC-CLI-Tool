package cmd

import (
	"fmt"
	"github.com/LiamH97/SMC-CLI-Tool/internal"
	"github.com/spf13/cobra"
	"strconv"
)

const StandardTaxCutOff = 44000.0

// taxCmd represents the tax command
var taxCmd = &cobra.Command{
	Use:   "tax [earnings for the year]",
	Short: "Tells you how much income tax you owe based on your earnings. Will tell you if you are breaching the higher bracket or how close you are to it",
	Args:  cobra.ExactArgs(1), // Ensure only one argument (earnings) is provided
	Run: func(cmd *cobra.Command, args []string) {
		earnings, err := strconv.ParseFloat(args[0], 64)
		if err != nil {
			fmt.Printf("Error: Invalid earnings amount provided. Please enter a number, e.g., 'tax 15000'.\n")
			return
		}

		// Get the detailed breakdown from your internal package functions
		// (Assuming internal.ReturnOwedTax now internally calculates these components or you add new public functions)
		// For a more detailed breakdown in the CLI, it's better to expose the individual tax components
		// from your internal package if they aren't already returned by ReturnOwedTax.
		// For this example, I'll call the internal helper functions directly.

		// Income Tax Calculation Components
		grossIncomeTax := internal.CalculateGrossIncomeTax(earnings)                   // Assuming this is public
		totalTaxCredits := internal.PersonalTaxCredit + internal.EarnedIncomeTaxCredit // Access constants from internal
		incomeTaxDue := 0.0
		if grossIncomeTax > totalTaxCredits {
			incomeTaxDue = grossIncomeTax - totalTaxCredits
		}

		// USC Calculation
		uscDue := internal.CalculateUSC(earnings) // Assuming this is public

		// PRSI Calculation
		prsiDue := internal.CalculatePRSI(earnings) // Assuming this is public

		// Total Owed Tax
		totalTaxOwed := incomeTaxDue + uscDue + prsiDue

		// Determine higher bracket status
		diffFromHigherBracket, _ := internal.ReturnAmountFromHigherBracket(earnings) // This checks if BELOW StandardTaxCutOff
		amountOverBracket := 0.0
		isPayingHigherRate := false
		if earnings > internal.StandardTaxCutOff { // Access constant from internal
			amountOverBracket = earnings - internal.StandardTaxCutOff
			isPayingHigherRate = true
		}

		// --- Detailed Output to the User ---
		fmt.Printf("\n--- Estimated Annual Tax for Earnings of €%.2f ---\n", earnings)
		fmt.Println("  (Based on 2025 rates for a single, self-employed individual, net of expenses)")
		fmt.Println("--------------------------------------------------------------------")

		// 1. Income Tax Breakdown
		fmt.Println("\n**1. Income Tax:**")
		fmt.Printf("   - Gross Income Tax (before credits): €%.2f\n", grossIncomeTax)
		fmt.Printf("     (Calculated at %.0f%% on income up to €%.0f and %.0f%% on income above that)\n",
			internal.StandardTaxRate*100, internal.StandardTaxCutOff, internal.HigherTaxRate*100)
		fmt.Printf("   - Less Tax Credits: (€%.2f Personal Credit + €%.2f Earned Income Credit) = -€%.2f\n",
			internal.PersonalTaxCredit, internal.EarnedIncomeTaxCredit, totalTaxCredits)
		fmt.Printf("     (These credits reduce your income tax liability directly)\n")
		fmt.Printf("   - **Income Tax Due:** €%.2f\n", incomeTaxDue)

		// Higher Bracket Information
		if isPayingHigherRate {
			fmt.Printf("     * You are already paying tax at the higher rate (%.0f%%) on €%.2f of your income.\n",
				internal.HigherTaxRate*100, amountOverBracket)
		} else {
			fmt.Printf("     * You are currently paying tax only at the standard rate (%.0f%%).\n",
				internal.StandardTaxRate*100)
			fmt.Printf("     * You are €%.2f away from reaching the higher tax bracket of €%.0f.\n",
				diffFromHigherBracket, internal.StandardTaxCutOff)
		}

		// 2. Universal Social Charge (USC) Breakdown
		fmt.Println("\n**2. Universal Social Charge (USC):**")
		if earnings <= internal.UscexemptionThreshold {
			fmt.Printf("   - Your income of €%.2f is below the USC exemption threshold of €%.0f. No USC is due.\n",
				earnings, internal.UscexemptionThreshold)
			fmt.Printf("   - **USC Due:** €%.2f\n", uscDue)
		} else {
			fmt.Printf("   - USC is a progressive tax on gross income (above €%.0f).\n", internal.UscexemptionThreshold)
			fmt.Printf("     - %.1f%% on the first €%.0f\n", internal.Uscband1Rate*100, internal.Uscband1Threshold)
			fmt.Printf("     - %.0f%% on income from €%.0f.01 to €%.0f\n", internal.Uscband2Rate*100, internal.Uscband1Threshold, internal.Uscband2UpperBound)
			fmt.Printf("     - %.0f%% on income from €%.0f.01 to €%.0f\n", internal.Uscband3Rate*100, internal.Uscband2UpperBound, internal.Uscband3UpperBound)
			fmt.Printf("     - %.0f%% on income above €%.0f.01\n", internal.UschigherRate*100, internal.Uscband3UpperBound)
			fmt.Printf("   - **USC Due:** €%.2f\n", uscDue)
		}

		// 3. Pay Related Social Insurance (PRSI) Breakdown (Class S)
		fmt.Println("\n**3. Pay Related Social Insurance (PRSI - Class S):**")
		if earnings < internal.PrsiClassSExemptionThreshold {
			fmt.Printf("   - Your income of €%.2f is below the PRSI Class S exemption threshold of €%.0f. No PRSI is due.\n",
				earnings, internal.PrsiClassSExemptionThreshold)
			fmt.Printf("   - **PRSI Due:** €%.2f\n", prsiDue)
		} else {
			fmt.Printf("   - As a self-employed individual, you pay Class S PRSI.\n")
			fmt.Printf("   - Calculated at %.1f%% of your earnings.\n", internal.PrsiClassSRate*100)
			fmt.Printf("   - A minimum annual contribution of €%.0f applies if your calculated amount is lower.\n", internal.PrsiClassSMinimumAnnual)
			fmt.Printf("   - **PRSI Due:** €%.2f\n", prsiDue)
		}

		fmt.Println("\n--------------------------------------------------------------------")
		fmt.Printf("**Total Estimated Annual Tax (Income Tax + USC + PRSI): €%.2f**\n", totalTaxOwed)
		fmt.Println("--------------------------------------------------------------------")

		// Add a disclaimer at the end
		fmt.Println("\nDisclaimer: This is an estimation. Always verify with official Revenue documentation or a tax professional.")
	},
}

func init() {
	rootCmd.AddCommand(taxCmd)
}
