package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/matnich89/checkoutkata/internal/checkout"
)

type app struct {
	router   *chi.Mux
	checkout checkout.Checkout
}

func newApp(router *chi.Mux, checkout checkout.Checkout) *app {
	return &app{
		router:   router,
		checkout: checkout,
	}
}

func (a *app) routes() *chi.Mux {

	a.router.Use(middleware.Recoverer)
	a.router.Get("/checkout/scan/{itemSKU}", a.ScanItem)
	a.router.Get("/checkout/total", a.GetItemTotal)

	return a.router
}
