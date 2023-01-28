package iteration

import "testing"

func TestRepeat(t *testing.T, n int) {
	repeated := Repeat("a", n)
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
		Repeat("a")
	}
}
