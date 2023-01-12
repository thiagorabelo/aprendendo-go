package repository

import (
	"shortfy/models"

	"gorm.io/gorm"
)

type ShortURIRepository struct {
	db *gorm.DB
}

func NewShortURIRepository(db *gorm.DB) ShortURIRepository {
	return ShortURIRepository{db}
}

func (repo *ShortURIRepository) ListAll() ([]models.ShortURI, error) {
	var uris []models.ShortURI
	result := repo.db.Find(&uris)
	if result.Error != nil {
		return nil, result.Error
	}

	return uris, nil
}
