package emaps_test

import (
	"testing"

	"github.com/bylexus/go-stdlib/emaps"
)

type getMapKeysTestdata struct {
	inMap map[string]string
	keys  []string
}

func TestGetMapKeys(t *testing.T) {
	testdata := []getMapKeysTestdata{
		{
			inMap: map[string]string{},
			keys:  []string{},
		},
		{
			inMap: map[string]string{
				"one": "One",
				"two": "Two",
			},
			keys: []string{"one", "two"},
		},
	}

	for _, data := range testdata {
		keys := emaps.GetMapKeys(&data.inMap)
		if len(keys) != len(data.keys) {
			t.Errorf("nr of keys wrong in %#v\n", data.inMap)
		}
		for _, testKey := range data.keys {
			if _, ok := data.inMap[testKey]; ok != true {
				t.Errorf("Key '%s' not in %#v\n", testKey, data.inMap)
			}
		}
	}
}
