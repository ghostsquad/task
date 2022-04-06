package taskfile

type Sh struct {
	// Run defines a shell or binary command to run.
	Run string `json:"run,omitempty"`

	// Args are an optional array of additional args for the command defined in Run
	Args []string `json:"args,omitempty"`

	// Dir is the desired working directory in which this command should execute in.
	Dir string `json:"dir,omitempty"`
}
