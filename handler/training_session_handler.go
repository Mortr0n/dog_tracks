package handler

import (
	"dog_tracks/helper"
	"dog_tracks/service"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type TrainingSessionHandler struct {
	TrainingService service.TrainingService
}

func NewTrainingSessionHandler(s service.TrainingService) TrainingSessionHandler {
	return TrainingSessionHandler{
		TrainingService: s,
	}
}

func (th TrainingSessionHandler) GetTrainingSessionByIDHandler(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	deleted := r.PathValue("deleted")

	includeDeleted, _ := strconv.ParseBool(deleted)

	id, err := strconv.Atoi(idString)
	if err != nil {
		errMsg := fmt.Sprintf("unable to convert id: %v to int err: %v", idString, err)
		helper.ErrorResponse(http.StatusBadRequest, errMsg, w)
		return
	}

	session, err := th.TrainingService.GetTrainingSessionByID(id, includeDeleted)
	if err != nil {
		errMsg := fmt.Sprintf("unable to get dog by id: %v, err: %v", idString, err)
		helper.ErrorResponse(http.StatusInternalServerError, errMsg, w)
		return
	}

	if session == nil {
		errMsg := fmt.Sprintf("invalid session returned from id: %v", id)
		helper.ErrorResponse(http.StatusInternalServerError, errMsg, w)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(session); err != nil {
		errMsg := fmt.Sprintf("unable to write training session response from session: %+v, err: %v", session, err)
		helper.ErrorResponse(http.StatusInternalServerError, errMsg, w)
	}

}

func (th TrainingSessionHandler) GetAllTrainingSessionsByDogIDHandler(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	deleted := r.PathValue("deleted")

	id, err := strconv.Atoi(idString)
	if err != nil {
		errMsg := fmt.Sprintf("unable to convert ID: %v to int, err: %v", idString, err)
		helper.ErrorResponse(http.StatusBadRequest, errMsg, w)
		return
	}

	includeDeleted, _ := strconv.ParseBool(deleted)

	sessions, err := th.TrainingService.GetAllTrainingSessionsByDogID(id, includeDeleted)
	if err != nil {
		errMsg := fmt.Sprintf("unable to get training sessions for dog ID: %v, err: %v", id, err)
		helper.ErrorResponse(http.StatusInternalServerError, errMsg, w)
		return
	}

	if len(sessions) < 1 {
		errMsg := fmt.Sprintf("no sessions returned for dog ID: %v", id)
		helper.ErrorResponse(http.StatusOK, errMsg, w)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(sessions); err != nil {
		errMsg := fmt.Sprintf("unable to write session response err: %v", err)
		helper.ErrorResponse(http.StatusInternalServerError, errMsg, w)
		return
	}
}
