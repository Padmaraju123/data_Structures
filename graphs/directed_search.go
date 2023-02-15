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


type Node struct{
	element int
	next *Node
}

type queue struct{
	front *Node
	rear *Node
	size int
}

type graph struct {
	vertices_size int
	adjmat        [7][7]int
	queue 
}




func (obj1 *queue)isempty()bool{
	if obj1.size ==0{
		return true
	}
	return false
}

func (obj1 *queue)length()int{
	return obj1.size
}

func (obj1 *queue)enqueue(val int){

	newnode := Node{element: val}
	if obj1.isempty(){
		obj1.front = &newnode
	}else{
		obj1.rear.next = &newnode
	}
	obj1.rear = &newnode
	obj1.size++
}

func (obj1 *queue)dequeue()int{

	if obj1.isempty(){
		fmt.Println("no elements are found")
	}else{
		e := obj1.front.element
		obj1.front = obj1.front.next
		obj1.size--
		if obj1.isempty(){
			obj1.rear = nil
		}
		return e
	}
	return -1
}

func (obj1 *queue)first()int{
	if obj1.isempty(){
		fmt.Println("queue is empty...")
		return -1
	}
	return obj1.front.element
}

func (obj1 *queue)displaying(){
	comm := obj1.front
	for comm!=nil{
		fmt.Print(comm.element,"--->")
		comm = comm.next
	}
}






func (obj *graph) insert_edge(u, v int) {
	obj.adjmat[u][v] = 1
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

func (obj *graph) display_mat() [7][7]int {
	return obj.adjmat
}

func (obj *graph )bfs(s int,obj1 *queue){
	i:= s
	visited := make([]int,4)
	fmt.Print(i," ")
	visited[i] = 1
	obj1.enqueue(i)
	vv := obj1.isempty()
	for vv != vv{
		i = obj1.dequeue()
		for j := 0 ;j<obj.vertices_size;j++{
			if obj.adjmat[i][j] == 1 && visited[j]==0{
				 fmt.Print(j," ")
				 visited[j] = 1
				 obj1.enqueue(j)
			}
		}

	}
}

func main() {
	obj := graph{}
	obj1 := queue{}

	fmt.Println("enter the no of vertices...")
	fmt.Scanln(&obj.vertices_size)

	fmt.Println("display the matrix....")
	fmt.Println(obj.display_mat())

	fmt.Printf("no of edge count... %d\n", obj.edge_count())

	obj.insert_edge(0, 1)
	obj.insert_edge(0, 5)
	obj.insert_edge(0, 6)
	obj.insert_edge(1, 0)
	obj.insert_edge(1, 2)
	obj.insert_edge(1, 5)
	obj.insert_edge(1, 6)
	obj.insert_edge(2, 3)
	obj.insert_edge(2, 4)
	obj.insert_edge(2, 6)
	obj.insert_edge(3, 4)
	obj.insert_edge(4, 2)
	obj.insert_edge(4, 5)
	obj.insert_edge(5, 2)
	obj.insert_edge(5, 3)
	obj.insert_edge(6, 3)


	fmt.Println("display the matrix....")
	fmt.Println(obj.display_mat())

	fmt.Printf("no of edge count... %d\n", obj.edge_count())

	fmt.Println("showing the edges between vertices of graphs...")
	obj.edges()

	fmt.Printf("the out-degree of vertex A %d\n", obj.outdegree(0))
	fmt.Printf("the in-degree of vertex A %d\n", obj.indegree(0))

	fmt.Printf("checking whether edge exist between the vertices of 2-3 : %v\n", obj.exist_edge(2, 3))
	fmt.Printf("checking whether edge exist between the vertices of 1-3 : %v\n", obj.exist_edge(1, 3))

	obj.bfs(0)

}
