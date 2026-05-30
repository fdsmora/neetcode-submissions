

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
        if delimit {
            sb.WriteString(delimiter)
        }
        //encodedWord := encode(w)
        sb.WriteString(w)
        delimit = true
    }
    return sb.String()
}

func (s *Solution) Decode(str string) []string {
    if str == "_" {
        return []string{}
    }
    return strings.Split(str, "|")
}
