package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	n := 10
	got := Repeat("a", n)
	expected := "aaaaaaaaaa"

	if got != expected {
		t.Errorf("expected %q, but got %q", expected, got)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}

func ExampleRepeat() {
	repeated := Repeat("b", 3)
	fmt.Println(repeated)

	// Output: bbb
}