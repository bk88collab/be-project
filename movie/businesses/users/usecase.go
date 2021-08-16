package users

import (
	"context"
	"movie/businesses"
)

type UserUseCase struct {
	userRepository Repository
}

func NewUserUseCase(ur Repository) UseCase {
	return &UserUseCase{
		userRepository: ur,
	}
}

func (uc UserUseCase) Register(ctx context.Context, domain *Domain) (Domain, error) {

	// is existed user?
	existed, err := uc.userRepository.GetUserByUsername(ctx, domain.Username)
	if err != nil {
		return Domain{}, err
	}

	// if data existed
	if existed != (Domain{}) {
		return Domain{}, businesses.ErrDuplicateData
	}

	// store user data
	domainRegister, err := uc.userRepository.Insert(ctx, domain)
	if err != nil {
		return Domain{}, err
	}
	return domainRegister, nil

}

func (uc UserUseCase) Update(ctx context.Context, domain *Domain) (Domain, error) {

	// is found user id?
	_, err := uc.userRepository.GetUserByID(ctx, domain.Id)
	if err != nil {
		return Domain{}, businesses.ErrUserIdNotFound
	}

	domainUpdate, err := uc.userRepository.UpdateUser(ctx, domain)
	if err != nil {
		return Domain{}, err
	}
	return domainUpdate, nil

}
