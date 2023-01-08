package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/CeoFred/go-book-api-mysqldb/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r) // register routes
	http.Handle("/",r)
	log.Fatal(http.ListenAndServe(":8080", r))
}
