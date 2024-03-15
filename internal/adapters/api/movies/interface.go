package movies

import (
	"auth-api/internal/domain/movie"
	"context"
)

type MovieService interface {
	CreateMovie(ctx context.Context, dto *movie.CreateMovieDTO) (*movie.Movie, error)
}
