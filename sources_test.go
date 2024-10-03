package iterutil_test

import (
	"cmp"
	"fmt"
	"slices"
	"strings"
	"testing"

	"github.com/jub0bs/iterutil"
)

func ExampleEmpty() {
	for i := range iterutil.Empty[int]() {
		fmt.Println(i)
	}
	// Output:
}

func ExampleSeqOf() {
	for i := range iterutil.SeqOf(1, 2, 3) {
		fmt.Println(i)
	}
	// Output:
	// 1
	// 2
	// 3
}

func TestSeqOf(t *testing.T) {
	cases := []struct {
		desc      string
		elems     []string
		breakWhen func(string) bool
		want      []string
	}{
		{
			desc:      "no break",
			elems:     []string{"one", "two", "three"},
			breakWhen: alwaysFalse[string],
			want:      []string{"one", "two", "three"},
		}, {
			desc:      "break early",
			elems:     []string{"one", "two", "three"},
			breakWhen: equal("three"),
			want:      []string{"one", "two"},
		},
	}
	for _, tc := range cases {
		f := func(t *testing.T) {
			got := iterutil.SeqOf(tc.elems...)
			assertEqual(t, got, tc.want, tc.breakWhen)
		}
		t.Run(tc.desc, f)
	}
}

func ExampleBetween() {
	for i := range iterutil.Between(2, 9, 3) {
		fmt.Println(i)
	}
	// Output:
	// 2
	// 5
	// 8
}

func TestBetween(t *testing.T) {
	cases := []struct {
		desc       string
		n, m, step int
		breakWhen  func(int) bool
		want       []int
		wantPanic  bool
	}{
		{
			desc:      "no break positive step n less than m",
			n:         1,
			m:         11,
			step:      3,
			breakWhen: alwaysFalse[int],
			want:      []int{1, 4, 7, 10},
		}, {
			desc:      "break positive step n less than m",
			n:         1,
			m:         11,
			step:      3,
			breakWhen: equal(7),
			want:      []int{1, 4},
		}, {
			desc:      "no break positive step n equal to m",
			n:         11,
			m:         11,
			step:      3,
			breakWhen: alwaysFalse[int],
		}, {
			desc:      "break positive step n equal to m",
			n:         11,
			m:         11,
			step:      3,
			breakWhen: equal(7),
		}, {
			desc:      "no break positive step n greater than m",
			n:         11,
			m:         1,
			step:      3,
			breakWhen: alwaysFalse[int],
		}, {
			desc:      "break positive step n greater than m",
			n:         11,
			m:         1,
			step:      3,
			breakWhen: equal(7),
		}, {
			desc:      "no break negative step n less than m",
			n:         1,
			m:         11,
			step:      -3,
			breakWhen: alwaysFalse[int],
		}, {
			desc:      "break negative step n less than m",
			n:         1,
			m:         11,
			step:      -3,
			breakWhen: equal(7),
		}, {
			desc:      "no break negative step n equal to m",
			n:         11,
			m:         11,
			step:      -3,
			breakWhen: alwaysFalse[int],
		}, {
			desc:      "break negative step n equal to m",
			n:         11,
			m:         11,
			step:      -3,
			breakWhen: equal(7),
		}, {
			desc:      "no break negative step n greater than m",
			n:         11,
			m:         1,
			step:      -3,
			breakWhen: alwaysFalse[int],
			want:      []int{11, 8, 5, 2},
		}, {
			desc:      "break negative step n greater than m",
			n:         11,
			m:         1,
			step:      -3,
			breakWhen: equal(8),
			want:      []int{11},
		}, {
			desc:      "zero step n less than m",
			n:         1,
			m:         11,
			step:      0,
			wantPanic: true,
		}, {
			desc:      "zero step n equal to m",
			n:         11,
			m:         1,
			step:      0,
			wantPanic: true,
		}, {
			desc:      "zero step n greater than m",
			n:         11,
			m:         11,
			step:      0,
			wantPanic: true,
		},
	}
	for _, tc := range cases {
		f := func(t *testing.T) {
			if tc.wantPanic {
				defer func() {
					if recover() == nil {
						t.Fatalf("got no panic; want panic")
					}
				}()
			}
			got := iterutil.Between(tc.n, tc.m, tc.step)
			assertEqual(t, got, tc.want, tc.breakWhen)
		}
		t.Run(tc.desc, f)
	}
}

func ExampleRepeat() {
	var count int
	for s := range iterutil.Repeat("foo", -1) {
		count++
		if count > 3 {
			break
		}
		fmt.Println(s)
	}
	// Output:
	// foo
	// foo
	// foo
}

func TestRepeat(t *testing.T) {
	cases := []struct {
		desc      string
		elem      string
		count     int
		breakWhen func(string) bool
		want      []string
	}{
		{
			desc:      "finite no break",
			elem:      "foo",
			count:     3,
			breakWhen: alwaysFalse[string],
			want:      []string{"foo", "foo", "foo"},
		}, {
			desc:      "finite break early",
			elem:      "foo",
			count:     3,
			breakWhen: falseAfterN[string](2),
			want:      []string{"foo", "foo"},
		}, {
			desc:      "infinite",
			elem:      "foo",
			count:     -1,
			breakWhen: falseAfterN[string](2),
			want:      []string{"foo", "foo"},
		},
	}
	for _, tc := range cases {
		f := func(t *testing.T) {
			got := iterutil.Repeat(tc.elem, tc.count)
			assertEqual(t, got, tc.want, tc.breakWhen)
		}
		t.Run(tc.desc, f)
	}
}

func ExampleIterate() {
	double := func(i int) int { return i + i }
	for i := range iterutil.Iterate(1, double) {
		if i > 20 {
			break
		}
		fmt.Println(i)
	}
	// Output:
	// 1
	// 2
	// 4
	// 8
	// 16
}

func ExampleCycle() {
	seq := slices.Values([]int{1, 2, 3})
	var count int
	for i := range iterutil.Cycle(seq) {
		count++
		if count > 5 {
			break
		}
		fmt.Println(i)
	}
	// Output:
	// 1
	// 2
	// 3
	// 1
	// 2
}

func ExampleSortedFromMap() {
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	for k, v := range iterutil.SortedFromMap(m) {
		fmt.Println(k, v)
	}
	// Output:
	// one 1
	// three 3
	// two 2
}

func TestSortedFromMap(t *testing.T) {
	cases := []struct {
		desc      string
		m         map[string]int
		breakWhen func(string, int) bool
		want      []Pair[string, int]
	}{
		{
			desc: "no break",
			m: map[string]int{
				"one":   1,
				"two":   2,
				"three": 3,
			},
			breakWhen: alwaysFalse2[string, int],
			want: []Pair[string, int]{
				{"one", 1},
				{"three", 3},
				{"two", 2},
			},
		}, {
			desc: "break early",
			m: map[string]int{
				"one":   1,
				"two":   2,
				"three": 3,
			},
			breakWhen: equal2("three", 3),
			want: []Pair[string, int]{
				{"one", 1},
			},
		},
	}
	for _, tc := range cases {
		f := func(t *testing.T) {
			got := iterutil.SortedFromMap(tc.m)
			assertEqual2(t, got, tc.want, tc.breakWhen)
		}
		t.Run(tc.desc, f)
	}
}

func ExampleSortedFuncFromMap() {
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	for k, v := range iterutil.SortedFuncFromMap(m, strings.Compare) {
		fmt.Println(k, v)
	}
	// Output:
	// one 1
	// three 3
	// two 2
}

func ExampleSortedFuncFromMap_incorrect() {
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	lenCmp := func(s1, s2 string) int { return cmp.Compare(len(s1), len(s2)) }
	// Note that lenCmp does not correspond to a total order on strings.
	// More specifically, lenCmp is not antisymmetric:
	// for example, lenCmp("one", "two") = 0 and lenCmp("two", "one") = 0,
	// but "one" != "two".
	for k, v := range iterutil.SortedFuncFromMap(m, lenCmp) {
		fmt.Println(k, v)
	}
	// Consequently, the output is undeterministic; it may be either
	//
	// one 1
	// two 2
	// three 3
	//
	// or
	//
	// two 2
	// one 1
	// three 3
}

func TestSortedFuncFromMap(t *testing.T) {
	cases := []struct {
		desc      string
		m         map[string]int
		breakWhen func(string, int) bool
		want      []Pair[string, int]
	}{
		{
			desc: "no break",
			m: map[string]int{
				"one":   1,
				"two":   2,
				"three": 3,
			},
			breakWhen: alwaysFalse2[string, int],
			want: []Pair[string, int]{
				{"one", 1},
				{"three", 3},
				{"two", 2},
			},
		}, {
			desc: "break early",
			m: map[string]int{
				"one":   1,
				"two":   2,
				"three": 3,
			},
			breakWhen: equal2("three", 3),
			want: []Pair[string, int]{
				{"one", 1},
				{"two", 2},
			},
		},
	}
	for _, tc := range cases {
		f := func(t *testing.T) {
			got := iterutil.SortedFuncFromMap(tc.m, strings.Compare)
			assertEqual2(t, got, tc.want, tc.breakWhen)
		}
		t.Run(tc.desc, f)
	}
}
