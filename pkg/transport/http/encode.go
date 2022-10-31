package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/mahdi-asadzadeh/go-kit-accounts/pkg/errors"
)

func EncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func EncodeErrorResponse(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func codeFrom(err error) int {
	switch err {
	case errors.NotFound404:
		return http.StatusNotFound
	case errors.BadRequest400:
	    return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
