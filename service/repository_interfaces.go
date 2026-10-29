package service

import "dog_tracks/object"

type DogRepositoryInterface interface {
	GetDogByID(id int, includeDeleted bool) (*object.Dog, error)
	CreateDog(dog object.Dog, userId int) (int, error)
	UpdateDog(id int, dog object.Dog, userId int) (object.Dog, error)
	DeleteDog(id int, userId int) error
}

type TrainSessionRepositoryInterface interface {
	GetSessionByID(id int, includeDeleted bool) (*object.Training_Session, error)
}
