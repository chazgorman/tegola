package list

import (
	"fmt"
	"testing"

	"github.com/terranodo/tegola/maths"
)

func checkListLen(t *testing.T, desc string, l *List, len int) bool {
	if n := l.Len(); n != len {
		t.Errorf("%v: l.Len() = %d, want %d: %#v", desc, n, len, l)

		return false
	}
	return true
}

func checkListPointers(t *testing.T, desc string, l *List, es []*Pt) {

	if !checkListLen(t, desc, l, len(es)) {
		return
	}
	if len(es) == 0 {
		return
	}
	var didError bool

	// check internal and external prev/ext connection
	for i, e := range es {
		var Next, Prev *Pt
		if i > 0 {
			Prev = es[i-1]
		}

		if p := e.Prev(); (p != nil || Prev != nil) && p != Prev {
			t.Errorf("%s: elt[%d](%p).Prev() = %p, want %p", desc, i, e, p, Prev)
			didError = true
		}
		if i < len(es)-1 {
			Next = es[i+1]
		}
		if n := e.Next(); (n != nil || Next != nil) && n != Next {
			t.Errorf("%s: elt[%d](%p).Next() = %p, want %p", desc, i, e, n, Next)
			didError = true
		}
	}
	if didError {
		t.Errorf("list:%#v", l)
	}
}

func checkListInBetween(t *testing.T, desc string, i maths.Pt, loc int, mpts ...maths.Pt) {

	l := New()
	pts := NewPointSlice(mpts...)
	insert := NewPoint(i.X, i.Y)
	offset := 1
	if loc < 0 {
		offset = 0
	}

	cpts := make([]*Pt, len(pts)+offset)
	for i, p := range pts {
		l.PushBack(p)
		switch {
		case i < loc:
			cpts[i] = p
		case i == loc:
			cpts[i] = insert
			fallthrough
		case i > loc:
			cpts[i+offset] = p
		}

	}
	checkListPointers(t, fmt.Sprintf("list check: %v", desc), l, pts)
	l.PushInBetween(pts[0], pts[len(pts)-1], insert)
	checkListPointers(t, desc, l, cpts)
}

func TestPushInBetween(t *testing.T) {
	checkListInBetween(t, "Simple two point(3,1), after 1,1.", maths.Pt{3, 1}, 1, maths.Pt{1, 1}, maths.Pt{4, 1})
	checkListInBetween(t, "Simple three point(3,1), after 2,1.", maths.Pt{3, 1}, 2, maths.Pt{1, 1}, maths.Pt{2, 1}, maths.Pt{4, 1})
	checkListInBetween(t, "Simple three point(-1,1), Not included.", maths.Pt{-1, 1}, -1, maths.Pt{1, 1}, maths.Pt{2, 1}, maths.Pt{4, 1})
	checkListInBetween(t, "Simple three point(5,1), Not included.", maths.Pt{4, 1}, 2, maths.Pt{1, 1}, maths.Pt{2, 1}, maths.Pt{4, 1})
	checkListInBetween(t, "Dup three point(1,1), after 1,1.", maths.Pt{1, 1}, 1, maths.Pt{1, 1}, maths.Pt{2, 1}, maths.Pt{4, 1})
	checkListInBetween(t, "Dup three point(2,1), after 2,1.", maths.Pt{2, 1}, 1, maths.Pt{1, 1}, maths.Pt{2, 1}, maths.Pt{4, 1})
	checkListInBetween(t, "Dup three point(4,1), after 2,1.", maths.Pt{4, 1}, 2, maths.Pt{1, 1}, maths.Pt{2, 1}, maths.Pt{4, 1})
}
