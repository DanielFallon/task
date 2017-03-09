package producer

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/go-task/helper"
)

// ExecProducer is the default producer for a task run
type ExecProducer struct {
}

// StartRun is called before the tasks are run
func (ep ExecProducer) StartRun() error {
	fmt.Println("Starting task")
	return nil
}

// RunTask is called for each task that is about to be executed
func (ep ExecProducer) RunTask(task string) error {
	fmt.Printf("\n*** %s ***\n\n", task)
	return nil
}

// RunCommand is called for each command of a task
func (ep ExecProducer) RunCommand(command string, variables map[string]string, options CommandOptions) error {
	c, err := helper.ReplaceVariables(command, variables)
	if err != nil {
		return err
	}
	fmt.Printf("--> %s\n", c)
	dir, err := helper.ReplaceVariables(options.Dir, variables)
	if err != nil {
		return err
	}
	var cmd *exec.Cmd
	if helper.ShExists {
		cmd = exec.Command(helper.ShPath, "-c", c)
	} else {
		cmd = exec.Command("cmd", "/C", c)
	}
	if dir != "" {
		cmd.Dir = dir
	}
	if options.Env != nil {
		cmd.Env = os.Environ()
		for key, value := range options.Env {
			replacedValue, err := helper.ReplaceVariables(value, variables)
			if err != nil {
				return err
			}
			replacedKey, err := helper.ReplaceVariables(key, variables)
			if err != nil {
				return err
			}
			cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", replacedKey, replacedValue))
		}
	}
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	if options.Set != "" {
		bytes, err := cmd.Output()
		if err != nil {
			return err
		}
		os.Setenv(options.Set, strings.TrimSpace(string(bytes)))
		return nil
	}
	cmd.Stdout = os.Stdout
	if err = cmd.Run(); err != nil {
		return err
	}
	return nil
}

// FinishRun is called at the end
func (ep ExecProducer) FinishRun() error {
	fmt.Println("Finished task")
	return nil
}
