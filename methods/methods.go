package methods

import (
	"bytes"
	"io"
	"os"
	"strings"
)

func StringFromFile(filename string) (data string) {
	// heavy use of https://stackoverflow.com/questions/1821811/how-to-read-write-from-to-a-file-using-go
	fi, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	// read to buffer
	buf := make([]byte, 1024)
	for {
		// read a chunk
		n, err := fi.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}
	}
	// buffer -> string
	buf = bytes.Trim(buf, "\x00")
	data = strings.TrimSpace(string(buf))
	// close fi on exit and check for its returned error
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()
	return
}
