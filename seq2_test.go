package iterutil_test

import (
	"fmt"
	"iter"
	"slices"
	"testing"

	"github.com/jub0bs/iterutil"
)

func ExampleLen2() {
	seq := slices.All([]int(nil))
	fmt.Println(iterutil.Len2(seq))
	seq = slices.All([]int{1, 2, 3, 4})
	fmt.Println(iterutil.Len2(seq))
	// Output:
	// 0
	// 4
}

func ExampleFilter2() {
	seq := slices.All([]string{"zero", "one", "two", "three", "four"})
	isShort := func(_ int, s string) bool { return len(s) < 5 }
	for i, s := range iterutil.Filter2(seq, isShort) {
		fmt.Println(i, s)
	}
	// Output:
	// 0 zero
	// 1 one
	// 2 two
	// 4 four
}

func TestFilter2(t *testing.T) {
	cases := []struct {
		desc      string
		elems     []string
		p         func(int, string) bool
		breakWhen func(int, string) bool
		want      []Pair[int, string]
	}{
		{
			desc:      "no break",
			elems:     []string{"zero", "one", "two", "three", "four"},
			p:         func(_ int, s string) bool { return len(s) < 5 },
			breakWhen: alwaysFalse2[int, string],
			want: []Pair[int, string]{
				{0, "zero"},
				{1, "one"},
				{2, "two"},
				{4, "four"},
			},
		}, {
			desc:      "no break",
			elems:     []string{"zero", "one", "two", "three", "four"},
			p:         func(_ int, s string) bool { return len(s) < 5 },
			breakWhen: equal2(4, "four"),
			want: []Pair[int, string]{
				{0, "zero"},
				{1, "one"},
				{2, "two"},
			},
		},
	}
	for _, tc := range cases {
		f := func(t *testing.T) {
			seq := slices.All(tc.elems)
			got := iterutil.Filter2(seq, tc.p)
			assertEqual2(t, got, tc.want, tc.breakWhen)
		}
		t.Run(tc.desc, f)
	}
}

func ExampleSwap() {
	seq := slices.All([]string{"foo", "bar", "baz"})
	for s, i := range iterutil.Swap(seq) {
		fmt.Println(s, i)
	}
	// Output:
	// foo 0
	// bar 1
	// baz 2
}

func TestSwap(t *testing.T) {
	cases := []struct {
		desc      string
		elems     []string
		breakWhen func(string, int) bool
		want      []Pair[string, int]
	}{
		{
			desc:      "no break",
			elems:     []string{"foo", "bar", "baz"},
			breakWhen: alwaysFalse2[string, int],
			want: []Pair[string, int]{
				{"foo", 0},
				{"bar", 1},
				{"baz", 2},
			},
		}, {
			desc:      "break early",
			elems:     []string{"foo", "bar", "baz"},
			breakWhen: equal2("baz", 2),
			want: []Pair[string, int]{
				{"foo", 0},
				{"bar", 1},
			},
		},
	}
	for _, tc := range cases {
		f := func(t *testing.T) {
			seq := slices.All(tc.elems)
			got := iterutil.Swap(seq)
			assertEqual2(t, got, tc.want, tc.breakWhen)
		}
		t.Run(tc.desc, f)
	}
}

func ExamplePush2() {
	seq := slices.All([]string{"foo", "bar", "baz"})
	next, stop := iter.Pull2(seq)
	for i, s := range iterutil.Push2(next, stop) {
		fmt.Println(i, s)
	}
	// Output:
	// 0 foo
	// 1 bar
	// 2 baz
}
