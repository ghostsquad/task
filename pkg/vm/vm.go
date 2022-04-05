package vm

import (
	"context"
	"fmt"

	"github.com/go-task/task/v4/pkg/pat"
	"github.com/go-task/task/v4/pkg/taskfile"
)

// VM is the core interpreter and is the touchpoint used to parse and execute a Taskfile.
type VM struct {
	f *taskfile.Taskfile
}

type VMOptions struct {
}

// New compiles the yaml of a Task file into the Taskfile API object
func New(yamlSource string, options ...pat.Option[*VMOptions]) (*VM, error) {
	opts := &VMOptions{}

	err := pat.ApplyOptions[pat.Option[*VMOptions]](opts, options)
	if err != nil {
		return nil, err
	}

	return &VM{}, nil
}

// Run evaluates
func (vm *VM) Run(ctx context.Context, name string) error {
	if t, ok := vm.f.Tasks[name]; ok {
		return vm.RunT(ctx, t)
	}

	return fmt.Errorf("unknown task: %s", name)
}

func (vm *VM) RunT(ctx context.Context, t *taskfile.Task) error {
	return nil
}
