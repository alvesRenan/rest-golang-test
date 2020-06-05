package main

import (
	"log"
	"net/http"

	"github.com/alvesRenan/rest-golang-test/conf"
	"github.com/alvesRenan/rest-golang-test/routes"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	conf.InitializeDB()
	log.Println("Server started on port 8000")
	log.Fatal(http.ListenAndServe(":8000", routes.NewRouter()))
}
