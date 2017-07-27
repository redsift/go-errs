package errs

import (
	"fmt"

	"github.com/redsift/go-foodfans"
)

const (
	reset = "\033[0m"
	cyan  = "\033[36m"
)

func WrapErrors(errs []*PropagatedError) *PropagatedError {
	//TODO: Work this through
	return errs[0]
}

func WrapWithCode(code InternalState, err error) *PropagatedError {
	if err == nil {
		return nil
	}

	if cast, ok := err.(*PropagatedError); ok {
		return cast
	}

	id := foodfans.New()
	message := code.Message()
	detail := err.Error()
	link := code.LookupURL()

	return &PropagatedError{id, code, message, detail, link, nil, 500}
}

func WrapAsParameterError(param string) *PropagatedError {
	perr := WrapWithCode(Cappuccino, fmt.Errorf("Parameter error: %q", param))
	perr.Source = &ErrorSource{"", param}
	return perr
}

func WrapAsConfigIssue(err error) *PropagatedError {
	return WrapWithCode(Affogato, err)
}

func WrapAsAssert(err error) *PropagatedError {
	return WrapWithCode(Yuanyang, err)
}

func Wrap(err error) *PropagatedError {
	return WrapWithCode(stateForErr(err), err)
}

func stateForErr(err error) InternalState {
	if err == nil {
		return None
	}

	return Unknown
}

type ErrorSource struct {
	Pointer   string `json:"pointer"`
	Parameter string `json:"parameter"`
}

type PropagatedError struct {
	Id     string        `json:"id"`
	Code   InternalState `json:"code"`
	Title  string        `json:"title"`
	Detail string        `json:"detail"`
	Link   string        `json:"link"`
	Source *ErrorSource  `json:"source"`
	Status int           `json:"-"`
}

func (s *ErrorSource) String() string {
	if s == nil {
		return ""
	}

	if s.Pointer != "" {
		return "json:" + s.Pointer
	}

	return s.Parameter
}

func (s *PropagatedError) Error() string {
	return fmt.Sprintf("[id:"+cyan+"%s"+reset+"] %s / %s: %s, %s", s.Id, s.Code, s.Title, s.Detail, s.Source)
}

func (s *PropagatedError) Reason() string {
	return fmt.Sprintf("[id:%s] %s / %s: %s, %s", s.Id, s.Code, s.Title, s.Detail, s.Source)
}

func (s *PropagatedError) StatusCode() int {
	return s.Status
}

func (pe *PropagatedError) Retry() bool {
	if pe == nil {
		return false
	}

	switch pe.Code {
	case Kopitubruk, Macchiato, Turkish, Mocha:
		return true
	default:
		return false
	}
}

func (pe *PropagatedError) Aerospike() bool {
	if pe == nil {
		return false
	}

	switch pe.Code {
	case Turkish, Mocha:
		return true
	default:
		return false
	}
}
