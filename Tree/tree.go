package main

import "fmt"

/*
tree
tree is "acyclic graph"
tree is a graph without a loop，tree是沒有循環的。
tree must have a root,and the only root,tree每個node只能有一個root

--- left - root - right ----

樹有很多種
Binary Tree(二元樹)：每個root最多只能有兩個子元素。
Binary Search Tree：二元樹中， parent node > 右邊子節點 > 左邊子節點，稱為Binary Search Tree。
以二元樹原則組成的資料結構在找尋資料的速度是很快的，為O(log n)

Complete Binary Tree：是個幾乎滿的二元樹。
Full Binary Tree：滿的二元樹，所有的leaf(最外層的子節點)都有一樣的深度。
Max Heap：是個complete binary tree，且每個node都要比自己的child來得大，最大的值必定在第一個root。
*/

/*
tree traversal 樹的遍歷法
1.breadth-first tree traversal 寬度優先遍歷法
2.depth-first tree traversl 深度優先遍歷法

第二種又有
2-1. preorder 中左右
2-2. inorder	左中右
2-3. PostOrder	左右中
*/

//Node
type Node struct{
	Key int //節點值
	Left,Right *Node			//左節點
}

//也可以用建構函式的方式建立Node
func createTreeNode(val int) *Node {
	return &Node{Key:val}
}

//Insert
//Node can not be already in the tree
func (n *Node) Insert (k int){
	if k == n.Key{//重複不安插新node
		return
	} 
	if n.Key < k {//比root大放右邊
		if n.Right == nil{//如果右側本是空則放右側
			n.Right = &Node{Key:k}
		}else{//如果右側不為空，則在放進右側的右側，這個會是個recursion
			n.Right.Insert(k)
		}
	}else if n.Key > k {//比root小放左邊
		if n.Left == nil{//如果左側本是空則放左側
			n.Left = &Node{Key:k}
		}else{//如果左側不為空，則在放進左側的左側，這個會是個recursion
			n.Left.Insert(k)
		}
	}
}
//Search
func (n *Node) Search (k int)bool{
	if n == nil{
		return false
	}
	if n.Key < k {//比root大放右邊
		return n.Right.Search(k)
	}else if n.Key > k {//比root小放左邊
		return n.Left.Search(k)
	}
	return true
}
//印出Key
func (n *Node) print(){
	fmt.Print(n.Key," ")
}

//preorder traversal 中左右
func (n *Node) preorder(){
	if n == nil{
		return
	}
	n.print()
	n.Left.preorder()
	n.Right.preorder()
}

//inorder traversal 左中右
func (n *Node) inorder(){
	if n == nil{
		return
	}
	n.Left.inorder()
	n.print()
	n.Right.inorder()
}

//postorder traversal 左右中


func main(){
	// root := new(Node)
	// root.Key = 100 以上兩行跟下一行同義
	root := &Node{Key:100}//因為要呼叫內部方法，所以得用pointer
	
	root.Insert(20)
	fmt.Println(root)
}