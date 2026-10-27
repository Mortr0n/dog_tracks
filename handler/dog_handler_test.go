package handler

import (
	"bytes"
	"dog_tracks/object"
	"dog_tracks/repository"
	"dog_tracks/service"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetDogByIDHandler(t *testing.T) {
	repository := repository.NewDogRepositoryFake()
	receivedID := 0

	repository.GetDogByIDFunc = func(id int) (*object.Dog, error) {
		receivedID = id
		return &object.Dog{ID: id, Name: "Freyja", Age: 5}, nil
	}

	service := service.NewDogService(repository)
	handler := NewDogHandler(service)

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/api/dog/1", nil)
	request.SetPathValue("id", "1")

	handler.GetDogByIDHandler(recorder, request)

	if receivedID != 1 {
		t.Errorf("expected receivedID to be 1 got %v", receivedID)
	}

	contentType := recorder.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("expected application/json got %v", contentType)
	}

	var dog object.Dog
	if err := json.NewDecoder(recorder.Body).Decode(&dog); err != nil {
		t.Errorf("unable to decode response err: %v", err)
	}

	if dog.ID != 1 || dog.Name != "Freyja" || dog.Age != 5 {
		t.Errorf("unexpected dog %+v", dog)
	}
}

func TestAddDogHandler(t *testing.T) {
	repository := repository.NewDogRepositoryFake()
	service := service.NewDogService(repository)
	handler := NewDogHandler(service)

	recorder := httptest.NewRecorder()
	newDog := object.Dog{
		ID:   0,
		Name: "Freyja",
		Age:  5,
	}

	dogBytes, err := json.Marshal(newDog)
	if err != nil {
		t.Fatalf("unable to marshal dogObject %+v", newDog)
	}

	body := bytes.NewReader(dogBytes)

	request := httptest.NewRequest("POST", "/api/dog/", body)

	handler.CreateDogHandler(recorder, request)

	// t.Logf("status: %d, body: %v", recorder.Code, recorder.Body.String())

	var response struct {
		Status   int
		Response object.Dog
	}
	var dogResponse object.Dog
	if err := json.NewDecoder(recorder.Body).Decode(&response); err != nil {
		t.Fatalf("unable to decode create response: %v", err)
	}

	dogResponse = response.Response
	// t.Logf("dogResponse: %+v, response %+v", dogResponse, response)

	createdDog, err := service.GetDogByID(dogResponse.ID)
	if err != nil {
		t.Fatalf("unable to get created dog by id %v err: %v", dogResponse.ID, err)
	}

	if recorder.Code != http.StatusOK {
		t.Errorf("expected status: %v, got %v", http.StatusOK, recorder.Code)
	}

	if createdDog.ID != dogResponse.ID || newDog.Age != createdDog.Age || newDog.Name != createdDog.Name {
		t.Errorf("expected %+v, got %+v", dogResponse, createdDog)
	}
}

func TestMalformedBodyAddDogHandler(t *testing.T) {
	repository := repository.NewDogRepositoryFake()
	service := service.NewDogService(repository)
	handler := NewDogHandler(service)

	recorder := httptest.NewRecorder()
	invalidBody := `{"name"`
	body := bytes.NewReader([]byte(invalidBody))

	request := httptest.NewRequest("POST", "/api/dog/", body)

	handler.CreateDogHandler(recorder, request)

	var response object.ErrorResponse

	if err := json.NewDecoder(recorder.Body).Decode(&response); err != nil {
		t.Fatalf("invalid response body expected err response err: %v", err)
	}

	if recorder.Code != http.StatusBadRequest {
		t.Errorf("expected %v, got %v", http.StatusBadRequest, recorder.Code)
	}

	t.Logf("status: %v, error message: %v", response.Status, response.Error)

	if response.Error == "" {
		t.Error("expected error message got blank")
	}
}

func TestRepositoryFailureAddDogHandler(t *testing.T) {
	dog := object.Dog{
		ID:   0,
		Name: "Fenrir",
		Age:  1,
	}

	testRunBool := false

	dogByte, err := json.Marshal(dog)
	if err != nil {
		t.Fatalf("unable to convert dog to bytes err: %v", err)
	}

	body := bytes.NewReader(dogByte)

	repository := repository.NewDogRepositoryFake()
	repository.CreateDogFunc = func(dog object.Dog) (int, error) {
		testRunBool = true
		return 0, errors.New("storage unavailable")
	}

	service := service.NewDogService(repository)
	handler := NewDogHandler(service)

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("POST", "/api/dog", body)
	var response object.ErrorResponse

	handler.CreateDogHandler(recorder, request)

	if testRunBool != true {
		t.Fatalf("test behavior did not run")
	}

	if err := json.NewDecoder(recorder.Body).Decode(&response); err != nil {
		t.Fatalf("invalid response body expected err response err: %v", err)
	}

	if recorder.Code != http.StatusInternalServerError {
		t.Errorf("unexpected http status, got %v, expected %v", recorder.Code, http.StatusInternalServerError)
	}

	if response.Status != http.StatusInternalServerError {
		t.Errorf("unexpected status got %v, expected %v", response.Status, http.StatusBadRequest)
	}

	if response.Error != "unable to create dog using {ID:0 Name:Fenrir Age:1}, err: storage unavailable" {
		t.Errorf("%v", response.Error)
	}
}
