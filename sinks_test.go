package iterutil_test

import (
	"cmp"
	"fmt"
	"slices"
	"strings"
	"testing"

	"github.com/jub0bs/iterutil"
)

func ExampleIsEmpty() {
	seq := slices.Values([]int{})
	fmt.Println(iterutil.IsEmpty(seq))
	seq = slices.Values([]int{1, 2, 3, 4})
	fmt.Println(iterutil.IsEmpty(seq))
	// Output:
	// true
	// false
}

func ExampleLen() {
	seq := slices.Values([]int{})
	fmt.Println(iterutil.Len(seq))
	seq = slices.Values([]int{1, 2, 3, 4})
	fmt.Println(iterutil.Len(seq))
	// Output:
	// 0
	// 4
}

func ExampleAt() {
	seq := slices.Values([]string{"foo", "bar", "baz", "qux"})
	fmt.Println(iterutil.At(seq, 2))
	// Output:
	// baz true
}

func TestAt(t *testing.T) {
	intCases := []struct {
		desc   string
		elems  []string
		n      int
		want   string
		ok     bool
		panics bool
	}{
		{
			desc:   "negative index",
			elems:  []string{"one", "two", "three"},
			n:      -1,
			panics: true,
		}, {
			desc:  "within bounds",
			elems: []string{"one", "two", "three"},
			n:     2,
			want:  "three",
			ok:    true,
		}, {
			desc:  "out of bounds",
			elems: []string{"one", "two", "three"},
			n:     4,
		},
	}
	for _, tc := range intCases {
		f := func(t *testing.T) {
			defer func() {
				if r := recover(); tc.panics && r == nil {
					t.Errorf("got no panic; want panic")
				}
			}()
			seq := slices.Values(tc.elems)
			s, ok := iterutil.At(seq, tc.n)
			if s != tc.want || ok != tc.ok {
				t.Fatalf("got %s, %t; want %s, %t", s, ok, tc.want, tc.ok)
			}

		}
		t.Run(tc.desc, f)
	}
	uintCases := []struct {
		desc  string
		elems []string
		n     uint
		want  string
		ok    bool
	}{
		{
			desc:  "within bounds",
			elems: []string{"one", "two", "three"},
			n:     2,
			want:  "three",
			ok:    true,
		}, {
			desc:  "out of bounds",
			elems: []string{"one", "two", "three"},
			n:     4,
		},
	}
	for _, tc := range uintCases {
		f := func(t *testing.T) {
			seq := slices.Values(tc.elems)
			s, ok := iterutil.At(seq, tc.n)
			if s != tc.want || ok != tc.ok {
				t.Fatalf("got %s, %t; want %s, %t", s, ok, tc.want, tc.ok)
			}

		}
		t.Run(tc.desc, f)
	}
}

func ExampleEqual() {
	seq1 := slices.Values([]string{"foo", "bar", "baz", "qux"})
	seq2 := slices.Values([]string{"foo", "bar", "baz", "qux"})
	fmt.Println(iterutil.Equal(seq1, seq2))
	// Output:
	// true
}

func TestEqual(t *testing.T) {
	cases := []struct {
		desc string
		seq1 []string
		seq2 []string
		want bool
	}{
		{
			desc: "equal",
			seq1: []string{"foo", "bar", "baz"},
			seq2: []string{"foo", "bar", "baz"},
			want: true,
		}, {
			desc: "not same size",
			seq1: []string{"foo", "bar", "baz", "qux"},
			seq2: []string{"foo", "bar", "baz"},
			want: false,
		}, {
			desc: "same size different values",
			seq1: []string{"foo", "bar", "baz", "qux"},
			seq2: []string{"foo", "bar", "baz", "quux"},
			want: false,
		},
	}
	for _, tc := range cases {
		f := func(t *testing.T) {
			seq1 := slices.Values(tc.seq1)
			seq2 := slices.Values(tc.seq2)
			got := iterutil.Equal(seq1, seq2)
			if got != tc.want {
				t.Errorf("got %t; want %t", got, tc.want)
			}
		}
		t.Run(tc.desc, f)
	}
}

func ExampleEqualFunc() {
	seq1 := slices.Values([]string{"foo", "bar", "baz", "qux"})
	seq2 := slices.Values([]string{"foO", "bAr", "Baz", "QUX"})
	fmt.Println(iterutil.EqualFunc(seq1, seq2, strings.EqualFold))
	// Output:
	// true
}

func ExampleContains() {
	seq := slices.Values([]int{1})
	fmt.Println(iterutil.Contains(seq, 2))
	seq = slices.Values([]int{1, 2, 3})
	fmt.Println(iterutil.Contains(seq, 2))
	// Output:
	// false
	// true
}

func ExampleContainsFunc() {
	isEven := func(i int) bool { return i%2 == 0 }
	seq := slices.Values([]int{1})
	fmt.Println(iterutil.ContainsFunc(seq, isEven))
	seq = slices.Values([]int{1, 2, 3})
	fmt.Println(iterutil.ContainsFunc(seq, isEven))
	// Output:
	// false
	// true
}

func ExampleMin() {
	seq := slices.Values([]int(nil))
	fmt.Println(iterutil.Min(seq))
	seq = slices.Values([]int{3, 5, 1, 42})
	fmt.Println(iterutil.Min(seq))
	// Output:
	// 0 false
	// 1 true
}

func ExampleMinFunc() {
	lenCmp := func(s1, s2 string) int { return cmp.Compare(len(s1), len(s2)) }
	seq := slices.Values([]string(nil))
	fmt.Println(iterutil.MinFunc(seq, lenCmp))
	seq = slices.Values([]string{"quux", "qux", "baz", "bar", "foo"})
	fmt.Println(iterutil.MinFunc(seq, lenCmp))
	// Output:
	//  false
	// qux true
}

func ExampleMax() {
	seq := slices.Values([]int(nil))
	fmt.Println(iterutil.Max(seq))
	seq = slices.Values([]int{3, 5, 1, 42})
	fmt.Println(iterutil.Max(seq))
	// Output:
	// 0 false
	// 42 true
}

func ExampleMaxFunc() {
	lenCmp := func(s1, s2 string) int { return cmp.Compare(len(s1), len(s2)) }
	seq := slices.Values([]string(nil))
	fmt.Println(iterutil.MaxFunc(seq, lenCmp))
	seq = slices.Values([]string{"qux", "quux", "corge", "grault", "garply"})
	fmt.Println(iterutil.MaxFunc(seq, lenCmp))
	// Output:
	//  false
	// grault true
}

func ExampleCompare() {
	seq1 := slices.Values([]string{"foo", "bar", "baz", "qux"})
	seq2 := slices.Values([]string{"foo", "bar", "baz", "qux", "quux"})
	fmt.Println(iterutil.Compare(seq1, seq2))
	// Output:
	// -1
}

func TestCompare(t *testing.T) {
	cases := []struct {
		desc string
		seq1 []string
		seq2 []string
		want int
	}{
		{
			desc: "equal",
			seq1: []string{"foo", "bar", "baz"},
			seq2: []string{"foo", "bar", "baz"},
			want: 0,
		}, {
			desc: "seq1 strict prefix of seq2",
			seq1: []string{"foo", "bar", "baz"},
			seq2: []string{"foo", "bar", "baz", "qux"},
			want: -1,
		}, {
			desc: "seq2 strict prefix of seq1",
			seq1: []string{"foo", "bar", "baz", "qux"},
			seq2: []string{"foo", "bar", "baz"},
			want: 1,
		}, {
			desc: "same values but last one",
			seq1: []string{"foo", "bar", "baz", "qux"},
			seq2: []string{"foo", "bar", "baz", "quux"},
			want: 1,
		},
	}
	for _, tc := range cases {
		f := func(t *testing.T) {
			seq1 := slices.Values(tc.seq1)
			seq2 := slices.Values(tc.seq2)
			got := iterutil.Compare(seq1, seq2)
			if got != tc.want {
				t.Errorf("got %d; want %d", got, tc.want)
			}
		}
		t.Run(tc.desc, f)
	}
}

func ExampleCompareFunc() {
	seq1 := slices.Values([]string{"foo", "bar", "baz", "qux", "quux"})
	seq2 := slices.Values([]string{"000", "111", "222", "333", "4444"})
	lenCmp := func(s1, s2 string) int { return cmp.Compare(len(s1), len(s2)) }
	fmt.Println(iterutil.CompareFunc(seq1, seq2, lenCmp))
	// Output:
	// 0
}

func ExampleIsSorted() {
	seq := slices.Values([]string{"bar", "baz", "foo", "quux", "qux"})
	fmt.Println(iterutil.IsSorted(seq))
	// Output:
	// true
}

func TestIsSorted(t *testing.T) {
	cases := []struct {
		desc  string
		elems []string
		want  bool
	}{
		{
			desc:  "not sorted",
			elems: []string{"foo", "bar", "baz", "qux", "quux"},
			want:  false,
		}, {
			desc:  "sorted",
			elems: []string{"bar", "baz", "foo", "quux", "qux"},
			want:  true,
		},
	}
	for _, tc := range cases {
		f := func(t *testing.T) {
			seq := slices.Values(tc.elems)
			got := iterutil.IsSorted(seq)
			if got != tc.want {
				t.Errorf("got %t; want %t", got, tc.want)
			}
		}
		t.Run(tc.desc, f)
	}
}

func ExampleIsSortedFunc() {
	seq := slices.Values([]string{"bar", "baz", "foo", "qux", "quux"})
	lenCmp := func(s1, s2 string) int { return cmp.Compare(len(s1), len(s2)) }
	fmt.Println(iterutil.IsSortedFunc(seq, lenCmp))
	// Output:
	// true
}

func TestIsSortedFunc(t *testing.T) {
	cases := []struct {
		desc  string
		elems []string
	}{
		{
			desc:  "not sorted",
			elems: []string{"bar", "baz", "foo", "quux", "qux"},
		}, {
			desc:  "sorted",
			elems: []string{"bar", "baz", "foo", "qux", "quux"},
		},
	}
	for _, tc := range cases {
		f := func(t *testing.T) {
			lenCmp := func(s1, s2 string) int { return cmp.Compare(len(s1), len(s2)) }
			seq := slices.Values(tc.elems)
			got := iterutil.IsSortedFunc(seq, lenCmp)
			want := slices.IsSortedFunc(tc.elems, lenCmp)
			if got != want {
				t.Errorf("got %t; want %t", got, want)
			}
		}
		t.Run(tc.desc, f)
	}
}

func ExampleReduce() {
	seq := slices.Values([]int{1, 2, 3, 4, 5, 6})
	plus := func(i, j int) int { return i + j }
	sum := iterutil.Reduce(seq, 0, plus)
	fmt.Println(sum)
	// Output: 21
}

func ExampleLen2() {
	seq := slices.All([]int(nil))
	fmt.Println(iterutil.Len2(seq))
	seq = slices.All([]int{1, 2, 3, 4})
	fmt.Println(iterutil.Len2(seq))
	// Output:
	// 0
	// 4
}
