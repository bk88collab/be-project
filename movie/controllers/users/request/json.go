package request

import "movie/businesses/users"

type Users struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (request *Users) ToDomain() *users.Domain {
	return &users.Domain{
		Firstname: request.Firstname,
		Lastname:  request.Lastname,
		Username:  request.Username,
		Email:     request.Email,
		Password:  request.Password,
	}
}
