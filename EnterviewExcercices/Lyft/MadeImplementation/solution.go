package main

import (
	"fmt"
	"math"
)

// These numbers are for testing only and may be changed by the intervieer.
// Do not use them in your solution
const MaxResults = 103
const PageSize = 10

// External API - should not be modified for solution

type Page struct {
	nextPage int
	results  []int
}

// FetchPage returns the next Page of results. If page == 0, starts from
// the begining. Otherwise, fetches the next records after the last page.
func FetchPage(page int) Page {
	if page*PageSize > MaxResults {
		return Page{
			nextPage: -1,
			results:  []int{},
		}
	}

	start := page * PageSize
	end := int(math.Min(float64(MaxResults), float64(page+1)*PageSize))
	results := make([]int, end-start)
	for i := range results {
		results[i] = start + i
	}
	return Page{
		nextPage: page + 1,
		results:  results,
	}
}

type Fetcher struct {
	fetchPage               func(page int) Page
	lastResultFetchReturned int
}

func NewFetcher() *Fetcher {
	return &Fetcher{
		fetchPage:               FetchPage,
		lastResultFetchReturned: 0,
	}
}

// TO IMPLEMENT: Fetch returns numResults amount of results from FetchPage.
func (f *Fetcher) Fetch(numResults int) []int {
	var results []int
	for i := 0; i < numResults; i++ {
		results = append(results, f.lastResultFetchReturned)
		f.lastResultFetchReturned++
	}
	return results
}

func TestCase(actual []int, testCaseNum, fetchNum, expectedStart, expectedEnd int) {
	expected := make([]int, expectedEnd-expectedStart)
	for i := range expected {
		expected[i] = expectedStart
		expectedStart += 1
	}
	if len(actual) != len(expected) {
		fmt.Printf("TEST CASE %d: Fetch %d, Start: %d End:%d, FAILED: \nexpected %v \ngot %v\n\n\n", testCaseNum, fetchNum, expectedStart, expectedEnd, expected, actual)
		return
	}

	for i, v := range actual {
		if v != expected[i] {
			fmt.Printf("TEST CASE %d: Fetch %d, Start: %d End:%d, FAILED: \nexpected %v \ngot %v\n\n\n", testCaseNum, fetchNum, expectedStart, expectedEnd, expected, actual)
			return
		}
	}

	fmt.Printf("TEST CASE %d: Fetch %d, Start: %d End:%d, SUCCESS: \nexpected %v \ngot %v\n\n\n", testCaseNum, fetchNum, expectedStart, expectedEnd, expected, actual)
}

func main() {
	fetcher := NewFetcher()
	TestCase(fetcher.Fetch(5), 1, 5, 0, 5)
	TestCase(fetcher.Fetch(2), 2, 2, 5, 7)
	TestCase(fetcher.Fetch(7), 3, 7, 7, 14)
	//The next test cases are not passing why?
	TestCase(fetcher.Fetch(103), 4, 103, 14, 103)
	/*
		//Confirm result
		TEST CASE 4: Fetch 103, Start: 103 End:103, FAILED:
		expected [14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47 48 49 50 51 52 53 54 55 56 57 58 59 60 61 62 63 64 65 66 67 68 69 70 71 72 73 74 75 76 77 78 79 80 81 82 83 84 85 86 87 88 89 90 91 92 93 94 95 96 97 98 99 100 101 102]
		got [14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47 48 49 50 51 52 53 54 55 56 57 58 59 60 61 62 63 64 65 66 67 68 69 70 71 72 73 74 75 76 77 78 79 80 81 82 83 84 85 86 87 88 89 90 91 92 93 94 95 96 97 98 99 100 101 102 103 104 105 106 107 108 109 110 111 112 113 114 115 116]
	*/
	TestCase(fetcher.Fetch(10), 5, 10, 0, 0)
	/*
		//Confirm result
		TEST CASE 5: Fetch 10, Start: 0 End:0, FAILED:
		expected []
		got [117 118 119 120 121 122 123 124 125 126]
	*/
	fetcher = NewFetcher()
	TestCase(fetcher.Fetch(200), 6, 200, 0, 103)
	/*
		//Confirm result
		TEST CASE 6: Fetch 200, Start: 103 End:103, FAILED:
		expected [0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47 48 49 50 51 52 53 54 55 56 57 58 59 60 61 62 63 64 65 66 67 68 69 70 71 72 73 74 75 76 77 78 79 80 81 82 83 84 85 86 87 88 89 90 91 92 93 94 95 96 97 98 99 100 101 102]
		got [0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47 48 49 50 51 52 53 54 55 56 57 58 59 60 61 62 63 64 65 66 67 68 69 70 71 72 73 74 75 76 77 78 79 80 81 82 83 84 85 86 87 88 89 90 91 92 93 94 95 96 97 98 99 100 101 102 103 104 105 106 107 108 109 110 111 112 113 114 115 116 117 118 119 120 121 122 123 124 125 126 127 128 129 130 131 132 133 134 135 136 137 138 139 140 141 142 143 144 145 146 147 148 149 150 151 152 153 154 155 156 157 158 159 160 161 162 163 164 165 166 167 168 169 170 171 172 173 174 175 176 177 178 179 180 181 182 183 184 185 186 187 188 189 190 191 192 193 194 195 196 197 198 199]
	*/
}
