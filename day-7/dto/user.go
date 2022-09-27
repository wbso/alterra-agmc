package dto

type UserResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name" `
	Email string `json:"email"`
}

type UserCreateRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserUpdateRequest struct {
	Name string `json:"name"`
}
