package iterutil_test

import (
	"errors"
	"fmt"

	"github.com/jub0bs/iterutil"
)

func ExampleAllErrors() {
	err := errors.Join(
		fmt.Errorf(
			"err2: %w",
			errors.New("err1"),
		),
		errors.Join(
			errors.New("err3"),
			errors.New("err4"),
		),
	)
	for i, err := range iterutil.AllErrors(err) {
		fmt.Println(i, err)
	}
	// Output:
	// 0 err2: err1
	// err3
	// err4
	// 1 err2: err1
	// 2 err1
	// 3 err3
	// err4
	// 4 err3
	// 5 err4
}

func ExampleAllLeafErrors() {
	err := errors.Join(
		fmt.Errorf(
			"err2: %w",
			errors.New("err1"),
		),
		errors.Join(
			errors.New("err3"),
			errors.New("err4"),
		),
	)
	for i, err := range iterutil.AllLeafErrors(err) {
		fmt.Println(i, err)
	}
	// Output:
	// 0 err1
	// 1 err3
	// 2 err4
}
