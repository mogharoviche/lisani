package main

import (
	"fmt"
	"lisani/router"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Mongo Api")
	r := router.Router()
	fmt.Println("server is getting started ...")
	log.Fatal(http.ListenAndServe(":8080", r))
	fmt.Println("Listening at port 8080 ...")
}
