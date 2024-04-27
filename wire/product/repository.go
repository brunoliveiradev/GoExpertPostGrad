package product

import "database/sql"

type Repository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetProductByID(id int) (*Product, error) {
	var p Product

	err := r.db.QueryRow("SELECT id, name FROM products WHERE id = $1", id).Scan(&p.ID, &p.Name)
	if err != nil {
		return &Product{
			ID:   id,
			Name: "Not found",
		}, nil
	}
	return &p, nil
}
