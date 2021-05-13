package gosort

import (
	"sort"
)

//堆排序
func Heap(data sort.Interface) {
	//调整为大根堆
	//从0-中间下标，依次往左右子节点调整
	for i := data.Len() >> 1; i >= 0; i-- {
		heapify(data, i, data.Len())
	}

	//堆调整完为大根堆之后，开始排序
	//1、收尾交互，最大值排在最后位置
	//2、再次调整，且 end 下标节点往前移
	//3、循环 data.Len() -1 之后，排序完毕
	for i := data.Len() - 1; i > 0; i-- {
		data.Swap(0, i)
		heapify(data, 0, i)
	}
}

//从指定根节点开始，到最大下标之间调整堆, 不包含 end
func heapify(data sort.Interface, root, end int) {
	for {
		//左右节点中最大节点，起初为左节点
		maxChild := root<<1 | 1

		//如果越界，则退出
		if maxChild >= end {
			break
		}

		//如果存在右节点，且右节点大于左节点，最大节点取右节点
		if maxChild+1 < end && data.Less(maxChild, maxChild+1) {
			maxChild++
		}

		//如果左右节点中最大节点比根节点小，则退出
		if data.Less(maxChild, root) {
			break
		}

		//否则根节点与左右节点最大节点互换
		data.Swap(maxChild, root)

		//同时根节点，调整为左右节点中最大节点,继续处理新根节点左右节点数据
		root = maxChild

	}
}
