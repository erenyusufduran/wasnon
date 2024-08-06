package repositories

import (
	"github.com/erenyusufduran/wasnon/internal/models"
	"gorm.io/gorm"
)

type GormEmployeeRepository struct {
	db *gorm.DB
}

func NewGormEmployeeRepository(db *gorm.DB) *GormEmployeeRepository {
	return &GormEmployeeRepository{db: db}
}

func (r *GormEmployeeRepository) Create(employee *models.Employee) error {
	return r.db.Create(employee).Error
}

func (r *GormEmployeeRepository) GetAll() ([]models.Employee, error) {
	var employees []models.Employee
	err := r.db.Find(&employees).Error
	return employees, err
}

func (r *GormEmployeeRepository) Update(employee *models.Employee) error {
	return r.db.Save(employee).Error
}
