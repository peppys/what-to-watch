package model

import (
	"time"
)

// DateRange type
type DateRange struct {
	Start        time.Time
	End          time.Time
	Display string
}

// Resume data model
type Resume struct {
	FirstName  string
	LastName   string
	Headline   string
	Summary    string
	Experience []struct {
		Title       string
		Company     string
		Location    string
		Description string
		DateRange
	}
	Education []struct {
		Degree   string
		Major    string
		School   string
		Location string
		DateRange
	}
}
