func checkInclusion(s1 string, s2 string) bool {
    if len(s2) < len(s1) {
        return false
    }

    s1Count := [26]int{}
    s2Count := [26]int{}

    for _, r := range s1 {
        i := r - 'a'
        s1Count[i]++
    }

    left, right := 0, 0

    for i := 0; i < len(s1); i++ {
        rIdx := s2[i] - 'a'
        s2Count[rIdx]++
        right++
    }

    matches := 0

    for i := 0; i < 26; i++ {
        if s1Count[i] == s2Count[i] {
            matches++
        }
    }

    for right < len(s2) && left < right {
        if matches == 26 {
            return true
        }
        
        
        // Add right character
        rightCharIdx := s2[right] - 'a'
        if s2Count[rightCharIdx] == s1Count[rightCharIdx] {
            matches-- // about to break this match
        }
        s2Count[rightCharIdx]++
        if s2Count[rightCharIdx] == s1Count[rightCharIdx] {
            matches++ // just created a match
        }

        // Remove left character
        leftCharIdx := s2[left] - 'a'
        if s2Count[leftCharIdx] == s1Count[leftCharIdx] {
            matches-- // about to break this match
        }
        s2Count[leftCharIdx]--
        if s2Count[leftCharIdx] == s1Count[leftCharIdx] {
            matches++ // just created a match
        }
        
        left++
        right++
    }

    return matches == 26
}