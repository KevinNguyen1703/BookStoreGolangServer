package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/KevinNguyen1703/FrameStoreGolangServer/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterFrameStoreRoutes(r)
	http.Handle("/", r)
	fmt.Println("Server is starting")
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
