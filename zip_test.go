package iterutil_test

import (
	"fmt"
	"slices"
	"testing"

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
