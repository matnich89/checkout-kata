package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/golang/mock/gomock"

	"github.com/matnich89/checkoutkata/internal/checkout"
)

type testServer struct {
	*httptest.Server
}

func newTestApplication(t *testing.T, checkout checkout.Checkout) *app {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	router := chi.NewRouter()
	app := newApp(router, checkout)
	return app
}

func newTestServer(t *testing.T, h http.Handler) *testServer {
	server := httptest.NewServer(h)
	return &testServer{server}
}
