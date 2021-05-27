package main

import (
	"log"
	"net/http"

	"github.com/tamagram/image-uploader/api"
)

func main() {
	log.Print("API Open")
	http.HandleFunc("/file", api.ImageHandler)
	http.HandleFunc("/sample", api.SampleHandler)
	http.ListenAndServe(":5000", nil)
}
