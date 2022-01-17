# Golog
A tiny logger that displays useful information about how things went wrong and at which point.
[Read the tutorial at Medium](https://medium.com/@ksrof/detailed-logging-with-golang-fa0074344f5c)

## Roadmap
- [ ] Add new logger methods
- [ ] Add color to output
- [ ] Improve project structure
- [ ] Improve code quality
- [ ] Implement unit tests
- [ ] Add logger documentation
- [ ] Setup CI

## Examples
* Simple Logger: displays a log contaning the File, Line and Timestamp of where it has been invocated.
	```go
	golog.Simple() // fires log.Print
	/*
	* Output:
	* File: path/to/file.go
	* Line: 45
	* Timestamp: 2022-01-13T16:38:46+01:00
	*/
	```
* Status Logger: displays a log and uses a different log method depending on the status given by the user.
	```go
	golog.Status("fatal") // fires log.Fatal
	/*
	* Output:
	* File: path/to/file.go
	* Line:	45
	* Timestamp: 2022-01-13T16:38:46+01:00
	* Status: fatal
	*/
	```
* Message Logger: displays a log containing the message provided by the user.
	```go
	golog.Message("beep beep boop") // fires log.Print
	/*
	* Output:
	* File: path/to/file.go
	* Line: 45
	* Timestamp: 2022-01-13T16:38:46+01:00
	* Message: beep beep boop
	*/
	```
* Error Logger: displays a log containing the error provided by the user.
	```go
	golog.Error("panic") // fires log.Panic
	/*
	* Output:
	* File: path/to/file.go
	* Line: 45
	* Timestamp: 2022-01-13T16:38:46+01:00
	* Error: invalid memory address or nil pointer dereference
	*/
	```
* Complete Logger: displays a log containing all the default and optional parameters.
	```go
	golog.Complete("success", "the logger is up and running", nil) // fires log.Print
	/*
	* Output:
	* File: path/to/file.go
	* Line: 45
	* Timestamp: 2022-01-13T16:38:46+01:00
	* Status: success
	* Message: the logger is up and running
	* Error: <nil>
	*/
	```