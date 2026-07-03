package handlers

import "net/http"

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method must be post", http.StatusInternalServerError)
	}
}
