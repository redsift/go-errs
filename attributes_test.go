package errs

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"go.opentelemetry.io/otel/attribute"
)

func TestErrorWithAttributes_Error(t *testing.T) {
	base := fmt.Errorf("something broke")
	wrapped := WrapWithAttributes(base,
		attribute.String("key1", "val1"),
		attribute.Int("longer_key", 42),
	)

	msg := wrapped.Error()

	if msg == base.Error() {
		t.Fatal("expected error message to include attributes")
	}
	for _, want := range []string{"something broke", "key1", "val1", "longer_key", "42"} {
		if !strings.Contains(msg, want) {
			t.Errorf("error message %q missing %q", msg, want)
		}
	}
}

func TestErrorWithAttributes_Unwrap(t *testing.T) {
	base := fmt.Errorf("root cause")
	wrapped := WrapWithAttributes(base, attribute.String("k", "v"))

	if !errors.Is(wrapped, base) {
		t.Error("errors.Is should find the underlying error")
	}
}

func TestErrorWithAttributes_Attributes(t *testing.T) {
	attrs := []attribute.KeyValue{
		attribute.String("a", "1"),
		attribute.Int("b", 2),
	}
	wrapped := WrapWithAttributes(fmt.Errorf("err"), attrs...)

	ewa, ok := errors.AsType[ErrorWithAttributes](wrapped)
	if !ok {
		t.Fatal("expected to extract ErrorWithAttributes from chain")
	}

	got := ewa.Attributes()
	if len(got) != len(attrs) {
		t.Fatalf("got %d attributes, want %d", len(got), len(attrs))
	}
	for i := range got {
		if got[i] != attrs[i] {
			t.Errorf("attribute[%d] = %v, want %v", i, got[i], attrs[i])
		}
	}
}

func TestWrapWithAttributes_NilError(t *testing.T) {
	if got := WrapWithAttributes(nil, attribute.String("k", "v")); got != nil {
		t.Errorf("expected nil, got %v", got)
	}
}

func TestWrapWithAttributes_NoAttributes(t *testing.T) {
	base := fmt.Errorf("err")
	if got := WrapWithAttributes(base); got != base {
		t.Error("expected original error returned when no attributes provided")
	}
}

func TestWrapWithAttributes_ChainPreservesLayers(t *testing.T) {
	inner := WrapWithAttributes(fmt.Errorf("inner"),
		attribute.String("layer", "inner"),
	)
	outer := WrapWithAttributes(inner,
		attribute.String("layer", "outer"),
	)

	// Outer layer
	ewa, ok := errors.AsType[ErrorWithAttributes](outer)
	if !ok {
		t.Fatal("expected ErrorWithAttributes on outer")
	}
	if ewa.Attributes()[0].Value.AsString() != "outer" {
		t.Error("outer layer should have 'outer' attribute")
	}

	// Inner layer reachable via Unwrap
	innerErr := errors.Unwrap(outer)
	ewaInner, ok := errors.AsType[ErrorWithAttributes](innerErr)
	if !ok {
		t.Fatal("expected ErrorWithAttributes on inner")
	}
	if ewaInner.Attributes()[0].Value.AsString() != "inner" {
		t.Error("inner layer should have 'inner' attribute")
	}
}

func TestPropagatedError_Attributes(t *testing.T) {
	t.Run("with status", func(t *testing.T) {
		pe := &PropagatedError{
			Code:   Espresso,
			Status: 503,
		}
		attrs := pe.Attributes()

		if len(attrs) != 2 {
			t.Fatalf("got %d attributes, want 2", len(attrs))
		}
		if attrs[0].Key != "redsift.error.code" || attrs[0].Value.AsString() != Espresso.String() {
			t.Errorf("attrs[0] = %v, want redsift.error.code=%s", attrs[0], Espresso.String())
		}
		if attrs[1].Key != "http.status_code" || attrs[1].Value.AsInt64() != 503 {
			t.Errorf("attrs[1] = %v, want http.status_code=503", attrs[1])
		}
	})

	t.Run("without status", func(t *testing.T) {
		pe := &PropagatedError{
			Code: Latte,
		}
		attrs := pe.Attributes()

		if len(attrs) != 1 {
			t.Fatalf("got %d attributes, want 1", len(attrs))
		}
		if attrs[0].Key != "redsift.error.code" {
			t.Errorf("attrs[0].Key = %q, want redsift.error.code", attrs[0].Key)
		}
	})

	t.Run("nil receiver", func(t *testing.T) {
		var pe *PropagatedError
		if attrs := pe.Attributes(); attrs != nil {
			t.Errorf("expected nil, got %v", attrs)
		}
	})
}

func TestPropagatedError_ImplementsErrorWithAttributes(t *testing.T) {
	pe := WrapWithCode(Cappuccino, fmt.Errorf("test")).(*PropagatedError)
	pe.Status = 400

	var ewa ErrorWithAttributes = pe // compile-time check
	attrs := ewa.Attributes()

	if len(attrs) != 2 {
		t.Fatalf("got %d attributes, want 2", len(attrs))
	}
}

func TestPropagatedError_ExtractableFromChain(t *testing.T) {
	pe := WrapWithCode(Latte, fmt.Errorf("base")).(*PropagatedError)
	wrapped := fmt.Errorf("context: %w", pe)

	ewa, ok := errors.AsType[ErrorWithAttributes](wrapped)
	if !ok {
		t.Fatal("expected to extract ErrorWithAttributes from wrapped PropagatedError")
	}
	attrs := ewa.Attributes()
	if len(attrs) == 0 {
		t.Fatal("expected at least one attribute")
	}
	if attrs[0].Value.AsString() != Latte.String() {
		t.Errorf("got code %q, want %q", attrs[0].Value.AsString(), Latte.String())
	}
}

func TestGatherAttributesRecursive_NilError(t *testing.T) {
	attrs := GatherAttributesRecursive(nil)
	if attrs != nil {
		t.Errorf("expected nil, got %v", attrs)
	}
}

func TestGatherAttributesRecursive_PlainError(t *testing.T) {
	attrs := GatherAttributesRecursive(fmt.Errorf("plain"))
	if len(attrs) != 0 {
		t.Errorf("expected no attributes from plain error, got %d", len(attrs))
	}
}

func TestGatherAttributesRecursive_SingleLayer(t *testing.T) {
	err := WrapWithAttributes(fmt.Errorf("base"),
		attribute.String("k", "v"),
	)
	attrs := GatherAttributesRecursive(err)
	if len(attrs) != 1 {
		t.Fatalf("got %d attributes, want 1", len(attrs))
	}
	if attrs[0].Key != "k" || attrs[0].Value.AsString() != "v" {
		t.Errorf("got %v, want k=v", attrs[0])
	}
}

func TestGatherAttributesRecursive_NestedLayers(t *testing.T) {
	inner := WrapWithAttributes(fmt.Errorf("base"),
		attribute.String("layer", "inner"),
	)
	outer := WrapWithAttributes(inner,
		attribute.String("layer", "outer"),
	)

	attrs := GatherAttributesRecursive(outer)
	if len(attrs) != 2 {
		t.Fatalf("got %d attributes, want 2", len(attrs))
	}
	// Inner first (bottom-up), then outer
	if attrs[0].Value.AsString() != "inner" {
		t.Errorf("attrs[0] = %q, want inner", attrs[0].Value.AsString())
	}
	if attrs[1].Value.AsString() != "outer" {
		t.Errorf("attrs[1] = %q, want outer", attrs[1].Value.AsString())
	}
}

func TestGatherAttributesRecursive_PropagatedError(t *testing.T) {
	pe := &PropagatedError{Code: Espresso, Status: 503}
	attrs := GatherAttributesRecursive(pe)
	if len(attrs) != 2 {
		t.Fatalf("got %d attributes, want 2", len(attrs))
	}
	if attrs[0].Key != "redsift.error.code" {
		t.Errorf("attrs[0].Key = %q, want redsift.error.code", attrs[0].Key)
	}
	if attrs[1].Key != "http.status_code" {
		t.Errorf("attrs[1].Key = %q, want http.status_code", attrs[1].Key)
	}
}

func TestGatherAttributesRecursive_WrappedPropagatedError(t *testing.T) {
	pe := &PropagatedError{Code: Latte, Status: 500}
	wrapped := WrapWithAttributes(pe,
		attribute.String("service", "test"),
	)

	attrs := GatherAttributesRecursive(wrapped)
	// PropagatedError contributes 2 (code + status), wrapper contributes 1 (service)
	if len(attrs) != 3 {
		t.Fatalf("got %d attributes, want 3", len(attrs))
	}
	// Inner (PropagatedError) first, then outer (WrapWithAttributes)
	if attrs[0].Key != "redsift.error.code" {
		t.Errorf("attrs[0].Key = %q, want redsift.error.code", attrs[0].Key)
	}
	if attrs[1].Key != "http.status_code" {
		t.Errorf("attrs[1].Key = %q, want http.status_code", attrs[1].Key)
	}
	if attrs[2].Key != "service" {
		t.Errorf("attrs[2].Key = %q, want service", attrs[2].Key)
	}
}

func TestGatherAttributesRecursive_JoinedErrors(t *testing.T) {
	err1 := WrapWithAttributes(fmt.Errorf("a"), attribute.String("from", "err1"))
	err2 := WrapWithAttributes(fmt.Errorf("b"), attribute.String("from", "err2"))
	joined := errors.Join(err1, err2)

	attrs := GatherAttributesRecursive(joined)
	if len(attrs) != 2 {
		t.Fatalf("got %d attributes, want 2", len(attrs))
	}
	if attrs[0].Value.AsString() != "err1" {
		t.Errorf("attrs[0] = %q, want err1", attrs[0].Value.AsString())
	}
	if attrs[1].Value.AsString() != "err2" {
		t.Errorf("attrs[1] = %q, want err2", attrs[1].Value.AsString())
	}
}

func TestGatherAttributesRecursive_FmtWrapped(t *testing.T) {
	inner := WrapWithAttributes(fmt.Errorf("base"), attribute.String("k", "v"))
	wrapped := fmt.Errorf("context: %w", inner)

	attrs := GatherAttributesRecursive(wrapped)
	if len(attrs) != 1 {
		t.Fatalf("got %d attributes, want 1", len(attrs))
	}
	if attrs[0].Key != "k" {
		t.Errorf("attrs[0].Key = %q, want k", attrs[0].Key)
	}
}
