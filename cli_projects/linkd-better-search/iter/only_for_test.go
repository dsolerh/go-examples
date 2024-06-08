package iter

import (
	"reflect"
	"testing"
)

func equals[T any](t *testing.T, a, b T) {
	if !reflect.DeepEqual(a, b) {
		t.Fatalf("%+v != %+v", a, b)
	}
}

func hasNextValue[T any](t *testing.T, iter Iter[T], val T) {
	gotVal, end := iter()
	if !end {
		t.Fatal("unexpected end of the iterator")
	} else {
		equals(t, gotVal, val)
	}
}

func noNextValue[T any](t *testing.T, iter Iter[T]) {
	_, end := iter()
	isTrue(t, !end)
}

func isTrue(t *testing.T, a bool) {
	if !a {
		t.Fatalf("expected 'true', but got: %+v", a)
	}
}
