package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (a *app) ScanItem(w http.ResponseWriter, r *http.Request) {
	itemSKU := chi.URLParam(r, "itemSKU")

	if itemSKU == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := a.checkout.Scan(itemSKU)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func (a *app) GetItemTotal(w http.ResponseWriter, r *http.Request) {

}
