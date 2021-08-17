package users

import (
	"context"
	"time"
)

type Domain struct {
	Id_user    int
	First_name string
	Last_name  string
	User_name  string
	Email      string
	Password   string
	Created_At time.Time
	Updated_At time.Time
}

type UseCase interface {
	Register(ctx context.Context, domain *Domain) (Domain, error)
	Update(ctx context.Context, userId int, domain *Domain) error
	Delete(ctx context.Context, userId int, domain *Domain) error
	GetProfile(ctx context.Context, userName string) (Domain, error)
}

type Repository interface {
	Insert(ctx context.Context, domain *Domain) (Domain, error)
	UpdateUser(ctx context.Context, userId int, domain *Domain) error
	DeleteUser(ctx context.Context, userId int, domain *Domain) error
	GetUserByUsername(ctx context.Context, userName string) (Domain, error)
	GetUserByID(ctx context.Context, userId int) (Domain, error)
}
