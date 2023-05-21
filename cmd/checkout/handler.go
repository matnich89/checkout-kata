package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/matnich89/checkoutkata/internal/model"
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

func (a *app) GetTotal(w http.ResponseWriter, r *http.Request) {
	totalPrice := a.checkout.GetTotalPrice()

	response := model.TotalResponse{Total: totalPrice}

	b, err := json.Marshal(response)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = w.Write(b)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

}
