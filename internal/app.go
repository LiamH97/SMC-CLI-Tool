package app

import (
	"errors"
	"log"
)

func main() {

}

const standardTaxCutOff = 44000.0
const standardTaxRate = 0.2
const higherTaxRate = 0.4

func returnOwedTax(earnings float64) float64 {
	return returnStandardRateTax(earnings) + returnHigherRateTax(earnings)
}

func returnStandardRateTax(earnings float64) float64 {
	if earnings >= standardTaxCutOff {
		return standardTaxCutOff * standardTaxRate

	}
	return earnings * standardTaxRate
}

func returnHigherRateTax(earnings float64) float64 {
	earningsAtHigherRate := earnings - standardTaxCutOff
	if earningsAtHigherRate > 0 {
		return earningsAtHigherRate * higherTaxRate
	}
	return 0
}

func returnAmountFromHigherBracket(earnings float64) (float64, error) {
	if earnings > standardTaxCutOff {
		log.Println("Your earnings are over the standard cutoff")
		return 0, errors.New("earnings exceed the standard tax cutoff")
	}

	amount := standardTaxCutOff - earnings
	return amount, nil
}
