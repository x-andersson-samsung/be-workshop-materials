package exercise7

// Write a function called GetArea accepting a Shape interface and returning the area.
type SSolution_hape interface {
	GetArea() float64
}

type Solution_Rectangle struct {
	Width  float64
	Height float64
}

func (r Solution_Rectangle) Solution_GetArea() float64 {
	return r.Height * r.Width
}
