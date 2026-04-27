func lengthOfLongestSubstring(s string) int {
    left, right := 0, 0
    maxLen := 0
    asciiMap := [256]bool{}

    for left <= right && right < len(s) {
        if !asciiMap[s[right]] {
            asciiMap[s[right]] = true
            right++
        } else {
            for left <= right && asciiMap[s[right]] {
                asciiMap[s[left]] = false
                left++
            }
        }
        subStrLen := right - left
        maxLen = max(maxLen, subStrLen)
    }

    return maxLen
}
