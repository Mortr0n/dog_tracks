package service

type TrainingService struct {
	TrainingSessionRepository TrainSessionRepositoryInterface
}

func NewTrainingSessionService(tr TrainSessionRepositoryInterface) TrainingService {
	return TrainingService{
		TrainingSessionRepository: tr,
	}
}
