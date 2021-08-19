package server

import (
	"fmt"
	"sync"
)

type Record struct {
	Value []byte `json:"value"`
	Offset uint64 `json:"offset"`
}

var ErrorOffsetNotFound = fmt.Errorf("Offset could not be found")

type Log struct {
	mutex sync.Mutex
	records []Record
}

func NewLog() *Log {
	return &Log{}
}

func (log *Log) Append(record Record) (uint64, error) {
	log.mutex.Lock()
	defer log.mutex.Unlock()
	record.Offset = uint64(len(log.records))
	log.records = append(log.records, record)
	return record.Offset, nil
}

func (log *Log) Read(offset uint64) (Record, error) {
	log.mutex.Lock()
	defer log.mutex.Unlock()
	if offset >= uint64(len(log.records)) {
		return Record{}, ErrorOffsetNotFound
	}
	return log.records[offset], nil
}

