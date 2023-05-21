package checkout

import (
	"fmt"
	"github.com/matnich89/checkoutkata/internal/model"
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

	for i := 0; i < len(initialItems); i++ {
		item := initialItems[i]
		availableItems[item.SKU] = &item
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
	var total int

	for sku, amount := range c.currentScannedItems {
		item := c.availableItems[sku]
		total += calculateTotalForItem(item, amount)
	}

	return total
}

func calculateTotalForItem(item *model.Item, amount int) int {

	var itemTotal int

	if item.SpecialPrice != nil {
		numberOfSpecialPrices := amount / item.SpecialPrice.AmountRequired
		remainingItems := amount % item.SpecialPrice.AmountRequired
		itemTotal += item.SpecialPrice.Price * numberOfSpecialPrices
		itemTotal += remainingItems * item.UnitPrice
	} else {
		itemTotal = item.UnitPrice * amount
	}

	return itemTotal
}
