package main

import "fmt"

func main() {
	var hari1, bulan1, hari2, bulan2 string
	var tanggal1, tanggal2, angkaBulan1, angkaBulan2, tahun1, tahun2 int

	fmt.Scan(&hari1, &tanggal1, &bulan1, &tahun1)
	angkaBulan1 = angkaBulan(bulan1)
	angkaBulan2 = angkaBulan(bulan2)

	pengambilan(tanggal1, angkaBulan1, tahun1, hari1, &tanggal2, &angkaBulan2, &tahun2, &hari2)

	bulan2 = bulanAngka(angkaBulan2)
	fmt.Println(hari2, tanggal2, bulan2, tahun2)

}

func kabisat(tahun int) bool {
	var th bool
	if tahun%4 == 0 || tahun%400 == 0 && tahun%100 != 0 {
		th = true
	} else {
		th = false
	}
	return th
}

func angkaBulan(bulan string) int {
	var a int
	if bulan == "januari" {
		a = 1
	} else if bulan == "februari" {
		a = 2
	} else if bulan == "maret" {
		a = 3
	} else if bulan == "april" {
		a = 4
	} else if bulan == "mei" {
		a = 5
	} else if bulan == "juni" {
		a = 6
	} else if bulan == "juli" {
		a = 7
	} else if bulan == "agustus" {
		a = 8
	} else if bulan == "september" {
		a = 9
	} else if bulan == "oktober" {
		a = 10
	} else if bulan == "november" {
		a = 11
	} else if bulan == "desember" {
		a = 12
	} else {
		a = 0
	}
	return a
}

func bulanAngka(angka int) string {
	var st string
	if angka == 1 {
		st = "januari"
	} else if angka == 2 {
		st = "februari"
	} else if angka == 3 {
		st = "maret"
	} else if angka == 4 {
		st = "april"
	} else if angka == 5 {
		st = "mei"
	} else if angka == 6 {
		st = "juni"
	} else if angka == 7 {
		st = "juli"
	} else if angka == 8 {
		st = "agustus"
	} else if angka == 9 {
		st = "september"
	} else if angka == 10 {
		st = "oktober"
	} else if angka == 11 {
		st = "november"
	} else if angka == 12 {
		st = "desember"
	} else {
		st = "invalid"
	}
	return st
}

func jumlahHari(bulan, tahun int) int {
	var h int
	if bulan == 1 {
		h = 31
	} else if bulan == 2 && kabisat(tahun) {
		h = 29
	} else if bulan == 2 {
		h = 28
	} else if bulan == 3 {
		h = 31
	} else if bulan == 4 {
		h = 30
	} else if bulan == 5 {
		h = 31
	} else if bulan == 6 {
		h = 30
	} else if bulan == 7 {
		h = 31
	} else if bulan == 8 {
		h = 31
	} else if bulan == 9 {
		h = 30
	} else if bulan == 10 {
		h = 31
	} else if bulan == 11 {
		h = 30
	} else if bulan == 12 {
		h = 31
	} else {
		h = -1
	}
	return h
}

func mencariDurasi(hari1 string, hari2 *string, durasi *int) {
	if hari1 == "senin" {
		*durasi = 2
		*hari2 = "rabu"
	} else if hari1 == "selasa" {
		*durasi = 2
		*hari2 = "kamis"
	} else if hari1 == "rabu" {
		*durasi = 2
		*hari2 = "jumat"
	} else if hari1 == "kamis" {
		*durasi = 4
		*hari2 = "senin"
	} else if hari1 == "jumat" {
		*durasi = 4
		*hari2 = "selasa"
	}
}

func pengambilan(tanggal1, bulan1, tahun1 int, hari1 string, tanggal2, bulan2, tahun2 *int, hari2 *string) {
	var durasi, jumlah int
	mencariDurasi(hari1, hari2, &durasi)
	*tanggal2 = tanggal1 + durasi
	*bulan2 = bulan1
	*tahun2 = tahun1

	jumlah = jumlahHari(bulan1, tahun1)
	if *tanggal2 > jumlah {
		*tanggal2 %= jumlah
		*bulan2++
	}

	if *bulan2 > 12 {
		*bulan2 %= 12
		*tahun2++
	}
}
