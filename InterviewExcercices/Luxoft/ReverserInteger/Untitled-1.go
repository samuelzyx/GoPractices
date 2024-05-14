/*
	Escenarios
	123
	321

	250
	052

	-171
	-171

	strconv
*/

func testInterview(param1 int) int { //123 ** -529

	//convert param1 into string
	paramS1 := fmt.PrintF("%d", param1) //"123" ** "-529"
	isNegative := param1 < 0            //boolean //false ** true
	var result string                   //"" ** ""

	if isNegative { //false ** true
		paramS1 = paramS1[1:len(paramS1)] //not remove symbol - ** "529"
	}

	//loop in order to reverse ther runes position
	for i := len(paramS1) - 1; i >= 0; i-- {
		//store/detect if have a negative simbol
		//made the reverse store in result variable with the simbol negative if exist
		result = result + paramS1[i]
		//3 ** 9
		//23 ** 92
		//123 ** 925
	}

	if isNegative { //false ** true
		result = "-" + result //do nothing ** "-925"
	}

	//it will return the result variable
	return strconv.Atoi(result) //123 ** -925
}