package checkout

import (
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

	t.Run("should scan multiple available items", func(t *testing.T) {

		checkout = NewStandardCheckout(setupItems())

		err := scanItemNoOfTimes(checkout, "A", 2)
		require.NoError(t, err)

		err = checkout.Scan("B")
		require.NoError(t, err)

		scannedAItems := checkout.currentScannedItems["A"]
		scannedBItems := checkout.currentScannedItems["B"]

		require.Equal(t, scannedAItems, 2)
		require.Equal(t, scannedBItems, 1)
	})

	t.Run("should scan multiple available items handle any order", func(t *testing.T) {
		checkout = NewStandardCheckout(setupItems())

		err := checkout.Scan("B")
		require.NoError(t, err)

		err = checkout.Scan("A")
		require.NoError(t, err)

		err = checkout.Scan("B")
		require.NoError(t, err)

		err = checkout.Scan("C")
		require.NoError(t, err)

		err = checkout.Scan("A")
		require.NoError(t, err)

		scannedAItems := checkout.currentScannedItems["A"]
		scannedBItems := checkout.currentScannedItems["B"]
		scannedCItems := checkout.currentScannedItems["C"]

		require.Equal(t, scannedAItems, 2)
		require.Equal(t, scannedBItems, 2)
		require.Equal(t, scannedCItems, 1)

	})

}

func TestGetTotal(t *testing.T) {

	var checkout Checkout

	t.Run("should calculate items with no special price correctly", func(t *testing.T) {

		checkout = NewStandardCheckout(setupItems())
		err := scanItemNoOfTimes(checkout, "C", 3)
		require.NoError(t, err)

		total := checkout.GetTotalPrice()

		require.Equal(t, 60, total)
	})

	t.Run("should calculate items with special prices correctly", func(t *testing.T) {
		checkout = NewStandardCheckout(setupItems())

		err := scanItemNoOfTimes(checkout, "A", 3)
		require.NoError(t, err)

		total := checkout.GetTotalPrice()

		require.Equal(t, 130, total)
	})

	t.Run("should calculate items with special prices correctly with items remaining", func(t *testing.T) {
		checkout = NewStandardCheckout(setupItems())

		err := scanItemNoOfTimes(checkout, "A", 4)
		require.NoError(t, err)

		total := checkout.GetTotalPrice()

		require.Equal(t, 180, total)

	})

	t.Run("should calculate items with special prices correctly with items remaining and with multi special price applied", func(t *testing.T) {
		checkout = NewStandardCheckout(setupItems())

		err := scanItemNoOfTimes(checkout, "A", 7)
		require.NoError(t, err)

		total := checkout.GetTotalPrice()

		require.Equal(t, 310, total)
	})

	t.Run("should calculate items with special prices correctly with items remaining and with multi special price applied", func(t *testing.T) {
		checkout = NewStandardCheckout(setupItems())

		err := scanItemNoOfTimes(checkout, "A", 7)
		require.NoError(t, err)

		err = scanItemNoOfTimes(checkout, "B", 2)
		require.NoError(t, err)

		total := checkout.GetTotalPrice()

		require.Equal(t, 355, total)
	})
}
