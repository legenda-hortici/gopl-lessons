package main

import (
	"fmt"
	temp "gopl/ch2/tempconv"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}

		f, c, k := temp.Fahrenheit(t), temp.Celsius(t), temp.Kelvin(t)
		g, kg := temp.Gram(t), temp.Kilogram(t)
		mill, cm := temp.Millimetr(t), temp.Centimetr(t)

		fmt.Printf("%s = %s, %s = %s, %s = %s\n", f, temp.FToC(f), c, temp.CToF(c), k, temp.KToC(k))
		fmt.Printf("%s = %s, %s = %s, %s = %s\n", g, temp.GrToKg(g), kg, temp.KgToT(kg), g, temp.GrToT(g))
		fmt.Printf("%s = %s, %s = %s, %s = %s\n", mill, temp.MillToCm(mill), cm, temp.CmToM(cm), mill, temp.MillToM(mill))
	}
}
