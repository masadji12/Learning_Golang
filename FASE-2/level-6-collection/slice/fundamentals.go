// Slice adalah struktur data GO yang digunakan untuk menyimpan kumpulan element dengan ukuran yang bisa bertambah dan berkurang

package main

import "fmt"

func main() {
	atks := []string {
		"Pulpen",
		"Buku",
		"Penghapus",
	}

	// READ: Menampilkan semua data
	fmt.Println(atks)

	// READ: Menampilkan satu data
	fmt.Println(atks[0])

	// READ: Menampilkan dengan loop
	for index, name := range atks {
		fmt.Println(index + 1, name)
	}

	// CREATE/ADD
	atks = append(atks, "Pensil")
	fmt.Println(atks)

	// UPDATE/EDIT
	atks[1] = "Buku Tulis"
	fmt.Println(atks)

	// DELETE
	// atks[:2] -> Mengambil elemen sebelum index 2, yaitu index 0 dan 1
	// atks[3:] -> Mengambil elemen mulai index 3 sampai akhir
	// ... -> Membuka elemen dari slice kedua agar dapat diteruskan sebagai argumen ke append()
	atks = append(atks[:2], atks[3:]...)
	fmt.Println(atks)
}