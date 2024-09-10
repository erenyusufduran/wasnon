package main

import (
	"github.com/erenyusufduran/wasnon/internal/branch"
	"github.com/erenyusufduran/wasnon/internal/company"
	"github.com/erenyusufduran/wasnon/internal/product"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repositories struct {
	CompanyRepository company.CompanyRepository
	ProductRepository product.ProductRepository
	BranchRepository  branch.BranchRepository
}

func NewRepositories(db *pgxpool.Pool) *Repositories {
	return &Repositories{
		CompanyRepository: company.NewCompanyRepositoryImpl(db),
		ProductRepository: product.NewProductRepositoryImpl(db),
		BranchRepository:  branch.NewBranchRepositoryImpl(db),
	}
}
