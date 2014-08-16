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

	f, err := os.OpenFile("test_db", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	b.StartTimer()

	var person P
	person.Name = "juuso"
	person.Visits = 7

	for i := 0; i < b.N; i++ {
		Insert(person, f)
	}

}

func BenchmarkRead(b *testing.B) {

	b.StopTimer()

	f, err := os.OpenFile("test_db", os.O_RDONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	b.StartTimer()

	data, err := Read("test_db")
	if err != nil {
		panic(err)
	}

	for i := 0; i < b.N; i++ {
		var q P
		//fmt.Println("Bytes left:", len(data.Bytes()))
		err := Scan(&q, data)
		if err == io.EOF {
			break
		}
		//fmt.Println(q)
	}

}

func BenchmarkCleanUp(b *testing.B) {
	os.Remove("test_db")
}
