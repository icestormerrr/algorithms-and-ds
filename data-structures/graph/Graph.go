package graph

var VERY_BIG_INT = 1000000

type Graph struct {
	vertices map[int]*Vertex
}

func NewGraph() *Graph {
	return &Graph{vertices: make(map[int]*Vertex)}
}

func (this *Graph) Clear() {
	this.vertices = make(map[int]*Vertex)
}

func (this *Graph) Build(adjacencyMatrix [][]int) {
	this.Clear()
	this.initVertices(len(adjacencyMatrix))
	this.addEdgesToVertices(adjacencyMatrix)
}

func (this *Graph) initVertices(numberOfVertices int) {
	for vertexId := 0; vertexId < numberOfVertices; vertexId++ {
		this.vertices[vertexId] = NewVertex(vertexId, 0)
	}
}

func (this *Graph) addEdgesToVertices(adjacencyMatrix [][]int) {
	for vertexId, edgeCosts := range adjacencyMatrix {
		for neighborId, edgeCost := range edgeCosts {
			if vertexId == neighborId {
				continue
			}
			if edgeCost != -1 {
				var newEdge = NewEdge(vertexId, neighborId, edgeCost)
				this.vertices[vertexId].AddEdge(*newEdge)
			}
		}
	}
}

// FindCostsToVerticesFrom - Dijkstra's algorithm
func (this *Graph) FindCostsToVerticesFrom(fromId int) map[int]int {
	var checkedVerticesIds = this.createInitCheckedVerticesMap()
	var verticesCosts = this.createInitVerticesCostsMap(fromId)
	var currentVertex *Vertex

	for {
		currentVertex = this.chooseCurrentVertex(verticesCosts, checkedVerticesIds)
		if currentVertex == nil {
			return verticesCosts
		}

		this.updateVerticesCosts(currentVertex, verticesCosts)
		checkedVerticesIds[currentVertex.GetId()] = true
	}
}

func (this *Graph) createInitVerticesCostsMap(fromId int) map[int]int {
	var verticesCosts = make(map[int]int)

	for vertexId := range this.vertices {
		if vertexId == fromId {
			verticesCosts[vertexId] = 0
			continue
		}

		verticesCosts[vertexId] = VERY_BIG_INT
	}

	return verticesCosts
}

func (this *Graph) createInitCheckedVerticesMap() map[int]bool {
	var checkedVerticesIds = make(map[int]bool)

	for vertexId := range this.vertices {
		checkedVerticesIds[vertexId] = false
	}

	return checkedVerticesIds
}

func (this *Graph) chooseCurrentVertex(verticesCosts map[int]int, checkedVerticesIds map[int]bool) *Vertex {
	var vertex *Vertex = nil
	var lowestCost = VERY_BIG_INT

	for vertexId, vertexCost := range verticesCosts {
		if !checkedVerticesIds[vertexId] && vertexCost < lowestCost {
			lowestCost = vertexCost
			vertex = this.vertices[vertexId]
		}
	}

	if lowestCost == VERY_BIG_INT {
		return nil
	}

	return vertex
}

func (this *Graph) updateVerticesCosts(currentVertex *Vertex, verticesCosts map[int]int) {
	for _, edge := range currentVertex.GetEdges() {
		var costToVertex = edge.cost + verticesCosts[currentVertex.GetId()]
		if verticesCosts[edge.toId] > costToVertex {
			verticesCosts[edge.toId] = costToVertex
		}
	}
}

func (this *Graph) FindCostFromTo(fromId, toId int) int {
	var verticesCosts = this.FindCostsToVerticesFrom(fromId)
	return verticesCosts[toId]
}
