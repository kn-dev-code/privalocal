package api

import (
	"fmt"
	"net/http"
)

type AuthState struct {

}


func AuthHandler(auth *AuthState) http.Handler {
	mux := http.NewServeMux()
	
}