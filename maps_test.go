package iterutil_test

import (
	"cmp"
	"fmt"
	"strings"
	"testing"

	"github.com/jub0bs/iterutil"
)

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
