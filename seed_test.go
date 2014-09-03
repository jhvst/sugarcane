package sugarcane

import (
	"io"
	"os"
	"testing"
)

type P struct {
	Name   string
	Visits int
}

func BenchmarkInsert(b *testing.B) {

	b.StopTimer()

	w, err := Open("test_db")
	if err != nil {
		panic(err)
	}

	b.StartTimer()

	var person P
	person.Name = "juuso"
	person.Visits = 7

	for i := 0; i < b.N; i++ {
		w.Insert(person)
	}

}

func BenchmarkRead(b *testing.B) {

	b.StopTimer()

	w, err := Open("test_db")
	if err != nil {
		panic(err)
	}

	b.StartTimer()

	data, err := w.Read()
	if err != nil {
		panic(err)
	}

	for i := 0; i < b.N; i++ {
		var q P
		//fmt.Println("Bytes left:", len(data.Bytes()))
		err := w.Scan(&q, data)
		if err == io.EOF {
			break
		}
		//fmt.Println(q)
	}

}

func BenchmarkCleanUp(b *testing.B) {
	os.Remove("test_db")
}
