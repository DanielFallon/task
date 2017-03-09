package producer

import (
	"fmt"

	"strings"

	"github.com/go-task/helper"
)

// BashProducer writes out a bash script
type BashProducer struct {
}

// StartRun is called before the tasks are run
func (bp BashProducer) StartRun() error {
	fmt.Println("#! /bin/bash")
	return nil
}

// RunTask is called for each task that is about to be executed
func (bp BashProducer) RunTask(task string) error {
	fmt.Println("echo 'task'")
	return nil
}

// RunCommand is called for each command of a task
func (bp BashProducer) RunCommand(command string, variables map[string]string, options CommandOptions) error {
	c, err := helper.ReplaceVariables(command, variables)
	if err != nil {
		return err
	}
	dir, err := helper.ReplaceVariables(options.Dir, variables)
	if err != nil {
		return err
	}
	additionalEnvironment := ""
	if options.Env != nil {
		for key, value := range options.Env {
			additionalEnvironment = fmt.Sprintf("%s%s=%s ", additionalEnvironment, key, value)
		}
	}
	if "" == dir {
		printCommand(c, additionalEnvironment, options)
	} else {
		fmt.Printf("cd \"%s\"\n", dir)
		printCommand(c, additionalEnvironment, options)
		fmt.Println("cd -")
	}
	return nil
}

func printCommand(c, additionalEnvironment string, options CommandOptions) {
	if options.Set == "" {
		fmt.Printf("sh -c '%s; %s'", strings.TrimSpace(additionalEnvironment), c)
	} else {
		fmt.Printf("export %s=`sh -c '\"%s; %s\"'`", options.Set, strings.TrimSpace(additionalEnvironment), c)
	}
}

// FinishRun is called at the end
func (bp BashProducer) FinishRun() error {
	return nil
}
