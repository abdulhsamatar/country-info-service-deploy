package handlers

import "net/http"

func InfoHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Info endpoint coming next"))
}
