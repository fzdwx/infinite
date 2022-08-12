package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"os"
	"unicode/utf16"
)

func main() {
	// Write a file containing UTF-16 encoded data.
	const file = "./foo.data"
	x := utf16.Encode([]rune("世界"))
	of, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal(err)
	}
	if err := binary.Write(of, binary.LittleEndian, x); err != nil {
		log.Fatal(err)
	}
	of.Close()

	// Read back the bytes from that file.
	content, err2 := os.ReadFile(file)
	if err2 != nil {
		log.Fatal(err2)
	}

	// Convert to string (assumes UTF-8).
	fmt.Printf("%s\n", string(content))

	// Better conversion.
	var nx [2]uint16
	b := bytes.NewBuffer(content)
	if err := binary.Read(b, binary.LittleEndian, &nx); err != nil {
		log.Fatalf("binary.Read: %v", err)
	}
	fmt.Printf("after binary.Read: %+v\n", nx)
	v := utf16.Decode(nx[:])
	fmt.Printf("%c %c\n", v[0], v[1])
}
