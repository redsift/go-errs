package errs

import (
	"errors"
	"fmt"
	"strings"

	"github.com/redsift/go-foodfans"
	"go.opentelemetry.io/otel/attribute"
)

//go:generate go tool msgp -io=false
//msgp:ignore Retry RetryIncrement RetryFlag
type Retry bool          // error is retryable
type RetryIncrement bool // error should result in consuming a retry attempt
type RetryFlag bool      // retry should be made visible to the sift

func IsCode(err error, code InternalState) bool {
	if err == nil {
		return false
	}

	cast, ok := errors.AsType[*PropagatedError](err)
	if !ok {
		return false
	}

	return cast.Code == code
}

// ContainsCode reports whether any error in err's chain, including [errors.Join]
// trees, is a [PropagatedError] with the given code. Unlike [IsCode], which only
// inspects the outermost [PropagatedError], a code that has since been wrapped
// under a different one is still found.
func ContainsCode(err error, code InternalState) (found bool) {
	visitRecursive(err, func(cast *PropagatedError) bool {
		found = found || cast.Code == code
		return !found
	})
	return found
}

// CollectCodes returns the code of every [PropagatedError] in err's chain,
// including [errors.Join] trees, outermost first. Codes are not deduplicated,
// so a code applied at several layers appears several times. The result is nil
// when err carries no [PropagatedError].
func CollectCodes(err error) (out []InternalState) {
	if err == nil {
		return nil
	}
	visitRecursive(err, func(err *PropagatedError) bool {
		out = append(out, err.Code)
		return true
	})
	return
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

	cast, ok := errors.AsType[*PropagatedError](err)
	if !ok {
		return false, false, false
	}

	return cast.RetryWithIncrementAndFlag()
}

func RetryError(err error) bool {
	if err == nil {
		return false
	}

	cast, ok := errors.AsType[*PropagatedError](err)
	if !ok {
		return false
	}

	return cast.Retry()
}

func AerospikeError(err error) bool {
	if err == nil {
		return false
	}

	cast, ok := errors.AsType[*PropagatedError](err)
	if !ok {
		return false
	}

	return cast.Aerospike()
}

func NodeTimeoutError(err error) bool {
	if err == nil {
		return false
	}

	cast, ok := errors.AsType[*PropagatedError](err)
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

	if cast, ok := errors.AsType[*PropagatedError](err); ok && cast.Code == code {
		return err
	}

	return wrapWithCode(code, err)
}

func wrapWithCode(code InternalState, err error) *PropagatedError {
	return &PropagatedError{
		Id:     foodfans.New(),
		Code:   code,
		Title:  code.Message(),
		Detail: err.Error(),
		Link:   code.LookupURL(),
		Status: 500,
		cause:  err,
	}
}

func WrapAsParameterError(param string) error {
	perr := wrapWithCode(Cappuccino, fmt.Errorf("Parameter error: %q", param))
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

// Attributes returns OpenTelemetry attributes for this error.
// It always includes the error code as "redsift.error.code" and
// includes "http.status_code" when Status is non-zero.
func (s *PropagatedError) Attributes() []attribute.KeyValue {
	if s == nil {
		return nil
	}
	attrs := make([]attribute.KeyValue, 0, 2)
	attrs = append(attrs, attribute.String("redsift.error.code", s.Code.String()))
	if s.Status != 0 {
		attrs = append(attrs, attribute.Int("http.status_code", s.Status))
	}
	return attrs
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
	case
		Bicerin,    // overloaded
		Kopitubruk, // nanomsg transport error
		Latte,      // service shutting down
		Macchiato,  // explicit retry requested
		Mochasippi: // service unavailable
		return true, false, false
	case Flatwhite: // service shutting down, request _may_ have been visible to the service
		return true, false, true
	case Lungo: // deadline exceeded, work _may_ have partially run
		return true, true, true
	default:
		return false, false, false
	}
}

func visitRecursive[T error](err error, visitor func(T) bool) bool {
	if cast, ok := err.(T); ok {
		if !visitor(cast) {
			return false
		}
	}

	switch e := err.(type) {
	case interface{ Unwrap() error }:
		return visitRecursive(e.Unwrap(), visitor)

	case interface{ Unwrap() []error }:
		for _, err := range e.Unwrap() {
			if !visitRecursive(err, visitor) {
				return false
			}
		}
	}

	return true
}
