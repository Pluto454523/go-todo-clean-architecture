package usecases

import (
	"context"
	"errors"
	"github.com/pluto454523/go-todo-list/internal/entity/task"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestTaskUseCase_CreateTask(t *testing.T) {

	t.Run("Success case #1", func(t *testing.T) {

		// Setup
		ctx := context.Background()
		tr, _, uc := newMocksRepository()
		tk := newValidTask()

		// Mocks
		tr.On("CreateTask", mock.Anything, tk).Return(uint(1), nil)

		// Execution
		taskId, err := uc.CreateTask(ctx, tk)

		// Assertions
		assertion := assert.New(t)
		assertion.NoError(err, "validation")
		assertion.Equal(uint(1), taskId)
		tr.AssertExpectations(t)
	})

	t.Run("Failure validation #2", func(t *testing.T) {

		// Setup
		ctx := context.Background()
		tr, _, uc := newMocksRepository()
		tk := newValidTask()

		// Mocks
		tr.On("CreateTask", mock.Anything, tk).Return(uint(1), nil)

		// Here, we'll deliberately introduce a validation error by providing an invalid task
		invalidTask := newValidTask()
		invalidTask.Status = ""

		// Execution
		taskId, err := uc.CreateTask(ctx, invalidTask)

		// Assertions
		assertion := assert.New(t)
		assertion.Errorf(err, "validation")
		assertion.Empty(taskId)
		tr.AssertNumberOfCalls(t, "CreateTask", 0)

	})

	t.Run("Failure case #3", func(t *testing.T) {

		// Setup
		ctx := context.Background()
		tr, _, uc := newMocksRepository()
		tk := newValidTask()

		// Mocks
		tr.On("CreateTask", mock.Anything, tk).Return(uint(1), errors.New("error"))

		// Execution
		taskId, err := uc.CreateTask(ctx, tk)

		// Assertions
		assertion := assert.New(t)
		assertion.Error(err, "CreateTask Failed")
		assertion.Empty(taskId)
		tr.AssertExpectations(t)

	})
}

func TestTaskUseCase_GetTaskByID(t *testing.T) {

	t.Run("Success get task by task_id #1", func(t *testing.T) {

		// Setup
		ctx := context.Background()
		tr, _, uc := newMocksRepository()
		tk := newValidTask()

		// Mocks
		tr.On("GetTaskByID", mock.Anything, uint(1)).Return(tk, nil)

		// Execution
		ntk, err := uc.GetTaskByID(ctx, 1)

		// Assertions
		assertion := assert.New(t)
		assertion.NoError(err)
		assertion.NotEmpty(ntk)

		tr.AssertExpectations(t)
	})

	t.Run("Failure task not found #2", func(t *testing.T) {

		// Setup
		ctx := context.Background()
		tr, _, uc := newMocksRepository()
		tk := newValidTask()

		// Mocks
		tr.On("GetTaskByID", mock.Anything, tk.ID).Return(
			task.TaskEntity{},
			errors.New("task not found"),
		)

		// Execution
		nkt, err := uc.GetTaskByID(ctx, tk.ID)

		// Assertions
		assertion := assert.New(t)
		assertion.Error(err, "Task not found")
		assertion.Empty(nkt)
		tr.AssertExpectations(t)

	})
}

func TestTaskUseCase_GetAllTask(t *testing.T) {

	t.Run("Success case with no filter and sort #1", func(t *testing.T) {

		// Setup
		ctx := context.Background()
		tr, _, uc := newMocksRepository()

		tks := []task.TaskEntity{
			newValidTask(),
			newValidTask(),
		}

		// Mocks
		tr.On("GetAllTask", mock.Anything, mock.AnythingOfType("*repository.CustomFieldValueFilterOption"), mock.AnythingOfType("*repository.CustomFieldSortOption")).Return(tks, nil)

		// Execution
		ntk, err := uc.GetAllTask(ctx, "", "", "", "")

		// Assertions
		assertion := assert.New(t)
		assertion.NoError(err)
		assertion.NotEmpty(ntk)
		assertion.Equal(tks, ntk)

		tr.AssertExpectations(t)
	})

	t.Run("success case with filter and sort #2", func(t *testing.T) {

		// Setup
		ctx := context.Background()
		tr, _, uc := newMocksRepository()

		tks := []task.TaskEntity{
			newValidTask(),
			newValidTask(),
		}

		// Mocks
		tr.On("GetAllTask", mock.Anything, mock.AnythingOfType("*repository.CustomFieldValueFilterOption"), mock.AnythingOfType("*repository.CustomFieldSortOption")).Return(tks, nil)

		// Execution
		ntk, err := uc.GetAllTask(ctx, "status", "desc", "id", "1")

		// Assertions
		assertion := assert.New(t)
		assertion.NoError(err)
		assertion.NotEmpty(ntk)
		assertion.Equal(tks, ntk)

		tr.AssertExpectations(t)
	})

	t.Run("Failure case #3", func(t *testing.T) {

		// Setup
		ctx := context.Background()
		tr, _, uc := newMocksRepository()

		tks := []task.TaskEntity{
			newValidTask(),
			newValidTask(),
		}

		// Mocks
		tr.On("GetAllTask", mock.Anything,
			mock.AnythingOfType("*repository.CustomFieldValueFilterOption"),
			mock.AnythingOfType("*repository.CustomFieldSortOption")).
			Return(tks, errors.New("error repository"))

		// Execution
		ntk, err := uc.GetAllTask(ctx, "status", "desc", "id", "1")

		// Assertions
		assertion := assert.New(t)
		assertion.Error(err)
		assertion.Empty(ntk)

		tr.AssertExpectations(t)
	})
}

func TestTaskUseCase_UpdateTask(t *testing.T) {

	t.Run("Success case #1", func(t *testing.T) {

		// Setup
		ctx := context.Background()
		tr, _, uc := newMocksRepository()

		tk := newValidTask()

		// Mocks
		tr.On("UpdateTask", mock.Anything, tk).Return(nil)

		// Execution
		ntk, err := uc.UpdateTask(ctx, tk)

		// Assertions
		assertion := assert.New(t)
		assertion.NoError(err)
		assertion.NotEmpty(ntk)
		assertion.Equal(tk, ntk)

		tr.AssertExpectations(t)
	})

	t.Run("Failure validation #2", func(t *testing.T) {

		// Setup
		ctx := context.Background()
		tr, _, uc := newMocksRepository()

		tk := newValidTask()
		tk.Status = ""

		// Mocks
		tr.On("UpdateTask", mock.Anything, tk).Return(nil)

		// Execution
		ntk, err := uc.UpdateTask(ctx, tk)

		// Assertions
		assertion := assert.New(t)
		assertion.Error(err)
		assertion.Empty(ntk)
		tr.AssertNumberOfCalls(t, "UpdateTask", 0)
	})

	t.Run("Failure case #3", func(t *testing.T) {

		// Setup
		ctx := context.Background()
		tr, _, uc := newMocksRepository()

		tk := newValidTask()

		// Mocks
		tr.On("UpdateTask", mock.Anything, tk).Return(errors.New("error repository"))

		// Execution
		ntk, err := uc.UpdateTask(ctx, tk)

		// Assertions
		assertion := assert.New(t)
		assertion.Error(err)
		assertion.Empty(ntk)
		tr.AssertNumberOfCalls(t, "UpdateTask", 1)
	})
}

func TestTaskUseCase_PatchTask(t *testing.T) {

	t.Run("Success case #1", func(t *testing.T) {

		// Setup
		ctx := context.Background()
		tr, _, uc := newMocksRepository()

		mockTk := newValidTask()
		mockPatchTk := task.TaskEntity{
			Title: "Patch Task 1",
		}
		expectaionTask := mockTk
		expectaionTask.Title = mockPatchTk.Title

		// Mocks
		tr.On("GetTaskByID", mock.Anything, uint(0)).Return(mockTk, nil)
		tr.On("UpdateTask", mock.Anything, mock.Anything).Return(nil)

		// Execution
		newTk, err := uc.PatchTask(ctx, mockPatchTk)

		// Assertions
		assertion := assert.New(t)
		assertion.NoError(err)
		assertion.NotEmpty(newTk)
		assertion.Equal(expectaionTask, newTk)
		tr.AssertExpectations(t)
	})

	t.Run("Failure case #2 task not found", func(t *testing.T) {

		// Setup
		ctx := context.Background()
		tr, _, uc := newMocksRepository()

		mockTk := newValidTask()
		mockPatchTk := task.TaskEntity{
			Title: "Patch Task 1",
		}
		expectaionTask := mockTk
		expectaionTask.Title = mockPatchTk.Title

		// Mocks
		tr.On("GetTaskByID", mock.Anything, uint(0)).Return(
			task.TaskEntity{}, errors.New("task not found"))
		tr.On("UpdateTask", mock.Anything, mock.Anything).Return(nil)

		// Execution
		newTk, err := uc.PatchTask(ctx, mockPatchTk)

		// Assertions
		assertion := assert.New(t)
		assertion.EqualError(err, "task not found")
		assertion.Empty(newTk)
		assertion.NotEqual(expectaionTask, newTk)
		tr.AssertNumberOfCalls(t, "GetTaskByID", 1)
		tr.AssertNumberOfCalls(t, "UpdateTask", 0)
	})

	t.Run("Failure case #2 update task error", func(t *testing.T) {

		// Setup
		ctx := context.Background()
		tr, _, uc := newMocksRepository()

		mockTk := newValidTask()
		mockPatchTk := task.TaskEntity{
			Title: "Patch Task 1",
		}
		expectaionTask := mockTk
		expectaionTask.Title = mockPatchTk.Title

		// Mocks
		tr.On("GetTaskByID", mock.Anything, uint(0)).Return(
			mockTk, nil)
		tr.On("UpdateTask", mock.Anything, mock.Anything).Return(
			errors.New("error repository"))

		// Execution
		newTk, err := uc.PatchTask(ctx, mockPatchTk)

		// Assertions
		assertion := assert.New(t)
		assertion.EqualError(err, "error repository")
		assertion.Empty(newTk)
		assertion.NotEqual(expectaionTask, newTk)
		tr.AssertNumberOfCalls(t, "GetTaskByID", 1)
		tr.AssertNumberOfCalls(t, "UpdateTask", 1)
	})
}

func TestTaskUseCase_DeleteTask(t *testing.T) {

	t.Run("Success case #1", func(t *testing.T) {

		// Setup
		ctx := context.Background()
		tr, _, uc := newMocksRepository()

		// Mocks
		tr.On("DeleteTask", mock.Anything, uint(0)).Return(nil)

		// Execution
		err := uc.DeleteTask(ctx, uint(0))

		// Assertions
		assertion := assert.New(t)
		assertion.NoError(err)
		tr.AssertExpectations(t)
	})

	t.Run("Failure case #2 repository error", func(t *testing.T) {

		// Setup
		ctx := context.Background()
		tr, _, uc := newMocksRepository()

		// Mocks
		tr.On("DeleteTask", mock.Anything, uint(0)).Return(errors.New("task not found"))

		// Execution
		err := uc.DeleteTask(ctx, uint(0))

		// Assertions
		assertion := assert.New(t)
		assertion.Errorf(err, "task not found")
		tr.AssertExpectations(t)
	})

}

func TestTaskUseCase_ChangeStatus(t *testing.T) {

	t.Run("Success case #1", func(t *testing.T) {

		// Setup
		ctx := context.Background()
		tr, _, uc := newMocksRepository()

		mockTk := newValidTask()
		mockTk.Status = "todo"
		ExpectedTk := mockTk
		ExpectedTk.Status = "doing"

		// Mocks
		tr.On("GetTaskByID", mock.Anything, mockTk.ID).Return(mockTk, nil)
		tr.On("UpdateTask", mock.Anything, ExpectedTk).Return(nil)

		// Execution
		ntk, err := uc.ChangeStatus(ctx, mockTk.ID, ExpectedTk.Status)

		// Assertions
		assertion := assert.New(t)
		assertion.NoError(err)
		assertion.NotEmpty(ntk)
		assertion.Equal(ExpectedTk, ntk)
		tr.AssertExpectations(t)
	})

	t.Run("Failure case #2 get task not found", func(t *testing.T) {

		// Setup
		ctx := context.Background()
		tr, _, uc := newMocksRepository()

		mockTk := newValidTask()
		mockTk.Status = "todo"
		ExpectedTk := mockTk
		ExpectedTk.Status = "doing"

		// Mocks
		tr.On("GetTaskByID", mock.Anything, mockTk.ID).Return(task.TaskEntity{}, errors.New("task not found"))
		tr.On("UpdateTask", mock.Anything, ExpectedTk).Return(nil)

		// Execution
		ntk, err := uc.ChangeStatus(ctx, mockTk.ID, ExpectedTk.Status)

		// Assertions
		assertion := assert.New(t)
		assertion.Errorf(err, "task not found")
		assertion.Empty(ntk)
		tr.AssertNumberOfCalls(t, "GetTaskByID", 1)
		tr.AssertNumberOfCalls(t, "UpdateTask", 0)
	})

	t.Run("Failure case #2 update task fail", func(t *testing.T) {

		// Setup
		ctx := context.Background()
		tr, _, uc := newMocksRepository()

		mockTk := newValidTask()
		mockTk.Status = "todo"
		ExpectedTk := mockTk
		ExpectedTk.Status = "doing"

		// Mocks
		tr.On("GetTaskByID", mock.Anything, mockTk.ID).Return(mockTk, nil)
		tr.On("UpdateTask", mock.Anything, ExpectedTk).Return(errors.New("update fail"))

		// Execution
		ntk, err := uc.ChangeStatus(ctx, mockTk.ID, ExpectedTk.Status)

		// Assertions
		assertion := assert.New(t)
		assertion.Errorf(err, "update fail")
		assertion.Empty(ntk)
		tr.AssertExpectations(t)
	})

	t.Run("Failure case #3 change status fail", func(t *testing.T) {

		// Setup
		ctx := context.Background()
		tr, _, uc := newMocksRepository()

		mockTk := newValidTask()
		mockTk.Status = "todo"
		ExpectedTk := mockTk
		ExpectedTk.Status = "done"

		// Mocks
		tr.On("GetTaskByID", mock.Anything, mockTk.ID).Return(mockTk, nil)
		tr.On("UpdateTask", mock.Anything, ExpectedTk).Return(nil)

		// Execution
		ntk, err := uc.ChangeStatus(ctx, mockTk.ID, ExpectedTk.Status)

		// Assertions
		assertion := assert.New(t)
		assertion.Error(err)
		assertion.Empty(ntk)
		tr.AssertNumberOfCalls(t, "GetTaskByID", 1)
		tr.AssertNumberOfCalls(t, "UpdateTask", 0)
	})
}
