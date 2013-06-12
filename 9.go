package main

import "fmt"
import "math"

func main() {
	for a:= 1; ; a++ {
		for b := a + 1; b <= 1000; b++ {
			temp := math.Sqrt(float64(a*a+b*b))
			if math.Floor(temp) == temp {
				c := int(temp)
				if a + b + c == 1000 && a < b && b < c {
					fmt.Printf("a: %d\nb: %d\nc: %d\nabc: %d\n", a, b, c, a * b * c)
					return
				}
			}
		}
	}
}
