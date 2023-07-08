package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("repeats a character", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})
	t.Run("repeats a character x times", func(t *testing.T) {
		repeated := Repeat("x", 3)
		expected := "xxx"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 20)
	}
}

func ExampleRepeat() {
	hey := Repeat("hey ", 4)
	fmt.Println(hey)
	// Output: hey hey hey hey
}
