import "sort"

func checkExist(arr *[]int, val int) bool {
	exist := false
	for e := 0; e < len(*arr); e++ {
		if (*arr)[e] == val {
			exist = true
		}
	}
	return exist
}

func deleteElement(arr *[]int, val int) bool {
	index := -1
	for e := 0; e < len(*arr); e++ {
		if (*arr)[e] == val {
			index = e
		}
	}
	founded := index != -1
	if founded {
		*arr = append((*arr)[:index], (*arr)[index+1:]...)
	}
	return founded
}

func getNext(arr *[]int, val int) string {
	next := -1
	sort.Slice((*arr), func(i, j int) bool { return (*arr)[i] < (*arr)[j] })
	for e := 0; e < len(*arr); e++ {
		if (*arr)[e] > val {
			next = (*arr)[e]
			break
		}
	}
	if next == -1 {
		return ""
	} else {
		return strconv.Itoa(next)
	}
}

func solution(queries [][]string) []string {
	var result []string
	var stored []int

	for i := 0; i < len(queries); i++ {
		parse, _ := strconv.Atoi(queries[i][1])
		if queries[i][0] == "ADD" {
			result = append(result, "")
			stored = append(stored, parse)
		} else if queries[i][0] == "EXISTS" {
			checkExistConv := strconv.FormatBool(checkExist(&stored, parse))
			result = append(result, checkExistConv)
		} else if queries[i][0] == "REMOVE" {
			deleteElementConv := strconv.FormatBool(deleteElement(&stored, parse))
			result = append(result, deleteElementConv)
		} else if queries[i][0] == "GET_NEXT" {
			result = append(result, getNext(&stored, parse))
		}
	}
	return result
}
