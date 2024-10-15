package tool

import (
	"bytes"
	"encoding/json"
)

func Compact(bs []byte) (s string, e error) {
	var buf bytes.Buffer
	if e = json.Compact(&buf, bs); e != nil {
		return
	}
	return buf.String(), e
}
