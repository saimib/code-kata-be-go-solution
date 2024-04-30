package todoist

import (
	"fmt"

	"github.com/saimib/code-kata-be-go-solution/models"
	"github.com/saimib/code-kata-be-go-solution/utils"
)

type Todoist interface {
	GetEvenTodos() []models.Todo
}

func NewTodoist(utils utils.Utils) Todoist {
	return todoist{
		utils: utils,
	}
}

type todoist struct {
	utils utils.Utils
}

func (cm todoist) GetEvenTodos() []models.Todo {

	todos := &[]models.Todo{}
	fetchErr := cm.utils.Fetch(todos)
	if fetchErr != nil {
		fmt.Printf("Could not fetch Todos. Error %s", fetchErr.Error())
		return []models.Todo{}
	}

	result := []models.Todo{}
	for _, todo := range *todos {
		if todo.ID%2 == 0 {
			result = append(result, todo)
		}
		if len(result) == 20 {
			return result
		}
	}
	return result
}
