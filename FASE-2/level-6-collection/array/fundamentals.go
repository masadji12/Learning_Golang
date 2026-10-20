// Array ukurannya tetap setelah dideklarasikan
// CREATE: Mengisi elemen array melalui indeks yang tersedia
// READ: Mengambil elemen menggunakan index, or reading of element with range
// UPDATE/EDIT: Mengganti nilai pada index tertentu.
// DELETE: Array tidak bisa diperkecil secara langsung. hanya bisa mengosongkan nilai elemen, tetapi elemennya tidak ada.

package main

import "fmt"

func main() {
	var atks [3]string

	atks[0] = "Pulpen"
	atks[1] = "Buku"
	atks[2] = "Penghapus"

	// Read
	fmt.Println(atks)
	for i, atk := range atks {
		fmt.Println(i + 1, atk)
	}

	// Edit
	atks[1] = "Buku Tulis"
	fmt.Println(atks)

	// Read once data
	fmt.Println(atks[1])
}