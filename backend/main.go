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
	if r.Method == "POST" {
		src_file, reader, err := r.FormFile("imageFile")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer src_file.Close()

		// Open and copy the file
		dst_file, err := os.OpenFile("./path_to_save_image/"+reader.Filename, os.O_RDONLY|os.O_CREATE, 0666)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer dst_file.Close()
		io.Copy(dst_file, src_file)

		// response
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Header", "Content-Type")
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("success"))
		w.WriteHeader(200)
	}
	// response
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("error"))
	w.WriteHeader(400)
}

func main() {
	http.HandleFunc("/file", imageHandler)
	http.HandleFunc("/sample", sampleHandler)
	http.ListenAndServe(":5000", nil)
}
