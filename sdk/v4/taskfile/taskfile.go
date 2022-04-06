package taskfile

type PatternVar struct {
	Pattern string
	Var     VarValue
}

type Taskfile struct {
	// Version is the version of the Task schema
	// It is the first thing that is parsed by Task in order to inform how the rest of the file should be parsed.
	Version string `json:"version,omitempty"`

	// Features are flags for task
	// It enables users to change the behavior of task by enabling/disabling features of the engine
	Features Features `json:"features,omitempty"`

	// Vars are global variables, usually used as default values
	Vars Vars `json:"vars,omitempty"`

	// Includes allows other Taskfiles to be included in this one. Each other include is becomes "namespaced".
	// More on namespacing later...
	Includes Includes `json:"includes,omitempty"`

	// Tasks is a set of task definitions
	Tasks map[string]Task `json:"tasks,omitempty"`
}
