package main

import (
	"github.com/erenyusufduran/wasnon/internal/company"
	"github.com/erenyusufduran/wasnon/internal/employee"
	"github.com/erenyusufduran/wasnon/internal/product"
	"gorm.io/gorm"
)

type Repositories struct {
	EmployeeRepository employee.EmployeeRepository
	CompanyRepository  company.CompanyRepository
	ProductRepository  product.ProductRepository
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		EmployeeRepository: employee.NewEmployeeRepositoryImpl(db),
		CompanyRepository:  company.NewCompanyRepositoryImpl(db),
		ProductRepository:  product.NewProductRepositoryImpl(db),
	}
}
