// 169. Majority Element 
//
//


func majorityElement(nums []int) int {
    sort.Ints(nums)
    
    count := 1 
    
    
    for i := 0; i < len(nums)-1; i++ {
        if nums[i] == nums[i+1] {
            count++
        } else {
            count = 1 
        }
        
        if count > len(nums)/2 {
            return nums[i]
        }
    }
    
    
    return nums[0]
}
