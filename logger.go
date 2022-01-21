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
	output         string
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

// JSON returns the log message output in JSON format.
func JSON(logger Logger) (string, error) {
	marshaled, err := json.MarshalIndent(logger, "", " ")
	if err != nil {
		log.Fatalf("unable to marshal the log message: %v", err)
		return "", err
	}

	return string(marshaled), nil
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

	// Fill the struct fields.
	// Apply colors to each field.
	logger.File = gocolors.Cyan(filename, "")
	logger.Line = gocolors.Cyan(strconv.Itoa(line), "")
	logger.Timestamp = gocolors.Cyan(time.Now().Format(time.RFC3339), "")

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

	// Format the data.
	output = fmt.Sprintf(simpleFormat, logger.File, logger.Line, logger.Timestamp)

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

	logger.File = filename
	logger.Line = strconv.Itoa(line)
	logger.Timestamp = time.Now().Format(time.RFC3339)
	logger.Status = strings.ToLower(status)

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
// (Optional) the output can be saved to a log file as JSON format.
func Message(message string, save bool) {
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))

	_, filename, line, ok := runtime.Caller(1)

	if !ok {
		log.Fatal("unable to recover information")
	}

	logger.File = filename
	logger.Line = strconv.Itoa(line)
	logger.Timestamp = time.Now().Format(time.RFC3339)
	logger.Message = strings.ToLower(message)

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

	output = fmt.Sprintf(messageFormat, logger.File, logger.Line, logger.Timestamp, logger.Message)

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

	logger.File = filename
	logger.Line = strconv.Itoa(line)
	logger.Timestamp = time.Now().Format(time.RFC3339)
	logger.Fault = fmt.Sprint(fault)

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
// (Optional) the output can be saved to a log file as JSON format.
func Complete(status, message string, fault error, save bool) {
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))

	_, filename, line, ok := runtime.Caller(1)

	if !ok {
		log.Fatal("unable to recover information")
	}

	logger.File = filename
	logger.Line = strconv.Itoa(line)
	logger.Timestamp = time.Now().Format(time.RFC3339)
	logger.Status = strings.ToLower(status)
	logger.Message = strings.ToLower(message)
	logger.Fault = fmt.Sprint(fault)

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
