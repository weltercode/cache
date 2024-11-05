package cache

import (
	"testing"
)

func TestNew(t *testing.T) {
	// Create a new cache instance using the `New()` function
	c := New()

	// Check if the returned type is `Cache`
	if _, ok := interface{}(c).(Cache); !ok {
		t.Errorf("Result was incorrect, got type %T, want type Cache", c)
	}

	// Ensure the items map is initialized
	if c.items == nil {
		t.Errorf("Cache items map was not initialized")
	}
}

func TestSet(t *testing.T) {
	c := New()

	key := "test_key_zz"
	c.Set(key, 10)
	d, err := c.Get(key)

	if err != nil {
		t.Errorf("Err was returned `%s`", err.Error())
	}

	if d != 10 {
		t.Errorf("Returned %s, but expected %d", d.(string), 10)
	}
}
