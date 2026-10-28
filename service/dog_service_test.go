package service

import (
	"dog_tracks/object"
	"dog_tracks/repository"
	"testing"
)

func TestGetDogByID(t *testing.T) {
	dogRepository := repository.NewDogRepositoryFake()
	dogService := NewDogService(dogRepository)

	dog, err := dogService.GetDogByID(1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if dog.ID != 1 {
		t.Errorf("Expected dog with ID 1, got %v", dog.ID)
	}
	if dog.Name != "Freyja" {
		t.Errorf("Expected dog with name 'Freyja', got %v", dog.Name)
	}
	if dog.Age != 5 {
		t.Errorf("Expected dog with age 5, got %v", dog.Age)
	}
}

func TestGetDogByIDNotFound(t *testing.T) {
	dogRepository := repository.NewDogRepositoryFake()
	dogService := NewDogService(dogRepository)

	dog, err := dogService.GetDogByID(999)
	if err == nil {
		t.Errorf("Expected error for non-existent dog, got nil")
	}
	if dog != nil {
		t.Errorf("Expected nil dog for non-existent dog, got %v", dog)
	}
}

func TestCreateDog(t *testing.T) {
	dogRepository := repository.NewDogRepositoryFake()
	dogService := NewDogService(dogRepository)

	newDog := object.Dog{
		ID:   0,
		Name: "Fenrir",
		Age:  1,
	}

	dogId, err := dogService.CreateDog(newDog)
	if err != nil {
		t.Fatalf("unable to create dog using %+v err:%v", newDog, err)
	}

	dog, err := dogRepository.GetDogByID(dogId)
	if err != nil {
		t.Fatalf("unable to get dog at ID: %v err: %v", dogId, err)
	}

	if dog.ID != dogId {
		t.Errorf("expected ID: %v, got %v", dogId, dog.ID)
	}

	if dog.Name != newDog.Name {
		t.Errorf("expected Name: %v, got %v", newDog.Name, dog.Name)
	}

	if dog.Age != newDog.Age {
		t.Errorf("expected Age: %v, got %v", newDog.Age, dog.Age)
	}
}

func TestUpdateDog(t *testing.T) {
	dogRepository := repository.NewDogRepositoryFake()
	dogService := NewDogService(dogRepository)

	dogUpdate := object.Dog{
		ID:   1,
		Name: "Fenris",
		Age:  1,
	}

	dog, err := dogService.UpdateDog(dogUpdate.ID, dogUpdate)
	if err != nil {
		t.Fatalf("unable to update dog at ID: %v with err: %v", dogUpdate.ID, err)
	}

	if dog.Name != dogUpdate.Name || dog.Age != dogUpdate.Age {
		t.Errorf("dog expected %+v, got %+v", dogUpdate, dog)
	}
}

func TestUpdateDogNotFound(t *testing.T) {
	dogRepository := repository.NewDogRepositoryFake()
	dogService := NewDogService(dogRepository)

	dogUpdate := object.Dog{
		ID:   -1,
		Name: "Fenris",
		Age:  1,
	}

	_, err := dogService.UpdateDog(dogUpdate.ID, dogUpdate)
	if err == nil {
		t.Fatalf("expected dog not found error")
	}

	if err.Error() != "dog at ID: -1 not found" {
		t.Fatalf("different error than expected %v", err.Error())
	}
}

func TestDeleteDog(t *testing.T) {
	r := repository.NewDogRepositoryFake()
	s := NewDogService(r)

	dogId := 1
	err := s.DeleteDog(dogId)
	if err != nil {
		t.Fatalf("error deleting dog at ID: %v, err: %v", dogId, err)
	}

	dog, err := s.GetDogByID(1)
	if err == nil {
		t.Fatalf("dog %+v retrieved after deletion", dog)
	}
}
