package errs

import (
	"fmt"
	"strings"

	"github.com/redsift/go-foodfans"
)

const (
	reset = "\033[0m"
	cyan  = "\033[36m"
)

type Retry bool
type RetryIncrement bool
type RetryFlag bool

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

	return &PropagatedError{id, code, message, detail, link, nil, 500}
}

func WrapAsParameterError(param string) error {
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
	//TODO move color to zaputil
	return fmt.Sprintf("[id:"+cyan+"%s"+reset+"] %s / %s: %s, %s", s.Id, s.Code, s.Title, s.Detail, s.Source)
}

func (s *PropagatedError) Reason() string {
	// TODO remove when color went to zaputil
	return fmt.Sprintf("[id:%s] %s / %s: %s, %s", s.Id, s.Code, s.Title, s.Detail, s.Source)
}

func (s *PropagatedError) StatusCode() int {
	return s.Status
}

func (pe *PropagatedError) Retry() bool {
	if pe == nil {
		return false
	}

	if pe.Aerospike() || pe.NodeTimeout() {
		return true
	}

	switch pe.Code {
	case Kopitubruk, Macchiato:
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
		// Aerospike "Record too big" error is not worth retrying
		if strings.Contains(pe.Detail, "Record too big") {
			return false
		}

		return true
	default:
		return false
	}
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
