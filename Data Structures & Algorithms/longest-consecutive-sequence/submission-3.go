func longestConsecutive(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    numSet := map[int]struct{}{}
    count, maxCount := 1, 1

    for _, num := range nums {
        numSet[num] = struct{}{}
    }

    for _, num := range nums {
        nextNum := num+1
        count = 1
        for {
            prevNum := num-1
            if _, exist := numSet[nextNum]; !exist {
                break
            }
            if _, exist := numSet[prevNum]; exist {
                break
            }
            count++
            nextNum++
        }
        maxCount = max(maxCount, count)
    }
    return maxCount
}
