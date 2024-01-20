package tax

type Repository interface {
	SaveTax(float64) error
}

func CalculateTaxes(amount float64) float64 {
	if amount >= 1000 && amount < 2000 {
		return amount * 0.1
	}
	if amount >= 2000 {
		return amount * 0.25
	}
	if amount <= 0 {
		return 0
	}
	return amount * 0.05
}

func CalculateTaxesAndSave(repository Repository, amount float64) error {
	taxes := CalculateTaxesForForeigners(amount)
	return repository.SaveTax(taxes)
}

func CalculateTaxesForForeigners(amount float64) float64 {
	if amount <= 0 {
		return 0
	}
	if amount >= 1000 && amount < 2000 {
		return 15.0
	}
	if amount >= 2000 {
		return 25.0
	}
	return 5.0
}
