package main

import "fmt"

var x = []int{4, 7, 1, 2, 5, 3, 8, 10, 33, 6, 77, 100, 101}
var y = []int{10, 20, 30, 5, 7, 9, 11, 13, 15, 17, 35}

func main() {
	z := quickSortWithPartition(x, 0, len(x)-1)
	fmt.Println("quick sort", z)
	y := mergeSort(x)
	fmt.Println("merge sort", y)

	m := &MaxHeap{}
	res := []int{}
	for _, v := range x {
		m.Insert(v)
	}
	for i := 0; i < len(x); i++ {
		res = append(res, m.Extract())
	}
	fmt.Println("heap sort(max)", res)

	// n := &MinHeap{}
	// resMin := []int{}
	// for _, v := range x {
	// 	n.InsertMin(v)
	// }
	// for i := 0; i < len(x); i++ {
	// 	resMin = append(resMin, n.ExtractMin())
	// }
	// fmt.Println("heap sort(min)", resMin)
}

/* n^2 系列*/
/*
	BubbleSort
	由後到前兩兩比較，較小的數字往前排，由此排序n平方次。
	1.設定一個pointer，len(arr)，loop讓arr[len(arr)]往前跑去比較與arr[len(arr)-1]大小
	2.第一次迴圈找出陣列第一個值，也就是最小值，第二次迴圈找出陣列第二個值次小值。
	3.所以需要兩個迴圈一個由最末位往前前後比較，一個遍歷所有迴圈讓每個值都排序。
	Bubble Sort的BigO是n2
	因為每一位都要由1到n比對一次，所以(n-1)+(n-2)+(n-3)+(n-4)+.....+0+1=((n-1)*n)/2 => n2( 常數省略、小數省略 )
	最糟的狀況BubbleSort要O(n^2)
	最好的狀況(陣列本身就排好)，則只需要O(n)
*/
func bubbleSort(arr []int) []int {
	var j = 0
	for j < len(arr) {
		for i := len(arr) - 2; i >= 0; i-- {
			if arr[i+1] < arr[i] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
		j++
	}
	return arr
}

/*
	Insertion Sort插入排序法
	由前到後，排列第i項往前到0位的順序
	for i loop to len(arr)-1
	j = i - 1
	then sort j to 0 position

	Big O notation n*(n-1)/2
	複雜度n^2
*/
func insertionSort(arr []int) []int {
	var n = len(arr)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
			j--
		}
	}
	return arr
}

/*
	Selection Sort
	The principle of selection sort is - select the smallest value in unsorted array,
	and swap it with the left most value in this unsorted array.
	在陣列中找到最小值，依序排入index 0 ~ len(arr)-1

	selection sort一樣worst case 是 O(n^2), best case 也是 O(n^2)
*/
func selectionSort(arr []int) []int {
	var n = len(arr)
	for i := 0; i < n; i++ {
		var minIdx = i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
	return arr
}

/*
	Merge sort
	merge sort是很經典的"divide and merge"案例。
	merge sort利用拆分與合併兩個排列好的陣列兩個優點實現O(n logn n)的時間複雜度
	原理就是把未排序的大陣列切成以一個元素為單位的小陣列，再將之一個一個合併回排列好的大陣列。
	其實你會發現，merge sort就是Preorder Traversal。
	如果今天陣列的長度為n，要拆分log(2,n)次，且每一層要排序n次，所以mergesort的複雜度是n*log(2,n)
	keywords:recursion
	recursion會把原初陣列切成len(arr)個單一值的陣列，再將其一一合併。
	*會進入merge的左右兩個array都已經是sorted array
*/

func mergeSort(arr []int) []int {
	// fmt.Println("拆分 arr", arr)
	//拆分
	if len(arr) > 1 {
		midIdx := len(arr) / 2
		arr_left := mergeSort(arr[:midIdx])
		arr_right := mergeSort(arr[midIdx:])
		//合併
		return merge(arr_left, arr_right)
	}
	return arr
}

func merge(left_arr, right_arr []int) []int {
	// fmt.Println("合併left_arr", left_arr, "right_arr", right_arr)
	final := []int{}
	i := 0
	j := 0
	for i < len(left_arr) && j < len(right_arr) {
		if left_arr[i] < right_arr[j] {
			final = append(final, left_arr[i])
			i++
		} else {
			final = append(final, right_arr[j])
			j++
		}
	}
	//上面迴圈跑完，左右陣列一定還會遺留一些值，大於final內的任何值，下面用loop將之加入
	// fmt.Println("合併run after for", final)
	for ; i < len(left_arr); i++ {
		final = append(final, left_arr[i])
	}
	for ; j < len(right_arr); j++ {
		final = append(final, right_arr[j])
	}
	// fmt.Println("合併run after append", final)
	return final
}

/* log
拆分 arr [4 7 1 2 5 3 8 10 33 6 77 100 101]
拆分 arr [4 7 1 2 5 3]
拆分 arr [4 7 1]
拆分 arr [4]
拆分 arr [7 1]
拆分 arr [7]
拆分 arr [1]
合併left_arr [7] right_arr [1]
合併run after for [1]
合併run after append [1 7]
合併left_arr [4] right_arr [1 7]
合併run after for [1 4]
合併run after append [1 4 7]
拆分 arr [2 5 3]
拆分 arr [2]
拆分 arr [5 3]
拆分 arr [5]
拆分 arr [3]
合併left_arr [5] right_arr [3]
合併run after for [3]
合併run after append [3 5]
合併left_arr [2] right_arr [3 5]
合併run after for [2]
合併run after append [2 3 5]
合併left_arr [1 4 7] right_arr [2 3 5]
合併run after for [1 2 3 4 5]
合併run after append [1 2 3 4 5 7]
拆分 arr [8 10 33 6 77 100 101]
拆分 arr [8 10 33]
拆分 arr [8]
拆分 arr [10 33]
拆分 arr [10]
拆分 arr [33]
合併left_arr [10] right_arr [33]
合併run after for [10]
合併run after append [10 33]
合併left_arr [8] right_arr [10 33]
合併run after for [8]
合併run after append [8 10 33]
拆分 arr [6 77 100 101]
拆分 arr [6 77]
拆分 arr [6]
拆分 arr [77]
合併left_arr [6] right_arr [77]
合併run after for [6]
合併run after append [6 77]
拆分 arr [100 101]
拆分 arr [100]
拆分 arr [101]
合併left_arr [100] right_arr [101]
合併run after for [100]
合併run after append [100 101]
合併left_arr [6 77] right_arr [100 101]
合併run after for [6 77]
合併run after append [6 77 100 101]
合併left_arr [8 10 33] right_arr [6 77 100 101]
合併run after for [6 8 10 33]
合併run after append [6 8 10 33 77 100 101]
合併left_arr [1 2 3 4 5 7] right_arr [6 8 10 33 77 100 101]
合併run after for [1 2 3 4 5 6 7]
合併run after append [1 2 3 4 5 6 7 8 10 33 77 100 101]
[1 2 3 4 5 6 7 8 10 33 77 100 101]
*/

/*
藉由max heap tree來進行sorting的排序法。
Heap sort
1. build MaxHeap struct
2. func Insert which append the new element to the final slice,then maxHeapifyUp it to the right place
3. func Extract the root element which is the largest num in the max heap tree, swap the root and the final element,then remove the root,maxHeapifyDown the new root to the right place.
4. func HeapifyUp,let the node of index get from parameter compare with its parent node,if larger than parent node,swap it, until it find the right place to go
5. func HeapifyDown, let the node of index get from parameter compare with their child nodes,if smaller than child node, swap it ,until node of index find the right place to go.
6. func swap , switch two element's position in slice.
7. parent, left, right,got index and find their parent, left child, right child.

maxHeapifyDown 是O(log n)
maxHeapifyUp 也是O(log n)
*/

type MaxHeap struct {
	arr []int
}

func (h *MaxHeap) Insert(value int) {
	h.arr = append(h.arr, value)
	h.maxHeapifyUp(len(h.arr) - 1)
}

func (h *MaxHeap) Extract() int {
	if len(h.arr) == 0 {
		fmt.Println("there is no more element can be extracted.")
		return -1
	}
	extracted := h.arr[0]
	l := len(h.arr) - 1
	h.arr[0] = h.arr[l]
	h.arr = h.arr[:l]
	h.maxHeapifyDown(0)
	return extracted
}

func (h *MaxHeap) maxHeapifyUp(i int) {
	for h.arr[i] > h.arr[parent(i)] {
		h.swap(i, parent(i))
		i = parent(i)
	}
}

func (h *MaxHeap) maxHeapifyDown(i int) {
	l, r := left(i), right(i)
	lastIndex := len(h.arr) - 1
	compareIndx := 0
	for l <= lastIndex {
		if l == lastIndex {
			compareIndx = l
		} else if h.arr[l] > h.arr[r] {
			compareIndx = l
		} else {
			compareIndx = r
		}

		if h.arr[i] < h.arr[compareIndx] {
			h.swap(i, compareIndx)
			i = compareIndx
			l, r = left(i), right(i)
		} else {
			return
		}
	}
}

func (h *MaxHeap) swap(i1, i2 int) {
	h.arr[i1], h.arr[i2] = h.arr[i2], h.arr[i1]
}

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return i*2 + 1
}

func right(i int) int {
	return i*2 + 2
}

type MinHeap struct {
	arr []int
}

func (h *MinHeap) InsertMin(value int) {
	h.arr = append(h.arr, value)
	h.minHeapifyUp(len(h.arr) - 1)
}

func (h *MinHeap) ExtractMin() int {
	if len(h.arr) == 0 {
		fmt.Println("there is no more element to extracted!")
		return -1
	}
	l := len(h.arr) - 1
	extracted := h.arr[0]
	h.arr[0] = h.arr[l]

	h.arr = h.arr[:l]
	h.minHeapifyDown(0)
	return extracted
}

func (h *MinHeap) minHeapifyUp(i int) {
	for h.arr[i] < h.arr[parent(i)] {
		h.swapMin(i, parent(i))
		i = parent(i)
	}
}

func (h *MinHeap) minHeapifyDown(i int) {
	l, r := left(i), right(i)
	lastIndex := len(h.arr) - 1
	compareIdx := 0

	for l <= lastIndex {
		if l == lastIndex {
			compareIdx = l
		} else if h.arr[l] > h.arr[r] {
			compareIdx = r
		} else {
			compareIdx = l
		}

		if h.arr[i] > h.arr[compareIdx] {
			h.swapMin(i, compareIdx)
			i = compareIdx
			l, r = left(i), right(i)
		} else {
			return
		}
	}
}

func (h *MinHeap) swapMin(i1, i2 int) {
	h.arr[i1], h.arr[i2] = h.arr[i2], h.arr[i1]
}

/*
Quick Sort
Partition:是QuickSort的根本algorithm。
設定最後一位為pivot，
put number less than pivot partition in pivot left side,
greater than pivot partition in the right side.
than recursion Partition function while less and greater slice finished.
將slice中小於pivot的數字放在pivot左邊，大於pivot的數字放在pivot右邊。

Quick Sort有兩種，一種是有用Partition的，一種沒有，沒有的會需要比較多的空間複雜度。

quick sort 
worse case performance:O(n^2)
Best case performance:O(n log n)
Average performance:O(n log n)
*/
//Without Partition
func quickSortWithoutPartition(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	pivot := arr[len(arr)-1]
	// fmt.Println("pivot",pivot)
	final, left, right := []int{}, []int{}, []int{}

	for i, v := range arr {
		if i == len(arr)-1 {
			break
		} else if v < pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}

	left = quickSortWithoutPartition(left)
	right = quickSortWithoutPartition(right)

	final = append(final, left...)
	final = append(final, pivot)
	final = append(final, right...)

	return final
}

func quicksort_swap(arr []int, i1, i2 int) []int {
	arr[i1], arr[i2] = arr[i2], arr[i1]
	return arr
}

/*
with partition
partition給定範圍start和end，和兩個pointer i 和 j
在start和end的範圍中去比較其內的元素跟pivot之間的大小關係，
j遍歷array，如果遇到較pivot小的值，j和i就swap，交換位置，然後i往前一格。
j遍歷完陣列後，pivot再跟最後的i值交換位置，如此就會形成pivot左邊比pivot小
pivot右邊比pivot大的形勢。
*/

func partition(arr []int, start, end int) ([]int, int) {
	pivot := arr[end]
	i := start
	for j := start; j < end; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[end] = arr[end], arr[i]
	return arr, i
}

func quickSortWithPartition(arr []int, start, end int) []int {
	if start < end {
		var pivot int
		arr, pivot = partition(arr, start, end)
		arr = quickSortWithPartition(arr, start, pivot-1)
		arr = quickSortWithPartition(arr, pivot+1, end)
	}
	return arr
}

/*
lower bound searching最小就是logn，不能再更小了
master theorem
*/
