package iterutil_test

import (
	"fmt"
	"iter"
	"slices"

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

func ExampleHead() {
	seq := slices.Values([]int{})
	fmt.Println(iterutil.Head(seq))
	seq = slices.Values([]int{1, 2, 3, 4})
	fmt.Println(iterutil.Head(seq))
	// Output:
	// 0 false
	// 1 true
}

func ExampleTail() {
	seq := slices.Values([]int{})
	tail, ok := iterutil.Tail(seq)
	if ok {
		fmt.Println(slices.Collect(tail))
	}
	seq = slices.Values([]int{1, 2, 3, 4})
	tail, ok = iterutil.Tail(seq)
	if ok {
		fmt.Println(slices.Collect(tail))
	}
	// Output: [2 3 4]
}

func ExampleUncons() {
	seq := slices.Values([]int{})
	head, tail, ok := iterutil.Uncons(seq)
	if ok {
		fmt.Println(head, slices.Collect(tail))
	}
	seq = slices.Values([]int{1, 2, 3, 4})
	head, tail, ok = iterutil.Uncons(seq)
	if ok {
		fmt.Println(head, slices.Collect(tail))
	}
	// Output: 1 [2 3 4]
}

func ExampleAppend() {
	seq1 := slices.Values([]string{"foo", "bar"})
	seq2 := slices.Values([]string{"baz", "qux"})
	for s := range iterutil.Append(seq1, seq2) {
		fmt.Println(s)
	}
	// Output:
	// foo
	// bar
	// baz
	// qux
}

func ExampleConcat() {
	seq1 := slices.Values([]string{"foo", "bar"})
	seq2 := slices.Values([]string{"baz", "qux"})
	seqs := slices.Values([]iter.Seq[string]{seq1, seq2})
	for s := range iterutil.Concat(seqs) {
		fmt.Println(s)
	}
	// Output:
	// foo
	// bar
	// baz
	// qux
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

func ExampleDrop() {
	seq := slices.Values([]string{"foo", "bar", "baz", "qux"})
	for s := range iterutil.Drop(seq, 3) {
		fmt.Println(s)
	}
	// Output:
	// qux
}

func ExampleAt() {
	seq := slices.Values([]string{"foo", "bar", "baz", "qux"})
	fmt.Println(iterutil.At(seq, 2))
	// Output:
	// baz true
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
