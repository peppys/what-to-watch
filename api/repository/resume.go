package repository

import (
	"github.com/PeppyS/what-to-watch/api/model"
)

// ResumePostgresRepository defines repo structure that interacts with postgres database
type ResumePostgresRepository struct {
	// TODO - Include postgres DB connection here
}

// NewResume instantiates ResumePostgresRepository
func NewResume() *ResumePostgresRepository {
	return &ResumePostgresRepository{}
}

// Get retrieves active resume TODO: get from postgres
func (r *ResumePostgresRepository) Get() (model.Resume, error) {
	return model.Resume{
		FirstName: "Peppy",
		LastName:  "Sisay",
		Headline:  "",
		Summary:   "",
		Experience: []struct {
			Title, Company, Location, Description string
			model.DateRange
		}{
			{"Junior Software Engineer", "Tradesy", "Los Angeles, CA", "", model.DateRange{Display: "January 2016 - Present"}},
		},
	}, nil
}
