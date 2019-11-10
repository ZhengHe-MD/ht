package lh

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"testing"
)

func TestLHT_basic_usage(t *testing.T) {
	var lht = NewLHT()
	lht.Put("a", "v1")
	v, ok := lht.Get("a")
	assert.True(t, ok)
	assert.Equal(t, "v1", v)
	v, ok = lht.Get("b")
	assert.False(t, ok)
	assert.Equal(t, nil, v)
	lht.Put("b", "v2")
	v, ok = lht.Get("b")
	assert.True(t, ok)
	assert.Equal(t, "v2", v)
	err := lht.Remove("a")
	assert.NoError(t, err)
	v, ok = lht.Get("a")
	assert.False(t, ok)
	assert.Equal(t, nil, v)
}

func TestLHT_mobydick(t *testing.T) {
	pwd, err := os.Getwd()
	assert.NoError(t, err)
	f, err := os.Open(path.Join(pwd, "resource/mobydick.txt"))
	assert.NoError(t, err)

	bs, err := ioutil.ReadAll(f)
	assert.NoError(t, err)

	words := strings.Fields(string(bs))

	var oht = NewLHT()
	var naive = make(map[string]interface{})
	for _, word := range words {
		oht.Put(word, struct{}{})
		naive[word] = struct{}{}
	}

	assert.True(t, oht.nentry > 0)
	assert.True(t, len(oht.table) > 0)
	assert.Equal(t, len(naive), int(oht.nentry))
}
