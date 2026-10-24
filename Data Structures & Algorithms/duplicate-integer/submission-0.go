func hasDuplicate(nums []int) bool {
	intMap := make(map[int]int)
    for _, num := range nums{
		_, ok := intMap[num]
		if ok{
			return true
		}
		intMap[num] = 1
	}
	return false

}
