package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("stok.txt")
	if err != nil {
		fmt.Println("Failed open file:", err)
		return
	}

	defer file.Close()
}

// Setelah file berhasil dibuka, defer file.Close() menjadwalkan penutupan file saat fungsi yang memuatnya selesai.
// Mengapa ditulis setelah pemeriksaan error? karena jika os.Open() failed, file tidak berhasil dibuka dan kita tidak boleh melanjutkan seolah olah file tersedia
// Catatan penting: argumen fungsi pada defer umumnya dievaluasi saat perintah defer dijalankan, bukan ketika fungsi benar-benar dipanggil di akhir. Ini bisa berpengaruh jika nilai variabel berubah setelahnya.