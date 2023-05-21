package main

import (
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	mockcheckout "github.com/matnich89/checkoutkata/internal/checkout/mocks"
)

func TestScan(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mockcheckout.NewMockCheckout(ctrl)
	app := newTestApplication(t, m)
	server := newTestServer(t, app.routes())
	defer server.Close()

	t.Run("should return 200", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, server.URL+"/checkout/scan/A", nil)
		resp, err := server.Client().Do(req)
		require.NoError(t, err)
		require.Equal(t, resp.StatusCode, http.StatusOK)
	})

}
