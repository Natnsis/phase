package handlers

import "net/http"

var maxUploadSize = 100 << 20

func UploadFile(w http.ResponseWriter, r *http.Request) {
	// check method to be only post
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// cap total request body size
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxUploadSize))
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		http.Error(w, "file too large or bad form", http.StatusBadRequest)
		return
	}

	// file form field name
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "missing file", http.StatusBadRequest)
		return
	}
	defer file.Close()
}
