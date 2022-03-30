package main

import (
	"bufio"
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

	return r
}

func main() {

	var str string
	var reverseString string

	fmt.Print("Enter string: ")
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	str, _ = reader.ReadString('\n')

	str = strings.Replace(str, "\n", "", -1)
	str = strings.Replace(str, "\r", "", -1)

	r := NewReverseStringReader(str)

	buff := make([]byte, len(str))

	for {
		_, err := r.Read(buff)

		if err == io.EOF {
			break
		}
	}

	fmt.Println()

	for i := len(buff) - 1; i >= 0; i-- {
		reverseString += string(buff[i])
	}

	r = NewReverseStringReader(reverseString)

	res, _ := io.Copy(writer, r)

	fmt.Println(string(res))

}
