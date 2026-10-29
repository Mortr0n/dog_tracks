package main

import (
	"dog_tracks/handler"
	"dog_tracks/repository"
	"dog_tracks/route"
	"dog_tracks/service"
	"fmt"
	"net/http"
)

func main() {
	port := ":8000"

	dogRepository := repository.NewDogRepositoryFake()
	dogService := service.NewDogService(dogRepository)
	dogHandler := handler.NewDogHandler(dogService)
	dogRoute := route.NewDogRoute(dogHandler)

	trainingRepository := repository.NewTrainingSessionRepositoryFake()
	trainingService := service.NewTrainingSessionService(trainingRepository)

	dogRoute.RegisterRoutes()

	fmt.Println("Server starting on port " + port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Error starting server: ", err)
	}
}
