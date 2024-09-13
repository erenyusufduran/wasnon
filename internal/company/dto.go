package company

import "github.com/erenyusufduran/wasnon/internal/branch"

type CreateCompanyRequest struct {
	Name string `json:"name"`
}

type CompanyWithBranches struct {
	ID       uint            `json:"id"`
	Name     string          `json:"name"`
	Branches []branch.Branch `json:"branches"`
}
