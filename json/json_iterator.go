package json

import (
	"fmt"

	"github.com/tidwall/gjson"
)

// Iterator allows to traverse a tree data structure
type Iterator interface {
	Next() bool
	Value() *PathElem
}

// PathElem is an element inside a JSON tree, identified by a JSON value and a path inside the
// overall JSON document
type PathElem struct {
	curMap []*kvPair
	curArr []gjson.Result
	index  int
	parent *PathElem
	Path   []string
	Value  *gjson.Result
}

type iterator struct {
	Data       *gjson.Result
	OnlyLeaves bool
	current    *PathElem
	done       bool
}

type kvPair struct {
	k string
	v gjson.Result
}

// NewIterator initializes an iterator to a JSON document taken as input
func NewIterator(data *gjson.Result) Iterator {
	return &iterator{Data: data}
}

func (j *iterator) Next() bool {

	if j.done {
		return false
	}

	if j.current == nil {

		if j.Data.Type != gjson.JSON {
			return false
		}

		j.current = &PathElem{
			Value: j.Data,
			Path:  []string{},
			index: 0,
		}

		if j.Data.IsObject() {
			j.current.curMap = mapToKVPairs(j.current.Value.Map())
		} else if j.Data.IsArray() {
			j.current.curArr = j.current.Value.Array()
		}

		return true
	}

	var nextPathElem *PathElem

	if j.current.Value.Type == gjson.JSON {

		if j.current.Value.IsObject() {
			j.current.curMap = mapToKVPairs(j.current.Value.Map())
		} else if j.current.Value.IsArray() {
			j.current.curArr = j.current.Value.Array()
		}

		child := &PathElem{
			parent: j.current,
			index:  -1,
		}
		nextPathElem = child.nextPathElemFromParent()
	} else {
		nextPathElem = j.current.nextPathElemFromParent()
	}

	if nextPathElem == nil {
		j.done = true
		return false
	}

	j.current = nextPathElem

	return true
}

// This might be a little bit hacky, but needs to be done in order for the map to be ordered
func mapToKVPairs(in map[string]gjson.Result) []*kvPair {

	out := make([]*kvPair, 0, len(in))

	for k, v := range in {
		out = append(out, &kvPair{k, v})
	}

	return out
}

func (j *iterator) Value() *PathElem {
	return j.current
}

func (p *PathElem) nextPathElemFromParent() *PathElem {

	if p.parent == nil {
		return nil
	}

	nextPathElem := &PathElem{
		parent: p.parent,
		index:  p.index + 1,
	}

	if p.parent.curMap != nil {

		if nextPathElem.index == len(nextPathElem.parent.curMap) {
			// we are done on this level, go recursively upwards
			return p.parent.nextPathElemFromParent()
		}

		// go to next item in parent map
		nextItem := nextPathElem.parent.curMap[nextPathElem.index]
		nextPathElem.Value = &nextItem.v
		nextPathElem.Path = append(p.parent.Path, nextItem.k)

		return nextPathElem

	} else if p.parent.curArr != nil {

		if nextPathElem.index == len(p.parent.curArr) {
			// we are done on this level, go recursively upwards
			return p.parent.nextPathElemFromParent()
		}
		// go to next item in parent array
		nextItem := nextPathElem.parent.curArr[nextPathElem.index]
		nextPathElem.Value = &nextItem
		nextPathElem.Path = append(p.parent.Path, fmt.Sprintf("[%d]", nextPathElem.index))

		return nextPathElem
	}

	return nil
}
