// Copyright (c) 2021 Andrej Giesbrecht

package store

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
)

// marshal is a function that marshals the object into an io.Reader.
// By default, it uses the JSON marshaller.
var marshal = func(v interface{}) (io.Reader, error) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(b), nil
}

// unmarshal is a function that unmarshals the data
// from the reader into the specified value.
var unmarshal = func(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}

// ReadJSON loads the file into v and returns an error, if any.
func ReadJSON(file string, v interface{}) error {
	lock.Lock()
	defer lock.Unlock()

	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	return unmarshal(f, v)
}

// WriteJSON saves a representation of v to the file.
// It returns the number of bytes written and an error, if any.
func WriteJSON(file string, v interface{}) (int64, error) {
	lock.Lock()
	defer lock.Unlock()

	f, err := os.Create(file)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	r, err := marshal(v)
	if err != nil {
		return 0, err
	}

	return io.Copy(f, r)
}
