package lh

const (
	utilization     = 0.75
	recordsPerBlock = 8
)

type LHT struct {
	table []*entry

	nblock uint
	nbit   uint
	nentry uint
}

func NewLHT() *LHT {
	return &LHT{
		table:  make([]*entry, 2),
		nblock: 2,
		nbit:   1,
		nentry: 0,
	}
}

func (h *LHT) bucket(hk Hashable) uint {
	m := uint(hk.Hash() & ((1 << h.nbit) - 1))
	if m < h.nblock {
		return m
	} else {
		return m ^ (1 << (h.nbit - 1))
	}
}

func (h *LHT) Get(k interface{}) (v interface{}, ok bool) {
	hk := toHashable(k)
	return h.table[h.bucket(hk)].get(hk)
}

func (h *LHT) Put(k interface{}, v interface{}) {
	hk := toHashable(k)
	h.put(hk, v)
}

func (h *LHT) put(hk Hashable, v interface{}) {
	bucket := h.bucket(hk)
	var appended bool
	if h.table[bucket], appended = h.table[bucket].put(hk, v); appended {
		h.nentry += 1
		// utilization = 0.75
		if h.nentry > uint(utilization*float64(h.nblock)*recordsPerBlock) {
			h.split()
		}
	}
}

func (h *LHT) Remove(k interface{}) (err error) {
	hk := toHashable(k)
	bucket := h.bucket(hk)
	h.table[bucket], err = h.table[bucket].remove(hk)
	if err != nil {
		return
	}
	h.nentry -= 1
	return nil
}

func (h *LHT) split() {
	bucket := h.nblock % (1 << (h.nbit - 1))
	block := h.table[bucket]
	h.nblock += 1
	if h.nblock > (1 << h.nbit) {
		h.nbit += 1
	}

	var currBlock *entry
	var nextBlock *entry
	for ee := block; ee != nil; ee = ee.next {
		if h.bucket(ee.k) == bucket {
			currBlock, _ = currBlock.put(ee.k, ee.v)
		} else {
			nextBlock, _ = nextBlock.put(ee.k, ee.v)
		}
	}

	h.table[bucket] = currBlock
	h.table = append(h.table, nextBlock)

	return
}
