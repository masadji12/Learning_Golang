package main

import "fmt"

func main() {
	tinggi := 5

	one(tinggi)
	two(tinggi)
	three(tinggi)
	four(tinggi)
	// fmt.Println("1. SETENGAH SEGITIGA:")
	// cetakSetengahSegitiga(tinggi)
	// fmt.Println()

	// fmt.Println("2. SEGITIGA TERBALIK (SETENGAH):")
	// cetakSetengahTerbalik(tinggi)
	// fmt.Println()

	// fmt.Println("3. SEGITIGA SAMA KAKI UTUH:")
	// cetakSegitigaUtuh(tinggi)
	// fmt.Println()
	
	// fmt.Println("4. SEGITIGA SAMA KAKI TERBALIK UTUH:")
	// cetakSegitigaUtuhTerbalik(tinggi)
}

// 1. Setengah Segitiga
func cetakSetengahSegitiga(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

// 2. Setengah Segitiga Terbalik
func cetakSetengahTerbalik(n int) {
	for i := n; i >= 1; i-- {
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

// 3. Segitiga Sama Kaki Utuh
func cetakSegitigaUtuh(n int) {
	for i := 1; i <= n; i++ {
		// Cetak spasi di kiri agar membentuk gunungan
		for j := 1; j <= n-i; j++ {
			fmt.Print(" ")
		}
		// Cetak bintang (rumus jumlah bintang: 2*i - 1)
		for k := 1; k <= (2*i - 1); k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

// 4. Segitiga Sama Kaki Terbalik Utuh
func cetakSegitigaUtuhTerbalik(n int) {
	for i := n; i >= 1; i-- {
		// Cetak spasi di kiri
		for j := 1; j <= n-i; j++ {
			fmt.Print(" ")
		}
		// Cetak bintang
		for k := 1; k <= (2*i - 1); k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func one(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func two(n int) {
	for i := n; i >= 1; i-- {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func three(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= n - i; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= (2 * i - 1); k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func four(n int) {
	for i := n - 1; i >= 1; i-- {
		for j := 1; j <= n - i; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= (2 * i -1); k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}