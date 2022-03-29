package main

import "fmt"

//graph struct
//Graph內有個陣列，裝著每個node
type Graph struct{
	vertices []*Vertex
}
//vertex struct
type Vertex struct {
	key int
	adjacent []*Vertex
	visited bool
}
//AddVertex
func (g *Graph) AddVertex(k int){
	if contains(g.vertices,k){
		err:=fmt.Errorf("Vertex %v not added because it is an existing key",k)
		fmt.Println(err.Error())
	}
	g.vertices = append(g.vertices,&Vertex{key:k,visited: false})
}
//contain
func contains(s []*Vertex,n int)bool{
	for _,v:=range s{
		if v.key == n{
			return true
		}
	}
	return false
}
//PrintAllGraph
func (g *Graph) PrintAllGraph(){
	for _,v := range g.vertices{
		fmt.Printf("\nVertex %v :",v.key)
		for _,v := range v.adjacent{
			fmt.Printf(" %v",v.key)
		}
	}
}

func (g *Graph) getVertex(k int)*Vertex{
	for _,v:=range g.vertices{
		if v.key == k{
			return v
		}
	}
	return nil
}

func (g *Graph) AddEdge(from, to int){
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)
	if fromVertex == nil || toVertex == nil{
		err:=fmt.Errorf("Invalid edge (%v --> %v)",from,to)
		fmt.Println(err.Error())
	}else if contains(fromVertex.adjacent,to){
		err:=fmt.Errorf("Existing edge (%v --> %v)",from,to)
		fmt.Println(err.Error())
	}else{
		//add edge
		fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
	}
}

//BFS寬度優先法
func (g *Graph) BFS(starter *Vertex){
	count := 0//計算輪到幾次node了，次數必須等於圖上的節點數。
	queue := make([]*Vertex,0)//先進先出法把所有的節點放進res中。
	queue = append(queue, starter)
	for len(queue) > 0 {
		//將queue內第一個Node取出
		firstNode := queue[0]
		queue = queue[1:]
		if !firstNode.visited{//若節點還未被訪問過，則將之加入res，標誌已訪。
				firstNode.visited = true
				count += 1
				if count == len(g.vertices){
					fmt.Println(firstNode.key)
					return
				}
				fmt.Printf("%v -> ",firstNode.key)
				for _,v := range firstNode.adjacent{
					if !v.visited{
						queue = append(queue,v)
					}
				}
		}
	}
}

func (g *Graph) DFS(starter *Vertex,count int){
	starter.visited = true
	if count == len(g.vertices){
		fmt.Printf("%v ",starter.key)
		return
	}
	fmt.Printf("%v -> ",starter.key)
	for _,v:= range starter.adjacent{
		if !v.visited{
			count += 1
			//遞迴
			g.DFS(v,count)
		}
	}
}

func (g *Graph) DFSTriversal(starter *Vertex){
	count := 1
	g.DFS(starter,count)
}

func main(){
	g:=Graph{}   
	for i := 0; i < 6; i++{
		g.AddVertex(i)
	}
	g.AddEdge(0,2)
	g.AddEdge(2,0)
	
	g.AddEdge(0,3)
	g.AddEdge(3,0)

	g.AddEdge(2,4)
	g.AddEdge(4,2)

	g.AddEdge(1,2)
	g.AddEdge(2,1)

	g.AddEdge(3,4)
	g.AddEdge(4,3)

	g.AddEdge(5,1)
	g.AddEdge(1,5)

	g.AddEdge(5,4)
	g.AddEdge(4,5)

	g.AddEdge(1,3)
	g.AddEdge(3,1)

	//深度優先法
	g.DFSTriversal(g.vertices[0])
	// 寬度優先法
	// g.BFS(g.vertices[0])
}