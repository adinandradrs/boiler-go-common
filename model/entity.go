package model

import "database/sql"

type BaseEntity struct {
	IsDeleted   bool           `json:"is_deleted"`
	CreatedBy   sql.NullString `json:"created_by"`
	CreatedDate sql.NullTime   `json:"created_date"`
	UpdatedBy   sql.NullString `json:"updated_by"`
	UpdatedDate sql.NullTime   `json:"updated_date"`
}