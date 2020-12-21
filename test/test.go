// Golang program to illustrate the usage of
// io.TeeReader() function

// Including main package
package main

// Importing fmt, io, strings, bytes, and os
import (
	"bytes"
	"io"
	"os"
	"strings"
)

// Calling main
func main() {

	// Defining reader using NewReader method
	reader := strings.NewReader("GfGsadjhaskjd")

	f, _ := os.Open("gg.php")

	// Defining buffer
	var buf bytes.Buffer

	// Calling TeeReader method with its parameters
	r := io.TeeReader(reader, &buf)

	// Calling Copy method with its parameters
	_, err := io.Copy(f, r)

	// If error is not nil then panics
	if err != nil {
		panic(err)
	}

	// Prints output
	//fmt.Printf("n: %v\n", Reader)

	//fmt.Println(buf)

}
