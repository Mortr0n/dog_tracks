package repository

import (
	"dog_tracks/object"
	"errors"
	"fmt"
)

var dog = object.Dog{
	ID:   1,
	Name: "Freyja",
	Age:  5,
}

var dogMap = map[int]object.Dog{
	1: dog,
}

type DogRepositoryFake struct {
	DogMap         map[int]object.Dog
	GetDogByIDFunc func(id int) (*object.Dog, error)
	CreateDogFunc  func(dog object.Dog) (int, error)
	UpdateDogFunc  func(id int, dog object.Dog) (object.Dog, error)
	DeleteDogFunc  func(id int) error
}

func NewDogRepositoryFake() DogRepositoryFake {
	return DogRepositoryFake{
		DogMap: map[int]object.Dog{
			1: dog,
		},
	}
}

func (dr DogRepositoryFake) GetDogByID(id int) (*object.Dog, error) {

	if dr.GetDogByIDFunc != nil {
		return dr.GetDogByIDFunc(id)
	}

	dog, exists := dr.DogMap[id]
	if !exists {
		return nil, fmt.Errorf("Unable to find dog at id %v", id)
	}

	return &dog, nil
}

func (dr DogRepositoryFake) CreateDog(dog object.Dog) (int, error) {

	if dr.CreateDogFunc != nil {
		return dr.CreateDogFunc(dog)
	}

	dogId := len(dr.DogMap) + 1
	newDog := object.Dog{
		ID:   dogId,
		Name: dog.Name,
		Age:  dog.Age,
	}

	dr.DogMap[newDog.ID] = newDog
	// need to return the actual dog id
	return newDog.ID, nil
}

func (dr DogRepositoryFake) UpdateDog(id int, dog object.Dog) (object.Dog, error) {
	if dr.UpdateDogFunc != nil {
		return dr.UpdateDogFunc(id, dog)
	}

	_, exists := dr.DogMap[id]
	if !exists {
		errMsg := fmt.Sprintf("dog at ID: %v not found", id)
		return object.Dog{ID: 0, Name: "", Age: 0}, errors.New(errMsg)
	}

	mappedDog := dr.DogMap[id]
	mappedDog.Age = dog.Age
	mappedDog.Name = dog.Name

	dr.DogMap[id] = mappedDog

	return dr.DogMap[id], nil
}

func (dr DogRepositoryFake) DeleteDog(id int) error {
	if dr.DeleteDogFunc != nil {
		return dr.DeleteDogFunc(id)
	}

	delete(dr.DogMap, id)
	return nil
}
