package users

import (
	"movie/businesses/users"
	"time"

	"gorm.io/gorm"
)

type Users struct {
	Id        int
	Firstname string
	Lastname  string
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdateAt  time.Time
	DeletedAt gorm.DeletedAt
}

func (record *Users) toDomain() users.Domain {
	return users.Domain{
		Id:        record.Id,
		Firstname: record.Firstname,
		Lastname:  record.Lastname,
		Username:  record.Username,
		Email:     record.Email,
		Password:  record.Password,
		CreatedAt: record.CreatedAt,
		UpdateAt:  record.UpdateAt,
	}
}

func fromDomainToRecord(userDomain users.Domain) *Users {
	return &Users{
		Id:        userDomain.Id,
		Firstname: userDomain.Firstname,
		Lastname:  userDomain.Lastname,
		Username:  userDomain.Username,
		Email:     userDomain.Email,
		Password:  userDomain.Password,
		CreatedAt: userDomain.CreatedAt,
		UpdateAt:  userDomain.UpdateAt,
	}
}
