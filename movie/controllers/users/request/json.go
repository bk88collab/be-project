package request

import "movie/businesses/users"

type Users struct {
	Id_user    int    `json:"id_user"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	User_name  string `json:"user_name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}

func (request *Users) ToDomain() *users.Domain {
	return &users.Domain{
		Id_user:    request.Id_user,
		First_name: request.First_name,
		Last_name:  request.Last_name,
		User_name:  request.User_name,
		Email:      request.Email,
		Password:   request.Password,
	}
}
