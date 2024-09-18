package iterutil

import "iter"

type (
	bwrapper interface{ Unwrap() []error }
	dwrapper interface{ Unwrap() error }
)

// AllErrors performs a preorder traversal of err and returns an iterator
// over its index-error pairs. For more context, see the [errors] package.
func AllErrors(err error) iter.Seq2[int, error] {
	return func(yield func(int, error) bool) {
		var i int
		if !yield(i, err) {
			return
		}
		i++
		switch err := err.(type) {
		case bwrapper:
			for _, err := range err.Unwrap() {
				for _, err := range AllErrors(err) {
					if !yield(i, err) {
						return
					}
					i++
				}
			}
		case dwrapper:
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
