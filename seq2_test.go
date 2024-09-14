package iterutil_test

import (
	"fmt"
	"iter"
	"slices"

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
