package employee

import (
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	Name      string `json:"name"`
	Email     string `json:"email"`
	CompanyID uint   `json:"company_id"`
}

func NewEmployee(name, email string, companyID uint) *Employee {
	return &Employee{
		Name:      name,
		Email:     email,
		CompanyID: companyID,
	}
}
