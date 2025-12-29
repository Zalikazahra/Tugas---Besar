package main

import "fmt"

func analisisLinear(nilai []int) {
	for i := 1; i < len(nilai); i++ {
		if nilai[i] > nilai[i-1] {
			fmt.Println("Nilai ke-", i+1, "mengalami kenaikan")
		} else if nilai[i] < nilai[i-1] {
			fmt.Println("Nilai ke-", i+1, "mengalami penurunan")
		} else {
			fmt.Println("Nilai ke-", i+1, "tetap")
		}
	}
}

func main() {
	nilaiMahasiswa := []int{70, 75, 78, 80, 82, 85, 88}
	analisisLinear(nilaiMahasiswa)
}
