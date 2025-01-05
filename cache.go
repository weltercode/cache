package cache

import (
	"errors"
	"fmt"
)

type Cache struct {
	items map[string]any
}

func New() Cache {
	return Cache{
		items: make(map[string]any),
	}
}

func (c Cache) Set(key string, value any) Cache {
	c.items[key] = value
	return c
}

func (c Cache) Get(key string) (any, error) {
	_, ok := c.items[key]
	if ok {
		return c.items[key], nil
	}
	message := fmt.Sprintf("key %s not found", key)
	return nil, errors.New(message)
}

func (c Cache) Delete(key string) error {
	_, ok := c.items[key]
	if ok {
		delete(c.items, key)
		return nil
	}
	message := fmt.Sprintf("key %s not found, can`t delete", key)
	return errors.New(message)
}
