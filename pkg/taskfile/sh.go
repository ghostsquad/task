package taskfile

type Sh struct {
	// Run defines a shell or binary command to run.
	// Run is mutually exclusive with `task`. If both are defined, and error will be produced.
	Run string `json:"run,omitempty"`

	// Dir is the desired working directory in which this step's `run` or `task` should execute in.
	// Noting that if you are calling a task, this directory supersedes the directory described in that task.
	Args string `json:"args,omitempty"`

	// Args are an optional array of additional args for command if `run` is used.
	// Args has no effect if `task` is empty.
	Dir string `json:"dir,omitempty"`
}
