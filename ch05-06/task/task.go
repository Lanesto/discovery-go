package task

import (
	"strconv"
	"time"

	"example/status"
)

// Deadline is struct embedding time.Time
type Deadline struct {
	time.Time
}

// Task is struct containing information about task
type Task struct {
	Title    string        `json:"title,"`
	Status   status.Status `json:"status,"`
	Deadline *Deadline     `json:"deadline,omitempty"`
	SubTasks []Task        `json:"subTasks,omitempty"`
}

// MarshalJSON implements json.Marshaler for Deadline type; convert to unix time.
func (d Deadline) MarshalJSON() ([]byte, error) {
	return strconv.AppendInt(nil, d.Unix(), 10), nil
}

// UnmarshalJSON implements json.Unmarshaler for Deadline type; convert from unix time.
func (d *Deadline) UnmarshalJSON(data []byte) error {
	t, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return err
	}
	d.Time = time.Unix(t, 0)
	return nil
}

// OverDue returns whether passed deadline of task.
func (t *Task) OverDue() bool {
	return t.Deadline.Before(time.Now())
}
