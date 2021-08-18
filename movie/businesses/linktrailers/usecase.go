package linktrailers

import (
	"context"
	"movie/businesses"
	"time"

	"gorm.io/gorm"
)

type LinkUseCase struct {
	linkRepository Repository
}

func NewLinkUseCase(linkrepo Repository) UseCase {
	return &LinkUseCase{
		linkRepository: linkrepo,
	}
}

func (luc LinkUseCase) Create(ctx context.Context, domain *Domain) (Domain, error) {
	// is existed user?
	existed, err := luc.linkRepository.GetAllLinkRepository(ctx, domain)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return Domain{}, businesses.ErrInternalServer
		}
	}
	// if data existed
	if existed != (Domain{}) {
		return Domain{}, businesses.ErrLinkExisted
	}
	domain.Created_At = time.Now()
	// store user data
	domainRegister, err := luc.linkRepository.CreateLinkRepository(ctx, domain)
	if err != nil {
		return Domain{}, err
	}
	return domainRegister, nil
}

func (luc LinkUseCase) Update(ctx context.Context, idLink int, domain *Domain) error {
	// is user id existed ?
	_, err := luc.linkRepository.GetLinkbyId(ctx, idLink)
	if err != nil {
		return err
	}
	// update user data
	domain.Updated_At = time.Now()
	err = luc.linkRepository.UpdateLinkRepository(ctx, idLink, domain)
	if err != nil {
		return err
	}
	return nil
}

func (luc LinkUseCase) Delete(ctx context.Context, idLink int, domain *Domain) error {
	// is user id existed ?
	_, err := luc.linkRepository.GetLinkbyId(ctx, idLink)
	if err != nil {
		return err
	}
	// delete user data
	err = luc.linkRepository.DeleteLinkRepository(ctx, idLink, domain)
	if err != nil {
		return err
	}
	return nil
}

func (luc LinkUseCase) GetUrl(ctx context.Context, domain *Domain) (Domain, error) {
	// is existed user?
	existed, err := luc.linkRepository.GetAllLinkRepository(ctx, domain)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return Domain{}, businesses.ErrInternalServer
		}
	}
	// if data existed
	if existed != (Domain{}) {
		return *domain, nil
	}
	return *domain, nil
}
