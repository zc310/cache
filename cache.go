package cache

import (
	"bytes"
	"encoding/binary"
	"time"
)

type Cache interface {
	Get(key []byte) ([]byte, bool)
	GetRange(key []byte, low, high int64) ([]byte, bool)
	Set(key []byte, value []byte) error
	SetTimeout(key []byte, value []byte, timeout time.Duration) error
	Delete(key []byte) error
	Clean() error
	Empty() error
}

type Value struct {
	Timeout int64
	Size    int64
	Value   []byte
}

func (p *Value) Valid() bool {
	return p.Timeout == 0 || p.Timeout > time.Now().Unix()
}

func Unmarshal(b []byte, cv *Value) error {
	buf := bytes.NewBuffer(b)

	err := binary.Read(buf, binary.LittleEndian, &cv.Timeout)
	if err != nil {
		return err
	}
	err = binary.Read(buf, binary.LittleEndian, &cv.Size)
	if err != nil {
		return err
	}

	cv.Value = make([]byte, cv.Size)
	err = binary.Read(buf, binary.LittleEndian, cv.Value)
	return err
}

func Marshal(p *Value) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, p.Timeout)
	if err != nil {
		return nil, err
	}

	err = binary.Write(buf, binary.LittleEndian, p.Size)
	if err != nil {
		return nil, err
	}
	buf.Write(p.Value)
	return buf.Bytes(), nil
}
