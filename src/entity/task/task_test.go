package task

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

// Mock และ Spy อาจจะไม่จำเป็นในตัวอย่างนี้ เพราะ method ChangeStatus ไม่มี dependency ภายนอก
// แต่ถ้ามี dependency เราสามารถใช้ mock และ spy ได้

// Stub: ตัวอย่างนี้เป็นการสร้างสถานการณ์ทดสอบ
func TestChangeStatus(t *testing.T) {

	testCases := []struct {
		initialStatus string
		newStatus     string
		expectedError string
	}{
		{"", "todo", ""},
		{"todo", "doing", ""},
		{"doing", "done", ""},
		{"todo", "done", "must change to doing before done"},
		{"done", "todo", "can't change done status"},
		{"doing", "doing", "can't change current status"},
	}

	for _, tc := range testCases {
		t.Run(tc.initialStatus+" to "+tc.newStatus, func(t *testing.T) {
			taskEntity := TaskEntity{
				ID:          1,
				Title:       "Test Task",
				Description: "This is a test task",
				DueDate:     time.Now().Add(24 * time.Hour),
				Status:      tc.initialStatus,
			}

			updatedTask, err := taskEntity.ChangeStatus(tc.newStatus)

			if tc.expectedError != "" {
				assert.EqualError(t, err, tc.expectedError)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.newStatus, updatedTask.Status)
			}
		})
	}
}
