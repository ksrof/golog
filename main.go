package main

import (
	"errors"
	"fmt"
	"log"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// TODO's
// 1.- Make a function that takes Status, Message and Fault(error) as parameters.
// 2.- Determine the type of log method using the Status parameter.
// 3.- Print the following information: File, Line, Timestamp, Status, Message and Fault(error).

// The string will be used to format the
// data provided by the Info struct.
var formatString = "\nFile: %s\nLine: %s\nTimestamp: %s\nStatus: %s\nMessage: %s\nFault: %s\n"

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

// Logger returns different log methods depending on the status
// and prints the following information: File, Line, Timestamp, Message, Status and Fault(err).
func Logger(status, message string, fault error) {
	// Get rid of some log flags.
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
		Message:   message,
		Fault:     fmt.Sprint(fault),
	}

	// Format the data.
	output := fmt.Sprintf(formatString, info.File, info.Line, info.Timestamp, info.Status, info.Message, info.Fault)

	// Change the log method depending
	// on the provided status.
	switch info.Status {
	case "success":
		log.Print(output)
	case "fatal":
		log.Fatal(output)
	}
}

func main() {
	// Let's try our logger
	alive := true

	// Fatal status
	if alive != true {
		err := errors.New("not alive")
		Logger("fatal", "The subject is not alive anymore", err)
	}

	// Success status
	if alive == true {
		Logger("success", "The subject is alive", nil)
	}
}