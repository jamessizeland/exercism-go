package airportrobot

import "fmt"

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.

type Counter interface {
	Add(increment int)
	Value() int
}

func implCounter(counter Counter) error { return nil }

type Stats struct {
	value int
}

var _ = implCounter(&Stats{}) // design pattern to prompt Interface method coverage.

func (s *Stats) Add(v int) {
	s.value += v
}
func (s *Stats) Value() int {
	return s.value
}

type Greeter interface {
	LanguageName() string
	Greet(name string) string
}

func implGreeter(greeter Greeter) error { return nil }

func SayHello(name string, greeter Greeter) string {
	return fmt.Sprintf("I can speak %s: %s", greeter.LanguageName(), greeter.Greet(name))
}

type Italian struct{}
type Portuguese struct{}

var _ = implGreeter(Italian{})
var _ = implGreeter(Portuguese{})

func (i Italian) LanguageName() string {
	return "Italian"
}
func (i Italian) Greet(name string) string {
	return fmt.Sprintf("Ciao %s!", name)
}

func (p Portuguese) LanguageName() string {
	return "Portuguese"
}
func (p Portuguese) Greet(name string) string {
	return fmt.Sprintf("Olá %s!", name)
}
