package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	got := Repeat("a", 5)
	want := "aaaaa"
	if got != want {
		t.Errorf("got %s, but want %s", got, want)
	}
	t.Run("repeat any times", func(t *testing.T) {
		repeated := Repeat("a", 6)
		expected := "aaaaaa"
		if expected != repeated {
			t.Errorf("expected %s, but repeated %s", expected, repeated)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", 5)
	fmt.Println(repeated)
	// Output: aaaaa
}
