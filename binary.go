// Copyright (c) 2021 Andrej Giesbrecht

package store

import (
	"bytes"
	"encoding/gob"
	"os"
)

// ReadBinary loads the file into v and returns an error, if any.
func ReadBinary(file string, v interface{}) error {
	lock.Lock()
	defer lock.Unlock()

	var buf bytes.Buffer

	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = buf.ReadFrom(f)
	if err != nil {
		return err
	}

	dec := gob.NewDecoder(&buf)

	return dec.Decode(v)
}

// WriteBinary saves a representation of v to the file at path.
// It returns the number of bytes written and an error, if any.
func WriteBinary(file string, v interface{}) (int64, error) {
	lock.Lock()
	defer lock.Unlock()

	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)

	err := enc.Encode(v)
	if err != nil {
		return 0, err
	}

	f, err := os.Create(file)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	return buf.WriteTo(f)
}
