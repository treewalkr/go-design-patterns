package strategy

import (
	"fmt"
	"os"
	"time"
)

// LoggerStrategy defines the strategy interface for logging
type LoggerStrategy interface {
	Log(message string)
}

// MockLogger implements LoggerStrategy and mocks logging
type MockLogger struct{}

func (m *MockLogger) Log(message string) {
	fmt.Printf("[MOCK] %s\n", message)
}

// FileLogger implements LoggerStrategy and logs to a file
type FileLogger struct {
	file *os.File
}

func NewFileLogger(fileName string) *FileLogger {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(fmt.Sprintf("Failed to open log file: %s", err))
	}
	return &FileLogger{file: file}
}

func (f *FileLogger) Log(message string) {
	logMessage := fmt.Sprintf("[%s] %s\n", time.Now().Format(time.RFC3339), message)
	_, err := f.file.WriteString(logMessage)
	if err != nil {
		fmt.Printf("[ERROR] Failed to write to log file: %s\n", err)
	}
}

// BetterStackLogger implements LoggerStrategy and mocks logging to BetterStack log tail
type BetterStackLogger struct{}

func (b *BetterStackLogger) Log(message string) {
	fmt.Printf("[BETTERSTACK] %s\n", message) // In a real implementation, this would send to BetterStack API
}

// Logger is the context that uses a LoggerStrategy
type Logger struct {
	strategy LoggerStrategy
}

// SetStrategy sets the logging strategy dynamically
func (l *Logger) SetStrategy(strategy LoggerStrategy) {
	l.strategy = strategy
}

// Log logs a message using the current strategy
func (l *Logger) Log(message string) {
	if l.strategy == nil {
		fmt.Println("[ERROR] No logging strategy set")
		return
	}
	l.strategy.Log(message)
}

func Example() {
	// Initialize Logger context
	logger := &Logger{}

	// Use MockLogger
	mockLogger := &MockLogger{}
	logger.SetStrategy(mockLogger)
	logger.Log("This is a mock log message.")

	// Use FileLogger
	fileLogger := NewFileLogger("app.log")
	defer fileLogger.file.Close()
	logger.SetStrategy(fileLogger)
	logger.Log("This message is logged to a file.")

	// Use BetterStackLogger
	betterStackLogger := &BetterStackLogger{}
	logger.SetStrategy(betterStackLogger)
	logger.Log("This message is sent to BetterStack log tail.")
}
