package service

// MovieService defines service structure
type MovieService struct {
	// TODO add ES Client
}

// NewMovie instantiates MovieService
func NewMovie() *MovieService {
	return &MovieService{}
}

// Get returns movie
func (rs *MovieService) Get() (error) {
	return nil
}
