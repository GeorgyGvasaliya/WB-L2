package memory

import (
	"11/internal/model"
	"sync"
)

type MemoryDB struct {
	sync.RWMutex
	data   map[int]model.Event
	lastID int
}

func New() *MemoryDB {
	return &MemoryDB{
		RWMutex: sync.RWMutex{},
		data:    make(map[int]model.Event, 0),
	}
}

func (m *MemoryDB) nextID() int {
	m.lastID++
	return m.lastID
}
