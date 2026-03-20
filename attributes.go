package errs

import (
	"fmt"
	"strings"

	"go.opentelemetry.io/otel/attribute"
)

// ErrorWithAttributes is an error that carries OpenTelemetry attributes.
// Use [errors.AsType] to extract attributes from an error chain.
type ErrorWithAttributes interface {
	error
	Attributes() []attribute.KeyValue
}

type errorWithAttributes struct {
	underlying error
	attributes []attribute.KeyValue
}

// Error returns the underlying error message followed by a formatted
// table of key-value attributes.
func (e errorWithAttributes) Error() string {
	maxKeyLen := 0
	for _, attr := range e.attributes {
		maxKeyLen = max(maxKeyLen, len(attr.Key))
	}
	var b strings.Builder
	b.Grow(64 * len(e.attributes))
	b.WriteString(e.underlying.Error())
	format := fmt.Sprintf("\n%%-%ds: %%v", maxKeyLen)
	for _, attr := range e.attributes {
		fmt.Fprintf(&b, format, attr.Key, attr.Value.AsInterface())
	}
	return b.String()
}

// Unwrap returns the underlying error, allowing [errors.Is] and
// [errors.AsType] to traverse the error chain.
func (e errorWithAttributes) Unwrap() error {
	return e.underlying
}

// Attributes returns the OpenTelemetry attributes attached to this error.
func (e errorWithAttributes) Attributes() []attribute.KeyValue {
	return e.attributes
}

// WrapWithAttributes wraps err with the given OpenTelemetry attributes.
// If err is nil or no attributes are provided, err is returned unchanged.
// Wrapping an error that already carries attributes nests rather than merges;
// use [errors.AsType] to collect attributes from the full error chain.
func WrapWithAttributes(err error, attrs ...attribute.KeyValue) error {
	if err == nil || len(attrs) == 0 {
		return err
	}
	return errorWithAttributes{
		underlying: err,
		attributes: attrs,
	}
}

// GatherAttributesRecursive walks the full error chain (including [errors.Join] trees)
// and collects [attribute.KeyValue] from every error implementing [ErrorWithAttributes].
// Attributes are returned bottom-up: innermost errors first, outermost last.
func GatherAttributesRecursive(err error) (attrs []attribute.KeyValue) {
	if u, ok := err.(interface{ Unwrap() []error }); ok {
		for _, one := range u.Unwrap() {
			attrs = append(attrs, GatherAttributesRecursive(one)...)
		}
	} else if u, ok := err.(interface{ Unwrap() error }); ok {
		attrs = append(attrs, GatherAttributesRecursive(u.Unwrap())...)
	}
	if aerr, ok := err.(ErrorWithAttributes); ok {
		attrs = append(attrs, aerr.Attributes()...)
	}
	return attrs
}
