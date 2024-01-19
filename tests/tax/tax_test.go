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

func TestCalculateTaxBatch(t *testing.T) {
	type calcTax struct {
		amount   float64
		expected float64
	}

	table := []calcTax{
		{1000.0, 100.0},
		{500.0, 25.0},
		{200.0, 10.0},
		{1200.0, 120.0},
	}

	for _, test := range table {
		tax := CalculateTax(test.amount)
		if tax != test.expected {
			t.Errorf("Expected tax of %v for amount of %v, but got %v", test.expected, test.amount, tax)
		}
	}
}

func BenchmarkCalculateTax(b *testing.B) {
	amount := 1000.0
	expected := 100.0

	for i := 0; i < b.N; i++ {
		tax := CalculateTax(amount)

		if tax != expected {
			b.Errorf("Expected tax of %v for amount of %v, but got %v", expected, amount, tax)
		}
	}
}
