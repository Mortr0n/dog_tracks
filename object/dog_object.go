package object

import "time"

type Dog struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	Age       int        `json:"age"`
	CreatedAt time.Time  `json:"createdat"`
	CreatedBy int        `json:createdby"`
	UpdatedAt time.Time  `json:"updatedat"`
	UpdatedBy int        `json:"updatedby"`
	DeletedAt *time.Time `json:"deletedat"`
	DeletedBy *int       `json:"deletedby"`
}
