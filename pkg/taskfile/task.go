package taskfile

type Task struct {
	// Name is the task identifier
	Name string `json:"name"`

	// NeedOf makes this task an "implicit" need of other tasks based on the label selector
	// similar to that of a "pattern rule" in a Makefile
	NeedOf Selector `json:"needOf,omitempty"`

	// Labels are intended to be used to specify identifying attributes of tasks that are
	// meaningful and relevant to users, but do not directly imply semantics to the core system.
	// Labels can be used to organize and to select subsets of tasks.
	Labels map[string]string `json:"labels,omitempty"`

	// Short is the short description shown in the 'help' output.
	Short string `json:"short,omitempty"`

	// Long is the long message shown in the 'help <this-task>' output.
	Long string `json:"long,omitempty"`

	// Example is examples of how to use the task.
	Example string `json:"example,omitempty"`

	// Steps are the things that will run in serial
	Do Steps `json:"steps,omitempty"`

	// Dir is the directory in which steps will be executed from (default: Taskfile directory)
	Dir string `json:"dir,omitempty"`

	// Vars are values that are evaluated/available for use elsewhere in this task.
	// Vars declared here will take precedence over Vars declared at the `Taskfile` (global) level.
	// Vars are _lazily_ evaluated.
	//
	// Late binding embodies some features of a lazy language.
	// Errors do not occur unless a field is actually dereferenced, and cyclic structures can be created.
	// For example the following is valid even in an eager version of the language:
	//
	//  local x = {a: "apple", b: y.b},
	//  y = {a: x.a, b: "banana"};
	//  x
	//
	// It would therefore be confusing if the following was not also valid, which leads us to lazy semantics for arrays.
	//
	//  local x = ["apple", y[1]],
	//  y = [x[0], "banana"];
	//  x
	//
	// Therefore, for consistency, the whole language is lazy. It does not harm the language to be lazy:
	// Performance is not significantly affected, stack traces are still possible,
	// and it doesn't interfere with I/O (because there is no I/O).
	// There is also a precedent for laziness, e.g. in Makefiles and the Nix expression language.
	//
	// Arguably, laziness brings real benefits in terms of abstraction and modularity.
	// It is possible to build infinite data-structures, and there is unrestricted beta expansion.
	// For example, the following 2 snippets of code are only equivalent in a lazy language.
	//
	// if x == 0 then 0 else if x > 0 then 1 / x else -1/x
	// local r = 1 / x;
	// if x == 0 then 0 else if x > 0 then r else -r
	Vars Vars `json:"vars,omitempty"`

	// Env is a map of environment variables that are available to all steps in the task.
	// You can also set environment variables for the entire Taskfile or an individual step.
	//
	// When more than one environment variable is defined with the same name (in different locations),
	// Task uses the most specific environment variable.
	//
	// For example, an environment variable defined in a step will override task and global (Taskfile) variables
	// with the same name, while the step executes.
	// An environment variable defined for a task will override a global (Taskfile) variable
	// with the same name, while the task executes.
	Env Vars `json:"env,omitempty"`

	// Needs are other tasks that this task depends on. A directed acyclic graph is built based on dependencies.
	// Independent needs are executed in parallel.
	Needs Needs `json:"needs,omitempty"`

	// SerialGroups is an array of arbitrary tag-like strings. Executions of this task
	// and other tasks referencing the same tags will be serialized.
	SerialGroups SerialGroups `json:"serialGroups,omitempty"`

	// From matched files have their contents hashed.
	// If the content hashes change between runs, this task will be marked as "out-of-date".
	From FileGlobs `json:"from,omitempty"`

	// Makes matched files are checked to exist.
	// If these files do not exist, this task will be marked as "out-of-date".
	// When providing a glob, only 1 match is required to keep this task "up-to-date".
	Makes FileGlobs `json:"makes,omitempty"`

	// Matrix is a step/task hook to define a matrix of step/task configurations.
	//
	// A matrix allows you to create multiple steps/tasks by performing variable substitution
	// in a single step/task definition.
	//
	// When defined on a Task, it dynamically creates several tasks, suffixing the task name with the matrix values.
	//
	// Each individual task can be called. The original task becomes a "virtual task" that "needs" the matrix of tasks.
	// Allowing you to call all matrix tasks from the single "parent" virtual task.
	//
	// For example, you can use a matrix to create tasks for more than one supported version of a programming language,
	// variable values, or tool, etc. A matrix reuses the step/task's configuration
	// and creates a step/task for each matrix you configure.
	//
	// Example
	//
	//	task:
	//	  echo:
	//		matrix:
	//		- env: ["dev", "staging", "prod"]
	//		steps:
	//		- run: "echo {{.Matrix.env}}"
	//
	// Tasks defined are:
	//  echo
	//  echo:dev
	//  echo:staging
	//  echo:prod
	Matrix Matrix `json:"matrix,omitempty"`

	// Timeout is the umbrella bounding time limit (duration) for the task before signalling for termination via SIGINT.
	Timeout Duration `json:"timeout,omitempty"`

	// GracefulTermination is the bounding time limit (duration) for this task before sending subprocesses a SIGKILL.
	GracefulTermination Duration `json:"gracefulTermination,omitempty"`

	// OnSuccess is a step/task hook for a series of steps to run if the current step/task succeeds.
	OnSuccess Steps `json:"onSuccess,omitempty"`

	// OnFailure is a step/task hook for a series of steps to run if the current step/task fails.
	OnFailure Steps `json:"onFailure,omitempty"`

	// Defer is a step/task hook that defers a series of steps until
	// the current step/task execution completes, whether success or failure.
	Defer Steps `json:"defer,omitempty"`

	// IfSpec or just `if` in the Taskfile yaml, defines rules for whether this task should run.
	IfSpec IfSpec `json:"if,omitempty"`
}
