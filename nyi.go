package errs

import "fmt"

type nyiError string

func NotYetImplemented(feature string) nyiError { return nyiError(feature) }

func (e nyiError) Error() string { return fmt.Sprintf("NOT-YET-IMPLEMENTED: %s", string(e)) }
