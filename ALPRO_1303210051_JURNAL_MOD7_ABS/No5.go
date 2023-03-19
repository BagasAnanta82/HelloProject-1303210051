package main

import "fmt"

func main() {
	var nama, nama_min, nama_max string
	var j, p int
	var faktor, jumlah, min, max, mean float64
	nama_min = ""
	nama_max = ""
	fmt.Scan(&nama)
	for nama != "STOP" {
		fmt.Scan(&p)
		jumlah = 0
		for j = 0; j < 3*p; j++ {
			fmt.Scan(&faktor)
			jumlah += faktor
		}
		mean = jumlah / float64(3*p)
		if nama_min == "" || nama_min != "" && min > mean {
			min = mean
			nama_min = nama
		}
		if nama_max == "" || nama_max != "" && max < mean {
			max = mean
			nama_max = nama
		}
		fmt.Scan(&nama)
	}
	fmt.Printf("%s %.2f\n", nama_max, max)
	fmt.Printf("%s %.2f\n", nama_min, min)
}
