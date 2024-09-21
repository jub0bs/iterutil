package iterutil_test

import (
	"fmt"
	"iter"
	"slices"

	"github.com/jub0bs/iterutil"
)

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
