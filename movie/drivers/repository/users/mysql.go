package users

import (
	"context"
	"movie/businesses/users"

	"gorm.io/gorm"
)

type userRepository struct {
	dbConn *gorm.DB
}

func NewUserRepository(db *gorm.DB) users.Repository {
	return &userRepository{
		dbConn: db,
	}
}

func (repository *userRepository) Insert(ctx context.Context, domain *users.Domain) (users.Domain, error) {
	record := fromDomainToRecord(domain)

	result := repository.dbConn.Create(&record)
	if result.Error != nil {
		return users.Domain{}, result.Error
	}

	return record.toDomain(), nil
}

func (repository *userRepository) UpdateUser(ctx context.Context, userId int, domain *users.Domain) error {
	record := fromDomainToRecord(domain)

	record.Id_user = userId
	err := repository.dbConn.Where("id_user = ?", userId).Updates(&record).Error
	if err != nil {
		return err
	}

	return nil
}

func (repository *userRepository) GetUserByUsername(ctx context.Context, userName string) (users.Domain, error) {
	record := Users{}
	err := repository.dbConn.Where("user_name = ?", userName).First(&record).Error
	if err != nil {
		return users.Domain{}, err
	}
	return record.toDomain(), nil
}

func (repository *userRepository) GetUserByID(ctx context.Context, userId int) (users.Domain, error) {
	record := Users{}
	err := repository.dbConn.Where("id_user = ?", userId).First(&record).Error
	if err != nil {
		return users.Domain{}, err
	}
	return record.toDomain(), nil
}

func (repository *userRepository) DeleteUser(ctx context.Context, userId int, domain *users.Domain) error {
	record := fromDomainToRecord(domain)

	record.Id_user = userId
	err := repository.dbConn.Where("id_user = ?", userId).Delete(&record).Error
	if err != nil {
		return err
	}

	return nil
}
