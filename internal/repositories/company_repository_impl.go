package repositories

import (
	"github.com/erenyusufduran/wasnon/internal/models"
	"gorm.io/gorm"
)

type GormCompanyRepository struct {
	db *gorm.DB
}

func NewGormCompanyRepository(db *gorm.DB) *GormCompanyRepository {
	return &GormCompanyRepository{db: db}
}

func (r *GormCompanyRepository) Create(company *models.Company) error {
	return r.db.Create(company).Error
}

func (r *GormCompanyRepository) GetAll() ([]models.Company, error) {
	var companies []models.Company
	err := r.db.Preload("Employees").Preload("Products").Find(&companies).Error
	return companies, err
}
