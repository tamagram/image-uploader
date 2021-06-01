package api

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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

type Image struct {
	Name string `json:"name"`
	Data string `json:"data"`
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

	// log.Println("create formfile")
	// writer := multipart.NewWriter(w)
	// part, err := writer.CreateFormFile("file", filepath.Base(f.Name()))
	// if err != nil{
	// 	log.Println(err.Error())
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	// io.Copy(part, f)
	// writer.Close()

	// w.Header().Add("Content-Type", writer.FormDataContentType())
	// log.Println("successfull")

	log.Println("generate json")
	var imgJson []Image
	var mImgJson []byte
	for _, file := range files {
		f, err := os.Open(dir + "/images/" + file.Name())
		if err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()

		binary, _ := ioutil.ReadAll(f)
		imgJson = append(imgJson, Image{
			Name: file.Name(),
			Data: base64.StdEncoding.EncodeToString(binary),
		})
	}

	log.Println("marshal")
	mImgJson, err = json.Marshal(imgJson)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(mImgJson)

	log.Println("successfull")
}
