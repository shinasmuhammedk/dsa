package maps

func TwoSum(arr []int, target int)[]int{
    seen := make(map[int]int)
    
    for i,v := range arr{
        complement := target - v
        
        if index, ok := seen[complement];ok{
            return []int{index,i}
        }
        
        seen[v] = i
    }
    
    return nil
    
}