package gotest

import (
	"testing"
)

func Test_Division_1(t *testing.T) {
	if i, e := Division(6, 2); i != 3 || e != nil {
		t.Error("Division case fail")
	} else {
		t.Log("The first case passed!")
	}
}

func Test_Division_2(t *testing.T) {
	t.Error("Case not pass.")
}

/*
➜  gotest git:(master) ✗ go test
--- FAIL: Test_Division_2 (0.00s)
        gotest_test.go:16: Case not pass.
FAIL
exit status 1
FAIL    go_json/gotest  0.002s

➜  gotest git:(master) ✗ go test -v
=== RUN   Test_Division_1
--- PASS: Test_Division_1 (0.00s)
        gotest_test.go:11: The first case passed!
=== RUN   Test_Division_2
--- FAIL: Test_Division_2 (0.00s)
        gotest_test.go:16: Case not pass.
FAIL
exit status 1
FAIL    go_json/gotest  0.004s
 */
