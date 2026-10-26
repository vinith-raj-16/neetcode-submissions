func topKFrequent(nums []int, k int) []int {

	dataSet := map[int]int{}

	for  _ , i := range nums {
		dataSet[i]++
	}

	temp := []int{}
	for _ , val := range dataSet {
		temp = append(temp , val)
	}
	sort.Ints(temp)
	val := temp[len(temp)-k:]
	resp := []int{}
	fmt.Println(val)
	for key , values := range dataSet {
		
		if Find(val , values) {
			resp = append(resp , key)

		}

	}
	return resp

}

func Find(i []int , k int)bool{
	for  _,j :=  range i {
		if j == k {
			return true
		}
		
	}
	return false
}
