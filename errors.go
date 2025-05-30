package errs

import (
	"fmt"
	"strings"

	"github.com/redsift/go-foodfans"
)

//go:generate msgp -io=false
//msgp:ignore Retry RetryIncrement RetryFlag
type Retry bool
type RetryIncrement bool
type RetryFlag bool

func IsCode(err error, code InternalState) bool {
	if err == nil {
		return false
	}

	cast, ok := err.(*PropagatedError)
	if !ok {
		return false
	}

	return cast.Code == code
}

// RetryWithCounter returns a bool indicating retry,
// and a counter which increments if applicable
func RetryWithCounter(err error, n int) (Retry, int) {
	retry, retryIncr, _ := RetryWithIncrementAndFlag(err)
	ctr := n
	if bool(retryIncr) {
		ctr++
	}
	return retry, ctr
}

// RetryWithIncrement returns a bool indicating retry,
// and a bool indicating if the counter should be incremented (based on error)
func RetryWithIncrement(err error) (Retry, RetryIncrement) {
	retry, retryIncr, _ := RetryWithIncrementAndFlag(err)
	return retry, retryIncr
}

// RetryWithIncrementAndFlag returns a bool indicating retry,
// a bool indicating if the counter should be incremented (based on error),
// and a bool indicating retryFlag
func RetryWithIncrementAndFlag(err error) (Retry, RetryIncrement, RetryFlag) {
	if err == nil {
		return false, false, false
	}

	cast, ok := err.(*PropagatedError)
	if !ok {
		return false, false, false
	}

	return cast.RetryWithIncrementAndFlag()
}

func RetryError(err error) bool {
	if err == nil {
		return false
	}

	cast, ok := err.(*PropagatedError)
	if !ok {
		return false
	}

	return cast.Retry()
}

func AerospikeError(err error) bool {
	if err == nil {
		return false
	}

	cast, ok := err.(*PropagatedError)
	if !ok {
		return false
	}

	return cast.Aerospike()
}

func NodeTimeoutError(err error) bool {
	if err == nil {
		return false
	}

	cast, ok := err.(*PropagatedError)
	if !ok {
		return false
	}

	return cast.NodeTimeout()
}

func WrapErrors(errs []*PropagatedError) error {
	//TODO: Work this through
	return errs[0]
}

func WrapWithCode(code InternalState, err error) error {
	if err == nil {
		return nil
	}

	if cast, ok := err.(*PropagatedError); ok {
		if cast.Code == code {
			return cast
		}
	}

	id := foodfans.New()
	message := code.Message()
	detail := err.Error()
	link := code.LookupURL()

	return &PropagatedError{Id: id, Code: code, Title: message, Detail: detail, Link: link, Status: 500, cause: err}
}

func WrapAsParameterError(param string) error {
	//goland:noinspection GoTypeAssertionOnErrors
	perr := WrapWithCode(Cappuccino, fmt.Errorf("Parameter error: %q", param)).(*PropagatedError)
	perr.Source = &ErrorSource{"", param}
	return perr
}

func WrapAsConfigIssue(err error) error {
	return WrapWithCode(Affogato, err)
}

func WrapAsAssert(err error) error {
	return WrapWithCode(Yuanyang, err)
}

func Wrap(err error) error {
	return WrapWithCode(stateForErr(err), err)
}

func stateForErr(err error) InternalState {
	if err == nil {
		return None
	}

	return Unknown
}

type ErrorSource struct {
	Pointer   string `json:"pointer" msg:"pointer"`
	Parameter string `json:"parameter" msg:"parameter"`
}

type PropagatedError struct {
	Id     string        `json:"id" msg:"id"`
	Code   InternalState `json:"code" msg:"code"`
	Title  string        `json:"title" msg:"title"`
	Detail string        `json:"detail" msg:"detail"`
	Link   string        `json:"link" msg:"link"`
	Source *ErrorSource  `json:"source" msg:"source"`
	Status int           `json:"-" msg:"-"`
	cause  error
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
	if s == nil {
		return ""
	}

	return fmt.Sprintf("[id:%s] %s / %s: %s, %s", s.Id, s.Code, s.Title, s.Detail, s.Source)
}

// Unwrap returns the cause of the error if it was wrapped
func (s *PropagatedError) Unwrap() error {
	return s.cause
}

func (s *PropagatedError) StatusCode() int {
	return s.Status
}

func (pe *PropagatedError) Retry() bool {
	if pe == nil {
		return false
	}

	retry, _ := pe.RetryWithIncrement()
	return bool(retry)
}

func (pe *PropagatedError) Aerospike() bool {
	if pe == nil {
		return false
	}

	switch pe.Code {
	case Turkish, Mocha:
		// Aerospike "Record too big" error is not worth retrying
		if strings.Contains(pe.Detail, "Record too big") {
			return false
		}

		return true
	default:
		return false
	}
}

func (pe *PropagatedError) Mongo() bool {
	if pe == nil {
		return false
	}

	if pe.Code == Guillermo {
		return true
	}

	return false
}

func (pe *PropagatedError) NodeTimeout() bool {
	if pe == nil {
		return false
	}

	if pe.Code == Espresso {
		return true
	}

	return false
}

func (pe *PropagatedError) RetryWithIncrement() (Retry, RetryIncrement) {
	if pe == nil {
		return false, false
	}

	retry, retryIncr, _ := pe.RetryWithIncrementAndFlag()
	return retry, retryIncr
}

func (pe *PropagatedError) RetryWithIncrementAndFlag() (Retry, RetryIncrement, RetryFlag) {
	if pe == nil {
		return false, false, false
	}

	if pe.Aerospike() {
		return true, false, true
	}

	if pe.Mongo() {
		return true, false, true
	}

	if pe.NodeTimeout() {
		return true, true, true
	}

	switch pe.Code {
	case Kopitubruk, Macchiato:
		return true, false, false
	default:
		return false, false, false
	}
}
