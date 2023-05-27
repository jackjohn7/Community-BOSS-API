package v1

import (
	"time"
)

type Quarter struct {
	Id          uint16
	Year        int
	Season      string
	DateUpdated time.Time
}

type Subject struct {
	SubjectId string
	Name      string
	QuarterId uint16
}

type Course struct {
	CourseId  string
	Name      string
	SubjectId string
}
