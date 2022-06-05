package store

import "log"

type BFS struct {
	queue []string
	res   []string
	exist map[string]bool
}

func (b *BFS) toQueue(s string) {
	b.queue = append(b.queue, s)
}

func (b *BFS) popQueue() string {
	if len(b.queue) < 1 {
		return ""
	}
	res := b.queue[0]
	b.queue = b.queue[1:]
	return res
}

func (b *BFS) toExist(s string) {
	b.exist[s] = true
}
func (b *BFS) checkIfExist(s string) bool {
	return b.exist[s]
}

func (b *BFS) toRes(s string) {
	b.res = append(b.res, s)
}

func (b *BFS) Process(m map[string][]string, start string) {
	b.toQueue(start)
	for len(b.queue) > 0 {
		popS := b.popQueue()
		b.toRes(popS)
		b.toExist(popS)
		conn := m[popS]
		for _, s := range conn {
			if b.checkIfExist(s) {
				continue
			}
			b.toExist(s)
			b.toQueue(s)
		}
	}
	log.Println(b.res)

}
