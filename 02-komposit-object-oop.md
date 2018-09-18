## Komposit Objek
Komposit Objek merupakan lifecycle object yang bergantung pada object ownernya.
ketika ada suatu objek yang di buat di dalam objek lain (owner), dan jika objek lain ini di hapus maka objek tersebut juga hilang.

Dalam bahasa golang komposit objek tidak sama dengan bahasa java, caranya berbeda tapi memiliki sifat yang sama dengan komposit objek.
jika di Golang istilah komposit objek mungkin kurang tepat, saya lebih suka menyebutnya dengan `anonymouse field in struct`.

ilustrasi yang akan di buat adalah ada sebuah rumah, di dalam rumah ada ruangan dengan jumlah 10 ruangan. Di ruangan dapur ada banyak 5 piring.
program akan menampilkan jumlah ruangan dalam rumah dan menampilkan jumlah piring di ruangan dapur.

```
package main

import "fmt"

type Dapur struct {
	jmlPiring int
}

type Rumah struct {
	Dapur
	jmlRuangan int
}

func main() {
	r := Rumah{Dapur{5}, 10}
	fmt.Println("Anonymouse field in struct\n==========================")
	fmt.Printf("Rumah memiliki %d ruangan\n", r.jmlRuangan)
	fmt.Printf("Rumah memiliki %d piring\n", r.jmlPiring)
}
```

Response:
```
Anonymouse field in struct
==========================
Rumah memiliki 10 ruangan
Rumah memiliki 5 piring
```

### Ananonymouse fields - kesamaan nama variabel
Apa yang akan terjadi jika ada kesamaan nama variabel di dalam struct ?

dari contoh sebelumnya, kita rubah struct dapur dengan mengganti `jmlPiring` menjadi `jmlRuangan`. mungkin hal seperti ini akan sering muncul dalam projek yang besar.
cara pemanggilannya pun berbeda, untuk memanggil variabel `jmlPiring` di sturct `Dapur` dengan mengganti perintah menjadi `r.Dapur.jmlPiring`

```
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
```

Response:
```
Anonymouse field in struct
==========================
Rumah memiliki 10 ruangan
Rumah memiliki 5 ruangan di dapur
```
