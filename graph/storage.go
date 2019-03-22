package DataStructure

// Adjacent storage: note that at general this will cause sparse matrix to take up too many space, a tuple of 3 elements could be used
type MVertex struct {
	id   int
	info string
}

type MGraph struct {
	// representing the weight or cost from node a to node b
	edges            [][]float32
	vertexNo, edgeNo int
	vertices         []MVertex
}

// Adjacent table: each element is a head pointing to the single linked list where every node inside the list is the corresponding
// node next to the head
type ALVertex struct {
	id   int
	data string
	next *ALVertex
	cost float32
}

type ALGraph struct {
	adjlist      []ALVertex
	nodes, edges int
}

// conversion
func (g *ALGraph) ALGraph2MGraph() MGraph {
	// init
	edges := make([][]float32, g.nodes)
	for i := range edges {
		edges[i] = make([]float32, g.nodes)
	}
	vertices := make([]MVertex, g.nodes)
	edgesNo := 0

	for index := range g.adjlist {
		next := g.adjlist[index].next
		if next != nil {
			edgesNo++
		}
		for next != nil {
			nextIndex := next.id
			weight := next.cost
			edges[index][nextIndex] = weight
			next = next.next
		}
		vertices[index] = MVertex{index, g.adjlist[index].data}
	}
	return MGraph{edges, g.nodes, edgesNo, vertices}
}

// func (g *MGraph) MGraph2ALGraph() ALGraph {

// }
