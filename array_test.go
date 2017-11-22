package containers

import (
	"testing"
)

func Test_Array(t *testing.T) {
	arr := NewArray(10)
	t.Log("test")
	t.Logf("Array Length, %d\n", arr.GetLength())
}

func Test_Array_push(t *testing.T) {
	arr := NewArray(3)
	t.Log("test")
	t.Logf("Array Length, %d\n", arr.GetLength())
	arr.Push(1)
	arr.Push(2)
	arr.Push(3)
	t.Logf("Array Length, %d\n", arr.GetLength())
	arr.Push(4)
	a := arr.Pop()
	if a.(int) != 3 {
		t.FailNow()
	}
	t.Logf("pop : %d", a.(int)) // print 3
}
