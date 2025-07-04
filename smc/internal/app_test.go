package internal

import (
	"fmt"
	"testing"
)

func TestReturnStandardRateTax(t *testing.T) {
	t.Run("Getting standard tax amount", func(t *testing.T) {
		// tax should be 20% up to 44,000, ignoring any earnings higher than 44,000
		sum := ReturnStandardRateTax(55000.0)
		want := 44000.0 * 0.2
		returnSuccessOrFail(t, sum, want)
	})

}

func TestReturnHigherRateTax(t *testing.T) {
	t.Run("Getting higher tax amount", func(t *testing.T) {
		// tax should be 40% for any amount over 44,000
		sum := ReturnHigherRateTax(40000.0)
		want := 0.0
		returnSuccessOrFail(t, sum, want)
	})
}

func TestReturnOwedTax(t *testing.T) {
	t.Run("Getting Owed tax", func(t *testing.T) {
		sum := ReturnOwedTax(50000.0)
		want := 11200.0
		returnSuccessOrFail(t, sum, want)

	})
}

func TestReturnAmountFromHigherTax(t *testing.T) {
	t.Run("Getting amount from higher bracket", func(t *testing.T) {
		sum, _ := ReturnAmountFromHigherBracket(30000.0)
		want := 44000.0 - 30000.0
		returnSuccessOrFail(t, sum, want)
	})
}

func TestReturnAmountFromVat(t *testing.T) {
	t.Run("Getting amount from VAT threshold", func(t *testing.T) {
		sum := ReturnAmountFromVat(40000)
		want := 2500.0
		returnSuccessOrFail(t, sum, want)
	})
	t.Run("Getting over Threshold", func(t *testing.T) {
		sum := ReturnAmountFromVat(50000)
		want := 0.0
		returnSuccessOrFail(t, sum, want)
	})
}

func TestReturnVatOwed(t *testing.T) {
	t.Run("Getting owed VAT amount", func(t *testing.T) {
		sum := ReturnVatOwed(50000)
		want := 50000 * 0.135
		returnSuccessOrFail(t, sum, want)
	})
}

func TestReturnVatInformation(t *testing.T) {
	t.Run("Returning VAT information", func(t *testing.T) {
		fmt.Println(ReturnVatInformation(40000))
	})
}

func returnSuccessOrFail(t testing.TB, sum, want float64) string {
	fmt.Printf("Succusfully returned %v \n", sum)
	if sum != want {
		t.Errorf("failed: wanted %v but got %v ", want, sum)
	}
	return "Pass"
}
