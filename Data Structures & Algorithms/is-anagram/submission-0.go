func isAnagram(s string, t string) bool {
	
	length := len(s)
	if len(t) != length{
		return false
	}
	letterMapS := make(map[rune]int)
	letterMapT := make(map[rune]int)
	for _, letter := range s{
		letterMapS[letter] += 1 
	}
	for _, letter := range t{
		letterMapT[letter] +=1
	}
	for letter := range letterMapS{
		if letterMapS[letter] != letterMapT[letter]{
			return false
		}
	}
	return true

}
