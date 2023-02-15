/*      A
       /  \
	  /     \
	B -------C
	        /
		   /
		  D
*/

package main

import "fmt"

type graph struct {
	vertices_size int
	adjmat        [4][4]int
}

func (obj *graph) insert_edge(u, v,w int) {
	obj.adjmat[u][v] = w
}

func (obj *graph) remove_edge(u, v int) {
	obj.adjmat[u][v] = 0
}

func (obj *graph) exist_edge(u, v int) bool {
	if obj.adjmat[u][v] != 0 {
		return true
	}
	return false
}

//return the vertices count in the graph
func (obj *graph) vertex_count() int {
	return obj.vertices_size
}

//return the edges count in the graph
func (obj *graph) edge_count() int {
	var count int
	for i := 0; i < obj.vertices_size; i++ {
		for j := 0; j < obj.vertices_size; j++ {
			if obj.adjmat[i][j] != 0 {
				count++
			}
		}
	}
	return count
}

//printing the edge weights in the graphs
func (obj *graph) edges() {
	for i := 0; i < obj.vertices_size; i++ {
		for j := 0; j < obj.vertices_size; j++ {
			if obj.adjmat[i][j] != 0 {
				fmt.Printf("%d -- %d\n", i, j)
			}
		}
	}
	fmt.Println()
}

//outdegree of given vertex of graph
func (obj *graph) outdegree(u int) int {
	var out_deg_count int

	for i := 0; i < obj.vertices_size; i++ {
		if obj.adjmat[u][i] != 0 {
			out_deg_count++
		}
	}

	return out_deg_count

}

//indegree of given vertex of graph
func (obj *graph) indegree(v int) int {
	var in_deg_count int

	for j := 0; j < obj.vertices_size; j++ {
		if obj.adjmat[j][v] != 0 {
			in_deg_count++
		}
	}

	return in_deg_count

}

func (obj *graph) display_mat() [4][4]int {
	return obj.adjmat
}

func main() {
	obj := graph{}

	fmt.Println("enter the no of vertices...")
	fmt.Scanln(&obj.vertices_size)

	//slice of weight of edges
	//note : if you doesn't assign the weights to the edge then it is not a weighted directed graph
	//sli := []int{0,1,1,0,0,0,1,0,0,0,0,1,0,0,0,0}

	fmt.Println("display the matrix....")
	fmt.Println(obj.display_mat())

	fmt.Printf("no of edge count... %d\n", obj.edge_count())

	obj.insert_edge(0, 1,24)
	obj.insert_edge(0, 2,10)
	obj.insert_edge(1, 2,15)
	obj.insert_edge(2, 3,30)

	fmt.Println("display the matrix....")
	fmt.Println(obj.display_mat())

	fmt.Printf("no of edge count... %d\n", obj.edge_count())

	fmt.Println("showing the edges between vertices of graphs...")
	obj.edges()

	fmt.Printf("the out-degree of vertex A %d\n", obj.outdegree(0))
	fmt.Printf("the in-degree of vertex A %d\n", obj.indegree(0))

	fmt.Printf("checking whether edge exist between the vertices of 2-3 : %v\n", obj.exist_edge(2, 3))
	fmt.Printf("checking whether edge exist between the vertices of 1-3 : %v\n", obj.exist_edge(1, 3))

}
