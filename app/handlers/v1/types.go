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

type Section struct {
	SectionId         string
	CourseId          string
	CallNumber        string
	SectionTitle      string
	CreditHours       int
	Activity          string
	Modality          string
	Days              string
	Location          string
	Instructor        string
	Status            string
	CombinedDays      string
	CombinedLocation  string
	CombinedTimeStart string
	CombinedTimeEnd   string
	IsCombined        bool
	Note              string
	TimeStart         string
	TimeStop          string
}
