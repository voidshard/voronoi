package voronoi

import (
	"fmt"
	"image"
	"math"
)

// Site exposes useful functions of a voronoi diagram Site
type Site interface {
	ID() int
	X() int
	Y() int

	Edges() [][2]image.Point
	Vertices() []image.Point
	AllContains() PointGenerator
	Bounds() image.Rectangle
	Neighbours() []*Neighbour
}

// vSite is a wrapper around Voronoi & voronoiCell
type vSite struct {
	id     int
	parent *Voronoi
	cell   *voronoiCell
	edges  [][2]image.Point
	points []image.Point
	bounds image.Rectangle
}

// AllContains returns a PointGenerator for the given site
func (s *vSite) AllContains() PointGenerator {
	return &vPGen{
		parent: s.parent,
		me:     s,
		x:      -1,
		y:      -1,
	}
}

// Neighbours returns all Sites that share an edge with this site.
func (s *vSite) Neighbours() []*Neighbour {
	toEdgeId := func(e [2]image.Point) string {
		a, b := e[0], e[1]
		if b.X < a.X {
			a, b = b, a
		}
		if a.X == b.X && b.Y < a.Y {
			a, b = b, a
		}
		return fmt.Sprintf("%d.%d-%d.%d", a.X, a.Y, b.X, b.Y)
	}

	myedges := map[string][2]image.Point{}
	for _, e := range s.Edges() {
		myedges[toEdgeId(e)] = e
	}

	ls := []*Neighbour{}
	for _, syte := range s.parent.Sites() {
		if syte.ID() == s.ID() {
			continue
		}

		n := &Neighbour{Site: syte, Edges: [][2]image.Point{}}
		for _, e := range syte.Edges() {
			eid := toEdgeId(e)
			_, ok := myedges[eid]
			if !ok {
				continue
			}
			n.Edges = append(n.Edges, e)
		}
		if len(n.Edges) > 0 {
			ls = append(ls, n)
		}
	}

	return ls
}

// ID of this site
func (s *vSite) ID() int {
	return s.id
}

// X value of site centre
func (s *vSite) X() int {
	return int(s.cell.Center.X)
}

// Y value of site centre
func (s *vSite) Y() int {
	return int(s.cell.Center.Y)
}

// build deduces all points / edges
func (s *vSite) build() {
	s.points = []image.Point{}
	s.edges = [][2]image.Point{}

	minX := -1
	maxX := -1
	minY := -1
	maxY := -1

	min := func(a, b int) int {
		if a == -1 || b < a {
			return b
		}
		return a
	}

	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}

	check := func(p image.Point) {
		minX = min(minX, p.X)
		minY = min(minY, p.Y)
		maxX = max(maxX, p.X)
		maxY = max(maxY, p.Y)
	}

	for _, edge := range s.cell.Edges {
		start := image.Pt(int(math.Round(edge[0].X)), int(math.Round(edge[0].Y)))
		end := image.Pt(int(math.Round(edge[1].X)), int(math.Round(edge[1].Y)))

		s.points = append(s.points, start)
		s.edges = append(s.edges, [2]image.Point{start, end})

		check(start)
		check(end)
	}

	s.bounds = image.Rect(minX, minY, maxX, maxY)
}

// Edges returns all edges surrounding this site
func (s *vSite) Edges() [][2]image.Point {
	if s.points == nil {
		s.build()
	}
	return s.edges
}

// Vertices returns all vertexes (through which edges pass) of the site
func (s *vSite) Vertices() []image.Point {
	if s.points == nil {
		s.build()
	}
	return s.points
}

// Bounds returns a rectangle that necessarily contains all points in the site
// and more besides.
func (s *vSite) Bounds() image.Rectangle {
	if s.points == nil {
		s.build()
	}
	return s.bounds
}
