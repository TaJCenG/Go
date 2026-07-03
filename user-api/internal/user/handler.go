package user

import (
	response "Day1Utils/user-api/internal/common"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/users", h.Users)
	mux.HandleFunc("/users/", h.UserByID)
}

func (h *Handler) Users(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		users, err := h.service.GetAll()
		if err != nil {
			response.Error(w, http.StatusInternalServerError, err.Error())
			return
		}
		response.JSON(w, http.StatusOK, users)
	case http.MethodPost:
		var u User
		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
			response.Error(w, http.StatusBadRequest, "Invalid JSON")
			return
		}
		if err := h.service.Create(u); err != nil {
			response.Error(w, http.StatusBadRequest, err.Error())
			return
		}
		response.JSON(w, http.StatusCreated, u)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *Handler) UserByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/users/")
	id, _ := strconv.Atoi(idStr)

	switch r.Method {
	case http.MethodGet:
		u, err := h.service.GetByID(id)
		if err != nil {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		response.JSON(w, http.StatusOK, u)
	case http.MethodPut:
		var u User
		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
			response.Error(w, http.StatusBadRequest, "Invalid JSON")
			return
		}
		u.ID = id
		if err := h.service.Update(id, u); err != nil {
			response.Error(w, http.StatusBadRequest, err.Error())
			return
		}
		response.JSON(w, http.StatusOK, u)
	case http.MethodDelete:
		if err := h.service.Delete(id); err != nil {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		response.JSON(w, http.StatusOK, "User deleted")
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context() // request-scoped context

	idStr := strings.TrimPrefix(r.URL.Path, "/users/")
	id, _ := strconv.Atoi(idStr)

	u, err := h.service.GetUserCtx(ctx, id)
	if err != nil {
		response.Error(w, http.StatusNotFound, err.Error())
		return
	}
	response.JSON(w, http.StatusOK, u)
}
