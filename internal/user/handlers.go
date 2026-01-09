package user

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"example.com/contact/internal/contact"
)

type userBody struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type userResponseWithContacts struct {
	Firstname string            `json:"firstname"`
	Lastname  string            `json:"lastname"`
	Email     string            `json:"email"`
	Contacts  []contact.Contact `json:"contacts"`
}

type UserHandler struct {
	userRepo    *UserRepository
	contactRepo *contact.ContactRepo
}

func NewUserHandler(userRepo *UserRepository, contactRepo *contact.ContactRepo) *UserHandler {
	return &UserHandler{
		userRepo:    userRepo,
		contactRepo: contactRepo,
	}
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	pageLimit := r.URL.Query().Get("limit")
	pageNumber := r.URL.Query().Get("page")

	limit := 10
	page := 1

	if pageLimit != "" {
		if l, e := strconv.Atoi(pageLimit); e == nil {
			limit = l
		}
	}

	if pageNumber != "" {
		if p, e := strconv.Atoi(pageNumber); e == nil {
			page = p
		}
	}

	offset := limit * (page - 1)

	users, err := h.userRepo.FindAll(offset, limit)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(users)
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	pathUserId := r.PathValue("userId")

	id, err := strconv.ParseInt(pathUserId, 10, 64)

	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	user, err := h.userRepo.FindById(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) GetUserWithContacts(w http.ResponseWriter, r *http.Request) {
	pathUserId := r.PathValue("userId")

	id, err := strconv.ParseInt(pathUserId, 10, 64)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := h.userRepo.FindById(id)

	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Can't find user", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		return
	}

	contacts, err := h.contactRepo.FindByUserId(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userResponse := userResponseWithContacts{
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Email:     user.Email,
		Contacts:  contacts,
	}

	json.NewEncoder(w).Encode(userResponse)
}

func (h *UserHandler) PostUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)
	defer r.Body.Close()

	ctx := r.Context()
	var body userBody

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	userData := UserModel{
		Firstname: body.Firstname,
		Lastname:  body.Lastname,
		Email:     body.Email,
		Password:  body.Password,
	}

	userId, err := h.userRepo.Save(ctx, userData)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userData.ID = userId

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(userData)
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ctx := r.Context()

	pathUserId := r.PathValue("userId")

	id, err := strconv.ParseInt(pathUserId, 10, 64)

	if err != nil {
		http.Error(w, "Invalid user Id", http.StatusBadRequest)
		return
	}

	if err := h.userRepo.DeleteById(ctx, id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]any{
		"message": "Deleted successfully.",
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
