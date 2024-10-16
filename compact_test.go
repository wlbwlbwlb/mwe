package mwe

import (
	"fmt"
	"testing"
)

func TestCompact(t *testing.T) {
	buf := []byte(`{
	"name":"fa",
	"age":18
}`)
	fmt.Println(string(buf))
	fmt.Println(Compact(buf))
}
