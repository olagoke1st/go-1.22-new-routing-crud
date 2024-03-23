package main

import (
	"fmt"
	"net/http"
)

func main() {
	app := http.NewServeMux()

	// before go-1.22, you will have to write conditional statements for the routes
	app.HandleFunc("POST /book")
	app.HandleFunc("GET /books")
	app.HandleFunc("GET /book/{id}")
	app.HandleFunc("PUT /book/{id}")
	app.HandleFunc("DELETE /book/{id}")

	port := "7000"

	err := http.ListenAndServe(":"+port, app)
	if err != nil {
		panic(err)
	}
	fmt.Println("App listening on localhost:" + port)
}
