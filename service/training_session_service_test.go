package service

import (
	"dog_tracks/helper"
	"dog_tracks/object"
	"dog_tracks/repository"
	"reflect"
	"testing"
	"time"
)

var userId = 0
var dogTime = time.Date(2026, time.October, 28, 11, 0, 0, 0, time.UTC)
var deleteTime = time.Date(2026, time.October, 28, 3, 0, 0, 0, time.UTC)
var trainTime = time.Date(2026, time.October, 28, 1, 0, 0, 0, time.UTC)

func TestGetTrainingSessionById(t *testing.T) {
	includeDeleted := false
	id := 1

	expectedSession := object.Training_Session{
		ID:              1,
		DogID:           1,
		TrainingTypeID:  1,
		StartedAt:       trainTime,
		Durationseconds: 1500,
		Notes:           "dog recalled first",
		CreatedAt:       dogTime,
		CreatedBy:       userId,
		UpdatedAt:       dogTime,
		UpdatedBy:       userId,
		DeletedAt:       &deleteTime,
		DeletedBy:       &userId,
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
