func isPalindrome(s string) bool {
    s = strings.ToLower(s)
    rArr := []rune(s)
    left, right := 0, len(rArr)-1

    for left <= right {
        if !isAlphaNumeric(rArr[left]) {
            left++
            continue
        } 
        if !isAlphaNumeric(rArr[right]) {
            right--
            continue
        }

        if rArr[left] != rArr[right] {
            return false
        }
        left++
        right--
    }
    return true
}

func isAlphaNumeric(r rune) bool {
    if (r >= 'a' && r <= 'z') ||
        (r >= '0' && r <= '9') {
            return true
        }
    return false
}