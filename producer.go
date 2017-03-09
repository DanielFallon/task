package task

// Producer executes a task definition
type Producer interface {
	StartRun() error
	TaskStart() error
	Command(exec string) error
	TaskDone() error
	FinishRun() error
}
