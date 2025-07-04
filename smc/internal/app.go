package internal

import (
	"errors"
	"fmt"
	"log"
)

func main() {

}

const standardTaxCutOff = 44000.0
const standardTaxRate = 0.2
const higherTaxRate = 0.4
const vatCutOff = 42500
const vatRate = .135

func ReturnOwedTax(earnings float64) float64 {
	return ReturnStandardRateTax(earnings) + ReturnHigherRateTax(earnings)
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

func ReturnStandardRateTax(earnings float64) float64 {
	if earnings >= standardTaxCutOff {
		return standardTaxCutOff * standardTaxRate

	}
	return earnings * standardTaxRate
}

func ReturnHigherRateTax(earnings float64) float64 {
	earningsAtHigherRate := earnings - standardTaxCutOff
	if earningsAtHigherRate > 0 {
		return earningsAtHigherRate * higherTaxRate
	}
	return 0
}

func ReturnAmountFromHigherBracket(earnings float64) (float64, error) {
	if earnings > standardTaxCutOff {
		log.Println("Your earnings are over the standard tax cutoff")
		return 0, errors.New("earnings exceed the standard tax cutoff")
	}

	amount := standardTaxCutOff - earnings
	return amount, nil
}

func ReturnIfVatOwed(earnings float64) bool {
	return earnings > vatCutOff
}

func ReturnAmountFromVat(earnings float64) float64 {
	if !ReturnIfVatOwed(earnings) {
		return vatCutOff - earnings
	}
	log.Println("\nYour earnings exceed the VAT threshold. You must pay VAT at 13.5% on all services. \nYou pay a reduced rate of VAT as a Psychotherapist.")
	return 0
}

func ReturnVatOwed(earnings float64) float64 {
	if ReturnIfVatOwed(earnings) {
		return earnings * vatRate
	}
	return 0
}
