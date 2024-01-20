package tax

func CalculateTax(amount float64) float64 {
	if amount >= 1000 {
		return amount * 0.1
	}
	if amount <= 0 {
		return 0
	}
	return amount * 0.05
}
