package iterutil

import (
	"iter"
)

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

// Push2 converts the “pull-style” iterator
// accessed by the two functions next and stop
// into a “push-style” iterator sequence.
// Push2 essentially is the inverse of [iter.Pull2].
// Note that you must consume the resulting iterator;
// otherwise, the underlying pull-based iterator may leak.
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
