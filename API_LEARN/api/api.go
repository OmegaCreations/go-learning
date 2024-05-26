package api

import (
	"encoding/json"
	"net/http"
)

// Parameters required to check coin balance
type CoinBalanceParams struct {
	Username string
}

// Coin balance response
type CoinBalanceResponse struct {
	// http res code
	Code int

	// Account balance
	Balance int
}

// Error Response
type Error struct {
	// Error code
	Code int

	// Error msg
	Message string
}

func writeError (w http.ResponseWriter, message string, code int) {
	resp := Error{
		Code: code,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(resp)
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest)
	}
	InternalErrorHandler = func(w http.ResponseWriter) {
		writeError(w, "An unexpected error occures.", http.StatusInternalServerError)
	}
)

