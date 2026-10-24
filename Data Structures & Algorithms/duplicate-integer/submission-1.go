func hasDuplicate(nums []int) bool {
	intMap := make(map[int]struct{})
    for _, num := range nums{
		_, ok := intMap[num]
		if ok{
			return true
		}
		intMap[num] = struct{}{}
	}
	return false

}
