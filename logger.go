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

var (
	logger         Logger
	simpleFormat   = "\n| File: %s\n| Line: %s\n| Timestamp: %s\n"
	statusFormat   = "\n| File: %s\n| Line: %s\n| Timestamp: %s\n| Status: %s\n"
	messageFormat  = "\n| File: %s\n| Line: %s\n| Timestamp: %s\n| Message: %s\n"
	faultFormat    = "\n| File: %s\n| Line: %s\n| Timestamp: %s\n| Fault: %s\n"
	completeFormat = "\n| File: %s\n| Line: %s\n| Timestamp: %s\n| Status: %s\n| Message: %s\n| Fault: %s\n"
)

// Logger represents the structure of the
// information contained in the log message.
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

	// Create the log file.
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

// Find looks for a log file in the current directory.
func Find() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("unable to get working directory: %v", err)
		return "", err
	}

	// Look for the log file.
	file, err := filepath.Glob(path.Join(dir, "golog.log"))
	if err != nil {
		log.Fatalf("unable to find the log file: %v", err)
		return "", err
	}

	logFile := strings.Join(file, "")

	return logFile, nil
}

// Save stores the json log outputs to the log file.
func Save(jsonLog string) error {
	// Find the log file.
	logFile, err := Find()
	if err != nil {
		log.Fatalf("unable to find the log file: %v", err)
		return err
	}

	// Save to log file.
	file, err := os.OpenFile(filepath.Clean(logFile), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatalf("unable to create/append/open log file: %v", err)
		return err
	}

	_, err = file.Write([]byte(jsonLog))
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

// Format formats the logger struct to a readable output.
func Format(logger Logger) string {
	if len(logger.Status) <= 0 && len(logger.Message) <= 0 && len(logger.Fault) <= 0 {
		output := fmt.Sprintf(simpleFormat, logger.File, logger.Line, logger.Timestamp)
		return output
	}

	if len(logger.Status) > 0 && len(logger.Message) > 0 && len(logger.Fault) > 0 {
		output := fmt.Sprintf(completeFormat, logger.File, logger.Line, logger.Timestamp, logger.Status, logger.Message, logger.Fault)
		return output
	}

	if len(logger.Status) > 0 {
		output := fmt.Sprintf(statusFormat, logger.File, logger.Line, logger.Timestamp, logger.Status)
		return output
	}

	if len(logger.Message) > 0 {
		output := fmt.Sprintf(messageFormat, logger.File, logger.Line, logger.Timestamp, logger.Message)
		return output
	}

	if len(logger.Fault) > 0 {
		output := fmt.Sprintf(faultFormat, logger.File, logger.Line, logger.Timestamp, logger.Fault)
		return output
	}

	return ""
}

// Simple fires a log.Print containing the following information:
// File, Line and Timestamp.
// (Optional) the output can be saved to a log file as JSON format.
func Simple(save bool) {
	// Get rid of some log flags.
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))

	// Get the name of the file and the line
	// from where the method is being called.
	_, filename, line, ok := runtime.Caller(1)

	// Check if the caller was able to recover the information.
	if !ok {
		log.Fatal("unable to recover information")
	}

	// Fill the struct fields without color.
	// (for the json log file)
	logger = Logger{
		File:      filename,
		Line:      strconv.Itoa(line),
		Timestamp: time.Now().Format(time.RFC3339),
	}

	// (Optional) save JSON output to log file.
	if save {
		jsonLog, err := JSON(logger)
		if err != nil {
			log.Fatalf("unable to marshal log message: %v", err)
		}

		err = Save(jsonLog)
		if err != nil {
			log.Fatalf("unable to save log message: %v", err)
		}
	}

	// Fill the struct fields with color.
	logger = Logger{
		File:      gocolors.Cyan(filename, ""),
		Line:      gocolors.Cyan(strconv.Itoa(line), ""),
		Timestamp: gocolors.Cyan(time.Now().Format(time.RFC3339), ""),
	}

	// Format the log message.
	output := Format(logger)

	// Display the information.
	log.Print(output)
}

// Status fires a different log method depending on a given status
// and contains the following information: File, Line, Timestamp and Status.
// (Optional) the output can be saved to a log file as JSON format.
func Status(status string, save bool) {
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))

	_, filename, line, ok := runtime.Caller(1)

	if !ok {
		log.Fatal("unable to recover information")
	}

	logger = Logger{
		File:      filename,
		Line:      strconv.Itoa(line),
		Timestamp: time.Now().Format(time.RFC3339),
		Status:    strings.ToLower(status),
	}

	if save {
		jsonLog, err := JSON(logger)
		if err != nil {
			log.Fatalf("unable to marshal log message: %v", err)
		}

		err = Save(jsonLog)
		if err != nil {
			log.Fatalf("unable to save log message: %v", err)
		}
	}

	switch strings.ToLower(status) {
	case "success":
		logger = Logger{
			File:      gocolors.Green(filename, ""),
			Line:      gocolors.Green(strconv.Itoa(line), ""),
			Timestamp: gocolors.Green(time.Now().Format(time.RFC3339), ""),
			Status:    gocolors.Green(strings.ToLower(status), ""),
		}

		output := Format(logger)
		log.Print(output)
	case "info":
		logger = Logger{
			File:      gocolors.Blue(filename, ""),
			Line:      gocolors.Blue(strconv.Itoa(line), ""),
			Timestamp: gocolors.Blue(time.Now().Format(time.RFC3339), ""),
			Status:    gocolors.Blue(strings.ToLower(status), ""),
		}

		output := Format(logger)
		log.Print(output)
	case "error":
		logger = Logger{
			File:      gocolors.Yellow(filename, ""),
			Line:      gocolors.Yellow(strconv.Itoa(line), ""),
			Timestamp: gocolors.Yellow(time.Now().Format(time.RFC3339), ""),
			Status:    gocolors.Yellow(strings.ToLower(status), ""),
		}

		output := Format(logger)
		log.Print(output)
	case "fatal":
		logger = Logger{
			File:      gocolors.Magenta(filename, ""),
			Line:      gocolors.Magenta(strconv.Itoa(line), ""),
			Timestamp: gocolors.Magenta(time.Now().Format(time.RFC3339), ""),
			Status:    gocolors.Magenta(strings.ToLower(status), ""),
		}

		output := Format(logger)
		log.Fatal(output)
	case "panic":
		logger = Logger{
			File:      gocolors.Red(filename, ""),
			Line:      gocolors.Red(strconv.Itoa(line), ""),
			Timestamp: gocolors.Red(time.Now().Format(time.RFC3339), ""),
			Status:    gocolors.Red(strings.ToLower(status), ""),
		}

		output := Format(logger)
		log.Panic(output)
	default:
		logger = Logger{
			File:      filename,
			Line:      strconv.Itoa(line),
			Timestamp: time.Now().Format(time.RFC3339),
			Status:    strings.ToLower(status),
		}

		output := Format(logger)
		log.Print(output)
	}
}

// Message fires a log.Print containing the following information:
// File, Line, Timestamp and Message.
// (Optional) the output can be saved to a log file as JSON format.
func Message(message string, save bool) {
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))

	_, filename, line, ok := runtime.Caller(1)

	if !ok {
		log.Fatal("unable to recover information")
	}

	logger = Logger{
		File:      filename,
		Line:      strconv.Itoa(line),
		Timestamp: time.Now().Format(time.RFC3339),
		Message:   strings.ToLower(message),
	}

	if save {
		jsonLog, err := JSON(logger)
		if err != nil {
			log.Fatalf("unable to marshal log message: %v", err)
		}

		err = Save(jsonLog)
		if err != nil {
			log.Fatalf("unable to save log message: %v", err)
		}
	}

	logger = Logger{
		File:      gocolors.Cyan(filename, ""),
		Line:      gocolors.Cyan(strconv.Itoa(line), ""),
		Timestamp: gocolors.Cyan(time.Now().Format(time.RFC3339), ""),
		Message:   gocolors.Cyan(strings.ToLower(message), ""),
	}

	output := Format(logger)
	log.Print(output)
}

// Fault fires a different log method depending on a given class
// and contains the following information: File, Line, Timestamp and Fault.
// (Optional) the output can be saved to a log file as JSON format.
func Fault(class string, fault error, save bool) {
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))

	_, filename, line, ok := runtime.Caller(1)

	if !ok {
		log.Fatal("unable to recover information")
	}

	logger = Logger{
		File:      filename,
		Line:      strconv.Itoa(line),
		Timestamp: time.Now().Format(time.RFC3339),
		Fault:     fmt.Sprint(fault),
	}

	if save {
		jsonLog, err := JSON(logger)
		if err != nil {
			log.Fatalf("unable to marshal log message: %v", err)
		}

		err = Save(jsonLog)
		if err != nil {
			log.Fatalf("unable to save log message: %v", err)
		}
	}

	switch strings.ToLower(class) {
	case "fatal":
		logger = Logger{
			File:      gocolors.Magenta(filename, ""),
			Line:      gocolors.Magenta(strconv.Itoa(line), ""),
			Timestamp: gocolors.Magenta(time.Now().Format(time.RFC3339), ""),
			Fault:     gocolors.Magenta(fmt.Sprint(fault), ""),
		}

		output := Format(logger)
		log.Fatal(output)
	case "panic":
		logger = Logger{
			File:      gocolors.Red(filename, ""),
			Line:      gocolors.Red(strconv.Itoa(line), ""),
			Timestamp: gocolors.Red(time.Now().Format(time.RFC3339), ""),
			Fault:     gocolors.Red(fmt.Sprint(fault), ""),
		}

		output := Format(logger)
		log.Panic(output)
	case "error":
		logger = Logger{
			File:      gocolors.Yellow(filename, ""),
			Line:      gocolors.Yellow(strconv.Itoa(line), ""),
			Timestamp: gocolors.Yellow(time.Now().Format(time.RFC3339), ""),
			Fault:     gocolors.Yellow(fmt.Sprint(fault), ""),
		}

		output := Format(logger)
		log.Print(output)
	}
}

// Complete fires a different log method depending on a given status
// and contains the following information: File, Line, Timestamp, Status, Message and Fault.
// (Optional) the output can be saved to a log file as JSON format.
func Complete(status, message string, fault error, save bool) {
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))

	_, filename, line, ok := runtime.Caller(1)

	if !ok {
		log.Fatal("unable to recover information")
	}

	logger = Logger{
		File:      filename,
		Line:      strconv.Itoa(line),
		Timestamp: time.Now().Format(time.RFC3339),
		Status:    strings.ToLower(status),
		Message:   strings.ToLower(message),
		Fault:     fmt.Sprint(fault),
	}

	if save {
		jsonLog, err := JSON(logger)
		if err != nil {
			log.Fatalf("unable to marshal log message: %v", err)
		}

		err = Save(jsonLog)
		if err != nil {
			log.Fatalf("unable to save log message: %v", err)
		}
	}

	switch strings.ToLower(status) {
	case "success":
		logger = Logger{
			File:      gocolors.Green(filename, ""),
			Line:      gocolors.Green(strconv.Itoa(line), ""),
			Timestamp: gocolors.Green(time.Now().Format(time.RFC3339), ""),
			Status:    gocolors.Green(strings.ToLower(status), ""),
			Message:   gocolors.Green(strings.ToLower(message), ""),
			Fault:     gocolors.Green(fmt.Sprint(fault), ""),
		}

		output := Format(logger)
		log.Print(output)
	case "info":
		logger = Logger{
			File:      gocolors.Blue(filename, ""),
			Line:      gocolors.Blue(strconv.Itoa(line), ""),
			Timestamp: gocolors.Blue(time.Now().Format(time.RFC3339), ""),
			Status:    gocolors.Blue(strings.ToLower(status), ""),
			Message:   gocolors.Blue(strings.ToLower(message), ""),
			Fault:     gocolors.Blue(fmt.Sprint(fault), ""),
		}

		output := Format(logger)
		log.Print(output)
	case "error":
		logger = Logger{
			File:      gocolors.Yellow(filename, ""),
			Line:      gocolors.Yellow(strconv.Itoa(line), ""),
			Timestamp: gocolors.Yellow(time.Now().Format(time.RFC3339), ""),
			Status:    gocolors.Yellow(strings.ToLower(status), ""),
			Message:   gocolors.Yellow(strings.ToLower(message), ""),
			Fault:     gocolors.Yellow(fmt.Sprint(fault), ""),
		}

		output := Format(logger)
		log.Print(output)
	case "fatal":
		logger = Logger{
			File:      gocolors.Magenta(filename, ""),
			Line:      gocolors.Magenta(strconv.Itoa(line), ""),
			Timestamp: gocolors.Magenta(time.Now().Format(time.RFC3339), ""),
			Status:    gocolors.Magenta(strings.ToLower(status), ""),
			Message:   gocolors.Magenta(strings.ToLower(message), ""),
			Fault:     gocolors.Magenta(fmt.Sprint(fault), ""),
		}

		output := Format(logger)
		log.Fatal(output)
	case "panic":
		logger = Logger{
			File:      gocolors.Red(filename, ""),
			Line:      gocolors.Red(strconv.Itoa(line), ""),
			Timestamp: gocolors.Red(time.Now().Format(time.RFC3339), ""),
			Status:    gocolors.Red(strings.ToLower(status), ""),
			Message:   gocolors.Red(strings.ToLower(message), ""),
			Fault:     gocolors.Red(fmt.Sprint(fault), ""),
		}

		output := Format(logger)
		log.Panic(output)
	default:
		logger = Logger{
			File:      filename,
			Line:      strconv.Itoa(line),
			Timestamp: time.Now().Format(time.RFC3339),
			Status:    strings.ToLower(status),
			Message:   strings.ToLower(message),
			Fault:     fmt.Sprint(fault),
		}

		output := Format(logger)
		log.Print(output)
	}
}

// JSON returns the log message output in JSON format.
func JSON(logger Logger) (string, error) {
	marshaled, err := json.MarshalIndent(logger, "", " ")
	if err != nil {
		log.Fatalf("unable to marshal the log message: %v", err)
		return "", err
	}

	return string(marshaled), nil
}
