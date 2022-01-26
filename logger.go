package golog

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/ksrof/gocolors"
)

// Format strings
const (
	simpleFormat   = "\n| File: %s\n| Line: %s\n| Timestamp: %s\n"
	faultFormat    = "\n| File: %s\n| Line: %s\n| Timestamp: %s\n| Fault: %s\n"
	completeFormat = "\n| File: %s\n| Line: %s\n| Timestamp: %s\n| Status: %s\n| Message: %s\n| Fault: %s\n"
)

/*
Logger represents the structure of the
information contained in the log message.
*/
type Logger struct {
	File      string `json:"file"`
	Line      string `json:"line"`
	Timestamp string `json:"timestamp"`
	Status    string `json:"status,omitempty"`
	Message   string `json:"message,omitempty"`
	Fault     string `json:"fault,omitempty"`
}

// Start creates a log file at the root of the current directory.
func Start() error {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("unable to get working directory: %v", err)
		return err
	}

	file, err := os.OpenFile(filepath.Clean(path.Join(dir, "golog.log")), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatalf("unable to create log file: %v", err)
		return err
	}

	err = file.Close()
	if err != nil {
		log.Fatalf("unable to close file: %v", err)
		return err
	}

	return nil
}

// Find looks for the log file in the current directory.
func Find() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("unable to get working directory: %v", err)
		return "", err
	}

	file, err := filepath.Glob(path.Join(dir, "golog.log"))
	if err != nil {
		log.Fatalf("unable to find the log file: %v", err)
		return "", err
	}

	logFile := strings.Join(file, "")

	return logFile, nil
}

/*
Save marshals the log struct and writes
it to the previosuly created log file.
*/
func Save(logger Logger) error {
	logFile, err := Find()
	if err != nil {
		log.Fatalf("unable to find the log file: %v", err)
		return err
	}

	file, err := os.OpenFile(filepath.Clean(logFile), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatalf("unable to create/append/open log file: %v", err)
		return err
	}

	marshaled, err := json.MarshalIndent(logger, "", " ")
	if err != nil {
		log.Fatalf("unable to marshal the log message: %v", err)
		return err
	}

	_, err = file.Write([]byte(marshaled))
	if err != nil {
		log.Fatalf("unable to write to file: %v", err)
		return err
	}

	err = file.Close()
	if err != nil {
		log.Fatalf("unable to close file: %v", err)
		return err
	}

	return nil
}

/*
Simple returns a formated string containing the following
fields: File, Line and Timestamp with a specific color.
It also takes a save parameter to determine whether or
not it should save the output to a log file.
*/
func Simple(save bool) {
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))

	_, filename, line, ok := runtime.Caller(1)
	if !ok {
		log.Fatal("unable to recover information")
	}

	if save {
		logger := Logger{
			File:      filename,
			Line:      strconv.Itoa(line),
			Timestamp: time.Now().Format(time.RFC3339),
		}

		err := Save(logger)
		if err != nil {
			log.Fatalf("unable to save log to file: %v", err)
		}
	}

	logger := Logger{
		File:      gocolors.Color(filename, "cyan", ""),
		Line:      gocolors.Color(strconv.Itoa(line), "cyan", ""),
		Timestamp: gocolors.Color(time.Now().Format(time.RFC3339), "cyan", ""),
	}

	output := fmt.Sprintf(simpleFormat, logger.File, logger.Line, logger.Timestamp)
	log.Print(output)
}

/*
Complete returns a formated string containing the following
fields: File, Line, Timestamp, Status, Message and Fault with
a specific color. Depending on the status parameter a different
log method will be called. It also takes a save parameter to
determine whether or not it should save the output to a log file.
*/
func Complete(status, message string, fault error, save bool) {
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))

	_, filename, line, ok := runtime.Caller(1)
	if !ok {
		log.Fatal("unable to recover information")
	}

	if save {
		logger := Logger{
			File:      filename,
			Line:      strconv.Itoa(line),
			Timestamp: time.Now().Format(time.RFC3339),
			Status:    strings.ToLower(status),
			Message:   strings.ToLower(message),
			Fault:     fmt.Sprint(fault),
		}

		err := Save(logger)
		if err != nil {
			log.Fatalf("unable to save log to file: %v", err)
		}
	}

	switch strings.ToLower(status) {
	case "success":
		logger := Logger{
			File:      gocolors.Color(filename, "green", ""),
			Line:      gocolors.Color(strconv.Itoa(line), "green", ""),
			Timestamp: gocolors.Color(time.Now().Format(time.RFC3339), "green", ""),
			Status:    gocolors.Color(strings.ToLower(status), "green", ""),
			Message:   gocolors.Color(strings.ToLower(message), "green", ""),
			Fault:     gocolors.Color(fmt.Sprint(fault), "green", ""),
		}

		output := fmt.Sprintf(completeFormat, logger.File, logger.Line, logger.Timestamp, logger.Status, logger.Message, logger.Fault)
		log.Print(output)
	case "info":
		logger := Logger{
			File:      gocolors.Color(filename, "blue", ""),
			Line:      gocolors.Color(strconv.Itoa(line), "blue", ""),
			Timestamp: gocolors.Color(time.Now().Format(time.RFC3339), "blue", ""),
			Status:    gocolors.Color(strings.ToLower(status), "blue", ""),
			Message:   gocolors.Color(strings.ToLower(message), "blue", ""),
			Fault:     gocolors.Color(fmt.Sprint(fault), "blue", ""),
		}

		output := fmt.Sprintf(completeFormat, logger.File, logger.Line, logger.Timestamp, logger.Status, logger.Message, logger.Fault)
		log.Print(output)
	case "warning":
		logger := Logger{
			File:      gocolors.Color(filename, "yellow", ""),
			Line:      gocolors.Color(strconv.Itoa(line), "yellow", ""),
			Timestamp: gocolors.Color(time.Now().Format(time.RFC3339), "yellow", ""),
			Status:    gocolors.Color(strings.ToLower(status), "yellow", ""),
			Message:   gocolors.Color(strings.ToLower(message), "yellow", ""),
			Fault:     gocolors.Color(fmt.Sprint(fault), "yellow", ""),
		}

		output := fmt.Sprintf(completeFormat, logger.File, logger.Line, logger.Timestamp, logger.Status, logger.Message, logger.Fault)
		log.Print(output)
	case "fatal":
		logger := Logger{
			File:      gocolors.Color(filename, "magenta", ""),
			Line:      gocolors.Color(strconv.Itoa(line), "magenta", ""),
			Timestamp: gocolors.Color(time.Now().Format(time.RFC3339), "magenta", ""),
			Status:    gocolors.Color(strings.ToLower(status), "magenta", ""),
			Message:   gocolors.Color(strings.ToLower(message), "magenta", ""),
			Fault:     gocolors.Color(fmt.Sprint(fault), "magenta", ""),
		}

		output := fmt.Sprintf(completeFormat, logger.File, logger.Line, logger.Timestamp, logger.Status, logger.Message, logger.Fault)
		log.Fatal(output)
	case "panic":
		logger := Logger{
			File:      gocolors.Color(filename, "red", ""),
			Line:      gocolors.Color(strconv.Itoa(line), "red", ""),
			Timestamp: gocolors.Color(time.Now().Format(time.RFC3339), "red", ""),
			Status:    gocolors.Color(strings.ToLower(status), "red", ""),
			Message:   gocolors.Color(strings.ToLower(message), "red", ""),
			Fault:     gocolors.Color(fmt.Sprint(fault), "red", ""),
		}

		output := fmt.Sprintf(completeFormat, logger.File, logger.Line, logger.Timestamp, logger.Status, logger.Message, logger.Fault)
		log.Panic(output)
	default:
		logger := Logger{
			File:      filename,
			Line:      strconv.Itoa(line),
			Timestamp: time.Now().Format(time.RFC3339),
			Status:    strings.ToLower(status),
			Message:   strings.ToLower(message),
			Fault:     fmt.Sprint(fault),
		}

		output := fmt.Sprintf(completeFormat, logger.File, logger.Line, logger.Timestamp, logger.Status, logger.Message, logger.Fault)
		log.Print(output)
	}
}

/*
Fault returns a formated string containing the following
fields: File, Line, Timestamp and Fault with a specific color.
Depending on the status parameter a different log method will be called.
It also takes a save parameter to determine whether or not it should
save the output to a log file.
*/
func Fault(status string, fault error, save bool) {
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))

	_, filename, line, ok := runtime.Caller(1)
	if !ok {
		log.Fatal("unable to recover information")
	}

	if save {
		logger := Logger{
			File:      filename,
			Line:      strconv.Itoa(line),
			Timestamp: time.Now().Format(time.RFC3339),
			Fault:     fmt.Sprint(fault),
		}

		err := Save(logger)
		if err != nil {
			log.Fatalf("unable to save log to file: %v", err)
		}
	}

	switch strings.ToLower(status) {
	case "warning":
		logger := Logger{
			File:      gocolors.Color(filename, "yellow", ""),
			Line:      gocolors.Color(strconv.Itoa(line), "yellow", ""),
			Timestamp: gocolors.Color(time.Now().Format(time.RFC3339), "yellow", ""),
			Fault:     gocolors.Color(fmt.Sprint(fault), "yellow", ""),
		}

		output := fmt.Sprintf(faultFormat, logger.File, logger.Line, logger.Timestamp, logger.Fault)
		log.Print(output)
	case "fatal":
		logger := Logger{
			File:      gocolors.Color(filename, "magenta", ""),
			Line:      gocolors.Color(strconv.Itoa(line), "magenta", ""),
			Timestamp: gocolors.Color(time.Now().Format(time.RFC3339), "magenta", ""),
			Fault:     gocolors.Color(fmt.Sprint(fault), "magenta", ""),
		}

		output := fmt.Sprintf(faultFormat, logger.File, logger.Line, logger.Timestamp, logger.Fault)
		log.Fatal(output)
	case "panic":
		logger := Logger{
			File:      gocolors.Color(filename, "red", ""),
			Line:      gocolors.Color(strconv.Itoa(line), "red", ""),
			Timestamp: gocolors.Color(time.Now().Format(time.RFC3339), "red", ""),
			Fault:     gocolors.Color(fmt.Sprint(fault), "red", ""),
		}

		output := fmt.Sprintf(faultFormat, logger.File, logger.Line, logger.Timestamp, logger.Fault)
		log.Panic(output)
	default:
		logger := Logger{
			File:      filename,
			Line:      strconv.Itoa(line),
			Timestamp: time.Now().Format(time.RFC3339),
			Fault:     fmt.Sprint(fault),
		}

		output := fmt.Sprintf(faultFormat, logger.File, logger.Line, logger.Timestamp, logger.Fault)
		log.Print(output)
	}
}
