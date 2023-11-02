package hello_world

import "fmt"

const english = "english"
const englishPrefix = "Hello"
const spanish = "spanish"
const spanishPrefix = "Hola"

func hello_world(language string, name string) string {
	var prefix string
	switch language {
	case spanish:
		if name == "" {
			name = "mundo"
		}
		prefix = spanishPrefix
	case english:
		if name == "" {
			name = "world"
		}
		prefix = englishPrefix
	default:
		panic("Language not recognized.")
	}
	return fmt.Sprintf("%s, %s!", prefix, name)
}
