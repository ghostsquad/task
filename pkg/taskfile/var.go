package taskfile

type Vars map[string]Var

type Var struct {
	// Eval can be used with the go templating in order make a decision
	// we use strconv.ParseBool to evaluate the result
	Eval string `json:"eval,omitempty"`

	// Sh defines a command to run, returns true if the command returns exit code 0, otherwise returns false
	Sh Sh `json:"sh,omitempty"`
}
