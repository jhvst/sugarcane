package main

import (
	"code.google.com/p/go-uuid/uuid"
	"io"
	"log"
	"sugarcane"
)

type P struct {
	Name  string
	Age   int
	ID    string
	City  string
	Games int
}

// This example shows the basic usage of the package: Create an encoder,
// transmit some values, receive them with a decoder.
func main() {

	w, err := sugarcane.Open("./db")
	if err != nil {
		panic(err)
	}

	for i := 0; i < 30; i++ {
		var person P

		person.Name = "juuso"
		person.City = uuid.New()
		person.ID = uuid.New()
		person.Age = 18
		person.Games = i

		err := w.Insert(person)
		if err != nil {
			panic(err)
		}
	}

	for i := 0; ; i++ {
		var q P
		err := w.Scan(&q)
		if q.Games == 2 {
			err := w.Delete(&q)
			if err != nil {
				panic(err)
			}
		}
		if err == io.EOF {
			log.Println("Everything read. Found", i, "occurances.")
			break
		}
		log.Println(q)
	}

}
