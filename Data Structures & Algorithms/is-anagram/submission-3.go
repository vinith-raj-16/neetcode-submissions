func isAnagram(s string, t string) bool {

	sData := map[byte]int{}
	tData := map[byte]int{}
	if len(s) != len(t) {
		return false
	}

	for i  :=  0 ;i < len(s) ; i++  {
		sData[s[i]]++
		tData[t[i]]++
	}
	
	for key , value := range sData {
		if tData[key]!=value{
			return false
		}
	}
	return true
}
