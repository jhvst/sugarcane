package sugarcane

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"io"
	"os"
)

// Prepare encodes struct into byte.Buffer.
func prepare(p interface{}) (bytes.Buffer, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(p)
	if err != nil {
		return buf, err
	}
	return buf, nil
}

// write writes byte buffer to file f.
func write(f *os.File, buf bytes.Buffer) error {
	w := bufio.NewWriter(f)
	_, err := w.Write(buf.Bytes())
	if err != nil {
		return err
	}
	w.Flush()
	return nil
}

// Open opens a file for writing.
func Open(filename string) (*os.File, error) {
	w, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		return w, err
	}
	return w, nil
}

// Read reads a file
func Read(filename string) (*bytes.Buffer, error) {
	f, err := os.OpenFile(filename, os.O_RDONLY, 0600)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		return nil, err
	}

	buf := make([]byte, fi.Size()) // make the buffer as big as the file
	r := bufio.NewReader(f)
	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			return nil, err
		}
		if n == 0 {
			break
		}
	}
	return bytes.NewBuffer(buf), nil
}

// Insert prepares structure for disk saving.
func Insert(p interface{}, f *os.File) error {
	binary, err := prepare(p)
	if err != nil {
		return err
	}
	write(f, binary)
	return nil
}

// Scan returns a first structure from byte buffer and decodes it according given structure.
func Scan(p interface{}, data *bytes.Buffer) error {
	dec := gob.NewDecoder(data)
	err := dec.Decode(p)
	if err != nil {
		return err
	}
	return err
}
