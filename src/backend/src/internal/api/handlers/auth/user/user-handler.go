package api

import (
	"context"
	"encoding/json"
	"net/http"
	"privalocal/internal/store/auth"
	"strconv"
	"errors"
)

type UserGetter interface {
	GetUserById(ctx context.Context, id int64) (auth.User, error)
}

type Handler struct {
	User UserGetter
}



// GET ME
func (h *Handler) GetMe(res http.ResponseWriter, req *http.Request) {
	userId, err := strconv.ParseInt(req.PathValue("id"), 10, 64)
	if err != nil || userId <= 0 {
		http.Error(res, "Invalid user id", http.StatusBadRequest)
		return 
	}

	user, err := h.User.GetUserById(req.Context(), userId)
	if err != nil {
		if errors.Is(err, auth.ErrNotFound) {
			http.Error(res, "User not found", http.StatusNotFound)
			return
		}
		http.Error(res, "Failed to fetch user", http.StatusInternalServerError)
		return

		res.Header().Set("Content-Type", "application/json")
		json.NewEncoder(res).Encode(user)
	}
}

// POST ME
func (h *Handler) PostMe(res http.Response, req http.Request) {

}

// PATCH ME
func (h *Handler) UpdateMe(res http.Response, req http.Request) {

}

// DELETE ME
func (h *Handler) DeleteMe(res http.Response, req http.Request) {

}



