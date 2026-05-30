

func isPalindrome(s string) bool {
    l, r := 0, len(s)-1
    for l<=r {
        charLeft := unicode.ToLower(rune(s[l]))
        charRight := unicode.ToLower(rune(s[r]))

        if !((charLeft >= 'a' && charLeft <='z') || (charLeft >= '0' && charLeft <= '9')) {
            l++
            continue
        }
       if !((charRight >= 'a' && charRight <='z') || (charRight >= '0' && charRight <= '9')) {
            r--
            continue
        }
        if unicode.ToLower(rune(s[l]))!=unicode.ToLower(rune(s[r])){
            return false
        }
        l++
        r--
    }
    return true
}
