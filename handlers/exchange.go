package handlers

import "net/http"

func ExchangeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Exchange endpoint"))
}
