package movie

type MovieStorage interface {
	GetMovie() (*Movie, error)
	CreateMovie() (*Movie, error)
	UpdateMovie() (*Movie, error)
}