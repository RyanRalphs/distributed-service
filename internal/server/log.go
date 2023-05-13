package server

import (
	"fmt"
	"sync"
)

type Record struct {
	Value []byte `json:"value"`
	Index uint64 `json:"index"`
}

type Log struct {
	mu      sync.Mutex
	records []Record
}

func NewLog() *Log {
	return &Log{}
}

var ErrIndexNotFound = fmt.Errorf("index not found")

func (c *Log) Append(record Record) (uint64, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	record.Index = uint64(len(c.records))
	c.records = append(c.records, record)
	return record.Index, nil
}

func (c *Log) Read(index uint64) (Record, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if index >= uint64(len(c.records)) {
		return Record{}, ErrIndexNotFound
	}
	return c.records[index], nil
}
