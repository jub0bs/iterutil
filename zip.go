package iterutil

import "iter"

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
