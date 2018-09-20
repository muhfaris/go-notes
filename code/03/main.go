package main

import "fmt"

type Car struct {
	jmlRoda int
}

// method
func (c Car) hitungRoda() int {
	return c.jmlRoda
}

type Avanza struct {
	Car // anonymouse field
}

func (a Avanza) sayAvanza() {
	fmt.Println("Hai, saya avanza")
}

type Ferari struct {
	Car
}

func (f Ferari) sayFerari() {
	fmt.Println("Hallo, saya ferari")
}

func main() {
	a := Avanza{Car{4}}
	f := Ferari{Car{4}}
	fmt.Println("Inheritance\n===============")
	fmt.Println("Avanza")
	fmt.Println("Jumlah roda :", a.hitungRoda())
	a.sayAvanza()
	fmt.Println("===============\nFerari")
	fmt.Println("Jumlah roda:", f.hitungRoda())
	f.sayFerari()

}
