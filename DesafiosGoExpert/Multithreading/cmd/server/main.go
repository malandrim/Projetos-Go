package main

import (
	"net/http"

	"github.com/malandrim/Projetos-Go/tree/main/DesafiosGoExpert/Multithreading/internal/infra/webserver/handlers"
)

func main() {

	http.HandleFunc("/", handlers.BuscaCepHandler)
	http.ListenAndServe(":8080", nil)

}
