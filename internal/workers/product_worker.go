package workers

import (
	"fmt"
	"time"

	"github.com/erenyusufduran/wasnon/internal/repositories"
)

// disableExpiredProducts checks for expired products and updates their status
func disableExpiredProducts(repo repositories.ProductRepository) error {
	fmt.Println("Running expiration check...")

	products, err := repo.GetActiveExpiredProducts(time.Now(), 100)
	if err != nil {
		return err
	}

	err = repo.UpdateProductsStatus(products, "past")
	if err != nil {
		return err
	}

	return nil
}
