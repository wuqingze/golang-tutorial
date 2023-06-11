package main

import (
    "fmt"
    "os"
    "strconv"
)

var CI = []float64{-0.0000000001, 0.0000000001}

func Sqrt(x float64) float64 {
    xn := x/2
    for xn*xn-x>CI[1] || xn*xn-x<CI[0]{
        fmt.Printf("xn is %f\n", xn)
        xn = (xn+x/xn)/2
    }
    return xn
}

func main() {
    if len(os.Args) < 2{
        fmt.Println("please provide a integer argument")
    }
    x, err := strconv.Atoi(os.Args[1])
	if err == nil {
        fmt.Println(Sqrt(float64(x)))
    }
}

