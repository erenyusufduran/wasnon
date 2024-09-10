package company

import (
	"github.com/erenyusufduran/wasnon/internal/branch"
)

type Company struct {
	ID       uint
	Name     string          `json:"name"`
	Branches []branch.Branch `json:"branches"`
}

func NewCompanyWithName(name string) *Company {
	return &Company{
		Name: name,
	}
}

func NewCompany(name string, branches []branch.Branch) *Company {
	return &Company{
		Name:     name,
		Branches: branches,
	}
}
