package api

import (
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func SampleHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("sample request ok"))
	w.WriteHeader(200)
}

func ImageReceiveHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	if r.Method == "POST" {

		log.Print("src_file open")
		src_file, header, err := r.FormFile("image")
		if err != nil {
			log.Print("failed")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer src_file.Close()

		log.Print("dst_file open | create")
		dir, _ := os.Getwd()
		dst_file, err := os.OpenFile(dir+"/images/"+header.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			log.Print("failed")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer dst_file.Close()

		log.Print("file copy")
		_, err = io.Copy(dst_file, src_file)
		if err != nil {
			log.Print("failed")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(200)
		log.Print("successfull")
	}
}

func ImageSendHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	dir, _ := os.Getwd()

	// Get all file names inside a directory
	log.Print("read filenames")
	files, err := ioutil.ReadDir(dir + "/images")
	if err != nil {
		log.Print("failed")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("open imagefile")
	f, err := os.Open(dir + "/images/" + files[0].Name())
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer f.Close()

	log.Println("create formfile")
	writer := multipart.NewWriter(w)
	part, err := writer.CreateFormFile("file", filepath.Base(f.Name()))
	if err != nil{
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.Copy(part, f)
	writer.Close()

	w.Header().Add("Content-Type", writer.FormDataContentType())
	log.Println("successfull")

	// for _, file := range files {
	// 	log.Print(file.Name())
	// }

	// // Get all image data from a file name
	// var images []image.Image
	// log.Print("get all imagedata")
	// for _, file := range files {
	// 	f, err := os.Open(dir + "/" + string(file))
	// 	if err != nil{
	// 		log.Print("failed")
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	}
	// 	defer f.Close()
	// 	image, _, err := image.Decode(f)
	// }
}
