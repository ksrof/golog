package golog

import (
	"fmt"
	"log"
	"runtime"
	"strings"
	"time"
)

var (
	logger         Logger
	output         string
	simpleFormat   = "\n| File: %s\n| Line: %d\n| Timestamp: %s\n"
	statusFormat   = "\n| File: %s\n| Line: %d\n| Timestamp: %s\n| Status: %s\n"
	messageFormat  = "\n| File: %s\n| Line: %d\n| Timestamp: %s\n| Message: %s\n"
	faultFormat    = "\n| File: %s\n| Line: %d\n| Timestamp: %s\n| Fault: %s\n"
	completeFormat = "\n| File: %s\n| Line: %d\n| Timestamp: %s\n| Status: %s\n| Message: %s\n| Fault: %s\n"
)

// Logger represents the structure of the
// information contained in the log message.
type Logger struct {
	File      string
	Line      int
	Timestamp string
	Status    string
	Message   string
	Fault     string
}

// Simple fires a log.Print containing the following information:
// File, Line and Timestamp.
func Simple() {
	// Get rid of some log flags.
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))

	// Get the name of the file and the line
	// from where the method is being called.
	_, filename, line, ok := runtime.Caller(1)

	// Check if the caller was able to recover the information.
	if !ok {
		log.Fatal("unable to recover information")
	}

	// Fill the struct fields.
	logger.File = filename
	logger.Line = line
	logger.Timestamp = time.Now().Format(time.RFC3339)

	// Format the data.
	output = fmt.Sprintf(simpleFormat, logger.File, logger.Line, logger.Timestamp)

	// Display the information.
	log.Print(output)
}

// Status fires a different log method depending on a given status
// and contains the following information: File, Line, Timestamp and Status.
func Status(status string) {
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))

	_, filename, line, ok := runtime.Caller(1)

	if !ok {
		log.Fatal("unable to recover information")
	}

	logger.File = filename
	logger.Line = line
	logger.Timestamp = time.Now().Format(time.RFC3339)
	logger.Status = strings.ToLower(status)

	output = fmt.Sprintf(statusFormat, logger.File, logger.Line, logger.Timestamp, logger.Status)

	switch logger.Status {
	case "success":
		log.Print(output)
	case "info":
		log.Print(output)
	case "error":
		log.Print(output)
	case "fatal":
		log.Fatal(output)
	case "panic":
		log.Panic(output)
	default:
		log.Print(output)
	}
}

// Message fires a log.Print containing the following information:
// File, Line, Timestamp and Message.
func Message(message string) {
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))

	_, filename, line, ok := runtime.Caller(1)

	if !ok {
		log.Fatal("unable to recover information")
	}

	logger.File = filename
	logger.Line = line
	logger.Timestamp = time.Now().Format(time.RFC3339)
	logger.Message = strings.ToLower(message)

	output = fmt.Sprintf(messageFormat, logger.File, logger.Line, logger.Timestamp, logger.Message)

	log.Print(output)
}

// Fault fires a different log method depending on a given class
// and contains the following information: File, Line, Timestamp and Fault.
func Fault(class string, fault error) {
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))

	_, filename, line, ok := runtime.Caller(1)

	if !ok {
		log.Fatal("unable to recover information")
	}

	logger.File = filename
	logger.Line = line
	logger.Timestamp = time.Now().Format(time.RFC3339)
	logger.Fault = fmt.Sprint(fault)

	output = fmt.Sprintf(faultFormat, logger.File, logger.Line, logger.Timestamp, logger.Fault)

	switch strings.ToLower(class) {
	case "fatal":
		log.Fatal(output)
	case "panic":
		log.Panic(output)
	default:
		log.Print(output)
	}
}

// Complete fires a different log method depending on a given status
// and contains the following information: File, Line, Timestamp, Status, Message and Fault.
func Complete(status, message string, fault error) {
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))

	_, filename, line, ok := runtime.Caller(1)

	if !ok {
		log.Fatal("unable to recover information")
	}

	logger.File = filename
	logger.Line = line
	logger.Timestamp = time.Now().Format(time.RFC3339)
	logger.Status = strings.ToLower(status)
	logger.Message = strings.ToLower(message)
	logger.Fault = fmt.Sprint(fault)

	output = fmt.Sprintf(completeFormat, logger.File, logger.Line, logger.Timestamp, logger.Status, logger.Message, logger.Fault)

	switch logger.Status {
	case "success":
		log.Print(output)
	case "info":
		log.Print(output)
	case "error":
		log.Print(output)
	case "fatal":
		log.Fatal(output)
	case "panic":
		log.Panic(output)
	default:
		log.Print(output)
	}
}
