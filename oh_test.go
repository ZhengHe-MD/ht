package lh

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"testing"
)

func TestOHT_basic_usage(t *testing.T) {
	var oht = NewOHT(16)
	oht.Put("a", "v1")
	v, ok := oht.Get("a")
	assert.True(t, ok)
	assert.Equal(t, "v1", v)
	v, ok = oht.Get("b")
	assert.False(t, ok)
	assert.Equal(t, nil, v)
	oht.Put("b", "v2")
	v, ok = oht.Get("b")
	assert.True(t, ok)
	assert.Equal(t, "v2", v)
	err := oht.Remove("a")
	assert.NoError(t, err)
	v, ok = oht.Get("a")
	assert.False(t, ok)
	assert.Equal(t, nil, v)
}

func TestOHT_mobydick(t *testing.T) {
	pwd, err := os.Getwd()
	assert.NoError(t, err)
	f, err := os.Open(path.Join(pwd, "resource/mobydick.txt"))
	assert.NoError(t, err)

	bs, err := ioutil.ReadAll(f)
	assert.NoError(t, err)

	words := strings.Fields(string(bs))

	var naive = make(map[string]interface{})
	var oht = NewOHT(16)
	for _, word := range words {
		oht.Put(word, struct{}{})
		naive[word] = struct{}{}
	}

	assert.True(t, oht.size > 0)
	assert.True(t, len(oht.table) > 0)
	assert.True(t, oht.size*2 <= len(oht.table))
	assert.Equal(t, len(naive), oht.size)
}
