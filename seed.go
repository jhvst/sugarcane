package sugarcane

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"io"
	"log"
	"os"
)

// Prepare encodes struct into byte.Buffer.
func prepare(p interface{}) bytes.Buffer {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(p)
	if err != nil {
		log.Fatal("encode error:", err)
	}
	return buf
}

// Write writes byte buffer to file f.
func Write(f *os.File, buf bytes.Buffer) {
	w := bufio.NewWriter(f)
	_, err := w.Write(buf.Bytes())
	if err != nil {
		panic(err)
	}
	w.Flush()
}

// Open opens a file for writing.
func Open(filename string) *os.File {
	w, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		panic(err)
	}
	return w
}

// Read reads a file
func Read(filename string) *bytes.Buffer {
	f, err := os.OpenFile(filename, os.O_RDONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		panic(err)
	}

	buf := make([]byte, fi.Size()) // make the buffer as big as the file
	r := bufio.NewReader(f)
	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}
	}
	return bytes.NewBuffer(buf)
}

// Insert prepares structure for disk saving.
func Insert(p interface{}, f *os.File) {
	binary := prepare(p)
	Write(f, binary)
}

// ReadOne returns a first structure from byte buffer and decodes it according given structure.
func ReadOne(p interface{}, data *bytes.Buffer) error {
	dec := gob.NewDecoder(data)
	err := dec.Decode(p)
	if err != nil {
		return err
	}
	return err
}
