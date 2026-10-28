package object

import "time"

type Training_Type struct {
	ID          int
	Name        string
	Description string
	CreatedAt   time.Time
	CreatedBy   int
	UpdatedAt   time.Time
	UpdatedBy   int
	DeletedAt   *time.Time
	DeletedBy   *int
}
