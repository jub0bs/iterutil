package iterutil

import (
	"cmp"
	"iter"
	"maps"
	"slices"
)

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
		for _, k := range slices.Sorted(maps.Keys(m)) {
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
