func hasDuplicate(nums []int) bool {

  data := map[int]bool{}

    for _,num := range nums{
        if ok := data[num] ; !ok{
            data[num]=true
        }else{
           return true
        }
    }

    return false  
    
}
