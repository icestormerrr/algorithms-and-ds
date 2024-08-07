package graph

type Edge struct {
	fromId int
	toId   int
	cost   int
}

func NewEdge(fromNodeId int, toNodeId int, cost int) *Edge {
	return &Edge{fromNodeId, toNodeId, cost}
}
