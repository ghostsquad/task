package taskfile

type Steps []Step

// Step is a specific command (or even another task) to execute
// The use of Run & Args is mutually exclusive with Task
type Step struct {
	Sh
	// Task allows you to call another task (instead of command via the `run` field).
	// Task is mutually exclusive with `run`. If both are defined, and error will be produced.
	Task string `json:"task,omitempty"`

	Vars Vars `json:"vars,omitempty"`

	// IfSpec or just `if` in the Taskfile yaml, defines rules for whether this step should run.
	IfSpec IfSpec `json:"if,omitempty"`

	// Matrix is a step/task hook to define a matrix of step/task configurations.
	//
	// A matrix allows you to create multiple steps/tasks by performing variable substitution
	// in a single step/task definition.
	//
	// When defined on a Step, it dynamically creates several steps. Each step is inlined, and are run serially.
	// Step ordering is based on the order of each provided matrix.
	// Example
	//
	//	task:
	//	  echo:
	//		matrix:
	//		- env: ["dev", "staging", "prod"]
	//		- ham: ["bacon", "eggs"]
	//		steps:
	//		- run: "echo {{.Matrix.env}} {{.Matrix.ham}}"
	//
	// Steps created defined are:
	//  - echo "dev bacon"
	//  - echo "dev eggs"
	//  - echo "staging bacon"
	//  - echo "staging eggs"
	//  - echo "prod bacon"
	//  - echo "prod eggs"
	Matrix Matrix `json:"matrix,omitempty"`

	// Timeout is the umbrella bounding time limit (duration) for the task before signalling for termination via SIGINT.
	Timeout Duration `json:"timeout,omitempty"`

	// GracefulTermination is the bounding time limit (duration) for this task before sending subprocesses a SIGKILL.
	GracefulTermination Duration `json:"gracefulTermination,omitempty"`

	// OnSuccess is a step/task hook for a series of steps to run if the current step/task succeeds.
	OnSuccess Steps `json:"onSuccess,omitempty"`

	// OnFailure is a step/task hook for a series of steps to run if the current step/task fails.
	OnFailure Steps `json:"onFailure,omitempty"`

	// Defer is a step/task hook that defers a series of steps until the current step/task execution completes, whether success or failure.
	Defer Steps `json:"defer,omitempty"`
}
