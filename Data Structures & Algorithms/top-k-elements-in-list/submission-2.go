
type Data struct {
	Key int
	Value int
}


func topKFrequent(nums []int, k int) []int {

	dataSet := map[int]int{}
	for  _ , i := range nums {
		dataSet[i]++
	}
	temp := []Data{}
	for key , val := range dataSet {
		temp = append(temp , Data{Key : key , Value :val })
	}

	sort.Slice(temp, func(a, b int) bool {
		return temp[a].Value > temp[b].Value
	})
	resp := []int{}
	for _ , d :=  range temp[:k] {
		resp = append(resp , d.Key)
	}
	
	return resp

}

