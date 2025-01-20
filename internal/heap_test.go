package internal_test

import (
	"cmp"
	"slices"
	"testing"

	"github.com/jub0bs/iterutil/internal"
)

func TestHeap(t *testing.T) {
	cases := []struct {
		desc      string
		s         []int
		breakWhen func(int) bool
	}{
		{
			desc:      "no break",
			s:         []int{1, 2, 3, 4, 5},
			breakWhen: alwaysFalse[int],
		}, {
			desc:      "break early",
			s:         []int{1, 2, 3, 4, 5},
			breakWhen: equal(3),
		},
	}
	for _, tc := range cases {
		f := func(t *testing.T) {
			seq := internal.NewHeap(tc.s, cmp.Compare).Iterator
			var got []int
			for v := range seq {
				if tc.breakWhen(v) {
					break
				}
				got = append(got, v)
			}
			if !slices.IsSortedFunc(got, cmp.Compare) {
				t.Errorf("got unsorted iterator: %v; want sorted iterator", got)
			}
		}
		t.Run(tc.desc, f)
	}
}

func alwaysFalse[E any](_ E) bool {
	return false
}

func equal[E comparable](target E) func(E) bool {
	return func(e E) bool {
		return e == target
	}
}
