package checkout

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestScan(t *testing.T) {
	checkout := NewStandardCheckout()

	err := checkout.Scan("A")

	require.NoError(t, err)
}

func TestGetTotal(t *testing.T) {
	checkout := NewStandardCheckout()

	total := checkout.GetTotalPrice()

	require.Equal(t, 0, total)
}