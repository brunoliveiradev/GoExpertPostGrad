package product

type ProductUseCase struct {
	ProductRepository *Repository
}

func NewProductUseCase(productRepository *Repository) *ProductUseCase {
	return &ProductUseCase{ProductRepository: productRepository}
}

// GetProductByID return a product by its ID
// Disclaimer: should return a DTO Product instead of an entity Product
func (uc *ProductUseCase) GetProductByID(id int) (*Product, error) {
	return uc.ProductRepository.GetProductByID(id)
}
