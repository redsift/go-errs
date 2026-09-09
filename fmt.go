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
		perr, ok = errors.AsType[*PropagatedError](err)
		if !ok {
			continue
		}

		if perr.cause != nil {
			args[i] = perr.cause
		} else {
			args[i] = errors.New(perr.Detail)
		}
	}

	if perr == nil {
		return fmt.Errorf(format, args...)
	}

	cause := fmt.Errorf(format, args...)

	return &PropagatedError{
		Id:     perr.Id,
		Code:   perr.Code,
		Title:  perr.Title,
		Detail: cause.Error(),
		Link:   perr.Link,
		Source: perr.Source,
		Status: perr.Status,
		cause:  cause,
	}
}
