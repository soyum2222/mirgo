package common

import (
	"fmt"
)

type Point struct {
	X uint32
	Y uint32
}

func (p Point) String() string {
	return fmt.Sprintf("%d,%d", p.X, p.Y)
}

func (p Point) NextPoint(direction MirDirection, step uint32) Point {
	x := p.X
	y := p.Y
	switch direction {
	case MirDirectionUp:
		y = y - step
	case MirDirectionUpRight:
		x = x + step
		y = y - step
	case MirDirectionRight:
		x = x + step
	case MirDirectionDownRight:
		x = x + step
		y = y + step
	case MirDirectionDown:
		y = y + step
	case MirDirectionDownLeft:
		x = x - step
		y = y + step
	case MirDirectionLeft:
		x = x - step
	case MirDirectionUpLeft:
		x = x - step
		y = y - step
	}
	return Point{X: x, Y: y}
}

func NewPoint(x, y int) Point {
	return Point{uint32(x), uint32(y)}
}

func (p Point) Equal(other Point) bool {
	return p.X == other.X && p.Y == other.Y
}

type SelectInfo struct {
	Index      uint32
	Name       string
	Level      uint16
	Class      MirClass
	Gender     MirGender
	LastAccess int64
}

type Color struct {
	R uint8
	G uint8
	B uint8
}

func (c Color) ToInt32() int32 {
	return int32(c.ToUint32())
}

func (c Color) ToUint32() uint32 {
	return BytesToUint32([]uint8{c.R, c.G, c.B, 255})
}
