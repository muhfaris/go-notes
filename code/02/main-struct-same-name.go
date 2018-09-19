package main

import "fmt"

type Dapur struct {
	jmlRuangan int
}

type Rumah struct {
	Dapur
	jmlRuangan int
}

func main() {
	r := Rumah{Dapur{5}, 10}
	fmt.Println("Anonymouse field in struct\n==========================")
	fmt.Printf("Rumah memiliki %d ruangan\n", r.jmlRuangan)
	fmt.Printf("Rumah memiliki %d ruangan di dapur\n", r.Dapur.jmlRuangan)
}
