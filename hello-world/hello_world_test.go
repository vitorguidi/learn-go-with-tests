package hello_world

import "testing"

func Test_hello_world(t *testing.T) {
	t.Run("Check if function call for empty string defaults to world in english", func(t *testing.T) {
		assertValue(t, "Hello, world!", hello_world("english", ""))
	})
	t.Run("Check if function call for empty string defaults to mundo in spanish", func(t *testing.T) {
		assertValue(t, "Hola, mundo!", hello_world("spanish", ""))
	})
	t.Run("Check if non empty word returns correctly in english", func(t *testing.T) {
		assertValue(t, "Hello, vitor!", hello_world("english", "vitor"))
	})
	t.Run("Check if non empty word returns correctly in spanish", func(t *testing.T) {
		assertValue(t, "Hola, vitor!", hello_world("spanish", "vitor"))
	})
	t.Run("Check if unrecognized language panics", func(t *testing.T) {
		defer func() {
			_ = recover()
		}()
		hello_world("mumbling", "")
		t.Error("Expected panic, got a valid result")
	})
}

func assertValue(t testing.TB, expected string, actual string) {
	t.Helper()
	if expected != actual {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}
