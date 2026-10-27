package service

import "dog_tracks/object"

type DogRepositoryInterface interface {
	GetDogByID(id int) (*object.Dog, error)
	CreateDog(dog object.Dog) (int, error)
	UpdateDog(id int, dog object.Dog) (object.Dog, error)
	DeleteDog(id int) error
}
