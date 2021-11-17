package my_append

func MyAppend(slice []int, newItem int) []int {
	sliceLength := len(slice) //切片长度
	sliceCap := cap(slice)    //切片容量
	if sliceLength >= sliceCap {
		//当前容量不够用，扩容
		var newSliceCap int
		if sliceLength*2 >= sliceLength+1 {
			newSliceCap = sliceLength * 2
		} else {
			newSliceCap = sliceLength + 1
		}
		newSlice := make([]int, sliceLength+1, newSliceCap)
		copy(newSlice, slice)
		slice = newSlice
	} else {
		//容量够，在原基础上增加切片长度
		slice = slice[:sliceLength+1]
	}
	slice[sliceLength] = newItem
	return slice
}
