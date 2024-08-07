package graph

type Vertex struct {
	id    int
	value int
	edges []Edge
}

func NewVertex(id int, value int) *Vertex {
	return &Vertex{id, value, []Edge{}}
}

func (this *Vertex) AddEdge(edge Edge) {
	this.edges = append(this.edges, edge)
}

func (this *Vertex) GetEdges() []Edge {
	return this.edges
}

func (this *Vertex) GetId() int {
	return this.id
}
