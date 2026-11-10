package route

import (
	"dog_tracks/handler"
	"net/http"
)

type Route struct {
	DogHandler             handler.DogHandler
	TrainingSessionHandler handler.TrainingSessionHandler
}

func NewRoute(d handler.DogHandler, th handler.TrainingSessionHandler) Route {
	return Route{
		DogHandler:             d,
		TrainingSessionHandler: th,
	}
}

func (dr Route) RegisterRoutes() {
	http.HandleFunc("/api/hello", dr.DogHandler.Hello)
	http.HandleFunc("GET /api/dog/{id}/{deleted}", dr.DogHandler.GetDogByID)
	http.HandleFunc("POST /api/dog", dr.DogHandler.CreateDog)
	http.HandleFunc("PATCH /api/dog/{id}", dr.DogHandler.UpdateDog)
	http.HandleFunc("DELETE /api/dog/{id}", dr.DogHandler.DeleteDog)
	http.HandleFunc("GET /api/trainingsession/{id}/{deleted}", dr.TrainingSessionHandler.GetTrainingSessionByID)
	http.HandleFunc("GET /api/trainingsessions/{id}/{deleted}", dr.TrainingSessionHandler.GetAllTrainingSessionsByDogID)
	http.HandleFunc("POST /api/trainingsession", dr.TrainingSessionHandler.CreateTrainingSession)
	http.HandleFunc("PATCH /api/trainingsession/{id}", dr.TrainingSessionHandler.UpdateTrainingSession)
	http.HandleFunc("DELETE /api/trainingsession/{id}", dr.TrainingSessionHandler.DeleteTrainingSession)
}
