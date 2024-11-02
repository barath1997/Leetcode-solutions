// O(n) solution , reference : https://m.youtube.com/watch?v=CZQGRp93K4k
func maxSlidingWindow(nums []int, k int) []int {
      
	dq := NewDeque()
	result := make([]int,0)
	n := len(nums)

	for i:=0;i<n;i++ {
       
		// code to remvove out of bound (values outside window size k) elements from the begin of dq
	   val := dq.PeekFront()
	   if dq.Size() != 0 && val == i-k {
		  dq.PopFront()
	   }
       
	   // to remove elements, those are less than the current index "i" elements (removal happens from behind)
	   for dq.Size() != 0 && nums[dq.PeekBack()] <= nums[i] {
		  dq.PopBack()
	   }
       
	   // adding index of elemts into dq
	   dq.PushBack(i)
       
	   // when index "i" becomes equal or greater than window size k , then we have to add the front element / greatest in the dq to the result 
	   if i >= k-1 {
		  result = append(result , nums[dq.PeekFront()])
	   }

	}
	
	return result


}

// Deque structure
type Deque struct {
	elements []int
}

// NewDeque creates a new Deque
func NewDeque() *Deque {
	return &Deque{elements: []int{}}
}

// PushFront adds an element to the front of the deque
func (d *Deque) PushFront(value int) {
	d.elements = append([]int{value}, d.elements...)
}

// PushBack adds an element to the back of the deque
func (d *Deque) PushBack(value int) {
	d.elements = append(d.elements, value)
}

// PopFront removes and returns the front element of the deque
func (d *Deque) PopFront() (int) {
	if len(d.elements) == 0 {
		return 0 // Deque is empty
	}
	front := d.elements[0]
	d.elements = d.elements[1:] // Remove the front element
	return front
}

// PopBack removes and returns the back element of the deque
func (d *Deque) PopBack() (int) {
	if len(d.elements) == 0 {
		return 0// Deque is empty
	}
	back := d.elements[len(d.elements)-1]
	d.elements = d.elements[:len(d.elements)-1] // Remove the back element
	return back
}

// PeekFront returns the front element without removing it
func (d *Deque) PeekFront() (int) {
	if len(d.elements) == 0 {
		return 0 // Deque is empty
	}
	return d.elements[0]
}

// PeekBack returns the back element without removing it
func (d *Deque) PeekBack() (int) {
	if len(d.elements) == 0 {
		return 0 // Deque is empty
	}
	return d.elements[len(d.elements)-1]
}

// Size returns the number of elements in the deque
func (d *Deque) Size() int {
	return len(d.elements)
}

/* O(n*k) solution
func maxSlidingWindow(nums []int, k int) []int {

	n := len(nums)
	currMax := math.MinInt
	currMaxIndex := -1
	start, end := 0, k-1
	result := make([]int, 0)
	for start <= n-k && end < n {
		if currMaxIndex >= start && currMaxIndex <= end {
			if nums[end] > currMax {
				currMax = nums[end]
				currMaxIndex = end
				start++
				end++
				result = append(result, currMax)
				continue
			} else {
				result = append(result, currMax)
				start++
				end++
				continue
			}
		}
        currMax = math.MinInt
		for j := start; j <= end; j++ {
			if nums[j] > currMax {
				currMax = nums[j]
				currMaxIndex = j
			}
		}
		result = append(result, currMax)
		start++
		end++
	}
	return result
}*/