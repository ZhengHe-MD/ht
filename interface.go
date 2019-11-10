package lh

import "math"

type Hashable interface {
	Hash() int
	Equals(b Hashable) bool
}

type HashTable interface {
	Size() int
	Get(key Hashable) (interface{}, bool)
	Put(key Hashable, value interface{})
	Remove(key Hashable) (err error)
}

type IntKey int

func (ik IntKey) Hash() int {
	return int(ik)
}

func (ik IntKey) Equals(b Hashable) bool {
	return ik.Hash() == b.Hash()
}

type StringKey string

// reference: http://www.cse.yorku.ca/~oz/hash.html
func (sk StringKey) Hash() int {
	var h = 5381

	// djb2
	for _, c := range sk {
		h = ((h<<5)+h) + int(c)
		h = h % math.MaxInt32
	}

	return h
}

func (sk StringKey) Equals(b Hashable) bool {
	return sk.Hash() == b.Hash()
}


