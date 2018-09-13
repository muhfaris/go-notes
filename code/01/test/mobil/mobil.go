package mobil

type jumlah int

// NewJumlah membuat dan mengembalikan object dari unexported tipe jumlah
func NewJumlah(value int) jumlah {
	return jumlah(value)
}
