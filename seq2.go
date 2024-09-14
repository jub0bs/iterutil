package iterutil

import "iter"

// Len2 returns the number of elements in seq.
// It terminates if and only if seq is finite.
func Len2[K, V any](seq iter.Seq2[K, V]) int {
	var n int
	for range seq {
		n++
	}
	return n
}

// Filter returns an iterator composed of the pairs of seq that
// satisfy predicate p.
func Filter2[L, R any](seq iter.Seq2[L, R], p func(L, R) bool) iter.Seq2[L, R] {
	return func(yield func(L, R) bool) {
		for l, r := range seq {
			if !p(l, r) {
				continue
			}
			if !yield(l, r) {
				return
			}
		}
	}
}

// Swap returns an iterator over the value-key pairs of seq.
func Swap[A, B any](seq iter.Seq2[A, B]) iter.Seq2[B, A] {
	return func(yield func(B, A) bool) {
		for a, b := range seq {
			if !yield(b, a) {
				return
			}
		}
	}
}

// Push2 converts the “pull-style” iterator
// accessed by the two functions next and stop
// into a “push-style” iterator sequence.
// Push essentially is the inverse of [iter.Pull2].
func Push2[K, V any](next func() (K, V, bool), stop func()) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		defer stop()
		for {
			k, v, ok := next()
			if !ok || !yield(k, v) {
				return
			}
		}
	}
}
