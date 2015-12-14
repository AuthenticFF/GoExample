package services

import (
	"github.com/AuthenticFF/GoExample/db"
	
	"log"
)

var Example ExampleService
var DB DBService

func init() {
	Example = ExampleService{}
	DB = DBService{db.Session}
	log.Printf("Services Initialized");
}
