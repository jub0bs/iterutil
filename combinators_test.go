package iterutil_test

import (
	"fmt"
	"iter"
	"slices"
	"testing"

	"github.com/jub0bs/iterutil"
)

func ExampleEnumerate() {
	seq := slices.Values([]string{"foo", "bar", "baz"})
	for i, v := range iterutil.Enumerate[int](seq) {
		fmt.Println(i, v)
	}
	// Output:
	// 0 foo
	// 1 bar
	// 2 baz
}

func TestEnumerate(t *testing.T) {
	cases := []struct {
		desc      string
		elems     []string
		breakWhen func(int, string) bool
		want      []Pair[int, string]
	}{
		{
			desc:      "no break",
			elems:     []string{"zero", "one", "two", "three"},
			breakWhen: alwaysFalse2[int, string],
			want: []Pair[int, string]{
				{0, "zero"},
				{1, "one"},
				{2, "two"},
				{3, "three"},
			},
		}, {
			desc:      "break early",
			elems:     []string{"zero", "one", "two", "three"},
			breakWhen: equal2(2, "two"),
			want: []Pair[int, string]{
				{0, "zero"},
				{1, "one"},
			},
		},
	}
	for _, tc := range cases {
		f := func(t *testing.T) {
			seq := slices.Values(tc.elems)
			got := iterutil.Enumerate[int](seq)
			assertEqual2(t, got, tc.want, tc.breakWhen)
		}
		t.Run(tc.desc, f)
	}
}

func ExampleConcat() {
	seq1 := slices.Values([]string{"foo", "bar"})
	seq2 := slices.Values([]string{"baz", "qux"})
	for s := range iterutil.Concat(seq1, seq2) {
		fmt.Println(s)
	}
	// Output:
	// foo
	// bar
	// baz
	// qux
}

func TestConcat(t *testing.T) {
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
			got := iterutil.Concat(seq1, seq2)
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
	}{
		{
			desc:      "negative count",
			elems:     []string{"one", "two", "three"},
			count:     -1,
			breakWhen: alwaysFalse[string],
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
	}{
		{
			desc:      "negative count",
			elems:     []string{"one", "two", "three"},
			count:     -1,
			breakWhen: alwaysFalse[string],
			want:      []string{"one", "two", "three"},
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
			seq := slices.Values(tc.elems)
			got := iterutil.Drop(seq, tc.count)
			assertEqual(t, got, tc.want, tc.breakWhen)
		}
		t.Run(tc.desc, f)
	}
}

func ExampleZip() {
	french := slices.Values([]string{"un", "deux", "trois", "quatre", "cinq"})
	english := slices.Values([]string{"one", "two", "three"})
	seq := iterutil.Zip(french, english)
	for f, e := range seq {
		fmt.Println(f, "=>", e)
	}
	// Output:
	// un => one
	// deux => two
	// trois => three
}

func TestZip(t *testing.T) {
	cases := []struct {
		desc      string
		keys      []string
		values    []string
		breakWhen func(string, string) bool
		want      []Pair[string, string]
	}{
		{
			desc:   "no break",
			keys:   []string{"un", "deux", "trois", "quatre", "cinq"},
			values: []string{"one", "two", "three"},
			want: []Pair[string, string]{
				{"un", "one"},
				{"deux", "two"},
				{"trois", "three"},
			},
			breakWhen: alwaysFalse2[string, string],
		}, {
			desc:   "break early",
			keys:   []string{"un", "deux", "trois", "quatre", "cinq"},
			values: []string{"one", "two", "three"},
			want: []Pair[string, string]{
				{"un", "one"},
				{"deux", "two"},
			},
			breakWhen: equal2("trois", "three"),
		},
	}
	for _, tc := range cases {
		f := func(t *testing.T) {
			keys := slices.Values(tc.keys)
			values := slices.Values(tc.values)
			got := iterutil.Zip(keys, values)
			assertEqual2(t, got, tc.want, tc.breakWhen)
		}
		t.Run(tc.desc, f)
	}
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
