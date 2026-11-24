package api

import (
	"net/http"
	"privalocal/internal/store/auth"
	"strconv"
)

type UserGetter interface {
	GetUserById(ctx context.Context, id int64) (User,error)
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



