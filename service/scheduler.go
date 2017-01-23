package service

import "time"

// Scheduler interface
type Scheduler interface {
	Execute()
	Register()
	Unregister()
}

// Task structure
type Task struct {
	Signals   []*Signal `json:"signals"`
	CreatedAt time.Time `json:"created_at"`
}

// Execute implement Scheduler
func (t *Task) Execute() error {
	for _, signal := range t.Signals {
		err := SendSignal(&SendSignalParam{Signal: signal})
		if err != nil {
			return err
		}
	}
	return nil
}

// OnceScheduledTask structure
type OnceScheduledTask struct {
	*Task
	Time time.Time `json:"time"`
}

// Register implement Scheduler
func (t *OnceScheduledTask) Register() error {
	return nil
}

// Unregister implement Scheduler
func (t *OnceScheduledTask) Unregister() error {
	return nil
}

// RoutineSchedulerTask structure
type RoutineSchedulerTask struct {
	*Task
	Rule string `json:"rule"`
}

// Register implement Scheduler
func (t *RoutineSchedulerTask) Register() error {
	return nil
}

// Unregister implement Scheduler
func (t *RoutineSchedulerTask) Unregister() error {
	return nil
}

// AddScheduler .
func AddScheduler(s Scheduler) {
	s.Register()
}

// RemoveScheduler .
func RemoveScheduler(s Scheduler) {

}
