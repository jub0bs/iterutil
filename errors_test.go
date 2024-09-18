package iterutil_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/jub0bs/iterutil"
)

func ExampleAllErrors() {
	err := errors.Join(
		fmt.Errorf(
			"err1: %w",
			errors.New("err0"),
		),
		errors.Join(
			errors.New("err2"),
			errors.New("err3"),
		),
	)
	for i, err := range iterutil.AllErrors(err) {
		fmt.Println(i, err)
	}
	// Output:
	// 0 err1: err0
	// err2
	// err3
	// 1 err1: err0
	// 2 err0
	// 3 err2
	// err3
	// 4 err2
	// 5 err3
}

func TestAllErrors(t *testing.T) {
	cases := []struct {
		desc      string
		err       error
		want      []Pair[int, error]
		breakWhen func(int, error) bool
	}{
		{
			desc: "singleton",
			err:  err0,
			want: []Pair[int, error]{
				{0, err0},
			},
			breakWhen: alwaysFalse2[int, error],
		}, {
			desc: "multi-error no break",
			err:  err4,
			want: []Pair[int, error]{
				{0, err4},
				{1, err2},
				{2, err3},
			},
			breakWhen: alwaysFalse2[int, error],
		}, {
			desc: "multi-error break early",
			err:  err4,
			want: []Pair[int, error]{
				{0, err4},
				{1, err2},
			},
			breakWhen: equal2(2, err3),
		}, {
			desc: "wrapped error no break",
			err:  err1,
			want: []Pair[int, error]{
				{0, err1},
				{1, err0},
			},
			breakWhen: alwaysFalse2[int, error],
		}, {
			desc: "wrapped error break early",
			err:  err1,
			want: []Pair[int, error]{
				{0, err1},
			},
			breakWhen: equal2(1, err0),
		}, {
			desc:      "complex error tree no break",
			err:       err5,
			breakWhen: alwaysFalse2[int, error],
			want: []Pair[int, error]{
				{0, err5},
				{1, err1},
				{2, err0},
				{3, err4},
				{4, err2},
				{5, err3},
			},
		},
	}
	for _, tc := range cases {
		f := func(t *testing.T) {
			got := iterutil.AllErrors(tc.err)
			assertEqual2(t, got, tc.want, tc.breakWhen)
		}
		t.Run(tc.desc, f)
	}
}

var (
	err0 = errors.New("err0")
	err1 = fmt.Errorf("err1: %w", err0)
	err2 = errors.New("err2")
	err3 = errors.New("err3")
	err4 = errors.Join(err2, err3)
	err5 = errors.Join(err1, err4)
)
