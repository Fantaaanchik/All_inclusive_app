package models

// TokenResponse represents an token response
type TokenResponse struct {
	Token string `json:"token"`
}

type GetUserResponse struct {
	User User `json:"user"`
}

type GetAllUsersResponse struct {
	Users []User `json:"users"`
}

// SuccessResponse represents an success response
type SuccessResponse struct {
	SuccessDescription string `json:"description"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	ErrorDescription string `json:"error"`
}

// CreateUserResponse represents a create user response
type CreateUserResponse struct {
	Message string `json:"message"`
}

// UpdateUserResponse represents a update user response
type UpdateUserResponse struct {
	Message string `json:"message"`
}

// DeleteUserResponse represents a delete user response
type DeleteUserResponse struct {
	Message string `json:"message"`
}
