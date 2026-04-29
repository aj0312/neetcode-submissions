func characterReplacement(s string, k int) int {
    left, right := 0, 0
    maxLen := 0
    charCount := [26]int{}

    for left <= right && right < len(s) {
        idx := s[right] - 'A'
        charCount[idx]++ //[3,2]
        maxCount := 0

        for _, count := range charCount {
            maxCount = max(maxCount, count) //3
        }

        for (right - left + 1 - maxCount) > k {
            charCount[s[left]-'A']--
            left++
        }
        
        maxLen = max(maxLen, right - left + 1) //4
        right++ //4
    }
    return maxLen
}
