package branch

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type BranchRepository interface {
	Create(branch *Branch) error
	GetAll() ([]Branch, error)
	GetOneById(id uint64) (*Branch, error)
}

type BranchRepositoryImpl struct {
	db *pgxpool.Pool
}

func NewBranchRepositoryImpl(db *pgxpool.Pool) *BranchRepositoryImpl {
	return &BranchRepositoryImpl{db: db}
}

func (r *BranchRepositoryImpl) Create(branch *Branch) error {
	query := `INSERT INTO branches (name, address, city, county, company_id, email) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.db.Exec(context.Background(), query, branch.Name, branch.Address, branch.City, branch.County, branch.CompanyID, branch.Email)
	if err != nil {
		return err
	}
	return nil
}

func (r *BranchRepositoryImpl) GetAll() ([]Branch, error) {
	query := `SELECT * FROM branches`
	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[Branch])

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

func (r *BranchRepositoryImpl) GetOneById(id uint64) (*Branch, error) {
	query := `SELECT id, name, address, city, county, email, company_id FROM branches WHERE id = $1`

	rows, err := r.db.Query(context.Background(), query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	branch, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[Branch])
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("branch with id %d not found", id)
		}
		return nil, err
	}

	return &branch, nil
}
