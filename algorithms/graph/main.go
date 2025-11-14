package main

import (
	"fmt"
)

type Vertex struct {
	Key      int
	Vertices []*Vertex
}

type Graph struct {
	vertices []*Vertex
}

func main() {
	g := Graph{}
	g.add_vertices(1)
	g.add_vertices(2)
	g.add_vertices(3)

	g.add_edges(1, 2)
	g.add_edges(1, 3)
	g.add_edges(2, 3)
	g.showGraph()
}

func (g *Graph) add_vertices(v int) {
	temp := &Vertex{Key: v}
	g.vertices = append(g.vertices, temp)
}

func (g *Graph) get_vertex(k int) *Vertex {
	for _, v := range g.vertices {
		if v.Key == k {
			return v
		}
	}
	return nil
}

// add edges to each verteses
func (g *Graph) add_edges(to, from int) {
	// change int to vertices
	toVertex := g.get_vertex(to)
	fromVertex := g.get_vertex(from)
	// in graph teake to and from points
	for _, v := range g.vertices {
		if v.Key == toVertex.Key {
			toVertex.Vertices = append(toVertex.Vertices, fromVertex)
			// append from.vertices to to
		}
	}

}

// print all vertex and list of related vertexes

func (g *Graph) showGraph() {

	for _, v := range g.vertices {
		fmt.Printf("%v", v.Key)
		for _, e := range v.Vertices {
			fmt.Printf("--> %v", e.Key)
		}
		fmt.Println()
	}
	// return nil
}
