package iterutil

import (
	"cmp"
	"iter"

	"golang.org/x/exp/constraints"
)

// Empty returns an empty iterator.
func Empty[E any]() iter.Seq[E] {
	return func(_ func(E) bool) {}
}

// IsEmpty reports whether seq is an empty iterator.
func IsEmpty[E any](seq iter.Seq[E]) bool {
	for range seq {
		return false
	}
	return true
}

// SeqOf returns an iterator composed of elems.
func SeqOf[E any](elems ...E) iter.Seq[E] {
	return func(yield func(E) bool) {
		for _, e := range elems {
			if !yield(e) {
				return
			}
		}
	}
}

// Between, if step is nonzero, returns an iterator
// ranging from n (inclusive) to m (exclusive) in increments of step;
// otherwise, it panics.
func Between[I constraints.Integer](n, m, step I) iter.Seq[I] {
	switch cmp.Compare(step, 0) {
	default:
		panic("step cannot be zero")
	case 1: // ascending
		return func(yield func(I) bool) {
			for ; n < m && yield(n); n += step {
				// deliberately empty
			}
		}
	case -1: // descending
		return func(yield func(I) bool) {
			for ; n > m && yield(n); n += step {
				// deliberately empty
			}
		}
	}
}

// Enumerate returns an iterator over pairs of indices (starting at 0)
// and elements of seq.
func Enumerate[I constraints.Integer, E any](seq iter.Seq[E]) iter.Seq2[I, E] {
	return func(yield func(I, E) bool) {
		var i I
		for v := range seq {
			if !yield(i, v) {
				return
			}
			i++
		}
	}
}

// Concat returns an iterator concatenating the passed in iterators.
func Concat[E any](seqs ...iter.Seq[E]) iter.Seq[E] {
	return func(yield func(E) bool) {
		for _, seq := range seqs {
			for e := range seq {
				if !yield(e) {
					return
				}
			}
		}
	}
}

// Flatten returns an iterator resulting from the concatenation of all iterators
// in seqs.
func Flatten[E any](seqs iter.Seq[iter.Seq[E]]) iter.Seq[E] {
	return func(yield func(E) bool) {
		for seq := range seqs {
			for e := range seq {
				if !yield(e) {
					return
				}
			}
		}
	}
}

// Map returns the result of applying f to each element of seq.
func Map[A, B any](seq iter.Seq[A], f func(A) B) iter.Seq[B] {
	return func(yield func(B) bool) {
		for a := range seq {
			if !yield(f(a)) {
				return
			}
		}
	}
}

// Filter returns an iterator composed of the elements of seq that
// satisfy predicate p.
func Filter[E any](seq iter.Seq[E], p func(E) bool) iter.Seq[E] {
	return func(yield func(E) bool) {
		for e := range seq {
			if p(e) && !yield(e) {
				return
			}
		}
	}
}

// TakeWhile returns the longest prefix of seq of elements that satisfy p.
func TakeWhile[E any](seq iter.Seq[E], p func(E) bool) iter.Seq[E] {
	return func(yield func(E) bool) {
		for e := range seq {
			if !p(e) || !yield(e) {
				return
			}
		}
	}
}

// DropWhile returns the suffix remaining after the longest prefix of seq
// of elements that satisfy p.
func DropWhile[E any](seq iter.Seq[E], p func(E) bool) iter.Seq[E] {
	return func(yield func(E) bool) {
		var doneDropping bool
		for e := range seq {
			if !doneDropping && p(e) {
				continue
			}
			doneDropping = true
			if !yield(e) {
				return
			}
		}
	}
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

// Take returns the prefix of seq
// whose length is min(max(count, 0), seq.Len()).
func Take[E any](seq iter.Seq[E], count int) iter.Seq[E] {
	return func(yield func(E) bool) {
		for e := range seq {
			count--
			if count < 0 || !yield(e) {
				return
			}
		}
	}
}

// Drop returns the suffix of seq
// after the first min(max(count, 0), seq.Len()) elements.
func Drop[E any](seq iter.Seq[E], count int) iter.Seq[E] {
	return func(yield func(E) bool) {
		for e := range seq {
			count--
			if 0 <= count {
				continue
			}
			if !yield(e) {
				return
			}
		}
	}
}

// At, if count is non-negative, returns
// the element at index n in seq and true
// or the zero value and false if seq contains fewer than count elements;
// otherwise, it panics.
func At[E any](seq iter.Seq[E], n int) (e E, ok bool) {
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
func Equal[E comparable](seq1, seq2 iter.Seq[E]) bool {
	return EqualFunc(seq1, seq2, equal)
}

func equal[E comparable](e1, e2 E) bool { return e1 == e2 }

// EqualFunc reports whether two iterators are equal using eq as equality
// function on each pair of elements.
// If the lengths are different, EqualFunc returns false.
// Otherwise, the elements are compared sequentially,
// and the comparison stops at the first pair for which eq returns false.
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
func Contains[E comparable](seq iter.Seq[E], target E) bool {
	for e := range seq {
		if e == target {
			return true
		}
	}
	return false
}

// ContainsFunc reports whether at least one element e of seq satisfies p(e).
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
func Compare[E cmp.Ordered](seq1, seq2 iter.Seq[E]) int {
	return CompareFunc(seq1, seq2, cmp.Compare)
}

// CompareFunc is like [Compare] but uses a custom comparison function on each
// pair of elements.
// The result is the first non-zero result of cmp;
// if cmp always returns 0, the result is 0 if len(seq1) == len(seq2),
// -1 if len(seq1) < len(seq2), and +1 if len(seq1) > len(seq2).
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

// Foldl performs a [left-associative] [fold] of seq using
// b as the initial value and
// f as the left-associative binary operation.
//
// [fold]: https://en.wikipedia.org/wiki/Fold_(higher-order_function)
// [left-associative]: https://en.wikipedia.org/wiki/Associative_property#Notation_for_non-associative_operations
func Foldl[A, B any](seq iter.Seq[A], b B, f func(B, A) B) B {
	for a := range seq {
		b = f(b, a)
	}
	return b
}

// ZipWith zips seq1 and seq2 with function f.
func ZipWith[A, B, C any](seq1 iter.Seq[A], seq2 iter.Seq[B], f func(A, B) C) iter.Seq[C] {
	return func(yield func(C) bool) {
		next1, stop1 := iter.Pull(seq1)
		defer stop1()
		next2, stop2 := iter.Pull(seq2)
		defer stop2()
		for {
			a, ok1 := next1()
			b, ok2 := next2()
			if !ok1 || !ok2 {
				return
			}
			if !yield(f(a, b)) {
				return
			}
		}
	}
}

// Repeat returns an iterator whose values are invariably e.
// the resulting iterator, if count is non-negative, is of length count;
// otherwise, it's infinite.
func Repeat[E any](e E, count int) iter.Seq[E] {
	if 0 <= count {
		return func(yield func(E) bool) {
			for range count {
				if !yield(e) {
					return
				}
			}
		}
	}
	return func(yield func(E) bool) {
		for yield(e) {
			// deliberately empty
		}
	}
}

// Iterate returns an infinite iterator composed of repeated applications
// of f to e.
func Iterate[E any](e E, f func(E) E) iter.Seq[E] {
	return func(yield func(E) bool) {
		for yield(e) {
			e = f(e)
		}
	}
}

// Cycle returns an iterator that infinitely repeats seq.
func Cycle[E any](seq iter.Seq[E]) iter.Seq[E] {
	return func(yield func(E) bool) {
		for {
			for e := range seq {
				if !yield(e) {
					return
				}
			}
		}
	}
}

// Push converts the “pull-style” iterator
// accessed by the two functions next and stop
// into a “push-style” iterator sequence.
// Push essentially is the inverse of [iter.Pull].
// Note that you must consume the resulting iterator;
// otherwise, the underlying pull-based iterator may leak.
func Push[E any](next func() (E, bool), stop func()) iter.Seq[E] {
	return func(yield func(E) bool) {
		defer stop()
		for {
			e, ok := next()
			if !ok || !yield(e) {
				return
			}
		}
	}
}
