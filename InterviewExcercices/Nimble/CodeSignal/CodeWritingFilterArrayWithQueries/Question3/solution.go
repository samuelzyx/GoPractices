package main

import (
	"fmt"
	"sort"
	"strings"
)

func solution(data [][]string, query []string) []string {
	result := []string{}
	for i := 1; i < len(data); i++ {
		fmt.Printf("\ndata[i][0]:%v\n", data[i][0])
		//fmt.Printf("data[i][1]:%v\n", data[i][1])
		name := data[i][0]
		labels := data[i][1]
		allGood := true

		for q := 0; q < len(query); q++ {
			query := query[q]
			female := strings.Contains(query, "female")
			male := strings.Contains(query, "male")
			not := strings.Contains(query, "not")

			if not {
				fmt.Printf("not:%v\n", not)
				notQuery := strings.Split(query, "not ")
				query = notQuery[1]
				fmt.Printf("query:%v\n", query)
				if strings.Contains(query, "female") && strings.Contains(labels, "female") {
					allGood = false
				} else if strings.Contains(query, "male") && strings.Contains(labels, ", male") && !female {
					allGood = false
				} else if strings.Contains(query, "sanctioned") && strings.Contains(labels, "sanctioned") {
					allGood = false
				} else if strings.Contains(query, "takes_new_patients") && strings.Contains(labels, "takes_new_patients") {
					allGood = false
				}
			} else if female || male {
				if female && !strings.Contains(labels, "female") {
					allGood = false
				} else if male && !strings.Contains(labels, ", male") && !female {
					allGood = false
				}
			} else if !strings.Contains(labels, query) {
				allGood = false
			}
			fmt.Printf("Name %v allGood:%v\n", name, allGood)
		}
		if allGood {
			result = append(result, name)
		}
	}
	sort.Slice(result, func(i, j int) bool { return result[i] < result[j] })

	fmt.Printf("\nresult returned:%v\n", result)
	return result
}

func main() {
	doctors := [][]string{
		{"name", "labels"},
		{"Matt", "board certified, primary care, male, takes_new_patients"},
		{"Belda", "board certified, internal medicine, female"},
		{"Wyatt", "primary care, male, takes_new_patients"},
		{"Emma", "board certified, oncology"},
		{"Aaron", "sanctioned, primary care"},
		{"Josh", "board certified, internal medicine, takes_new_patients"},
		{"Adrien", "oncology, board certified, takes_new_patients"},
		{"Andy", "internal medicine, male, sanctioned"},
	}
	queryTest1 := []string{
		"primary care",
		"not sanctioned",
	}
	queryTest2 := []string{
		"male",
	}
	queryTest3 := []string{
		"not male",
		"board certified",
	}
	queryTest4 := []string{
		"not takes_new_patients",
		"female",
	}
	result1 := solution(doctors, queryTest1)
	result2 := solution(doctors, queryTest2)
	result3 := solution(doctors, queryTest3)
	result4 := solution(doctors, queryTest4)

	answer1 := []string{
		"Matt",
		"Wyatt",
	}
	answer2 := []string{
		"Andy",
		"Matt",
		"Wyatt",
	}
	answer3 := []string{
		"Adrien",
		"Belda",
		"Emma",
		"Josh",
	}
	answer4 := []string{
		"Belda",
	}

	TestAnswers(&answer1, &result1, 1)
	TestAnswers(&answer2, &result2, 2)
	TestAnswers(&answer3, &result3, 3)
	TestAnswers(&answer4, &result4, 4)
}

func TestAnswers(answer *[]string, result *[]string, numAns int) {
	allGood := true

	if len(*answer) != len(*result) {
		allGood = false
	} else {
		for i := 0; i < len(*answer); i++ {
			if (*answer)[i] != (*result)[i] {
				allGood = false
			}
		}
	}

	if allGood {
		fmt.Printf("Answer %d correct\n", numAns)
	} else {
		fmt.Printf("Answer %d incorrect\n", numAns)
	}
}
