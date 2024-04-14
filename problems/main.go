package main

import (
	"fmt"
)

func main() {
	s := []string{"a", "a", "a", "b", "b", "a", "a"}
	bs := []byte{}
	for _, ss := range s {
		fmt.Printf("[]byte(ss): %v\n", []byte(ss))
		b := []byte(ss)
		bs = append(bs, b[0])
	}
	fmt.Printf("bs: %v\n", bs)

	fmt.Println(compress(bs))
	fmt.Printf("bs: %v\nv: ", bs)
	for _, v := range bs {
		fmt.Printf("%v, ", string(v))
	}
	fmt.Println("")
}
