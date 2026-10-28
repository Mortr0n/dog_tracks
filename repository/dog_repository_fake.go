package repository

import (
	"dog_tracks/object"
	"errors"
	"fmt"
	"time"
)

var dogTime = time.Date(2026, time.October, 28, 11, 0, 0, 0, time.UTC)
var deleteTime = time.Date(2026, time.October, 28, 12, 0, 0, 0, time.UTC)
var userId = 0

var Freyja = object.Dog{
	ID:        1,
	Name:      "Freyja",
	Age:       5,
	CreatedAt: dogTime,
	CreatedBy: userId,
	UpdatedAt: dogTime,
	UpdatedBy: userId,
	DeletedAt: nil,
	DeletedBy: nil,
}
var Fenrir = object.Dog{
	ID:        2,
	Name:      "Fenrir",
	Age:       1,
	CreatedAt: dogTime,
	CreatedBy: userId,
	UpdatedAt: dogTime,
	UpdatedBy: userId,
	DeletedAt: nil,
	DeletedBy: nil,
}

var Badger = object.Dog{
	ID:        3,
	Name:      "Badger",
	Age:       10,
	CreatedAt: dogTime,
	CreatedBy: userId,
	UpdatedAt: dogTime,
	UpdatedBy: userId,
	DeletedAt: nil,
	DeletedBy: nil,
}

var DeletedDog = object.Dog{
	ID:        4,
	Name:      "Deleter",
	Age:       62,
	CreatedAt: dogTime,
	CreatedBy: userId,
	UpdatedAt: dogTime,
	UpdatedBy: userId,
	DeletedAt: &deleteTime,
	DeletedBy: &userId,
}

var dogMap = map[int]object.Dog{
	1: Freyja,
	2: Fenrir,
	3: Badger,
	4: DeletedDog,
}

type DogRepositoryFake struct {
	DogMap         map[int]object.Dog
	DogMapCounter  int
	GetDogByIDFunc func(id int, includeDeleted bool) (*object.Dog, error)
	CreateDogFunc  func(dog object.Dog) (int, error)
	UpdateDogFunc  func(id int, dog object.Dog) (object.Dog, error)
	DeleteDogFunc  func(id int) error
}

func NewDogRepositoryFake() DogRepositoryFake {
	return DogRepositoryFake{
		DogMap: map[int]object.Dog{
			1: Freyja,
			2: Fenrir,
			3: Badger,
			4: DeletedDog,
		},
		DogMapCounter: 4,
	}
}

func (dr DogRepositoryFake) GetDogByID(id int, includeDeleted bool) (*object.Dog, error) {

	if dr.GetDogByIDFunc != nil {
		return dr.GetDogByIDFunc(id, includeDeleted)
	}

	dog, exists := dr.DogMap[id]
	if !exists {
		return nil, fmt.Errorf("Unable to find dog at id %v", id)
	}

	returnDog := dog.DeletedAt == nil

	if dog.DeletedAt != nil && includeDeleted {
		returnDog = true
	}

	if !returnDog {
		return nil, fmt.Errorf("no undeleted dogs at id")
	}

	return &dog, nil
}

func (dr DogRepositoryFake) CreateDog(dog object.Dog, userId int) (int, error) {

	if dr.CreateDogFunc != nil {
		return dr.CreateDogFunc(dog)
	}

	dogId := dr.DogMapCounter + 1
	newDog := object.Dog{
		ID:        dogId,
		Name:      dog.Name,
		Age:       dog.Age,
		CreatedAt: time.Now(),
		CreatedBy: userId,
		UpdatedAt: time.Now(),
		UpdatedBy: userId,
	}
	dr.DogMap[newDog.ID] = newDog
	// need to return the actual dog id
	return newDog.ID, nil
}

func (dr DogRepositoryFake) UpdateDog(id int, dog object.Dog, userId int) (object.Dog, error) {
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
	mappedDog.UpdatedAt = time.Now()
	mappedDog.UpdatedBy = userId

	dr.DogMap[id] = mappedDog

	return dr.DogMap[id], nil
}

func (dr DogRepositoryFake) DeleteDog(id int, userId int) error {
	if dr.DeleteDogFunc != nil {
		return dr.DeleteDogFunc(id)
	}

	dog, exists := dr.DogMap[id]
	if !exists {
		errMsg := fmt.Sprintf("dog at ID: %v not found", id)
		return errors.New(errMsg)
	}

	dog = dr.DogMap[id]
	now := time.Now()
	dog.DeletedAt = &now
	dog.DeletedBy = &userId
	dr.DogMap[id] = dog

	return nil
}
