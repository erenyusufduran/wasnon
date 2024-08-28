package repositories

import (
	"gorm.io/gorm"
)

func New(db *gorm.DB) *Repositories {
	return &Repositories{
		EmployeeRepository: NewEmployeeRepositoryImpl(db),
		CompanyRepository:  NewCompanyRepositoryImpl(db),
		ProductRepository:  NewProductRepositoryImpl(db),
	}
}

type Repositories struct {
	EmployeeRepository EmployeeRepository
	CompanyRepository  CompanyRepository
	ProductRepository  ProductRepository
}
