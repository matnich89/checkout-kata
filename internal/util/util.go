package util

import (
	"github.com/matnich89/checkoutkata/internal/model"
)

func SetupItems() []model.Item {
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
