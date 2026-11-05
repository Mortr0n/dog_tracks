package handler

import (
	"dog_tracks/helper"
	"dog_tracks/object"
	"dog_tracks/repository"
	"dog_tracks/service"
	"encoding/json"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

var userId = 1
var dogTime = time.Date(2026, time.October, 28, 11, 0, 0, 0, time.UTC)
var deleteTime = time.Date(2026, time.October, 28, 3, 0, 0, 0, time.UTC)
var trainTime = time.Date(2026, time.October, 28, 1, 0, 0, 0, time.UTC)

func TestGetTrainingSessionByIDHandler(t *testing.T) {
	r := repository.NewTrainingSessionRepositoryFake()
	receivedID := 0

	expectedSession := &object.Training_Session{
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
		DeletedAt:       nil,
		DeletedBy:       nil,
	}

	r.GetTrainingSessionByIDFunc = func(id int, includeDeleted bool) (*object.Training_Session, error) {
		receivedID = id
		return expectedSession, nil
	}

	s := service.NewTrainingSessionService(r)
	h := NewTrainingSessionHandler(s)

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/api/trainingsession/1/false", nil)

	request.SetPathValue("id", "1")
	request.SetPathValue("deleted", "false")

	h.GetTrainingSessionByIDHandler(recorder, request)

	if receivedID != 1 {
		t.Errorf("expected received ID = 1 got ID: %v", receivedID)
	}

	contentType := recorder.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("expected application/json content-type, got: %v", contentType)
	}

	var ts object.Training_Session
	if err := json.NewDecoder(recorder.Body).Decode(&ts); err != nil {
		t.Fatalf("unable to decode request err: %v", err)
	}

	es := *expectedSession
	if !reflect.DeepEqual(ts, es) {
		helper.ObjectCheckFormatter(t, ts, es)
	}
}

func TestGetAllTrainingSessionsByDogIDHandler(t *testing.T) {
	r := repository.NewTrainingSessionRepositoryFake()
	receivedID := 0
	deletedIncl := false

	expectedSessions := []object.Training_Session{
		{
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
			DeletedAt:       nil,
			DeletedBy:       nil,
		},
		{
			ID:              2,
			DogID:           1,
			TrainingTypeID:  1,
			StartedAt:       trainTime,
			Durationseconds: 1500,
			Notes:           "dog recalled first",
			CreatedAt:       dogTime,
			CreatedBy:       userId,
			UpdatedAt:       dogTime,
			UpdatedBy:       userId,
			DeletedAt:       nil,
			DeletedBy:       nil,
		},
	}

	r.GetAllTrainingSessionsByDogIDFunc = func(id int, includeDeleted bool) ([]object.Training_Session, error) {
		receivedID = id
		deletedIncl := includeDeleted
		return expectedSessions, nil
	}

	s := service.NewTrainingSessionService(r)
	h := NewTrainingSessionHandler(s)

	idString := "1"

	recorder := httptest.NewRecorder()

	request := httptest.NewRequest("GET", "/api/sessions/1/true", nil)
	request.SetPathValue("id", idString)
	request.SetPathValue("deleted", "true")

	h.GetAllTrainingSessionsByDogIDHandler(recorder, request)

	var sessions []object.Training_Session
	if err := json.NewDecoder(recorder.Body).Decode(&sessions); err != nil {
		t.Fatalf("unable to decode response from body: %+v, err: %v", recorder.Body, err)
	}

	if len(sessions) != 2 {
		t.Errorf("invalid length of sessions length: %v", len(sessions))
	}

	if !reflect.DeepEqual(sessions[0], expectedSessions[0]) {
		helper.ObjectCheckFormatter(t, sessions[0], expectedSessions[0])
	}

	if !reflect.DeepEqual(sessions[1], expectedSessions[1]) {
		helper.ObjectCheckFormatter(t, sessions[1], expectedSessions[1])
	}
}
