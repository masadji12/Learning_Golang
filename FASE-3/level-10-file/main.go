package main

import (
	"fmt"
	"os"
)

func main() {
	// []byte("...") => Mengubah teks menjadi kumpulan byte
	data := []byte("Pulpen: 100 pcs\n")

	// os.WriteFile() => Menulis data ke file; membuat file jika belum ada dan menimpa isinya jika sudah ada.
	// "stok.txt" => Nama atau path file tujuan
	// 0064 => Permission mode file pada sistem yang mendukung permission Unix
	// err != nil => Memeriksa apakah operasi gagal
	err := os.WriteFile("stock.txt", data, 0064)
	if err != nil {
		fmt.Println("Failed to writing file:", err)
		return
	}

	content, err := os.ReadFile("stock.txt")
	if err != nil {
		fmt.Println("Failed reading file:", err)
		return
	}

	fmt.Println("File Contents:")
	fmt.Println(string(content))
}