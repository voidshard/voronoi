# Voronoi diagrams in Go

A Implementation of of Steven J. Fortune's algorithm to efficiently compute Voronoi diagrams in Go language. Based on a Raymond Hill's javascript implementation (https://raw.github.com/gorhill/Javascript-Voronoi).

Forked from https://github.com/pzsz/voronoi -- mostly because I wanted to push in some fixes and the original repo seems a bit dead.


## Usage

```
import "github.com/voidshard/voronoi"

func useVoronoi() {
     	// Sites of voronoi diagram
	sites := []voronoi.Vertex{
		voronoi.Vertex{4, 5},
		voronoi.Vertex{6, 5},
		...
	}

	// Create bounding box of [0, 20] in X axis
	// and [0, 10] in Y axis
	bbox := voronoi.NewBBox(0, 20, 0, 10)

	// Compute diagram and close cells (add half edges from bounding box)
	diagram := voronoi.ComputeDiagram(sites, bbox, true)

	// Iterate over cells
	for _, cell := diagram.Cells {
		for _, hedge := cell.Halfedges {
		    ...
		}	
	}

	// Iterate over all edges
	for _, edge := diagram.Edge {
	    ...
	}
}
```


## Fixes

- fixed documentation
- added go mod files
- fixed null pointer panic when certain vertex conditions met. +tests
- fixed potential infinite loop bug when certain vertex conditions met. +regression tests

