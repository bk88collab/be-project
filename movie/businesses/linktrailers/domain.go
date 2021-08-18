package linktrailers

import (
	"context"
	"time"
)

type Domain struct {
	Id_trailer int
	Url        string
	Created_At time.Time
	Updated_At time.Time
}

type UseCase interface {
	Create(ctx context.Context, domain *Domain) (Domain, error)
	Update(ctx context.Context, idLink int, domain *Domain) error
	Delete(ctx context.Context, idLink int, domain *Domain) error
	GetUrl(ctx context.Context, domain *Domain) (Domain, error)
}

type Repository interface {
	CreateLinkRepository(ctx context.Context, domain *Domain) (Domain, error)
	UpdateLinkRepository(ctx context.Context, idLink int, domain *Domain) error
	DeleteLinkRepository(ctx context.Context, idLink int, domain *Domain) error
	GetAllLinkRepository(ctx context.Context, domain *Domain) (Domain, error)
	GetLinkbyId(ctx context.Context, idLink int) (Domain, error)
}
