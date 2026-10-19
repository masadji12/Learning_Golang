// cd in directory modules
// go mod init <module> (Read Dokumentation go.mod.txt)
// create file main.go
// git run . or git run <spesification_file>
// git build (build file compiler)

package main

import "fmt"

func main() {

	// fmt.Println("Hello") Menampilkan nilai dan mengakhiri output dengan baris baru
	// fmt.Print() Menampilkan nilai without otomatis menambahkan baris baru
	// fmt.Printf() Menampilkan teks menggunakan format tertentu $s, $d, dan %f
	fmt.Println("Hello, Golang!")
	fmt.Println("I current studying beckend development")
	fmt.Println("My first project successfull running")
}

// Statement = Intruksi yang melakukan suatu tindakan
// Expression = bagian kode yang menghasilkan suat nilai
// Komentar = memberikan penjelasan kode

// Code formatting menggunakan gofmt (gofmt -w main.go)
