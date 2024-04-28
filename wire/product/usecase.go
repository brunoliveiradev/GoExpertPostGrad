package product

type ProductUseCase struct {
	RepositoryInterface ProductRepositoryInterface
}

func NewProductUseCase(productRepository ProductRepositoryInterface) *ProductUseCase {
	return &ProductUseCase{RepositoryInterface: productRepository}
}

// GetProductByID return a product by its ID
// Disclaimer: should return a DTO Product instead of an entity Product
func (uc *ProductUseCase) GetProductByID(id int) (*Product, error) {
	return uc.RepositoryInterface.GetProductByID(id)
}
