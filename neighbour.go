package voronoi

import (
	"image"
)

// Neighbour is a Site & Edge that is shared by another cell.
// See Neighbours()
type Neighbour struct {
	Site  Site
	Edges [][2]image.Point
}
