package model

import "time"

type ReferenceList struct {
	ID          int64     `json:"id" gorm:"column:id;type:bigint;"`
	Name        string    `json:"name"`
	Code        string    `json:"code"`
	DateCreated time.Time `json:"dateCreated"`
	DateUpdated time.Time `json:"dateUpdated"`
	CreatedBy   int64     `json:"createdBy"`
	UpdatedBy   int64     `json:"updatedBy"`
	Type        string    `json:"type"`
	Description string    `json:"description"`
	Status      int       `json:"status"`
}