package company

import (
	"context"

	"github.com/erenyusufduran/wasnon/internal/branch"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CompanyRepository interface {
	Create(product *Company) error
	GetAll() ([]Company, error)
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

	companyMap := make(map[uint]*Company)
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
			companyMap[companyID] = &Company{
				ID:       companyID,
				Name:     nullableString(companyName),
				Branches: []branch.Branch{},
			}
		}

		// Add branch to the company if the branch ID is valid (not null)
		if branchID != nil {
			companyMap[companyID].Branches = append(companyMap[companyID].Branches, branch.Branch{
				ID:        *branchID,
				CompanyID: companyID,
				Name:      nullableString(name),
				Email:     nullableString(email),
				Address:   nullableString(address),
				City:      nullableString(city),
				County:    nullableString(county),
			})
		}
	}

	// Convert the map to a slice of companies
	var companies []Company
	for _, company := range companyMap {
		companies = append(companies, *company)
	}

	return companies, nil
}

func nullableString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
