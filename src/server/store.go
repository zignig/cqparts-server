package main

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"os"
	"sync"
	"time"

	bolt "github.com/coreos/bbolt"
)

// storage for uploaded files
// make an interface so it can be swapped out for other storage later
// TODO make file storate that is persistent

type Storage interface {
	Load(name string) (data []byte, err error)
	Save(name string, data []byte) (err error)
	List() (items []string)
	Multi(name string) (files map[string][]byte, err error)
}

// bbolt store
type BBoltStore struct {
	db *bolt.DB
}

func NewBBoltStore(name string) (bb *BBoltStore) {
	bb = &BBoltStore{}
	bb.db, _ = bolt.Open(name+".db", 0600, nil)
	bb.db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte("models"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
	return
}

func (bb BBoltStore) Load(name string) (data []byte, err error) {
	bb.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("models"))
		data = b.Get([]byte(name))
		return err
	})
	return data, err
}

func (bb BBoltStore) Save(name string, data []byte) (err error) {
	err = bb.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("models"))
		err := b.Put([]byte(name), data)
		return err
	})
	return
}

func (bb BBoltStore) List() (items []string) {
	return
}

func (bb BBoltStore) Multi(name string) (files map[string][]byte, err error) {
	files = make(map[string][]byte)
	err = bb.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("models"))
		c := b.Cursor()
		prefix := []byte("/" + name)
		for k, v := c.Seek(prefix); k != nil && bytes.HasPrefix(k, prefix); k, v = c.Next() {
			fmt.Println(string(k))
			files[string(k)] = v
		}
		return nil
	})
	return
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

func (ms *MemStore) Multi(name string) (files map[string][]byte, err error) {
	return
}

func (ms *MemStore) List() (items []string) {
	keys := make([]string, len(ms.Items))
	for i, _ := range ms.Items {
		keys = append(keys, i)
	}
	return keys
}
