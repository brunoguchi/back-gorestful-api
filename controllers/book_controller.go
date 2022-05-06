package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	restful "github.com/emicklei/go-restful"
)

type Book struct {
	Id          int
	Description string
}

func MockedData() []Book {
	books := []Book{
		{Id: 1, Description: "book 1"},
		{Id: 2, Description: "book 2"},
		{Id: 3, Description: "book 3"},
		{Id: 4, Description: "book 4"},
		{Id: 5, Description: "book 5"},
	}

	return books
}

func GetAllBooks(request *restful.Request, response *restful.Response) {
	response.Header().Set("Content-Type", "application/json")

	dados := MockedData()

	encoder := json.NewEncoder(response)
	encoder.Encode(dados)
}

func GetBookById(request *restful.Request, response *restful.Response) {
	idParam := request.PathParameter("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		response.WriteError(http.StatusBadRequest, err)
		return
	}

	books := MockedData()

	for _, book := range books {
		if book.Id == id {
			response.WriteAsJson(book)
			return
		}
	}

	response.WriteErrorString(http.StatusNotFound, "Book not found.")
}
