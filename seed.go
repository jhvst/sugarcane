package sugarcane

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"io"
	"os"
)

type Database struct {
	File     *os.File
	Filename string
	Cache    *bytes.Buffer
}

// prepare encodes structs into buffers
func prepare(p interface{}) (bytes.Buffer, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(p)
	if err != nil {
		return buf, err
	}
	return buf, nil
}

// write writes data to disk
func write(f *os.File, buf bytes.Buffer) error {
	w := bufio.NewWriter(f)
	_, err := w.Write(buf.Bytes())
	if err != nil {
		return err
	}
	w.Flush()
	return nil
}

// binappend appends data to cache
func (d Database) binappend(binary bytes.Buffer) error {
	w := bufio.NewWriter(d.Cache)
	_, err := w.Write(binary.Bytes())
	if err != nil {
		return err
	}
	w.Flush()
	return nil
}

// Open opens a file for writing.
func Open(filename string) (Database, error) {
	var d Database
	w, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		return d, err
	}
	d.File = w
	d.Filename = filename

	fi, err := w.Stat()
	if err != nil {
		return d, err
	}

	// read the old file
	buf := make([]byte, fi.Size()) // make the buffer as big as the file
	r := bufio.NewReader(w)
	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			return d, err
		}
		if n == 0 {
			break
		}
	}

	// save buffer to cache
	d.Cache = bytes.NewBuffer(buf)
	return d, nil
}

// Insert prepares structure for disk saving.
func (d Database) Insert(p interface{}) error {
	binary, err := prepare(p)
	if err != nil {
		return err
	}
	err = d.binappend(binary)
	if err != nil {
		return err
	}
	write(d.File, binary)
	return nil
}

// Scan returns a first structure from byte buffer and decodes it according given structure.
func (d Database) Scan(p interface{}) error {
	dec := gob.NewDecoder(d.Cache)
	err := dec.Decode(p)
	if err != nil {
		return err
	}
	return nil
}
