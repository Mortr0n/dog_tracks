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

func (ds DogService) GetDogByID(id int) (*object.Dog, error) {
	return ds.DogRepository.GetDogByID(id)
}

func (ds DogService) CreateDog(dog object.Dog) (int, error) {
	return ds.DogRepository.CreateDog(dog)
}

func (ds DogService) UpdateDog(id int, dog object.Dog) (object.Dog, error) {
	return ds.DogRepository.UpdateDog(id, dog)
}

func (ds DogService) DeleteDog(id int) error {
	return ds.DogRepository.DeleteDog(id)
}
