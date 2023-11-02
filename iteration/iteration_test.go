package iteration

import "testing"

func TestIteration(t *testing.T) {
	expected := "bababa"
	actual := Repeat("ba", 3)
	if expected != actual {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("ba", 100)
	}
}
