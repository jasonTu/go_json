package coloredpoint

import (
	"image/color"
	"go_json/gopl/ch6/geometry"
)

type ColoredPoint struct {
	geometry.Point
	Color color.RGBA
}
