package repositories

import (
	"github.com/erenyusufduran/wasnon/internal/models"
	"gorm.io/gorm"
)

type EmployeeRepository interface {
	Create(product *models.Employee) error
	GetAll() ([]models.Employee, error)
	Update(product *models.Employee) error
}

type EmployeeRepositoryImpl struct {
	db *gorm.DB
}

func NewEmployeeRepositoryImpl(db *gorm.DB) *EmployeeRepositoryImpl {
	return &EmployeeRepositoryImpl{db: db}
}

func (r *EmployeeRepositoryImpl) Create(employee *models.Employee) error {
	return r.db.Create(employee).Error
}

func (r *EmployeeRepositoryImpl) GetAll() ([]models.Employee, error) {
	var employees []models.Employee
	err := r.db.Find(&employees).Error
	return employees, err
}

func (r *EmployeeRepositoryImpl) Update(employee *models.Employee) error {
	return r.db.Save(employee).Error
}
