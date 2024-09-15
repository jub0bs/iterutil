package iterutil_test

import (
	"fmt"
	"iter"
	"testing"
)

func assertEqual[E comparable](
	t *testing.T,
	got iter.Seq[E],
	want []E,
	breakWhen func(E) bool,
) {
	t.Helper()
	var es []E
	var i int
	for e := range got {
		if breakWhen(e) {
			return
		}
		es = append(es, e)
		if len(want) <= i {
			t.Fatalf("too many elements: got %v...; want %v", es, want)
		}
		if e != want[i] {
			t.Fatalf("got %v...; want %v...", es, want)
		}
		i++
	}
	// i should now be equal to len(want)
	if i != len(want) {
		t.Fatalf("not enough elements: got %v; want %v...", es, want)
	}
}

func alwaysFalse[E any](_ E) bool {
	return false
}

func equal[E comparable](target E) func(E) bool {
	return func(e E) bool {
		return e == target
	}
}

func assertEqual2[K, V comparable](
	t *testing.T,
	got iter.Seq2[K, V],
	want []Pair[K, V],
	breakWhen func(K, V) bool,
) {
	t.Helper()
	var pairs []Pair[K, V]
	var i int
	for k, v := range got {
		if breakWhen(k, v) {
			return
		}
		pairs = append(pairs, Pair[K, V]{k, v})
		if len(want) <= i {
			t.Fatalf("too many pairs: got %v...; want %v", pairs, want)
		}
		if k != want[i].k || v != want[i].v {
			t.Fatalf("got %v...; want %v...", pairs, want)
		}
		i++
	}
	// i should now be equal to len(want)
	if i != len(want) {
		t.Fatalf("not enough pairs: got %v; want %v...", pairs, want)
	}
}

type Pair[K, V any] struct {
	k K
	v V
}

func (p Pair[K, V]) String() string {
	return fmt.Sprintf("(%v,%v)", p.k, p.v)
}

// falseAfterN, n is non-negative, returns a function that returns
// true for the first n invocations and
// false for subsequent invocations,
// regardless of the value of its argument;
// otherwise, it returns a function that invariably returns false.
func falseAfterN[E any](n int) func(E) bool {
	if n < 0 {
		return func(E) bool {
			return true
		}
	}
	var count int
	return func(E) bool {
		if count < n {
			count++
			return true
		}
		return false
	}
}

func alwaysFalse2[K, V any](_ K, _ V) bool {
	return false
}

func equal2[K, V comparable](key K, value V) func(K, V) bool {
	return func(k K, v V) bool {
		return k == key && v == value
	}
}
