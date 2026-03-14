package heapp

type MaxHeap struct {
	Data []int
}

func (h *MaxHeap) Insert(val int) {
	h.Data = append(h.Data, val)
	h.heapifyUp(len(h.Data) - 1)
}

func (h *MaxHeap) heapifyUp(index int) {
	for index > 0 {
		parent := (index - 1) / 2
		if h.Data[parent] > h.Data[index] {
			break
		}
		h.Data[parent], h.Data[index] = h.Data[index], h.Data[parent]
		index = parent
	}
}

func (h *MaxHeap) ExtractMax() int {
	if len(h.Data) == 0 {
		return -1
	}
	max := h.Data[0]
	last := h.Data[len(h.Data)-1]
	h.Data[0] = last
	h.Data = h.Data[:len(h.Data)-1]
	h.heapifyDown(0)
	return max
}

func (h *MaxHeap) heapifyDown(index int) {
	lastIndex := len(h.Data) - 1
	for {
		left := 2*index + 1
		right := 2*index + 2
        largest := index
        
        if left <= lastIndex && h.Data[left] > h.Data[largest]{
            largest = left
        }
        if right <= lastIndex && h.Data[right] > h.Data[largest]{
            largest = right
        }
        if largest == index{
            break
        }
        h.Data[index] , h.Data[largest] = h.Data[largest],h.Data[index]
        index = largest
	}
}
