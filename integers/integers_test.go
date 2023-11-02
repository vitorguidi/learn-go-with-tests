package integers

import "testing"

func TestAdd(t *testing.T) {
	t.Run("2 plus 2 should be 4", func(t *testing.T) {
		expected := 4
		actual := Add(2, 2)
		if expected != actual {
			t.Errorf("Expected %d, got %d", expected, actual)
		}
	})
}
