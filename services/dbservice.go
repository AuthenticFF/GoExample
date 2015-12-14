package services
import (
	"github.com/AuthenticFF/GoExample/models"
	
    "gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
)

var dbService IDBService
type IDBService interface {
	SaveResult(newResult models.Result) (models.Result, error)
}
type DBService struct {
	session *mgo.Session
}

func (s *DBService) SaveResult(newResult models.Result) (models.Result, error) {

    // Add an Id
    newResult.Id = bson.NewObjectId()
    s.session.DB("goExample").C("result").Insert(newResult)
	return newResult, nil
	
}