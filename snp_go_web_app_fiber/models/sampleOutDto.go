package models

import "time"

type SampleOutDto struct {
	SampleId          int       `db:"SampleId" json:"sampleId"`
	SampleProgramId   int       `db:"SampleProgramId" json:"sampleProgramId"`
	SampleProgramName string    `db:"SampleProgramName" json:"sampleProgramName"`
	Title             string    `db:"Title" json:"title"`
	Content           string    `db:"Content" json:"content"`
	IsActive          bool      `db:"IsActive" json:"isActive"`
	CreatedOn         time.Time `db:"CreatedOn" json:"createdOn"`
	LastModifiedOn    time.Time `db:"LastModifiedOn" json:"lastModifiedOn"`
}
