package stringutils_test

import (
	"fmt"
	"testing"

	"github.com/umuttopalak/stringutils"
)

func TestReverse(t *testing.T) {
	tcs := map[string]struct {
		input []string
		want  []string
	}{
		"none Turkish letters": {
			input: []string{"hello", "this is vigo"},
			want:  []string{"olleh", "ogiv si siht"},
		},
		"with Turkish letters": {
			input: []string{"uğur", "kırmızı şapka ve ÖĞRENCİ"},
			want:  []string{"ruğu", "İCNERĞÖ ev akpaş ızımrık"},
		},
		"with German letters": {
			input: []string{"Präzisionsmeßgerät"},
			want:  []string{"täregßemsnoisizärP"},
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			for i, in := range tc.input {
				got := stringutils.Reverse(in)

				if got != tc.want[i] {
					fmt.Println(len(got), len(tc.want[i]))
					t.Errorf("want: %v; got: %v", tc.want[i], got)
				}
			}
		})
	}
}

var gs string

func BenchmarkReverse(b *testing.B) {
	var s string
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s = stringutils.Reverse("merhaba dünya!")
	}

	gs = s
}

func ExampleReverse() {
	fmt.Println(stringutils.Reverse("vigo"))
	// Output: ogiv
}
