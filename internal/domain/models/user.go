package models

type User struct {
	UID      string `json:"uid"`
	Name     string `json:"name" validate:"required,min=2,max=50"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,containsany=!@#$%^&*()"`
}

func (u User) ID() string {
	return u.UID
}

func (u User) ToResponse() UserResponse {
	return UserResponse {
		UID: u.UID,
		Name: u.Name,
		Email: u.Email,
	}
}

type UserResponse struct {
	UID   string `json:"uid"`
	Name  string `json:"name"`
	Email string `json:"email"`
}