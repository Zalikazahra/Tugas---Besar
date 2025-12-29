package main

import "fmt"

func analisisDivideConquer(nilai []int, kiri int, kanan int) {
	if kiri >= kanan {
		return
	}

	tengah := (kiri + kanan) / 2

	analisisDivideConquer(nilai, kiri, tengah)
	analisisDivideConquer(nilai, tengah+1, kanan)

	if nilai[tengah] < nilai[tengah+1] {
		fmt.Println("Kenaikan nilai antara indeks", tengah, "dan", tengah+1)
	} else if nilai[tengah] > nilai[tengah+1] {
		fmt.Println("Penurunan nilai antara indeks", tengah, "dan", tengah+1)
	} else {
		fmt.Println("Nilai tetap antara indeks", tengah, "dan", tengah+1)
	}
}

func main() {
	nilaiMahasiswa := []int{70, 75, 78, 80, 82, 85, 88}
	analisisDivideConquer(nilaiMahasiswa, 0, len(nilaiMahasiswa)-1)
}
