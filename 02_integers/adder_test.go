package integers

import "testing"
import "fmt"

func TestAdder(t *testing.T) {
	sum := Add(1, 2)
	expected := 3

	if sum != expected {
		t.Errorf("Sum was %d expected %d", sum, expected)
	}
}

func ExampleAdd() {
	sum := Add(1, 8)
    fmt.Println(sum)
    // Output: 9
}
