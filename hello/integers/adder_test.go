package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("Expected '%d' but got '%d", expected, sum)
	}
}

func ExampleAdd() {
	sum := Add(2, 6)
	fmt.Println(sum)
	// Output: 8
}
