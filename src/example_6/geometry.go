package example_6

import (
	"fmt"
	"image/color"
	"math"
)

type Point struct {
	X, Y float64
}

type ColorPoint struct {
	Point
	Color color.RGBA
}

type Path []Point

func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p Path) Distance() float64 {
	sum := 0.0

	for i, _ := range p {
		if i > 0 {
			sum += p[i-1].Distance(p[i])
		}
	}
	return sum
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func TestDistance() {

	p := Point{1, 2}
	q := Point{4, 6}

	fmt.Println(Distance(p, q))
	fmt.Println(p.Distance(q))

	perim := Path{{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println(perim.Distance())

	r := &p
	//测试指针对象的方法
	p.ScaleBy(2)
	r.ScaleBy(2)
	Point{3, 5}.Distance(q)
	fmt.Println(p)

	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}

	colorP := ColorPoint{Point: Point{1, 1}, Color: red}
	colorQ := ColorPoint{Point: Point{5, 4}, Color: blue}

	fmt.Println(colorP.Distance(colorQ.Point))
	colorP.ScaleBy(2)
	colorQ.ScaleBy(2)

	fmt.Println(colorP.Distance(colorQ.Point))

}
