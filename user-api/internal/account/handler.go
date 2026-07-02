package account

import (
	response "Day1Utils/user-api/internal/common"
	"encoding/json"
	"net/http"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

type TransferRequest struct {
	FromID int     `json:"from_id"`
	ToID   int     `json:"to_id"`
	Amount float64 `json:"amount"`
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/accounts/transfer", h.Transfer)
}

func (h *Handler) Transfer(w http.ResponseWriter, r *http.Request) {
	var req TransferRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	if err := h.service.TransferMoney(req.FromID, req.ToID, req.Amount); err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.JSON(w, http.StatusOK, "Transfer successful")
}
