package controllers

import (
	"github.com/AuthenticFF/GoExample/models"
	"github.com/AuthenticFF/GoExample/services"

	"net/http"

	"github.com/julienschmidt/httprouter"
)
type exampleController struct {
	exampleService services.ExampleService
	dbService services.DBService
}

func (c *exampleController) Init(router *httprouter.Router) *httprouter.Router {
	router.GET("/api/example", PublicRoute(Example.Analyze))

	return router

}
/*
Main API method
*/
func (c *exampleController) Analyze(writer http.ResponseWriter, req *http.Request, params httprouter.Params) (interface{}, httpStatus) {
	var result models.Result;
	var err error;

	result, err = c.exampleService.GetData(result)
	if err != nil {
		return nil, ServerError(err)
	}

	result, err = c.dbService.SaveResult(result)
	if err != nil {
		return nil, ServerError(err)
	}


	return result , StatusOk(http.StatusOK)
}
