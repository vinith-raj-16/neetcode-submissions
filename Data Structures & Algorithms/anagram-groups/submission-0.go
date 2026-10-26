func groupAnagrams(strs []string) [][]string {

	resp := [][]string{}

	dataSet := map[string][]string{}

	for _ , str := range strs {
		s := Sort(str)
		dataSet[s]=append(dataSet[s] , str)
	}

	for _ , val := range dataSet {
		resp = append(resp , val)
	}



	return resp

}


func Sort(s string) string {
    r := []rune(s) // convert to runes
    sort.Slice(r, func(i, j int) bool {
        return r[i] < r[j]
    })
	return string(r)
}
