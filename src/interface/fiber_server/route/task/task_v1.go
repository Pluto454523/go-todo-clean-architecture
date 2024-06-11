package task_spec

import (
	"github.com/pluto454523/go-todo-list/src/entity/task"
	"github.com/pluto454523/go-todo-list/src/usecases"
	"github.com/pluto454523/go-todo-list/src/usecases/model"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type (
	TaskHandlerDependency struct {
		useCase *usecases.UsecaseDependency
	}
)

func NewTaskHandler(uc *usecases.UsecaseDependency) *TaskHandlerDependency {
	return &TaskHandlerDependency{
		useCase: uc,
	}
}

func (h TaskHandlerDependency) CreateTask(c *fiber.Ctx) error {

	pl := model.TaskPayload{}
	if err := c.BodyParser(&pl); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Code":    400,
			"Message": err.Error(),
		})
	}

	if pl.Title == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Code":    400,
			"Message": "title are required",
		})
	}

	if pl.Status == "" {
		pl.Status = "todo"
	}

	if pl.DueDate.IsZero() {
		pl.DueDate = time.Now()
	} else if pl.DueDate.Before(time.Now()) {
		pl.DueDate = time.Now()
	}

	t := task.TaskEntity{
		Title:       pl.Title,
		Description: pl.Description,
		Status:      pl.Status,
		DueDate:     pl.DueDate,
	}

	tid, err := h.useCase.CreateTask(c.Context(), t)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Code":    400,
			"Message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(model.TaskResponse{
		ID:          tid,
		Title:       t.Title,
		Description: t.Description,
		Status:      t.Status,
		DueDate:     t.DueDate,
	})
}

func (h TaskHandlerDependency) GetTaskByID(c *fiber.Ctx) error {

	//ctx := c.Locals("otel_trace_context").(context.Context)
	ctx := c.Context()

	taskID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Code":    400,
			"Message": err.Error(),
		})
	}

	var t task.TaskEntity
	t, err = h.useCase.GetTaskByID(ctx, uint(taskID))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Code":    400,
			"Message": err.Error(),
		})
	}

	return c.Status(fiber.StatusAccepted).JSON(model.TaskResponse{
		ID:          t.ID,
		Title:       t.Title,
		Description: t.Description,
		Status:      t.Status,
		DueDate:     t.DueDate,
	})
}

func (h TaskHandlerDependency) GetAllTask(c *fiber.Ctx) error {

	mt := model.TaskOptional{}
	if err := c.QueryParser(&mt); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Code":    400,
			"Message": err.Error(),
		})
	}

	validate := validator.New()
	if err := validate.Struct(mt); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Code":    400,
			"Message": err.Error(),
		})
	}

	ts, err := h.useCase.GetAllTask(c.Context(), mt.Order, mt.Sort, mt.Filter, mt.Value)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Code":    400,
			"Message": err.Error(),
		})
	}

	var rt []model.TaskResponse
	for _, t := range ts {
		rt = append(rt, model.TaskResponse{
			ID:          t.ID,
			Title:       t.Title,
			Description: t.Description,
			Status:      t.Status,
			DueDate:     t.DueDate,
		})
	}

	return c.Status(fiber.StatusAccepted).JSON(rt)
}

func (h TaskHandlerDependency) UpdateTask(c *fiber.Ctx) error {

	taskID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Code":    400,
			"Message": err.Error(),
		})
	}

	pl := model.TaskPayload{}
	if err := c.BodyParser(&pl); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Code":    400,
			"Message": err.Error(),
		})
	}

	t := task.TaskEntity{
		ID:          uint(taskID),
		Title:       pl.Title,
		Description: pl.Description,
		Status:      pl.Status,
		DueDate:     pl.DueDate,
	}

	t, err = h.useCase.UpdateTask(c.Context(), t)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Code":    400,
			"Message": err.Error(),
		})
	}

	return c.Status(fiber.StatusAccepted).JSON(t)
}

func (h TaskHandlerDependency) PatchTask(c *fiber.Ctx) error {

	taskID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Code":    400,
			"Message": err.Error(),
		})
	}

	pl := model.TaskPayload{}
	if err := c.BodyParser(&pl); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Code":    400,
			"Message": err.Error(),
		})
	}

	t := task.TaskEntity{
		ID:          uint(taskID),
		Title:       pl.Title,
		Description: pl.Description,
		Status:      pl.Status,
		DueDate:     pl.DueDate,
	}

	t, err = h.useCase.PatchTask(c.Context(), t)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Code":    400,
			"Message": err.Error(),
		})
	}

	return c.Status(fiber.StatusAccepted).JSON(t)

}

func (h TaskHandlerDependency) DeleteTask(c *fiber.Ctx) error {

	id := c.Params("id")
	taskID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Code":    400,
			"Message": err.Error(),
		})
	}

	if err := h.useCase.DeleteTask(c.Context(), uint(taskID)); err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Code":    400,
			"Message": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h TaskHandlerDependency) ChangeStatus(c *fiber.Ctx) error {

	id := c.Params("id")
	taskID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Code":    400,
			"Message": err.Error(),
		})
	}

	pl := model.TaskPayload{}
	if err := c.BodyParser(&pl); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Code":    400,
			"Message": err.Error(),
		})
	}

	if pl.Status == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Code":    400,
			"Message": "status invalid",
		})
	}

	t, err := h.useCase.GetTaskByID(c.Context(), uint(taskID))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Code":    400,
			"Message": err.Error(),
		})
	}

	t, err = h.useCase.ChangeStatus(c.Context(), uint(taskID), pl.Status)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Code":    400,
			"Message": err.Error(),
		})
	}

	return c.Status(fiber.StatusAccepted).JSON(t)
}
