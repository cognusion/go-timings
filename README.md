
# go-timings

[![GoDoc](https://godoc.org/github.com/cognusion/go-timings?status.svg)](https://godoc.org/github.com/cognusion/go-timings)

Several helpers for timing functions or sections of functions. e.g.

```go
func MyFunc() {
  defer timings.Track("MyFunc", timings.Now(), timings.Stderr)
  // Do stuff here
  
}
```
