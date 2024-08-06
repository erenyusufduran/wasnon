package repositories

import "github.com/erenyusufduran/wasnon/internal/models"

type CompanyRepository interface {
	Create(product *models.Company) error
	GetAll() ([]models.Company, error)
}
