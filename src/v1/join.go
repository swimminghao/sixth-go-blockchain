package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main01() {
	strs := []string{"hello", "world", "!"}
	res := strings.Join(strs, "")
	fmt.Println(res)
	res1 := bytes.Join([][]byte{[]byte("hello"), []byte("world"), []byte("!")}, []byte{})
	fmt.Printf("res1: %s\n", res1)
}
