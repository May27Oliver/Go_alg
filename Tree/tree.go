package main

func main(){
	tree := Node{val:100}
}

/*
tree
tree is "acyclic graph"
tree must have a root 
--- left - root - right ----
樹有很多種
Binary Tree(二元樹)：每個root最多只能有兩個子元素。
Binary Search Tree：二元樹中，右邊子節點 > parent node > 左邊子節點，稱為Binary Search Tree。
Complete Binary Tree：是個幾乎滿的二元樹
Full Binary Tree：滿的二元樹，所有的leaf(最外層的子節點)都有一樣的深度。
Max Heap：是個complete binary tree，且每個node都要比自己的child來得大，最大的值必定在第一個root。
*/

//Node
//Insert
//Search
type Node struct{
	val interface{} //節點值
	left *Node			//左節點
	right *Node			//右節點
}
