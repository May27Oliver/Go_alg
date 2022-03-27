package main

import "fmt"

/*
1.graph是個抽象的資料類別
2.graph是個有限的集合(a finite set of vertices)
3.在graph的點和點之間相連的線，稱之為edges，
如果線是有方向性的，則稱為arrow，graph則稱為directive graph。
4.graph就是有形成cycle的tree
5.adjacency list可以用來實現graph，類似一個hashTable的結構，首先有個陣列，陣列裡每個值都是一個陣列，用來表達塗上相鄰的節點vertex
6.adjacency matrix用雙重陣列來記載graph，每個節點在matrix上都得成對一個節點，紀錄彼此的相臨關係。
*/

//Graph represents an adjacency list graph
type Graph struct {
	vertices []*Vertex
}

//Vertex represents a graph vertex
type Vertex struct {
	key      int
	adjacent []*Vertex
}

//AddVertex adds Vertex to the Graph
func (g *Graph) AddVertex(n int) {
	if contains(g.vertices, n) {
		err := fmt.Errorf("Vertex %v not added because it is an existing key", n)
		fmt.Println(err.Error())
	} else {
		g.vertices = append(g.vertices, &Vertex{key: n})
	}
}

//Add Edge
//getVertex
//contain
func contains(s []*Vertex, n int) bool {
	for i := 0; i < len(s); i++ {
		if s[i].key == n {
			return true
		}
	}
	return false
}

//print all the adjacent list for each vertex of graph
func (g *Graph) PrintAllVertexAdjacent() {
	for _, v := range g.vertices {
		fmt.Printf("\nVertex %v :", v.key)
		for _, v := range v.adjacent {
			fmt.Printf(" %v", v.key)
		}
	}
}

//AddEdge adds an edge to the graph
func (g *Graph) AddEdge(from,to int){
	//get vertex
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)
	//check error
	if fromVertex == nil || toVertex == nil{
		err := fmt.Errorf("Invalid edge (%v --> %v)",from,to)
		fmt.Println(err.Error())
	}else if contains(fromVertex.adjacent,to){
		err := fmt.Errorf("Existing edge (%v --> %v)",from,to)
		fmt.Println(err.Error())
	}	else{
		//add edge
		fromVertex.adjacent = append(fromVertex.adjacent,toVertex)
	}
}
//getVertex returns a pointer to the vertex with a key integer
func (g *Graph) getVertex(k int) *Vertex{
	for i,v := range g.vertices{
		if v.key == k{
			return g.vertices[i]
		}
	}
	return nil
}

func main() {
	graphTest := Graph{}
	for i := 0; i < 5; i++ {
		graphTest.AddVertex(i)
	}
	graphTest.AddEdge(1,2)
	graphTest.AddEdge(1,2)
	graphTest.AddEdge(6,2)
	graphTest.AddEdge(3,2)

	graphTest.PrintAllVertexAdjacent()
}
