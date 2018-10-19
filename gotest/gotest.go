package gotest

import (
	"errors"
)

func Division(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Division 0 forbidden")
	}

	return a / b, nil
}

/*
Unit test:
➜  gotest git:(master) ✗ go test gotest.go gotest_test.go
--- FAIL: Test_Division_2 (0.00s)
        gotest_test.go:16: Case not pass.
FAIL
FAIL    command-line-arguments  0.003s

Unit test verbose mode:
➜  gotest git:(master) ✗ go test -v gotest.go gotest_test.go
=== RUN   Test_Division_1
--- PASS: Test_Division_1 (0.00s)
        gotest_test.go:11: The first case passed!
=== RUN   Test_Division_2
--- FAIL: Test_Division_2 (0.00s)
        gotest_test.go:16: Case not pass.
FAIL
FAIL    command-line-arguments  0.002s

Performance test:
➜  gotest git:(master) ✗ go test -test.bench=".*" gotest.go webbench_test.go
goos: linux
goarch: amd64
Benchmark_Division-4                    2000000000               0.39 ns/op
Benchmark_TimeConsumingFunction-4       2000000000               0.39 ns/op
PASS
ok      command-line-arguments  1.651s
 */

