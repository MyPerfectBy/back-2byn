package main

import (
	"./db"
	"./handler"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
)


func main() {

	db.GetDataBaseInstance()

	//db.CreateTable(model.Product{})

	r := mux.NewRouter()
	r.HandleFunc("/", handler.InfoHandler).Methods("GET")

	//products
	r.HandleFunc("/api/products/new", handler.NewProductHandler).Methods("POST")
	r.HandleFunc("/api/products/approve", handler.ApproveProduct).Methods("POST")
	r.HandleFunc("/api/products/delete", handler.DeleteProduct).Methods("POST")
	r.HandleFunc("/api/products/not-approved", handler.GetNotApprovedProducts).Methods("GET")
	r.HandleFunc("/api/products/approved", handler.GetApprovedProducts).Methods("GET")
	r.HandleFunc("/api/products/{id}", handler.GetProductByID).Methods("GET")

	//photos
	r.HandleFunc("/photos/{id}", handler.RenderPhotoHandler).Methods("GET")
	r.HandleFunc("/api/uploadPhoto", handler.FileUploadHandler).Methods("POST")

	http.ListenAndServe(":7755", handlers.CORS()(r))
}