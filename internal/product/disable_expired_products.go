package product

import (
	"time"
)

// disableExpiredProducts checks for expired products and updates their status
func DisableExpiredProducts(repo ProductRepository) error {
	products, err := repo.GetActiveExpiredProducts(time.Now(), 100)
	if err != nil {
		return err
	}

	err = repo.UpdateProductsStatus(products, Past)
	if err != nil {
		return err
	}

	return nil
}
