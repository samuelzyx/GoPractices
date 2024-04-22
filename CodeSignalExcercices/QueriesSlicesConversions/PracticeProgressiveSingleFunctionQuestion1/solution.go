func solution(queries [][]string) []string {
    var result []string
    var stored []int
    
	for i := 0; i < len(queries); i++ {
        parse, _ := strconv.Atoi(queries[i][1])
        if queries[i][0]=="ADD"{
            result = append(result,"")
            stored = append(stored, parse)
        } else if queries[i][0]=="EXISTS"{
            exist:=false
            for e:=0;e<len(stored);e++{
                if(stored[e] == parse){
                    exist = true
                }        
            }
            result = append(result,strconv.FormatBool(exist))                       
        }
    }
    return result
}
