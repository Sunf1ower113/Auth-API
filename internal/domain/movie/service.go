package movie

import "context"

type MovieService interface {
	GetMovieByID(ctx context.Context, id int) (*Movie, error)
}

type movieService struct {
	storage MovieStorage
}

func NewMovieService(storage MovieStorage) MovieService {
	return &movieService{
		storage: storage,
	}
}

func (u *movieService) GetMovieByID(ctx context.Context, id int) (*Movie, error) {
	//TODO implement me
	panic("implement me")
}
