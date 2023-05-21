package checkout

import (
	"checkoutkata/internal/model"
	"fmt"
)

type Checkout interface {
	Scan(itemSku string) error
	GetTotalPrice() int
}

type StandardCheckout struct {
	currentScannedItems map[string]int
	availableItems      map[string]*model.Item
}

func NewStandardCheckout(initialItems []model.Item) *StandardCheckout {

	availableItems := map[string]*model.Item{}

	for _, it := range initialItems {
		availableItems[it.SKU] = &it
	}

	return &StandardCheckout{
		currentScannedItems: map[string]int{},
		availableItems:      availableItems,
	}
}

func (c *StandardCheckout) Scan(itemSku string) error {

	if _, ok := c.availableItems[itemSku]; ok {
		c.currentScannedItems[itemSku] = c.currentScannedItems[itemSku] + 1
		return nil
	}

	return fmt.Errorf("invalid item SKU: %s", itemSku)
}

func (c *StandardCheckout) GetTotalPrice() int {
	return 0
}
