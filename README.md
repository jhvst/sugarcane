sugarcane
=========

The sweet embedded database in Go

Under heavy development!

Sugarcane is a persistent data-store for Go structures. Data is saved in binary and loaded as byte buffer, so you don't need to encode data to different format, say JSON. Databases saved by sugarcane are also lightweight; one million structures of three field structures weight around 50-100MB.

Sugarcane has no dependencies outside of the Go base library, so it will compile on every platform supported by Go. 

Saving object into disk is as easy as

	type Person struct {
		Name string
		Age
	}

	var p Person
	p.Name = "foo"
	p.Age = 18

	sugarcane.Insert(p)

You can then read a single structure with

	sugarcane.ReadOne(&p)

You can also read the whole file with for loop

	for i := 0; ; i++ {
		var q Person
		err := sugarcane.ReadOne(&q)
		if err == io.EOF {
			break
		}
		fmt.Println(q.Name) // will print "foo"
	}

##Performance

At the moment, sugarcane is rather naive implementation and does not include any optimizations whatsoever. Therefore, it should be no surprise that sugarcane is not as fast as, say PostgreSQL. To get a better look at the performance, git clone the repository and run `go test -bench=".*"`. On my Macbook I get the following results:

	BenchmarkInsert	  200000	      8642 ns/op
	BenchmarkRead	   50000	     50118 ns/op
	BenchmarkDecode	   50000	     50275 ns/op 
	ok	8.019s

Inserting one million lines of three field structures takes:

	real	0m9.574s
	user	0m7.942s
	sys		0m1.944s

Reading however...

	real	1m13.020s
	user	1m9.370s
	sys		0m3.823s

This makes sugarcane about 50 times slower than PostgreSQL.

##Known bugs

1. Buffer is currently drained upon reading

##License

MIT