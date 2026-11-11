package service

import (
	"dog_tracks/helper"
	"dog_tracks/object"
	"dog_tracks/repository"
	"fmt"
	"reflect"
	"testing"
	"time"
)

var userID = 0
var dogTime = time.Date(2026, time.October, 28, 11, 0, 0, 0, time.UTC)
var deleteTime = time.Date(2026, time.October, 28, 3, 0, 0, 0, time.UTC)
var trainTime = time.Date(2026, time.October, 28, 1, 0, 0, 0, time.UTC)
var newTime = time.Date(2026, time.October, 28, 1, 0, 0, 0, time.UTC)

func TestGetTrainingSessionById(t *testing.T) {
	includeDeleted := false
	id := 1

	expectedSession := object.Training_Session{
		ID:              1,
		DogID:           1,
		TrainingTypeID:  1,
		StartedAt:       trainTime,
		DurationSeconds: 1500,
		Notes:           "dog recalled first",
		CreatedAt:       dogTime,
		CreatedBy:       userID,
		UpdatedAt:       dogTime,
		UpdatedBy:       userID,
		DeletedAt:       nil,
		DeletedBy:       nil,
	}

	r := repository.NewTrainingSessionRepositoryFake()
	s := NewTrainingSessionService(r)

	ts, err := s.GetTrainingSessionByID(id, includeDeleted)
	if err != nil {
		t.Fatalf("unable to get training session at ID: %v, err: %v", id, err)
	}

	if !reflect.DeepEqual(*ts, expectedSession) {
		helper.ObjectCheckFormatter(t, expectedSession, *ts)
	}
}

func TestGetTrainingSessionByIdNotFound(t *testing.T) {
	r := repository.NewTrainingSessionRepositoryFake()
	s := NewTrainingSessionService(r)

	includeDeleted := false
	id := 42

	ts, err := s.GetTrainingSessionByID(id, includeDeleted)
	if ts != nil {
		t.Fatalf("received unexpected session %+v at ID: %v", ts, id)
	}

	if err == nil {
		t.Fatalf("expected error for no session at ID: %v", id)
	}
}

func TestCreateTrainingSession(t *testing.T) {
	includeDeleted := false
	r := repository.NewTrainingSessionRepositoryFake()
	s := NewTrainingSessionService(r)

	newSession := object.Training_Session{
		ID:              1,
		DogID:           1,
		TrainingTypeID:  1,
		StartedAt:       trainTime,
		DurationSeconds: 1500,
		Notes:           "training session test",
	}
	// for updating the time stamp variables for comparison after checking
	// they fall within expected time range for creation.
	expectedSession := newSession
	expectedSession.CreatedBy = tempUserID
	expectedSession.UpdatedBy = tempUserID

	before := time.Now()
	newID, err := s.CreateTrainingSession(newSession)
	if err != nil {
		t.Fatalf("training service creation fail err: %v", err)
	}
	after := time.Now()

	newSession.ID = newID
	expectedSession.ID = newID

	createdSession, err := s.GetTrainingSessionByID(newID, includeDeleted)
	if err != nil {
		t.Fatalf("unable to retrieve newly created dog at ID: %v, err: %v", newID, err)
	}

	if createdSession.CreatedAt.Before(before) ||
		createdSession.CreatedAt.After(after) ||
		createdSession.UpdatedAt.Before(before) ||
		createdSession.UpdatedAt.After(after) {
		t.Errorf("creation time outside of expected range, created: %v, updated: %v", createdSession.CreatedAt, createdSession.UpdatedAt)
	}

	// setting expected session time after comparison validating time was
	// within range and before the deep equal
	expectedSession.CreatedAt = createdSession.CreatedAt
	expectedSession.UpdatedAt = createdSession.UpdatedAt

	if !reflect.DeepEqual(*createdSession, expectedSession) {
		helper.ObjectCheckFormatter(t, expectedSession, *createdSession)
	}
}

func TestUpdateTrainingSession(t *testing.T) {
	r := repository.NewTrainingSessionRepositoryFake()
	s := NewTrainingSessionService(r)

	sessionUpdate := object.Training_Session{
		ID:              3,
		DogID:           2,
		TrainingTypeID:  3,
		StartedAt:       trainTime,
		DurationSeconds: 1500,
		Notes:           "training session test",
	}

	expectedSessionUpdate := sessionUpdate
	expectedSessionUpdate.UpdatedBy = 1

	before := time.Now()
	updatedSession, err := s.UpdateTrainingSession(sessionUpdate)
	if err != nil {
		t.Fatalf("unable to update session for test with session: %v, err: %v", sessionUpdate, err)
	}
	after := time.Now()

	if updatedSession.UpdatedAt.Before(before) || updatedSession.CreatedAt.After(after) {
		t.Errorf("updated session updatedAt outside expected range updatedAt: %v", updatedSession.UpdatedAt)
	}

	expectedSessionUpdate.CreatedAt = updatedSession.CreatedAt
	expectedSessionUpdate.UpdatedAt = updatedSession.UpdatedAt

	if !reflect.DeepEqual(updatedSession, expectedSessionUpdate) {
		helper.ObjectCheckFormatter(t, expectedSessionUpdate, updatedSession)
	}
}

func TestUpdateTrainingSessionIDNotFound(t *testing.T) {
	r := repository.NewTrainingSessionRepositoryFake()
	s := NewTrainingSessionService(r)

	sessionUpdate := object.Training_Session{
		ID:              42,
		DogID:           2,
		TrainingTypeID:  3,
		StartedAt:       trainTime,
		DurationSeconds: 1500,
		Notes:           "training session test",
	}

	updatedSession, err := s.UpdateTrainingSession(sessionUpdate)
	if err == nil {
		t.Fatalf("unexpected successful session update updated: %v, sessionUpdating: %v", updatedSession, sessionUpdate)
	}

	expectedErrMsg := fmt.Sprintf("no session at ID: %v", sessionUpdate.ID)

	if err.Error() != expectedErrMsg {
		t.Fatalf("expected error: %v, got %v", expectedErrMsg, err.Error())
	}
}

func TestDeleteTrainingSession(t *testing.T) {
	r := repository.NewTrainingSessionRepositoryFake()
	s := NewTrainingSessionService(r)

	idToDelete := 1

	before := time.Now()
	err := s.DeleteTrainingSession(idToDelete)
	if err != nil {
		t.Fatalf("error deleting ID: %v, err: %v", idToDelete, err)
	}
	after := time.Now()

	deletedSession, err := s.GetTrainingSessionByID(idToDelete, true)
	if err != nil {
		t.Fatalf("unable to get deleted session at id %v, err: %v", idToDelete, err)
	}

	if deletedSession.DeletedAt.Before(before) || deletedSession.DeletedAt.After(after) {
		t.Errorf("invalid session Deletion time deletedAt: %v", deletedSession.DeletedAt)
	}

	if *deletedSession.DeletedBy != userID {
		t.Errorf("invalid userID for deletion ID: %v, deletedBy: %v", userID, deletedSession.DeletedBy)
	}
}

func TestDeleteTrainingSessionIDNotFound(t *testing.T) {
	r := repository.NewTrainingSessionRepositoryFake()
	s := NewTrainingSessionService(r)

	sessionToDeleteID := 42

	err := s.DeleteTrainingSession(sessionToDeleteID)
	if err == nil {
		t.Fatalf("expected error for deleting ID: %v", sessionToDeleteID)
	}

	expectedErr := fmt.Sprintf("no session found at ID: %v", sessionToDeleteID)
	if err.Error() != expectedErr {
		t.Fatalf("err expected: %v, got: %v", expectedErr, err.Error())
	}
}
