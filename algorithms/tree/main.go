package main

import "fmt"

type Tree struct {
	Data  int
	Left  *Tree
	Right *Tree
}

func main() {
	t := &Tree{Data: 50, Left: nil, Right: nil}
	t.Insert(100)
	t.Insert(20)
	t.Insert(55)
	t.Insert(23)
	t.Insert(11)
	t.Insert(43)
	t.Insert(788)
	t.Insert(454)
	t.Insert(246)

	// fmt.Println(t.lenght)
	// fmt.Println(t.Right.Right.Data)
	// fmt.Println(t.Left.Left.Data)
	// fmt.Println(t.Lookup(20))
	// out := Traverse(t)
	// fmt.Println(out)
	out := BFST(t)
	fmt.Println(out)
	// t.Root.Traverse()
	// print(os.Stdout, t.Root, 0, 'M')

}

func BFST(tree *Tree) []int {
	result := []int{}
	queue := []*Tree{}
	queue = append(queue, tree)
	return bfstUtil(queue, result)
}

func bfstUtil(queue []*Tree, result []int) []int {
	if len(queue) == 0 {
		return result
	}

	result = append(result, queue[0].Data)

	if queue[0].Right != nil {
		queue = append(queue, queue[0].Right)
	}
	if queue[0].Left != nil {
		queue = append(queue, queue[0].Left)
	}
	return bfstUtil(queue[1:], result)
}

func DFS(tree *Tree) []int {
	// mark node visited
	visited := []*Tree{}
	result := []int{}
	visited = append(visited, tree)
	DFSUtil(visited, result)
	return result
}

func DFSUtil(visited []*Tree, result []int) []int {
	if len(visited) == 0 {
		return result
	}
	result = append(result, visited[0].Data)
	if visited[0].Right != nil {

	}
	return result
}

func (t *Tree) Insert(data int) *Tree {
	temp := Tree{Data: data, Left: nil, Right: nil}
	if t.Data == 0 {
		t.Left = &temp
		return &temp
	} else {
		current := t
		for {
			if data < current.Data {
				// left
				if current.Left == nil {
					current.Left = &temp
					return current
				}
				current = current.Left
			} else {
				// right
				if current.Right == nil {
					current.Right = &temp
					return current

				}
				current = current.Right
			}
		}
	}
}

func (t *Tree) Lookup(data int) *Tree {
	if t == nil {
		fmt.Println("non tree")
	} else {
		current := t
		for current != nil {
			if data < current.Data {
				current = current.Left
			} else if data > current.Data {
				current = current.Right
			} else if current.Data == data {
				return current
			}
		}
	}
	return nil
}

// in order Traverse
func Traverse(root *Tree) []int {
	if root == nil {
		return nil
	}
	output := make([]int, 0)
	left := Traverse(root.Left)
	right := Traverse(root.Right)

	output = append(output, left...)
	output = append(output, root.Data)
	output = append(output, right...)
	return output
}
