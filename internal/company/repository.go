package company

import (
	"gorm.io/gorm"
)

type CompanyRepository interface {
	Create(product *Company) error
	GetAll() ([]Company, error)
}

type CompanyRepositoryImpl struct {
	db *gorm.DB
}

func NewCompanyRepositoryImpl(db *gorm.DB) *CompanyRepositoryImpl {
	return &CompanyRepositoryImpl{db: db}
}

func (r *CompanyRepositoryImpl) Create(company *Company) error {
	return r.db.Create(company).Error
}

func (r *CompanyRepositoryImpl) GetAll() ([]Company, error) {
	var companies []Company
	err := r.db.Preload("Employees").Preload("Products").Find(&companies).Error
	return companies, err
}
