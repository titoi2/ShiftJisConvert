package main

import (
	"./sjisconv"
	"bufio"
	"io"
	"os"
	"fmt"
	"unicode/utf16"
	"unicode/utf8"
)


func readLn(p *bufio.Reader) ([]byte, error) {
	input, err := p.ReadString('\n')
	return []byte(input), err
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		var strbuf string = ""
		src, err := readLn(reader)
		if err == nil || err == io.EOF {
			utf16str := sjisconv.SjistoUtf16(src)
			runes := utf16.Decode(utf16str)
			buf := make([]byte, 6)
			for i :=0; i<len(runes) ; i++ {
				n := utf8.EncodeRune(buf, runes[i])
				strbuf += string(buf[0:n])
			}
			
			fmt.Print(strbuf)
		}
		if err != nil {
			break
		}
	}
}
