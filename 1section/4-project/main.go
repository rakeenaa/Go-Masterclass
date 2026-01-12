package main

import "fmt"

type LogLevel int

const (
	LevelTrace LogLevel = iota
	LevelDebug
	LevelInfo
	LevelWarning
	LevelError
	LevelCritical
)

var levelNames = []string{"Trace", "Debug", "Info", "Warning", "Error", "Critical"}

func (l LogLevel) String() string {
	if l < LevelTrace || l > LevelCritical {
		return "Unknown"
	}

	return levelNames[l]
}

func printLogLevel(level LogLevel) {
	fmt.Printf("Log level: %d %s\n", level, level.String())
}
func main() {

	printLogLevel(LevelTrace)
	printLogLevel(LevelDebug)
	printLogLevel(LevelInfo)
	printLogLevel(LevelWarning)
	printLogLevel(LevelError)
	printLogLevel(LevelCritical)
	printLogLevel(10)

}
