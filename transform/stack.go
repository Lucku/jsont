package transform

import (
	"fmt"
	"reflect"
	"strings"
)

// Stack is the interface for a LIFO data structure that can hold values of arbitrary type
type Stack interface {
	Push(elem interface{})
	Pop() interface{}
	Peek() interface{}
	Size() int
}

type typeSafeStack struct {
	contents []interface{}
	pos      int
	t        reflect.Type
}

func newTypeSafeStack(size int) *typeSafeStack {

	s := make([]interface{}, 0, size)

	return &typeSafeStack{contents: s, pos: -1}
}

func (s *typeSafeStack) Push(elem interface{}) {

	t := reflect.TypeOf(elem)

	if s.t != nil && t.PkgPath()+"/"+t.Name() != s.t.PkgPath()+"/"+s.t.Name() {
		panic("pushing element of inconsistent type to stack")
	}

	s.contents = append(s.contents, elem)

	s.pos++

	if s.pos == 0 {
		s.t = reflect.TypeOf(s.contents[0])
	}
}

func (s *typeSafeStack) Pop() interface{} {

	if s.pos < 0 {
		return nil
	}

	res := s.contents[s.pos]

	s.contents = s.contents[:s.pos]

	s.pos--

	return res
}

func (s *typeSafeStack) Peek() interface{} {

	if s.pos < 0 {
		return nil
	}

	return s.contents[s.pos]
}

func (s typeSafeStack) Size() int {
	return len(s.contents)
}

func (s typeSafeStack) String() string {

	sb := strings.Builder{}

	sb.WriteString(fmt.Sprintf("Stack (%d) [", s.Size()))

	for i, elem := range s.contents {

		if i == len(s.contents)-1 {
			sb.WriteString(fmt.Sprintf("%v]", elem))
		} else {
			sb.WriteString(fmt.Sprintf("%v ", elem))
		}
	}

	return sb.String()
}
