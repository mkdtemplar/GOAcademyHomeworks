package main

import (
	"fmt"
	"io"
	"os"
	strings "strings"
)

type ReverseStringReader struct {
	io.Reader
}

func NewReverseStringReader(input string) *ReverseStringReader {

	r := &ReverseStringReader{strings.NewReader(input)}

	//ourReverseStringReader, err := io.Copy(os.Stdout, r)

	return r
}

func (receiver *ReverseStringReader) ReverseString() string {

	var input string
	receiver = NewReverseStringReader(input)

	//buffer := make([]byte, 1)

	result, err := io.Copy(os.Stdout, receiver)

	if err == nil {
		return string(result)
	} else {
		return fmt.Sprintf("%s", err)
	}

}

func main() {

	r := NewReverseStringReader("Ivan Markovski")

	res, _ := io.Copy(os.Stdout, *r)
	for _, ch := range string(res) {
		fmt.Println(ch)
	}
	fmt.Println()
	fmt.Println(string(res))
}
