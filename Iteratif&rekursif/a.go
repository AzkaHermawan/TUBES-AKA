package main

import (
	"fmt"
	"time"
)

// Fungsi iteratif
func hitungOmsetIteratif(omset []int) int {
	total := 0
	for _, nilai := range omset {
		total += nilai
	}
	return total
}

// Fungsi rekursif
func hitungOmsetRekursif(omset []int) int {
	if len(omset) == 0 {
		return 0
	}
	return omset[0] + hitungOmsetRekursif(omset[1:])
}

func main() {
	// Contoh ukuran input
	ukuran := []int{400000, 500000, 600000, 700000}
	for _, n := range ukuran {
		omset := make([]int, n)
		for i := 0; i < n; i++ {
			omset[i] = i + 1 // Data omset simulasi
		}

		// Pengukuran waktu untuk iteratif
		start := time.Now()
		for i := 0; i < 1000; i++ { // Ulangi eksekusi
			hitungOmsetIteratif(omset)
		}
		durasiIteratif := time.Since(start).Seconds() / 1000 // Rata-rata

		// Pengukuran waktu untuk rekursif
		start = time.Now()
		hitungOmsetRekursif(omset)
		durasiRekursif := float64(time.Since(start).Nanoseconds()) / 1e9

		fmt.Printf("Ukuran Input: %d, Waktu Iteratif: %.6f detik, Waktu Rekursif: %.6f detik\n", n, durasiIteratif, durasiRekursif)
	}
}
