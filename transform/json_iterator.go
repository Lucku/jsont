package transform

import (
	"fmt"

	json "github.com/tidwall/gjson"
)

type JSONIterator interface {
	Next() bool
	Value() *PathElem
}

type PathElem struct {
	curMap []*kvPair
	curArr []json.Result
	index  int
	parent *PathElem
	Path   []string
	Value  *json.Result
}

type jsonIterator struct {
	Data       *json.Result
	OnlyLeaves bool
	current    *PathElem
	done       bool
}

type kvPair struct {
	k string
	v json.Result
}

func (j *jsonIterator) Next() bool {

	if j.done {
		return false
	}

	if j.current == nil {

		if j.Data.Type != json.JSON {
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

	if j.current.Value.Type == json.JSON {

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
func mapToKVPairs(in map[string]json.Result) []*kvPair {

	out := make([]*kvPair, 0, len(in))

	for k, v := range in {
		out = append(out, &kvPair{k, v})
	}

	return out
}

func (j *jsonIterator) Value() *PathElem {
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
