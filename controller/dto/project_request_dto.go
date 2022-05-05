package dto

type ProjectRequestDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Nda_accept  bool   `json:"nda_accept"`
}
