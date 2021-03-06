// MIT License: See https://github.com/voidshard/voronoi/LICENSE.md

// Author: Przemyslaw Szczepaniak (przeszczep@gmail.com)
// Port of Raymond Hill's (rhill@raymondhill.net) javascript implementation
// of Steven Forune's algorithm to compute Voronoi diagrams

package voronoi

import (
	"math/rand"
	"testing"
)

func verifyDiagram(diagram *Diagram, edgesCount, cellsCount, perCellCount int, t *testing.T) {
	if len(diagram.Edges) != edgesCount {
		t.Errorf("Expected %d edges not %d", edgesCount, len(diagram.Edges))
	}

	if len(diagram.Cells) != cellsCount {
		t.Errorf("Expected %d cells not %d", cellsCount, len(diagram.Cells))
	}

	if perCellCount > 0 {
		for _, cell := range diagram.Cells {
			if len(cell.Halfedges) != perCellCount {
				t.Errorf("Expected per cell edge count expected %d, not %d", perCellCount, len(cell.Halfedges))
			}
		}
	}
}

func TestHorizontal(t *testing.T) {
	sites := make([]Vertex, 0)
	for i := 0; i < 100; i++ {
		sites = append(sites, Vertex{float64(i), 1})
	}
	verifyDiagram(ComputeDiagram(sites, NewBBox(0, 100, 0, 100), true),
		301, 100, 4, t)
}

func TestVertical(t *testing.T) {
	sites := make([]Vertex, 0)
	for i := 0; i < 100; i++ {
		sites = append(sites, Vertex{1, float64(i)})
	}
	verifyDiagram(ComputeDiagram(sites, NewBBox(0, 100, 0, 100), true),
		301, 100, 4, t)
}

func TestSquare(t *testing.T) {
	sites := make([]Vertex, 0)
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			sites = append(sites, Vertex{float64(i), float64(j)})
		}
	}
	ComputeDiagram(sites, NewBBox(0, 10, 0, 10), true)
}

func TestRegressionSetEdgeStartPointPanic(t *testing.T) {
	//	There's a weird edge case where an edge isn't set so a panic occurs.
	//	Unsure exactly why .. but this magic set of Verts triggers it ..
	sites := []Vertex{
		Vertex{210, 551},
		Vertex{87, 580},
		Vertex{471, 770},
		Vertex{519, 0},
		Vertex{659, 329},
		Vertex{569, 441},
		Vertex{575, 524},
		Vertex{242, 564},
		Vertex{569, 0},
		Vertex{305, 317},
		Vertex{417, 334},
		Vertex{586, 391},
	}

	diag := ComputeDiagram(sites, NewBBox(0, 1000, 0, 1000), true)
	verifyDiagram(diag, 17, 5, 0, t)

	diag = ComputeDiagram(sites, NewBBox(0, 1000, 0, 1000), false)
	verifyDiagram(diag, 4, 5, 0, t)
}

func TestVoronoi2Points(t *testing.T) {
	sites := []Vertex{
		Vertex{4, 5},
		Vertex{6, 5},
	}

	verifyDiagram(ComputeDiagram(sites, NewBBox(0, 10, 0, 10), true),
		7, 2, 4, t)
	verifyDiagram(ComputeDiagram(sites, NewBBox(0, 10, 0, 10), false),
		1, 2, 1, t)
}

func TestVoronoi3Points(t *testing.T) {
	sites := []Vertex{
		Vertex{4, 5},
		Vertex{6, 5},
		Vertex{5, 8},
	}

	verifyDiagram(ComputeDiagram(sites, NewBBox(0, 10, 0, 10), true),
		10, 3, -1, t)
	verifyDiagram(ComputeDiagram(sites, NewBBox(0, 10, 0, 10), false),
		3, 3, 2, t)
}

func Benchmark1000(b *testing.B) {
	rand.Seed(1234567)
	b.StopTimer()
	sites := make([]Vertex, 100)
	for j := 0; j < 100; j++ {
		sites[j].X = rand.Float64() * 100
		sites[j].Y = rand.Float64() * 100
	}
	b.StartTimer()
	ComputeDiagram(sites, NewBBox(0, 100, 0, 100), true)
}
