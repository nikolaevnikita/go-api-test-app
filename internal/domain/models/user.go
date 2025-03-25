package models

type User struct {
	UID      string	 `json:"uid"`
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
}

func (u User) ID() string {
	return u.UID
}