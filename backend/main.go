package main

import (
	"io"
	"log"
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
	if r.Method == "POST" {
		// log.Print("ParseMultipartForm")
		// err := r.ParseMultipartForm(32 << 20)
		// if err != nil {
		// 	log.Print("failed...")
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	return
		// }

		// dump, _ := httputil.DumpRequest(r, true)
		// fmt.Println(string(dump))

		log.Print("src_file open")
		src_file, header, err := r.FormFile("image")
		if err != nil {
			log.Print("failed")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer src_file.Close()

		log.Print("dst_file open | create")
		dst_file, err := os.OpenFile("./images/"+header.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil{
			log.Print("failed")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer dst_file.Close()

		log.Print("file copy")
		_, err = io.Copy(dst_file, src_file)
		if err != nil{
			log.Print("failed")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		
		log.Print("successfull")


		// m := r.MultipartForm
		// for _, fileHeaders := range m.File {
		// 	for _, header := range fileHeaders {
		// 		src_file, err := header.Open()
		// 		log.Print("header.Open")
		// 		fmt.Print("fmt")
		// 		if err != nil {
		// 			log.Print("failed")
		// 			http.Error(w, err.Error(), http.StatusInternalServerError)
		// 		}
		// 		defer src_file.Close()

		// 		// Open and copy the file
		// 		dst_file, err := os.OpenFile("./images/"+header.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		// 		log.Print("os.OpenFile")
		// 		if err != nil {
		// 			log.Print("failed")
		// 			http.Error(w, err.Error(), http.StatusInternalServerError)
		// 			return
		// 		}
		// 		defer dst_file.Close()
		// 		io.Copy(dst_file, src_file)
		// 	}
		// }
	}
}

func main() {
	log.Print("API Open")
	http.HandleFunc("/file", imageHandler)
	http.HandleFunc("/sample", sampleHandler)
	http.ListenAndServe(":5000", nil)
}
