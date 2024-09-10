package branch

type Branch struct {
	ID        uint    `json:"id"`
	CompanyID uint    `json:"company_id"`
	Name      string  `json:"name"`
	Address   string  `json:"address"`
	City      string  `json:"city"`
	County    string  `json:"county"`
	Email     string  `json:"email"`
	Lat       float64 `json:"latitude" db:"-"`
	Lng       float64 `json:"longitude" db:"-"`
}

func NewBranch(companyID uint, name, address, city, county, email string) *Branch {
	return &Branch{
		CompanyID: companyID,
		Name:      name,
		Address:   address,
		City:      city,
		County:    county,
		Email:     email,
	}
}
