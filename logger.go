package golog

import (
	"fmt"
	"log"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// Functionallity
// 1.- Simple Log
// displays a log contaning the File, Line and Timestamp of where it has been invocated.
// 2.- Status Log
// displays a log and uses a different log method depending on the status given by the user.
// 3.- Message Log
// displays a log containing the message provided by the user.
// 4.- Fault Log
// displays a log with an error message and uses a different log method depending on the class given by the user.
// 5.- Complete Log
// displays a log containing all the default and optional parameters.

// The following strings will be used
// to format the data provided by the Info struct.
var (
	output         string
	simpleFormat   = "\n| File: %s\n| Line: %s\n| Timestamp: %s\n"
	statusFormat   = "\n| File: %s\n| Line: %s\n| Timestamp: %s\n| Status: %s\n"
	messageFormat  = "\n| File: %s\n| Line: %s\n| Timestamp: %s\n| Message: %s\n"
	faultFormat    = "\n| File: %s\n| Line: %s\n| Timestamp: %s\n| Fault: %s\n"
	completeFormat = "\n| File: %s\n| Line: %s\n| Timestamp: %s\n| Status: %s\n| Message: %s\n| Fault: %s\n"
)

// Info represents the data structure of
// the information contained in the log message.
type Info struct {
	File      string
	Line      string
	Timestamp string
	Status    string
	Message   string
	Fault     string
}

// Simple fires a log.Print containing the following information:
// File, Line and Timestamp.
func Simple() {
	// Get rid if some log flags.
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))

	// Get the name of the file and the line
	// from where the method is being called.
	_, filename, line, _ := runtime.Caller(1)

	// Fill the struct fields with data.
	info := &Info{
		File:      filename,
		Line:      strconv.Itoa(line),
		Timestamp: time.Now().Format(time.RFC3339),
	}

	// Format the data.
	output = fmt.Sprintf(simpleFormat, info.File, info.Line, info.Timestamp)

	// Display the information.
	log.Print(output)

	// Output.
	// | File: path/to/file.go
	// | Line: 69
	// | Timestamp: 2022-01-13T16:38:46+01:00
}

// Status fires a different log method depending on the given status
// and contains the following information: File, Line, Timestamp and Status.
func Status(status string) {
	// Get rid if some log flags.
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))

	// Get the name of the file and the line
	// from where the method is being called.
	_, filename, line, _ := runtime.Caller(1)

	// Fill the struct fields with data.
	info := &Info{
		File:      filename,
		Line:      strconv.Itoa(line),
		Timestamp: time.Now().Format(time.RFC3339),
		Status:    strings.ToLower(status),
	}

	// Format the data.
	output = fmt.Sprintf(statusFormat, info.File, info.Line, info.Timestamp, info.Status)

	// Fires a different log method
	// depending on the status.
	switch info.Status {
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

	// Output.
	// | File: path/to/file.go
	// | Line: 107
	// | Timestamp: 2022-01-13T16:38:46+01:00
	// | Status: info
}

// Message fires a log print containing the following information:
// File, Line, Timestamp and Message.
func Message(message string) {
	// Get rid if some log flags.
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))

	// Get the name of the file and the line
	// from where the method is being called.
	_, filename, line, _ := runtime.Caller(1)

	// Fill the struct fields with data.
	info := &Info{
		File:      filename,
		Line:      strconv.Itoa(line),
		Timestamp: time.Now().Format(time.RFC3339),
		Message:   strings.ToLower(message),
	}

	// Format the data.
	output = fmt.Sprintf(messageFormat, info.File, info.Line, info.Timestamp, info.Message)

	// Display the information.
	log.Print(output)

	// Output.
	// | File: path/to/file.go
	// | Line: 107
	// | Timestamp: 2022-01-13T16:38:46+01:00
	// | Message: beep beep boop
}

// Fault fires a different log method depending on the given class
// and contains the following information: File, Line, Timestamp and Fault.
func Fault(class string, fault error) {
	// Get rid if some log flags.
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))

	// Get the name of the file and the line
	// from where the method is being called.
	_, filename, line, _ := runtime.Caller(1)

	// Fill the struct fields with data.
	info := &Info{
		File:      filename,
		Line:      strconv.Itoa(line),
		Timestamp: time.Now().Format(time.RFC3339),
		Fault:     fmt.Sprint(fault),
	}

	// Format the data.
	output = fmt.Sprintf(faultFormat, info.File, info.Line, info.Timestamp, info.Fault)

	// Fires a different log method
	// depending on the status.
	switch strings.ToLower(class) {
	case "fatal":
		log.Fatal(output)
	case "panic":
		log.Panic(output)
	default:
		log.Print(output)
	}

	// Output.
	// | File: path/to/file.go
	// | Line: 107
	// | Timestamp: 2022-01-13T16:38:46+01:00
	// | Fault: error message
}

// Complete fires a different log method depending on the given status
// and contains the following information: File, Line, Timestamp, Status, Message and Fault
func Complete(status, message string, fault error) {
	// Get rid if some log flags.
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))

	// Get the name of the file and the line
	// from where the method is being called.
	_, filename, line, _ := runtime.Caller(1)

	// Fill the struct fields with data.
	info := &Info{
		File:      filename,
		Line:      strconv.Itoa(line),
		Timestamp: time.Now().Format(time.RFC3339),
		Status:    strings.ToLower(status),
		Message:   strings.ToLower(message),
		Fault:     fmt.Sprint(fault),
	}

	// Format the data.
	output = fmt.Sprintf(completeFormat, info.File, info.Line, info.Timestamp, info.Status, info.Message, info.Fault)

	// Fires a different log method
	// depending on the status.
	switch info.Status {
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

	// Output.
	// | File: path/to/file.go
	// | Line: 69
	// | Timestamp: 2022-01-13T16:38:46+01:00
	// | Status: info
	// | Message: the logger is up and running
	// | Fault: <nil>
}
