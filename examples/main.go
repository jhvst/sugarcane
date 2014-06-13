package main

import (
	"io"
	"log"
	"sugarcane"
)

type P struct {
	Name   string
	Age    int
	Visits int
}

// This example shows the basic usage of the package: Create an encoder,
// transmit some values, receive them with a decoder.
func main() {

	w := sugarcane.Open("./db")
	defer w.Close()

	var person P
	person.Name = "juuso"
	person.Visits = 7
	person.Age = 18

	for i := 0; i < 1000000; i++ {
		sugarcane.Insert(person, w)
	}

	// data returned as byte buffer from database
	data := sugarcane.Read("./db")

	for i := 0; ; i++ {
		var q P
		//log.Println("Bytes left:", len(data.Bytes()))
		err := sugarcane.ReadOne(&q, data)
		if err == io.EOF {
			log.Println("Everything read. Found", i, "occurances.")
			break
		}
		//log.Println(q)
	}
}
