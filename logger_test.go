package golog

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"testing"
	"time"
)

// TestStart checks whether or not the log file has been created.
func TestStart(t *testing.T) {
	err := Start()
	if err != nil {
		t.Fatalf("\n❌ Start failed: %v\n", err)
	}

	dir, err := os.Getwd()
	if err != nil {
		t.Fatalf("\n❌ Unable to get the working directory: %v\n", err)
	}

	// Look for the log file.
	file, err := filepath.Glob(path.Join(dir, "golog.log"))
	if err != nil {
		t.Fatalf("\n❌ Unable to find the log file: %v\n", err)
	}

	filename := strings.Join(file, "")
	if filename != path.Join(dir, "golog.log") {
		t.Fatalf("\n❌ Unable to match the log file: %v\n", filename)
	}

	t.Log("\n✅ Start test passed...")
}

// TestFind checks whether or not the log file can be found.
func TestFind(t *testing.T) {
	filename, err := Find()
	if err != nil {
		t.Fatalf("\n❌ Find failed: %v\n", err)
	}

	dir, err := os.Getwd()
	if err != nil {
		t.Fatalf("\n❌ Unable to get working directory: %v\n", err)
	}

	if filename != path.Join(dir, "golog.log") {
		t.Fatalf("\n❌ Unable to match the log file: %v\n", err)
	}

	t.Log("\n✅ Find test passed...\n")
}

// TestSave checks whether or not the log file has content.
func TestSave(t *testing.T) {
	// Get information about the file and the line.
	_, filename, line, _ := runtime.Caller(1)

	logger := Logger{
		File:      filename,
		Line:      strconv.Itoa(line),
		Timestamp: time.Now().Format(time.RFC3339),
		Status:    "success",
		Message:   "golog is up and running",
		Fault:     fmt.Sprint(nil),
	}

	err := Save(logger)
	if err != nil {
		t.Fatalf("\n❌ Unable to save the log message to the log file: %v\n", err)
	}

	dir, err := os.Getwd()
	if err != nil {
		t.Fatalf("\n❌ Unable to get working directory: %v\n", err)
	}

	file, err := os.ReadFile(filepath.Clean(path.Join(dir, "golog.log")))
	if err != nil {
		t.Fatalf("\n❌ Unable to read the log file: %v\n", err)
	}

	var content map[string]interface{}
	_ = json.Unmarshal(file, &content)
	if err != nil {
		t.Fatalf("\n❌ Unable to unmarshal the log file content: %v\n", err)
	}

	if content["file"] != logger.File {
		t.Fatalf("\n❌ Unable to match the log file name: %v\n", err)
	}

	t.Log("\n✅ Save test passed...\n")
}

// TestLogger returns the output of Complete logger method
func TestLogger(t *testing.T) {
	// Complete log type
	Complete("success", "golog is up and running", nil, false)

	dir, err := os.Getwd()
	if err != nil {
		t.Fatalf("\n❌ Unable to get working directory: %v\n", err)
	}

	// Remove previously created log file
	err = os.Remove(path.Join(dir, "golog.log"))
	if err != nil {
		t.Fatalf("\n❌ Unable to remove the log file: %v\n", err)
	}

	t.Log("\n✅ TestLogger test passed...\n")
}

// BenchmarkSimple tests the performance of the Simple method.
func BenchmarkSimple(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Simple(false)
	}
}

// BenchmarkComplete tests the performance of the Complete method.
func BenchmarkComplete(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Complete("info", "benchmarking complete method", nil, false)
	}
}

// BenchmarkFault tests the performance of the Fault method.
func BenchmarkFault(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fault("warning", errors.New("benchmarking fault method"), false)
	}
}
