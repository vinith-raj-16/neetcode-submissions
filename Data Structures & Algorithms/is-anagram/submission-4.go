func isAnagram(s string, t string) bool {

	sData := map[byte]int{}
	if len(s) != len(t) {
		return false
	}

	for i  :=  0 ;i < len(s) ; i++  {
		sData[s[i]]++
		sData[t[i]]--
	}
	
	for _ , value := range sData {
		if value != 0{
			return false
		}
	}
	return true
}
