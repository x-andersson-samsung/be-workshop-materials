package exercise7

// Write a function called GetArea accepting a Shape interface and returning the area.
type Shape interface {
	GetArea() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) GetArea() float64 {
	return r.Height * r.Width
}
