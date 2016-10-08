package gotest

import "testing"
import "fmt"

func Test_Division(t *testing.T) {
	var a, b float64
	a, b = 1.0, 0
	if _, e := Division(a, b); e != nil {
		t.Error(fmt.Sprintf("除法测试%f, %f没通过", a, b))
	}
	a, b = 1.0, 1.0
	if _, e := Division(a, b); e != nil {
		t.Error(fmt.Sprintf("除法测试%f, %f没通过", a, b))
	} else {
		t.Log(fmt.Sprintf("除法测试%f, %f通过", a, b))
	}
}
