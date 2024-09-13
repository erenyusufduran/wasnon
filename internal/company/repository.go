package company

import (
	"context"

	"github.com/erenyusufduran/wasnon/internal/branch"
	"github.com/erenyusufduran/wasnon/shared/stringpkg"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CompanyRepository interface {
	Create(product *Company) error
	GetAll() ([]Company, error)
	GetManyWithBranches() ([]CompanyWithBranches, error)
}

type CompanyRepositoryImpl struct {
	db *pgxpool.Pool
}

func NewCompanyRepositoryImpl(db *pgxpool.Pool) *CompanyRepositoryImpl {
	return &CompanyRepositoryImpl{db: db}
}

func (r *CompanyRepositoryImpl) Create(company *Company) error {
	query := `INSERT INTO companies (name) VALUES ($1)`
	_, err := r.db.Exec(context.Background(), query, company.Name)
	if err != nil {
		return err
	}

	return nil
}

func (r *CompanyRepositoryImpl) GetAll() ([]Company, error) {
	query := `SELECT * FROM companies`
	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[Company])

	// var branches []Branch
	// for rows.Next() {
	// 	var branch Branch
	// 	err := rows.Scan(&branch.ID, &branch.Name, &branch.Email, &branch.CompanyID)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	branches = append(branches, branch)
	// }
	// return branches, nil
}

func (r *CompanyRepositoryImpl) GetManyWithBranches() ([]CompanyWithBranches, error) {
	query := `
		SELECT 
			c.id AS company_id, 
			c.name AS company_name, 
			b.id AS branch_id,
			b.name AS name,
			b.address AS address,
			b.city AS city,
			b.county AS county,
			b.email AS email
		FROM companies c
		LEFT JOIN branches b ON b.company_id = c.id
		ORDER BY c.id;
	`

	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	companyMap := make(map[uint]*CompanyWithBranches)
	for rows.Next() {
		var companyID uint
		var branchID *uint
		var companyName, name, address, city, county, email *string

		err := rows.Scan(
			&companyID, &companyName,
			&branchID, &name, &address,
			&city, &county, &email,
		)
		if err != nil {
			return nil, err
		}

		// If the company is not yet added, create it
		if _, exists := companyMap[companyID]; !exists {
			companyMap[companyID] = &CompanyWithBranches{
				ID:       companyID,
				Name:     stringpkg.NullableString(companyName),
				Branches: []branch.Branch{},
			}
		}

		// Add branch to the company if the branch ID is valid (not null)
		if branchID != nil {
			companyMap[companyID].Branches = append(companyMap[companyID].Branches, branch.Branch{
				ID:        *branchID,
				CompanyID: companyID,
				Name:      stringpkg.NullableString(name),
				Email:     stringpkg.NullableString(email),
				Address:   stringpkg.NullableString(address),
				City:      stringpkg.NullableString(city),
				County:    stringpkg.NullableString(county),
			})
		}
	}

	// Convert the map to a slice of companies
	var companies []CompanyWithBranches
	for _, company := range companyMap {
		companies = append(companies, *company)
	}

	return companies, nil
}
