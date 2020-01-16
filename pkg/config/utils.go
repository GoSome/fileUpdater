package config

import (
	"bytes"
	"encoding/gob"
)

func deepCopy(target, source interface{}) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(source); err != nil {
		return err
	}
	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(target)
}
