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
	"strconv"
	"testing"
)

func TestGetDogByIDHandler(t *testing.T) {
	repository := repository.NewDogRepositoryFake()
	receivedID := 0

	repository.GetDogByIDFunc = func(id int, includeDeleted bool) (*object.Dog, error) {
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

	createdDog, err := service.GetDogByID(dogResponse.ID, false)
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
	testRunBool := false

	repository := repository.NewDogRepositoryFake()
	repository.CreateDogFunc = func(dog object.Dog) (int, error) {
		testRunBool = true
		return 0, errors.New("storage unavailable")
	}

	service := service.NewDogService(repository)
	handler := NewDogHandler(service)

	dog := object.Dog{
		ID:   0,
		Name: "Fenrir",
		Age:  1,
	}

	dogByte, err := json.Marshal(dog)
	if err != nil {
		t.Fatalf("unable to convert dog to bytes err: %v", err)
	}

	body := bytes.NewReader(dogByte)

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

	if response.Error != "unable to create dog using {ID:0 Name:Fenrir Age:1 CreatedAt:0001-01-01 00:00:00 +0000 UTC CreatedBy:0 UpdatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedBy:0 DeletedAt:<nil> DeletedBy:<nil>}, err: storage unavailable" {
		t.Errorf("%v", response.Error)
	}
}

func TestUpdateDogHandler(t *testing.T) {
	r := repository.NewDogRepositoryFake()
	s := service.NewDogService(r)
	h := NewDogHandler(s)

	dogUpdate := object.Dog{
		ID:   1,
		Name: "Fenrir",
		Age:  1,
	}

	bodyBytes, err := json.Marshal(dogUpdate)
	if err != nil {
		t.Fatalf("unable to process dog bytes from object %+v, err: %v", dogUpdate, err)
	}

	bodyBytesReader := bytes.NewReader(bodyBytes)

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("PATCH", "/api/dog/{id}", bodyBytesReader)
	idString := strconv.Itoa(dogUpdate.ID)
	request.SetPathValue("id", idString)

	h.UpdateDogHandler(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Errorf("invalid http response got %v, expected %v", recorder.Code, http.StatusOK)
	}

	var response object.SuccessResponse
	if err := json.NewDecoder(recorder.Body).Decode(&response); err != nil {
		t.Errorf("unable to decode response err %v", err)
	}

	if response.Status != http.StatusOK {
		t.Errorf("invalid response status: got %v, expected %v", response.Status, http.StatusOK)
	}

	updatedDog, err := s.GetDogByID(dogUpdate.ID, false)
	if err != nil {
		t.Fatalf("unable to get dog at ID: %v, err: %v", dogUpdate.ID, err)
	}

	if updatedDog.ID != dogUpdate.ID || updatedDog.Age != dogUpdate.Age || updatedDog.Name != dogUpdate.Name {
		t.Fatalf("after updated expected %+v, got %+v", dogUpdate, updatedDog)
	}

}

func TestMalformedBodyUpdateDogHandler(t *testing.T) {
	r := repository.NewDogRepositoryFake()
	s := service.NewDogService(r)
	h := NewDogHandler(s)

	invalidDog := `{"Name:"`
	invalidDogReader := bytes.NewReader([]byte(invalidDog))

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("PATCH", "/api/dog/{id}", invalidDogReader)
	request.SetPathValue("id", "1")

	h.UpdateDogHandler(recorder, request)

	if recorder.Code != http.StatusBadRequest {
		t.Errorf("invalid http status code, expected %v, got %v", http.StatusInternalServerError, recorder.Code)
	}

	var response object.ErrorResponse
	if err := json.NewDecoder(recorder.Body).Decode(&response); err != nil {
		t.Fatalf("unable to decode response body got %+v, err: %v", recorder.Body, err)
	}

	if response.Status != http.StatusBadRequest {
		t.Errorf("unexpected status, got %v, expected %v", response.Status, http.StatusInternalServerError)
	}

	expectedError := "invalid dog object body provided err: unexpected EOF"
	if response.Error != expectedError {
		t.Errorf("Invalid error: got: %v, expected %v", response.Error, expectedError)
	}
}

func TestRepositoryFailureUpdateDogHandler(t *testing.T) {
	r := repository.NewDogRepositoryFake()
	testRan := false
	dog := object.Dog{
		ID:   0,
		Name: "NA",
		Age:  0,
	}
	r.UpdateDogFunc = func(id int, dog object.Dog) (object.Dog, error) {
		testRan = true
		return dog, errors.New("repository update dog failure")
	}
	s := service.NewDogService(r)
	h := NewDogHandler(s)

	recorder := httptest.NewRecorder()

	dogBytes, err := json.Marshal(dog)
	if err != nil {
		t.Fatalf("unable to convert dog to bytes dog: %+v, err: %v", dog, err)
	}

	dogByteReader := bytes.NewReader(dogBytes)

	request := httptest.NewRequest("PATCH", "/api/dog/{id}", dogByteReader)
	idInt := strconv.Itoa(dog.ID)
	request.SetPathValue("id", idInt)

	h.UpdateDogHandler(recorder, request)

	if !testRan {
		t.Fatalf("update dog repository failure test never ran testRan = %v", testRan)
	}

	var response object.ErrorResponse
	if err := json.NewDecoder(recorder.Body).Decode(&response); err != nil {
		t.Fatalf("unable to decode response body %+v, err: %v", recorder.Body, err)
	}

	if recorder.Code != http.StatusInternalServerError {
		t.Errorf("expected status: %v, got: %v", http.StatusInternalServerError, recorder.Code)
	}

	if response.Status != http.StatusInternalServerError {
		t.Errorf("expected status: %v, got: %v", http.StatusInternalServerError, recorder.Code)
	}

	expectedErr := "unable to update dog using supplied values {ID:0 Name:NA Age:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC CreatedBy:0 UpdatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedBy:0 DeletedAt:<nil> DeletedBy:<nil>} for id 0, err: repository update dog failure"
	if response.Error != expectedErr {
		t.Errorf("expected: %v, got: %v", expectedErr, response.Error)
	}
}

func TestDeleteDogHandler(t *testing.T) {
	r := repository.NewDogRepositoryFake()
	s := service.NewDogService(r)
	h := NewDogHandler(s)

	dogId := 1

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("DELETE", "/api/dog/{id}", nil)
	dogIdString := strconv.Itoa(dogId)
	request.SetPathValue("id", dogIdString)

	h.DeleteDogHandler(recorder, request)

	var response object.SuccessResponse
	if err := json.NewDecoder(recorder.Body).Decode(&response); err != nil {
		t.Fatalf("unable to decode response body from %+v, err: %v", recorder.Body, err)
	}

	if recorder.Code != http.StatusOK || response.Status != http.StatusOK {
		t.Errorf("invalid http status code expected: %v, got %v and %v", http.StatusOK, recorder.Code, response.Status)
	}

	validResponse := "Deleted dog at ID: 1"
	if response.Response != validResponse {
		t.Errorf("got: %v, expected %v", response.Response, validResponse)
	}
}
