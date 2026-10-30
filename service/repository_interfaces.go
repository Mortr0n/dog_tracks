package service

import "dog_tracks/object"

type DogRepositoryInterface interface {
	GetDogByID(id int, includeDeleted bool) (*object.Dog, error)
	CreateDog(dog object.Dog, userId int) (int, error)
	UpdateDog(id int, dog object.Dog, userId int) (object.Dog, error)
	DeleteDog(id int, userId int) error
}

type TrainSessionRepositoryInterface interface {
	GetTrainingSessionByID(id int, includeDeleted bool) (*object.Training_Session, error)
	GetAllTrainingSessionsByDogID(id int, includeDeleted bool) ([]object.Training_Session, error)
	CreateTrainingSession(session object.Training_Session) (int, error)
	UpdateTrainingSession(session object.Training_Session) (object.Training_Session, error)
	DeleteTrainingSession(id int) error
}
