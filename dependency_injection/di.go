package dependency_injection

import (
	"fmt"
	"io"
)

func Greet(writer io.Writer, text string) {
	fmt.Fprint(writer, text)
}
