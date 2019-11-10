package lh

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOHEntry_put(t *testing.T) {
	cases := []struct {
		e                *entry
		k                Hashable
		v                interface{}
		expectedEntry    *entry
		expectedAppended bool
	}{
		{
			nil,
			IntKey(1),
			"v1",
			&entry{IntKey(1), "v1", nil},
			true,
		},
		{
			&entry{IntKey(1), "v1", nil},
			IntKey(1),
			"v2",
			&entry{IntKey(1), "v2", nil},
			false,
		},
		{
			&entry{IntKey(1), "v1", nil},
			IntKey(2),
			"v2",
			&entry{
				IntKey(1),
				"v1",
				&entry{IntKey(2), "v2", nil},
			},
			true,
		},
	}

	for _, c := range cases {
		oe, appended := c.e.put(c.k, c.v)
		assert.Equal(t, c.expectedEntry, oe)
		assert.Equal(t, c.expectedAppended, appended)
	}
}

func TestOHEntry_get(t *testing.T) {
	cases := []struct {
		e          *entry
		k          Hashable
		expectedV  interface{}
		expectedOk bool
	}{
		{nil, IntKey(1), nil, false},
		{&entry{IntKey(1), "v1", nil}, IntKey(1), "v1", true},
		{&entry{IntKey(1), "v1", nil}, IntKey(2), nil, false},
		{&entry{IntKey(1), "v1", &entry{IntKey(2), "v2", nil}}, IntKey(2), "v2", true},
	}

	for _, c := range cases {
		v, ok := c.e.get(c.k)
		assert.Equal(t, c.expectedV, v)
		assert.Equal(t, c.expectedOk, ok)
	}
}

func TestOHEntry_remove(t *testing.T) {
	cases := []struct {
		e             *entry
		k             Hashable
		expectedEntry *entry
		expectedError error
	}{
		{nil, IntKey(1), nil, ErrEntryNotFound},
		{&entry{IntKey(1), "v1", nil}, IntKey(1), nil, nil},
		{
			&entry{
				IntKey(1),
				"v1",
				&entry{IntKey(2), "v2", nil},
			},
			IntKey(2),
			&entry{IntKey(1), "v1", nil},
			nil,
		},
	}

	for _, c := range cases {
		oe, err := c.e.remove(c.k)
		assert.Equal(t, c.expectedEntry, oe)
		assert.Equal(t, c.expectedError, err)
	}
}

func TestOHT_expand(t *testing.T) {
	var oht = NewOHT(2)

	oht.Put(1, "v1")
	assert.Equal(t, 2, len(oht.table))
	oht.Put(2, "v2")
	assert.Equal(t, 4, len(oht.table))
	oht.Put(3, "v3")
	oht.Put(4, "v4")
	oht.Put(5, "v5")
	assert.Equal(t, 8, len(oht.table))
}
