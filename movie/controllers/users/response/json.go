package response

import "movie/businesses/users"

type Users struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func FromDomain(domain users.Domain) Users {
	return Users{
		Firstname: domain.Firstname,
		Lastname:  domain.Lastname,
		Username:  domain.Username,
		Email:     domain.Email,
		Password:  domain.Password,
	}
}
