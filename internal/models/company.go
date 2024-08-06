package models

import (
	"gorm.io/gorm"
)

type Company struct {
	gorm.Model
	Name      string     `json:"name"`
	Employees []Employee `json:"employees"`
	Products  []Product  `json:"products"`
}

func NewCompanyWithName(name string) *Company {
	return &Company{
		Name: name,
	}
}

func NewCompany(name string, employees []Employee, products []Product) *Company {
	return &Company{
		Name:      name,
		Employees: employees,
		Products:  products,
	}
}
