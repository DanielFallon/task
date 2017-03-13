package task

import "github.com/go-task/task/producer"

var (
	currentProducer Producer
)

// Producer executes a task definition
type Producer interface {
	StartRun() error
	RunTask(task string) error
	RunCommand(command string, variables map[string]string, options producer.CommandOptions) error
	FinishRun() error
}

func setProducer() error {
	switch ProducerFlag {
	case "exec":
		currentProducer = producer.ExecProducer{}
	case "bash":
		currentProducer = producer.BashProducer{}
	}
	return nil
}

func init() {
	currentProducer = producer.ExecProducer{}
}
