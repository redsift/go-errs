package errs

import (
	"errors"
	"fmt"
)

func Errorf(format string, args ...any) error {
	var perr *PropagatedError

	for i, arg := range args {
		err, ok := arg.(error)
		if !ok {
			continue
		}
		perr, ok = err.(*PropagatedError)
		if !ok {
			continue
		}

		args[i] = errors.New(perr.Detail)
	}

	if perr == nil {
		return fmt.Errorf(format, args...)
	}

	return &PropagatedError{
		Id:     perr.Id,
		Code:   perr.Code,
		Title:  perr.Title,
		Detail: fmt.Errorf(format, args...).Error(),
		Link:   perr.Link,
		Source: perr.Source,
		Status: perr.Status,
	}
}
