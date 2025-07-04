package internal

import (
	"errors"
	"log"
)

func main() {

}

const standardTaxCutOff = 44000.0
const standardTaxRate = 0.2
const higherTaxRate = 0.4

func ReturnOwedTax(earnings float64) float64 {
	return ReturnStandardRateTax(earnings) + ReturnHigherRateTax(earnings)
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
		log.Println("Your earnings are over the standard cutoff")
		return 0, errors.New("earnings exceed the standard tax cutoff")
	}

	amount := standardTaxCutOff - earnings
	return amount, nil
}
