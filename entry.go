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

func (e *entry) len() (count uint) {
	if e == nil {
		return
	}
	for ee := e; ee != nil; ee = ee.next {
		count += 1
	}
	return
}
