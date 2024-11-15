package iterutil

import (
	"cmp"
	"iter"
	"slices"

	"golang.org/x/exp/constraints"
)

// Empty returns an empty iterator.
func Empty[E any]() iter.Seq[E] {
	return func(_ func(E) bool) {}
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
func Between[I constraints.Signed](n, m, step I) iter.Seq[I] {
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

// Repeat returns an iterator whose values are invariably e.
// The resulting iterator, if count is non-negative, is of length count;
// otherwise, it's infinite.
func Repeat[I constraints.Integer, E any](e E, count I) iter.Seq[E] {
	if 0 <= count {
		return func(yield func(E) bool) {
			for i := I(0); i < count; i++ {
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

// SortedFromMap returns an iterator over the key-value pairs in m
// ordered by its keys.
func SortedFromMap[M ~map[K]V, K cmp.Ordered, V any](m M) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		// One possibility would be to simply iterate over
		// slices.Sorted(maps.Keys(m)),
		// but doing so would incur unnecessary allocations;
		// we can do better since we already know the number of keys.
		// See https://github.com/golang/go/issues/61899#issuecomment-2198727055.
		ks := keys(m)
		slices.Sort(ks)
		for _, k := range ks {
			if !yield(k, m[k]) {
				return
			}
		}
	}
}

// SortedFuncFromMap returns an iterator over the key-value pairs in m
// ordered by its keys, using cmp as comparison function.
//
// Note that, for a deterministic behavior,
// cmp must define a [total order] on K;
// for more details, see the testable example labeled "incorrect".
//
// [total order]: https://en.wikipedia.org/wiki/Total_order
func SortedFuncFromMap[M ~map[K]V, K comparable, V any](m M, cmp func(K, K) int) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		// see implementation comment in SortedFuncFromMap
		ks := keys(m)
		slices.SortFunc(ks, cmp)
		for _, k := range ks {
			if !yield(k, m[k]) {
				return
			}
		}
	}
}

func keys[K comparable, V any](m map[K]V) []K {
	ks := make([]K, 0, len(m))
	for k := range m {
		ks = append(ks, k)
	}
	return ks
}
