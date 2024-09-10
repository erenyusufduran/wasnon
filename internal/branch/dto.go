package branch

type CreateBranchRequest struct {
	Name      string `json:"name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	CompanyID uint   `json:"companyid" validate:"required"`
	Address   string `json:"address"`
	City      string `json:"city"`
	County    string `json:"county"`
}
