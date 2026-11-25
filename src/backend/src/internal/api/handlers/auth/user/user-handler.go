package api

import (
	"context"
	"encoding/json"
	"net/http"
	"privalocal/internal/store/auth"
	"strconv"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserGetter interface {
	GetUserById(ctx context.Context, id int64) (auth.User, error)
}

type Handler struct {
	User UserGetter
}



// GET ME
func (h *Handler) GetMe(res http.ResponseWriter, req http.Request) {
	ctx := req.Context()
	_ = ctx
	userId, err := strconv.ParseInt(req.PathValue("id"), 10, 64)
	if err != nil || userId <= 0 {
		http.Error(res, "Invalid user id", http.StatusBadRequest)
		return 
	}

	user, err := h.User.GetUserById(ctx, userId)
	if err != nil {
		if err.Is(err, pgx.ErrNoRows) {
			http.Error(res, "User not found", http.StatusNotFound)
			return
		}
		http.Error(res, "Failed to fetch user", http.StatusInternalServerError)
		return

		res.Header().Set("Content Type", "application/json")
		if err := json.NewEncoder(res).Encode(user); err != nil {
			http.Error(res, "Failed to encode application", http.StatusInternalServerError)
			return
		}
	}


	return 
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



