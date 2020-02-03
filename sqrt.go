package main

import (
                "fmt"
)

func Sqrt(x float64) float64 {
    z := float64(1)
    var comp float64
    for i := 0; i < 1000; i++ {
        z = z- (z*z - x) / (2*z)
        comp = z*z-x
        if comp < 9e-15 {
            fmt.Println("Itterations:", i, "Deviation:" , comp)
            return z
        }
    }
    fmt.Println("Itterations: 1000,  Deviation:",z*z-x)
    return z

}

func main() {
    var a float64
    a=125
    fmt.Println("The square root of",a,"is",Sqrt(a))
}
