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
	"strconv"
)

type ProductHandler struct {
	ProductDB database.ProductInterface
}

func NewProductHandler(db database.ProductInterface) *ProductHandler {
	return &ProductHandler{ProductDB: db}
}

// CreateProduct godoc
// @Summary 	Create a new product
// @Description Create a new product given a name and a price
// @Tags 		products
// @Accept 		json
// @Produce 	json
// @Param 		request body 		dto.CreateProductInput 	true 	"product request"
// @Success 	201
// @Failure 	400		{object} 	entity.ErrorResponse 	"Bad request"
// @Failure 	500		{object} 	entity.ErrorResponse 	"Internal server error"
// @Router 		/products [post]
// @Security 	ApiKeyAuth
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

// GetAllProducts godoc
// @Summary 	Get all products
// @Description Get all products with pagination and sorting
// @Tags 		products
// @Accept 		json
// @Produce 	json
// @Param 		page 	query 	string 	false 	"page number"
// @Param 		limit 	query 	string 	false 	"limit per page"
// @Param 		sort 	query 	string 	false 	"sort by field"
// @Success 	200 	{object} 	[]domain.Product
// @Failure 	500		{object} 	entity.ErrorResponse 	"Internal server error"
// @Router 		/products [get]
// @Security 	ApiKeyAuth
func (h *ProductHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	inputPage := r.URL.Query().Get("page")
	inputLimit := r.URL.Query().Get("limit")
	sort := r.URL.Query().Get("sort")

	page, err := strconv.Atoi(inputPage)
	if err != nil {
		page = 0
	}
	limit, err := strconv.Atoi(inputLimit)
	if err != nil {
		limit = 0
	}

	products, err := h.ProductDB.FindAll(page, limit, sort)
	if err != nil {
		log.Printf("Error getting products from database: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(products); err != nil {
		log.Printf("Error encoding products to JSON: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

// GetProduct godoc
// @Summary 	Get a product by ID
// @Description Get a product by ID
// @Tags 		products
// @Accept 		json
// @Produce 	json
// @Param 		id 		path 		string 	true 	"product ID" Format(uuid)
// @Success 	200 	{object} 	domain.Product
// @Failure 	404		{object} 	entity.ErrorResponse 	"Not found"
// @Failure 	500		{object} 	entity.ErrorResponse 	"Internal server error"
// @Router 		/products/{id} [get]
// @Security 	ApiKeyAuth
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

// UpdateProduct godoc
// @Summary 	Update a product by ID
// @Description Update a product by ID given a name and a price
// @Tags 		products
// @Accept 		json
// @Produce 	json
// @Param 		id 	path 	string 	true 	"product ID" Format(uuid)
// @Param 		request 	body 		dto.CreateProductInput 	true 	"product request"
// @Success 	200
// @Failure 	400			{object} 	entity.ErrorResponse 	"Bad request"
// @Failure 	404			{object} 	entity.ErrorResponse 	"Not found"
// @Failure 	500			{object} 	entity.ErrorResponse 	"Internal server error"
// @Router 		/products/{id} [put]
// @Security 	ApiKeyAuth
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

	if inputProduct.Price <= 0 {
		log.Printf("Price is required to be greater than 0: %v", err)
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

// DeleteProduct godoc
// @Summary 	Delete a product by ID
// @Description Delete a product by ID
// @Tags 		products
// @Accept 		json
// @Produce 	json
// @Param 		id 	path 	string 	true 	"product ID"
// @Success 	200
// @Failure 	400		{object} 	entity.ErrorResponse 	"Bad request"
// @Failure 	404		{object} 	entity.ErrorResponse 	"Not found"
// @Failure 	500		{object} 	entity.ErrorResponse 	"Internal server error"
// @Router 		/products/{id} [delete]
// @Security 	ApiKeyAuth
func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	_, err := entity.ParseID(id)
	if err != nil {
		log.Printf("Error parsing ID: %v", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	err = h.ProductDB.Delete(id)
	if err != nil {
		if errors.Is(err, database.ErrProductNotFound) {
			log.Printf("Product not found: %v", err)
			http.Error(w, "Not found", http.StatusNotFound)
			return
		}
		log.Printf("Error deleting product in database: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
