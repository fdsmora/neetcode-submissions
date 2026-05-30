func isValid(s string) bool {
    if len(s)%2 == 1{
        return false
    }
    st := []rune{}
    pars := map[rune]rune{'(':')', '{':'}', '[':']'}

    for _, c := range s {
        if v, ok := pars[c]; ok {
            st = append(st, v)
        }else if len(st)>0{
            closing := st[len(st)-1]
            if closing != c {
                return false
            }
            st = st[:len(st)-1]
        }else{
            return false
        }
    }
    return len(st)==0
}