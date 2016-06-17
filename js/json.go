package js

import (
	"bytes"
	"encoding/json"
)

func JsonObjectToString(i interface{}) (string, error) {
	buf := &bytes.Buffer{}
	jencoder := json.NewEncoder(buf)
	err := jencoder.Encode(i)
	return buf.String(), err
}
