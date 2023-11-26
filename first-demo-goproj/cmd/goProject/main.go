package main

import (
	"fmt"
	"net/http"

	"github.com/pradeep/go_projs/internal/routes"
)

func main() {
	router := routes.NewRouter()

	port := 8080
	addr := fmt.Sprintf(":%d", port)

	fmt.Printf("Serving on port http://localhost:%d\n", port)
	err := http.ListenAndServe(addr, router)

	if err != nil {
		panic(err)
	}
}
