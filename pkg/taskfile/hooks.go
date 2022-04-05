package taskfile

import (
	"encoding/json"
	"fmt"
	"time"
)

// Matrix is a step/task hook to define a matrix of step/task configurations.
// A matrix allows you to create multiple steps/tasks by performing variable substitution
// in a single step/task definition.
type Matrix struct {
	KVs KeyValues `json:"kvs,omitempty"`

	// Includes lets you add additional configuration options to a build matrix step/task that already exists.
	Includes KeyValues `json:"includes,omitempty"`
}

type KeyValues struct {
	Key    string   `json:"key"`
	Values []string `json:"values"`
}

// Duration uses time.ParseDuration (see https://pkg.go.dev/time#ParseDuration) for unmarshalling.
type Duration time.Duration

// MarshalJSON implements the json.Marshaler interface.
func (t Duration) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Duration(t).String())
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (t *Duration) UnmarshalJSON(data []byte) error {
	var txt string
	err := json.Unmarshal(data, &txt)
	if err != nil {
		return fmt.Errorf("unmarshalling timeout: %w", err)
	}

	d, err := time.ParseDuration(txt)
	if err != nil {
		return err
	}

	*t = Duration(d)
	return nil
}

// IfSpec is a predicate for the purpose of deciding whether to run this step/task
// The result of evaluation of Eval, Sh, and NeedsFulfilled conditions are OR'd.
type IfSpec struct {
	// Eval can be used with the go templating in order make a decision
	// we use strconv.ParseBool to evaluate the result
	Eval string `json:"eval,omitempty"`

	// EvalLenient allows Eval to return "false" in the case that
	// strconv.ParseBool returns an error.
	// By default, an error produced during Eval will cause a crash.
	EvalLenient bool `json:"evalStrict,omitempty"`

	// Sh defines a command to run, returns true if the command returns exit code 0, otherwise returns false
	Sh string `json:"sh,omitempty"`

	// NeedsFulfilled is condition for a task that has "needs".returns true if any needs of the current task must run. Has no effect on steps.
	// When this is set to false (default), if any need is unfulfilled, this returns true.
	// When this is set to true, unfulfilled needs will not be used to determine if this task should run.
	// In other words, even if a need of this task were to be "out-of-date", this current task is still eligible to be "up-to-date"
	// based on the other conditions of this if spec.
	//
	// This has no effect for steps.
	NeedsFulfilled bool `json:"needsUnfulfilled,omitempty"`
}
