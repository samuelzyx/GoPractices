package main

import (
	"fmt"
	"sort"
	"strings"
)

func solution(data [][]string) []string {
	result := []string{}
	for i := 1; i < len(data); i++ {
		//fmt.Printf("data[i][0]:%v\n", data[i][0])
		//fmt.Printf("data[i][1]:%v\n", data[i][1])
		name := data[i][0]
		labels := data[i][1]
		if strings.Contains(labels, "internal medicine") && strings.Contains(labels, "board certified") {
			result = append(result, name)
		}
	}
	sort.Slice(result, func(i, j int) bool { return result[i] < result[j] })

	//fmt.Printf("result returned:%v\n", result)
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
	result := solution(doctors)

	answer := []string{
		"Belda",
		"Josh",
	}

	allGood := true

	if len(answer) != len(result) {
		allGood = false
	} else {
		for i := 0; i < len(answer); i++ {
			if answer[i] != result[i] {
				allGood = false
			}
		}
	}

	if allGood {
		fmt.Print("Answer correct")
	} else {
		fmt.Print("Answer incorrect")
	}
}
