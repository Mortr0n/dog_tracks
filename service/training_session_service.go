package service

import (
	"dog_tracks/object"
)

type TrainingService struct {
	TrainingSessionRepository TrainSessionRepositoryInterface
}

func NewTrainingSessionService(tr TrainSessionRepositoryInterface) TrainingService {
	return TrainingService{
		TrainingSessionRepository: tr,
	}
}

func (ts TrainingService) GetTrainingSessionByID(id int, includeDeleted bool) (*object.Training_Session, error) {
	return ts.TrainingSessionRepository.GetTrainingSessionByID(id, includeDeleted)
}

func (ts TrainingService) GetAllTrainingSessionsByDogID(id int, includeDeleted bool) ([]object.Training_Session, error) {
	return ts.TrainingSessionRepository.GetAllTrainingSessionsByDogID(id, includeDeleted)
}

func (ts TrainingService) CreateTrainingSession(session object.Training_Session) (int, error) {
	return ts.TrainingSessionRepository.CreateTrainingSession(session)
}

func (ts TrainingService) UpdateTrainingSession(session object.Training_Session) (object.Training_Session, error) {
	return ts.TrainingSessionRepository.UpdateTrainingSession(session)
}

func (ts TrainingService) DeleteTrainingSession(id int) error {
	return ts.TrainingSessionRepository.DeleteTrainingSession(id)
}
