package users

import (
	"movie/businesses/users"
	"time"
)

type Users struct {
	Id_user    int
	First_name string
	Last_name  string
	User_name  string
	Email      string
	Password   string
	Created_At time.Time
	Updated_At time.Time
}

func (record *Users) toDomain() users.Domain {
	return users.Domain{
		Id_user:    record.Id_user,
		First_name: record.First_name,
		Last_name:  record.Last_name,
		User_name:  record.User_name,
		Email:      record.Email,
		Password:   record.Password,
		Created_At: record.Created_At,
		Updated_At: record.Updated_At,
	}
}

func fromDomainToRecord(userDomain *users.Domain) *Users {
	return &Users{
		Id_user:    userDomain.Id_user,
		First_name: userDomain.First_name,
		Last_name:  userDomain.Last_name,
		User_name:  userDomain.User_name,
		Email:      userDomain.Email,
		Password:   userDomain.Password,
		Created_At: userDomain.Created_At,
		Updated_At: userDomain.Updated_At,
	}
}
