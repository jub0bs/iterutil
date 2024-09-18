package iterutil_test

import (
	"fmt"
	"iter"
	"slices"
	"testing"

	"github.com/jub0bs/iterutil"
)

func ExampleEmpty() {
	for i := range iterutil.Empty[int]() {
		fmt.Println(i)
	}
	// Output:
}

func ExampleIsEmpty() {
	seq := slices.Values([]int{})
	fmt.Println(iterutil.IsEmpty(seq))
	seq = slices.Values([]int{1, 2, 3, 4})
	fmt.Println(iterutil.IsEmpty(seq))
	// Output:
	// true
	// false
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

func ExampleCons() {
	seq := iterutil.Cons(0, slices.Values([]int{1, 2, 3}))
	for i := range seq {
		fmt.Println(i)
	}
	// Output:
	// 0
	// 1
	// 2
	// 3
}

func TestCons(t *testing.T) {
	cases := []struct {
		desc      string
		first     string
		rest      []string
		breakWhen func(string) bool
		want      []string
	}{
		{
			desc:      "no break",
			first:     "zero",
			rest:      []string{"one", "two", "three"},
			breakWhen: alwaysFalse[string],
			want:      []string{"zero", "one", "two", "three"},
		}, {
			desc:      "break early",
			first:     "zero",
			breakWhen: equal("zero"),
			rest:      []string{"one", "two", "three"},
		}, {
			desc:      "break early but later",
			first:     "zero",
			rest:      []string{"one", "two", "three"},
			breakWhen: equal("two"),
			want:      []string{"zero", "one"},
		},
	}
	for _, tc := range cases {
		f := func(t *testing.T) {
			rest := slices.Values(tc.rest)
			got := iterutil.Cons(tc.first, rest)
			assertEqual(t, got, tc.want, tc.breakWhen)
		}
		t.Run(tc.desc, f)
	}
}

func ExampleFlatten() {
	seq1 := slices.Values([]string{"foo", "bar"})
	seq2 := slices.Values([]string{"baz", "qux"})
	seqs := slices.Values([]iter.Seq[string]{seq1, seq2})
	for s := range iterutil.Flatten(seqs) {
		fmt.Println(s)
	}
	// Output:
	// foo
	// bar
	// baz
	// qux
}

func TestFlatten(t *testing.T) {
	cases := []struct {
		desc      string
		seq1      []string
		seq2      []string
		breakWhen func(string) bool
		want      []string
	}{
		{
			desc:      "empty",
			seq1:      []string{},
			seq2:      []string{},
			breakWhen: alwaysFalse[string],
		}, {
			desc:      "no break",
			seq1:      []string{"one", "two", "three"},
			seq2:      []string{"four", "five", "six"},
			breakWhen: alwaysFalse[string],
			want:      []string{"one", "two", "three", "four", "five", "six"},
		}, {
			desc:      "break early",
			seq1:      []string{"one", "two", "three"},
			seq2:      []string{"four", "five", "six"},
			breakWhen: equal("three"),
			want:      []string{"one", "two"},
		},
	}
	for _, tc := range cases {
		f := func(t *testing.T) {
			seq1 := slices.Values(tc.seq1)
			seq2 := slices.Values(tc.seq2)
			seq := slices.Values([]iter.Seq[string]{seq1, seq2})
			got := iterutil.Flatten(seq)
			assertEqual(t, got, tc.want, tc.breakWhen)
		}
		t.Run(tc.desc, f)
	}
}

func ExampleMap() {
	seq := slices.Values([]string{"one", "two", "three"})
	length := func(s string) int { return len(s) }
	for s := range iterutil.Map(seq, length) {
		fmt.Println(s)
	}
	// Output:
	// 3
	// 3
	// 5
}

func TestMap(t *testing.T) {
	cases := []struct {
		desc      string
		elems     []string
		f         func(string) int
		breakWhen func(int) bool
		want      []int
	}{
		{
			desc:      "no break",
			elems:     []string{"one", "two", "three"},
			f:         func(s string) int { return len(s) },
			breakWhen: alwaysFalse[int],
			want:      []int{3, 3, 5},
		}, {
			desc:      "break early",
			elems:     []string{"one", "two", "three"},
			f:         func(s string) int { return len(s) },
			breakWhen: equal(5),
			want:      []int{3, 3},
		},
	}
	for _, tc := range cases {
		f := func(t *testing.T) {
			seq := slices.Values(tc.elems)
			got := iterutil.Map(seq, tc.f)
			assertEqual(t, got, tc.want, tc.breakWhen)
		}
		t.Run(tc.desc, f)
	}
}

func ExampleFilter() {
	seq := slices.Values([]int{1, 42, 99, 100})
	isOdd := func(i int) bool { return i%2 != 0 }
	for s := range iterutil.Filter(seq, isOdd) {
		fmt.Println(s)
	}
	// Output:
	// 1
	// 99
}

func TestFilter(t *testing.T) {
	cases := []struct {
		desc      string
		elems     []string
		p         func(string) bool
		breakWhen func(string) bool
		want      []string
	}{
		{
			desc:      "no break",
			elems:     []string{"one", "two", "three"},
			p:         func(s string) bool { return len(s) == 3 },
			breakWhen: alwaysFalse[string],
			want:      []string{"one", "two"},
		}, {
			desc:      "break early",
			elems:     []string{"one", "two", "three"},
			p:         func(s string) bool { return len(s) == 3 },
			breakWhen: equal("two"),
			want:      []string{"one"},
		},
	}
	for _, tc := range cases {
		f := func(t *testing.T) {
			seq := slices.Values(tc.elems)
			got := iterutil.Filter(seq, tc.p)
			assertEqual(t, got, tc.want, tc.breakWhen)
		}
		t.Run(tc.desc, f)
	}
}

func ExampleFlatMap() {
	seq := slices.Values([]int{0, 1, 2, 3})
	repeatN := func(i int) iter.Seq[int] {
		return slices.Values(slices.Repeat([]int{i}, i))
	}
	for i := range iterutil.FlatMap(seq, repeatN) {
		fmt.Println(i)
	}
	// Output:
	// 1
	// 2
	// 2
	// 3
	// 3
	// 3
}

func TestFlatMap(t *testing.T) {
	cases := []struct {
		desc      string
		elems     []string
		f         func(string) iter.Seq[byte]
		breakWhen func(byte) bool
		want      []byte
	}{
		{
			desc:      "no break",
			elems:     []string{"one", "two", "three"},
			f:         func(s string) iter.Seq[byte] { return slices.Values([]byte(s)) },
			breakWhen: alwaysFalse[byte],
			want:      []byte("one" + "two" + "three"),
		}, {
			desc:      "break early",
			elems:     []string{"one", "two", "three"},
			f:         func(s string) iter.Seq[byte] { return slices.Values([]byte(s)) },
			breakWhen: equal(byte('w')),
			want:      []byte("one" + "t"),
		},
	}
	for _, tc := range cases {
		f := func(t *testing.T) {
			seq := slices.Values(tc.elems)
			got := iterutil.FlatMap(seq, tc.f)
			assertEqual(t, got, tc.want, tc.breakWhen)
		}
		t.Run(tc.desc, f)
	}
}

func ExampleTakeWhile() {
	seq := slices.Values([]string{"foo", "bar", "baz", "qux"})
	isNotBaz := func(s string) bool { return s != "baz" }
	for s := range iterutil.TakeWhile(seq, isNotBaz) {
		fmt.Println(s)
	}
	// Output:
	// foo
	// bar
}

func ExampleDropWhile() {
	seq := slices.Values([]string{"foo", "bar", "baz", "qux"})
	isNotBaz := func(s string) bool { return s != "baz" }
	for s := range iterutil.DropWhile(seq, isNotBaz) {
		fmt.Println(s)
	}
	// Output:
	// baz
	// qux
}

func TestDropWhile(t *testing.T) {
	cases := []struct {
		desc      string
		elems     []string
		p         func(string) bool
		breakWhen func(string) bool
		want      []string
	}{
		{
			desc:      "no break",
			elems:     []string{"one", "two", "three"},
			p:         func(s string) bool { return len(s) == 3 },
			breakWhen: alwaysFalse[string],
			want:      []string{"three"},
		}, {
			desc:      "break early",
			elems:     []string{"one", "two", "three", "four"},
			p:         func(s string) bool { return len(s) == 3 },
			breakWhen: equal("four"),
			want:      []string{"three"},
		},
	}
	for _, tc := range cases {
		f := func(t *testing.T) {
			seq := slices.Values(tc.elems)
			got := iterutil.DropWhile(seq, tc.p)
			assertEqual(t, got, tc.want, tc.breakWhen)
		}
		t.Run(tc.desc, f)
	}
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

func ExampleTake() {
	seq := slices.Values([]string{"foo", "bar", "baz", "qux"})
	for s := range iterutil.Take(seq, 2) {
		fmt.Println(s)
	}
	// Output:
	// foo
	// bar
}

func TestTake(t *testing.T) {
	cases := []struct {
		desc      string
		elems     []string
		count     int
		breakWhen func(string) bool
		want      []string
		panics    bool
	}{
		{
			desc:      "negative count",
			elems:     []string{"one", "two", "three"},
			count:     -1,
			breakWhen: alwaysFalse[string],
			panics:    true,
		}, {
			desc:      "no break",
			elems:     []string{"one", "two", "three"},
			count:     2,
			breakWhen: alwaysFalse[string],
			want:      []string{"one", "two"},
		}, {
			desc:      "break early",
			elems:     []string{"one", "two", "three"},
			count:     3,
			breakWhen: equal("two"),
			want:      []string{"one"},
		},
	}
	for _, tc := range cases {
		f := func(t *testing.T) {
			defer func() {
				if r := recover(); tc.panics && r == nil {
					t.Errorf("got no panic; want panic")
				}
			}()
			seq := slices.Values(tc.elems)
			got := iterutil.Take(seq, tc.count)
			assertEqual(t, got, tc.want, tc.breakWhen)
		}
		t.Run(tc.desc, f)
	}
}

func ExampleDrop() {
	seq := slices.Values([]string{"foo", "bar", "baz", "qux"})
	for s := range iterutil.Drop(seq, 3) {
		fmt.Println(s)
	}
	// Output:
	// qux
}

func TestDrop(t *testing.T) {
	cases := []struct {
		desc      string
		elems     []string
		count     int
		breakWhen func(string) bool
		want      []string
		panics    bool
	}{
		{
			desc:      "negative count",
			elems:     []string{"one", "two", "three"},
			count:     -1,
			breakWhen: alwaysFalse[string],
			panics:    true,
		}, {
			desc:      "no break",
			elems:     []string{"one", "two", "three"},
			count:     2,
			breakWhen: alwaysFalse[string],
			want:      []string{"three"},
		}, {
			desc:      "break early",
			elems:     []string{"one", "two", "three", "four"},
			count:     2,
			breakWhen: equal("four"),
			want:      []string{"three"},
		},
	}
	for _, tc := range cases {
		f := func(t *testing.T) {
			defer func() {
				if r := recover(); tc.panics && r == nil {
					t.Errorf("got no panic; want panic")
				}
			}()
			seq := slices.Values(tc.elems)
			got := iterutil.Drop(seq, tc.count)
			assertEqual(t, got, tc.want, tc.breakWhen)
		}
		t.Run(tc.desc, f)
	}
}

func ExampleAt() {
	seq := slices.Values([]string{"foo", "bar", "baz", "qux"})
	fmt.Println(iterutil.At(seq, 2))
	// Output:
	// baz true
}

func TestAt(t *testing.T) {
	cases := []struct {
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
	for _, tc := range cases {
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

func ExampleFoldl() {
	seq := slices.Values([]int{1, 2, 3, 4, 5, 6})
	plus := func(i, j int) int { return i + j }
	sum := iterutil.Foldl(seq, 0, plus)
	fmt.Println(sum)
	// Output: 21
}

func ExampleZipWith() {
	french := slices.Values([]string{"un", "deux", "trois", "quatre", "cinq"})
	english := slices.Values([]string{"one", "two", "three", "four"})
	join := func(fr, en string) string { return fr + " => " + en }
	seq := iterutil.ZipWith(french, english, join)
	for s := range seq {
		fmt.Println(s)
	}
	// Output:
	// un => one
	// deux => two
	// trois => three
	// quatre => four
}

func TestZipWith(t *testing.T) {
	cases := []struct {
		desc      string
		keys      []string
		values    []string
		f         func(string, string) string
		breakWhen func(string) bool
		want      []string
	}{
		{
			desc:      "no break",
			keys:      []string{"un", "deux", "trois", "quatre", "cinq"},
			values:    []string{"one", "two", "three"},
			f:         func(fr, en string) string { return fr + " => " + en },
			breakWhen: alwaysFalse[string],
			want: []string{
				"un => one",
				"deux => two",
				"trois => three",
			},
		}, {
			desc:      "break early",
			keys:      []string{"un", "deux", "trois", "quatre", "cinq"},
			values:    []string{"one", "two", "three"},
			f:         func(fr, en string) string { return fr + " => " + en },
			breakWhen: equal("trois => three"),
			want: []string{
				"un => one",
				"deux => two",
			},
		},
	}
	for _, tc := range cases {
		f := func(t *testing.T) {
			keys := slices.Values(tc.keys)
			values := slices.Values(tc.values)
			got := iterutil.ZipWith(keys, values, tc.f)
			assertEqual(t, got, tc.want, tc.breakWhen)
		}
		t.Run(tc.desc, f)
	}
}

func ExampleRepeat() {
	seq := iterutil.Repeat("foo", -1)
	var count int
	for s := range seq {
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
	seq := iterutil.Iterate(1, double)
	for i := range seq {
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
	cycle := iterutil.Cycle(seq)
	var count int
	for i := range cycle {
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

func ExamplePush() {
	seq := slices.Values([]int{1, 2, 3})
	next, stop := iter.Pull(seq)
	for i := range iterutil.Push(next, stop) {
		fmt.Println(i)
	}
	// Output:
	// 1
	// 2
	// 3
}
