package requests

import "errors"

// UserRequest represents the request payload for user-related operations.
type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *UserRequest) Validate() error {
	if u.Username == "" || len(u.Username) < 3 || len(u.Username) > 50 {
		return errors.New("username must be between 3 and 50 characters long")
	}
	if u.Password == "" || len(u.Password) < 8 || len(u.Password) > 50 {
		return errors.New("password must be between 8 and 50 characters long")
	}
	return nil
}

// UserUpdateRequest represents the request payload for updating user information.
type UserUpdateRequest struct {
	Username string `json:"username"`
	Password string `json:"password,omitempty"`
	Role     string `json:"role,omitempty"`
}
