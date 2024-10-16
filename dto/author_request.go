package dto

type AuthorRequest struct {
	Name      string `json:"name" validate:"required,min=1"`
	BirthDate string `json:"birth_date" validate:"required"`
}
