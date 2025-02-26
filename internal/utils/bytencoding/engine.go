package bytencoding

import (
	"bytes"
	"encoding/gob"
)

func Encode(p interface{}) ([]byte, error) {
	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)

	err := enc.Encode(p)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func Decode(b []byte, dest interface{}) error {
	dec := gob.NewDecoder(bytes.NewReader(b))
	err := dec.Decode(dest)
	if err != nil {
		return err
	}

	return nil
}
