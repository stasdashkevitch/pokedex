package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const intrval = time.Second * 5

	cases := []struct {
		key   string
		value []byte
	}{
		{
			key:   "https://example.com",
			value: []byte("testdata"),
		},
		{
			key:   "https://example.com/path",
			value: []byte("moretestdata"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %d", i), func(t *testing.T) {
			cache := NewCache(intrval)
			cache.Add(c.key, c.value)

			value, ok := cache.Get(c.key)
			if !ok {
				t.Error("Expected to find key")
				return
			}

			if string(value) != string(c.value) {
				t.Errorf("Expected to find value")
				return
			}
		})
	}
}

func TestReadLoop(t *testing.T) {
	const baseTime = time.Millisecond * 5
	const waitTime = time.Millisecond * baseTime
	cache := NewCache(baseTime)
	cache.Add("some url", []byte("testdata"))

	if _, ok := cache.Get("some url"); !ok {
		fmt.Errorf("Expected to find key")
	}

	time.Sleep(waitTime)

	if _, ok := cache.Get("some url"); ok {
		fmt.Errorf("Expected to find not key")
	}
}
