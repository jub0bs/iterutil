package iterutil_test

import (
	"fmt"
	"slices"

	"github.com/jub0bs/iterutil"
)

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
