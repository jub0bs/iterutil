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
func Filter2[K, V any](seq iter.Seq2[K, V], p func(K, V) bool) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range seq {
			if p(k, v) && !yield(k, v) {
				return
			}
		}
	}
}

// Swap returns an iterator over the value-key pairs of seq.
func Swap[K, V any](seq iter.Seq2[K, V]) iter.Seq2[V, K] {
	return func(yield func(V, K) bool) {
		for k, v := range seq {
			if !yield(v, k) {
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
