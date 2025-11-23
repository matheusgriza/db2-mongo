package handlers

import (
	"encoding/json"
	"net/http"
	"task-api/internal/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h Handlers) registerTaskEndpoints() {
	http.HandleFunc("POST /tasks", h.createTask)
	http.HandleFunc("GET /tasks/{id}", h.getTaskById)
	http.HandleFunc("GET /tasks", h.getTask)
	http.HandleFunc("PUT /tasks/{id}", h.updateTask)
	http.HandleFunc("DELETE /tasks/{id}", h.deleteTask)

	//Manage "invite" array from a task (Add or remove it)
	http.HandleFunc("POST /tasks/{id}/invite", h.addInvited)
	http.HandleFunc("DELETE /tasks/{id}/invite", h.removeInvited)
}

func (h Handlers) createTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req models.CreateTaskRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Reason: err.Error()})
		return
	}

	id, err := h.useCases.AddTask(ctx, req)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{Reason: err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(models.CreateTaskResponse{Id: id})
}

func (h Handlers) getTaskById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	urlParam := r.PathValue("id")

	id, err := primitive.ObjectIDFromHex(urlParam)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{Reason: err.Error()})
		return
	}

	task, err := h.useCases.GetTask(ctx, id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)

}

func (h Handlers) getTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	tasks, err := h.useCases.GetAllTask(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tasks)

}

func (h Handlers) deleteTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	urlParam := r.PathValue("id")

	id, err := primitive.ObjectIDFromHex(urlParam)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{Reason: err.Error()})
		return
	}

	task, err := h.useCases.DeleteTask(ctx, id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)

}

func (h Handlers) updateTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req models.UpdateTaskRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Reason: err.Error()})
		return
	}

	urlParam := r.PathValue("id")

	taskId, parseErr := primitive.ObjectIDFromHex(urlParam)

	if parseErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{Reason: parseErr.Error()})
		return
	}

	task, err := h.useCases.UpdateTask(ctx, taskId, req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)

}

func (h Handlers) addInvited(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	urlParam := r.PathValue("id")
	var req models.ManageTaskInvited

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Reason: err.Error()})
		return
	}

	id, err := primitive.ObjectIDFromHex(urlParam)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{Reason: err.Error()})
		return
	}

	task, err := h.useCases.AddInvited(ctx, id, req.Invited)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{Reason: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)

}

func (h Handlers) removeInvited(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	urlParam := r.PathValue("id")
	var req models.ManageTaskInvited

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Reason: err.Error()})
		return
	}

	id, err := primitive.ObjectIDFromHex(urlParam)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{Reason: err.Error()})
		return
	}

	task, err := h.useCases.RemoveInvited(ctx, id, req.Invited)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{Reason: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)
}
