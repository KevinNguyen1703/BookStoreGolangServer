package routes

import (
	"github.com/KevinNguyen1703/BookStoreGolangServer/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/Data/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/Data/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/Data/{bookID}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/Data/{bookID}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/Data/{bookID}", controllers.DeleteBook).Methods("DELETE")
}
