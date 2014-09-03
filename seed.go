package sugarcane

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"io"
	"runtime"
	"os"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

type Database struct {
	File *os.File
	Filename string
}

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
func Open(filename string) (Database, error) {
	var d Database
	w, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		return d, err
	}
	d.File = w
	d.Filename = filename
	return d, nil
}

// Read reads a file
func (d Database) Read() (*bytes.Buffer, error) {
	f, err := os.OpenFile(d.Filename, os.O_RDONLY, 0600)
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
func (d Database) Insert(p interface{}) error {
	binary, err := prepare(p)
	if err != nil {
		return err
	}
	write(d.File, binary)
	return nil
}

// Scan returns a first structure from byte buffer and decodes it according given structure.
func (d Database) Scan(p interface{}, data *bytes.Buffer) error {
	dec := gob.NewDecoder(data)
	err := dec.Decode(p)
	if err != nil {
		return err
	}
	return err
}
