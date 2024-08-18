package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"discounts/internal/services"
	"discounts/pkg/models"

	"github.com/gorilla/mux"
)

type ClientHandler struct {
	Service *services.ClientService
}

// CreateClient
func (h *ClientHandler) CreateClient(w http.ResponseWriter, r *http.Request) {
	var client models.Client
	if err := json.NewDecoder(r.Body).Decode(&client); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := h.Service.CreateClient(&client); err != nil {
		http.Error(w, "Failed to create client", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(client)
}

// GetClientWithCoupons
func (h *ClientHandler) GetClientWithCoupons(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid client ID", http.StatusBadRequest)
		return
	}

	client, err := h.Service.GetClientWithCoupons(uint(id))
	if err != nil {
		http.Error(w, "Client not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(client)
}

// UpdateClient
func (h *ClientHandler) UpdateClient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid client ID", http.StatusBadRequest)
		return
	}

	var client models.Client
	if err := json.NewDecoder(r.Body).Decode(&client); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	client.ID = uint(id)
	if err := h.Service.UpdateClient(&client); err != nil {
		http.Error(w, "Failed to update client", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(client)
}

// DeleteClient
func (h *ClientHandler) DeleteClient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid client ID", http.StatusBadRequest)
		return
	}

	if err := h.Service.DeleteClient(uint(id)); err != nil {
		http.Error(w, "Failed to delete client", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GetAllClients
func (h *ClientHandler) GetAllClients(w http.ResponseWriter, r *http.Request) {
	clients, err := h.Service.GetAllClients()
	if err != nil {
		http.Error(w, "Failed to get clients", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(clients)
}
