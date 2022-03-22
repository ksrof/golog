# Golog
Logger library for `Golang` which works on top of the standard library logger. It gives a nicely color-coded output (when used alongisde a `TTY`) otherwise it can write the contents to a file in `JSON` format.

![Golog output](https://github.com/ksrof/golog/blob/main/images/golog_output.png)
## Roadmap
- [x] New logger methods
  - [x] Fault log
  - [x] Complete log
- [x] (Optional) Save log to file
- [x] Colored output
- [x] Implement tests
- [x] Documentation
- [x] CI workflow

The following ideas should be included in further releases, but the intentions are to have all this features added in the first stable version of the **golog logger library**.

**Important note** This project is still in its early stages so expect breaking changes in the API, experiment, play and contribute but don't use it in production code or something that is serious enough.

### Golog Version 1.0.0
- [ ] Logger method with custom fields
- [ ] Plain text and JSON output
- [ ] Save to specific log file
- [ ] Multi level logging
- [ ] Environment option
- [ ] Multiple terminal outputs
- [ ] Logger assertion on tests
- [ ] Fatal and Panic graceful handling

## Examples
The simplest way to use Golog is the Simple method:
```go
package main

import (
  log "github.com/ksrof/golog"
)

func main() {
  // Simple returns File, Line and Timestamp
  // it also takes a boolean parameter to determine
  // whether or not it should save the output to a file.
  log.Simple(false)
}
```
As you might be seeing it is fully compatible with the standard library logger meaning that you can replace all your `log` imports with `log "github.com/ksrof/golog"` and profit from the cool features that it has!

## License
The MIT License (MIT) - see [`license`](https://github.com/ksrof/golog/blob/main/LICENSE) for more details.
