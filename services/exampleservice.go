package services
import (
    "github.com/AuthenticFF/GoExample/models"
)

var exampleService IExampleService
type IExampleService interface {
    GetData() (models.Result, error)
}
type ExampleService struct {}

func (s *ExampleService) GetData(result models.Result) (models.Result, error){
	/* do something neat */
	result.Data = "Lorem Ipsum"
    return result, nil 
}