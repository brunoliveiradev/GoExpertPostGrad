package tax

import "testing"

// CalculateTax calculates the tax for the given amount.
func TestCalculateTax(t *testing.T) {
	amount := 1000.0
	expected := 100.0

	tax := CalculateTax(amount)

	if tax != expected {
		t.Errorf("Expected tax of %v for amount of %v, but got %v", expected, amount, tax)
	}
}
