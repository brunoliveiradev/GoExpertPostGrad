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
	UserDB database.UserInterface
}

func NewUserHandler(db database.UserInterface) *UserHandler {
	return &UserHandler{
		UserDB: db,
	}
}

// CreateUser godoc
// @Summary 	Create a new user
// @Description Create a new user
// @Tags 		users
// @Accept 		json
// @Produce 	json
// @Param 		request 	body 		dto.CreateUserInput 	true 	"user request"
// @Success 	201
// @Failure 	400			{object} 	entity.ErrorResponse 	"Bad request"
// @Failure 	500			{object} 	entity.ErrorResponse 	"Internal server error"
// @Router 		/users [post]
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

// GetJWT godoc
// @Summary 	Generate a JWT token for user
// @Description Generate a JWT token for a given user credentials
// @Tags 		users
// @Accept 		json
// @Produce 	json
// @Param 		request 	body 		dto.GetJWTInput 		true 	"user credentials"
// @Success 	200 		{object} 	dto.GetJWTOutput
// @Failure 	400			{object} 	entity.ErrorResponse 	"Bad request"
// @Failure 	401			{object} 	entity.ErrorResponse 	"Unauthorized"
// @Failure 	404			{object} 	entity.ErrorResponse 	"User not found"
// @Failure 	500			{object} 	entity.ErrorResponse 	"Internal server error"
// @Router 		/users/generate_token [post]
func (h *UserHandler) GetJWT(w http.ResponseWriter, r *http.Request) {
	jwt := r.Context().Value("jwt").(*jwtauth.JWTAuth)
	jwtExpireTime := r.Context().Value("jwtExpireTime").(int)

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
			http.Error(w, "User not found", http.StatusNotFound)
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

	_, tokenString, err := jwt.Encode(map[string]interface{}{
		"sub": user.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(jwtExpireTime)).Unix(),
	})
	accessToken := dto.GetJWTOutput{AccessToken: tokenString}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)
}
