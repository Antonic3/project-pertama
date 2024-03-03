package main

import (
	"fmt"
	"os"
)

type Teman struct {
	Absen     int
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

var dataTeman = []Teman{
	{1, "Avant", "Jl. Melati no 56", "Web Developer", "Ingin mencoba hal baru dalam programing"},
	{2, "Azis", "Jl. Sawah Baru No. 33", "Mobile Developer", "Ingin menambah skill dalam programing"},
	{3, "Ali", "Jl. Kramat batu No. 66", "Mahasiswa", "Mencari ilmu baru buat unlock new skills"},
	{4, "Aldian", "Jl. Batu Lampa No. 88", "Mahasiswa", "Mencari skill baru untuk kerja nanti"},
	{5, "Christian", "Jl. Damai No. 46", "Web Developer", "Menambah skill dalam programing"},
}

// Fungsi untuk menampilkan data teman berdasarkan nomor absen
func tampilkanData(absen int) {
	for _, teman := range dataTeman {
		if teman.Absen == absen {
			fmt.Println("Data Teman dengan Absen", absen)
			fmt.Println("Nama:", teman.Nama)
			fmt.Println("Alamat:", teman.Alamat)
			fmt.Println("Pekerjaan:", teman.Pekerjaan)
			fmt.Println("Alasan ingin belajar pemrograman GOLANG:", teman.Alasan)
			return
		}
	}
	fmt.Println("Tidak ditemukan data teman dengan absen", absen)
}

func main() {
	// Cek apakah argumen sudah disediakan
	if len(os.Args) < 2 {
		fmt.Println("Gunakan: go run main.go [nomor absen]")
		return
	}

	// Dapatkan nomor absen dari argumen
	absen := os.Args[1]

	// Konversi nomor absen dari string ke integer
	var absenInt int
	_, err := fmt.Sscanf(absen, "%d", &absenInt)
	if err != nil {
		fmt.Println("Nomor absen harus berupa angka")
		return
	}

	tampilkanData(absenInt)
}
