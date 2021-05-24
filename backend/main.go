package main

import (
	"io"
	"net/http"
	"os"
)

func sampleHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("sample request ok"))
	w.WriteHeader(200)
}

func imageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Methods", "POST")

	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	m := r.MultipartForm
	for _, fileHeaders := range m.File {
		for _, header := range fileHeaders {
			src_file, err := header.Open()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			defer src_file.Close()

			// Open and copy the file
			dst_file, err := os.OpenFile("../path_to_save_image/"+header.Filename, os.O_RDONLY|os.O_CREATE, 0666)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			defer dst_file.Close()
			io.Copy(dst_file, src_file)
		}
	}

	// src_file, reader, err := r.FormFile("file")
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	// defer src_file.Close()

	// response
	w.Write([]byte("success"))
	w.WriteHeader(200)
}

func main() {
	http.HandleFunc("/file", imageHandler)
	http.HandleFunc("/sample", sampleHandler)
	http.ListenAndServe(":5000", nil)
}
