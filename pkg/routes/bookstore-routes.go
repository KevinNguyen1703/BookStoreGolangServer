package routes

import (
	"github.com/KevinNguyen1703/FrameStoreGolangServer/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterFrameStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/Data/", controllers.CreateFrame).Methods("POST")
	router.HandleFunc("/Data/", controllers.GetFrame).Methods("GET")
	router.HandleFunc("/Data/{frameID}", controllers.GetFrameById).Methods("GET")
	router.HandleFunc("/Data/{frameID}", controllers.UpdateFrame).Methods("PUT")
	router.HandleFunc("/Data/{frameID}", controllers.DeleteFrame).Methods("DELETE")
	router.HandleFunc("/DeleteAll/", controllers.DeleteAll).Methods("DELETE")
}
