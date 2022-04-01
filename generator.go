package voronoi

import (
	"image"
)

// PointGenerator returns all points within a Site, in a somewhat
// sane fashion that doesn't involve pre-computing a huge set
type PointGenerator interface {
	Next() *image.Point
}

// vPGen satisties PointGenerator
type vPGen struct {
	parent *Voronoi
	me     Site
	x      int
	y      int
}

// Next return the next point contained in a Site.
// A nil value indicates that there are no more.
func (v *vPGen) Next() *image.Point {
	if v.x == -1 && v.y == -1 {
		v.x = v.bounds.Min.X
		v.y = v.bounds.Min.Y
	}

	bounds := v.me.Bounds()
	for ; v.y < bounds.Max.Y; v.y++ {
		for x := v.x; x < bounds.Max.X; x++ {
			s := v.parent.SiteFor(x, v.y)
			if s.ID() == v.me.ID() {
				v.x = x + 1
				return &p
			}
		}
		v.x = bounds.Min.X
	}

	return nil
}