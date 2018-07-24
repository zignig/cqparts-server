package main

import (
	"github.com/satori/go.uuid"
	"time"
)

// the construct for a single model
// Flesh this out so it has all the info for a given model

// have segmented update and dirty write down
// get as much information as possible from the cqparts display post

type Model struct {
	Name    string    `json:"name"`
	Img     string    `json:"img"`
	View    Render    `json:"view"`
	Status  string    `json:"status"`
	UUID    uuid.UUID `json:"uuid"`
	Created time.Time `json:"created"`
	Pinned  bool      `json:"pinned"`
}

type ModelCollection struct {
	data Storage
}

func NewModelCollection(s Storage) (mc *ModelCollection) {
	mc = &ModelCollection{
		data: s,
	}
	return mc
}
