// open hashing
package lh

type entry struct {
	k    Hashable
	v    interface{}
	next *entry
}

func (e *entry) put(k Hashable, v interface{}) (oe *entry, appended bool) {
	if e == nil {
		return &entry{k, v, nil}, true
	}

	if e.k.Equals(k) {
		e.v = v
		return e, false
	} else {
		e.next, appended = e.next.put(k, v)
		return e, appended
	}
}

func (e *entry) get(k Hashable) (v interface{}, ok bool) {
	if e == nil {
		return nil, false
	}

	if e.k.Equals(k) {
		return e.v, true
	} else {
		return e.next.get(k)
	}
}

func (e *entry) remove(k Hashable) (*entry, error) {
	if e == nil {
		return nil, ErrEntryNotFound
	}

	if e.k.Equals(k) {
		return e.next, nil
	} else {
		var err error
		e.next, err = e.next.remove(k)
		return e, err
	}
}

type OHT struct {
	table []*entry
	size int
}

func NewOHT(initSize int) *OHT {
	return &OHT{
		table: make([]*entry, initSize),
		size:  0,
	}
}

func (h *OHT) Size() int {
	return h.size
}

func (h *OHT) bucket(hk Hashable) int {
	return hk.Hash() % len(h.table)
}

func (h *OHT) Get(k interface{}) (v interface{}, ok bool) {
	hk := toHashable(k)
	return h.table[h.bucket(hk)].get(hk)
}

func (h *OHT) Put(k interface{}, v interface{}) {
	hk := toHashable(k)
	h.put(hk, v)
}

func (h *OHT) put(hk Hashable, v interface{}) {
	bucket := h.bucket(hk)
	var appended bool
	if h.table[bucket], appended = h.table[bucket].put(hk, v); appended {
		h.size += 1
		// keep load <= 50%
		if h.size * 2 > len(h.table) {
			h.expand()
		}
	}
	return
}

func (h *OHT) Remove(k interface{}) (err error) {
	hk := toHashable(k)
	bucket := h.bucket(hk)
	h.table[bucket], err = h.table[bucket].remove(hk)
	if err == ErrEntryNotFound {
		h.size -= 1
	}
	return
}

func (h *OHT) expand() {
	prevTable := h.table
	h.table = make([]*entry, len(prevTable)*2)
	h.size = 0
	for _, e := range h.table {
		for ee := e; ee != nil; ee = ee.next {
			h.put(ee.k, ee.v)
		}
	}
}
