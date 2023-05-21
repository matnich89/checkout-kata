package main

import (
	"log"

	"github.com/go-chi/chi/v5"

	"github.com/matnich89/checkoutkata/internal/checkout"
	"github.com/matnich89/checkoutkata/internal/util"
)

func main() {
	router := chi.NewRouter()

	checkout := checkout.NewStandardCheckout(util.SetupItems())

	app := newApp(router, checkout)

	log.Fatal(app.serve())
}
