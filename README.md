# golang simple benchmark

```
$ go version
go version go1.11 darwin/amd64

$ go test -bench .
goos: darwin
goarch: amd64
pkg: github.com/tkhk-gh/golang-replace-bench
BenchmarkReplaceWithStrings-4           10000000               163 ns/op
BenchmarkReplaceWithRegex-4               500000              3429 ns/op
PASS
ok      github.com/tkhk-gh/golang-replace-bench 3.572s
```
