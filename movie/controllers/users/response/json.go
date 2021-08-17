package response

import "movie/businesses/users"

type Users struct {
	Id_user    int    `json:"id_user"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	User_name  string `json:"user_name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}

func FromDomain(domain users.Domain) Users {
	return Users{
		Id_user:    domain.Id_user,
		First_name: domain.First_name,
		Last_name:  domain.Last_name,
		User_name:  domain.User_name,
		Email:      domain.Email,
		Password:   domain.Password,
	}
}
