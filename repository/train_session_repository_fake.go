package repository

import (
	"dog_tracks/object"
	"fmt"
	"time"
)

var zeroSession = object.Training_Session{
	ID:              0,
	DogId:           0,
	TrainingTypeID:  0,
	StartedAt:       time.Time{},
	Durationseconds: 0,
	Notes:           "",
	CreatedAt:       time.Time{},
	CreatedBy:       0,
	UpdatedAt:       time.Time{},
	UpdatedBy:       0,
	DeletedAt:       nil,
	DeletedBy:       nil,
}

type TrainSessionRepositoryFake struct {
	SessionMap                        map[int]object.Training_Session
	SessionMapCounter                 int
	GetTrainingSessionByIDFunc        func(id int, includeDeleted bool) (*object.Training_Session, error)
	GetAllTrainingSessionsByDogIDFunc func(id int, includeDeleted bool) ([]object.Training_Session, error)
	CreateTrainingSessionFunc         func(session object.Training_Session) (int, error)
	UpdateTrainingSessionFunc         func(session object.Training_Session) (object.Training_Session, error)
}

func NewTrainingSessionRepositoryFake() *TrainSessionRepositoryFake {
	return &TrainSessionRepositoryFake{
		SessionMap: map[int]object.Training_Session{
			1: recallTrain1,
			2: recallTrain2,
			3: looseLeashTrain1,
			4: looseLeashTrain3,
			5: looseLeashTrain2,
			6: leaveitTrain1,
		},
		SessionMapCounter: 4,
	}
}

func (ts *TrainSessionRepositoryFake) GetTrainingSessionByID(id int, includeDeleted bool) (*object.Training_Session, error) {
	if ts.GetTrainingSessionByIDFunc != nil {
		return ts.GetTrainingSessionByIDFunc(id, includeDeleted)
	}

	session, exists := ts.SessionMap[id]
	if !exists {
		return &zeroSession, fmt.Errorf("session not found at id: %v", id)
	}

	return &session, nil
}

func (ts *TrainSessionRepositoryFake) GetAllTrainingSessionsByDogID(id int, includeDeleted bool) ([]object.Training_Session, error) {
	if ts.GetAllTrainingSessionsByDogIDFunc != nil {
		return ts.GetAllTrainingSessionsByDogIDFunc(id, includeDeleted)
	}

	var sessions []object.Training_Session

	for _, session := range ts.SessionMap {
		if session.DogId == id {
			sessions = append(sessions, session)
		}
	}

	if len(sessions) < 1 {
		return sessions, fmt.Errorf("no sessions found for dogID: %v", id)
	}

	return sessions, nil
}

func (ts *TrainSessionRepositoryFake) CreateTrainingSession(session object.Training_Session) (int, error) {
	if ts.CreateTrainingSessionFunc != nil {
		return ts.CreateTrainingSessionFunc(session)
	}

	ts.SessionMapCounter = ts.SessionMapCounter + 1
	session.ID = ts.SessionMapCounter

	ts.SessionMap[session.ID] = session

	return session.ID, nil
}

func (ts *TrainSessionRepositoryFake) UpdateTrainingSession(session object.Training_Session) (object.Training_Session, error) {
	if ts.UpdateTrainingSessionFunc != nil {
		return ts.UpdateTrainingSessionFunc(session)
	}

	updateSession, exists := ts.SessionMap[session.ID]
	if !exists {
		return zeroSession, fmt.Errorf("no session at ID: %v", session.ID)
	}

	updateSession.DogId = session.DogId
	updateSession.Durationseconds = session.Durationseconds
	updateSession.Notes = session.Notes
	updateSession.StartedAt = session.StartedAt
	updateSession.TrainingTypeID = session.TrainingTypeID
	updateSession.UpdatedAt = time.Now()
	updateSession.UpdatedBy = session.UpdatedBy

	ts.SessionMap[session.ID] = updateSession
	return ts.SessionMap[session.ID], nil
}

func (ts *TrainSessionRepositoryFake) DeleteTrainingSession(id int, userId int) error {
	session, exists := ts.SessionMap[id]
	if !exists {
		return fmt.Errorf("no session found at ID: %v", id)
	}
	now := time.Now()

	session.DeletedAt = &now
	session.DeletedBy = &userId

	ts.SessionMap[id] = session

	return nil
}
