func isAnagram(s string, t string) bool {
    charCount := make([]int, 26)
    for _, c := range s {
        charCount[c-'a']+=1
    }
    for _, c := range t {
        charCount[c-'a']-=1
        if charCount[c-'a']<0{
            return false
        }
    }
    for _, e := range charCount {
        if e != 0{
            return false
        }
    }
    return true
}
