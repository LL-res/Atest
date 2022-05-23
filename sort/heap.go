package sort

//堆排序

/*值得注意的是下面这个函数没有构建出一个十分标准的大顶堆
它的作用就跟它的名字一样，只负责管选出最大的那个，整个树的结构就好似一场场的擂台赛，两个子加上父一打，打出一个冠军
之后就马上开始向上层继续打，打出的冠军提出来，放到队列的最后，把最后那个元素放到堆首继续打一波*/
//这个评注是最开始写的，但过程跟这个很相似，只不过新写的函数实现了完整的大根堆

func HeapSort(arr []int) {
	HeapBuild(arr)
	for i := len(arr) - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		HeapSink(arr[:i], 0)
	}
}

//这个sort的过程就是把堆化后的数组的堆顶，也就是最大的那个元素，抽出来放到队尾部，形成最大的数组尾部
//然后把原来尾部那个元素再放到队首，进行一波下沉，再一次形成一个堆，如此往复，也就完成了排序

func HeapSink(arr []int, i int) {
	max := i
	left, right := 2*i+1, 2*i+2
	if left < len(arr) && arr[left] > arr[max] {
		max = left
	}
	if right < len(arr) && arr[right] > arr[max] {
		max = right
	}
	if max != i {
		arr[i], arr[max] = arr[max], arr[i]
		HeapSink(arr, max)
	}
}
func HeapBuild(arr []int) {
	for i := len(arr)/2 - 1; i >= 0; i-- { //这个i的开始的点是最后一个有叶子节点的点。还可以打擂台的地方
		HeapSink(arr, i)
	}
}

//这个build复杂度只有O（N），注意一点，一定要从下往上进行更新，也就是打擂台的过程
