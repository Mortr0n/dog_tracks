package repository

import (
	"dog_tracks/object"
	"time"
)

// ! Dog object
var userId = 0
var dogTime = time.Date(2026, time.October, 28, 11, 0, 0, 0, time.UTC)
var deleteTime = time.Date(2026, time.October, 28, 3, 0, 0, 0, time.UTC)
var trainTime = time.Date(2026, time.October, 28, 1, 0, 0, 0, time.UTC)

var Freyja = object.Dog{
	ID:        1,
	Name:      "Freyja",
	Age:       5,
	CreatedAt: dogTime,
	CreatedBy: userId,
	UpdatedAt: dogTime,
	UpdatedBy: userId,
	DeletedAt: nil,
	DeletedBy: nil,
}
var Fenrir = object.Dog{
	ID:        2,
	Name:      "Fenrir",
	Age:       1,
	CreatedAt: dogTime,
	CreatedBy: userId,
	UpdatedAt: dogTime,
	UpdatedBy: userId,
	DeletedAt: nil,
	DeletedBy: nil,
}

var Badger = object.Dog{
	ID:        3,
	Name:      "Badger",
	Age:       10,
	CreatedAt: dogTime,
	CreatedBy: userId,
	UpdatedAt: dogTime,
	UpdatedBy: userId,
	DeletedAt: nil,
	DeletedBy: nil,
}

var DeletedDog = object.Dog{
	ID:        4,
	Name:      "Deleter",
	Age:       62,
	CreatedAt: dogTime,
	CreatedBy: userId,
	UpdatedAt: dogTime,
	UpdatedBy: userId,
	DeletedAt: &deleteTime,
	DeletedBy: &userId,
}

var dogMap = map[int]object.Dog{
	1: Freyja,
	2: Fenrir,
	3: Badger,
	4: DeletedDog,
}

// ! Training Type
var recall = object.Training_Type{
	ID:          1,
	Name:        "Recall",
	Description: "Training for dog recall and name recognition",
	CreatedAt:   dogTime,
	CreatedBy:   userId,
	UpdatedAt:   dogTime,
	UpdatedBy:   userId,
	DeletedAt:   &deleteTime,
	DeletedBy:   &userId,
}

var looseLeash = object.Training_Type{
	ID:          2,
	Name:        "Loose Leash",
	Description: "Training for loose leash walking",
	CreatedAt:   dogTime,
	CreatedBy:   userId,
	UpdatedAt:   dogTime,
	UpdatedBy:   userId,
	DeletedAt:   &deleteTime,
	DeletedBy:   &userId,
}

var leaveIt = object.Training_Type{
	ID:          3,
	Name:        "Leave it",
	Description: "Training for dog to leave unsafe or otherwise objects",
	CreatedAt:   dogTime,
	CreatedBy:   userId,
	UpdatedAt:   dogTime,
	UpdatedBy:   userId,
	DeletedAt:   &deleteTime,
	DeletedBy:   &userId,
}

var stay = object.Training_Type{
	ID:          4,
	Name:        "Stay",
	Description: "Training for dog to no movey movey",
	CreatedAt:   dogTime,
	CreatedBy:   userId,
	UpdatedAt:   dogTime,
	UpdatedBy:   userId,
	DeletedAt:   &deleteTime,
	DeletedBy:   &userId,
}

var trainingTypeMap = map[int]object.Training_Type{
	1: recall,
	2: looseLeash,
	3: leaveIt,
	4: stay,
}

// ! Training Session
var recallTrain1 = object.Training_Session{
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

var recallTrain2 = object.Training_Session{
	ID:              2,
	DogID:           1,
	TrainingTypeID:  1,
	StartedAt:       trainTime,
	Durationseconds: 500,
	Notes:           "dog recalled second",
	CreatedAt:       dogTime,
	CreatedBy:       userId,
	UpdatedAt:       dogTime,
	UpdatedBy:       userId,
	DeletedAt:       &deleteTime,
	DeletedBy:       &userId,
}

var leaveitTrain1 = object.Training_Session{
	ID:              3,
	DogID:           2,
	TrainingTypeID:  3,
	StartedAt:       trainTime,
	Durationseconds: 900,
	Notes:           "dog left it first",
	CreatedAt:       dogTime,
	CreatedBy:       userId,
	UpdatedAt:       dogTime,
	UpdatedBy:       userId,
	DeletedAt:       nil,
	DeletedBy:       nil,
}

var looseLeashTrain1 = object.Training_Session{
	ID:              4,
	DogID:           3,
	TrainingTypeID:  2,
	StartedAt:       trainTime,
	Durationseconds: 1500,
	Notes:           "dog loose leashed first",
	CreatedAt:       dogTime,
	CreatedBy:       userId,
	UpdatedAt:       dogTime,
	UpdatedBy:       userId,
	DeletedAt:       nil,
	DeletedBy:       nil,
}

var looseLeashTrain2 = object.Training_Session{
	ID:              5,
	DogID:           3,
	TrainingTypeID:  2,
	StartedAt:       trainTime,
	Durationseconds: 900,
	Notes:           "dog loose leashed second",
	CreatedAt:       dogTime,
	CreatedBy:       userId,
	UpdatedAt:       dogTime,
	UpdatedBy:       userId,
	DeletedAt:       nil,
	DeletedBy:       nil,
}

var looseLeashTrain3 = object.Training_Session{
	ID:              6,
	DogID:           2,
	TrainingTypeID:  2,
	StartedAt:       trainTime,
	Durationseconds: 1500,
	Notes:           "dog loose leashed third",
	CreatedAt:       dogTime,
	CreatedBy:       userId,
	UpdatedAt:       dogTime,
	UpdatedBy:       userId,
	DeletedAt:       &deleteTime,
	DeletedBy:       &userId,
}

var TrainingSessionMap = map[int]object.Training_Session{
	1: recallTrain1,
	2: recallTrain2,
	3: looseLeashTrain1,
	4: looseLeashTrain3,
	5: looseLeashTrain2,
	6: leaveitTrain1,
}
