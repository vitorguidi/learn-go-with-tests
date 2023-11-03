package mocking

import (
	"bytes"
	"reflect"
	"testing"
)

func TestCountdown(t *testing.T) {
	t.Run("All operations should run in correct order", func(t *testing.T) {
		s := &SpyCountDownOperations{make([]string, 0)}
		Countdown(s, s)

		actualOps := s.calls
		expectedOps := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(expectedOps, actualOps) {
			t.Errorf("Ops assertion failed: expected %s, got %s", expectedOps, actualOps)
		}
	})

	t.Run("Should produce correct text output", func(t *testing.T) {
		b := bytes.Buffer{}
		s := &SpyCountDownOperations{make([]string, 0)}
		Countdown(&b, s)
		actualText := b.String()
		expectedText := `3
2
1
Go!
`
		if expectedText != b.String() {
			t.Errorf("Text assertion failed: expected %s, got %s", expectedText, actualText)
		}
	})

}
