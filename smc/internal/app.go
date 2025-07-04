package internal

import (
	"errors"
	"fmt"
	"log"
	"math"
)

func main() {

}

const (
	// Note: minimumTaxableIncome is typically covered by tax credits,
	// but kept here if it represents a specific threshold for other purposes.
	minimumTaxableIncome = 20000.0 // Income below which no income tax is typically paid due to credits
	StandardTaxCutOff    = 44000.0 // Standard rate cut-off point for a single person
	StandardTaxRate      = 0.2     // 20%
	HigherTaxRate        = 0.4     // 40%
)

const (
	vatCutOff = 42500.0
	vatRate   = 0.135
)

const (
	PersonalTaxCredit     = 2000.0
	EarnedIncomeTaxCredit = 2000.0
)

const (
	UscexemptionThreshold = 13000.0 // Income below this pays no USC

	Uscband1Threshold = 12012.0
	Uscband1Rate      = 0.005 // 0.5%

	Uscband2LowerBound = Uscband1Threshold + 0.01 // Starts from 12012.01
	Uscband2UpperBound = 27382.0                  // 12012 + 15370
	Uscband2Rate       = 0.02                     // 2%

	Uscband3LowerBound = Uscband2UpperBound + 0.01 // Starts from 27382.01
	Uscband3UpperBound = 70044.0                   // 27382 + 42662
	Uscband3Rate       = 0.03                      // 3%

	UschigherRate = 0.08 // 8% (applies to income over 70044.01)
)

const (
	PrsiClassSRate               = 0.041  // 4.1% of reckonable income (effective from Oct 1, 2024)
	PrsiClassSMinimumAnnual      = 650.0  // Minimum annual contribution for Class S (effective from Oct 1, 2024)
	PrsiClassSExemptionThreshold = 5000.0 // If income is below this, no PRSI is due
)

func ReturnOwedTax(earnings float64) float64 {
	grossIncomeTax := CalculateGrossIncomeTax(earnings)
	totalTaxCredits := PersonalTaxCredit + EarnedIncomeTaxCredit
	incomeTaxDue := math.Max(0, grossIncomeTax-totalTaxCredits) // Income tax cannot be negative

	uscDue := CalculateUSC(earnings)

	prsiDue := CalculatePRSI(earnings)

	return incomeTaxDue + uscDue + prsiDue
}
func ReturnVatInformation(earnings float64) string {
	vatOwed := ReturnVatOwed(earnings)
	amountFromVat := ReturnAmountFromVat(earnings)

	if ReturnIfVatOwed(earnings) {
		fmt.Printf("Based on your current earnings you owe %v in VAT. \n", vatOwed)
		return ""
	}
	fmt.Printf("You are %v from the VAT threshold. You do not need to register for VAT yet. \nKeep an eye on your income as you will need to register for VAT if your income grows by %v \n", amountFromVat, amountFromVat)
	return ""
}

func CalculateGrossIncomeTax(earnings float64) float64 {
	var tax float64

	if earnings <= StandardTaxCutOff {
		tax = earnings * StandardTaxRate
	} else {
		tax = StandardTaxCutOff * StandardTaxRate
		tax += (earnings - StandardTaxCutOff) * HigherTaxRate
	}
	return tax
}
func CalculateUSC(earnings float64) float64 {
	if earnings <= UscexemptionThreshold {
		return 0.0 // No USC if income is below the exemption threshold
	}

	var usc float64
	remainingEarnings := earnings

	// Band 1: 0.5% on first €12,012
	if remainingEarnings > Uscband1Threshold {
		usc += Uscband1Threshold * Uscband1Rate
		remainingEarnings -= Uscband1Threshold
	} else {
		usc += remainingEarnings * Uscband1Rate
		remainingEarnings = 0
	}

	// Band 2: 2% on next €15,370 (up to €27,382 total)
	if remainingEarnings > 0 && earnings > Uscband2LowerBound { // Ensure we are in this band range for total earnings
		band2Amount := math.Min(remainingEarnings, Uscband2UpperBound-Uscband1Threshold)
		usc += band2Amount * Uscband2Rate
		remainingEarnings -= band2Amount
	}

	// Band 3: 3% on next €42,662 (up to €70,044 total)
	if remainingEarnings > 0 && earnings > Uscband3LowerBound { // Ensure we are in this band range for total earnings
		band3Amount := math.Min(remainingEarnings, Uscband3UpperBound-Uscband2UpperBound)
		usc += band3Amount * Uscband3Rate
		remainingEarnings -= band3Amount
	}

	// Higher Rate: 8% on balance
	if remainingEarnings > 0 {
		usc += remainingEarnings * UschigherRate
	}
	return usc
}

func ReturnStandardRateTax(earnings float64) float64 {
	if earnings >= StandardTaxCutOff {
		return StandardTaxCutOff * StandardTaxRate

	}
	return earnings * StandardTaxRate
}

func ReturnHigherRateTax(earnings float64) float64 {
	earningsAtHigherRate := earnings - StandardTaxCutOff
	if earningsAtHigherRate > 0 {
		return earningsAtHigherRate * HigherTaxRate
	}
	return 0
}
func CalculatePRSI(earnings float64) float64 {
	if earnings < PrsiClassSExemptionThreshold {
		return 0.0 // No PRSI if income is below the exemption threshold
	}

	calculatedPRSI := earnings * PrsiClassSRate
	// Pay the higher of the calculated PRSI or the minimum annual contribution
	return math.Max(calculatedPRSI, PrsiClassSMinimumAnnual)
}

func ReturnAmountFromHigherBracket(earnings float64) (float64, error) {
	if earnings > StandardTaxCutOff {
		log.Println("Your earnings are over the standard tax cutoff")
		return 0, errors.New("earnings exceed the standard tax cutoff")
	}

	amount := StandardTaxCutOff - earnings
	return amount, nil
}

func ReturnIfVatOwed(earnings float64) bool {
	return earnings > vatCutOff
}

func ReturnAmountFromVat(earnings float64) float64 {
	if !ReturnIfVatOwed(earnings) {
		return vatCutOff - earnings
	}
	log.Println("\nYour earnings exceed the VAT threshold. You must pay vat at 13.5% on all services. \nYou pay a reduced rate of VAT as a Psychotherapist.")
	return 0
}

func ReturnVatOwed(earnings float64) float64 {
	if ReturnIfVatOwed(earnings) {
		return earnings * vatRate
	}
	return 0
}
