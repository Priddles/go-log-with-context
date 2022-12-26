package logger

type LogLevel string

const (
	DEBUG LogLevel = "debug"
	INFO  LogLevel = "info"
	WARN  LogLevel = "warn"
	ERROR LogLevel = "error"
	FATAL LogLevel = "fatal"
)

type logEntry struct {
	Level     LogLevel        `json:"level"`
	Message   string          `json:"message"`
	Tracing   *TracingContext `json:"tracing"`
	Data      []interface{}   `json:"data"`
	LogErrors []string        `json:"logErrors,omitempty"`
}

type TracingContext struct {
	// Event bridge event ID.
	EventID string `json:"eventId,omitempty"`

	// Step function execution ID.
	ExecutionID string `json:"executionId,omitempty"`

	// API gateway request ID.
	RequestID string `json:"requestId,omitempty"`

	// X-Ray trace ID.
	TraceID string `json:"traceId,omitempty"`
}
