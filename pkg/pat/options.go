package pat

import (
	"fmt"

	"github.com/hashicorp/go-multierror"
)

func ApplyOptions[U ~func(*T) error, V ~[]U, T any](options *T, optFns V) error {
	var errResult *multierror.Error

	for _, o := range optFns {
		errResult = multierror.Append(errResult, o(options))
	}

	err := errResult.ErrorOrNil()
	if err != nil {
		return fmt.Errorf("applying options: %w", err)
	}

	return nil
}

type Option[T any] func(T) error
