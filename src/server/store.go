package main

import (
	"errors"
	"fmt"
	"sync"
)

// storage for uploaded files
// make an interface so it can be swapped out for other storage later
// TODO make file storate that is persistent

type Storage interface {
	Load(name string) (data []byte, err error)
	Save(name string, data []byte) (err error)
	List() (items []string)
}

// basic in memory storage

type MemStore struct {
	name  string
	items map[string][]byte
	lock  sync.RWMutex
}

func NewMemStore(name string) (ms *MemStore) {
	ms = &MemStore{
		name:  name,
		items: make(map[string][]byte),
	}
	return
}

func (ms *MemStore) Load(name string) (data []byte, err error) {
	ms.lock.RLock()
	defer ms.lock.RUnlock()
	data, ok := ms.items[name]
	if !ok {
		return nil, errors.New("key not found")
	}
	return data, nil
}

func (ms *MemStore) Save(name string, data []byte) (err error) {
	fmt.Println(ms.items)
	fmt.Println("save --> ", name)
	ms.lock.Lock()
	ms.items[name] = data
	ms.lock.Unlock()
	return
}

func (ms *MemStore) List() (items []string) {
	keys := make([]string, len(ms.items))
	for i, _ := range ms.items {
		keys = append(keys, i)
	}
	return keys
}
