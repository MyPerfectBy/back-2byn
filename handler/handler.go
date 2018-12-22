package handler

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"os"
	"strings"
)

type PhotoUploadResponse struct {
	Success bool `json:"success"`
	URL string `json:"url"`
}

func responseJSON(data interface{}, response http.ResponseWriter, request *http.Request) {
	js, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}

	response.Header().Set("Content-Type", "application/json")
	response.Write(js)
}

func randPhotoID() string {
	b := make([]byte, 4)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func ClearString(str string) string {
	str = str[1 : len(str)-1]
	return str
}


var InfoHandler = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Doroy pidor, server works")
})

var FileUploadHandler = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	var req PhotoUploadResponse

	request.ParseMultipartForm(32 << 20)
	file, handler, err := request.FormFile("uploadfile")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	splitStrings := strings.Split(handler.Filename, ".")
	extension := splitStrings[len(splitStrings) - 1]
	photoURL := randPhotoID()
	req.URL = photoURL + "." + extension
	f, err := os.OpenFile("./photos/" + req.URL, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)
	req.Success = true

	responseJSON(req, response, request)
})

var RenderPhotoHandler = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {

	id := strings.ToLower(mux.Vars(request)["id"])
	http.ServeFile(response, request, "./photos/" + id)
})