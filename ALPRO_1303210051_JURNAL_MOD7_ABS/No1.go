package main

import "fmt"

func main() {
	var n, a, b, positif, negatif int
	positif = 0
	negatif = 0
	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&a, &b)
		if a > 0 {
			positif += a
		} else {
			negatif += a
		}
		if b > 0 {
			positif += b
		} else {
			negatif += b
		}
	}
	fmt.Println("negatif :", negatif)
	fmt.Println("positif :", positif)
}
