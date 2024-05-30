package models

import "time"

type Sample struct {
	SampleId        int       `db:"SampleId" json:"SampleId"`
	SampleProgramId int       `db:"SampleProgramId" json:"SampleProgramId"`
	Title           string    `db:"Title" json:"Title"`
	Content         string    `db:"Content" json:"Content"`
	IsActive        bool      `db:"IsActive" json:"IsActive"`
	CreatedOn       time.Time `db:"CreatedOn" json:"CreatedOn"`
	LastModifiedOn  time.Time `db:"LastModifiedOn" json:"LastModifiedOn"`
}
