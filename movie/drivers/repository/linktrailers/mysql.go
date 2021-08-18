package linktrailers

import (
	"context"
	"movie/businesses/linktrailers"

	"gorm.io/gorm"
)

type linkRepository struct {
	dbConn *gorm.DB
}

func NewLinkRepository(db *gorm.DB) linktrailers.Repository {
	return &linkRepository{
		dbConn: db,
	}
}

func (repository *linkRepository) CreateLinkRepository(ctx context.Context, domain *linktrailers.Domain) (linktrailers.Domain, error) {
	record := fromDomainToRecord(domain)
	result := repository.dbConn.Create(&record)
	if result.Error != nil {
		return linktrailers.Domain{}, result.Error
	}
	return record.toDomain(), nil
}

func (repository *linkRepository) UpdateLinkRepository(ctx context.Context, idLink int, domain *linktrailers.Domain) error {
	record := fromDomainToRecord(domain)
	record.Id_trailer = idLink
	err := repository.dbConn.Where("id_trailer = ?", idLink).Updates(&record).Error
	if err != nil {
		return err
	}
	return nil
}

func (repository *linkRepository) DeleteLinkRepository(ctx context.Context, idLink int, domain *linktrailers.Domain) error {
	record := fromDomainToRecord(domain)
	record.Id_trailer = idLink
	err := repository.dbConn.Where("id_trailer = ?", idLink).Delete(&record).Error
	if err != nil {
		return err
	}
	return nil
}

func (repository *linkRepository) GetAllLinkRepository(ctx context.Context, domain *linktrailers.Domain) (linktrailers.Domain, error) {
	record := LinkUrl{}
	err := repository.dbConn.First(&record).Error
	if err != nil {
		return linktrailers.Domain{}, err
	}
	return record.toDomain(), nil
}

func (repository *linkRepository) GetLinkbyId(ctx context.Context, idLink int) (linktrailers.Domain, error) {
	record := LinkUrl{}
	err := repository.dbConn.Where("id_trailer = ?", idLink).First(&record).Error
	if err != nil {
		return linktrailers.Domain{}, err
	}
	return record.toDomain(), nil
}
