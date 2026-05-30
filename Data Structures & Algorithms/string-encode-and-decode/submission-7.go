

type Solution struct{}

func (s *Solution) Encode(strs []string) string {
    if len(strs)==0{
        return "_"
    }
    var sb strings.Builder

    var delimit bool
    var delimiter = "|"
    for i:=0; i<len(strs); i++ {
        w := strs[i]
        encodW := encode(w)
        if delimit {
            sb.WriteString(delimiter)
        }
        //encodedWord := encode(w)
        sb.WriteString(encodW)
        delimit = true
    }
    return sb.String()
}

func encode(s string) string {
    var sb strings.Builder

    for _, c := range s {
        if c == '/' || c == '|'{
            sb.WriteRune('/')
        }
        sb.WriteRune(c)
    }
    return sb.String()
}

func (s *Solution) Decode(str string) []string {
    if str == "_" {
        return []string{}
    }

    var sb strings.Builder
    var res []string
    i := 0
    for i<len(str) {
        ch := str[i]
        if ch == '/' {
            i++
            ch = str[i]
            sb.WriteByte(ch)
        } else if ch == '|' {
            res = append(res, sb.String())
            sb.Reset()
        } else {
            sb.WriteByte(ch)
        } 
        i++          
    }
    res = append(res, sb.String())
    return res
}


