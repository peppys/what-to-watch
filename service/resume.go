package service

import (
	"github.com/PeppyS/personal-site-api/model"
)

type resumeRepository interface {
	Get() (model.Resume, error)
}

// ResumeService defines service structure
type ResumeService struct {
	repository resumeRepository
}

// NewResume instantiates ResumeService
func NewResume(rp resumeRepository) *ResumeService {
	return &ResumeService{
		repository: rp,
	}
}

// Get returns resume from repository
func (rs *ResumeService) Get() (model.Resume, error) {
	return rs.repository.Get()
}
