package main

import (
	"log"
	"net/http"

	"com.zoyu/user/router"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", router.Router()))
}
