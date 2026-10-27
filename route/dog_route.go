package route

import (
	"dog_tracks/handler"
	"net/http"
)

type DogRoute struct {
	DogHandler handler.DogHandler
}

func NewDogRoute(d handler.DogHandler) DogRoute {
	return DogRoute{
		DogHandler: d,
	}
}

func (dr DogRoute) RegisterRoutes() {
	http.HandleFunc("/api/hello", dr.DogHandler.HelloHandler)
	http.HandleFunc("GET /api/dog/{id}", dr.DogHandler.GetDogByIDHandler)
	http.HandleFunc("POST /api/dog", dr.DogHandler.CreateDogHandler)
	http.HandleFunc("PATCH /api/dog/{id}", dr.DogHandler.UpdateDogHandler)
	http.HandleFunc("DELETE /api/dog/{id}", dr.DogHandler.DeleteDogHandler)
}
