package dependency_injection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	writer := &bytes.Buffer{}
	msg := "DI is neat! Can even do it by passing arguments!"
	Greet(writer, msg)
	actual := writer.String()
	if writer.String() != msg {
		t.Errorf("Expected %s, got %s", msg, actual)
	}
}
