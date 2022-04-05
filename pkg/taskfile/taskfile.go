package taskfile

type PatternVar struct {
	Pattern string
	Var     Var
}

type Taskfile struct {
	Version  string          `json:"version,omitempty"`
	Vars     Vars            `json:"vars,omitempty"`
	Includes Includes        `json:"includes,omitempty"`
	Features Features        `json:"features,omitempty"`
	Tasks    map[string]Task `json:"tasks,omitempty"`
}
