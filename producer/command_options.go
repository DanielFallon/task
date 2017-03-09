package producer

// CommandOptions provide the producer with some options for the command to be executed
type CommandOptions struct {
	Dir  string
	Vars map[string]string
	Set  string
	Env  map[string]string
}
