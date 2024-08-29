package company

import (
	"github.com/erenyusufduran/wasnon/internal/employee"
	"github.com/erenyusufduran/wasnon/internal/product"

	"gorm.io/gorm"
)

type Company struct {
	gorm.Model
	Name      string              `json:"name"`
	Employees []employee.Employee `json:"employees"`
	Products  []product.Product   `json:"products"`
}

func NewCompanyWithName(name string) *Company {
	return &Company{
		Name: name,
	}
}

func NewCompany(name string, employees []employee.Employee, products []product.Product) *Company {
	return &Company{
		Name:      name,
		Employees: employees,
		Products:  products,
	}
}
