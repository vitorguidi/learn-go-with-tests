package mocking

import (
	"fmt"
	"io"
	"time"
)

const write = "write"
const sleep = "sleep"

type Sleeper interface {
	Sleep(duration time.Duration)
}

type SpyCountDownOperations struct {
	calls []string
}

func (s *SpyCountDownOperations) Sleep(_ time.Duration) {
	s.calls = append(s.calls, sleep)
}

func (s *SpyCountDownOperations) Write(_ []byte) (n int, err error) {
	s.calls = append(s.calls, write)
	return
}

func Countdown(writer io.Writer, sleeper Sleeper) {
	fmt.Fprintln(writer, "3")
	sleeper.Sleep(time.Second)
	fmt.Fprintln(writer, "2")
	sleeper.Sleep(time.Second)
	fmt.Fprintln(writer, "1")
	sleeper.Sleep(time.Second)
	fmt.Fprintln(writer, "Go!")
}
