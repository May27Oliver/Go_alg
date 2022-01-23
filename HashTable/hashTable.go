package main

import "fmt"

/*
HashTable
HashTable指透過一個Hash function產生一組數字
當作array的index把要儲存的資料放在該index，
如果有資料的index重複，則用Linked List儲存在同一個index下。
集合linked list優點與array優點的資料結構
*/

const ArraySize = 7 
//如果你的arraysize為7，那麼只要ascii 總和除7餘數相同的都會進同一個index的linked list內

//HashTable Structure(array)
type HashTable struct{
	array [ArraySize]*bucket
}

//Insert
func (h *HashTable) Insert(s string){
	index := hash(s)
	h.array[index].insert(s)
}

//Search
func (h *HashTable) Search(s string)bool{
	index := hash(s)
	return h.array[index].search(s)
}

//Delete
func (h *HashTable) Delete(s string){
	index := hash(s)
	h.array[index].delete(s)
}

//hash
func hash(val string)int{
	sum := 0
	for _,v := range val{
		sum +=int(v)
	}
	return sum % ArraySize
}



/*
Bucket Structure(Linked list)
insert
search
delete
*/
type bucket struct {
	head *bucketNode
}

type bucketNode struct{
	key string
	next *bucketNode
}

//Bucket insert
func (b *bucket) insert(s string){
	if !b.search(s){
		newNode := &bucketNode{key:s}
		newNode.next = b.head
		b.head = newNode 
		return
	}
	fmt.Printf("%s is already in this slot",s)
}

//Bucket search
func (b *bucket) search(s string)bool{
	pointer := b.head
	for pointer != nil{
		if pointer.key == s {
			return true
		}
		pointer = pointer.next
	}
	return false
}

////Bucket delete
func (b *bucket) delete(s string){
	if b.head.key == s {
		b.head = b.head.next
		return
	}
	pointer := b.head
	for pointer.next != nil{
		if pointer.next.key == s {
			pointer.next = pointer.next.next
			return
		}
		pointer = pointer.next
	}
	return
}
func printAll(bucket *bucket){
	pointer := bucket.head
	for pointer != nil{
		fmt.Println(pointer.key)
		pointer = pointer.next
	}
}

func Init() *HashTable{
	result:= &HashTable{}
	for i := range result.array{
		result.array[i] = &bucket{}
	}
	return result
}

func main(){
	hashTable := Init()
	var list = []string{
		"ERIC",
		"KENNY",
		"KYLE",
		"STAN",
		"RANDY",
		"BUTTERS",
		"TOKEN",
	}
	for _,v := range list{
		hashTable.Insert(v)
	}
	index := hash("KENNY")
	printAll(hashTable.array[index])
	hashTable.Delete("ERIC")
	printAll(hashTable.array[index])
}