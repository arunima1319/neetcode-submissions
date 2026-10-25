func twoSum(nums []int, target int) []int {
    
	complementMap := make(map[int]int)

	for i, num := range nums{
		if _, ok := complementMap[num]; ok{
			return []int{complementMap[num], i}
		}
		complementMap[target - num] = i
	
	}

	return []int{0, 0}
}