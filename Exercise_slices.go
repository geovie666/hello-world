// A Tour of Go 18/27
package main

import (
	"golang.org/x/tour/pic"
	"math"
)

func Pic(dx, dy int) [][]uint8 {

	a := make([][]uint8, dy)
	for i:=0;i<dy;i++{
		a[i] = make([]uint8, dx)
	}

	for i:=0;i<dx;i++{
		for j:=0;j<dy;j++{
			c:=float64(math.Pow(float64(dx/2-i),2)+math.Pow(float64(j-dy/2),2))
			if math.Sqrt(c)<80{
				a[i][j]=0
			} else {
				a[i][j]=180
			}
		}
	}

	return a
}

func main() {
	pic.Show(Pic)
}
