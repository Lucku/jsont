package transform

import (
	"testing"
)

func TestStackPush(t *testing.T) {

	s := newTypeSafeStack(5)

	s.Push(24)
	s.Push(25)
	s.Push(12)
	s.Push(1)

	t.Logf("%v\n", s)

	t.Logf("%v\n", s.Pop())
	t.Logf("%v\n", s.Pop())
	t.Logf("%v\n", s.Pop())
	t.Logf("%v\n", s.Pop())
}
