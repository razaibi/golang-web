package models

import "time"

type SampleProgram struct {
	SampleProgramId int       `db:"SampleProgramId" json:"sampleProgramId"`
	Name            string    `db:"Name" json:"name"`
	IsActive        bool      `db:"IsActive" json:"isActive"`
	CreatedOn       time.Time `db:"CreatedOn" json:"createdOn"`
	LastModifiedOn  time.Time `db:"LastModifiedOn" json:"lastModifiedOn"`
}
