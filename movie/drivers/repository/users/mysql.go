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
	record := fromDomainToRecord(*domain)

	result := repository.dbConn.Create(&record)
	if result.Error != nil {
		return users.Domain{}, result.Error
	}

	return record.toDomain(), nil
}

func (repository *userRepository) UpdateUser(ctx context.Context, domain *users.Domain) (users.Domain, error) {
	record := fromDomainToRecord(users.Domain{})

	result := repository.dbConn.Model(&record).Updates(&record)
	err := result.Error
	if err != nil {
		return users.Domain{}, err
	}
	return record.toDomain(), nil
}

func (repository *userRepository) GetUserByUsername(ctx context.Context, username string) (users.Domain, error) {
	record := Users{}
	err := repository.dbConn.Where("username = ?", username).First(&record).Error
	if err != nil {
		return users.Domain{}, err
	}
	return record.toDomain(), nil
}

func (repository *userRepository) GetUserByID(ctx context.Context, id int) (users.Domain, error) {
	record := Users{}
	err := repository.dbConn.Where("id = ?", id).First(&record).Error
	if err != nil {
		return users.Domain{}, err
	}
	return record.toDomain(), nil
}
