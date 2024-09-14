package iterutil_test

import (
	"fmt"
	"slices"

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
