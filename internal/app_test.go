package app

import (
	"fmt"
	"testing"
)

func TestReturnStandardRateTax(t *testing.T) {
	t.Run("Getting standard tax amount", func(t *testing.T) {
		// tax should be 20% up to 44,000, ignoring any earnings higher than 44,000
		sum := returnStandardRateTax(55000.0)
		want := 44000.0 * 0.2
		returnSuccessOrFail(t, sum, want)
	})

}

func TestReturnHigherRateTax(t *testing.T) {
	t.Run("Getting higher tax amount", func(t *testing.T) {
		// tax should be 40% for any amount over 44,000
		sum := returnHigherRateTax(40000.0)
		want := 0.0
		returnSuccessOrFail(t, sum, want)
	})
}

func TestReturnOwedTax(t *testing.T) {
	t.Run("Getting Owed tax", func(t *testing.T) {
		sum := returnOwedTax(50000.0)
		want := 11200.0
		returnSuccessOrFail(t, sum, want)

	})
}

func TestReturnAmountFromHigherTax(t *testing.T) {
	t.Run("Getting amount from higher bracket", func(t *testing.T) {
		sum, _ := returnAmountFromHigherBracket(30000.0)
		want := 44000.0 - 30000.0
		returnSuccessOrFail(t, sum, want)
	})
}

func returnSuccessOrFail(t testing.TB, sum, want float64) string {
	fmt.Printf("Succusfully returned %v \n", sum)
	if sum != want {
		t.Errorf("failed: wanted %v but got %v ", want, sum)
	}
	return "Pass"
}
