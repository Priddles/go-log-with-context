package log

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
)

func Log(
	ctx context.Context,
	level LogLevel,
	message string,
	data ...interface{},
) {
	entry := logEntry{
		Level:   level,
		Message: message,
		Data:    data,
	}

	tracing, exists := GetTracing(ctx)
	if !exists {
		entry.LogErrors = append(entry.LogErrors, "missing tracing information")
	}
	entry.Tracing = tracing

	entryB, err := json.Marshal(entry)
	if err != nil {
		entry.LogErrors = append(entry.LogErrors, fmt.Sprintf("dropped data due to marshal error: %v", err))

		entry.Data = nil
		entryB, err = json.Marshal(entry)
		if err != nil {
			log.Printf("log: could not marshal log entry: %v", entry)
			return
		}
	}

	log.Println(string(entryB))
}
