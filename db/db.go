package db
import (
	"log"
	"os"

	"gopkg.in/mgo.v2"
		)
var Session *mgo.Session
var mongohosts = "goExample_database_1"

func init() {
	log.Println("Datastore Initialized");
	var err error
	Session, err = mgo.Dial(mongohosts)
	if err != nil {
		log.Fatalf("Error connecting to the database: %s\n", err.Error())
	}
	Session.SetMode(mgo.Monotonic, false)
    Session.DB(os.Getenv("DB_NAME"))
}