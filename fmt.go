package errs

import (
	"fmt"
)

// Errorf is now an alias for fmt.Errorf
//
// Deprecated: no longer needed now that PropagatedError implements Unwrap
func Errorf(format string, args ...any) error {
	return fmt.Errorf(format, args...)
}
