package main

import (
	"github.com/brunoliveiradev/courseGoExpert/APIs/config"
	"github.com/brunoliveiradev/courseGoExpert/APIs/internal/domain"
	"github.com/brunoliveiradev/courseGoExpert/APIs/internal/infra/database"
	"github.com/brunoliveiradev/courseGoExpert/APIs/internal/infra/http/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

func main() {
	// Change the current working directory to the server directory
	err := os.Chdir("/Users/brunooliveira/GolandProjects/courseGoExpert/APIs/cmd/server")
	if err != nil {
		panic(err)
	}

	cfg, err := config.LoadConfig("./.env")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&domain.Product{}, &domain.User{})

	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	userDB := database.NewUser(db)
	userHandler := handlers.NewUserHandler(userDB)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)                       // absorb panics and log the errors
	r.Use(middleware.WithValue("jwt", cfg.TokenAuth)) // set jwt in context
	r.Use(middleware.WithValue("jwtExpireTime", cfg.JWTExpireTime))
	//r.Use(LogRequest)

	// Product routes
	r.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(cfg.TokenAuth)) // jwtauth.Verifier is a middleware to verify JWT tokens
		r.Use(jwtauth.Authenticator)           // check if token is valid
		r.Post("/", productHandler.CreateProduct)
		r.Get("/", productHandler.GetAllProducts)
		r.Get("/{id}", productHandler.GetProduct)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
	})

	// User routes
	r.Post("/users", userHandler.CreateUser)
	r.Post("/users/generate_token", userHandler.GetJWT)

	http.ListenAndServe(":8000", r)
}

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
