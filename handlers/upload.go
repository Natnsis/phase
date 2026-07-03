package handlers

import "net/http"

const (
	maxSize   = 100 << 20
	uploadDir = "~/Desktop/uploads"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	// only post request method
	if r.Method != http.MethodPost {
		http.Error(w, "method must be post", http.StatusInternalServerError)
		return
	}

	// control request size
	r.Body = http.MaxBytesReader(w, r.Body, maxSize)

	// parse multipart headers into memory (32kb)
	if err := r.ParseMultipartForm(32 << 10); err != nil {
		http.Error(w, "file to large malformed request", http.StatusBadRequest)
		return
	}

	// check the files existance
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "no file found", http.StatusBadRequest)
		return
	}
	defer file.Close()

	savedPath, err := saveUploadedFile(file, header.Filename)
}
