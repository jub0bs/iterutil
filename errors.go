package iterutil

import "iter"

// AllErrors performs a [pre-order traversal] of err and returns an iterator
// over its index-error pairs. For more context, see the [errors] package.
//
// [pre-order traversal]: https://en.wikipedia.org/wiki/Tree_traversal#Arbitrary_trees
func AllErrors(err error) iter.Seq2[int, error] {
	return func(yield func(int, error) bool) {
		var i int
		if !yield(i, err) {
			return
		}
		i++
		switch err := err.(type) {
		case interface{ Unwrap() []error }:
			for _, err := range err.Unwrap() {
				for _, err := range AllErrors(err) {
					if !yield(i, err) {
						return
					}
					i++
				}
			}
		case interface{ Unwrap() error }:
			for _, err := range AllErrors(err.Unwrap()) {
				if !yield(i, err) {
					return
				}
				i++
			}
		default:
			return
		}
	}
}
