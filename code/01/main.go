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
