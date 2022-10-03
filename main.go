package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/an1l4/stockAPI/router"
)

func main() {
	route := router.Router()

	fmt.Println("server running at 8080...")
	log.Fatal(http.ListenAndServe(":8080", route))
}
