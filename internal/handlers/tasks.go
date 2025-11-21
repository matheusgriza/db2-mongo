package handlers

import (
	"net/http"
)

func (h Handlers) registerTaskEndpoints() {
	http.HandleFunc("POST /tasks/", h.createPerson)
	http.HandleFunc("GET /tasks", h.getAll)
	http.HandleFunc("GET /tasks/{id}", h.getPerson)
	http.HandleFunc("PUT /tasks/{id}", h.createPerson)
	http.HandleFunc("DELETE /tasks/{id}", h.createPerson)

	//Manage "invite" array from a task (Add or remove it)
	http.HandleFunc("POST /tasks/{task}/invite", h.createPerson)
	http.HandleFunc("DELETE /tasks/{task}/invite", h.createPerson)
}
