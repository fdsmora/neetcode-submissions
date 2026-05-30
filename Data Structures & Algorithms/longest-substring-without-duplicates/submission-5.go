func lengthOfLongestSubstring(s string) int {
    var maxLen, currLen, front, back int
    chars := map[byte]int{}

    for front = 0; front< len(s); front++ {
        if idx, ok := chars[s[front]]; ok {
            back = max(back, idx+1)
        }
        chars[s[front]]=front
        currLen = front - back + 1
        if currLen > maxLen {
            maxLen = currLen
        }
    }
    return maxLen
}

/*
0 1 2 3 4 5 6 7 8 9 10
a b c d e f g a z y x
                      f
  b
chars:   
a 7  
b 1
c 2
d 3
e 4
f 5
g 6
z 8
y 9 
x 10

cL = 11-1 = 10
mL = 10
idx = 0
*/
