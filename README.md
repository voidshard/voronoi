## Voronoi


Simple library to build Voronoi diagrams from random points, chosen points or a mixture of both.
```golang
// prepare to make a new graph
build := NewBuilder(image.Rect(0, 0, 1000, 1000))

// add two sites by hand
build.AddSite(100, 150)
build.AddSite(800, 700)

// a filter so the builder will not add points that are too close together
build.SetSiteFilters(voronoi.MinDistance(100)) 

// add ~10 random sites (respecting the above filters)
for i := 0; i < 10; i++ {
        build.AddRandomSite()
}

// compute the voronoi diagram
graph, err := build.Voronoi()

// ... profit
```

There are also a few extra features; calculating the external edges of sets of sites, finding all points between two points etc.
Fixes / features will be added as I find and/or need them.

There are a few voronoi libraries out there but some occasionally panic (?!?), others didn't supply all of the functionality I wanted, others made it hard to add a mixture of random & chosen points .. so I figured I'd break my own lib out of another project for ease of reuse.
