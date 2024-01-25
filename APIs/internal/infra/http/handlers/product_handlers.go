package handlers

import (
	"encoding/json"
	"errors"
	"github.com/brunoliveiradev/courseGoExpert/APIs/internal/domain"
	"github.com/brunoliveiradev/courseGoExpert/APIs/internal/dto"
	"github.com/brunoliveiradev/courseGoExpert/APIs/internal/infra/database"
	"github.com/brunoliveiradev/courseGoExpert/APIs/pkg/entity"
	"github.com/go-chi/chi"
	"log"
	"net/http"
)

type ProductHandler struct {
	ProductDB database.ProductInterface
}

func NewProductHandler(db database.ProductInterface) *ProductHandler {
	return &ProductHandler{ProductDB: db}
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var input dto.CreateProductInput

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		log.Printf("Error decoding request body: %v", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	product, err := domain.NewProduct(input.Name, input.Price)
	if err != nil {
		log.Printf("Error creating new product: %v", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	err = h.ProductDB.Create(product)
	if err != nil {
		log.Printf("Error creating product in database: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	product, err := h.ProductDB.FindByID(id)
	if err != nil {
		if errors.Is(err, database.ErrProductNotFound) {
			log.Printf("Product not found: %v", err)
			http.Error(w, "Not found", http.StatusNotFound)
			return
		}
		log.Printf("Error getting product from database: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(product); err != nil {
		log.Printf("Error encoding product to JSON: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	var inputProduct domain.Product
	err := json.NewDecoder(r.Body).Decode(&inputProduct)
	if err != nil {
		log.Printf("Error decoding request body: %v", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	inputProduct.ID, err = entity.ParseID(id)
	if err != nil {
		log.Printf("Error parsing ID: %v", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	err = h.ProductDB.Update(&inputProduct)
	if err != nil {
		if errors.Is(err, database.ErrProductNotFound) {
			log.Printf("Product not found: %v", err)
			http.Error(w, "Not found", http.StatusNotFound)
			return
		}
		log.Printf("Error updating product in database: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
