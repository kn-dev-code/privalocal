package routes

import (
	"net/http"
)


func RouteFunction(h *Handler) http.Handler {
	mux := http.NewServeMux()
}