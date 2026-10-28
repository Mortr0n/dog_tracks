package service

import "dog_tracks/object"

type DogService struct {
	DogRepository DogRepositoryInterface
}

func NewDogService(dr DogRepositoryInterface) DogService {
	return DogService{
		DogRepository: dr,
	}
}

func (ds DogService) GetDogByID(id int, includeDeleted bool) (*object.Dog, error) {
	return ds.DogRepository.GetDogByID(id, includeDeleted)
}

func (ds DogService) CreateDog(dog object.Dog, userID int) (int, error) {
	return ds.DogRepository.CreateDog(dog, userID)
}

func (ds DogService) UpdateDog(id int, dog object.Dog, userID int) (object.Dog, error) {
	return ds.DogRepository.UpdateDog(id, dog, userID)
}

func (ds DogService) DeleteDog(id int, userID int) error {
	return ds.DogRepository.DeleteDog(id, userID)
}
