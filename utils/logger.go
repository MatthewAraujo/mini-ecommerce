package utils

import (
	"fmt"
	"time"
)

type ParentLogger struct {
	context string
	route   string
}

func NewParentLogger(context string) *ParentLogger {
	fmt.Printf("Logger inicializado para: %s\n", context)
	return &ParentLogger{context: context}
}

func (p *ParentLogger) Info(routeOrMessage string, message ...string) {
	p.log("INFO", routeOrMessage, message...)
}

func (p *ParentLogger) Warn(routeOrMessage string, message ...string) {
	p.log("WARN", routeOrMessage, message...)
}

func (p *ParentLogger) log(level, routeOrMessage string, message ...string) {
	if len(message) > 0 {
		p.route = routeOrMessage
		routeOrMessage = message[0]
	}

	timestamp := time.Now().Format("02/01/2006 15:04:05")
	fmt.Printf("[%s] %s - Context: %s | Router: %s - %s\n",
		level, timestamp, p.context, p.route, routeOrMessage)
}

func (p *ParentLogger) LogError(routeOrMessage string, err error, message ...string) {
	if len(message) > 0 {
		p.route = routeOrMessage
		routeOrMessage = message[0]
	}

	timestamp := time.Now().Format("02/01/2006 15:04:05")
	fmt.Printf("[ERROR] %s - Context: %s | Router: %s - %s | ERROR: %v\n",
		timestamp, p.context, p.route, routeOrMessage, err)
}
