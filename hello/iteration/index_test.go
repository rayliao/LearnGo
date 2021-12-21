package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	const times = 6
	repeated := Repeat("b", times)
	expected := "bbbbbb"

	if repeated != expected {
		t.Errorf("expected '%q' but got '%q'", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("d", 4)
	}
}

func ExampleRepeat() {
	repeated := Repeat("d", 5)
	fmt.Println(repeated)
	// Output: ddddd
}
