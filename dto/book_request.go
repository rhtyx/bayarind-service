package dto

type BookRequest struct {
	ISBN     string `json:"isbn" validate:"required,isbn"`
	Title    string `json:"title" validate:"required,min=1"`
	AuthorID int64  `json:"author_id" validate:"required,numeric"`
}
