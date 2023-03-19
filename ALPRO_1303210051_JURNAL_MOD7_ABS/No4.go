package main

import "fmt"

func main() {
	var j, pengunjung, sebelum, penurunan, jumlah int
	penurunan = 0
	sebelum = 0
	j = 0
	jumlah = 0
	for penurunan != 3 {
		j++
		fmt.Print("Hari ", j, " : ")
		fmt.Scan(&pengunjung)
		for pengunjung < 0 || pengunjung > 100 {
			fmt.Print("Hari ", j, " : ")
			fmt.Scan(&pengunjung)
		}
		jumlah += pengunjung
		if pengunjung < sebelum {
			penurunan++
		} else {
			penurunan = 0
		}
		sebelum = pengunjung
	}
	fmt.Println("Museum buka selama", j, "hari")
	fmt.Printf("Rata-rata pengunjung %.2f orang", float64(jumlah)/float64(j))
}
