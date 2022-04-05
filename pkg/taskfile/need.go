package taskfile

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-multierror"
)

type Needs struct {
	needs []Need
	tpl   string
}

type Need struct{}

func (n *Needs) UnmarshalJSON(data []byte) error {
	tmp := []Need{}

	errResult := &multierror.Error{}

	// first try to unmarshal simply as a list of needs
	listUnmarshalErr := json.Unmarshal(data, &tmp)
	if listUnmarshalErr == nil {
		*n = Needs{
			needs: tmp,
		}
		return nil
	}
	errResult = multierror.Append(errResult, fmt.Errorf("unmarshalling needs as list: %w", listUnmarshalErr))

	// if the user provided a string, it means that they are _dynamically_ (via templating) providing their list
	// so, let's hold on to it, but we don't need to evaluate it yet
	var tpl string
	stringUnmarshalErr := json.Unmarshal(data, &tpl)
	if stringUnmarshalErr == nil {
		*n = Needs{
			tpl: tpl,
		}
		return nil
	}
	errResult = multierror.Append(errResult, fmt.Errorf("unmarshalling needs as string: %w", stringUnmarshalErr))

	return errResult.ErrorOrNil()
}
