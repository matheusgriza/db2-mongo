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
	var persons = h.useCases.GetAllPerson()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(persons)

}

func (h Handlers) getPerson(w http.ResponseWriter, r *http.Request) {
	urlParam := r.PathValue("id")

	id, err := uuid.Parse(urlParam)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{Reason: err.Error()})
		return
	}

	var persons = h.useCases.GetAllPerson()

	for _, p := range persons {
		if p.Id == id {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(p)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(nil)

}

func (h Handlers) createPerson(w http.ResponseWriter, r *http.Request) {

	var req models.CreatePersonRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Reason: err.Error()})
		return
	}

	id, err := h.useCases.AddPerson(req)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{Reason: err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(models.CreatePersonResponse{NewPersonId: id})
}
