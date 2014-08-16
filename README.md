sugarcane [![GoDoc](https://godoc.org/github.com/9uuso/sugarcane?status.svg)](https://godoc.org/github.com/9uuso/sugarcane)
=========

The sweet embedded database in Go

Sugarcane is a persistent data-store for Go structures. Data is saved with the output of `encoding/gob` to disk, from which it can be loaded as byte buffer. You could for example save application progress in native Go structure, so you don't need to encode your data to different format, say JSON.

Databases saved by sugarcane are also lightweight; one million lines of three field structures weight only around 50-100MB!

Sugarcane also has no dependencies outside of the Go standard library.

Saving object into disk is as easy as

	type Person struct {
		Name string
		Age
	}

	var p Person
	p.Name = "foo"
	p.Visits = 3

	w, _ := sugarcane.Open("./person_table")

	sugarcane.Insert(p, w)

You can then read a single structure with

	data, _ := sugarcane.Read("person_table")

	sugarcane.Scan(&p, data)

You can also read the whole file with for loop

	data, _ := sugarcane.Read("person_table")

	var persons []Person

	for i := 0; ; i++ {
		var q Person
		err := sugarcane.Scan(&q, data)
		if err == io.EOF {
			break
		}
		persons = append(persons, p)
	}

	fmt.Println(persons)

##Performance

At the moment, sugarcane is rather naive implementation and does not include any optimizations whatsoever. Therefore, it should be no surprise that sugarcane is not as fast as, say PostgreSQL. To get a better look at the performance, git clone the repository and run `go test -bench=".*"`. On my Macbook I get the following results:

	BenchmarkInsert	  200000	      8642 ns/op
	BenchmarkRead	   50000	     50118 ns/op
	ok	4.873s

Inserting one million lines of three field structures takes:

	real	0m9.574s
	user	0m7.942s
	sys		0m1.944s

Reading however...

	real	1m13.020s
	user	1m9.370s
	sys		0m3.823s

This makes sugarcane about 50 times slower than PostgreSQL.

##License

MIT
