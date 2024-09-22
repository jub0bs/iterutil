package iterutil

import (
	"iter"

	"golang.org/x/exp/constraints"
)

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

// Take returns the prefix of seq
// whose length is min(max(count, 0), Len(seq)).
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
// after the first min(max(count, 0), Len(seq)) elements.
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

// Zip zips seq1 and seq2 into a sequence of corresponding pairs.
func Zip[K, V any](seq1 iter.Seq[K], seq2 iter.Seq[V]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		next1, stop1 := iter.Pull(seq1)
		defer stop1()
		next2, stop2 := iter.Pull(seq2)
		defer stop2()
		for {
			k, ok1 := next1()
			v, ok2 := next2()
			if !ok1 || !ok2 {
				return
			}
			if !yield(k, v) {
				return
			}
		}
	}
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

// Left return an iterator composed of the keys of the pairs in seq.
func Left[K, V any](seq iter.Seq2[K, V]) iter.Seq[K] {
	return func(yield func(K) bool) {
		for k := range seq {
			if !yield(k) {
				return
			}
		}
	}
}

// Right return an iterator composed of the values of the pairs in seq.
func Right[K, V any](seq iter.Seq2[K, V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, v := range seq {
			if !yield(v) {
				return
			}
		}
	}
}
