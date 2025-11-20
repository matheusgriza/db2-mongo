package handlers

import (
	"encoding/json"
	"net/http"
	"task-api/internal/models"

	"github.com/google/uuid"
)

func (h Handlers) registerUserEndpoints() {
	http.HandleFunc("GET /persons", h.getAll)
	http.HandleFunc("GET /persons/{id}", h.getPerson)
	http.HandleFunc("POST /persons", h.createPerson)
}

func (h Handlers) getAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	persons, err := h.useCases.GetAllPerson(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(persons)

}

func (h Handlers) getPerson(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	urlParam := r.PathValue("id")

	id, err := uuid.Parse(urlParam)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{Reason: err.Error()})
		return
	}

	person, err := h.useCases.GetPerson(ctx, id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(person)

}

func (h Handlers) createPerson(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req models.CreatePersonRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Reason: err.Error()})
		return
	}

	id, err := h.useCases.AddPerson(ctx, req)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{Reason: err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(models.CreatePersonResponse{NewPersonId: id})
}
