package model

type Item struct {
	SKU          string
	UnitPrice    int
	SpecialPrice *SpecialPrice
}

type SpecialPrice struct {
	AmountRequired int
	Price          int
}
