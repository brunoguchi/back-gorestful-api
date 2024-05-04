package server

import (
	"log"
	"net/http"

	"github.com/brunoguchi/back-gorestful-api/routes"

	restful "github.com/emicklei/go-restful"
)

func ServerRun() {
	ws := new(restful.WebService)
	ws.Route(routes.GetAllBooksRoute(ws))
	ws.Route(routes.GetBookByIdRoute(ws))
	restful.Add(ws)

	log.Print("Server running at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
