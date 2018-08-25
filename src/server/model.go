package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/satori/go.uuid"
	"time"
)

// the construct for a single model
// Flesh this out so it has all the info for a given model

// have segmented update and dirty write down
// get as much information as possible from the cqparts display post

// Todo , covert model to storage interface
// models are dereferenced via UUID

type Model struct {
	Name    string    `json:"name"`
	Img     string    `json:"img"`
	View    *Render   `json:"view"`
	Status  string    `json:"status"`
	UUID    uuid.UUID `json:"uuid"`
	Created time.Time `json:"created"`
	Pinned  bool      `json:"pinned"`
	Thumb bool `json:"thumb"`

}

type ModelCollection struct {
	data  Storage
	names Storage
}

func NewModel(name string) (m *Model) {
	m = &Model{
		Name:    name,
		Created: time.Now(),
		View:    NewRender("base"),
		Img:     " ",
		UUID:    uuid.Must(uuid.NewV4()),
	}
	return m
}

func NewModelCollection(name string, bs *BBoltStore) (mc *ModelCollection) {
	mc = &ModelCollection{
		data:  bs.NewBucket(name + "_model_data"),
		names: bs.NewBucket(name + "_model_names"),
	}
	return mc
}

// find a named model and create if it does not exist
func (mc *ModelCollection) Find(name string) (m *Model) {
	d, err := mc.names.Load(name)
	// create if missing
	if err != nil {
		m := NewModel(name)
		mc.names.Save(name, []byte(m.UUID.String()))
		jm, _ := json.Marshal(m)
		mc.data.Save(m.UUID.String(), jm)
		d = []byte(m.UUID.String())
		fmt.Println(m)
	}
	// return if found (including just built)
	data, err := mc.data.Load(string(d))
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &m)
	if err != nil {
		panic(err)
	}
	return m
}

func (mc *ModelCollection) Put(m *Model) (err error) {
	fmt.Println("update", m.Name)
	jm, err := json.Marshal(m)
	if err != nil {
		return err
	}
	mc.data.Save(m.UUID.String(), jm)
	return err
}

func (mc *ModelCollection) Get(name string) (m *Model, err error) {
	d, err := mc.names.Load(name)
	if err != nil {
		return nil, errors.New("Model " + name + " not found")
	}
	data, _ := mc.data.Load(string(d))
	err = json.Unmarshal(data, &m)
	return m, err
}

func (mc *ModelCollection) List() (ml []*Model, err error) {
	ml = make([]*Model, 0)
	model_list := mc.names.List()
	for _, i := range model_list {
		md, err := mc.Get(i)
		if err != nil {
			return nil, err
		}
		ml = append(ml, md)
	}
	return ml, err
}
