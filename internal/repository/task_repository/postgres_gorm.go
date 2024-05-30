package task_repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/pluto454523/go-todo-list/internal/entity/task"
	"github.com/pluto454523/go-todo-list/internal/repository/migrations"
	"github.com/pluto454523/go-todo-list/internal/usecases/repository"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type TaskRepositoryDependency struct {
	DB *gorm.DB
}

func NewTaskRepository(db *gorm.DB) repository.TaskRepository {

	// ? Create a new migration instance
	migrator := migrations.NewMigration(db, &taskCollectionSchema{})

	// ? Drop table
	//if err := migrator.DropTable(); err != nil {
	//	defer log.Fatal().
	//		Err(err).
	//		Msg("drop table failed")
	//}

	// ? Start the migration
	if err := migrator.Start(); err != nil {
		defer log.Fatal().
			Err(err).
			Msg("migration failed")
	}

	return &TaskRepositoryDependency{
		DB: db,
	}
}

func (r TaskRepositoryDependency) CreateTask(ctx context.Context, et task.TaskEntity) (uint, error) {

	ctx, sp := tracer.Start(ctx, "TaskRepositoryDependency.CreateTask")
	defer sp.End()

	//info
	s := taskCollectionSchema{
		Title:       et.Title,
		Description: et.Description,
		Status:      et.Status,
		DueDate:     et.DueDate,
	}

	tx := r.DB.WithContext(ctx).Create(&s)

	if tx.Error != nil {
		return 0, fmt.Errorf("failed to create task: %v", tx.Error)
	}

	//sp.SpanContext()

	return s.ID, nil
}

func (r TaskRepositoryDependency) GetTaskByID(ctx context.Context, id uint) (task.TaskEntity, error) {

	ctx, sp := tracer.Start(ctx, "TaskRepositoryDependency.GetTaskByID")
	defer sp.End()

	s := taskCollectionSchema{}

	tx := r.DB.WithContext(ctx).Where("is_deleted = ?", false).First(&s, id)

	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return task.TaskEntity{}, fmt.Errorf("task not found")
	}

	if tx.Error != nil {
		return task.TaskEntity{}, tx.Error
	}

	return task.TaskEntity{
		ID:          s.ID,
		Title:       s.Title,
		Description: s.Description,
		DueDate:     s.DueDate,
		Status:      s.Status,
	}, nil
}

func (r TaskRepositoryDependency) GetAllTask(ctx context.Context, fo repository.FilterOption, so repository.SortOption) ([]task.TaskEntity, error) {

	ctx, sp := tracer.Start(ctx, "TaskRepositoryDependency.GetAllTask")
	defer sp.End()
	var s []taskCollectionSchema
	tx := r.DB.Where("is_deleted = ?", false)

	tx, err := applyPgFilterOption(tx, fo)
	if err != nil {
		log.Error().Err(err).Msg("failed to apply filter option")
		//WithAppJsonData("filter_option", f)
		return nil, err
	}

	tx, err = applyPgSortOption(tx, so)
	if err != nil {
		log.Error().Err(err).Msg("failed to apply sort option")
		return nil, err
	}

	if err := tx.WithContext(ctx).Find(&s).Error; err != nil {
		return nil, err
	}

	var ets []task.TaskEntity
	for _, ts := range s {
		ets = append(ets, task.TaskEntity{
			ID:          ts.ID,
			Title:       ts.Title,
			Description: ts.Description,
			DueDate:     ts.DueDate,
			Status:      ts.Status,
		})
	}

	return ets, nil
}

func (r TaskRepositoryDependency) UpdateTask(ctx context.Context, t task.TaskEntity) (err error) {

	ctx, sp := tracer.Start(ctx, "TaskRepositoryDependency.UpdateTask")
	defer sp.End()

	s := &taskCollectionSchema{
		ID:          t.ID,
		Title:       t.Title,
		Description: t.Description,
		Status:      t.Status,
		DueDate:     t.DueDate,
	}

	// result := r.DB.Model(task).Updates(*task)
	result := r.DB.WithContext(ctx).Save(s)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r TaskRepositoryDependency) DeleteTask(ctx context.Context, id uint) error {

	ctx, sp := tracer.Start(ctx, "TaskRepositoryDependency.DeleteTask")
	defer sp.End()

	var s taskCollectionSchema
	// result := r.DB.Delete(&task, id)
	result := r.DB.WithContext(ctx).Model(&s).
		Where("id = ?", id).
		Update("is_deleted", true)

	if result.RowsAffected == 0 {
		return fmt.Errorf("task not found")
	}
	
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r TaskRepositoryDependency) HardDeleteTask(ctx context.Context, id uint) error {

	ctx, sp := tracer.Start(ctx, "TaskRepositoryDependency.HardDeleteTask")
	defer sp.End()

	var s taskCollectionSchema
	// result := r.DB.Unscoped().Delete(&task, id)
	result := r.DB.WithContext(ctx).Delete(&s, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
