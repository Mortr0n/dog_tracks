package handler

import (
	"bytes"
	"dog_tracks/helper"
	"dog_tracks/object"
	"dog_tracks/repository"
	"dog_tracks/service"
	"encoding/json"
	"fmt"
	"net/http"
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
		DurationSeconds: 1500,
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

	expectedStatus := http.StatusOK
	if recorder.Code != expectedStatus {
		t.Errorf("http status expected: %v, got %v", expectedStatus, recorder.Code)
	}

	var response struct {
		Status   int
		Response object.Training_Session
	}
	if err := json.NewDecoder(recorder.Body).Decode(&response); err != nil {
		t.Fatalf("invalid response returned for request err: %v", err)
	}

	if response.Status != expectedStatus {
		t.Errorf("http status expected: %v, got %v", expectedStatus, response.Status)
	}

	if receivedID != 1 {
		t.Errorf("expected received ID = 1 got ID: %v", receivedID)
	}

	contentType := recorder.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("expected application/json content-type, got: %v", contentType)
	}

	ts := &response.Response

	if !reflect.DeepEqual(ts, expectedSession) {
		helper.ObjectCheckFormatter(t, expectedSession, ts)
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
			DurationSeconds: 1500,
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
			DurationSeconds: 1500,
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
		deletedIncl = includeDeleted
		return expectedSessions, nil
	}

	s := service.NewTrainingSessionService(r)
	h := NewTrainingSessionHandler(s)

	idString := "1"
	deletedString := "true"
	targetString := fmt.Sprintf("/api/sessions/%v/%v", idString, deletedString)

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", targetString, nil)
	request.SetPathValue("id", idString)

	request.SetPathValue("deleted", deletedString)

	h.GetAllTrainingSessionsByDogIDHandler(recorder, request)

	expectedStatus := http.StatusOK
	if recorder.Code != expectedStatus {
		t.Errorf("http status expected: %v, got %v", expectedStatus, recorder.Code)
	}

	if receivedID != 1 {
		t.Fatalf("received ID got: %v, expected %v", receivedID, idString)
	}

	if deletedIncl != true {
		t.Fatalf("deleted inclusion got: %v, expected: %v", deletedIncl, true)
	}

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

func TestCreateTrainingSessionHandler(t *testing.T) {
	r := repository.NewTrainingSessionRepositoryFake()
	s := service.NewTrainingSessionService(r)
	h := NewTrainingSessionHandler(s)

	expectedSession := object.Training_Session{
		ID:              0,
		DogID:           1,
		TrainingTypeID:  2,
		StartedAt:       trainTime,
		DurationSeconds: 2000,
		Notes:           "Notes from training creation",
	}

	recorder := httptest.NewRecorder()

	tsBytes, err := json.Marshal(expectedSession)
	if err != nil {
		t.Fatalf("unable to marshal bytes from expected session: %+v, err: %v", expectedSession, err)
	}

	body := bytes.NewReader(tsBytes)

	request := httptest.NewRequest("POST", "/api/trainingsession", body)

	before := time.Now()
	h.CreateTrainingSessionHandler(recorder, request)
	after := time.Now()

	expectedStatus := http.StatusOK
	if recorder.Code != expectedStatus {
		t.Errorf("http status expected: %v, got %v", expectedStatus, recorder.Code)
	}

	var response struct {
		Response int
		Status   int
	}

	if err := json.NewDecoder(recorder.Body).Decode(&response); err != nil {
		t.Fatalf("unable to decode request body: %+v err: %v", &request.Body, err)
	}

	sessionID := response.Response

	session, err := s.GetTrainingSessionByID(sessionID, false)
	if err != nil {
		t.Fatalf("unable to get newly created training session at ID: %v, err: %v", sessionID, err)
	}

	if session == nil {
		t.Fatalf("invalid session gotten for created ID: %v", sessionID)
	}

	if session.CreatedAt.Before(before) ||
		session.CreatedAt.After(after) ||
		session.UpdatedAt.Before(before) ||
		session.UpdatedAt.After(after) {
		t.Errorf("invalid creation or update times should be after %v and before %v got createdAt: %v and updatedAt: %v", after, before, session.CreatedAt, session.UpdatedAt)
	}

	expectedSession.ID = 7
	expectedSession.CreatedAt = session.CreatedAt
	expectedSession.UpdatedAt = session.UpdatedAt

	es := &expectedSession

	if !reflect.DeepEqual(session, es) {
		helper.ObjectCheckFormatter(t, session, es)
	}
}
