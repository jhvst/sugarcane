package main

import (
	"code.google.com/p/go-uuid/uuid"
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

	// Open file for reading. Only use a single file for a single structure.
	// Consider one file as one table.
	w, err := sugarcane.Open("./db")
	if err != nil {
		panic(err)
	}

	// Lets start by inserting 30 rows of data to the said table.
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

	// Now we start reading our data from the cache file.
	for i := 0; ; i++ {

		var q P
		w, err = w.Scan(&q) // This line will magically convert the empty q struct into a stored struct in our cache.

		if q.Games == 2 { // Let's manipulate our data and tell the program that if the user has two games, delete the whole struct.
			err := w.Delete(&q)
			if err != nil {
				panic(err)
			}
		}

		if q.Name == "juuso" { // Actually we made a mistake earlier and we want to replace every user's name to contain a capital J

			// For that though, we need to create a new structure, to which we assign the old one and replace the field we want.
			var newperson P
			newperson = q
			newperson.Name = "Juuso"

			err := w.Update(q, newperson) // Here we make a call to replace the old struct q with our new one, newperson.
			if err != nil {
				panic(err)
			}

		}

		if err == sugarcane.Empty { // Once we have drained the whole cache, we log the amount of structs we read. For fun.
			break
		}
	}

	// Finally we iterate trough the file to check that the changes we applied earlier have been made.
	for i := 0; ; i++ {
		var q P
		w, err = w.Scan(&q)
		if err == sugarcane.Empty {
			log.Println("Everything read. Found", i, "occurances.")
			break
		}
		log.Println(q)
	}

}
