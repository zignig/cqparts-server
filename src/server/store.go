package main

import (
	"encoding/gob"
	"errors"
	"fmt"
	"os"
	"sync"
	"time"
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
	Name  string
	Items map[string][]byte
	lock  sync.RWMutex
	dirty bool
}

func NewMemStore(name string) (ms *MemStore) {
	ms = &MemStore{
		Name:  name,
		Items: make(map[string][]byte),
	}
	err := ms.Restore()
	if err != nil {
		fmt.Println(err)
	}
	ms.Periodic()
	return
}

func (ms *MemStore) Periodic() {
	ticker := time.NewTicker(5 * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				if ms.dirty {
					ms.dirty = false
					fmt.Println("saved down " + ms.Name)
					ms.Dump()
				}
			}
		}
	}()
}

func (ms *MemStore) Load(name string) (data []byte, err error) {
	ms.lock.RLock()
	defer ms.lock.RUnlock()
	data, ok := ms.Items[name]
	if !ok {
		return nil, errors.New("key not found")
	}
	return data, nil
}

func (ms *MemStore) Save(name string, data []byte) (err error) {
	fmt.Println("save --> ", name)
	ms.lock.Lock()
	ms.dirty = true
	ms.Items[name] = data
	ms.lock.Unlock()
	return
}

func (ms *MemStore) Restore() (err error) {
	file, err := os.Open(ms.Name + ".dmp")
	if err == nil {
		ms.lock.Lock()
		decoder := gob.NewDecoder(file)
		err = decoder.Decode(ms)
		ms.lock.Unlock()
	}
	file.Close()
	return err
}
func (ms *MemStore) Dump() (err error) {
	file, err := os.Create(ms.Name + ".dmp")
	if err == nil {
		ms.lock.Lock()
		encoder := gob.NewEncoder(file)
		encoder.Encode(ms)
		ms.lock.Unlock()
	}
	file.Close()
	return err
}

func (ms *MemStore) List() (items []string) {
	keys := make([]string, len(ms.Items))
	for i, _ := range ms.Items {
		keys = append(keys, i)
	}
	return keys
}
