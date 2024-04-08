package main

type position struct {
	x int
	y int
}

func (pos position) isInside(topLeft position, bottomRight position) bool {
	return pos.x >= topLeft.x && pos.x <= bottomRight.x && pos.y >= topLeft.y && pos.y <= bottomRight.y
}
