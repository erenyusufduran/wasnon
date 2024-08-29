package employee

import (
	"gorm.io/gorm"
)

type EmployeeRepository interface {
	Create(product *Employee) error
	GetAll() ([]Employee, error)
	Update(product *Employee) error
}

type EmployeeRepositoryImpl struct {
	db *gorm.DB
}

func NewEmployeeRepositoryImpl(db *gorm.DB) *EmployeeRepositoryImpl {
	return &EmployeeRepositoryImpl{db: db}
}

func (r *EmployeeRepositoryImpl) Create(employee *Employee) error {
	return r.db.Create(employee).Error
}

func (r *EmployeeRepositoryImpl) GetAll() ([]Employee, error) {
	var employees []Employee
	err := r.db.Find(&employees).Error
	return employees, err
}

func (r *EmployeeRepositoryImpl) Update(employee *Employee) error {
	return r.db.Save(employee).Error
}
