package iterutil

import (
	"cmp"
	"iter"
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

// Cons returns an iterator whose head is e and whose tail is seq.
func Cons[E any](e E, seq iter.Seq[E]) iter.Seq[E] {
	return func(yield func(E) bool) {
		if !yield(e) {
			return
		}
		for e := range seq {
			if !yield(e) {
				return
			}
		}
	}
}

// Head, if seq is non-empty, returns the head of seq and true;
// otherwise, it returns the zero value and false.
func Head[E any](seq iter.Seq[E]) (E, bool) {
	for e := range seq {
		return e, true
	}
	var zero E
	return zero, false
}

// Tail, if seq is non-empty, returns an iterator composed of
// all the elements of seq after the latter's head and true;
// otherwise, it returns nil and false.
func Tail[E any](seq iter.Seq[E]) (iter.Seq[E], bool) {
	next, stop := iter.Pull(seq)
	if _, ok := next(); !ok {
		return nil, false
	}
	f := func(yield func(E) bool) {
		defer stop()
		for {
			e, ok := next()
			if !ok {
				return
			}
			if !yield(e) {
				return
			}
		}
	}
	return f, true
}

// Uncons, if seq is non-empty, returns the head and tail of seq and true;
// otherwise, it returns the zero value, nil, and false.
func Uncons[E any](seq iter.Seq[E]) (E, iter.Seq[E], bool) {
	next, stop := iter.Pull(seq)
	head, ok := next()
	if !ok {
		return head, nil, false
	}
	tail := func(yield func(E) bool) {
		defer stop()
		for {
			e, ok := next()
			if !ok {
				return
			}
			if !yield(e) {
				return
			}
		}
	}
	return head, tail, true
}

// Append returns an iterator resulting from the concatenation of seq1 and
// seq2.
func Append[E any](seq1, seq2 iter.Seq[E]) iter.Seq[E] {
	return func(yield func(E) bool) {
		for e := range seq1 {
			if !yield(e) {
				return
			}
		}
		for e := range seq2 {
			if !yield(e) {
				return
			}
		}
	}
}

// Concat returns an iterator resulting from the concatenation of all iterators
// in seqs.
func Concat[E any](seqs iter.Seq[iter.Seq[E]]) iter.Seq[E] {
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
			if !p(e) {
				continue
			}
			if !yield(e) {
				return
			}
		}
	}
}

// FlatMap maps f over seq and concatenates the resulting iterators.
func FlatMap[A, B any](seq iter.Seq[A], f func(A) iter.Seq[B]) iter.Seq[B] {
	return func(yield func(B) bool) {
		for a := range seq {
			for b := range f(a) {
				if !yield(b) {
					return
				}
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

// Take returns the prefix of seq of length count.
// If count is negative, Take returns an empty iterator.
// If count is larger than the number of elements in seq, Take returns seq.
func Take[E any](seq iter.Seq[E], count int) iter.Seq[E] {
	if count < 0 {
		return Empty[E]()
	}
	return func(yield func(E) bool) {
		for e := range seq {
			count--
			if count < 0 || !yield(e) {
				return
			}
		}
	}
}

// Drop returns the suffix of seq after the first count elements.
// If count is negative, Drop returns seq.
// If count is larger than the number of elements in seq,
// Drop returns an empty iterator.
func Drop[E any](seq iter.Seq[E], count int) iter.Seq[E] {
	if count < 0 {
		return seq
	}
	return func(yield func(E) bool) {
		for e := range seq {
			count--
			if count >= 0 {
				continue
			}
			if !yield(e) {
				return
			}
		}
	}
}

// At, if seq has at least n elements,
// returns the element at index n in seq and true;
// otherwise, it returns the zero value and false.
func At[E any](seq iter.Seq[E], n int) (e E, ok bool) {
	if n < 0 {
		return
	}
	for v := range seq {
		switch cmp.Compare(n, 0) {
		case -1:
			return
		case 0:
			e = v
			ok = true
			return
		case 1:
			n--
			continue
		}
	}
	return
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
func ContainsFunc[E comparable](seq iter.Seq[E], p func(E) bool) bool {
	for e := range seq {
		if p(e) {
			return true
		}
	}
	return false
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

// Replicate returns an iterator of length count whose values are invariably e.
// If count is negative, Replicate returns an empty iterator.
func Replicate[E any](e E, count int) iter.Seq[E] {
	if count < 0 {
		return Empty[E]()
	}
	return func(yield func(E) bool) {
		for range count {
			if !yield(e) {
				return
			}
		}
	}
}

// Repeat returns an infinite iterator whose values are invariably e.
func Repeat[E any](e E) iter.Seq[E] {
	return func(yield func(E) bool) {
		for {
			if !yield(e) {
				return
			}
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