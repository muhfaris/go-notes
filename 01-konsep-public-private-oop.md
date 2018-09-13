## Konsep identifikasi public dan private OOP di Golang
GO menerapkan cara yang berbeda untuk hal ini. jika pernah belajar java untuk menerapkan modifier oop dalam sebuah kode java, cukup menggunakan modifier private atau public atau protected.
secara umum ada 3 modifier dalam java :

|Modifier | Class | Package | SubClass | World|
|---------|-------|---------|----------|------|
|public | Y | Y | Y | Y |
|Protected | Y | Y | Y | N |
|no modifier | Y | Y | N | N |
|private | Y | N | N | N | N |


*SubClass : class anak*

*World : seluruh package di aplikasi*

kembali ke GO, bagaimana dengan golang ?

coba perhatikan baik-baik, ketika mulai belajar golang akan kita temui penamaan di golang menggunakan `uppercase` dan `lowercase` di huruf pertamanya dalam penamaan `type`, `variabel`, `function`.
hal ini memberikan penjelasan bahwa huruf pertama kapital adalah sebagai `public` yang mana dapat digunakan oleh setiap kode yang akan menggunakannya. sedangkan ketika huruf pertama kecil adalah sebagai `private` dan hanya dapat di akses dalam satu `package`.

menurut ku penggunakan kata `public` dann `private` masih kurang akurat di golang, melainkan lebih ke kata `exported` dan `unexported`.
ketika menggunakan `exported` di sebuah `package` itu berarti identifier dapat di akses dari `package` lain. dan sebaliknya untuk `unexported`, identifier hanya bisa di akses di `package` itu sendiri.


### Example Exported
```
package mobil
type  Jumlah int
```
dari contoh diatas, jumlah bertipe data `int` didefinisikan didalam package `mobil`. penamaannya juga menggunakan huruf kapital `J`, yang mana adalah tipe `exported`, dapat di akses di package lain.

```
package main

import (
	"fmt"

	"github.com/muhfaris/go-notes/code/01/test/mobil"
)

func main() {
	fmt.Println("Pengenalanan Exported:")
	jml := mobil.Jumlah(10)
	fmt.Printf("Jumlah:%d", jml)
}
```
ketika dijalankan maka program akan berjalan dengan semestinya tidak ada error dan akan menampilkan value `10`.

```
Pengenalanan Exported:
Jumlah:10
[Process exited 0]
```

sekarnag cobalah untuk mengganti type `Jumlah` di package `mobil` menjadi `unexported`.
```
package mobil
type  jumlah int
```

kemudian rubah juga bagian main.go, menjadi seperti berikut:
```
package main

import (
	"fmt"

	"github.com/muhfaris/go-notes/code/01/test/mobil"
)

func main() {
	fmt.Println("Pengenalanan Exported:")
	jml := mobil.jumlah(10)
	fmt.Printf("Jumlah:%d", jml)
}
```
ketika di jalankan pasti akan mendapatkan error dengan keterangan variabel dari `jumlah` bukan tipe `Exported` melainkan `Unexported`
```
# command-line-arguments
./main.go:11:9: cannot refer to unexported name mo
bil.jumlah
./main.go:11:9: undefined: mobil.jumlah

[Process exited 2]
```

agar bisa mengakses variabel `unexported`, kita bisa menggunakan cara lain yaitu dengan  membuat fungsi baru yang mengembalikan nilai dengan tipe `unexported` variabelnya tadi  di dalam package `mobil`.
jadi si `unexported` ini tidak bisa di akses secara langsung, tapi masih bisa di akses dengan cara lain.

```
//mobile.go
package mobil

type jumlah int

// NewJumlah membuat dan mengembalikan object dari unexported tipe jumlah
func NewJumlah(value int) jumlah {
	return jumlah(value)
}
```
```
// main.go
package main

import (
	"fmt"

	"github.com/muhfaris/go-notes/code/01/test/mobil"
)

func main() {
	fmt.Println("Pengenalanan Exported:")
	jml := mobil.NewJumlah(10)
	fmt.Printf("Jumlah:%d", jml)
}
```
jika ingin mempelajari lebih dalam ada satu library yang bagus untuk pengenalan `exported` dan `unexported`,
librarynya itu adalah `time.Time`
*modifier: tingkatan akses suatu fungsi, variabel dalam sebuah kode - pn*


