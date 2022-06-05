package store

import "sort"

type Unit struct {
	point       string
	disFromLast int
	disFromOri  int
	last        string
}

func (u *Unit) setDis(n int) {
	u.disFromOri = n
}

func (u *Unit) setLastPorint(s string) {
	u.last = s
}

type PriorityQ struct {
	que []Unit
}

func (pq *PriorityQ) append(u Unit) {
	pq.que = append(pq.que, u)
	sort.SliceStable(pq.que, func(i, j int) bool {
		return pq.que[i].disFromOri <= pq.que[j].disFromOri
	})
}

func (pq *PriorityQ) pop() Unit {
	if len(pq.que) < 1 {
		return Unit{}
	}
	unit := pq.que[0]
	pq.que = pq.que[1:]
	return unit
}

type Bfss struct {
	pq    PriorityQ
	res   []Unit
	exist map[string]bool
}

func (b *Bfss) pop() {
	u := b.pq.pop()
	if !b.exist[u.point] {
		b.res = append(b.res, u)
		b.exist[u.point] = true
	}
}

func (b *Bfss) append(u Unit, lastUnit Unit) {
	u.last = lastUnit.point
	u.disFromOri = lastUnit.disFromOri + u.disFromLast
	b.pq.append(u)
}
