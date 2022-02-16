# compare
Functions that help compare and output test results.

# How to use

```Compare(t *testing.T, actual interface{}, expected interface{})))```

# Sample
Refer the [test code of  this package itself](https://github.com/UedaTakeyuki/compare/blob/main/test/compare_test.go). 

The result of running above code is as follows:

```
MacBook-Air-Early-2015:test takeyuki$ go test compare_test.go -run 01 
ok      command-line-arguments  0.015s

MacBook-Air-Early-2015:test takeyuki$ go test compare_test.go -run 02
--- FAIL: Test_02 (0.00s)
    compare.go:13: got: kaeru, string
        want: 🐸, string
        called from command-line-arguments.Test_02 on /Volumes/devtmp/MyGo/compare/test/compare_test.go:14
FAIL
FAIL    command-line-arguments  0.015s
FAIL
```
