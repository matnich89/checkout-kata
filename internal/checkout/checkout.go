package checkout

import "checkoutkata/internal/model"

type Checkout interface {
	Scan(itemSku string) error
	GetTotalPrice() int
}

type StandardCheckout struct {
	currentScannedItems []model.Item
	availableItems      map[string]model.Item
}

func NewStandardCheckout() *StandardCheckout {
	return &StandardCheckout{}
}

func (c *StandardCheckout) Scan(itemSku string) error {
	return nil
}

func (c *StandardCheckout) GetTotalPrice() int {
	return 0
}
