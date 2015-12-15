package controllers
import (	
	"github.com/AuthenticFF/GoExample/services"

	"net/http"
	"encoding/json"
	"log"
	"fmt"

	"github.com/julienschmidt/httprouter"
)

var Example exampleController

func Init(router *httprouter.Router) *httprouter.Router {
	Example = exampleController{services.Example, services.DB}
	router = Example.Init(router)
	log.Println("Controllers Initialized");
	return router
}


/*
helper funcs/structs
*/
type httpStatus struct {
	err    error
	status int
}
func ServerError(err error) httpStatus {
	return httpStatus{err, http.StatusInternalServerError}
}

func StatusOk(status int) httpStatus {
	return httpStatus{nil, status}
}
type controllerRoute func(http.ResponseWriter, *http.Request, httprouter.Params) (interface{}, httpStatus);

func PublicRoute(r controllerRoute) httprouter.Handle {
	return func(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
		result, status := r(w, req, p)
		WriteResponse(w, result, status)
	}
}

func WriteResponse(w http.ResponseWriter, result interface{}, httpStatus httpStatus) {
	var responseBody string
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Powered-By", "Authentic F&F")
	w.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
	if httpStatus.err != nil {
		w.WriteHeader(httpStatus.status)
		jsonBody, _ := json.Marshal(httpStatus.err.Error())
		responseBody = string(jsonBody)
	} else {
		w.WriteHeader(httpStatus.status)
		jsonBody, _ := json.Marshal(result)
		responseBody = string(jsonBody)
	}
	fmt.Fprintf(w, responseBody)
}