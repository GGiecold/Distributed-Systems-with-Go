package log

import (
	"bufio"
	"encoding/binary"
	"os"
	"sync"
)

var (
	encoding = binary.BigEndian
)

const (
	lenWidth = 8
)

type store struct {
	*os.File
	mutex sync.Mutex
	buffer *bufio.Writer
	size uint64
}

func newStore(f *os.File) (*store, error) {
	info, err := os.Stat(f.Name())
	if err != nil {
		return nil, nil
	}
	size := uint64(info.Size())
	return &store{
		File: f,
		size: size,
		buffer: bufio.NewWriter(f),
	}, nil
}

func (s *store) Append(p []byte) (n uint64, pos uint64, err error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	pos := s.size
	if err := binary.Write(s.buffer, encoding, uint64(len(p))); err != nil {
		return 0, 0, err
	}
	w, err := s.buffer.Write(p)
	if err != nil {
		return 0, 0, err
	}
	w += lenWidth
	s.size += uint64(w)
	return uint64(w), pos, nil
}


