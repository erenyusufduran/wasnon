package company

import (
	"github.com/erenyusufduran/wasnon/internal/branch"
)

type Company struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func NewCompanyWithName(name string) *Company {
	return &Company{
		Name: name,
	}
}

func NewCompany(name string, branches []branch.Branch) *Company {
	return &Company{
		Name: name,
	}
}
