# Golog
Golog lets you display errors with different types of characteristics and give context when things go wrong.

## Roadmap
- [x] New logger methods
  - [x] Status log
  - [x] Message log
  - [x] Fault log
  - [x] Complete log
- [x] (Optional) Save log to file
- [ ] Colored output
- [ ] Implement tests
- [ ] Documentation
- [ ] Setup CI

## Examples
**Simple Logger:** displays a log contaning the File, Line and Timestamp of where it has been invocated.
```go
golog.Simple() // fires log.Print
/*
* Output:
* File: path/to/file.go
* Line: 45
* Timestamp: 2022-01-13T16:38:46+01:00
*/
```

**Status Logger:** displays a log and uses a different log method depending on the status given by the user.
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

**Message Logger:** displays a log containing the message provided by the user.
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

**Fault Logger:** displays a log with an error message and uses a different log method depending on the class given by the user.
```go
golog.Fault("panic", err) // fires log.Panic
/*
* Output:
* File: path/to/file.go
* Line: 45
* Timestamp: 2022-01-13T16:38:46+01:00
* Fault: invalid memory address or nil pointer dereference
*/
```

**Complete Logger:** displays a log containing all the default and optional parameters.
```go
golog.Complete("success", "the logger is up and running", nil) // fires log.Print
/*
* Output:
* File: path/to/file.go
* Line: 45
* Timestamp: 2022-01-13T16:38:46+01:00
* Status: success
* Message: the logger is up and running
* Fault: <nil>
*/
```

## Credits
* [Kevin Suñer](https://github.com/ksrof)

## License
The MIT License (MIT) - see [`license`](https://github.com/ksrof/golog/blob/main/LICENSE) for more details.