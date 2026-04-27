import "slices"

func threeSum(nums []int) [][]int {
    slices.Sort(nums)
    output := [][]int{}
    // [-4,-1,-1,0,1,2]
    idx := 0

    for idx < len(nums) {
        if idx > 0 && nums[idx] == nums[idx-1] {
            idx++
            continue
        }
        num := nums[idx]
        if num > 0 {
            break
        }
        left, right := idx+1, len(nums)-1
        target := 0 - num
        for left < right {
            sum := nums[left] + nums[right]
            if sum == target {
                group := []int{num, nums[left], nums[right]}
                output = append(output, group)
                left++
                for nums[left] == nums[left-1] && left < right {
                    left++
                }
                continue
            }
            if sum > target {
                right--
                continue
            }
            left++
        }

        idx++
    }

    return output
}
