package vcb

import (
	"testing"
)

func TestGreet(t *testing.T) {
	result := Greet()
	criteria := "Let's Go!"
	if result != criteria {
		t.Fatalf(`Greet() = %q, should be %q`, result, criteria)
	}
}
