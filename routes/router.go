package routes

import (
	"github.com/brunoguchi/back-gorestful-api/controllers"
	restful "github.com/emicklei/go-restful"
)

func GetAllBooksRoute(ws *restful.WebService) *restful.RouteBuilder {
	return ws.GET("/books").To(controllers.GetAllBooks)
}

func GetBookByIdRoute(ws *restful.WebService) *restful.RouteBuilder {
	return ws.GET("/book/{id}").To(controllers.GetBookById)
}
