## Inheritence / Pewarisan
Pewarisan didalam bahasan pemograman adalah sebuah kelas yang dapat menurunkan `property` dan `method` ke kelas turunannya.
Permisalannya:

ada kelas `mobil` yang memiliki property kecepatan, warna, bahan bakar dan juga memiliki method jumlah roda, maju, mundur, bip.
dari kelas mobil nanti diturunkan ke kelas lain panter, kijang, avanza.

### example
```
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

func main() {
	a := Avanza{Car{4}}
	fmt.Println("Inheritance\n===============")
	fmt.Println("Jumlah roda :", a.hitungRoda())
}

```

```
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
```
Response:
```
Inheritance
===============
Avanza
Jumlah roda : 4
Hai, saya avanza
===============
Ferari
Jumlah roda: 4
Hallo, saya ferari
```
