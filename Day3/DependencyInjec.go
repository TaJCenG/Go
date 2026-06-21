package main

import "fmt"

// Interface
type Logger interface {
	Log(msg string)
}

// Implementation
type ConsoleLogger struct{}

func (c ConsoleLogger) Log(msg string) {
	fmt.Println("LOG:", msg)
}

// Service depends on Logger
type Service struct {
	logger Logger
}

func (s Service) DoWork() {
	s.logger.Log("Work started")
	// business logic...
	s.logger.Log("Work finished")
}

func main() {
	logger := ConsoleLogger{}
	service := Service{logger: logger} // Inject dependency
	service.DoWork()
}

// Logger is an interface.

// ConsoleLogger implements it.

// Service has a logger field of type Logger.

// In main, we inject ConsoleLogger into Service.
