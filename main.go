package main

import (
	"fmt"
	"net/http"
	"os"

	dbConnManager "github.com/davidmm07/ordenes_api/database"
	routerManager "github.com/davidmm07/ordenes_api/routers"
)

func main() {
	dbConnManager.InitDB()
	dbConnManager.RunMigrations()
	routerManager.InitRouteHandlers()
	fmt.Println("graphql api server is started at: http://localhost:" + os.Getenv("PORT") + "/ordenes")
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
