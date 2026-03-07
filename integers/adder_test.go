package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(1, 2)
	expected := 3
	if sum != expected {
		t.Errorf("expected result %d but got %d", expected, sum)
	}
}

func ExampleAdd() {
	sum := Add(6, 7)
	fmt.Println(sum)
	// Output: 13
}
