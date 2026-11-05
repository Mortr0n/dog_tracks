package route

import (
	"dog_tracks/handler"
	"net/http"
)

type Route struct {
	DogHandler             handler.DogHandler
	TrainingSessionHandler handler.TrainingSessionHandler
}

func NewDogRoute(d handler.DogHandler, th handler.TrainingSessionHandler) Route {
	return Route{
		DogHandler:             d,
		TrainingSessionHandler: th,
	}
}

func (dr Route) RegisterRoutes() {
	http.HandleFunc("/api/hello", dr.DogHandler.HelloHandler)
	http.HandleFunc("GET /api/dog/{id}/deleted", dr.DogHandler.GetDogByIDHandler)
	http.HandleFunc("POST /api/dog", dr.DogHandler.CreateDogHandler)
	http.HandleFunc("PATCH /api/dog/{id}", dr.DogHandler.UpdateDogHandler)
	http.HandleFunc("DELETE /api/dog/{id}", dr.DogHandler.DeleteDogHandler)
	http.HandleFunc("GET /api/trainingsession/{id}/{deleted}", dr.TrainingSessionHandler.GetTrainingSessionByIDHandler)
	http.HandleFunc("GET /api/trainingsessions/{id}/{deleted}", dr.TrainingSessionHandler.GetAllTrainingSessionsByDogIDHandler)
}
