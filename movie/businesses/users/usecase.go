package users

import (
	"context"
	"movie/businesses"
	"time"

	"gorm.io/gorm"
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
	// is existed user ?
	existed, err := uc.userRepository.GetUserByUsername(ctx, domain.User_name)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return Domain{}, businesses.ErrInternalServer
		}
	}
	// if data existed
	if existed != (Domain{}) {
		return Domain{}, businesses.ErrDuplicateData
	}
	domain.Created_At = time.Now()
	// store user data
	domainRegister, err := uc.userRepository.Insert(ctx, domain)
	if err != nil {
		return Domain{}, err
	}
	return domainRegister, nil
}

func (uc *UserUseCase) Update(ctx context.Context, userId int, domain *Domain) error {
	// is user id existed ?
	_, err := uc.userRepository.GetUserByID(ctx, userId)
	if err != nil {
		return err
	}
	// update user data
	domain.Updated_At = time.Now()
	err = uc.userRepository.UpdateUser(ctx, userId, domain)
	if err != nil {
		return err
	}
	return nil
}

func (uc *UserUseCase) Delete(ctx context.Context, userId int, domain *Domain) error {
	// is user id existed ?
	_, err := uc.userRepository.GetUserByID(ctx, userId)
	if err != nil {
		return err
	}
	// delete user data
	err = uc.userRepository.DeleteUser(ctx, userId, domain)
	if err != nil {
		return err
	}
	return nil
}

func (uc *UserUseCase) GetProfile(ctx context.Context, userName string) (Domain, error) {
	// is existed user ?
	existed, err := uc.userRepository.GetUserByUsername(ctx, userName)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return Domain{}, businesses.ErrInternalServer
		}
	}
	// get data user
	if existed != (Domain{}) {
		return existed, nil
	}
	return existed, nil
}
