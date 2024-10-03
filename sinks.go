package iterutil

import (
	"cmp"
	"iter"

	"golang.org/x/exp/constraints"
)

// IsEmpty reports whether seq is an empty iterator.
func IsEmpty[E any](seq iter.Seq[E]) bool {
	for range seq {
		return false
	}
	return true
}

// Len returns the number of elements in seq.
// It terminates if and only if seq is finite.
func Len[E any](seq iter.Seq[E]) int {
	var n int
	for range seq {
		n++
	}
	return n
}

// At, if count is non-negative, returns
// the element at index n in seq and true
// or the zero value and false if seq contains fewer than count elements;
// otherwise, it panics.
func At[I constraints.Integer, E any](seq iter.Seq[E], n I) (e E, ok bool) {
	if n < 0 {
		panic("cannot be negative")
	}
	for v := range seq {
		if 0 < n {
			n--
			continue
		}
		e = v
		ok = true
		return
	}
	return
}

// Equal reports whether two iterators are equal:
// the same length and all elements equal.
// If the lengths are different, Equal returns false.
// Otherwise, the elements are compared sequentially,
// and the comparison stops at the first unequal pair.
// Floating point NaNs are not considered equal.
// Equal may not terminate if seq1 or seq2 or both are infinite.
func Equal[E comparable](seq1, seq2 iter.Seq[E]) bool {
	return EqualFunc(seq1, seq2, equal)
}

func equal[E comparable](e1, e2 E) bool { return e1 == e2 }

// EqualFunc reports whether two iterators are equal using eq as equality
// function on each pair of elements.
// If the lengths are different, EqualFunc returns false.
// Otherwise, the elements are compared sequentially,
// and the comparison stops at the first pair for which eq returns false.
// EqualFunc may not terminate if seq1 or seq2 or both are infinite.
func EqualFunc[A, B comparable](seq1 iter.Seq[A], seq2 iter.Seq[B], eq func(A, B) bool) bool {
	next1, stop1 := iter.Pull(seq1)
	defer stop1()
	next2, stop2 := iter.Pull(seq2)
	defer stop2()
	for {
		v1, ok1 := next1()
		v2, ok2 := next2()
		if !ok1 {
			return !ok2
		}
		if ok1 != ok2 || !eq(v1, v2) {
			return false
		}
	}
}

// Contains report whether target is present in seq.
// It may not terminate if seq is infinite.
func Contains[E comparable](seq iter.Seq[E], target E) bool {
	for e := range seq {
		if e == target {
			return true
		}
	}
	return false
}

// ContainsFunc reports whether at least one element e of seq satisfies p(e).
// It may not terminate if seq is infinite.
func ContainsFunc[E any](seq iter.Seq[E], p func(E) bool) bool {
	for e := range seq {
		if p(e) {
			return true
		}
	}
	return false
}

// Min, if seq is not empty, returns the minimal value in seq and true;
// otherwise, it returns the zero value and false.
// For floating-point numbers, Min propagates NaNs
// (any NaN value in seq forces the output to be NaN).
// Min terminates if and only if seq is finite.
func Min[E cmp.Ordered](seq iter.Seq[E]) (E, bool) {
	var (
		m         E
		firstSeen bool
	)
	for e := range seq {
		if !firstSeen {
			m = e
			firstSeen = true
			continue
		}
		m = min(e, m)
	}
	return m, firstSeen
}

// MinFunc, if seq is not empty, returns the minimal value
// (using cmp as comparison function) in seq and true;
// otherwise, it returns the zero value and false.
// If there is more than one minimal element according
// to the cmp function, MinFunc returns the first one.
// MinFunc terminates if and only if seq is finite.
func MinFunc[E any](seq iter.Seq[E], cmp func(E, E) int) (E, bool) {
	var (
		m         E
		firstSeen bool
	)
	for e := range seq {
		if !firstSeen {
			m = e
			firstSeen = true
			continue
		}
		if cmp(e, m) < 0 {
			m = e
		}
	}
	return m, firstSeen
}

// Max, if seq is not empty, returns the maximal value in seq and true;
// otherwise, it returns the zero value and false.
// For floating-point numbers, Max propagates NaNs
// (any NaN value in seq forces the output to be NaN).
// Max terminates if and only if seq is finite.
func Max[E cmp.Ordered](seq iter.Seq[E]) (E, bool) {
	var (
		m        E
		nonEmpty bool
	)
	for e := range seq {
		nonEmpty = true
		m = max(e, m)
	}
	return m, nonEmpty
}

// MaxFunc, if seq is not empty, returns the maximal value
// (using cmp as comparison function) in seq and true;
// otherwise, it returns the zero value and false.
// If there is more than one maximal element according
// to the cmp function, MaxFunc returns the first one.
// MaxFunc terminates if and only if seq is finite.
func MaxFunc[E any](seq iter.Seq[E], cmp func(E, E) int) (E, bool) {
	var (
		m        E
		nonEmpty bool
	)
	for e := range seq {
		nonEmpty = true
		if cmp(e, m) > 0 {
			m = e
		}
	}
	return m, nonEmpty
}

// Compare compares the elements of seq1 and seq2,
// using [cmp.Compare] on each pair of elements.
// The elements are compared sequentially until one element is not equal to
// the other.
// The result of comparing the first non-matching elements is returned.
// If seq1 and seq2 are equal until one of them ends,
// the shorter one is considered less than the longer one.
// The result is 0 if seq1 == seq2, -1 if seq1 < seq2, and +1 if seq1 > seq2.
// For floating-point types, a NaN is considered less than any non-NaN,
// and -0.0 is not less than (is equal to) 0.0.
// It may not terminate if seq1 or seq2 or both are infinite.
func Compare[E cmp.Ordered](seq1, seq2 iter.Seq[E]) int {
	return CompareFunc(seq1, seq2, cmp.Compare)
}

// CompareFunc is like [Compare] but uses a custom comparison function on each
// pair of elements.
// The result is the first non-zero result of cmp;
// if cmp always returns 0, the result is 0 if len(seq1) == len(seq2),
// -1 if len(seq1) < len(seq2), and +1 if len(seq1) > len(seq2).
// It may not terminate if seq1 or seq2 or both are infinite.
func CompareFunc[A, B any](seq1 iter.Seq[A], seq2 iter.Seq[B], cmp func(A, B) int) int {
	next1, stop1 := iter.Pull(seq1)
	defer stop1()
	next2, stop2 := iter.Pull(seq2)
	defer stop2()
	for {
		v1, ok1 := next1()
		v2, ok2 := next2()
		switch {
		case !ok1 && ok2:
			return -1
		case !ok1 && !ok2:
			return 0
		case ok1 && !ok2:
			return 1
		default:
			if c := cmp(v1, v2); c != 0 {
				return c
			}
		}
	}
}

// IsSorted reports whether seq is sorted in ascending order.
// For floating-point types, a NaN is considered less than any non-NaN,
// and -0.0 is not less than (is equal to) 0.0.
// It may not terminate if seq is infinite.
func IsSorted[E cmp.Ordered](seq iter.Seq[E]) bool {
	var (
		last      E
		firstSeen bool
	)
	for e := range seq {
		if firstSeen && cmp.Less(e, last) {
			return false
		}
		last = e
		firstSeen = true
	}
	return true
}

// IsSortedFunc reports whether seq is sorted in ascending order,
// using cmp as comparison function.
// It may not terminate if seq is infinite.
func IsSortedFunc[E any](seq iter.Seq[E], cmp func(E, E) int) bool {
	var (
		last      E
		firstSeen bool
	)
	for e := range seq {
		if firstSeen && cmp(e, last) < 0 {
			return false
		}
		last = e
		firstSeen = true
	}
	return true
}

// Reduce performs a [left-associative] [fold] of seq using
// b as the initial value and
// f as the left-associative binary operation.
// It terminates if and only if seq is finite.
//
// [fold]: https://en.wikipedia.org/wiki/Fold_(higher-order_function)
// [left-associative]: https://en.wikipedia.org/wiki/Associative_property#Notation_for_non-associative_operations
func Reduce[A, B any](seq iter.Seq[A], b B, f func(B, A) B) B {
	for a := range seq {
		b = f(b, a)
	}
	return b
}

// Len2 returns the number of elements in seq.
// It terminates if and only if seq is finite.
func Len2[K, V any](seq iter.Seq2[K, V]) int {
	var n int
	for range seq {
		n++
	}
	return n
}
