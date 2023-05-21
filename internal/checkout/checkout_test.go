package checkout

import (
	"checkoutkata/internal/model"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestScan(t *testing.T) {
	checkout := NewStandardCheckout(setupItems())

	t.Run("should scan available item", func(t *testing.T) {
		err := checkout.Scan("A")
		require.NoError(t, err)
	})

	t.Run("should handle invalid item", func(t *testing.T) {
		err := checkout.Scan("Z")
		require.Error(t, err)
	})
}

func TestGetTotal(t *testing.T) {

	var checkout Checkout

	t.Run("should calculate items with no special price correctly", func(t *testing.T) {

		checkout = NewStandardCheckout(setupItems())

		checkout.Scan("C")
		checkout.Scan("C")
		checkout.Scan("C")

		total := checkout.GetTotalPrice()

		require.Equal(t, 60, total)
	})

	t.Run("should calculate items with special prices correctly", func(t *testing.T) {
		checkout = NewStandardCheckout(setupItems())
		checkout.Scan("A")
		checkout.Scan("A")
		checkout.Scan("A")

		total := checkout.GetTotalPrice()

		require.Equal(t, 0, total)
	})

}

func setupItems() []model.Item {
	itemA := model.Item{
		SKU:       "A",
		UnitPrice: 50,
		SpecialPrice: &model.SpecialPrice{
			AmountRequired: 3,
			Price:          130,
		},
	}

	itemB := model.Item{
		SKU:       "B",
		UnitPrice: 30,
		SpecialPrice: &model.SpecialPrice{
			AmountRequired: 2,
			Price:          45,
		},
	}

	itemC := model.Item{
		SKU:       "C",
		UnitPrice: 20,
	}

	itemD := model.Item{
		SKU:       "D",
		UnitPrice: 15,
	}

	return []model.Item{itemA, itemB, itemC, itemD}
}
