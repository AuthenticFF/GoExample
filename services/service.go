package services

import (
	"github.com/AuthenticFF/GoExample/db"
	
	"log"
)

var DB DBService
var Example ExampleService

func init() {
	Example = ExampleService{}
	DB = DBService{db.Session}
	log.Printf("Services Initialized");
}
