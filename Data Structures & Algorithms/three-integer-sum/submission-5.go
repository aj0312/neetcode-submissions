import "slices"

func threeSum(nums []int) [][]int {
    slices.Sort(nums)
    output := [][]int{}

    idx := 0

    for idx < len(nums) {
        if idx > 0 && nums[idx] == nums[idx-1] {
            idx++
            continue
        }
        left, right := idx+1, len(nums)-1

        for left < right {
            sum := nums[idx] + nums[left] + nums[right]
            if sum == 0 {
                output = append(output, []int{nums[idx], nums[left], nums[right]})
                left++
                for nums[left] == nums[left-1] && left < right {
                    left++
                }
                continue
            }
            if sum > 0 {
                right--
                continue
            }
            left++
        }
        idx++
    }
    return output
}
