package contact

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

type ContactHandler struct {
	repo *ContactRepo
}

type contactBody struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	UserId      int64  `json:"user_id"`
}

func NewContactHandler(repo *ContactRepo) *ContactHandler {
	return &ContactHandler{
		repo: repo,
	}
}

func (h *ContactHandler) GetContacts(w http.ResponseWriter, r *http.Request) {
	pageLimit := r.URL.Query().Get("limit")
	pageNumber := r.URL.Query().Get("page")

	page := 1
	limit := 10

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

	contacts, err := h.repo.FindAll(offset, limit)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(contacts)
}

func (h *ContactHandler) GetContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	pathContactId := r.PathValue("contactId")

	id, err := strconv.ParseInt(pathContactId, 10, 64)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	contact, err := h.repo.FindById(id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Contact not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		return
	}

	json.NewEncoder(w).Encode(contact)
}

func (h *ContactHandler) PostContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)
	defer r.Body.Close()

	ctx := r.Context()
	var body contactBody

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	contactData := Contact{
		Name:        body.Name,
		Email:       body.Email,
		PhoneNumber: body.PhoneNumber,
		UserId:      body.UserId,
	}

	contactId, err := h.repo.Save(ctx, contactData)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	contactData.ID = contactId

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(contactData)
}

func (h *ContactHandler) DeleteContact(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	pathUserId := r.PathValue("contactId")

	id, err := strconv.ParseInt(pathUserId, 10, 64)

	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := h.repo.DeleteById(ctx, id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]any{
		"message": "Contact deleted successfully.",
	}

	json.NewEncoder(w).Encode(response)
}
