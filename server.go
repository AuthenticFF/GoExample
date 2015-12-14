package main
import (
	"github.com/AuthenticFF/GoExample/controllers"
	"github.com/AuthenticFF/GoExample/db"

	"net/http"
	"os"
	
	"github.com/julienschmidt/httprouter"
)
func main() {
	router := httprouter.New()
	defer db.Session.Close();

	router = controllers.Init(router)
	port := os.Getenv("PORT")
	if port == "" {
		port = "9091"
	}
	http.ListenAndServe(":"+port, router)
}