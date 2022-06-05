package store

import "log"

var example map[string][]string = map[string][]string{
	"a": []string{"b", "c"},
	"b": []string{"a", "c", "d"},
	"c": []string{"a", "b", "c", "d"},
	"d": []string{"b", "c", "e", "f"},
	"e": []string{"c", "d"},
	"f": []string{"d"},
}

type Stack struct {
	arr []string
}

func (s *Stack) length() int {
	return len(s.arr)
}

func (s *Stack) Append(v string) {
	s.arr = append(s.arr, v)
}

func (s *Stack) Pop() string {
	tail := len(s.arr) - 1
	if tail < 0 {
		return ""
	}
	popString := s.arr[tail]
	s.arr = s.arr[:tail]
	return popString
}

type DFS struct {
	stack Stack
	exist map[string]bool
	res   []string
}

func (b *DFS) toExist(s string) {
	b.exist[s] = true
}
func (b *DFS) checkIfExist(s string) bool {
	return b.exist[s]
}

func (b *DFS) toRes(s string) {
	b.res = append(b.res, s)
}

func (b *DFS) Process(a map[string][]string, start string) {
	b.stack.Append(start)
	b.toExist(start)
	for b.stack.length() > 0 {
		popString := b.stack.Pop()
		b.toRes(popString)

		newCome := a[popString]

		for _, s := range newCome {
			if b.checkIfExist(s) {
				continue
			}
			b.toExist(s)
			b.stack.Append(s)
		}
	}

	log.Println(b.res)
}
