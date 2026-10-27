package helper

import (
	"dog_tracks/object"
	"encoding/json"
	"log"
	"net/http"
)

func ErrorResponse(status int, error string, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	e := object.ErrorResponse{
		Status: status,
		Error:  error,
	}
	if err := json.NewEncoder(w).Encode(e); err != nil {
		log.Printf("unable to format error response : %v", err)
	}
}

func SuccessResponse(status int, resp any, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	r := object.SuccessResponse{
		Status:   status,
		Response: resp,
	}
	if err := json.NewEncoder(w).Encode(r); err != nil {
		log.Printf("unable to format success response %v", err)
	}
}
