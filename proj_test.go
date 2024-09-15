package iterutil_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/jub0bs/iterutil"
)

func ExampleLeft() {
	seq := slices.All([]string{"zero", "one", "two", "three", "four"})
	for i := range iterutil.Left(seq) {
		fmt.Println(i)
	}
	// Output:
	// 0
	// 1
	// 2
	// 3
	// 4
}

func TestLeft(t *testing.T) {
	cases := []struct {
		desc      string
		elems     []string
		breakWhen func(int) bool
		want      []int
	}{
		{
			desc:      "no break",
			elems:     []string{"foo", "bar", "baz"},
			breakWhen: alwaysFalse[int],
			want:      []int{0, 1, 2},
		}, {
			desc:      "break early",
			elems:     []string{"foo", "bar", "baz"},
			breakWhen: equal(1),
			want:      []int{0},
		},
	}
	for _, tc := range cases {
		f := func(t *testing.T) {
			seq := slices.All(tc.elems)
			got := iterutil.Left(seq)
			assertEqual(t, got, tc.want, tc.breakWhen)
		}
		t.Run(tc.desc, f)
	}
}

func ExampleRight() {
	seq := slices.All([]string{"zero", "one", "two", "three", "four"})
	for s := range iterutil.Right(seq) {
		fmt.Println(s)
	}
	// Output:
	// zero
	// one
	// two
	// three
	// four
}

func TestRight(t *testing.T) {
	cases := []struct {
		desc      string
		elems     []string
		breakWhen func(string) bool
		want      []string
	}{
		{
			desc:      "no break",
			elems:     []string{"foo", "bar", "baz"},
			breakWhen: alwaysFalse[string],
			want:      []string{"foo", "bar", "baz"},
		}, {
			desc:      "break early",
			elems:     []string{"foo", "bar", "baz"},
			breakWhen: falseAfterN[string](1),
			want:      []string{"foo"},
		},
	}
	for _, tc := range cases {
		f := func(t *testing.T) {
			seq := slices.All(tc.elems)
			got := iterutil.Right(seq)
			assertEqual(t, got, tc.want, tc.breakWhen)
		}
		t.Run(tc.desc, f)
	}
}
