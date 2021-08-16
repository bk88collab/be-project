package users

import (
	"context"
	"time"
)

type Domain struct {
	Id        int
	Firstname string
	Lastname  string
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdateAt  time.Time
}

type UseCase interface {
	Register(ctx context.Context, domain *Domain) (Domain, error)
	Update(ctx context.Context, domain *Domain) (Domain, error)
}

type Repository interface {
	Insert(ctx context.Context, domain *Domain) (Domain, error)
	UpdateUser(ctx context.Context, domain *Domain) (Domain, error)
	GetUserByUsername(ctx context.Context, username string) (Domain, error)
	GetUserByID(ctx context.Context, id int) (Domain, error)
}
