/*
堆排序

*/

package heap

type Heap struct {
	//堆长度
	size int

	//堆数值
	data []int
}

//创建一个堆对象
func New(args ...int) *Heap {
	var heap *Heap

	//如果有参数，则依次加入堆中
	if len(args) > 0 {
		heap = &Heap{data: make([]int, 1, len(args)+1)}
		for i := 0; i < len(args); i++ {
			heap.Push(args[i])
		}
	} else {
		heap = &Heap{data: make([]int, 1)}
	}

	return heap
}

//插入一个值
func (h *Heap) Push(v int) {
	//长度加1
	h.size++

	//放入到最后一个位置
	h.data = append(h.data, v)

	//调整堆数据
	h.heapInsert()
}

//根据新插入一个新值做调整
func (h *Heap) heapInsert() {
	index := h.size
	for j := index >> 1; j > 0; {
		//与父节点比较
		//小于或等于，直接退出
		if h.data[j] >= h.data[index] {
			break
		}

		//大于父节点，则交换数据，同时交换下标
		h.data[j], h.data[index] = h.data[index], h.data[j]
		index = j
		j = j >> 1
	}
}

//弹出最大值
func (h *Heap) Pop() int {
	//第一个位置为最大值
	result := h.data[1]

	//把最后一个值复制到第一个值
	h.data[1] = h.data[h.size]

	//长度减一
	h.size--

	//调整堆
	h.heapify()

	//返回结果
	return result
}

//根据弹出值做调整
func (h *Heap) heapify() {
	index := 1

	//比较左右子节点
	for left := index << 1; left <= h.size; {
		right := left | 1
		bigIndex := left
		//如果存在右节点，取左右节点中最大值的下标
		if right <= h.size && h.data[right] > h.data[left] {
			bigIndex = right
		}

		//如果左右节点最大值不比当前节点大，则退出循环
		if h.data[bigIndex] <= h.data[index] {
			break
		}

		//交换最大值下标与当前节点下标
		h.data[bigIndex], h.data[index] = h.data[index], h.data[bigIndex]
		index = bigIndex
		left = index << 1

	}

}

//排序
func (h *Heap) Sort() []int {
	size := h.Size()
	res := make([]int, size)
	for i := 1; i <= size; i++ {
		res[size-i] = h.Pop()
	}

	return res
}

//获取堆长度
func (h *Heap) Size() int {
	return h.size
}

//是否为空
func (h *Heap) IsEmpty() bool {
	return h.Size() == 0
}

//获取数据
func (h *Heap) Data() []int {
	data := make([]int, h.size)
	copy(data, h.data[1:h.size+1])
	return data
}
