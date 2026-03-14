package heapp

type MinHeap struct {
	Data []int
}

func (h *MinHeap) Insert(val int) {
	h.Data = append(h.Data, val)
	h.heapifyUp(len(h.Data) - 1)
}

func (h *MinHeap) heapifyUp(index int) {
	for index > 0 {
		parent := (index - 1) / 2
		if h.Data[parent] < h.Data[index] {
			break
		}
		h.Data[index], h.Data[parent] = h.Data[parent], h.Data[index]
		index = parent
	}
}

func (h *MinHeap) ExtractMin() int {
	if len(h.Data) == 0 {
		return -1
	}
	min := h.Data[0]
	last := h.Data[len(h.Data)-1]
	h.Data[0] = last
	h.Data = h.Data[:len(h.Data)-1]
	h.heapifyDown(0)
	return min
}

func (h *MinHeap) heapifyDown(index int) {
	lastIndex := len(h.Data) - 1
	for {
		left := 2*index + 1
		right := 2*index + 2
        smallest := index 
        
        if left <= lastIndex && h.Data[left] < h.Data[smallest]{
            smallest = left
        }
        if right <= lastIndex && h.Data[right]  < h.Data[smallest]{
            smallest = right
        }
        if smallest == index{
            break
        }
        h.Data[index],h.Data[smallest] = h.Data[smallest],h.Data[index]
        index = smallest
	}
}