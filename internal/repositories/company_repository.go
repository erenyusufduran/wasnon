package repositories

import (
	"github.com/erenyusufduran/wasnon/internal/models"
	"gorm.io/gorm"
)

type CompanyRepository interface {
	Create(product *models.Company) error
	GetAll() ([]models.Company, error)
}

type CompanyRepositoryImpl struct {
	db *gorm.DB
}

func NewCompanyRepositoryImpl(db *gorm.DB) *CompanyRepositoryImpl {
	return &CompanyRepositoryImpl{db: db}
}

func (r *CompanyRepositoryImpl) Create(company *models.Company) error {
	return r.db.Create(company).Error
}

func (r *CompanyRepositoryImpl) GetAll() ([]models.Company, error) {
	var companies []models.Company
	err := r.db.Preload("Employees").Preload("Products").Find(&companies).Error
	return companies, err
}
