package main

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	bolt "github.com/coreos/bbolt"
	"os"
	"strings"
	"sync"
	"time"
)

// Key value store(s) for data

// interface for golang awesomeness
type Storage interface {
	Load(name string) (data []byte, err error)
	Save(name string, data []byte) (err error)
	List() (items []string)
	Multi(name string) (files map[string][]byte, err error)
}

// a named kv store for bbolt
type databucket struct {
	db   *bolt.DB
	name string
}

// bbolt store
type BBoltStore struct {
	db      *bolt.DB
	names   *databucket
	buckets map[string]databucket
}

func NewBBoltStore(name string) (bb *BBoltStore) {
	bb = &BBoltStore{}
	bb.db, _ = bolt.Open(name+".db", 0600, nil)
	bb.buckets = make(map[string]databucket)
	bb.names = bb.NewBucket("names")
	return
}

func (bb BBoltStore) NewBucket(name string) (buk *databucket) {
	buk = &databucket{
		db:   bb.db,
		name: name,
	}
	// check it's not the name bucket
	if strings.Compare(name, "names") != 0 {
		bb.names.Save(name, []byte("_"))
	}
	bb.db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte(name))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
	return buk
}

// Make Storage interface for the databucket
func (bk databucket) Load(name string) (data []byte, err error) {
	err = bk.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bk.name))
		data = b.Get([]byte(name))
		if len(data) == 0 {
			return errors.New("key " + name + " does not exist")
		}
		return err
	})
	return data, err
}

func (bk databucket) Save(name string, data []byte) (err error) {
	err = bk.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bk.name))
		err := b.Put([]byte(name), data)
		return err
	})
	return
}

func (bk databucket) List() (items []string) {
	items = make([]string, 0)
	bk.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bk.name))
		c := b.Cursor()
		for k, _ := c.First(); k != nil; k, _ = c.Next() {
			items = append(items, string(k))
		}
		return nil
	})
	return items
}

func (bk databucket) Multi(name string) (files map[string][]byte, err error) {
	files = make(map[string][]byte)
	err = bk.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bk.name))
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
