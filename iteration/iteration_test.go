package iteration

import "testing"

func TestIteration(t *testing.T) {
	expected := "bababa"
	actual := Repeat("ba", 3)
	if expected != actual {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestBadLoopVarUsage(t *testing.T) {
	// what happens is that we store the pointer to the loop variable
	// which at the end is 3, so we get 3*3 = 9
	expected := 9
	actual := BadLoopVarUsage()
	if expected != actual {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("ba", 100)
	}
}
