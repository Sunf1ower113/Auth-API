package movie

type CreateMovieDTO struct {
	Title        string `json:"title"`
	Category     string `json:"category"`
	Type         string `json:"project_type"`
	AgeGroup     string `json:"age_group"`
	CreationYear string `json:"creation_year"`
	Timing       string `json:"timing"`
	Tags         string `json:"tags"`
	Description  string `json:"description"`
	Director     string `json:"director"`
	Producer     string `json:"producer"`
}
type UpdateMovieDTO struct {
	Id           int64  `json:"id"`
	Title        string `json:"title"`
	Category     string `json:"category"`
	Type         string `json:"project_type"`
	AgeGroup     string `json:"age_group"`
	CreationYear string `json:"creation_year"`
	Timing       string `json:"timing"`
	Tags         string `json:"tags"`
	Description  string `json:"description"`
	Director     string `json:"director"`
	Producer     string `json:"producer"`
}
