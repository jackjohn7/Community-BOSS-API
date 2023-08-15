package v1

import (
	"time"
)

type Quarter struct {
	Id          uint16    `json:"id"`
	Year        int       `json:"year"`
	Season      string    `json:"season"`
	DateUpdated time.Time `json:"dateUpdated"`
}

type Subject struct {
	SubjectId string `json:"subjectId"`
	Name      string `json:"name"`
	QuarterId uint16 `json:"quarterId"`
}

type Course struct {
	CourseId  string `json:"courseId"`
	Name      string `json:"name"`
	SubjectId string `json:"subjectId"`
}

type Section struct {
	SectionId         string `json:"sectionId"`
	CourseId          string `json:"courseId"`
	CallNumber        string `json:"callNumber"`
	SectionTitle      string `json:"sectionTitle"`
	CreditHours       int    `json:"creditHours"`
	Activity          string `json:"activity"`
	Modality          string `json:"modality"`
	Days              string `json:"days"`
	Location          string `json:"location"`
	Instructor        string `json:"instructor"`
	Status            string `json:"status"`
	CombinedDays      string `json:"combinedDays"`
	CombinedLocation  string `json:"combinedLocation"`
	CombinedTimeStart string `json:"combinedTimeStart"`
	CombinedTimeEnd   string `json:"combinedTimeEnd"`
	IsCombined        bool   `json:"isCombined"`
	Note              string `json:"note"`
	TimeStart         string `json:"timeStart"`
	TimeStop          string `json:"timeStop"`
}
