package repositories

import "github.com/erenyusufduran/wasnon/internal/models"

type EmployeeRepository interface {
	Create(product *models.Employee) error
	GetAll() ([]models.Employee, error)
	Update(product *models.Employee) error
}
