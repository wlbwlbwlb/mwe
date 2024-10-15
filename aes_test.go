package tool

import (
	"testing"
)

// 长度必须为16, 24或者32
var key = []byte("1234567812345678")

func TestAesEncrypt(t *testing.T) {
	enc, e := AesEncrypt([]byte("hello world"), key)
	if e != nil {
		t.Error(e)
	}
	dec, e := AesDecrypt(enc, key)
	if e != nil {
		t.Error(e)
	}
	t.Logf("%s", dec)
}

func BenchmarkAesEncrypt(b *testing.B) {
	for n := 0; n < b.N; n++ {
		enc, e := AesEncrypt([]byte("hello world"), key)
		if e != nil {
			b.Error(e)
		}
		_, e = AesDecrypt(enc, key)
		if e != nil {
			b.Error(e)
		}
	}
}
