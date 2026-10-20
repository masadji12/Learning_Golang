package main

import "fmt"

func main() {
	// 1. Buat array bernama 'atks' yang memiliki 4 elemen bertipe string
	var atks [4]string

	// 2. Isi array dengan Pulpen, Buku, Penghapus, dan Pensi
	atks[0] = "Pulpen"
	atks[1] = "Buku"
	atks[2] = "Penghapus"
	atks[3] = "Pensil"

	// 3. Tampilkan seluruh data menggunakn 'fmt.Println'
	fmt.Println(atks)

	// 4. Ubah Buku menjadi Buku Tulis
	atks[1] = "Buku Tulis"

	// 5. Kosongkan data penghapus
	atks[2] = ""

	// Menampilkan data array new
	fmt.Println(atks)
}