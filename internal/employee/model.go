package employee

import "github.com/erenyusufduran/wasnon/shared"

type Employee struct {
	shared.CustomModel
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
