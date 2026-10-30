package object

import "time"

type Training_Type struct {
	ID          int        `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	CreatedAt   time.Time  `json:"createdat"`
	CreatedBy   int        `json:"createdby"`
	UpdatedAt   time.Time  `json:"updatedat"`
	UpdatedBy   int        `json:"updatedby"`
	DeletedAt   *time.Time `json:"deletedat"`
	DeletedBy   *int       `json:"deletedby"`
}

type Training_Session struct {
	ID              int        `json:"id"`
	DogID           int        `json:"dogid"`
	TrainingTypeID  int        `json:"trainingtypeid"`
	StartedAt       time.Time  `json:"startedat"`
	Durationseconds int        `json:"duration"`
	Notes           string     `json:"notes"`
	CreatedAt       time.Time  `json:"createdat"`
	CreatedBy       int        `json:"createdby"`
	UpdatedAt       time.Time  `json:"updatedat"`
	UpdatedBy       int        `json:"updatedby"`
	DeletedAt       *time.Time `json:"deletedat"`
	DeletedBy       *int       `json:"deletedby"`
}
