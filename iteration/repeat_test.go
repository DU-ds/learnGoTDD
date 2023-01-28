package iteration

import (
	"fmt"
	"testing"
)


func TestRepeat(t *testing.T) {
	nn := 5
	repeated := Repeat("a", nn)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
	if len(repeated) != len(expected) {
		t.Errorf(expected)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", b.N)
	}
}

func ExampleRepeat(){
	repeated := Repeat("abc", 5)
	fmt.Println(repeated)
	// Output: "abcabcabcabcabc"
}
