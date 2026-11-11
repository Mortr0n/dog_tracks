package handler

import (
	"bytes"
	"dog_tracks/helper"
	"dog_tracks/object"
	"dog_tracks/repository"
	"dog_tracks/service"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strconv"
	"testing"
	"time"
)

var userId = 1
var dogTime = time.Date(2026, time.October, 28, 11, 0, 0, 0, time.UTC)
var deleteTime = time.Date(2026, time.October, 28, 3, 0, 0, 0, time.UTC)
var trainTime = time.Date(2026, time.October, 28, 1, 0, 0, 0, time.UTC)

func TestGetTrainingSessionByID(t *testing.T) {
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

	h.GetTrainingSessionByID(recorder, request)

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

func TestGetAllTrainingSessionsByDogID(t *testing.T) {
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

	h.GetAllTrainingSessionsByDogID(recorder, request)

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

func TestCreateTrainingSession(t *testing.T) {
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
	h.CreateTrainingSession(recorder, request)
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
	expectedSession.CreatedBy = 1
	expectedSession.UpdatedBy = 1

	es := &expectedSession

	if !reflect.DeepEqual(session, es) {
		helper.ObjectCheckFormatter(t, session, es)
	}
}

func TestCreateTrainingSessionMalformedBody(t *testing.T) {
	r := repository.NewTrainingSessionRepositoryFake()
	s := service.NewTrainingSessionService(r)
	h := NewTrainingSessionHandler(s)

	recorder := httptest.NewRecorder()
	invalidObject := `{"ID": `

	body := bytes.NewReader([]byte(invalidObject))
	request := httptest.NewRequest("POST", "/api/trainingsession/", body)

	h.CreateTrainingSession(recorder, request)

	var response object.ErrorResponse
	if err := json.NewDecoder(recorder.Body).Decode(&response); err != nil {
		t.Fatalf("unable to decode error response for training session creation malformed body err: %v", err)
	}

	expectedStatus := http.StatusBadRequest
	if recorder.Code != expectedStatus {
		t.Errorf("invalid http status code expected %v, got %v", expectedStatus, recorder.Code)
	}

	if response.Status != expectedStatus {
		t.Errorf("invalid error http status code expected %v, got %v", expectedStatus, response.Status)
	}

	expectedError := "err decoding body from request err: unexpected EOF"
	if response.Error != expectedError {
		t.Errorf("invalid error response, expected: %v, got: %v", expectedError, response.Error)
	}
}

func TestCreatTrainingSessionRepositoryError(t *testing.T) {
	r := repository.NewTrainingSessionRepositoryFake()
	testRun := false
	r.CreateTrainingSessionFunc = func(session object.Training_Session) (int, error) {
		testRun = true
		return -1, errors.New("unable to store training session due to no space available")
	}

	s := service.NewTrainingSessionService(r)
	h := NewTrainingSessionHandler(s)

	testSession := object.Training_Session{
		ID:              1,
		DogID:           1,
		TrainingTypeID:  1,
		StartedAt:       trainTime,
		DurationSeconds: 1500,
		Notes:           "dog recalled test",
		CreatedAt:       dogTime,
		CreatedBy:       userId,
		UpdatedAt:       dogTime,
		UpdatedBy:       userId,
		DeletedAt:       nil,
		DeletedBy:       nil,
	}

	recorder := httptest.NewRecorder()

	testSessionBytes, err := json.Marshal(testSession)
	if err != nil {
		t.Fatalf("unable to marshal reader bytes %v", err)
	}
	bodyBytesReader := bytes.NewReader(testSessionBytes)

	request := httptest.NewRequest("POST", "/api/trainingsession/", bodyBytesReader)

	h.CreateTrainingSession(recorder, request)

	if testRun != true {
		t.Fatalf("test func was not run")
	}

	var response object.ErrorResponse
	if err := json.NewDecoder(recorder.Body).Decode(&response); err != nil {
		t.Fatalf("unable to decode response for training session repository error, err: %v", err)
	}

	expectedStatus := http.StatusInternalServerError
	if recorder.Code != expectedStatus {
		t.Errorf("expected http status code: %v, got %v", expectedStatus, recorder.Code)
	}

	if response.Status != expectedStatus {
		t.Errorf("expected http status code: %v, got %v", expectedStatus, response.Status)
	}

	expectedError := "unable to create training session from provided session: {ID:1 DogID:1 TrainingTypeID:1 StartedAt:2026-10-28 01:00:00 +0000 UTC DurationSeconds:1500 Notes:dog recalled test CreatedAt:2026-10-28 11:00:00 +0000 UTC CreatedBy:1 UpdatedAt:2026-10-28 11:00:00 +0000 UTC UpdatedBy:1 DeletedAt:<nil> DeletedBy:<nil>}, err: unable to store training session due to no space available"
	if response.Error != expectedError {
		t.Errorf("unexpected error response, expected: %v, got: %v", expectedError, response.Error)
	}
}

func TestUpdateTrainingSession(t *testing.T) {
	r := repository.NewTrainingSessionRepositoryFake()
	s := service.NewTrainingSessionService(r)
	h := NewTrainingSessionHandler(s)

	var newTime = time.Date(2026, time.October, 30, 11, 01, 0, 0, time.UTC)

	incorrectID := -42

	expectedSession := object.Training_Session{
		ID:              2,
		DogID:           2,
		TrainingTypeID:  2,
		StartedAt:       newTime,
		DurationSeconds: 5555,
		Notes:           "dog loose leash test",
		UpdatedBy:       incorrectID,
		CreatedBy:       1,
		CreatedAt:       time.Date(2026, time.October, 28, 11, 0, 0, 0, time.UTC),
		DeletedAt:       nil,
		DeletedBy:       nil,
	}

	recorder := httptest.NewRecorder()

	dogBytes, err := json.Marshal(expectedSession)
	if err != nil {
		t.Fatalf("unable to marshal dog bytes for test err: %v", err)
	}

	dogBytesReader := bytes.NewReader(dogBytes)

	requestUrl := fmt.Sprintf("/api/trainingsession/%v", expectedSession.ID)
	request := httptest.NewRequest("PATCH", requestUrl, dogBytesReader)
	request.SetPathValue("id", "2")

	before := time.Now()
	h.UpdateTrainingSession(recorder, request)
	after := time.Now()

	type resp struct {
		Status   int
		Response object.Training_Session
	}
	var response resp

	if err := json.NewDecoder(recorder.Body).Decode(&response); err != nil {
		t.Fatalf("unable to decode response from update err: %v", err)
	}

	expectedStatus := http.StatusOK
	if recorder.Code != expectedStatus {
		t.Errorf("invalid http status, expected: %v got: %v", expectedStatus, recorder.Code)
	}

	if response.Status != expectedStatus {
		t.Errorf("invalide http status, expected: %v got: %v", expectedStatus, response.Status)
	}

	us, err := s.GetTrainingSessionByID(expectedSession.ID, false)
	if err != nil {
		t.Fatalf("unable to get updated training session for ID: %v, err: %v", expectedSession.ID, err)
	}

	if us == nil {
		t.Fatalf("unable to return updated session from ID: %v", expectedSession.ID)
	}
	updatedSession := *us

	if updatedSession.UpdatedAt.Before(before) || updatedSession.UpdatedAt.After(after) {
		t.Errorf("training session updatedAt invalid, expected before: %v and after: %v, got %v", after, before, updatedSession.UpdatedAt)
	}

	expectedSession.UpdatedBy = userId //! correcting expected userID for testing
	expectedSession.UpdatedAt = updatedSession.UpdatedAt

	if !reflect.DeepEqual(expectedSession, updatedSession) {
		helper.ObjectCheckFormatter(t, expectedSession, updatedSession)
	}

	//
	if !expectedSession.UpdatedAt.Equal(response.Response.UpdatedAt) {
		t.Errorf("response updatedAt differs from expected updatedAt")
	}

	response.Response.UpdatedAt = updatedSession.UpdatedAt
	if !reflect.DeepEqual(expectedSession, response.Response) {
		helper.ObjectCheckFormatter(t, expectedSession, response.Response)
	}
}

func TestUpdateTrainingSessionMalformedBody(t *testing.T) {
	r := repository.NewTrainingSessionRepositoryFake()
	s := service.NewTrainingSessionService(r)
	h := NewTrainingSessionHandler(s)

	recorder := httptest.NewRecorder()

	badBodyString := `{"ID": `
	badBodyBytes, err := json.Marshal(badBodyString)
	if err != nil {
		t.Fatalf("unable to create bytes from body string err: %v", err)
	}

	bodyBytesReader := bytes.NewReader(badBodyBytes)

	request := httptest.NewRequest("PATCH", "/api/trainingsession/1/", bodyBytesReader)
	request.SetPathValue("id", "1")

	h.UpdateTrainingSession(recorder, request)

	var response object.ErrorResponse
	if err := json.NewDecoder(recorder.Body).Decode(&response); err != nil {
		t.Fatalf("unable to decode error response err: %v", err)
	}

	expectedStatus := http.StatusBadRequest
	if recorder.Code != expectedStatus {
		t.Errorf("invalid http status code expected: %v, got: %v", recorder.Code, expectedStatus)
	}

	if response.Status != expectedStatus {
		t.Errorf("invalid http status code expected: %v, got: %v", recorder.Code, expectedStatus)
	}

	expectedResponse := "invalid session body for update err: json: cannot unmarshal string into Go value of type object.Training_Session"
	if response.Error != expectedResponse {
		t.Errorf("invalid error response expected: %v, got: %v", expectedResponse, response.Error)
	}
}

func TestDeleteTrainingSession(t *testing.T) {
	r := repository.NewTrainingSessionRepositoryFake()
	s := service.NewTrainingSessionService(r)
	h := NewTrainingSessionHandler(s)

	idToDeleteStr := "2"

	recorder := httptest.NewRecorder()

	requestURL := fmt.Sprintf("/api/trainingsession/%v", idToDeleteStr)
	request := httptest.NewRequest("DELETE", requestURL, nil)
	request.SetPathValue("id", idToDeleteStr)

	h.DeleteTrainingSession(recorder, request)

	var response object.SuccessResponse
	if err := json.NewDecoder(recorder.Body).Decode(&response); err != nil {
		t.Fatalf("error decoding response for delete err: %v", err)
	}

	if recorder.Code != http.StatusOK {
		t.Errorf("http status expected: %v, got: %v", http.StatusOK, recorder.Code)
	}

	if response.Status != http.StatusOK {
		t.Errorf("response status expected: %v, got: %v", http.StatusOK, response.Status)
	}

	expectedResponse := "successful deletion of training session at ID: 2"
	if expectedResponse != response.Response {
		t.Errorf("Expected response: %v, got: %v", expectedResponse, response.Response)
	}

	id, _ := strconv.Atoi(idToDeleteStr)
	session, err := s.GetTrainingSessionByID(id, true)
	if err != nil {
		t.Fatalf("unable to retrieve deleted session at ID: %v, err: %v", id, err)
	}

	if session == nil {
		t.Fatalf("invalid session returned")
	}

	if session.DeletedAt == nil {
		t.Errorf("session deletedAt is nil")
	}

	if session.DeletedBy == nil {
		t.Errorf("session deletedBy is nil")
	}

	deletedByVal := *session.DeletedBy
	if deletedByVal != userId {
		t.Errorf("session deletedby set incorrectly ID: %v, expected: %v", deletedByVal, userId)
	}
}
