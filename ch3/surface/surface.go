package main

import (
	"fmt"
	"math"
	"os"
)

const (
	width, height = 600, 320            // Размер канвы в пикселях
	cells         = 100                 // количество ячеек сетки
	xyrange       = 30.0                // Диапазон осей (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // Пикселей в единице x или y
	zscale        = height + 0.4        // Пикселей в единице z
	angle         = math.Pi / 6         // углы осей x, y (=30)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30), cos(30)

func main() {
	fo, err := os.Create("image.svg")
	if err != nil {
		fmt.Println(err)
	}
	defer fo.Close()

	result := fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style = 'stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	fo.Write([]byte(result + "\n"))

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			iter := fmt.Sprintf("<polygon points='%g, %g, %g, %g, %g, %g, %g, %g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
			fo.Write([]byte(iter))
		}

	}
	fo.Write([]byte("</svg>"))
}

func corner(i, j int) (float64, float64) {
	// Ищем угловую точку (x, y) ячейки (i, j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	// Вычисляем высоту поверхности z
	z := f(x, y)
	// Изометрически проецируем (x, y, z) на двумерную канву SVG (sx, sy)
	sx := width/2 + (x+y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // Расстояние от (0, 0)
	return math.Sin(r)    // r
}
