package main

import "fmt"

func main() {
	var member string
	var a, b, c, d, e int
	var cashback, diskon float64

	fmt.Scan(&member, &a, &b, &c, &d, &e)

	diskon = 0
	cashback = 0
	if a%2 == 0 && b%2 == 0 && c%2 == 0 && d%2 == 0 && e%2 == 0 {
		cashback = 1.7 * float64(c+d+e)
	} else if a%2 != 0 && b%2 != 0 && c%2 != 0 && d%2 != 0 && e%2 != 0 {
		diskon = 1.7 * float64(c+d+e)
	} else {
		diskon = 1.7 * float64(c+d+e)
		cashback = 3.1 * float64(a+b+c)
	}

	if member == "yes" {
		diskon *= 1.15
		cashback *= 1.15
	}

	if cashback > 35 {
		cashback = 35
	}
	if diskon > 50 {
		diskon = 50
	}
	fmt.Print(cashback, diskon)
}
