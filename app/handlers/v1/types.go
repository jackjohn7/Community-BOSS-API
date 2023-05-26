package v1

import (
	"time"
)

type Quarter struct {
	Id          int
	Year        int
	Season      string
	DateUpdated time.Time
}

type Course struct {
}
