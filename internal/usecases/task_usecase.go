package usecases

import (
	"context"
	"github.com/pluto454523/go-todo-list/internal/entity/task"
	"github.com/pluto454523/go-todo-list/internal/usecases/repository"
	"go.opentelemetry.io/otel/attribute"
)

func (uc UsecaseDependency) CreateTask(ctx context.Context, tk task.TaskEntity) (uint, error) {

	ctx, sp := tracer.Start(ctx, "UsecaseDependency.CreateTask")
	defer sp.End()

	// Validate TaskEntity
	if err := tk.Validate(); err != nil {
		sp.RecordError(err)
		return 0, err
	}
	sp.AddEvent("validate task")

	taskId, err := uc.TaskRepository.CreateTask(ctx, tk)
	if err != nil {
		sp.RecordError(err)
		return 0, err
	}
	sp.AddEvent("create task")
	sp.SetAttributes(attribute.Int(
		"task_id",
		int(taskId),
	))

	return taskId, nil

}

func (uc UsecaseDependency) GetTaskByID(ctx context.Context, id uint) (et task.TaskEntity, err error) {

	ctx, sp := tracer.Start(ctx, "UsecaseDependency.GetTaskByID")
	defer sp.End()

	// Fetch the task by ID
	et, err = uc.TaskRepository.GetTaskByID(ctx, id)
	if err != nil {
		sp.RecordError(err)
		return task.TaskEntity{}, err
	}
	sp.AddEvent("get task by id")
	sp.SetAttributes(attribute.Int(
		"task_id",
		int(et.ID),
	))

	return task.TaskEntity{
		ID:          et.ID,
		Title:       et.Title,
		Description: et.Description,
		Status:      et.Status,
	}, nil
}

func (uc UsecaseDependency) GetAllTask(ctx context.Context, order string, sort string, filter string, value string) (tasks []task.TaskEntity, err error) {

	ctx, sp := tracer.Start(ctx, "UsecaseDependency.GetAllTask")
	defer sp.End()

	// List order
	//fo := repository.CreatedBetweenFilterOption{
	//	After:  time.Now().AddDate(-1, 0, 0),
	//	Before: time.Now(),
	//}

	fo := repository.CustomFieldValueFilterOption{
		Field: filter,
		Value: value,
	}

	if filter == "" {
		fo.Field = "title"
	}

	so := repository.CustomFieldSortOption{
		Field: order,
		Desc:  sort == "desc",
	}

	if order == "" {
		so.Field = "id"
	}

	tasks, err = uc.TaskRepository.GetAllTask(ctx, &fo, &so)
	if err != nil {
		sp.RecordError(err)
		return []task.TaskEntity{}, err
	}
	sp.AddEvent("get task by id")

	return tasks, nil
}

func (uc UsecaseDependency) UpdateTask(ctx context.Context, ntk task.TaskEntity) (task.TaskEntity, error) {

	ctx, sp := tracer.Start(ctx, "UsecaseDependency.UpdateTask")
	defer sp.End()

	// Validate new task entity
	if err := ntk.Validate(); err != nil {
		sp.RecordError(err)
		return task.TaskEntity{}, err
	}
	sp.AddEvent("validate TaskEntity")

	// save task to database
	err := uc.TaskRepository.UpdateTask(ctx, ntk)
	if err != nil {
		return task.TaskEntity{}, err
	}
	sp.AddEvent("save task")
	sp.SetAttributes(attribute.Int(
		"task_id",
		int(ntk.ID),
	))

	return ntk, nil
}

func (uc UsecaseDependency) PatchTask(ctx context.Context, tk task.TaskEntity) (task.TaskEntity, error) {

	ctx, sp := tracer.Start(ctx, "UsecaseDependency.PatchTask")
	defer sp.End()

	// Fetch the existingTask by ID
	etk, err := uc.TaskRepository.GetTaskByID(ctx, tk.ID)
	if err != nil {
		sp.RecordError(err)
		return task.TaskEntity{}, err
	}
	sp.AddEvent("get task by id")

	// Update entity
	etk = etk.PatchTask(tk)
	sp.AddEvent("patch task")

	// Update task to database
	err = uc.TaskRepository.UpdateTask(ctx, etk)
	if err != nil {
		return task.TaskEntity{}, err
	}
	sp.AddEvent("save task")
	sp.SetAttributes(attribute.Int(
		"task_id",
		int(tk.ID),
	))

	return etk, nil
}

func (uc UsecaseDependency) DeleteTask(ctx context.Context, id uint) error {

	ctx, sp := tracer.Start(ctx, "UsecaseDependency.DeleteTask")
	defer sp.End()

	err := uc.TaskRepository.DeleteTask(ctx, id)
	if err != nil {
		sp.RecordError(err)
		return err
	}
	sp.AddEvent("delete task")
	sp.SetAttributes(attribute.Int(
		"task_id",
		int(id),
	))

	return nil
}

func (uc UsecaseDependency) ChangeStatus(ctx context.Context, id uint, status string) (tk task.TaskEntity, err error) {

	ctx, sp := tracer.Start(ctx, "UsecaseDependency.ChangeStatus")
	defer sp.End()

	// Fetch the task by ID
	tk, err = uc.TaskRepository.GetTaskByID(ctx, id)
	if err != nil {
		sp.RecordError(err)
		return task.TaskEntity{}, err
	}
	sp.AddEvent("get task by id")

	// Change status
	tk, err = tk.ChangeStatus(status)
	if err != nil {
		sp.RecordError(err)
		return task.TaskEntity{}, err
	}
	sp.AddEvent("change status")

	// Save the updated task
	err = uc.TaskRepository.UpdateTask(ctx, tk)
	if err != nil {
		sp.RecordError(err)
		return task.TaskEntity{}, err
	}
	sp.AddEvent("update task")
	sp.SetAttributes(attribute.Int(
		"task_id",
		int(id),
	))

	return tk, nil
}
