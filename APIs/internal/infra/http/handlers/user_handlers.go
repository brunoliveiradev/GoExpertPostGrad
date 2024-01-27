package handlers

import (
	"encoding/json"
	"errors"
	"github.com/brunoliveiradev/courseGoExpert/APIs/internal/domain"
	"github.com/brunoliveiradev/courseGoExpert/APIs/internal/dto"
	"github.com/brunoliveiradev/courseGoExpert/APIs/internal/infra/database"
	"github.com/go-chi/jwtauth"
	"log"
	"net/http"
	"time"
)

type UserHandler struct {
	UserDB        database.UserInterface
	Jwt           *jwtauth.JWTAuth
	JwtExperiesIn int
}

func NewUserHandler(db database.UserInterface, jwt *jwtauth.JWTAuth, JwtExperiesIn int) *UserHandler {
	return &UserHandler{
		UserDB:        db,
		Jwt:           jwt,
		JwtExperiesIn: JwtExperiesIn,
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var input dto.CreateUserInput

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		log.Printf("Error decoding request body: %v", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	user, err := domain.NewUser(input.Name, input.Email, input.Password)
	if err != nil {
		log.Printf("Error creating new user: %v", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	err = h.UserDB.Create(user)
	if err != nil {
		log.Printf("Error creating user in database: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *UserHandler) GetJWT(w http.ResponseWriter, r *http.Request) {
	var input dto.GetJWTInput

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		log.Printf("Error decoding request body: %v", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	user, err := h.UserDB.FindByEmail(input.Email)
	if err != nil {
		if errors.Is(err, database.ErrUserNotFound) {
			log.Printf("User not found: %v", err)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		log.Printf("Error getting user from database: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if !user.ValidatedPassword(input.Password) {
		log.Printf("Invalid password")
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	_, tokenString, err := h.Jwt.Encode(map[string]interface{}{
		"sub": user.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(h.JwtExperiesIn)).Unix(),
	})

	accessToken := struct {
		AccessToken string `json:"tiger_token"`
	}{
		AccessToken: tokenString,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)
}
