package vertex

import "math"

//顶点
type Vertex struct {
	X float64
	Y float64
}

//按比例缩放
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

//平方根
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
