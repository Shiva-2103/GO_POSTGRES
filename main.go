package main

import (
	"fmt"
	"log"
	"net/http"
	"postgres/router"
)

func main() {
	r := router.Router()
	fmt.Println("Starting Server on the Port 3000...")

	log.Fatal(http.ListenAndServe(":3000", r))
}
