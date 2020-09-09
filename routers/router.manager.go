package routers

import (
	"net/http"

	"github.com/davidmm07/ordenes_api/handlers"
)

func InitRouteHandlers() {
	fs := http.FileServer(http.Dir("static"))
	http.HandleFunc("/ordenes", handlers.GraphqlHandler)
	http.Handle("/", fs)
}
