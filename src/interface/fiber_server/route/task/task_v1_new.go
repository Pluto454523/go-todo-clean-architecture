package task_spec

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pluto454523/go-todo-list/src/interface/fiber_server/helper"
	"github.com/pluto454523/go-todo-list/src/usecases"
)

type (
	routeTaskV1 struct {
		useCase *usecases.UsecaseDependency
	}
)

func NewRouteTaskV1(uc *usecases.UsecaseDependency) ServerInterface {
	//return routeTaskV1{
	//	useCase: uc,
	//}
	panic("maintainane")
}

func (rt routeTaskV1) GetTasks(c *fiber.Ctx, params GetTasksParams) error {

	ts, err := rt.useCase.GetAllTask(c.Context(), *params.Order, *params.Sort, *params.Filter, *params.Value)
	if err != nil {
		return helper.ErrorHandler(c, err)
	}

	var tr []TaskResponse
	for _, t := range ts {
		taskId := int(t.ID)
		tr = append(tr, TaskResponse{
			Id:          &taskId,
			Title:       &t.Title,
			Description: &t.Description,
			Status:      &t.Status,
			DueDate:     &t.DueDate,
		})
	}

	return c.Status(fiber.StatusAccepted).JSON(tr)
}
