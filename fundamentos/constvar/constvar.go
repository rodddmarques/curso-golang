package main

import (
	"fmt"
	"math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 // tipo (float64) inferido pelo compilador

	//forma reduzida de criar um var
	area := PI * math.Pow(raio, 2)
	fmt.Printf("A área da circunferência é %f m2 \n", area)

	var e, f bool = true, false
	fmt.Println(e, f)

	g, h, i := 2, false, "epa!!"
	fmt.Println(g, h, i)
}
