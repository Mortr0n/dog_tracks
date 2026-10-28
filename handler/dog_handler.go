package handler

import (
	"dog_tracks/helper"
	"dog_tracks/object"
	"dog_tracks/service"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type DogHandler struct {
	DogService service.DogService
}

func NewDogHandler(ds service.DogService) DogHandler {
	return DogHandler{
		DogService: ds,
	}
}

func (dh DogHandler) HelloHandler(w http.ResponseWriter, r *http.Request) {
	dogName := "Freyja"

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(dogName))
}

func (dh DogHandler) GetDogByIDHandler(w http.ResponseWriter, r *http.Request) {

	idString := r.PathValue("id")
	idInt, err := strconv.Atoi(idString)
	if err != nil {
		errMsg := fmt.Sprintf("Unable to convert id %v to string: err %v", idString, err)
		helper.ErrorResponse(http.StatusBadRequest, errMsg, w)
		return
	}

	dog, err := dh.DogService.GetDogByID(idInt)
	if err != nil {
		errMsg := fmt.Sprintf("unable to get dog by id %v: err %v", idInt, err)
		helper.ErrorResponse(http.StatusInternalServerError, errMsg, w)
		return
	}

	if dog != nil {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		// helper.SuccessResponse(http.StatusOK, dog, w)
		if err := json.NewEncoder(w).Encode(dog); err != nil {
			errMsg := fmt.Sprintf("unable to write dog response for %v: err %v", dog, err)
			helper.ErrorResponse(http.StatusInternalServerError, errMsg, w)
		}
	}
}

func (dh DogHandler) CreateDogHandler(w http.ResponseWriter, r *http.Request) {
	var dog object.Dog

	if err := json.NewDecoder(r.Body).Decode(&dog); err != nil {
		errMsg := fmt.Sprintf("invalid dog object body provided err: %v", err)
		helper.ErrorResponse(http.StatusBadRequest, errMsg, w)
		return
	}

	dogId, err := dh.DogService.CreateDog(dog)
	if err != nil {
		errMsg := fmt.Sprintf("unable to create dog using %+v, err: %v", dog, err)
		helper.ErrorResponse(http.StatusInternalServerError, errMsg, w)
		return
	}

	dog.ID = dogId
	helper.SuccessResponse(http.StatusOK, dog, w)
}

func (dh DogHandler) UpdateDogHandler(w http.ResponseWriter, r *http.Request) {
	var dog object.Dog

	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		errMsg := fmt.Sprintf("Unable to convert id %v to string: err %v", idString, err)
		helper.ErrorResponse(http.StatusBadRequest, errMsg, w)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&dog); err != nil {
		errMsg := fmt.Sprintf("invalid dog object body provided err: %v", err)
		helper.ErrorResponse(http.StatusBadRequest, errMsg, w)
		return
	}

	updatedDog, err := dh.DogService.UpdateDog(id, dog)
	if err != nil {
		errMsg := fmt.Sprintf("unable to update dog using supplied values %+v for id %v, err: %v", dog, id, err)
		helper.ErrorResponse(http.StatusInternalServerError, errMsg, w)
		return
	}

	helper.SuccessResponse(http.StatusOK, updatedDog, w)
}

func (dh DogHandler) DeleteDogHandler(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		errMsg := fmt.Sprintf("Unable to convert id %v to string: err %v", idString, err)
		helper.ErrorResponse(http.StatusBadRequest, errMsg, w)
		return
	}

	if err := dh.DogService.DeleteDog(id); err != nil {
		errMsg := fmt.Sprintf("unable to delete dog at id: %v, err: %v", id, err)
		helper.ErrorResponse(http.StatusInternalServerError, errMsg, w)
		return
	}

	delMsg := fmt.Sprintf("Deleted dog at ID: %v", id)
	helper.SuccessResponse(http.StatusOK, delMsg, w)
}
