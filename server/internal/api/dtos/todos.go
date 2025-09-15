package dtos

type CreateTodoDto struct {
	Title       string `json:"title" validate:"required,min=3"`
	Description string `json:"description" validate:"required,max=256"`
}
