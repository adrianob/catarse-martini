package models

type User struct {
	Id       int64
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Users struct {
	Collection []User `json:"users"`
}

func NewUser(name, email, password string) User {
	return User{
		Name:     name,
		Email:    email,
		Password: password,
	}
}
