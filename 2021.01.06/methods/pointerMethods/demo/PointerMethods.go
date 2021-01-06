//用于试验基于指针对象方法的使用，学习了解
package main

import (
	"fmt"
	"math"
)

type Point struct{ X, Y float64 }

func main() {
	r := Point{1,1}
	//编译器隐式的用&r调用了Scale方法
	r.Scale(2)
	fmt.Println(r)
	//真实情况
	(&r).Scale(3)
	fmt.Println(r)

	p := Point{1, 2}
	q := Point{4, 6}
	//p.Distance 会返回一个方法"值"
	distanceFromP := p.Distance        // 方法"值"
	//distanceFromP
	fmt.Println(distanceFromP(q))      // "5"
	var origin Point                   // {0, 0}
	fmt.Println(distanceFromP(origin)) //p（1，2） origin(0,0)  sqrt(5)= "2.23606797749979"

	scaleP := p.ScaleBy // method value
	scaleP(2)           // p becomes (2, 4)

}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func (p *Point) Scale(factor float64)  {
	p.X *= factor
	p.Y *= factor
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}