package auth

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"example.com/contact/internal/jsonwebtoken"
	"example.com/contact/internal/password"
	"example.com/contact/internal/user"
	"github.com/golang-jwt/jwt/v5"
)

type AuthHandler interface {
	Login(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	userRepo *user.UserRepository
}

type loginBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewAuthHandler(userRepo *user.UserRepository) AuthHandler {
	return &handler{
		userRepo: userRepo,
	}
}

func (h *handler) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)
	defer r.Body.Close()

	var body loginBody

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid JSON", http.StatusInternalServerError)
		return
	}

	user, err := h.userRepo.FindByEmail(body.Email)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Invalid Email", http.StatusBadRequest)
		} else {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}

	passwordMatched, err := password.VerifyPassword(body.Password, user.Password)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if !passwordMatched {
		http.Error(w, "Incorrect Credential", http.StatusBadRequest)
		return
	}

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"iss":     "what-the-fack",
	}

	token, err := jsonwebtoken.CreateToken(claims)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"token": *token,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
