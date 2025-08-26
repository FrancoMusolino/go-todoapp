package dtos

type RegisterUserDto struct {
	Username string `json:"username" validate:"required,min=4"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginDto struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginResponseDto struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Token    string `json:"token"`
}
